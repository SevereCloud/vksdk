/*
Package errors for api.

See more https://vk.com/dev/errors
*/
package errors // import "github.com/SevereCloud/vksdk/api/errors"

import (
	"fmt"

	"github.com/SevereCloud/vksdk/object"
	"github.com/pkg/errors"
)

// Error codes. See https://vk.com/dev/errors
const (
	NoType ErrorType = 0 // NoType error

	// Unknown error occurred
	//
	// Try again later.
	Unknown ErrorType = 1

	// Application is disabled. Enable your application or use test mode
	//
	// You need to switch on the app in Settings (https://vk.com/editapp?id={Your API_ID}
	// or use the test mode (test_mode=1).
	Disabled ErrorType = 2

	// Unknown method passed
	//
	// Check the method name: http://vk.com/dev/methods
	Method    ErrorType = 3
	Signature ErrorType = 4 // Incorrect signature

	// User authorization failed
	//
	// Make sure that you use a correct authorization type
	Auth ErrorType = 5

	// Too many requests per second
	//
	// Decrease the request frequency or use the execute method.
	// More details on frequency limits here:
	// https://vk.com/dev/api_requests
	TooMany ErrorType = 6

	// Permission to perform this action is denied
	//
	// Make sure that your have received required permissions during the authorization.
	// You can do it with the account.getAppPermissions method.
	// https://vk.com/dev/permissions
	Permission ErrorType = 7

	// Invalid request
	//
	// Check the request syntax and used parameters list (it can be found on a method description page)
	Request ErrorType = 8

	// Flood control
	//
	// You need to decrease the count of identical requests. For more efficient work you may use execute.
	Flood ErrorType = 9

	// Internal server error
	//
	// Try again later.
	Server ErrorType = 10

	// In test mode application should be disabled or user should be authorized
	//
	// Switch the app off in Settings: https://vk.com/editapp?id={Your API_ID}
	EnabledInTest ErrorType = 11

	// Captcha needed
	//
	// See https://vk.com/dev/captcha_error
	Captcha ErrorType = 14

	// Access denied
	//
	// Make sure that you use correct identifiers and the content is available for the user in the full version of the site.
	Access ErrorType = 15

	// HTTP authorization failed
	//
	// To avoid this error check if a user has the 'Use secure connection' option enabled with the account.getInfo method.
	AuthHTTPS ErrorType = 16

	// Validation required
	//
	// Make sure that you don't use a token received with http://vk.com/dev/auth_mobile for a request from the server. It's restricted.
	//
	// https://vk.com/dev/need_validation
	AuthValidation ErrorType = 17
	UserDeleted    ErrorType = 18 // User was deleted or banned
	Blocked        ErrorType = 19 // Content blocked

	// Permission to perform this action is denied for non-standalone applications
	//
	// If you see this error despite your app has the Standalone type, make sure that you use redirect_uri=https://oauth.vk.com/blank.html.
	// Details here: https://vk.com/dev/auth_mobile
	MethodPermission ErrorType = 20
	MethodAds        ErrorType = 21 // Permission to perform this action is allowed only for standalone and OpenAPI applications
	Upload           ErrorType = 22 // Upload error

	// This method was disabled
	//
	// All the methods available now are listed here: http://vk.com/dev/methods
	MethodDisabled ErrorType = 23

	// Confirmation required
	//
	// In some cases VK requires to request action confirmation from the user (for Standalone apps only).
	//
	// Following parameter is transmitted with the error message as well:
	//
	// confirmation_text – text of the message to be shown in the default confirmation window.
	//
	// The app should display the default confirmation windos
	// with text from confirmation_text and two buttons: "Continue" and "Cancel".
	// If user confirms the action repeat the request with an extra parameter: confirm = 1.
	//
	// https://vk.com/dev/need_confirmation
	NeedConfirmation      ErrorType = 24
	NeedTokenConfirmation ErrorType = 25 // Token confirmation required
	GroupAuth             ErrorType = 27 // Group authorization failed
	AppAuth               ErrorType = 28 // Application authorization failed

	// Rate limit reached
	//
	// More details on rate limits here: https://vk.com/dev/data_limits
	RateLimit      ErrorType = 29
	PrivateProfile ErrorType = 30 // This profile is private

	// One of the parameters specified was missing or invalid
	//
	// Check the reqired parameters list and their format on a method description page.
	Param ErrorType = 100

	// Invalid application API ID
	//
	// Find the app in the administrated list in settings: http://vk.com/apps?act=settings
	// And set the correct API_ID in the request.
	ParamAPIID   ErrorType = 101
	Limits       ErrorType = 103 // Out of limits
	NotFound     ErrorType = 104 // Not found
	SaveFile     ErrorType = 105 // Couldn't save file
	ActionFailed ErrorType = 106 // Unable to process action

	// Invalid user id
	//
	// Make sure that you use a correct id. You can get an id using a screen
	// name with the utils.resolveScreenName method
	ParamUserID        ErrorType = 113
	ParamAlbumID       ErrorType = 114 // Invalid album id
	ParamServer        ErrorType = 118 // Invalid server
	ParamTitle         ErrorType = 119 // Invalid title
	ParamPhotos        ErrorType = 122 // Invalid photos
	ParamHash          ErrorType = 121 // Invalid hash
	ParamPhoto         ErrorType = 129 // Invalid photo
	ParamGroupID       ErrorType = 125 // Invalid group id
	ParamPageID        ErrorType = 140 // Page not found
	AccessPage         ErrorType = 141 // Access to page denied
	MobileNotActivated ErrorType = 146 // The mobile number of the user is unknown
	InsufficientFunds  ErrorType = 147 // Application has insufficient funds
	AccessMenu         ErrorType = 148 // Access to the menu of the user denied

	// Invalid timestamp
	//
	// You may get a correct value with the utils.getServerTime method
	ParamTimestamp     ErrorType = 150
	FriendsListID      ErrorType = 171 // Invalid list id
	FriendsListLimit   ErrorType = 173 // Reached the maximum number of lists
	FriendsAddYourself ErrorType = 174 // Cannot add user himself as friend
	FriendsAddInEnemy  ErrorType = 175 // Cannot add this user to friends as they have put you on their blacklist
	FriendsAddEnemy    ErrorType = 176 // Cannot add this user to friends as you put him on blacklist
	FriendsAddNotFound ErrorType = 177 // Cannot add this user to friends as user not found
	ParamNoteID        ErrorType = 180 // Note not found
	AccessNote         ErrorType = 181 // Access to note denied
	AccessNoteComment  ErrorType = 182 // You can't comment this note
	AccessComment      ErrorType = 183 // Access to comment denied

	// Access to album denied
	//
	// Make sure you use correct ids (owner_id is always positive for users,
	// negative for communities) and the current user has access to the
	// requested content in the full version of the site.
	AccessAlbum ErrorType = 200

	// Access to audio denied
	//
	// Make sure you use correct ids (owner_id is always positive for users,
	// negative for communities) and the current user has access to the
	// requested content in the full version of the site.
	AccessAudio ErrorType = 201

	// Access to group denied
	//
	// Make sure that the current user is a member or admin of the community (for closed and private groups and events).
	AccessGroup             ErrorType = 203
	AccessVideo             ErrorType = 204 // Access denied
	AccessMarket            ErrorType = 205 // Access denied
	WallAccessPost          ErrorType = 210 // Access to wall's post denied
	WallAccessComment       ErrorType = 211 // Access to wall's comment denied
	WallAccessReplies       ErrorType = 212 // Access to post comments denied
	WallAccessAddReply      ErrorType = 213 // Access to status replies denied
	WallAddPost             ErrorType = 214 // Access to adding post denied
	WallAdsPublished        ErrorType = 219 // Advertisement post was recently added
	WallTooManyRecipients   ErrorType = 220 // Too many recipients
	StatusNoAudio           ErrorType = 221 // User disabled track name broadcast
	WallLinksForbidden      ErrorType = 222 // Hyperlinks are forbidden
	WallReplyOwnerFlood     ErrorType = 223 // Too many replies
	WallAdsPostLimitReached ErrorType = 224 // Too many ads posts
	DonutDisabled           ErrorType = 225 // Donut is disabled
	PollsAccess             ErrorType = 250 // Access to poll denied
	PollsAnswerID           ErrorType = 252 // Invalid answer id
	PollsPollID             ErrorType = 251 // Invalid poll id
	PollsAccessWithoutVote  ErrorType = 253 // Access denied, please vote first
	AccessGroups            ErrorType = 260 // Access to the groups list is denied due to the user's privacy settings

	// This album is full
	//
	// You need to delete the odd objects from the album or use another album.
	AlbumFull   ErrorType = 300
	AlbumsLimit ErrorType = 302 // Albums number limit is reached

	// Permission denied. You must enable votes processing in application settings
	//
	// Check the app settings: http://vk.com/editapp?id={Your API_ID}&section=payments
	VotesPermission                   ErrorType = 500
	Votes                             ErrorType = 503  // Not enough votes
	NotEnoughMoney                    ErrorType = 504  // Not enough money on owner's balance
	AdsPermission                     ErrorType = 600  // Permission denied. You have no access to operations specified with given object(s)
	WeightedFlood                     ErrorType = 601  // Permission denied. You have requested too many actions this day. Try later
	AdsPartialSuccess                 ErrorType = 602  // Some part of the request has not been completed
	AdsSpecific                       ErrorType = 603  // Some ads error occurred
	AdsObjectDeleted                  ErrorType = 629  // Object deleted
	GroupChangeCreator                ErrorType = 700  // Cannot edit creator role
	GroupNotInClub                    ErrorType = 701  // User should be in club
	GroupTooManyOfficers              ErrorType = 702  // Too many officers in club
	GroupNeed2fa                      ErrorType = 703  // You need to enable 2FA for this action
	GroupHostNeed2fa                  ErrorType = 704  // User needs to enable 2FA for this action
	GroupTooManyAddresses             ErrorType = 706  // Too many addresses in club
	GroupAppIsNotInstalledInCommunity ErrorType = 711  // "Application is not installed in community
	GroupInvalidInviteLink            ErrorType = 714  // Invite link is invalid - expired, deleted or not exists
	VideoAlreadyAdded                 ErrorType = 800  // This video is already added
	VideoCommentsClosed               ErrorType = 801  // Comments for this video are closed
	MessagesUserBlocked               ErrorType = 900  // Can't send messages for users from blacklist
	MessagesDenySend                  ErrorType = 901  // Can't send messages for users without permission
	MessagesPrivacy                   ErrorType = 902  // Can't send messages to this user due to their privacy settings
	MessagesTooOldPts                 ErrorType = 907  // Value of ts or pts is too old
	MessagesTooNewPts                 ErrorType = 908  // Value of ts or pts is too new
	MessagesEditExpired               ErrorType = 909  // Can't edit this message, because it's too old
	MessagesTooBig                    ErrorType = 910  // Can't sent this message, because it's too big
	MessagesKeyboardInvalid           ErrorType = 911  // Keyboard format is invalid
	MessagesChatBotFeature            ErrorType = 912  // This is a chat bot feature, change this status in settings
	MessagesTooLongForwards           ErrorType = 913  // Too many forwarded messages
	MessagesTooLongMessage            ErrorType = 914  // Message is too long
	MessagesChatUserNoAccess          ErrorType = 917  // You don't have access to this chat
	MessagesCantSeeInviteLink         ErrorType = 919  // You can't see invite link for this chat
	MessagesEditKindDisallowed        ErrorType = 920  // Can't edit this kind of message
	MessagesCantFwd                   ErrorType = 921  // Can't forward these messages
	MessagesCantDeleteForAll          ErrorType = 924  // Can't delete this message for everybody
	MessagesChatNotAdmin              ErrorType = 925  // You are not admin of this chat
	MessagesChatNotExist              ErrorType = 927  // Chat does not exist
	MessagesCantChangeInviteLink      ErrorType = 931  // You can't change invite link for this chat
	MessagesGroupPeerAccess           ErrorType = 932  // Your community can't interact with this peer
	MessagesChatUserNotInChat         ErrorType = 935  // User not found in chat
	MessagesContactNotFound           ErrorType = 936  // Contact not found
	MessagesMessageRequestAlreadySend ErrorType = 939  // Message request already send
	MessagesTooManyPosts              ErrorType = 940  // Too many posts in messages
	MessagesCantUseIntent             ErrorType = 943  // Cannot use this intent
	MessagesLimitIntent               ErrorType = 944  // Limits overflow for this intent
	MessagesChatDisabled              ErrorType = 945  // Chat was disabled
	MessagesChatNotSupported          ErrorType = 946  // Chat not support
	ParamPhone                        ErrorType = 1000 // Invalid phone number
	PhoneAlreadyUsed                  ErrorType = 1004 // This phone number is used by another user
	AuthFloodError                    ErrorType = 1105 // Too many auth attempts, try again later
	AuthDelay                         ErrorType = 1112 // Processing.. Try later
	ParamDocID                        ErrorType = 1150 // Invalid document id
	ParamDocDeleteAccess              ErrorType = 1151 // Access to document deleting is denied
	ParamDocTitle                     ErrorType = 1152 // Invalid document title
	ParamDocAccess                    ErrorType = 1153 // Access to document is denied
	PhotoChanged                      ErrorType = 1160 // Original photo was changed
	TooManyLists                      ErrorType = 1170 // Too many feed lists
	AppsAlreadyUnlocked               ErrorType = 1251 // This achievement is already unlocked
	AppsSubscriptionNotFound          ErrorType = 1256 // Subscription not found
	AppsSubscriptionInvalidStatus     ErrorType = 1257 // Subscription is in invalid status
	InvalidAddress                    ErrorType = 1260 // Invalid screen name
	CommunitiesCatalogDisabled        ErrorType = 1310 // Catalog is not available for this user
	CommunitiesCategoriesDisabled     ErrorType = 1311 // Catalog categories are not available for this user
	MarketRestoreTooLate              ErrorType = 1400 // Too late for restore
	MarketCommentsClosed              ErrorType = 1401 // Comments for this market are closed
	MarketAlbumNotFound               ErrorType = 1402 // Album not found
	MarketItemNotFound                ErrorType = 1403 // Item not found
	MarketItemAlreadyAdded            ErrorType = 1404 // Item already added to album
	MarketTooManyItems                ErrorType = 1405 // Too many items
	MarketTooManyItemsInAlbum         ErrorType = 1406 // Too many items in album
	MarketTooManyAlbums               ErrorType = 1407 // Too many albums
	MarketItemHasBadLinks             ErrorType = 1408 // Item has bad links in description
	MarketShopNotEnabled              ErrorType = 1409 // Shop not enabled
	MarketCartEmpty                   ErrorType = 1427 // Cart is empty
	StoryExpired                      ErrorType = 1600 // Story has already expired
	StoryIncorrectReplyPrivacy        ErrorType = 1602 // Incorrect reply privacy
	PrettyCardsCardNotFound           ErrorType = 1900 // Card not found
	PrettyCardsTooManyCards           ErrorType = 1901 // Too many cards
	PrettyCardsCardIsConnectedToPost  ErrorType = 1902 // Card is connected to post
	CallbackServersLimit              ErrorType = 2000 // Servers number limit is reached
	Recaptcha                         ErrorType = 3300 // Recaptcha needed
	PhoneValidation                   ErrorType = 3301 // Phone validation needed
	PasswordValidation                ErrorType = 3302 // Password validation needed
	OtpAppValidation                  ErrorType = 3303 // Otp app validation needed
	EmailConfirmation                 ErrorType = 3304 // Email confirmation needed
	AssertVotes                       ErrorType = 3305 // Assert votes
	TokenExtension                    ErrorType = 3609 // Token extension required
)

