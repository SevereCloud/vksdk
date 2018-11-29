package vksdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	version = "5.92"
	apiURL  = "https://api.vk.com/method/"
)

type VK struct {
	AccessToken string
	Version     string
}

// Type error VK
type Error struct {
	Code          int    `json:"error_code"`
	Message       string `json:"error_msg"`
	RequestParams []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"request_params"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code %d: %s", e.Code, e.Message)
}

// Init VK api
func Init(token string) VK {
	vk := VK{}
	vk.AccessToken = token
	vk.Version = version
	return vk
}

// Request provides access to VK API methods
func (vk *VK) Request(method string, params map[string]string) ([]byte, error) {
	// TODO ограничитель на запросы
	u, err := url.Parse(apiURL + method)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	for key, value := range params {
		query.Set(key, value)
	}
	query.Set("access_token", vk.AccessToken)
	query.Set("v", vk.Version)
	u.RawQuery = query.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var handler struct {
		Error    *Error
		Response json.RawMessage
	}
	err = json.Unmarshal(body, &handler)

	if handler.Error != nil {
		return nil, handler.Error
	}

	return handler.Response, nil

}

// TODO account.ban

// TODO account.changePassword Changes a user password after access is successfully restored with the auth.restore method.

// TODO account.getActiveOffers Returns a list of active ads (offers).If the user fulfill their conditions, he will be able to get the appropriate number of votes to his balance.

// TODO account.getAppPermissions Gets settings of the user in this application.

// TODO account.getBanned Returns a user's blacklist.

// TODO account.getCounters Returns non-null values of user counters.

// TODO account.getInfo Returns current account info.

// TODO account.getProfileInfo Returns the current account info.

// TODO account.getPushSettings Gets settings of push notifications.

// TODO account.registerDevice Subscribes an iOS/Android/Windows/Mac based device to receive push notifications

// TODO account.saveProfileInfo Edits current profile info.

// TODO account.setInfo Allows to edit the current account info.

// TODO account.setNameInMenu Sets an application screen name (up to 17 characters), that is shown to the user in the left menu.

// TODO account.setOffline Marks a current user as offline.

// TODO account.setOnline Marks the current user as online for 5 minutes.

// TODO account.setPushSettings Change push settings.

// TODO account.setSilenceMode Mutes push notifications for the set period of time.

// TODO account.unban

// TODO account.unregisterDevice Unsubscribes a device from push notifications.

// TODO appWidgets.getAppImageUploadServer Returns a URL for uploading a photo to the app collection for community app widgets.

// TODO appWidgets.getAppImages Returns an app collection of images for community app widgets.

// TODO appWidgets.getGroupImageUploadServer Returns a URL for uploading a photo to the community collection for community app widgets.

// TODO appWidgets.getGroupImages Returns a community collection of images for community app widgets.

// TODO appWidgets.getImagesById Returns an image for community app widgets by its ID.

// TODO appWidgets.saveAppImage Allows to save image into app collection for community app widgets.

// TODO appWidgets.saveGroupImage Allows to save image into community collection for community app widgets.

// TODO appWidgets.update Allows to update community app widget.

// TODO apps.deleteAppRequests Deletes all request notifications from the current app.

// TODO apps.get Returns applications data.

// TODO apps.getCatalog Returns a list of applications (apps) available to users in the App Catalog.

// TODO apps.getFriendsList Creates friends list for requests and invites in current app.

// TODO apps.getLeaderboard Returns players rating in the game.

// TODO apps.getScopes

// TODO apps.getScore Returns user score in app.

// TODO apps.sendRequest Sends a request to another user in an app that uses VK authorization.

// TODO auth.checkPhone Checks a user's phone number for correctness.

// TODO auth.restore Allows to restore account access using a code received via SMS.

// TODO board.addTopic Creates a new topic on a community's discussion board.

// TODO board.closeTopic Closes a topic on a community's discussion board so that comments cannot be posted.

// TODO board.createComment Adds a comment on a topic on a community's discussion board.

// TODO board.deleteComment Deletes a comment on a topic on a community's discussion board.

// TODO board.deleteTopic Deletes a topic from a community's discussion board.

// TODO board.editComment Edits a comment on a topic on a community's discussion board.

