/*
Package payments implements Payments API.

Documentation: https://vk.com/dev/payments

With Payments API applications can sell virtual products to users or provide
them with services using VK internal currency — votes.

Processing payment notifications

Documentation: https://vk.com/dev/payments_callbacks

Notifications are sent by the payment system server to the callback URL
indicated in app settings, according to HTTP or HTTPS protocol, depending on
the protocol specified in the callback URL with the POST method in UTF-8
encoding.

To avoid the possible forgery, notifications are signed with a secret key
known only to the app and payment system owner.

The app developer must process notifications and return either the
processing result if successful or an error description if unsuccessful.
The response must be sent in 10 seconds, or else the connection will be lost
and the attempt to send the notification will happen once more after some
time.

	cb := payments.NewCallback(secret)

	cb.OnGetItem(func(e payments.GetItemRequest) (*payments.GetItemResponse, *payments.Error) {
		// ...
	})

	cb.OnOrderStatusChange(func(e payments.OrderStatusChangeRequest) (*payments.OrderStatusChangeResponse, *payments.Error) {
		// ...
	})

	cb.OnGetSubscription(func(e payments.GetSubscriptionRequest) (*payments.GetSubscriptionResponse, *payments.Error) {
		// ...
	})

	cb.OnSubscriptionStatusChange(func(e payments.SubscriptionStatusChangeRequest) (*payments.SubscriptionStatusChangeResponse, *payments.Error) {
		// ...
	})

	http.HandleFunc("/payments", cb.HandleFunc)

	http.ListenAndServe(":8080", nil)

Test Mode

Documentation: https://vk.com/dev/payments_testmode

Test mode allows for testing an app's functionality for buying goods and
transferring votes without a real transfer of votes.

While your app has not been checked by the VK administration, test mode is
necessary to use. Payments are completed in "combat" mode immediately after
the app verification procedure has been completed.

To make payments in test mode, add all users who will make test payments to
the special list in app settings (the "Payments tester" field).

In test mode, the suffix '_test' is added to the NotificationType parameter.

	cb.OnGetItemTest(func(e payments.GetItemRequest) (*payments.GetItemResponse, *payments.Error) {
		// ...
	})

	cb.OnOrderStatusChangeTest(func(e payments.OrderStatusChangeRequest) (*payments.OrderStatusChangeResponse, *payments.Error) {
		// ...
	})

	cb.OnGetSubscriptionTest(func(e payments.GetSubscriptionRequest) (*payments.GetSubscriptionResponse, *payments.Error) {
		// ...
	})

	cb.OnSubscriptionStatusChangeTest(func(e payments.SubscriptionStatusChangeRequest) (*payments.SubscriptionStatusChangeResponse, *payments.Error) {
		// ...
	})

The order IDs in test mode (i.e. the notification parameter OrderID) can
overlap with the order IDs in operating mode.
*/
package payments // import "github.com/SevereCloud/vksdk/payments"

import (
	"crypto/md5" // nolint: gosec
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/schema"
)

// User language as language_country.
const (
	LangRussian    = "ru_RU" // Russian
	LangUkrainian  = "uk_UA" // Ukrainian
	LangBelarusian = "be_BY" // Belarusian
	LangEnglish    = "en_US" // English
)

// Status type.
type Status string

// Subscription status.
const (
	// Subscription is ready for payments. An order for the user
	// needs to be processed within the application. If the response is
	// successful, the payment system credits votes to the application
	// balance. If the response is a failure, the order is canceled.
	Chargeable Status = "chargeable"

	// Subscription is active;
	Active Status = "active"

	// Subscription is canceled.
	Cancelled Status = "cancelled"
)

// Reason type.
type Reason string

// Reason for cancellation.
const (
	// Subscription canceled by the user.
	UserDecision Reason = "user_decision"

	// Subscription canceled by the application (orders.cancelSubscription).
	AppDecision Reason = "app_decision"

	// Subscription canceled due to a failed payment.
	PaymentFail Reason = "payment_fail"

	// Subscription canceled for a different reason.
	Unknown Reason = "unknown"
)

