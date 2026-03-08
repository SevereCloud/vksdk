package api

import (
	"errors"
	"fmt"

	"github.com/SevereCloud/vksdk/v3/object"
)

const errorMessagePrefix = "api: "

// ErrorType is the type of an error.
type ErrorType int

// Error returns the message of a ErrorType.
func (e ErrorType) Error() string {
	return fmt.Sprintf(errorMessagePrefix+"error with code %d", e)
}

// Error codes. See https://dev.vk.com/ru/reference/errors
const (
	ErrNoType ErrorType = 0 // NoType error

	// ErrUnknown error occurred.
	//
	// Try again later.
	ErrUnknown ErrorType = 1

	// ErrDisabled application is disabled.
	//
	// Enable your application or use test mode.
	// You need to switch on the app in Settings
	// https://vk.com/editapp?id={Your API_ID}
	// or use the test mode (test_mode=1).
	ErrDisabled ErrorType = 2

	// ErrMethod unknown method passed.
	//
	// Check the method name: http://vk.com/dev/methods
	ErrMethod    ErrorType = 3
	ErrSignature ErrorType = 4 // Incorrect signature

	// ErrAuth user authorization failed.
	//
	// Make sure that you use a correct authorization type.
	ErrAuth ErrorType = 5

	// ErrTooMany requests per second.
	//
	// Decrease the request frequency or use the execute method.
	// More details on frequency limits here:
	// https://dev.vk.com/ru/api/api-requests
	ErrTooMany ErrorType = 6

	// ErrPermission denied to perform this action.
	//
	// Make sure that your have received required permissions during the
	// authorization.
	// You can do it with the account.getAppPermissions method.
	// https://dev.vk.com/ru/reference/access-rights
	ErrPermission ErrorType = 7

	// ErrRequest invalid request.
	//
	// Check the request syntax and used parameters list (it can be found on
	// a method description page).
	ErrRequest ErrorType = 8

	// ErrFlood control.
	//
	// You need to decrease the count of identical requests. For more efficient
	// work you may use execute.
	ErrFlood ErrorType = 9

	// ErrServer internal server error.
	//
	// Try again later.
	ErrServer ErrorType = 10

	// ErrEnabledInTest in test mode application should be disabled or user should be authorized.
	//
	// Switch the app off in Settings:
	//
	// 	https://vk.com/editapp?id={Your API_ID}
	//
	ErrEnabledInTest ErrorType = 11

	// ErrCompile unable to compile code.
	ErrCompile ErrorType = 12

	// ErrRuntime error occurred during code invocation.
	ErrRuntime ErrorType = 13

	// ErrCaptcha is needed.
	//
	// See https://dev.vk.com/ru/api/captcha-error
	ErrCaptcha ErrorType = 14

	// ErrAccess denied.
	//
	// Make sure that you use correct identifiers and the content is available
	// for the user in the full version of the site.
	ErrAccess ErrorType = 15

	// ErrAuthHTTPS HTTP authorization failed.
	//
	// To avoid this error check if a user has the 'Use secure connection'
	// option enabled with the account.getInfo method.
	ErrAuthHTTPS ErrorType = 16

	// ErrAuthValidation required.
	//
	// Make sure that you don't use a token received with
	// http://vk.com/dev/auth_mobile for a request from the server.
	// It's restricted.
	//
	// https://dev.vk.com/ru/api/validation-required-error
	ErrAuthValidation ErrorType = 17
	ErrUserDeleted    ErrorType = 18 // User was deleted or banned
	ErrBlocked        ErrorType = 19 // Content blocked

	// ErrMethodPermission denied for non-standalone applications.
	ErrMethodPermission ErrorType = 20

	// ErrMethodAds permission to perform this action is allowed only for standalone and
	// OpenAPI applications.
	ErrMethodAds ErrorType = 21
	ErrUpload    ErrorType = 22 // Upload error

	// ErrMethodDisabled this method was disabled.
	//
	// All the methods available now are listed here: http://vk.com/dev/methods
	ErrMethodDisabled ErrorType = 23

	// ErrNeedConfirmation confirmation required.
	//
	// In some cases VK requires to request action confirmation from the user
	// (for Standalone apps only).
	//
	// Following parameter is transmitted with the error message as well:
	//
	// confirmation_text – text of the message to be shown in the default
	// confirmation window.
	//
	// The app should display the default confirmation window
	// with text from confirmation_text and two buttons: "Continue" and
	// "Cancel".
	// If user confirms the action repeat the request with an extra parameter:
	//
	// 	confirm = 1.
	//
	// https://dev.vk.com/ru/api/confirmation-required-error
	ErrNeedConfirmation      ErrorType = 24
	ErrNeedTokenConfirmation ErrorType = 25 // Token confirmation required
	ErrGroupAuth             ErrorType = 27 // Group authorization failed
	ErrAppAuth               ErrorType = 28 // Application authorization failed

	// ErrRateLimit reached.
	//
	// More details on rate limits here: https://dev.vk.com/ru/reference/roadmap
	ErrRateLimit      ErrorType = 29
	ErrPrivateProfile ErrorType = 30 // This profile is private

	// ErrClientVersionDeprecated client version deprecated.
	ErrClientVersionDeprecated ErrorType = 34

	// ErrExecutionTimeout method execution was interrupted due to timeout.
	ErrExecutionTimeout ErrorType = 36

	// ErrUserBanned user was banned.
	ErrUserBanned ErrorType = 37

	// ErrUnknownApplication unknown application.
	ErrUnknownApplication ErrorType = 38

	// ErrUnknownUser unknown user.
	ErrUnknownUser ErrorType = 39

	// ErrUnknownGroup unknown group.
	ErrUnknownGroup ErrorType = 40

	// ErrAdditionalSignupRequired additional signup required.
	ErrAdditionalSignupRequired ErrorType = 41

	// ErrIPNotAllowed IP is not allowed.
	ErrIPNotAllowed ErrorType = 42

	// ErrParam one of the parameters specified was missing or invalid.
	//
	// Check the required parameters list and their format on a method
	// description page.
	ErrParam ErrorType = 100

	// ErrParamAPIID invalid application API ID.
	//
	// Find the app in the administrated list in settings:
	// http://vk.com/apps?act=settings
	// And set the correct API_ID in the request.
	ErrParamAPIID   ErrorType = 101
	ErrLimits       ErrorType = 103 // Out of limits
	ErrNotFound     ErrorType = 104 // Not found
	ErrSaveFile     ErrorType = 105 // Couldn't save file
	ErrActionFailed ErrorType = 106 // Unable to process action

	// ErrParamUserID invalid user id.
	//
	// Make sure that you use a correct id. You can get an id using a screen
	// name with the utils.resolveScreenName method.
	ErrParamUserID  ErrorType = 113
	ErrParamAlbumID ErrorType = 114 // Invalid album id
	ErrParamServer  ErrorType = 118 // Invalid server
	ErrParamTitle   ErrorType = 119 // Invalid title
	ErrParamPhotos  ErrorType = 122 // Invalid photos
	ErrParamHash    ErrorType = 121 // Invalid hash
	ErrParamPhoto   ErrorType = 129 // Invalid photo
	ErrParamGroupID ErrorType = 125 // Invalid group id
	ErrParamPageID  ErrorType = 140 // Page not found
	ErrAccessPage   ErrorType = 141 // Access to page denied

	// ErrMobileNotActivated the mobile number of the user is unknown.
	ErrMobileNotActivated ErrorType = 146

	// ErrInsufficientFunds application has insufficient funds.
	ErrInsufficientFunds ErrorType = 147

	// ErrAccessMenu access to the menu of the user denied.
	ErrAccessMenu ErrorType = 148

	// ErrParamTimestamp invalid timestamp.
	//
	// You may get a correct value with the utils.getServerTime method.
	ErrParamTimestamp ErrorType = 150
	ErrFriendsListID  ErrorType = 171 // Invalid list id

	// ErrFriendsListLimit reached the maximum number of lists.
	ErrFriendsListLimit ErrorType = 173

	// ErrFriendsAddYourself cannot add user himself as friend.
	ErrFriendsAddYourself ErrorType = 174

	// ErrFriendsAddInEnemy cannot add this user to friends as they have put you on their blacklist.
	ErrFriendsAddInEnemy ErrorType = 175

	// ErrFriendsAddEnemy cannot add this user to friends as you put him on blacklist.
	ErrFriendsAddEnemy ErrorType = 176

	// ErrFriendsAddNotFound cannot add this user to friends as user not found.
	ErrFriendsAddNotFound ErrorType = 177
	ErrParamNoteID        ErrorType = 180 // Note not found
	ErrAccessNote         ErrorType = 181 // Access to note denied
	ErrAccessNoteComment  ErrorType = 182 // You can't comment this note
	ErrAccessComment      ErrorType = 183 // Access to comment denied

	// ErrAccessAlbum access to album denied.
	//
	// Make sure you use correct ids (owner_id is always positive for users,
	// negative for communities) and the current user has access to the
	// requested content in the full version of the site.
	ErrAccessAlbum ErrorType = 200

	// ErrAccessAudio access to audio denied.
	//
	// Make sure you use correct ids (owner_id is always positive for users,
	// negative for communities) and the current user has access to the
	// requested content in the full version of the site.
	ErrAccessAudio ErrorType = 201

	// ErrAccessGroup access to group denied.
	//
	// Make sure that the current user is a member or admin of the community
	// (for closed and private groups and events).
	ErrAccessGroup ErrorType = 203

	// ErrAccessVideo access denied.
	ErrAccessVideo ErrorType = 204

	// ErrAccessMarket access denied.
	ErrAccessMarket ErrorType = 205

	// ErrWallAccessPost access to wall's post denied.
	ErrWallAccessPost ErrorType = 210

	// ErrWallAccessComment access to wall's comment denied.
	ErrWallAccessComment ErrorType = 211

	// ErrWallAccessReplies access to post comments denied.
	ErrWallAccessReplies ErrorType = 212

	// ErrWallAccessAddReply access to status replies denied.
	ErrWallAccessAddReply ErrorType = 213

	// ErrWallAddPost access to adding post denied.
	ErrWallAddPost ErrorType = 214

	// ErrWallAdsPublished advertisement post was recently added.
	ErrWallAdsPublished ErrorType = 219

	// ErrWallTooManyRecipients too many recipients.
	ErrWallTooManyRecipients ErrorType = 220

	// ErrStatusNoAudio user disabled track name broadcast.
	ErrStatusNoAudio ErrorType = 221

	// ErrWallLinksForbidden hyperlinks are forbidden.
	ErrWallLinksForbidden ErrorType = 222

	// ErrWallReplyOwnerFlood too many replies.
	ErrWallReplyOwnerFlood ErrorType = 223

	// ErrWallAdsPostLimitReached too many ads posts.
	ErrWallAdsPostLimitReached ErrorType = 224

	// ErrDonutDisabled donut is disabled.
	ErrDonutDisabled ErrorType = 225

	// ErrLikesReactionCanNotBeApplied reaction can not be applied to the object.
	ErrLikesReactionCanNotBeApplied ErrorType = 232

	// ErrPollsAccess access to poll denied.
	ErrPollsAccess ErrorType = 250

	// ErrPollsAnswerID invalid answer id.
	ErrPollsAnswerID ErrorType = 252

	// ErrPollsPollID invalid poll id.
	ErrPollsPollID ErrorType = 251

	// ErrPollsAccessWithoutVote access denied, please vote first.
	ErrPollsAccessWithoutVote ErrorType = 253

	// ErrAccessGroups access to the groups list is denied due to the user's privacy settings.
	ErrAccessGroups ErrorType = 260

	// ErrAlbumFull this album is full.
	//
	// You need to delete the odd objects from the album or use another album.
	ErrAlbumFull   ErrorType = 300
	ErrAlbumsLimit ErrorType = 302 // Albums number limit is reached

	// ErrVotesPermission permission denied.
	//
	// You must enable votes processing in application settings.
	// Check the app settings:
	//
	// 	http://vk.com/editapp?id={Your API_ID}&section=payments
	//
	ErrVotesPermission ErrorType = 500

	// ErrVotes not enough votes.
	ErrVotes ErrorType = 503

	// ErrNotEnoughMoney not enough money on owner's balance.
	ErrNotEnoughMoney ErrorType = 504

	// ErrAdsPermission permission denied.
	//
	// You have no access to operations specified with given object(s).
	ErrAdsPermission ErrorType = 600

	// ErrWeightedFlood permission denied.
	//
	// You have requested too many actions this day. Try later.
	ErrWeightedFlood ErrorType = 601

	// ErrAdsPartialSuccess some part of the request has not been completed.
	ErrAdsPartialSuccess ErrorType = 602

	// ErrAdsSpecific some ads error occurred.
	ErrAdsSpecific ErrorType = 603

	// ErrAdsDomainInvalid invalid domain.
	ErrAdsDomainInvalid ErrorType = 604

	// ErrAdsDomainForbidden domain is forbidden.
	ErrAdsDomainForbidden ErrorType = 605

	// ErrAdsDomainReserved domain is reserved.
	ErrAdsDomainReserved ErrorType = 606

	// ErrAdsDomainOccupied domain is occupied.
	ErrAdsDomainOccupied ErrorType = 607

	// ErrAdsDomainActive domain is active.
	ErrAdsDomainActive ErrorType = 608

	// ErrAdsDomainAppInvalid domain app is invalid.
	ErrAdsDomainAppInvalid ErrorType = 609

	// ErrAdsDomainAppForbidden domain app is forbidden.
	ErrAdsDomainAppForbidden ErrorType = 610

	// ErrAdsApplicationMustBeVerified application must be verified.
	ErrAdsApplicationMustBeVerified ErrorType = 611

	// ErrAdsApplicationMustBeInDomainsList application must be in domains list of site of ad unit.
	ErrAdsApplicationMustBeInDomainsList ErrorType = 612

	// ErrAdsApplicationBlocked application is blocked.
	ErrAdsApplicationBlocked ErrorType = 613

	// ErrAdsDomainTypeForbiddenInCurrentOffice domain of type specified is forbidden in current office type.
	ErrAdsDomainTypeForbiddenInCurrentOffice ErrorType = 614

	// ErrAdsDomainGroupInvalid domain group is invalid.
	ErrAdsDomainGroupInvalid ErrorType = 615

	// ErrAdsDomainGroupForbidden domain group is forbidden.
	ErrAdsDomainGroupForbidden ErrorType = 616

	// ErrAdsDomainAppBlocked domain app is blocked.
	ErrAdsDomainAppBlocked ErrorType = 617

	// ErrAdsDomainGroupNotOpen domain group is not open.
	ErrAdsDomainGroupNotOpen ErrorType = 618

	// ErrAdsDomainGroupNotPossibleJoined domain group is not possible to be joined to adsweb.
	ErrAdsDomainGroupNotPossibleJoined ErrorType = 619

	// ErrAdsDomainGroupBlocked domain group is blocked.
	ErrAdsDomainGroupBlocked ErrorType = 620

	// ErrAdsDomainGroupLinksForbidden domain group has restriction: links are forbidden.
	ErrAdsDomainGroupLinksForbidden ErrorType = 621

	// ErrAdsDomainGroupExcludedFromSearch domain group has restriction: excluded from search.
	ErrAdsDomainGroupExcludedFromSearch ErrorType = 622

	// ErrAdsDomainGroupCoverForbidden domain group has restriction: cover is forbidden.
	ErrAdsDomainGroupCoverForbidden ErrorType = 623

	// ErrAdsDomainGroupWrongCategory domain group has wrong category.
	ErrAdsDomainGroupWrongCategory ErrorType = 624

	// ErrAdsDomainGroupWrongName domain group has wrong name.
	ErrAdsDomainGroupWrongName ErrorType = 625

	// ErrAdsDomainGroupLowPostsReach domain group has low posts reach.
	ErrAdsDomainGroupLowPostsReach ErrorType = 626

	// ErrAdsDomainGroupWrongClass domain group has wrong class.
	ErrAdsDomainGroupWrongClass ErrorType = 627

	// ErrAdsDomainGroupCreatedRecently domain group is created recently.
	ErrAdsDomainGroupCreatedRecently ErrorType = 628

	// ErrAdsObjectDeleted object deleted.
	ErrAdsObjectDeleted ErrorType = 629

	// ErrAdsLookalikeRequestAlreadyInProgress lookalike request with same source already in progress.
	ErrAdsLookalikeRequestAlreadyInProgress ErrorType = 630

	// ErrAdsLookalikeRequestsLimit max count of lookalike requests per day reached.
	ErrAdsLookalikeRequestsLimit ErrorType = 631

	// ErrAdsAudienceTooSmall given audience is too small.
	ErrAdsAudienceTooSmall ErrorType = 632

	// ErrAdsAudienceTooLarge given audience is too large.
	ErrAdsAudienceTooLarge ErrorType = 633

	// ErrAdsLookalikeAudienceSaveAlreadyInProgress lookalike request audience save already in progress.
	ErrAdsLookalikeAudienceSaveAlreadyInProgress ErrorType = 634

	// ErrAdsLookalikeSavesLimit max count of lookalike request audience saves per day reached.
	ErrAdsLookalikeSavesLimit ErrorType = 635

	// ErrAdsRetargetingGroupsLimit max count of retargeting groups reached.
	ErrAdsRetargetingGroupsLimit ErrorType = 636

	// ErrAdsDomainGroupActiveNemesisPunishment domain group has active nemesis punishment.
	ErrAdsDomainGroupActiveNemesisPunishment ErrorType = 637

	// ErrGroupChangeCreator cannot edit creator role.
	ErrGroupChangeCreator ErrorType = 700

	// ErrGroupNotInClub user should be in club.
	ErrGroupNotInClub ErrorType = 701

	// ErrGroupTooManyOfficers too many officers in club.
	ErrGroupTooManyOfficers ErrorType = 702

	// ErrGroupNeed2fa you need to enable 2FA for this action.
	ErrGroupNeed2fa ErrorType = 703

	// ErrGroupHostNeed2fa user needs to enable 2FA for this action.
	ErrGroupHostNeed2fa ErrorType = 704

	// ErrGroupTooManyAddresses too many addresses in club.
	ErrGroupTooManyAddresses ErrorType = 706

	// ErrGroupAppIsNotInstalledInCommunity application is not installed in community.
	ErrGroupAppIsNotInstalledInCommunity ErrorType = 711

	// ErrGroupInvalidInviteLink invite link is invalid - expired, deleted or not exists.
	ErrGroupInvalidInviteLink ErrorType = 714

	// ErrVideoAlreadyAdded this video is already added.
	ErrVideoAlreadyAdded ErrorType = 800

	// ErrVideoCommentsClosed comments for this video are closed.
	ErrVideoCommentsClosed ErrorType = 801

	// ErrMessagesUserBlocked can't send messages for users from blacklist.
	ErrMessagesUserBlocked ErrorType = 900

	// ErrMessagesDenySend can't send messages for users without permission.
	ErrMessagesDenySend ErrorType = 901

	// ErrMessagesPrivacy can't send messages to this user due to their privacy settings.
	ErrMessagesPrivacy ErrorType = 902

	// ErrMessagesTooOldPts value of ts or pts is too old.
	ErrMessagesTooOldPts ErrorType = 907

	// ErrMessagesTooNewPts value of ts or pts is too new.
	ErrMessagesTooNewPts ErrorType = 908

	// ErrMessagesEditExpired can't edit this message, because it's too old.
	ErrMessagesEditExpired ErrorType = 909

	// ErrMessagesTooBig can't sent this message, because it's too big.
	ErrMessagesTooBig ErrorType = 910

	// ErrMessagesKeyboardInvalid keyboard format is invalid.
	ErrMessagesKeyboardInvalid ErrorType = 911

	// ErrMessagesChatBotFeature this is a chat bot feature, change this status in settings.
	ErrMessagesChatBotFeature ErrorType = 912

	// ErrMessagesTooLongForwards too many forwarded messages.
	ErrMessagesTooLongForwards ErrorType = 913

	// ErrMessagesTooLongMessage message is too long.
	ErrMessagesTooLongMessage ErrorType = 914

	// ErrMessagesChatUserNoAccess you don't have access to this chat.
	ErrMessagesChatUserNoAccess ErrorType = 917

	// ErrMessagesCantSeeInviteLink you can't see invite link for this chat.
	ErrMessagesCantSeeInviteLink ErrorType = 919

	// ErrMessagesEditKindDisallowed can't edit this kind of message.
	ErrMessagesEditKindDisallowed ErrorType = 920

	// ErrMessagesCantFwd can't forward these messages.
	ErrMessagesCantFwd ErrorType = 921

	// ErrMessagesCantDeleteForAll can't delete this message for everybody.
	ErrMessagesCantDeleteForAll ErrorType = 924

	// ErrMessagesChatNotAdmin you are not admin of this chat.
	ErrMessagesChatNotAdmin ErrorType = 925

	// ErrMessagesChatNotExist chat does not exist.
	ErrMessagesChatNotExist ErrorType = 927

	// ErrMessagesCantChangeInviteLink you can't change invite link for this chat.
	ErrMessagesCantChangeInviteLink ErrorType = 931

	// ErrMessagesGroupPeerAccess your community can't interact with this peer.
	ErrMessagesGroupPeerAccess ErrorType = 932

	// ErrMessagesChatUserNotInChat user not found in chat.
	ErrMessagesChatUserNotInChat ErrorType = 935

	// ErrMessagesContactNotFound contact not found.
	ErrMessagesContactNotFound ErrorType = 936

	// ErrMessagesMessageRequestAlreadySend message request already send.
	ErrMessagesMessageRequestAlreadySend ErrorType = 939

	// ErrMessagesTooManyPosts too many posts in messages.
	ErrMessagesTooManyPosts ErrorType = 940

	// ErrMessagesCantPinOneTimeStory cannot pin one-time story.
	ErrMessagesCantPinOneTimeStory ErrorType = 942

	// ErrMessagesCantUseIntent cannot use this intent.
	ErrMessagesCantUseIntent ErrorType = 943

	// ErrMessagesLimitIntent limits overflow for this intent.
	ErrMessagesLimitIntent ErrorType = 944

	// ErrMessagesChatDisabled chat was disabled.
	ErrMessagesChatDisabled ErrorType = 945

	// ErrMessagesChatNotSupported chat not support.
	ErrMessagesChatNotSupported ErrorType = 946

	// ErrMessagesMemberAccessToGroupDenied can't add user to chat, because user has no access to group.
	ErrMessagesMemberAccessToGroupDenied ErrorType = 947

	// ErrMessagesEditPinned can't edit pinned message yet.
	ErrMessagesEditPinned ErrorType = 949

	// ErrMessagesReplyTimedOut can't send message, reply timed out.
	ErrMessagesReplyTimedOut ErrorType = 950

	// ErrMessagesAccessDonutChat you can't access donut chat without subscription.
	ErrMessagesAccessDonutChat ErrorType = 962

	// ErrMessagesAccessWorkChat this user can't be added to the work chat, as they aren't an employe.
	ErrMessagesAccessWorkChat ErrorType = 967

	// ErrMessagesCantForwarded message cannot be forwarded.
	ErrMessagesCantForwarded ErrorType = 969

	// ErrMessagesPinExpiringMessage cannot pin an expiring message.
	ErrMessagesPinExpiringMessage ErrorType = 970

	// ErrParamPhone invalid phone number.
	ErrParamPhone ErrorType = 1000

	// ErrPhoneAlreadyUsed this phone number is used by another user.
	ErrPhoneAlreadyUsed ErrorType = 1004

	// ErrAuthFloodError too many auth attempts, try again later.
	ErrAuthFloodError ErrorType = 1105

	// ErrAuthDelay processing.. Try later.
	ErrAuthDelay ErrorType = 1112

	// ErrAnonymousTokenExpired anonymous token has expired.
	ErrAnonymousTokenExpired ErrorType = 1114

	// ErrAnonymousTokenInvalid anonymous token is invalid.
	ErrAnonymousTokenInvalid ErrorType = 1116

	// ErrAuthAccessTokenHasExpired access token has expired.
	ErrAuthAccessTokenHasExpired ErrorType = 1117

	// ErrAuthAnonymousTokenIPMismatch anonymous token ip mismatch.
	ErrAuthAnonymousTokenIPMismatch ErrorType = 1118

	// ErrParamDocID invalid document id.
	ErrParamDocID ErrorType = 1150

	// ErrParamDocDeleteAccess access to document deleting is denied.
	ErrParamDocDeleteAccess ErrorType = 1151

	// ErrParamDocTitle invalid document title.
	ErrParamDocTitle ErrorType = 1152

	// ErrParamDocAccess access to document is denied.
	ErrParamDocAccess ErrorType = 1153

	// ErrPhotoChanged original photo was changed.
	ErrPhotoChanged ErrorType = 1160

	// ErrTooManyLists too many feed lists.
	ErrTooManyLists ErrorType = 1170

	// ErrAppsAlreadyUnlocked this achievement is already unlocked.
	ErrAppsAlreadyUnlocked ErrorType = 1251

	// ErrAppsSubscriptionNotFound subscription not found.
	ErrAppsSubscriptionNotFound ErrorType = 1256

	// ErrAppsSubscriptionInvalidStatus subscription is in invalid status.
	ErrAppsSubscriptionInvalidStatus ErrorType = 1257

	// ErrInvalidAddress invalid screen name.
	ErrInvalidAddress ErrorType = 1260

	// ErrCommunitiesCatalogDisabled catalog is not available for this user.
	ErrCommunitiesCatalogDisabled ErrorType = 1310

	// ErrCommunitiesCategoriesDisabled catalog categories are not available for this user.
	ErrCommunitiesCategoriesDisabled ErrorType = 1311

	// ErrMarketRestoreTooLate too late for restore.
	ErrMarketRestoreTooLate ErrorType = 1400

	// ErrMarketCommentsClosed comments for this market are closed.
	ErrMarketCommentsClosed ErrorType = 1401

	// ErrMarketAlbumNotFound album not found.
	ErrMarketAlbumNotFound ErrorType = 1402

	// ErrMarketItemNotFound item not found.
	ErrMarketItemNotFound ErrorType = 1403

	// ErrMarketItemAlreadyAdded item already added to album.
	ErrMarketItemAlreadyAdded ErrorType = 1404

	// ErrMarketTooManyItems too many items.
	ErrMarketTooManyItems ErrorType = 1405

	// ErrMarketTooManyItemsInAlbum too many items in album.
	ErrMarketTooManyItemsInAlbum ErrorType = 1406

	// ErrMarketTooManyAlbums too many albums.
	ErrMarketTooManyAlbums ErrorType = 1407

	// ErrMarketItemHasBadLinks item has bad links in description.
	ErrMarketItemHasBadLinks ErrorType = 1408

	// ErrMarketShopNotEnabled extended market not enabled.
	ErrMarketShopNotEnabled ErrorType = 1409

	// ErrMarketGroupingItemsWithDifferentProperties grouping items with different properties.
	ErrMarketGroupingItemsWithDifferentProperties ErrorType = 1412

	// ErrMarketGroupingAlreadyHasSuchVariant grouping already has such variant.
	ErrMarketGroupingAlreadyHasSuchVariant ErrorType = 1413

	// ErrMarketVariantNotFound variant not found.
	ErrMarketVariantNotFound ErrorType = 1416

	// ErrMarketPropertyNotFound property not found.
	ErrMarketPropertyNotFound ErrorType = 1417

	// ErrMarketGroupingMustContainMoreThanOneItem grouping must have two or more items.
	ErrMarketGroupingMustContainMoreThanOneItem ErrorType = 1425

	// ErrMarketGroupingItemsMustHaveDistinctProperties item must have distinct properties.
	ErrMarketGroupingItemsMustHaveDistinctProperties ErrorType = 1426

	// ErrMarketOrdersNoCartItems cart is empty.
	ErrMarketOrdersNoCartItems ErrorType = 1427

	// ErrMarketInvalidDimensions specify width, length, height and weight all together.
	ErrMarketInvalidDimensions ErrorType = 1429

	// ErrMarketCantChangeVkpayStatus VK Pay status can not be changed.
	ErrMarketCantChangeVkpayStatus ErrorType = 1430

	// ErrMarketShopAlreadyEnabled market was already enabled in this group.
	ErrMarketShopAlreadyEnabled ErrorType = 1431

	// ErrMarketShopAlreadyDisabled market was already disabled in this group.
	ErrMarketShopAlreadyDisabled ErrorType = 1432

	// ErrMarketPhotosCropInvalidFormat invalid image crop format.
	ErrMarketPhotosCropInvalidFormat ErrorType = 1433

	// ErrMarketPhotosCropOverflow crop bottom right corner is outside of the image.
	ErrMarketPhotosCropOverflow ErrorType = 1434

	// ErrMarketPhotosCropSizeTooLow crop size is less than the minimum.
	ErrMarketPhotosCropSizeTooLow ErrorType = 1435

	// ErrMarketNotEnabled market not enabled.
	ErrMarketNotEnabled ErrorType = 1438

	// ErrMarketCartEmpty cart is empty.
	ErrMarketCartEmpty ErrorType = 1427

	// ErrMarketSpecifyDimensions specify width, length, height and weight all together.
	ErrMarketSpecifyDimensions ErrorType = 1429

	// ErrVKPayStatus VK Pay status can not be changed.
	ErrVKPayStatus ErrorType = 1430

	// ErrMarketAlreadyEnabled market was already enabled in this group.
	ErrMarketAlreadyEnabled ErrorType = 1431

	// ErrMarketAlreadyDisabled market was already disabled in this group.
	ErrMarketAlreadyDisabled ErrorType = 1432

	// ErrMainAlbumCantHidden main album can not be hidden.
	ErrMainAlbumCantHidden ErrorType = 1446

	// ErrStoryExpired story has already expired.
	ErrStoryExpired ErrorType = 1600

	// ErrStoryIncorrectReplyPrivacy incorrect reply privacy.
	ErrStoryIncorrectReplyPrivacy ErrorType = 1602

	// ErrPrettyCardsCardNotFound card not found.
	ErrPrettyCardsCardNotFound ErrorType = 1900

	// ErrPrettyCardsTooManyCards too many cards.
	ErrPrettyCardsTooManyCards ErrorType = 1901

	// ErrPrettyCardsCardIsConnectedToPost card is connected to post.
	ErrPrettyCardsCardIsConnectedToPost ErrorType = 1902

	// ErrCallbackServersLimit servers number limit is reached.
	ErrCallbackServersLimit ErrorType = 2000

	// ErrStickersNotPurchased stickers are not purchased.
	ErrStickersNotPurchased ErrorType = 2100

	// ErrStickersTooManyFavorites too many favorite stickers.
	ErrStickersTooManyFavorites ErrorType = 2101

	// ErrStickersNotFavorite stickers are not favorite.
	ErrStickersNotFavorite ErrorType = 2102

	// ErrWallCheckLinkCantDetermineSource specified link is incorrect (can't find source).
	ErrWallCheckLinkCantDetermineSource ErrorType = 3102

	// ErrRecaptcha recaptcha needed.
	ErrRecaptcha ErrorType = 3300

	// ErrPhoneValidation phone validation needed.
	ErrPhoneValidation ErrorType = 3301

	// ErrPasswordValidation password validation needed.
	ErrPasswordValidation ErrorType = 3302

	// ErrOtpAppValidation otp app validation needed.
	ErrOtpAppValidation ErrorType = 3303

	// ErrEmailConfirmation email confirmation needed.
	ErrEmailConfirmation ErrorType = 3304

	// ErrAssertVotes assert votes.
	ErrAssertVotes ErrorType = 3305

	// ErrTokenExtension token extension required.
	ErrTokenExtension ErrorType = 3609

	// ErrUserDeactivated user is deactivated.
	ErrUserDeactivated ErrorType = 3610

	// ErrServiceDeactivated service is deactivated for user.
	ErrServiceDeactivated ErrorType = 3611

	// ErrAliExpressTag can't set AliExpress tag to this type of object.
	ErrAliExpressTag ErrorType = 3800

	// ErrInvalidUploadResponse invalid upload response.
	ErrInvalidUploadResponse ErrorType = 5701

	// ErrInvalidUploadHash invalid upload hash.
	ErrInvalidUploadHash ErrorType = 5702

	// ErrInvalidUploadUser invalid upload user.
	ErrInvalidUploadUser ErrorType = 5703

	// ErrInvalidUploadGroup invalid upload group.
	ErrInvalidUploadGroup ErrorType = 5704

	// ErrInvalidCropData invalid crop data.
	ErrInvalidCropData ErrorType = 5705

	// ErrToSmallAvatar to small avatar.
	ErrToSmallAvatar ErrorType = 5706

	// ErrPhotoNotFound photo not found.
	ErrPhotoNotFound ErrorType = 5708

	// ErrInvalidPhoto invalid Photo.
	ErrInvalidPhoto ErrorType = 5709

	// ErrInvalidHash invalid hash.
	ErrInvalidHash ErrorType = 5710
)

