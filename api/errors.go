package api

import (
	"fmt"

	"github.com/SevereCloud/vksdk/object"
)

// ErrorType is the type of an error.
type ErrorType int

// Error returns the message of a ErrorType.
func (e ErrorType) Error() string {
	return fmt.Sprintf("api: error with code %d", e)
}

// Error codes. See https://vk.com/dev/errors
const (
	ErrNoType ErrorType = 0 // NoType error

	// Unknown error occurred
	//
	// Try again later.
	ErrUnknown ErrorType = 1

	// Application is disabled. Enable your application or use test mode
	//
	// You need to switch on the app in Settings (https://vk.com/editapp?id={Your API_ID}
	// or use the test mode (test_mode=1).
	ErrDisabled ErrorType = 2

	// Unknown method passed
	//
	// Check the method name: http://vk.com/dev/methods
	ErrMethod    ErrorType = 3
	ErrSignature ErrorType = 4 // Incorrect signature

	// User authorization failed
	//
	// Make sure that you use a correct authorization type
	ErrAuth ErrorType = 5

	// Too many requests per second
	//
	// Decrease the request frequency or use the execute method.
	// More details on frequency limits here:
	// https://vk.com/dev/api_requests
	ErrTooMany ErrorType = 6

	// Permission to perform this action is denied
	//
	// Make sure that your have received required permissions during the authorization.
	// You can do it with the account.getAppPermissions method.
	// https://vk.com/dev/permissions
	ErrPermission ErrorType = 7

	// Invalid request
	//
	// Check the request syntax and used parameters list (it can be found on a method description page)
	ErrRequest ErrorType = 8

	// Flood control
	//
	// You need to decrease the count of identical requests. For more efficient work you may use execute.
	ErrFlood ErrorType = 9

	// Internal server error
	//
	// Try again later.
	ErrServer ErrorType = 10

	// In test mode application should be disabled or user should be authorized
	//
	// Switch the app off in Settings: https://vk.com/editapp?id={Your API_ID}
	ErrEnabledInTest ErrorType = 11

	// Unable to compile code.
	ErrCompile ErrorType = 12

	// Runtime error occurred during code invocation.
	ErrRuntime ErrorType = 13

	// Captcha needed
	//
	// See https://vk.com/dev/captcha_error
	ErrCaptcha ErrorType = 14

	// Access denied
	//
	// Make sure that you use correct identifiers and the content is available for the user in the full version of the site.
	ErrAccess ErrorType = 15

	// HTTP authorization failed
	//
	// To avoid this error check if a user has the 'Use secure connection' option enabled with the account.getInfo method.
	ErrAuthHTTPS ErrorType = 16

	// Validation required
	//
	// Make sure that you don't use a token received with http://vk.com/dev/auth_mobile for a request from the server. It's restricted.
	//
	// https://vk.com/dev/need_validation
	ErrAuthValidation ErrorType = 17
	ErrUserDeleted    ErrorType = 18 // User was deleted or banned
	ErrBlocked        ErrorType = 19 // Content blocked

	// Permission to perform this action is denied for non-standalone applications
	//
	// If you see this error despite your app has the Standalone type, make sure that you use redirect_uri=https://oauth.vk.com/blank.html.
	// Details here: https://vk.com/dev/auth_mobile
	ErrMethodPermission ErrorType = 20
	ErrMethodAds        ErrorType = 21 // Permission to perform this action is allowed only for standalone and OpenAPI applications
	ErrUpload           ErrorType = 22 // Upload error

	// This method was disabled
	//
	// All the methods available now are listed here: http://vk.com/dev/methods
	ErrMethodDisabled ErrorType = 23

	// Confirmation required
	//
	// In some cases VK requires to request action confirmation from the user (for Standalone apps only).
	//
	// Following parameter is transmitted with the error message as well:
	//
	// confirmation_text – text of the message to be shown in the default confirmation window.
	//
	// The app should display the default confirmation window
	// with text from confirmation_text and two buttons: "Continue" and "Cancel".
	// If user confirms the action repeat the request with an extra parameter: confirm = 1.
	//
	// https://vk.com/dev/need_confirmation
	ErrNeedConfirmation      ErrorType = 24
	ErrNeedTokenConfirmation ErrorType = 25 // Token confirmation required
	ErrGroupAuth             ErrorType = 27 // Group authorization failed
	ErrAppAuth               ErrorType = 28 // Application authorization failed

	// Rate limit reached
	//
	// More details on rate limits here: https://vk.com/dev/data_limits
	ErrRateLimit      ErrorType = 29
	ErrPrivateProfile ErrorType = 30 // This profile is private

	// One of the parameters specified was missing or invalid
	//
	// Check the required parameters list and their format on a method description page.
	ErrParam ErrorType = 100

	// Invalid application API ID
	//
	// Find the app in the administrated list in settings: http://vk.com/apps?act=settings
	// And set the correct API_ID in the request.
	ErrParamAPIID   ErrorType = 101
	ErrLimits       ErrorType = 103 // Out of limits
	ErrNotFound     ErrorType = 104 // Not found
	ErrSaveFile     ErrorType = 105 // Couldn't save file
	ErrActionFailed ErrorType = 106 // Unable to process action

	// Invalid user id
	//
	// Make sure that you use a correct id. You can get an id using a screen
	// name with the utils.resolveScreenName method
	ErrParamUserID        ErrorType = 113
	ErrParamAlbumID       ErrorType = 114 // Invalid album id
	ErrParamServer        ErrorType = 118 // Invalid server
	ErrParamTitle         ErrorType = 119 // Invalid title
	ErrParamPhotos        ErrorType = 122 // Invalid photos
	ErrParamHash          ErrorType = 121 // Invalid hash
	ErrParamPhoto         ErrorType = 129 // Invalid photo
	ErrParamGroupID       ErrorType = 125 // Invalid group id
	ErrParamPageID        ErrorType = 140 // Page not found
	ErrAccessPage         ErrorType = 141 // Access to page denied
	ErrMobileNotActivated ErrorType = 146 // The mobile number of the user is unknown
	ErrInsufficientFunds  ErrorType = 147 // Application has insufficient funds
	ErrAccessMenu         ErrorType = 148 // Access to the menu of the user denied

	// Invalid timestamp
	//
	// You may get a correct value with the utils.getServerTime method
	ErrParamTimestamp     ErrorType = 150
	ErrFriendsListID      ErrorType = 171 // Invalid list id
	ErrFriendsListLimit   ErrorType = 173 // Reached the maximum number of lists
	ErrFriendsAddYourself ErrorType = 174 // Cannot add user himself as friend
	ErrFriendsAddInEnemy  ErrorType = 175 // Cannot add this user to friends as they have put you on their blacklist
	ErrFriendsAddEnemy    ErrorType = 176 // Cannot add this user to friends as you put him on blacklist
	ErrFriendsAddNotFound ErrorType = 177 // Cannot add this user to friends as user not found
	ErrParamNoteID        ErrorType = 180 // Note not found
	ErrAccessNote         ErrorType = 181 // Access to note denied
	ErrAccessNoteComment  ErrorType = 182 // You can't comment this note
	ErrAccessComment      ErrorType = 183 // Access to comment denied

	// Access to album denied
	//
	// Make sure you use correct ids (owner_id is always positive for users,
	// negative for communities) and the current user has access to the
	// requested content in the full version of the site.
	ErrAccessAlbum ErrorType = 200

	// Access to audio denied
	//
	// Make sure you use correct ids (owner_id is always positive for users,
	// negative for communities) and the current user has access to the
	// requested content in the full version of the site.
	ErrAccessAudio ErrorType = 201

	// Access to group denied
	//
	// Make sure that the current user is a member or admin of the community (for closed and private groups and events).
	ErrAccessGroup             ErrorType = 203
	ErrAccessVideo             ErrorType = 204 // Access denied
	ErrAccessMarket            ErrorType = 205 // Access denied
	ErrWallAccessPost          ErrorType = 210 // Access to wall's post denied
	ErrWallAccessComment       ErrorType = 211 // Access to wall's comment denied
	ErrWallAccessReplies       ErrorType = 212 // Access to post comments denied
	ErrWallAccessAddReply      ErrorType = 213 // Access to status replies denied
	ErrWallAddPost             ErrorType = 214 // Access to adding post denied
	ErrWallAdsPublished        ErrorType = 219 // Advertisement post was recently added
	ErrWallTooManyRecipients   ErrorType = 220 // Too many recipients
	ErrStatusNoAudio           ErrorType = 221 // User disabled track name broadcast
	ErrWallLinksForbidden      ErrorType = 222 // Hyperlinks are forbidden
	ErrWallReplyOwnerFlood     ErrorType = 223 // Too many replies
	ErrWallAdsPostLimitReached ErrorType = 224 // Too many ads posts
	ErrDonutDisabled           ErrorType = 225 // Donut is disabled
	ErrPollsAccess             ErrorType = 250 // Access to poll denied
	ErrPollsAnswerID           ErrorType = 252 // Invalid answer id
	ErrPollsPollID             ErrorType = 251 // Invalid poll id
	ErrPollsAccessWithoutVote  ErrorType = 253 // Access denied, please vote first
	ErrAccessGroups            ErrorType = 260 // Access to the groups list is denied due to the user's privacy settings

	// This album is full
	//
	// You need to delete the odd objects from the album or use another album.
	ErrAlbumFull   ErrorType = 300
	ErrAlbumsLimit ErrorType = 302 // Albums number limit is reached

	// Permission denied. You must enable votes processing in application settings
	//
	// Check the app settings: http://vk.com/editapp?id={Your API_ID}&section=payments
	ErrVotesPermission                   ErrorType = 500
	ErrVotes                             ErrorType = 503  // Not enough votes
	ErrNotEnoughMoney                    ErrorType = 504  // Not enough money on owner's balance
	ErrAdsPermission                     ErrorType = 600  // Permission denied. You have no access to operations specified with given object(s)
	ErrWeightedFlood                     ErrorType = 601  // Permission denied. You have requested too many actions this day. Try later
	ErrAdsPartialSuccess                 ErrorType = 602  // Some part of the request has not been completed
	ErrAdsSpecific                       ErrorType = 603  // Some ads error occurred
	ErrAdsObjectDeleted                  ErrorType = 629  // Object deleted
	ErrGroupChangeCreator                ErrorType = 700  // Cannot edit creator role
	ErrGroupNotInClub                    ErrorType = 701  // User should be in club
	ErrGroupTooManyOfficers              ErrorType = 702  // Too many officers in club
	ErrGroupNeed2fa                      ErrorType = 703  // You need to enable 2FA for this action
	ErrGroupHostNeed2fa                  ErrorType = 704  // User needs to enable 2FA for this action
	ErrGroupTooManyAddresses             ErrorType = 706  // Too many addresses in club
	ErrGroupAppIsNotInstalledInCommunity ErrorType = 711  // "Application is not installed in community
	ErrGroupInvalidInviteLink            ErrorType = 714  // Invite link is invalid - expired, deleted or not exists
	ErrVideoAlreadyAdded                 ErrorType = 800  // This video is already added
	ErrVideoCommentsClosed               ErrorType = 801  // Comments for this video are closed
	ErrMessagesUserBlocked               ErrorType = 900  // Can't send messages for users from blacklist
	ErrMessagesDenySend                  ErrorType = 901  // Can't send messages for users without permission
	ErrMessagesPrivacy                   ErrorType = 902  // Can't send messages to this user due to their privacy settings
	ErrMessagesTooOldPts                 ErrorType = 907  // Value of ts or pts is too old
	ErrMessagesTooNewPts                 ErrorType = 908  // Value of ts or pts is too new
	ErrMessagesEditExpired               ErrorType = 909  // Can't edit this message, because it's too old
	ErrMessagesTooBig                    ErrorType = 910  // Can't sent this message, because it's too big
	ErrMessagesKeyboardInvalid           ErrorType = 911  // Keyboard format is invalid
	ErrMessagesChatBotFeature            ErrorType = 912  // This is a chat bot feature, change this status in settings
	ErrMessagesTooLongForwards           ErrorType = 913  // Too many forwarded messages
	ErrMessagesTooLongMessage            ErrorType = 914  // Message is too long
	ErrMessagesChatUserNoAccess          ErrorType = 917  // You don't have access to this chat
	ErrMessagesCantSeeInviteLink         ErrorType = 919  // You can't see invite link for this chat
	ErrMessagesEditKindDisallowed        ErrorType = 920  // Can't edit this kind of message
	ErrMessagesCantFwd                   ErrorType = 921  // Can't forward these messages
	ErrMessagesCantDeleteForAll          ErrorType = 924  // Can't delete this message for everybody
	ErrMessagesChatNotAdmin              ErrorType = 925  // You are not admin of this chat
	ErrMessagesChatNotExist              ErrorType = 927  // Chat does not exist
	ErrMessagesCantChangeInviteLink      ErrorType = 931  // You can't change invite link for this chat
	ErrMessagesGroupPeerAccess           ErrorType = 932  // Your community can't interact with this peer
	ErrMessagesChatUserNotInChat         ErrorType = 935  // User not found in chat
	ErrMessagesContactNotFound           ErrorType = 936  // Contact not found
	ErrMessagesMessageRequestAlreadySend ErrorType = 939  // Message request already send
	ErrMessagesTooManyPosts              ErrorType = 940  // Too many posts in messages
	ErrMessagesCantUseIntent             ErrorType = 943  // Cannot use this intent
	ErrMessagesLimitIntent               ErrorType = 944  // Limits overflow for this intent
	ErrMessagesChatDisabled              ErrorType = 945  // Chat was disabled
	ErrMessagesChatNotSupported          ErrorType = 946  // Chat not support
	ErrMessagesMemberAccessToGroupDenied ErrorType = 947  // Can't add user to chat, because user has no access to group
	ErrMessagesEditPinned                ErrorType = 949  // Can't edit pinned message yet
	ErrMessagesReplyTimedOut             ErrorType = 950  // Can't send message, reply timed out
	ErrParamPhone                        ErrorType = 1000 // Invalid phone number
	ErrPhoneAlreadyUsed                  ErrorType = 1004 // This phone number is used by another user
	ErrAuthFloodError                    ErrorType = 1105 // Too many auth attempts, try again later
	ErrAuthDelay                         ErrorType = 1112 // Processing.. Try later
	ErrParamDocID                        ErrorType = 1150 // Invalid document id
	ErrParamDocDeleteAccess              ErrorType = 1151 // Access to document deleting is denied
	ErrParamDocTitle                     ErrorType = 1152 // Invalid document title
	ErrParamDocAccess                    ErrorType = 1153 // Access to document is denied
	ErrPhotoChanged                      ErrorType = 1160 // Original photo was changed
	ErrTooManyLists                      ErrorType = 1170 // Too many feed lists
	ErrAppsAlreadyUnlocked               ErrorType = 1251 // This achievement is already unlocked
	ErrAppsSubscriptionNotFound          ErrorType = 1256 // Subscription not found
	ErrAppsSubscriptionInvalidStatus     ErrorType = 1257 // Subscription is in invalid status
	ErrInvalidAddress                    ErrorType = 1260 // Invalid screen name
	ErrCommunitiesCatalogDisabled        ErrorType = 1310 // Catalog is not available for this user
	ErrCommunitiesCategoriesDisabled     ErrorType = 1311 // Catalog categories are not available for this user
	ErrMarketRestoreTooLate              ErrorType = 1400 // Too late for restore
	ErrMarketCommentsClosed              ErrorType = 1401 // Comments for this market are closed
	ErrMarketAlbumNotFound               ErrorType = 1402 // Album not found
	ErrMarketItemNotFound                ErrorType = 1403 // Item not found
	ErrMarketItemAlreadyAdded            ErrorType = 1404 // Item already added to album
	ErrMarketTooManyItems                ErrorType = 1405 // Too many items
	ErrMarketTooManyItemsInAlbum         ErrorType = 1406 // Too many items in album
	ErrMarketTooManyAlbums               ErrorType = 1407 // Too many albums
	ErrMarketItemHasBadLinks             ErrorType = 1408 // Item has bad links in description
	ErrMarketShopNotEnabled              ErrorType = 1409 // Shop not enabled
	ErrMarketCartEmpty                   ErrorType = 1427 // Cart is empty
	ErrMarketSpecifyDimensions           ErrorType = 1429 // Specify width, length, height and weight all together
	ErrStoryExpired                      ErrorType = 1600 // Story has already expired
	ErrStoryIncorrectReplyPrivacy        ErrorType = 1602 // Incorrect reply privacy
	ErrPrettyCardsCardNotFound           ErrorType = 1900 // Card not found
	ErrPrettyCardsTooManyCards           ErrorType = 1901 // Too many cards
	ErrPrettyCardsCardIsConnectedToPost  ErrorType = 1902 // Card is connected to post
	ErrCallbackServersLimit              ErrorType = 2000 // Servers number limit is reached
	ErrRecaptcha                         ErrorType = 3300 // Recaptcha needed
	ErrPhoneValidation                   ErrorType = 3301 // Phone validation needed
	ErrPasswordValidation                ErrorType = 3302 // Password validation needed
	ErrOtpAppValidation                  ErrorType = 3303 // Otp app validation needed
	ErrEmailConfirmation                 ErrorType = 3304 // Email confirmation needed
	ErrAssertVotes                       ErrorType = 3305 // Assert votes
	ErrTokenExtension                    ErrorType = 3609 // Token extension required
	ErrUserDeactivated                   ErrorType = 3610 // User is deactivated
	ErrServiceDeactivated                ErrorType = 3611 // Service is deactivated for user
	ErrAliExpressTag                     ErrorType = 3800 // Can't set AliExpress tag to this type of object
)