// NotificationType type of notification.
type NotificationType string

// List type of notification.
const (
	// For acquiring product information.
	//
	// https://vk.com/dev/payments_getitem
	GetItem NotificationType = "get_item"

	// For changing the order’s status.
	//
	// https://vk.com/dev/payments_status
	OrderStatusChange NotificationType = "order_status_change"

	// For receiving subscription information.
	//
	// https://vk.com/dev/payments_getsubscription
	GetSubscription NotificationType = "get_subscription"

	// For changing a subscription’s status.
	//
	// https://vk.com/dev/payments_subscriptionstatus
	SubscriptionStatusChange NotificationType = "subscription_status_change"
)

// Test return NotificationType for test mode.
//
// In test mode, the suffix '_test' is added to the notification_type parameter.
//
// See https://vk.com/dev/payments_testmode
func (t NotificationType) Test() NotificationType {
	return t + "_test"
}

// Notification structures.
//
// The selection of fields depends on the type of notification. However all
// types of notifications contain the following fields.
type Notification struct {
	// The type of notification.
	Type NotificationType `schema:"notification_type,required"`

	// The app’s ID.
	AppID int `schema:"app_id,required"`

	// The ID of the user who made the order.
	UserID int `schema:"user_id,required"`

	// The notification signature.
	Sig string `schema:"sig,required"`
}

// Numerical code error.
//
// Errors 100-999 are specified by the app. Such errors always include an
// error description.
//
// See https://vk.com/dev/payments_errors
const (
	// Common error.
	//
	// Severity: true/false.
	CommonError = 1

	// Temporary database error.
	//
	// Severity: false.
	TemporaryDatabaseError = 2

	// The calculated and sent signatures do not match.
	//
	// Severity: true.
	BadSignatures = 10

	// Request parameters do not comply with the specification.
	// No required fields in the request.
	// Other request integrity errors.
	//
	// Severity: true.
	BadRequest = 11

	// Product does not exist.
	//
	// Severity: true.
	ProductNotExist = 20

	// Product is out of stock.
	//
	// Severity: true.
	ProductOutOfStock = 21

	// User does not exist.
	//
	// Severity: true.
	UserNotExist = 22
)

// Error struct.
//
// When critical errors occur, the order cancels and the app event onOrderFail
// is sent.
// If the error is temporary, a notification will be resent after some time
// and the user will wait for the process is completed.
//
// See https://vk.com/dev/payments_errors
type Error struct {
	// Numerical code error
	Code int `json:"error_code"`

	// Error description in easy-to-read format
	// (required for error_code >= 100).
	Msg string `json:"error_msg,omitempty"`

	// Error severity. Possible values are:
	//
	// true — if a notification with identical parameters is passed
	// repeatedly, the same error will occur. For example, the indicated
	// product does not exist. The notification will not be resent as the user
	// will receive an error.
	//
	// false — if the error is temporary, a notification may be processed
	// later. For example, a temporary error in posting to the database.
	// The notification will be sent after some time and the user will wait
	// for the response.
	Critical bool `json:"critical"`
}

// Error return error message.
func (e Error) Error() string {
	if e.Msg != "" {
		return "payments: " + e.Msg
	}

	return "payments: error with code " + strconv.Itoa(e.Code)
}

// Callback struct.
type Callback struct {
	Secret string

	getItem                      func(e GetItemRequest) (*GetItemResponse, *Error)
	getItemTest                  func(e GetItemRequest) (*GetItemResponse, *Error)
	orderStatusChange            func(e OrderStatusChangeRequest) (*OrderStatusChangeResponse, *Error)
	orderStatusChangeTest        func(e OrderStatusChangeRequest) (*OrderStatusChangeResponse, *Error)
	getSubscription              func(e GetSubscriptionRequest) (*GetSubscriptionResponse, *Error)
	getSubscriptionTest          func(e GetSubscriptionRequest) (*GetSubscriptionResponse, *Error)
	subscriptionStatusChange     func(e SubscriptionStatusChangeRequest) (*SubscriptionStatusChangeResponse, *Error)
	subscriptionStatusChangeTest func(e SubscriptionStatusChangeRequest) (*SubscriptionStatusChangeResponse, *Error)
}