// ErrorSubtype is the subtype of an error.
type ErrorSubtype int

// Error returns the message of a ErrorSubtype.
func (e ErrorSubtype) Error() string {
	return fmt.Sprintf(errorMessagePrefix+"error with subcode %d", e)
}

// Error struct VK.
type Error struct {
	Code       ErrorType    `json:"error_code"`
	Subcode    ErrorSubtype `json:"error_subcode"`
	Message    string       `json:"error_msg"`
	Text       string       `json:"error_text"`
	CaptchaSID string       `json:"captcha_sid"`
	CaptchaImg string       `json:"captcha_img"`

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
	// confirms the action repeat the request with an extra parameter:
	// confirm = 1.
	//
	// See https://dev.vk.com/ru/api/confirmation-required-error
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
	// See https://dev.vk.com/ru/api/validation-required-error
	RedirectURI   string                    `json:"redirect_uri"`
	RequestParams []object.BaseRequestParam `json:"request_params"`
}

// Error returns the message of a Error.
func (e Error) Error() string {
	return errorMessagePrefix + e.Message
}

// Is unwraps its first argument sequentially looking for an error that matches
// the second.
func (e Error) Is(target error) bool {
	var tError *Error
	if errors.As(target, &tError) {
		return e.Code == tError.Code && e.Message == tError.Message
	}

	var tErrorType ErrorType
	if errors.As(target, &tErrorType) {
		return e.Code == tErrorType
	}

	return false
}