// TODO board.editTopic Edits the title of a topic on a community's discussion board.

// TODO board.fixTopic Pins a topic (fixes its place) to the top of a community's discussion board.

// TODO board.getComments Returns a list of comments on a topic on a community's discussion board.

// TODO board.getTopics Returns a list of topics on a community's discussion board.

// TODO board.openTopicRe-opens a previously closed topic on a community's discussion board.

// TODO board.restoreComment Restores a comment deleted from a topic on a community's discussion board.

// TODO board.unfixTopic Unpins a pinned topic from the top of a community's discussion board.

// TODO database.getChairs Returns list of chairs on a specified faculty.

// TODO database.getCities Returns a list of cities.
func (vk *VK) DatabaseGetCities(params map[string]string) (DatabaseGetCitiesResponse, error) {
	var response DatabaseGetCitiesResponse

	rawResponse, err := vk.Request("database.getCities", params)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}

// TODO database.getCitiesById Returns information about cities by their IDs.

// TODO database.getCountries Returns a list of countries.

// TODO database.getCountriesById Returns information about countries by their IDs.

// TODO database.getFaculties Returns a list of faculties (i.e., university departments).

// TODO database.getRegions Returns a list of regions.

// TODO database.getSchoolClasses Returns a list of school classes specified for the country.

// TODO database.getSchools Returns a list of schools.

// TODO database.getUniversities Returns a list of higher education institutions.

// TODO docs.add Copies a document to a user's or community's document list.

// TODO docs.delete Deletes a user or community document.

// TODO docs.edit Edits a document.

// TODO docs.get Returns detailed information about user or community documents.

// TODO docs.getById Returns information about documents by their IDs.

// TODO docs.getMessagesUploadServer Returns the server address for document upload.

// TODO docs.getTypes Returns documents types available for current user.

// TODO docs.getUploadServer Returns the server address for document upload.

// TODO docs.getWallUploadServer Returns the server address for document upload onto a user's or community's wall.

// TODO docs.save Saves a document after uploading it to a server.

// TODO docs.search Returns a list of documents matching the search criteria.

// TODO fave.addGroup Adds a community to user faves.

// TODO fave.addLink Adds a link to user faves.

// TODO fave.addUser Adds a profile to user faves.

// TODO fave.getLinks Returns a list of links that the current user has bookmarked.

// TODO fave.getMarketItems Returns market items bookmarked by current user.

// TODO fave.getPhotos Returns a list of photos that the current user has liked.

// TODO fave.getPosts Returns a list of wall posts that the current user has liked.

// TODO fave.getUsers Returns a list of users whom the current user has bookmarked.

// TODO fave.getVideos Returns a list of videos that the current user has liked.

// TODO fave.removeGroup Removes a community from user faves.

// TODO fave.removeLink Removes link from the user's faves.

// TODO fave.removeUser Removes a profile from user faves.

// TODO friends.add Approves or creates a friend request.

// TODO friends.addList Creates a new friend list for the current user.

// TODO friends.areFriends Checks the current user's friendship status with other specified users.

// TODO friends.delete Declines a friend request or deletes a user from the current user's friend list.

// TODO friends.deleteAllRequests Marks all incoming friend requests as viewed.

// TODO friends.deleteList Deletes a friend list of the current user.

// TODO friends.edit Edits the friend lists of the selected user.

// TODO friends.editList Edits a friend list of the current user.

// TODO friends.get Returns a list of user IDs or detailed information about a user's friends.
func (vk *VK) FriendsGet(params map[string]string) (FriendsGetResponse, error) {
	var response FriendsGetResponse

	rawResponse, err := vk.Request("friends.get", params)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}

// TODO friends.getAppUsers Returns a list of IDs of the current user's friends who installed the application.

// TODO friends.getByPhones Returns a list of the current user's friends whose phone numbers, validated or specified in a profile, are in a given list.

// TODO friends.getLists Returns a list of the user's friend lists.

// TODO friends.getMutual Returns a list of user IDs of the mutual friends of two users.

// TODO friends.getOnline Returns a list of user IDs of a user's friends who are online.

// TODO friends.getRecent Returns a list of user IDs of the current user's recently added friends.

