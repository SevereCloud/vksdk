package payments_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/SevereCloud/vksdk/v2/payments"
	"github.com/stretchr/testify/assert"
)

const secret = "secret"

// response object.
type response struct {
	Response interface{}    `json:"response,omitempty"`
	Error    payments.Error `json:"error,omitempty"`
}

func TestNewCallback(t *testing.T) {
	t.Parallel()

	cb := payments.NewCallback(secret)
	assert.Equal(t, secret, cb.Secret)

	f := func(v url.Values, wantResp response) {
		t.Helper()

		req, err := http.NewRequest(http.MethodPost, "/payments", strings.NewReader(v.Encode()))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(cb.HandleFunc)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json; encoding=utf-8", rr.Header().Get("Content-Type"))

		var resp response
		err = json.NewDecoder(rr.Body).Decode(&resp)
		assert.NoError(t, err)

		assert.Equal(t, wantResp, resp)
	}

	f(
		url.Values{
			"notification_type": {"1"},
			"app_id":            {"1"},
			"user_id":           {"1"},
			"receiver_id":       {"1"},
			"order_id":          {"1"},
			"subscription_id":   {"1"},
			"sig":               {"1"},
		},
		response{
			Error: payments.Error{
				Code:     payments.BadSignatures,
				Msg:      "The calculated and sent signatures do not match",
				Critical: true,
			},
		},
	)
	f(
		url.Values{
			"notification_type": {"not_processed"},
			"app_id":            {"1"},
			"user_id":           {"1"},
			"receiver_id":       {"1"},
			"order_id":          {"1"},
			"subscription_id":   {"1"},
			"sig":               {"4dd8fa646851ff3f756e6c26d321471b"},
		},
		response{
			Error: payments.Error{
				Code:     payments.CommonError,
				Msg:      "not_processed not processed",
				Critical: true,
			},
		},
	)
}

func TestCallback_Sign(t *testing.T) {
	t.Parallel()

	cb := payments.NewCallback(secret)

	f := func(v url.Values, want string) {
		assert.Equal(t, want, cb.Sign(v))
	}

	f(nil, "5ebe2294ecd0e0f08eab7690d2a6ee69")
	f(
		url.Values{
			"notification_type": {"not_processed"},
			"app_id":            {"1"},
			"user_id":           {"1"},
			"receiver_id":       {"1"},
			"order_id":          {"1"},
			"subscription_id":   {"1"},
			"sig":               {"4dd8fa646851ff3f756e6c26d321471b"},
		},
		"4dd8fa646851ff3f756e6c26d321471b",
	)
	f(
		url.Values{
			"notification_type": {"not_processed"},
			"app_id":            {"1"},
			"user_id":           {"1"},
			"receiver_id":       {"1"},
			"order_id":          {"1"},
			"subscription_id":   {"1"},
		},
		"4dd8fa646851ff3f756e6c26d321471b",
	)
}

func TestCallback_OnGetItem(t *testing.T) {
	t.Parallel()

	paymentErr := payments.Error{
		Code:     1,
		Msg:      "error message",
		Critical: true,
	}

	var v url.Values

	cb := payments.NewCallback(secret)
	assert.Equal(t, secret, cb.Secret)

	cb.OnGetItem(func(_ payments.GetItemRequest) (*payments.GetItemResponse, *payments.Error) {
		return nil, &paymentErr
	})

	cb.OnGetItemTest(func(_ payments.GetItemRequest) (*payments.GetItemResponse, *payments.Error) {
		return nil, &paymentErr
	})

	f := func(v url.Values, wantResp response) {
		t.Helper()

		req, err := http.NewRequest(http.MethodPost, "/payments", strings.NewReader(v.Encode()))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(cb.HandleFunc)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json; encoding=utf-8", rr.Header().Get("Content-Type"))

		var resp response
		err = json.NewDecoder(rr.Body).Decode(&resp)
		assert.NoError(t, err)

		assert.Equal(t, wantResp.Response, resp.Response)
		assert.Equal(t, wantResp.Error.Code, resp.Error.Code)
	}

	v = url.Values{
		"notification_type": {
			string(payments.GetItem),
		},
		"app_id":      {"1"},
		"user_id":     {"1"},
		"receiver_id": {"1"},
		"order_id":    {"1"},
		"lang":        {"1"},
		"item":        {"1"},
	}
	v.Add("sig", cb.Sign(v))
	f(
		v,
		response{
			Error: paymentErr,
		},
	)

	v = url.Values{
		"notification_type": {
			string(payments.GetItem.Test()),
		},
		"app_id":      {"1"},
		"user_id":     {"1"},
		"receiver_id": {"1"},
		"order_id":    {"1"},
		"lang":        {"1"},
		"item":        {"1"},
	}
	v.Add("sig", cb.Sign(v))
	f(
		v,
		response{
			Error: paymentErr,
		},
	)

	// Check error
	v = url.Values{
		"notification_type": {
			string(payments.GetItem),
		},
	}
	v.Add("sig", cb.Sign(v))
	f(
		v,
		response{
			Error: payments.Error{Code: payments.BadRequest},
		},
	)

	v = url.Values{
		"notification_type": {
			string(payments.GetItem.Test()),
		},
	}
	v.Add("sig", cb.Sign(v))
	f(
		v,
		response{
			Error: payments.Error{Code: payments.BadRequest},
		},
	)
}