// NewCallback return new Callback.
func NewCallback(secret string) *Callback {
	return &Callback{
		Secret: secret,
	}
}

// GetItemRequest notification parameters.
type GetItemRequest struct {
	Notification

	// The recipient’s ID.
	ReceiverID int `schema:"receiver_id,required"`

	// The ID of the order.
	OrderID int `schema:"order_id,required"`

	// User language as language_country.
	Lang string `schema:"lang,required"`

	// Product name passed to the purchase dialog box.
	Item string `schema:"item,required"`
}

// GetItemResponse reply parameters.
type GetItemResponse struct {
	// product name, max 48 characters
	//
	// Required parameter
	Title string `json:"title"`

	// URL of product image on the developer's server
	PhotoURL string `json:"photo_url,omitempty"`

	// product price, in votes
	//
	// Required parameter
	Price int `json:"price"`

	// the size of the discount on the goods in the votes. Should be within
	// 1 to 1000 votes and less than the price of goods.
	Discount int `json:"discount,omitempty"`

	// product ID in the application
	ItemID string `json:"item_id,omitempty"`

	// allows product caching for {Expiration} seconds. Allowed
	// range from 600 to 604800 seconds.
	//
	// Warning! In the absence of the parameter is possible to cache the goods
	// for 3600 seconds with a large number of consecutive identical
	// responses. For cancellation of caching it is necessary to pass 0 as
	// value of parameter.
	Expiration int `json:"expiration,omitempty"`
}

// OnGetItem notification is sent when the product purchase dialog box is
// opened in the application. Then, obtained information about the product
// is shown in the purchase dialog box.
//
// With item specifying the product name received in the notification, the
// developer shall return actual information about such product. When there
// is no product available, reply with Error 20 "Product does not exist".
//
// 		&payments.Error{
// 			Code:     payments.ProductNotExist,
// 			Msg:      "Product does not exist",
// 			Critical: true,
// 		}
//
// Please note! As item is passed at the client side, user can change this
// parameters.
//
// Please note! We recommend you to use product caching for all products,
// this will decrease the number of calls to the application server and users
// will not wait when information about them is retrieved.
//
// In case information about the product with item ID was cached for
// expiration period, following requests for such product will not be run
// within the given time.
//
// See https://vk.com/dev/payments_getitem
func (cb *Callback) OnGetItem(f func(e GetItemRequest) (*GetItemResponse, *Error)) {
	cb.getItem = f
}

// OnGetItemTest OnGetItem for test.
func (cb *Callback) OnGetItemTest(f func(e GetItemRequest) (*GetItemResponse, *Error)) {
	cb.getItemTest = f
}

// OrderStatusChangeRequest notification parameters.
type OrderStatusChangeRequest struct {
	Notification

	// The recipient’s ID.
	ReceiverID int `schema:"receiver_id,required"`

	// Order ID in VK payment system.
	OrderID int `schema:"order_id,required"`

	// Date when the order was created (as "unix timestamp").
	Date int `schema:"date,required"`

	// New status of the order.
	//
	// chargeable - order ready for payment. User shall complete the order
	// in the application. In case of successful reply the payment system will
	// charge votes to the application account. In case of error message the
	// order will be canceled.
	Status Status `schema:"status,required"`

	// Product name passed to the purchase dialog box.
	//
	// For special offer (call of the payment dialog box with type=offers), Item
	// parameter will include "offer_{offer_id}", item_price parameter is the
	// number of votes charged for such special offer.
	Item string `schema:"item,required"`

	// Product ID in the application.
	ItemID string `schema:"item_id"`

	// Product name.
	ItemTitle string `schema:"item_title,required"`

	// Product image.
	ItemPhotoURL string `schema:"item_photo_url"`

	// Product price.
	ItemPrice string `schema:"item_price,required"`

	// Cost in application currency.
	//
	// See https://vk.com/dev/payments_offers
	ItemCurrencyAmount string `schema:"item_currency_amount"`

	// Product discount.
	ItemDiscount string `schema:"item_discount"`
}