// TODO friends.getRequests Returns information about the current user's incoming and outgoing friend requests.

// TODO friends.getSuggestions Returns a list of profiles of users whom the current user may know.

// TODO friends.search Returns a list of friends matching the search criteria.

// TODO gifts.get Returns a list of user gifts.

// TODO groups.addCallbackServerAdds Callback API server to the community.

// TODO groups.addLink Allows to add a link to the community.

// TODO groups.approveRequest Allows to approve join request to the community.

// TODO groups.ban Adds a user or a group to the community blacklist.

// TODO groups.create Creates a new community.

// TODO groups.deleteCallbackServerDeletes Callback API server from the community.

// TODO groups.deleteLink Allows to delete a link from the community.

// TODO groups.disableOnline Disables "online" status in the community.

// TODO groups.edit Edits a community.

// TODO groups.editCallbackServer Edits Callback API server in the community.

// TODO groups.editLink Allows to edit a link in the community.

// TODO groups.editManager Allows to add, remove or edit the community manager .

// TODO groups.editPlace Edits the place in community.

// TODO groups.enableOnline Enables "online" status in the community.

// TODO groups.get Returns a list of the communities to which a user belongs.

// TODO groups.getAddresses

// TODO groups.getBanned Returns a list of users on a community blacklist.

// TODO groups.getById Returns information about communities by their IDs.

// TODO groups.getCallbackConfirmationCode Returns Callback API confirmation code for the community.

// TODO groups.getCallbackServers Receives a list of Callback API servers from the community.

// TODO groups.getCallbackSettings Returns Callback API notifications settings.

// TODO groups.getCatalog Returns communities list for a catalog category.

// TODO groups.getCatalogInfo Returns categories list for communities catalog

// TODO groups.getInvitedUsers Returns invited users list of a community

// TODO groups.getInvites Returns a list of invitations to join communities and events.

// TODO groups.getLongPollServer Returns data for Bots Long Poll API connection.

// TODO groups.getLongPollSettings Returns Bots Long Poll API settings.

// TODO groups.getMembers Returns a list of community members.
func (vk *VK) GroupsGetMembers(params map[string]string) (GroupsGetMembersResponse, error) {
	var response GroupsGetMembersResponse

	rawResponse, err := vk.Request("groups.getMembers", params)
	if err != nil {
		return response, err
	}
	// FIXME list if no filter
	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}

// TODO groups.getOnlineStatus Returns a community's online status.

// TODO groups.getRequests Returns a list of requests to the community.

// TODO groups.getSettings Returns community settings.

// TODO groups.getTokenPermissions Returns permissions scope for the community's access_token.

// TODO groups.invite Allows to invite friends to the community.

// TODO groups.isMember Returns information specifying whether a user is a member of a community.

// TODO groups.join With this method you can join the group or public page, and also confirm your participation in an event.

// TODO groups.leave With this method you can leave a group, public page, or event.

// TODO groups.removeUser Removes a user from the community.

// TODO groups.reorderLink Allows to reorder links in the community.

// TODO groups.search Returns a list of communities matching the search criteria.

// TODO groups.setCallbackSettings Allow to set notifications settings for Callback API.

// TODO groups.setLongPollSettings Allows to set Bots Long Poll API settings in the community.

// TODO groups.unban

// TODO leadForms.create

// TODO leadForms.delete

// TODO leadForms.get

// TODO leadForms.getLeads

// TODO leadForms.getUploadURL

// TODO leadForms.list

// TODO leadForms.update

// TODO likes.add Adds the specified object to the Likes list of the current user.

// TODO likes.delete Deletes the specified object from the Likes list of the current user.

// TODO likes.getList Returns a list of IDs of users who added the specified object to their Likes list.

// TODO likes.isLiked Checks for the object in the Likes list of the specified user.

// TODO market.add Adds a new item to the market.

// TODO market.addAlbum Creates new collection of items

// TODO market.addToAlbum Adds an item to one or multiple collections.

// TODO market.createComment Creates a new comment for an item.

// TODO market.delete Deletes an item.

// TODO market.deleteAlbum Deletes a collection of items.

// TODO market.deleteComment Deletes an item's comment

// TODO market.edit Edits an item.