func TestCallback_OnGetSubscription(t *testing.T) {
	t.Parallel()

	paymentErr := payments.Error{
		Code:     1,
		Msg:      "error message",
		Critical: true,
	}

	var v url.Values

	cb := payments.NewCallback(secret)
	assert.Equal(t, secret, cb.Secret)

	cb.OnGetSubscription(func(_ payments.GetSubscriptionRequest) (*payments.GetSubscriptionResponse, *payments.Error) {
		return nil, &paymentErr
	})

	cb.OnGetSubscriptionTest(func(_ payments.GetSubscriptionRequest) (*payments.GetSubscriptionResponse, *payments.Error) {
		return nil, &paymentErr
	})

	f := func(v url.Values, wantResp response) {
		t.Helper()

		req, err := http.NewRequest(http.MethodPost, "/payments", strings.NewReader(v.Encode()))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(cb.HandleFunc)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json; encoding=utf-8", rr.Header().Get("Content-Type"))

		var resp response
		err = json.NewDecoder(rr.Body).Decode(&resp)
		assert.NoError(t, err)

		assert.Equal(t, wantResp.Response, resp.Response)

		if !assert.Equal(t, wantResp.Error.Code, resp.Error.Code) {
			t.Error(resp.Error.Msg)
		}
	}

	v = url.Values{
		"notification_type": {
			string(payments.GetSubscription),
		},
		"app_id":      {"1"},
		"item":        {"1"},
		"user_id":     {"1"},
		"receiver_id": {"1"},
		"order_id":    {"1"},
		"lang":        {"1"},
	}
	v.Add("sig", cb.Sign(v))
	f(
		v,
		response{
			Error: paymentErr,
		},
	)

	v = url.Values{
		"notification_type": {
			string(payments.GetSubscription.Test()),
		},
		"app_id":      {"1"},
		"item":        {"1"},
		"user_id":     {"1"},
		"receiver_id": {"1"},
		"order_id":    {"1"},
		"lang":        {"1"},
	}
	v.Add("sig", cb.Sign(v))
	f(
		v,
		response{
			Error: paymentErr,
		},
	)

	// Check error
	v = url.Values{
		"notification_type": {
			string(payments.GetSubscription),
		},
	}
	v.Add("sig", cb.Sign(v))
	f(
		v,
		response{
			Error: payments.Error{Code: payments.BadRequest},
		},
	)

	v = url.Values{
		"notification_type": {
			string(payments.GetSubscription.Test()),
		},
	}
	v.Add("sig", cb.Sign(v))
	f(
		v,
		response{
			Error: payments.Error{Code: payments.BadRequest},
		},
	)
}

func TestCallback_OnOrderStatusChange(t *testing.T) {
	t.Parallel()

	paymentErr := payments.Error{
		Code:     1,
		Msg:      "error message",
		Critical: true,
	}

	var v url.Values

	cb := payments.NewCallback(secret)
	assert.Equal(t, secret, cb.Secret)

	cb.OnOrderStatusChange(func(_ payments.OrderStatusChangeRequest) (
		*payments.OrderStatusChangeResponse,
		*payments.Error,
	) {
		return nil, &paymentErr
	})

	cb.OnOrderStatusChangeTest(func(_ payments.OrderStatusChangeRequest) (
		*payments.OrderStatusChangeResponse,
		*payments.Error,
	) {
		return nil, &paymentErr
	})

	f := func(v url.Values, wantResp response) {
		t.Helper()

		req, err := http.NewRequest(http.MethodPost, "/payments", strings.NewReader(v.Encode()))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(cb.HandleFunc)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json; encoding=utf-8", rr.Header().Get("Content-Type"))

		var resp response
		err = json.NewDecoder(rr.Body).Decode(&resp)
		assert.NoError(t, err)

		assert.Equal(t, wantResp.Response, resp.Response)

		if !assert.Equal(t, wantResp.Error.Code, resp.Error.Code) {
			t.Error(resp.Error.Msg)
		}
	}

	v = url.Values{
		"notification_type": {
			string(payments.OrderStatusChange),
		},
		"app_id":         {"1"},
		"user_id":        {"1"},
		"receiver_id":    {"1"},
		"order_id":       {"1"},
		"date":           {"1"},
		"status":         {"1"},
		"item":           {"1"},
		"item_id":        {"1"},
		"item_title":     {"1"},
		"item_photo_url": {"1"},
		"item_price":     {"1"},
		"item_discount":  {"1"},
	}
	v.Add("sig", cb.Sign(v))
	f(
		v,
		response{
			Error: paymentErr,
		},
	)

	v = url.Values{
		"notification_type": {
			string(payments.OrderStatusChange.Test()),
		},
		"app_id":         {"1"},
		"user_id":        {"1"},
		"receiver_id":    {"1"},
		"order_id":       {"1"},
		"date":           {"1"},
		"status":         {"1"},
		"item":           {"1"},
		"item_id":        {"1"},
		"item_title":     {"1"},
		"item_photo_url": {"1"},
		"item_price":     {"1"},
		"item_discount":  {"1"},
	}
	v.Add("sig", cb.Sign(v))
	f(
		v,
		response{
			Error: paymentErr,
		},
	)

	// Check error
	v = url.Values{
		"notification_type": {
			string(payments.OrderStatusChange),
		},
	}
	v.Add("sig", cb.Sign(v))
	f(
		v,
		response{
			Error: payments.Error{Code: payments.BadRequest},
		},
	)

	v = url.Values{
		"notification_type": {
			string(payments.OrderStatusChange.Test()),
		},
	}
	v.Add("sig", cb.Sign(v))
	f(
		v,
		response{
			Error: payments.Error{Code: payments.BadRequest},
		},
	)
}