// Error struct VK.
type Error struct {
	Code       ErrorType `json:"error_code"`
	Message    string    `json:"error_msg"`
	Text       string    `json:"error_text"`
	CaptchaSID string    `json:"captcha_sid"`
	CaptchaImg string    `json:"captcha_img"`

	// In some cases VK requires to request action confirmation from the user
	// (for Standalone apps only). Following error will be returned:
	//
	// Error code: 24
	// Error text: Confirmation required
	//
	// Following parameter is transmitted with the error message as well:
	//
	// confirmation_text – text of the message to be shown in the default
	// confirmation window.
	//
	// The app should display the default confirmation window with text from
	// confirmation_text and two buttons: "Continue" and "Cancel". If user
	// confirms the action repeat the request with an extra parameter: confirm = 1.
	//
	// See https://vk.com/dev/need_confirmation
	ConfirmationText string `json:"confirmation_text"`

	// In some cases VK requires a user validation procedure. . As a result
	// starting from API version 5.0 (for the older versions captcha_error
	// will be requested) following error will be returned as a reply to any
	// API request:
	//
	// Error code: 17
	// Error text: Validation Required
	//
	// Following parameter is transmitted with an error message:
	// redirect_uri – a special address to open in a browser to pass the
	// validation procedure.
	//
	// After passing the validation a user will be redirected to the service
	// page:
	//
	// https://oauth.vk.com/blank.html#{Data required for validation}
	//
	// In case of successful validation following parameters will be
	// transmitted after #:
	//
	// https://oauth.vk.com/blank.html#success=1&access_token={NEW USER TOKEN}&user_id={USER ID}
	//
	// If a token was not received by https a new secret will be transmitted
	// as well.
	//
	// In case of unsuccessful validation following address is transmitted:
	//
	// https://oauth.vk.com/blank.html#fail=1
	//
	// See https://vk.com/dev/need_validation
	RedirectURI   string                    `json:"redirect_uri"`
	RequestParams []object.BaseRequestParam `json:"request_params"`
}

// Error returns the message of a Error.
func (e Error) Error() string {
	return fmt.Sprintf("api: %s", e.Message)
}

// Is unwraps its first argument sequentially looking for an error that matches
// the second.
func (e Error) Is(target error) bool {
	if t, ok := target.(*Error); ok {
		return e.Code == t.Code && e.Message == t.Message
	}

	if t, ok := target.(ErrorType); ok {
		return e.Code == t
	}

	return false
}

// As unwraps its first argument sequentially looking for an error that can
// be assigned to its second argument, which must be a pointer. If it
// succeeds, it performs the assignment and returns true. Otherwise, it
// returns false.
func (e Error) As(target error) bool {
	if t, ok := target.(Error); ok {
		return e.Code == t.Code
	}

	if t, ok := target.(ErrorType); ok {
		return e.Code == t
	}

	return false
}

// ExecuteError struct.
type ExecuteError struct {
	Method string `json:"method"`
	Code   int    `json:"error_code"`
	Msg    string `json:"error_msg"`
}

// ExecuteErrors type.
type ExecuteErrors []ExecuteError

// Error returns the message of a ExecuteErrors.
func (e ExecuteErrors) Error() string {
	return fmt.Sprintf("api: execute errors (%d)", len(e))
}