// OfferID returns a offer_id.
func (r OrderStatusChangeRequest) OfferID() int {
	offerID, _ := strconv.Atoi(strings.TrimPrefix(r.Item, "offer_"))
	return offerID
}

// OrderStatusChangeResponse reply parameters.
type OrderStatusChangeResponse struct {
	// order ID in VK payment system
	OrderID int `json:"order_id"`

	// ID of the order in the application. Shall be unique for each order
	AppOrderID int `json:"app_order_id,omitempty"`
}

// OnOrderStatusChange is sent when order status changes.
//
// Please note! In case of repeated notification of Changing order status type
// (with the same order_id), the reply shall be the exact copy of the reply
// for initial notification.
//
// For example, if such notification was sent and positive reply was received
// but it was not received by VK, or for any temporary reason the blocked
// funds could not be charged to the application account immediately after
// reply, such notification will be resent. And there is no need to place a
// new order, you just send the same reply as before (by keeping order_id and
// checking whether such notification was received).
//
// See https://vk.com/dev/payments_status
func (cb *Callback) OnOrderStatusChange(f func(e OrderStatusChangeRequest) (*OrderStatusChangeResponse, *Error)) {
	cb.orderStatusChange = f
}

// OnOrderStatusChangeTest OnOrderStatusChange for test.
func (cb *Callback) OnOrderStatusChangeTest(f func(e OrderStatusChangeRequest) (*OrderStatusChangeResponse, *Error)) {
	cb.orderStatusChangeTest = f
}

// GetSubscriptionRequest notification parameters.
type GetSubscriptionRequest struct {
	Notification

	// The recipient’s ID.
	ReceiverID int `schema:"receiver_id,required"`

	// Order identifier in the VK payment system.
	OrderID int `schema:"order_id,required"`

	// User language in the language_country format.
	//
	// ru_RU — Russian
	//
	// uk_UA — Ukrainian
	//
	// be_BY — Belarusian
	//
	// en_US — English
	Lang string `schema:"lang,required"`

	// Product name passed to the subscription dialog window.
	Item string `schema:"item,required"`
}

// GetSubscriptionResponse reply parameters.
//
// Using the item identifier received in the notification, developers should
// return current information about it. If there is no product, error 20
// “Subscription does not exist” should be returned in the response.
//
// 		&payments.Error{
// 			Code:     payments.ProductNotExist,
// 			Msg:      "Subscription does not exist",
// 			Critical: true,
// 		}
//
// Warning! Due to the item being sent on the client-side, this parameter can be changed by users.
type GetSubscriptionResponse struct {
	// Product identifier in the application.
	ItemID int `json:"item_id,omitempty"`

	// Subscription name.
	Title string `json:"title"`

	// Image URL on the developer’s server for the subscription.
	PhotoURL string `json:"photo_url,omitempty"`

	// Subscription price shown in votes.
	Price int `json:"price"`

	// Subscription period duration in days. Possible values: 3, 7, 30.
	Period int `json:"period"`

	// Trial period duration in days. Possible values: 3, 7, 30.
	TrialDuration int `json:"trial_duration,omitempty"`

	// Allows product caching for {expiration} seconds. Permitted range is
	// from 600 to 604,800 seconds.
	//
	// Warning! Without the parameter, caching of the product is possible for
	// 3,600 seconds if many identical consecutive responses occur. To disable
	// caching, it is necessary to pass 0 as the parameter value.
	Expiration int `json:"expiration,omitempty"`
}

// OnGetSubscription is sent when a subscription dialog window is opened via
// application.
//
// See https://vk.com/dev/payments_getsubscription
func (cb *Callback) OnGetSubscription(f func(e GetSubscriptionRequest) (*GetSubscriptionResponse, *Error)) {
	cb.getSubscription = f
}