// TODO market.editAlbum Edits a collection of items

// TODO market.editComment Changes item comment's text

// TODO market.get Returns items list for a community.

// TODO market.getAlbumById Returns items album's data

// TODO market.getAlbums Returns community's collections list.

// TODO market.getById Returns information about market items by their ids.

// TODO market.getCategories Returns a list of market categories.

// TODO market.getComments Returns comments list for an item.

// TODO market.removeFromAlbum Removes an item from one or multiple collections.

// TODO market.reorderAlbums Reorders the collections list.

// TODO market.reorderItems Changes item place in a collection.

// TODO market.report Sends a complaint to the item.

// TODO market.reportComment Sends a complaint to the item's comment.

// TODO market.restore Restores recently deleted item

// TODO market.restoreComment Restores a recently deleted comment

// TODO market.search Searches market items in a community's catalog

// TODO messages.addChatUser Adds a new user to a chat.

// TODO messages.allowMessagesFromGroup Allows sending messages from community to the current user.

// TODO messages.createChat Creates a chat with several participants.

// TODO messages.delete Deletes one or more messages.

// TODO messages.deleteChatPhoto Deletes a chat's cover picture.

// TODO messages.deleteConversation Deletes private messages in a conversation.

// TODO messages.denyMessagesFromGroup Denies sending message from community to the current user.

// TODO messages.edit Edits the message.

// TODO messages.editChat Edits the title of a chat.

// TODO messages.getByConversationMessageId

// TODO messages.getById Returns messages by their IDs.

// TODO messages.getChat Returns information about a chat.

// TODO messages.getChatPreview Allows to receive chat preview by the invitation link.

// TODO messages.getConversationMembers Returns a list of IDs of users participating in a conversation.

// TODO messages.getConversations Returns a list of conversations.

// TODO messages.getConversationsById Returns conversations by their IDs.

// TODO messages.getHistory Returns message history for the specified user or group chat.

// TODO messages.getHistoryAttachments Returns media files from the dialog or group chat.

// TODO messages.getImportantMessages

// TODO messages.getInviteLink Receives a link to invite a user to the chat.

// TODO messages.getLastActivity Returns a user's current status and date of last activity.

// TODO messages.getLongPollHistory Returns updates in user's private messages.

// TODO messages.getLongPollServer Returns data required for connection to a Long Poll server.

// TODO messages.isMessagesFromGroupAllowed Returns information whether sending messages from the community to current user is allowed.

// TODO messages.joinChatByInviteLink Allows to enter the chat by the invitation link.

// TODO messages.markAsAnsweredConversation

// TODO messages.markAsImportant Marks and un marks messages as important (starred).

// TODO messages.markAsImportantConversation

// TODO messages.markAsRead Marks messages as read.

// TODO messages.pin

// TODO messages.removeChatUser Allows the current user to leave a chat or, if the current user started the chat, allows the user to remove another user from the chat.

// TODO messages.restore Restores a deleted message.

// TODO messages.search Returns a list of the current user's private messages that match search criteria.

// TODO messages.searchConversations Returns a list of conversations that match search criteria.

// TODO messages.send Sends a message.
func (vk *VK) MessagesSend(params map[string]string) error {
	const method = "messages.send"
	// TODO if user_ids in params

	_, err := vk.Request(method, params)
	if err != nil {
		return err
	}

	return nil
}

// TODO messages.setActivity Changes the status of a user as typing in a conversation.

// TODO messages.setChatPhoto Sets a previously-uploaded picture as the cover picture of a chat.

// TODO messages.unpin

// TODO newsfeed.addBan Prevents news from specified users and communities from appearing in the current user's newsfeed.

// TODO newsfeed.deleteBan Allows news from previously banned users and communities to be shown in the current user's newsfeed.

// TODO newsfeed.deleteList

// TODO newsfeed.get Returns data required to show newsfeed for the current user.

// TODO newsfeed.getBanned Returns a list of users and communities banned from the current user's newsfeed.

// TODO newsfeed.getComments Returns a list of comments in the current user's newsfeed.

// TODO newsfeed.getDiscoverForContestant

// TODO newsfeed.getLists Returns a list of newsfeeds followed by the current user.