// ExecuteError struct.
type ExecuteError struct {
	Method string    `json:"method"`
	Code   ErrorType `json:"error_code"`
	Msg    string    `json:"error_msg"`
}

// ExecuteErrors type.
type ExecuteErrors []ExecuteError

// Error returns the message of a ExecuteErrors.
func (e ExecuteErrors) Error() string {
	return fmt.Sprintf("api: execute errors (%d)", len(e))
}

// InvalidContentType type.
type InvalidContentType struct {
	ContentType string
}

// Error returns the message of a InvalidContentType.
func (e InvalidContentType) Error() string {
	if e.ContentType == "" {
		return "api: empty content-type"
	}

	return fmt.Sprintf("api: invalid content-type(%s)", e.ContentType)
}

// UploadError type.
type UploadError struct {
	Err      string `json:"error"`
	Code     int    `json:"error_code"`
	Descr    string `json:"error_descr"`
	IsLogged bool   `json:"error_is_logged"`
}

// Error returns the message of a UploadError.
func (e UploadError) Error() string {
	if e.Err != "" {
		return errorMessagePrefix + e.Err
	}

	return fmt.Sprintf("api: upload code %d", e.Code)
}

// AdsError struct.
type AdsError struct {
	Code ErrorType `json:"error_code"`
	Desc string    `json:"error_desc"`
}