// OnGetSubscriptionTest OnGetSubscription for test.
func (cb *Callback) OnGetSubscriptionTest(f func(e GetSubscriptionRequest) (*GetSubscriptionResponse, *Error)) {
	cb.getSubscriptionTest = f
}

// SubscriptionStatusChangeRequest notification parameters.
type SubscriptionStatusChangeRequest struct {
	Notification

	// Reason for cancellation.
	CancelReason Reason `schema:"cancel_reason"`

	// Product identifier in the application.
	ItemID string `schema:"item_id,required"`

	// Product price.
	ItemPrice string `schema:"item_price,required"`

	// New subscription status. Possible values:
	//
	// chargeable — subscription is ready for payments. An order for the user
	// needs to be processed within the application. If the response is
	// successful, the payment system credits votes to the application
	// balance. If the response is a failure, the order is canceled.
	//
	// active — subscription is active;
	//
	// cancelled — subscription is canceled.
	Status Status `schema:"status,required"`

	// Subscription is active until the end of the paid period (status = active).
	//
	// integer, [1]
	PendingCancel int `schema:"pending_cancel"`

	// Subscription identifier.
	SubscriptionID int `schema:"subscription_id,required"`

	// Subscription next bill time
	NextBillTime int `schema:"next_bill_time"`
}

// SubscriptionStatusChangeResponse reply parameters.
type SubscriptionStatusChangeResponse struct {
	// Global subscription identifier.
	SubscriptionID int `json:"subscription_id"`

	// Order identifier within the app. Must be unique for each order.
	AppOrderID int `json:"app_order_id,omitempty"`
}

// OnSubscriptionStatusChange is sent the moment the subscription status
// changes. Please note that the subscription status doesn’t change after
// renewal (withdrawal of a new payment from a user). You won’t receive a
// payment notification about subscription renewal. It’s sent only if a user
// decides to not renew and thus cancels the subscription.
// Warning! If the notification of Change in subscription status type (with
// the same subscription_id) occurs repeatedly, the response should be
// identical to the response of the initial notification.
//
// For example, if this notification was sent and a positive response was
// received, but it didn’t reach VK servers or for some temporary reasons
// transferring frozen funds to the application’s account failed right after
// receiving the response, the same notification will be sent again. There is
// no need to issue a new order. It is necessary to send the same response as
// the previous time (having saved the subscription_id and verified through it
// as if the same notification had already been received).
//
// Note that if the subscription was canceled due to lack of funds on a user’s
// account, it can be renewed within five days (provided that the user
// replenishes their account balance within these five days). In this case, the
// existing subscription is the one renewed, and there is no need to create a
// new one.
//
// See https://vk.com/dev/payments_subscriptionstatus
func (cb *Callback) OnSubscriptionStatusChange(f func(e SubscriptionStatusChangeRequest) (*SubscriptionStatusChangeResponse, *Error)) {
	cb.subscriptionStatusChange = f
}

// OnSubscriptionStatusChangeTest OnSubscriptionStatusChange for test.
func (cb *Callback) OnSubscriptionStatusChangeTest(f func(e SubscriptionStatusChangeRequest) (*SubscriptionStatusChangeResponse, *Error)) {
	cb.subscriptionStatusChangeTest = f
}

// response struct.
type response struct {
	Response interface{} `json:"response,omitempty"`
	Error    *Error      `json:"error,omitempty"`
}