// ErrorType is the type of an error.
type ErrorType int

type customError struct {
	errorType     ErrorType
	originalError error
	context       object.Error
}

// New creates a new customError.
func (errorType ErrorType) New(msg string) error {
	return customError{errorType: errorType, originalError: errors.New(msg)}
}

// Newf creates a new customError with formatted message.
func (errorType ErrorType) Newf(msg string, args ...interface{}) error {
	return customError{errorType: errorType, originalError: fmt.Errorf(msg, args...)}
}

// Wrap creates a new wrapped error.
func (errorType ErrorType) Wrap(err error, msg string) error {
	return errorType.Wrapf(err, msg)
}

// Wrapf creates a new wrapped error with formatted message.
func (errorType ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	return customError{errorType: errorType, originalError: errors.Wrapf(err, msg, args...)}
}

// Error returns the mssage of a customError.
func (error customError) Error() string {
	return error.originalError.Error()
}

// New creates a no type error.
func New(vkErr object.Error) error {
	if vkErr.Code == 0 {
		return nil
	}

	return customError{
		errorType:     ErrorType(vkErr.Code),
		originalError: errors.New(vkErr.Message),
		context:       vkErr,
	}
}

// Cause gives the original error.
func Cause(err error) error {
	return errors.Cause(err)
}

// AddErrorContext adds a context to an error.
func AddErrorContext(err error, context object.Error) error {
	if customErr, ok := err.(customError); ok {
		return customError{
			errorType:     customErr.errorType,
			originalError: customErr.originalError,
			context:       context,
		}
	}

	return customError{errorType: NoType, originalError: err, context: context}
}

// GetErrorContext returns the error context.
func GetErrorContext(err error) object.Error {
	if customErr, ok := err.(customError); ok {
		return customErr.context
	}

	return object.Error{}
}

// GetType returns the error type.
func GetType(err error) ErrorType {
	if customErr, ok := err.(customError); ok {
		return customErr.errorType
	}

	return NoType
}