// Error returns the message of a AdsError.
func (e AdsError) Error() string {
	return errorMessagePrefix + e.Desc
}

// Is unwraps its first argument sequentially looking for an error that matches
// the second.
func (e AdsError) Is(target error) bool {
	var tAdsError *AdsError
	if errors.As(target, &tAdsError) {
		return e.Code == tAdsError.Code && e.Desc == tAdsError.Desc
	}

	var tErrorType ErrorType
	if errors.As(target, &tErrorType) {
		return e.Code == tErrorType
	}

	return false
}

// AuthSilentTokenError struct.
type AuthSilentTokenError struct {
	Token       string    `json:"token"`
	Code        ErrorType `json:"code"`
	Description string    `json:"description"`
}

// Error returns the description of a AuthSilentTokenError.
func (e AuthSilentTokenError) Error() string {
	return errorMessagePrefix + e.Description
}

// Is unwraps its first argument sequentially looking for an error that matches
// the second.
func (e AuthSilentTokenError) Is(target error) bool {
	var tError *AuthSilentTokenError
	if errors.As(target, &tError) {
		return e.Code == tError.Code && e.Description == tError.Description
	}

	var tErrorType ErrorType
	if errors.As(target, &tErrorType) {
		return e.Code == tErrorType
	}

	return false
}