// TODO newsfeed.getMentions Returns a list of posts on user walls in which the current user is mentioned.

// TODO newsfeed.getRecommended Returns a list of newsfeeds recommended to the current user.

// TODO newsfeed.getSuggestedSources Returns communities and users that current user is suggested to follow.

// TODO newsfeed.ignoreItemHides an item from the newsfeed.

// TODO newsfeed.saveList Creates and edits user newsfeed lists

// TODO newsfeed.search Returns search results by statuses.

// TODO newsfeed.unignoreItem Returns a hidden item to the newsfeed.

// TODO newsfeed.unsubscribe Unsubscribes the current user from specified newsfeeds.

// TODO notes.add Creates a new note for the current user.

// TODO notes.createComment Adds a new comment on a note.

// TODO notes.delete Deletes a note of the current user.

// TODO notes.deleteComment Deletes a comment on a note.

// TODO notes.edit Edits a note of the current user.

// TODO notes.editComment Edits a comment on a note.

// TODO notes.get Returns a list of notes created by a user.

// TODO notes.getById Returns a note by its ID.

// TODO notes.getComments Returns a list of comments on a note.

// TODO notes.restoreComment Restores a deleted comment on a note.

// TODO notifications.get Returns a list of notifications about other users' feedback to the current user's wall posts.

// TODO notifications.markAsViewed Resets the counter of new notifications about other users' feedback to the current user's wall posts.

// TODO notifications.sendMessage

// TODO pages.clearCache Allows to clear the cache of particular external pages which may be attached to VK posts.

// TODO pages.get Returns information about a wiki page.

// TODO pages.getHistory Returns a list of all previous versions of a wiki page.

// TODO pages.getTitles Returns a list of wiki pages in a group.

// TODO pages.getVersion Returns the text of one of the previous versions of a wiki page.

// TODO pages.parseWiki Returns HTML representation of the wiki markup.

// TODO pages.save Saves the text of a wiki page.

// TODO pages.saveAccess Saves modified read and edit access settings for a wiki page.

// TODO photos.confirmTag Confirms a tag on a photo.

// TODO photos.copy Allows to copy a photo to the "Saved photos" album

// TODO photos.createAlbum Creates an empty photo album.

// TODO photos.createComment Adds a new comment on the photo.

// TODO photos.delete Deletes a photo.

// TODO photos.deleteAlbum Deletes a photo album belonging to the current user.

// TODO photos.deleteComment Deletes a comment on the photo.

// TODO photos.edit Edits the caption of a photo.

// TODO photos.editAlbum Edits information about a photo album.

// TODO photos.editComment Edits a comment on a photo.

// TODO photos.get Returns a list of a user's or community's photos.

// TODO photos.getAlbums Returns a list of a user's or community's photo albums.

// TODO photos.getAlbumsCount Returns the number of photo albums belonging to a user or community.

// TODO photos.getAll Returns a list of photos belonging to a user or community, in reverse chronological order.

// TODO photos.getAllComments Returns a list of comments on a specific photo album or all albums of the user sorted in reverse chronological order.

// TODO photos.getById Returns information about photos by their IDs.

// TODO photos.getChatUploadServer Returns an upload link for chat cover pictures.

// TODO photos.getComments Returns a list of comments on a photo.

// TODO photos.getMarketAlbumUploadServer Returns the server address for market album photo upload.

// TODO photos.getMarketUploadServer Returns the server address for market photo upload.

// TODO photos.getMessagesUploadServer Returns the server address for photo upload in a private message for a user.

// TODO photos.getNewTags Returns a list of photos with tags that have not been viewed.

// TODO photos.getOwnerCoverPhotoUploadServer Receives server address for uploading community cover.

// TODO photos.getOwnerPhotoUploadServer Returns an upload server address for a profile or community photo.

// TODO photos.getTags Returns a list of tags on a photo.

// TODO photos.getUploadServer Returns the server address for photo upload.

// TODO photos.getUserPhotos Returns a list of photos in which a user is tagged.

// TODO photos.getWallUploadServer Returns the server address for photo upload onto a user's wall.

// TODO photos.makeCover Makes a photo into an album cover.

// TODO photos.moveMoves a photo from one album to another.