func (cb *Callback) handlerForm(form url.Values) (*response, error) {
	response := new(response)
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	t := NotificationType(form.Get("notification_type"))

	switch {
	case t == GetItem && cb.getItem != nil:
		var event GetItemRequest
		if err := decoder.Decode(&event, form); err != nil {
			return nil, err
		}

		response.Response, response.Error = cb.getItem(event)
	case t == GetItem.Test() && cb.getItemTest != nil:
		var event GetItemRequest
		if err := decoder.Decode(&event, form); err != nil {
			return nil, err
		}

		response.Response, response.Error = cb.getItemTest(event)
	case t == GetSubscription && cb.getSubscription != nil:
		var event GetSubscriptionRequest
		if err := decoder.Decode(&event, form); err != nil {
			return nil, err
		}

		response.Response, response.Error = cb.getSubscription(event)
	case t == GetSubscription.Test() && cb.getSubscriptionTest != nil:
		var event GetSubscriptionRequest
		if err := decoder.Decode(&event, form); err != nil {
			return nil, err
		}

		response.Response, response.Error = cb.getSubscriptionTest(event)
	case t == OrderStatusChange && cb.orderStatusChange != nil:
		var event OrderStatusChangeRequest
		if err := decoder.Decode(&event, form); err != nil {
			return nil, err
		}

		response.Response, response.Error = cb.orderStatusChange(event)
	case t == OrderStatusChange.Test() && cb.orderStatusChangeTest != nil:
		var event OrderStatusChangeRequest
		if err := decoder.Decode(&event, form); err != nil {
			return nil, err
		}

		response.Response, response.Error = cb.orderStatusChangeTest(event)
	case t == SubscriptionStatusChange && cb.subscriptionStatusChange != nil:
		var event SubscriptionStatusChangeRequest
		if err := decoder.Decode(&event, form); err != nil {
			return nil, err
		}

		response.Response, response.Error = cb.subscriptionStatusChange(event)
	case t == SubscriptionStatusChange.Test() && cb.subscriptionStatusChangeTest != nil:
		var event SubscriptionStatusChangeRequest
		if err := decoder.Decode(&event, form); err != nil {
			return nil, err
		}

		response.Response, response.Error = cb.subscriptionStatusChangeTest(event)
	default:
		response.Error = &Error{
			Code:     CommonError,
			Msg:      fmt.Sprintf("%s not processed", t),
			Critical: true,
		}
	}

	return response, nil
}

// HandleFunc a request handler that process notifications
// and return either the processing result if successful
// or an error description if unsuccessful.
func (cb *Callback) HandleFunc(w http.ResponseWriter, r *http.Request) {
	writeResponse := func(resp response) {
		w.Header().Add("Content-Type", "application/json; encoding=utf-8")
		w.WriteHeader(http.StatusOK)

		encoder := json.NewEncoder(w)
		_ = encoder.Encode(resp)
	}

	if err := r.ParseForm(); err != nil {
		// NOTE: what about net error?
		writeResponse(response{
			Error: &Error{
				Code:     BadRequest,
				Msg:      err.Error(),
				Critical: true,
			},
		})

		return
	}

	// If signatures do not match, give the error 10 in response
	sign := cb.Sign(r.PostForm)
	if r.PostForm.Get("sig") != sign {
		writeResponse(response{
			Error: &Error{
				Code:     BadSignatures,
				Msg:      "The calculated and sent signatures do not match",
				Critical: true,
			},
		})

		return
	}

	resp, err := cb.handlerForm(r.PostForm)
	if err != nil {
		writeResponse(response{
			Error: &Error{
				Code:     BadRequest,
				Msg:      err.Error(),
				Critical: true,
			},
		})

		return
	}

	writeResponse(*resp)
}

// Sign return signature.
//
// The parameter sig equals md5 from the concatenation of pairs parameter
// name=parameter value, placed in ascending alphabetical order according
// to the parameter name and the secret signature api_secret indicated in
// your app settings.
func (cb *Callback) Sign(values url.Values) string {
	h := md5.New() // nolint: gosec
	_, _ = io.WriteString(h, getConcatenationParams(values)+cb.Secret)

	return toHex(h.Sum(nil))
}

// getConcatenationParams concatenation of pairs parameter name=parameter
// value, placed in ascending alphabetical order according to the
// parameter name.
func getConcatenationParams(values url.Values) string {
	if values == nil {
		return ""
	}

	var buf strings.Builder

	keys := make([]string, 0, len(values))

	for k := range values {
		if k == "sig" {
			continue
		}

		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		vs := values[k]
		for _, v := range vs {
			_, _ = buf.WriteString(k)
			_ = buf.WriteByte('=')
			_, _ = buf.WriteString(v)
		}
	}

	return buf.String()
}

// toHex return hex string.
func toHex(src []byte) string {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	return string(dst)
}