func TestCallback_OnSubscriptionStatusChange(t *testing.T) {
	t.Parallel()

	paymentErr := payments.Error{
		Code:     1,
		Msg:      "error message",
		Critical: true,
	}

	var v url.Values

	cb := payments.NewCallback(secret)
	assert.Equal(t, secret, cb.Secret)

	cb.OnSubscriptionStatusChange(func(_ payments.SubscriptionStatusChangeRequest) (
		*payments.SubscriptionStatusChangeResponse,
		*payments.Error,
	) {
		return nil, &paymentErr
	})

	cb.OnSubscriptionStatusChangeTest(func(_ payments.SubscriptionStatusChangeRequest) (
		*payments.SubscriptionStatusChangeResponse,
		*payments.Error,
	) {
		return nil, &paymentErr
	})

	f := func(v url.Values, wantResp response) {
		t.Helper()

		req, err := http.NewRequest(http.MethodPost, "/payments", strings.NewReader(v.Encode()))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(cb.HandleFunc)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json; encoding=utf-8", rr.Header().Get("Content-Type"))

		var resp response
		err = json.NewDecoder(rr.Body).Decode(&resp)
		assert.NoError(t, err)

		assert.Equal(t, wantResp.Response, resp.Response)

		if !assert.Equal(t, wantResp.Error.Code, resp.Error.Code) {
			t.Error(resp.Error.Msg)
		}
	}

	v = url.Values{
		"notification_type": {
			string(payments.SubscriptionStatusChange),
		},
		"app_id":          {"1"},
		"user_id":         {"1"},
		"receiver_id":     {"1"},
		"cancel_reason":   {"1"},
		"item_id":         {"1"},
		"item_price":      {"1"},
		"status":          {"1"},
		"pending_cancel":  {"1"},
		"subscription_id": {"1"},
	}
	v.Add("sig", cb.Sign(v))
	f(
		v,
		response{
			Error: paymentErr,
		},
	)

	v = url.Values{
		"notification_type": {
			string(payments.SubscriptionStatusChange.Test()),
		},
		"app_id":          {"1"},
		"user_id":         {"1"},
		"receiver_id":     {"1"},
		"cancel_reason":   {"1"},
		"item_id":         {"1"},
		"item_price":      {"1"},
		"status":          {"1"},
		"pending_cancel":  {"1"},
		"subscription_id": {"1"},
	}
	v.Add("sig", cb.Sign(v))
	f(
		v,
		response{
			Error: paymentErr,
		},
	)

	// Check error
	v = url.Values{
		"notification_type": {
			string(payments.SubscriptionStatusChange),
		},
	}
	v.Add("sig", cb.Sign(v))
	f(
		v,
		response{
			Error: payments.Error{Code: payments.BadRequest},
		},
	)

	v = url.Values{
		"notification_type": {
			string(payments.SubscriptionStatusChange.Test()),
		},
	}
	v.Add("sig", cb.Sign(v))
	f(
		v,
		response{
			Error: payments.Error{Code: payments.BadRequest},
		},
	)
}

func TestParseForm(t *testing.T) {
	t.Parallel()

	cb := payments.NewCallback(secret)

	req, err := http.NewRequest(http.MethodPost, "/payments", strings.NewReader("body"))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "text/plain; boundary=")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(cb.HandleFunc)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json; encoding=utf-8", rr.Header().Get("Content-Type"))

	var resp response
	err = json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err)

	wantResp := response{
		Error: payments.Error{
			Code:     payments.BadRequest,
			Msg:      "mime: invalid media parameter",
			Critical: true,
		},
	}
	assert.Equal(t, wantResp, resp)
}