// TODO photos.putTag Adds a tag on the photo.

// TODO photos.removeTag Removes a tag from a photo.

// TODO photos.reorderAlbums Reorders the album in the list of user albums.

// TODO photos.reorderPhotos Reorders the photo in the list of photos of the user album.

// TODO photos.report Reports (submits a complaint about) a photo.

// TODO photos.reportComment Reports (submits a complaint about) a comment on a photo.

// TODO photos.restore Restores a deleted photo.

// TODO photos.restoreComment Restores a deleted comment on a photo.

// TODO photos.save Saves photos after successful uploading.

// TODO photos.saveMarketAlbum Photo Saves market album photos after successful uploading.

// TODO photos.saveMarketPhoto Saves market photos after successful uploading.

// TODO photos.saveMessagesPhoto Saves a photo after being successfully uploaded.URL obtained with photos.getMessagesUploadServer method.

// TODO photos.saveOwnerCoverPhoto Saves cover photo after successful uploading.

// TODO photos.saveOwnerPhoto Saves a profile or community photo.

// TODO photos.saveWallPhoto Saves a photo to a user's or community's wall after being uploaded.

// TODO photos.search Returns a list of photos.

// TODO polls.addVote Adds the current user's vote to the selected answer in the poll.

// TODO polls.create Creates polls that can be attached to the users' or communities' posts.

// TODO polls.deleteVote Deletes the current user's vote from the selected answer in the poll.

// TODO polls.edit Edits created polls

// TODO polls.getBackgrounds Return default backgrounds for polls.

// TODO polls.getById Returns detailed information about a poll by its ID.

// TODO polls.getPhotoUploadServer Returns a URL for uploading a photo to a poll.

// TODO polls.getVoters Returns a list of IDs of users who selected specific answers in the poll.

// TODO polls.savePhoto Allows to save poll's uploaded photo.

// TODO prettyCards.create

// TODO prettyCards.delete

// TODO prettyCards.edit

// TODO prettyCards.get

// TODO prettyCards.getById

// TODO prettyCards.getUploadURL

// TODO search.getHints Allows the programmer to do a quick search for any substring.

// TODO stats.get Returns statistics of a community or an application.

// TODO stats.getPostReach Returns stats for a wall post.

// TODO stats.trackVisitor Adds current session's data in the application statistics.

// TODO status.get Returns data required to show the status of a user or community.

// TODO status.set Sets a new status for the current user.

// TODO storage.get Returns a value of variable with the name set by key parameter.

// TODO storage.getKeys Returns the names of all variables.

// TODO storage.set Saves a value of variable with the name set by key parameter.

// TODO stories.banOwner Allows to hide stories from chosen sources from current user's feed.

// TODO stories.delete Allows to delete story.

// TODO stories.get Returns stories available for current user.

// TODO stories.getBanned Returns list of sources hidden from current user's feed.

// TODO stories.getById Returns story by its ID.

// TODO stories.getPhotoUploadServer Returns URL for uploading a story with photo.

// TODO stories.getReplies Returns replies to the story.

// TODO stories.getStats Return statictics data for the story.

// TODO stories.getVideoUploadServer Allows to receive URL for uploading story with video.

// TODO stories.getViewers Returns a list of story viewers.

// TODO stories.hideAllReplies Hides all replies in the last 24 hours from the user to current user's stories.

// TODO stories.hideReply Hides the reply to the current user's story.

// TODO stories.unbanOwner Allows to show stories from hidden sources in current user's feed.

// TODO streaming.getServerUrl Allows to receive data for the connection to Streaming API.

// TODO streaming.getSettings Allows to receive monthly tier for Streaming API.

// TODO streaming.getStats Allows to receive statistics for prepared and received events in Streaming API.

// TODO streaming.getStem

// TODO streaming.setSettings Allows to set monthly tier for Streaming API.

// users.get Returns detailed information on users.
func (vk *VK) UsersGet(params map[string]string) (UsersGetResponse, error) {
	var response UsersGetResponse

	rawResponse, err := vk.Request("users.get", params)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}

// TODO users.getFollowers Returns a list of IDs of followers of the user in question, sorted by date added, most recent first.

// TODO users.getSubscriptions Returns a list of IDs of users and public pages followed by the user.

// TODO users.isAppUser Returns information whether a user installed the application.

// TODO users.report Reports (submits a complain about) a user.

// TODO users.search Returns a list of users matching the search criteria.

// TODO utils.checkLink Checks whether a link is blocked in VK.

// TODO utils.deleteFromLastShortened Deletes shortened link from user's list.

// TODO utils.getLastShortenedLinks Returns a list of user's shortened links.

// TODO utils.getLinkStats Returns stats data for shortened link.

// TODO utils.getServerTime Returns the current time of the VK server.

// TODO utils.getShortLink Allows to receive a link shortened via vk.cc.

// TODO utils.resolveScreenName Detects a type of object (e.g., user, community, application) and its ID by screen name.
func (vk *VK) UtilsResolveScreenName(params map[string]string) (utilsResolveScreenNameResponse, error) {
	var response utilsResolveScreenNameResponse

	rawResponse, err := vk.Request("utils.resolveScreenName", params)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}

// TODO video.add Adds a video to a user or community page.

// TODO video.addAlbum Creates an empty album for videos.

// TODO video.addToAlbum

// TODO video.createComment Adds a new comment on a video.

// TODO video.delete Deletes a video from a user or community page.

// TODO video.deleteAlbum Deletes a video album.

// TODO video.deleteComment Deletes a comment on a video.

// TODO video.edit Edits information about a video on a user or community page.

// TODO video.editAlbum Edits the title of a video album.

// TODO video.editComment Edits the text of a comment on a video.

// TODO video.get Returns detailed information about videos.

// TODO video.getAlbumById Returns video album info

// TODO video.getAlbums Returns a list of video albums owned by a user or community.

// TODO video.getAlbumsByVideo

// TODO video.getComments Returns a list of comments on a video.

// TODO video.removeFromAlbum

// TODO video.reorderAlbums Reorders the album in the list of user video albums.

// TODO video.reorderVideos Reorders the video in the video album.

// TODO video.report Reports (submits a complaint about) a video.

// TODO video.reportComment Reports (submits a complaint about) a comment on a video.

// TODO video.restore Restores a previously deleted video.

// TODO video.restoreComment Restores a previously deleted comment on a video.

// TODO video.save Returns a server address (required for upload) and video data.

// TODO video.search Returns a list of videos under the set search criterion.

// TODO wall.closeComments

// TODO wall.createComment Adds a comment to a post on a user wall or community wall.

// TODO wall.delete Deletes a post from a user wall or community wall.

// TODO wall.deleteComment Deletes a comment on a post on a user wall or community wall.

// TODO wall.edit Edits a post on a user wall or community wall.

// TODO wall.editAdsStealth Allows to edit hidden post.

// TODO wall.editComment Edits a comment on a user wall or community wall.

// TODO wall.get Returns a list of posts on a user wall or community wall.

// TODO wall.getById Returns a list of posts from user or community walls by their IDs.

// TODO wall.getComment Allows to obtain wall comment info.

// TODO wall.getComments Returns a list of comments on a post on a user wall or community wall.

// TODO wall.getReposts Returns information about reposts of a post on user wall or community wall.

// TODO wall.openComments

// TODO wall.pin Pins the post on wall.

// TODO wall.post Adds a new post on a user wall or community wall.Can also be used to publish suggested or scheduled posts.

// TODO wall.postAdsStealth Allows to create hidden post which will not be shown on the community's wall and can be used for creating an ad with type "Community post".

// TODO wall.reportComment Reports (submits a complaint about) a comment on a post on a user wall or community wall.

// TODO wall.reportPost Reports (submits a complaint about) a post on a user wall or community wall.

// TODO wall.repost Reposts ( copies) an object to a user wall or community wall.

// TODO wall.restore Restores a post deleted from a user wall or community wall.

// TODO wall.restoreComment Restores a comment deleted from a user wall or community wall.

// TODO wall.search Allows to search posts on user or community walls.

// TODO wall.unpin Unpins the post on wall.

// TODO widgets.getComments Gets a list of comments for the page added through the Comments widget.

// TODO widgets.getPages Gets a list of application/site pages where the Comments widget or Like widget is installed.