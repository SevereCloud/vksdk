package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// VideoAddBulder builder
//
// Adds a video to a user or community page.
//
// https://vk.com/dev/video.add
type VideoAddBulder struct {
	api.Params
}

// NewVideoAddBulder func
func NewVideoAddBulder() *VideoAddBulder {
	return &VideoAddBulder{api.Params{}}
}

// TargetID identifier of a user or community to add a video to. Use a negative value to designate a community ID.
func (b *VideoAddBulder) TargetID(v int) {
	b.Params["target_id"] = v
}

// VideoID Video ID.
func (b *VideoAddBulder) VideoID(v int) {
	b.Params["video_id"] = v
}

// OwnerID ID of the user or community that owns the video. Use a negative value to designate a community ID.
func (b *VideoAddBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoAddAlbumBulder builder
//
// Creates an empty album for videos.
//
// https://vk.com/dev/video.addAlbum
type VideoAddAlbumBulder struct {
	api.Params
}

// NewVideoAddAlbumBulder func
func NewVideoAddAlbumBulder() *VideoAddAlbumBulder {
	return &VideoAddAlbumBulder{api.Params{}}
}

// GroupID Community ID (if the album will be created in a community).
func (b *VideoAddAlbumBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Title Album title.
func (b *VideoAddAlbumBulder) Title(v string) {
	b.Params["title"] = v
}

// Privacy new access permissions for the album. Possible values: , *'0' – all users,, *'1' – friends only,, *'2' – friends and friends of friends,, *'3' – "only me".
func (b *VideoAddAlbumBulder) Privacy(v []string) {
	b.Params["privacy"] = v
}

// VideoAddToAlbumBulder builder
//
// https://vk.com/dev/video.addToAlbum
type VideoAddToAlbumBulder struct {
	api.Params
}

// NewVideoAddToAlbumBulder func
func NewVideoAddToAlbumBulder() *VideoAddToAlbumBulder {
	return &VideoAddToAlbumBulder{api.Params{}}
}

// TargetID parameter
func (b *VideoAddToAlbumBulder) TargetID(v int) {
	b.Params["target_id"] = v
}

// AlbumID parameter
func (b *VideoAddToAlbumBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// AlbumIDs parameter
func (b *VideoAddToAlbumBulder) AlbumIDs(v []int) {
	b.Params["album_ids"] = v
}

// OwnerID parameter
func (b *VideoAddToAlbumBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoID parameter
func (b *VideoAddToAlbumBulder) VideoID(v int) {
	b.Params["video_id"] = v
}

// VideoCreateCommentBulder builder
//
// Adds a new comment on a video.
//
// https://vk.com/dev/video.createComment
type VideoCreateCommentBulder struct {
	api.Params
}

// NewVideoCreateCommentBulder func
func NewVideoCreateCommentBulder() *VideoCreateCommentBulder {
	return &VideoCreateCommentBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoCreateCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoID Video ID.
func (b *VideoCreateCommentBulder) VideoID(v int) {
	b.Params["video_id"] = v
}

// Message New comment text.
func (b *VideoCreateCommentBulder) Message(v string) {
	b.Params["message"] = v
}

// Attachments List of objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media attachment owner. '<media_id>' — Media attachment ID. Example: "photo100172_166443618,photo66748_265827614"
func (b *VideoCreateCommentBulder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// FromGroup '1' — to post the comment from a community name (only if 'owner_id'<0)
func (b *VideoCreateCommentBulder) FromGroup(v bool) {
	b.Params["from_group"] = v
}

// ReplyToComment parameter
func (b *VideoCreateCommentBulder) ReplyToComment(v int) {
	b.Params["reply_to_comment"] = v
}

// StickerID parameter
func (b *VideoCreateCommentBulder) StickerID(v int) {
	b.Params["sticker_id"] = v
}

// GUID parameter
func (b *VideoCreateCommentBulder) GUID(v string) {
	b.Params["guid"] = v
}

// VideoDeleteBulder builder
//
// Deletes a video from a user or community page.
//
// https://vk.com/dev/video.delete
type VideoDeleteBulder struct {
	api.Params
}

// NewVideoDeleteBulder func
func NewVideoDeleteBulder() *VideoDeleteBulder {
	return &VideoDeleteBulder{api.Params{}}
}

// VideoID Video ID.
func (b *VideoDeleteBulder) VideoID(v int) {
	b.Params["video_id"] = v
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoDeleteBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// TargetID parameter
func (b *VideoDeleteBulder) TargetID(v int) {
	b.Params["target_id"] = v
}

// VideoDeleteAlbumBulder builder
//
// Deletes a video album.
//
// https://vk.com/dev/video.deleteAlbum
type VideoDeleteAlbumBulder struct {
	api.Params
}

// NewVideoDeleteAlbumBulder func
func NewVideoDeleteAlbumBulder() *VideoDeleteAlbumBulder {
	return &VideoDeleteAlbumBulder{api.Params{}}
}

// GroupID Community ID (if the album is owned by a community).
func (b *VideoDeleteAlbumBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// AlbumID Album ID.
func (b *VideoDeleteAlbumBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// VideoDeleteCommentBulder builder
//
// Deletes a comment on a video.
//
// https://vk.com/dev/video.deleteComment
type VideoDeleteCommentBulder struct {
	api.Params
}

// NewVideoDeleteCommentBulder func
func NewVideoDeleteCommentBulder() *VideoDeleteCommentBulder {
	return &VideoDeleteCommentBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoDeleteCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID ID of the comment to be deleted.
func (b *VideoDeleteCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// VideoEditBulder builder
//
// Edits information about a video on a user or community page.
//
// https://vk.com/dev/video.edit
type VideoEditBulder struct {
	api.Params
}

// NewVideoEditBulder func
func NewVideoEditBulder() *VideoEditBulder {
	return &VideoEditBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoEditBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoID Video ID.
func (b *VideoEditBulder) VideoID(v int) {
	b.Params["video_id"] = v
}

// Name New video title.
func (b *VideoEditBulder) Name(v string) {
	b.Params["name"] = v
}

// Desc New video description.
func (b *VideoEditBulder) Desc(v string) {
	b.Params["desc"] = v
}

// PrivacyView Privacy settings in a [vk.com/dev/privacy_setting|special format]. Privacy setting is available for videos uploaded to own profile by user.
func (b *VideoEditBulder) PrivacyView(v []string) {
	b.Params["privacy_view"] = v
}

// PrivacyComment Privacy settings for comments in a [vk.com/dev/privacy_setting|special format].
func (b *VideoEditBulder) PrivacyComment(v []string) {
	b.Params["privacy_comment"] = v
}

// NoComments Disable comments for the group video.
func (b *VideoEditBulder) NoComments(v bool) {
	b.Params["no_comments"] = v
}

// Repeat '1' — to repeat the playback of the video, '0' — to play the video once,
func (b *VideoEditBulder) Repeat(v bool) {
	b.Params["repeat"] = v
}

// VideoEditAlbumBulder builder
//
// Edits the title of a video album.
//
// https://vk.com/dev/video.editAlbum
type VideoEditAlbumBulder struct {
	api.Params
}

// NewVideoEditAlbumBulder func
func NewVideoEditAlbumBulder() *VideoEditAlbumBulder {
	return &VideoEditAlbumBulder{api.Params{}}
}

// GroupID Community ID (if the album edited is owned by a community).
func (b *VideoEditAlbumBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// AlbumID Album ID.
func (b *VideoEditAlbumBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// Title New album title.
func (b *VideoEditAlbumBulder) Title(v string) {
	b.Params["title"] = v
}

// Privacy new access permissions for the album. Possible values: , *'0' – all users,, *'1' – friends only,, *'2' – friends and friends of friends,, *'3' – "only me".
func (b *VideoEditAlbumBulder) Privacy(v []string) {
	b.Params["privacy"] = v
}

// VideoEditCommentBulder builder
//
// Edits the text of a comment on a video.
//
// https://vk.com/dev/video.editComment
type VideoEditCommentBulder struct {
	api.Params
}

// NewVideoEditCommentBulder func
func NewVideoEditCommentBulder() *VideoEditCommentBulder {
	return &VideoEditCommentBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoEditCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID Comment ID.
func (b *VideoEditCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// Message New comment text.
func (b *VideoEditCommentBulder) Message(v string) {
	b.Params["message"] = v
}

// Attachments List of objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media attachment owner. '<media_id>' — Media attachment ID. Example: "photo100172_166443618,photo66748_265827614"
func (b *VideoEditCommentBulder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// VideoGetBulder builder
//
// Returns detailed information about videos.
//
// https://vk.com/dev/video.get
type VideoGetBulder struct {
	api.Params
}

// NewVideoGetBulder func
func NewVideoGetBulder() *VideoGetBulder {
	return &VideoGetBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video(s).
func (b *VideoGetBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Videos Video IDs, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", Use a negative value to designate a community ID. Example: "-4363_136089719,13245770_137352259"
func (b *VideoGetBulder) Videos(v []string) {
	b.Params["videos"] = v
}

// AlbumID ID of the album containing the video(s).
func (b *VideoGetBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// Count Number of videos to return.
func (b *VideoGetBulder) Count(v int) {
	b.Params["count"] = v
}

// Offset Offset needed to return a specific subset of videos.
func (b *VideoGetBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Extended '1' — to return an extended response with additional fields
func (b *VideoGetBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// VideoGetAlbumByIDBulder builder
//
// Returns video album info
//
// https://vk.com/dev/video.getAlbumById
type VideoGetAlbumByIDBulder struct {
	api.Params
}

// NewVideoGetAlbumByIDBulder func
func NewVideoGetAlbumByIDBulder() *VideoGetAlbumByIDBulder {
	return &VideoGetAlbumByIDBulder{api.Params{}}
}

// OwnerID identifier of a user or community to add a video to. Use a negative value to designate a community ID.
func (b *VideoGetAlbumByIDBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumID Album ID.
func (b *VideoGetAlbumByIDBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// VideoGetAlbumsBulder builder
//
// Returns a list of video albums owned by a user or community.
//
// https://vk.com/dev/video.getAlbums
type VideoGetAlbumsBulder struct {
	api.Params
}

// NewVideoGetAlbumsBulder func
func NewVideoGetAlbumsBulder() *VideoGetAlbumsBulder {
	return &VideoGetAlbumsBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video album(s).
func (b *VideoGetAlbumsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Offset Offset needed to return a specific subset of video albums.
func (b *VideoGetAlbumsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of video albums to return.
func (b *VideoGetAlbumsBulder) Count(v int) {
	b.Params["count"] = v
}

// Extended '1' — to return additional information about album privacy settings for the current user
func (b *VideoGetAlbumsBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// NeedSystem parameter
func (b *VideoGetAlbumsBulder) NeedSystem(v bool) {
	b.Params["need_system"] = v
}

// VideoGetAlbumsByVideoBulder builder
//
// https://vk.com/dev/video.getAlbumsByVideo
type VideoGetAlbumsByVideoBulder struct {
	api.Params
}

// NewVideoGetAlbumsByVideoBulder func
func NewVideoGetAlbumsByVideoBulder() *VideoGetAlbumsByVideoBulder {
	return &VideoGetAlbumsByVideoBulder{api.Params{}}
}

// TargetID parameter
func (b *VideoGetAlbumsByVideoBulder) TargetID(v int) {
	b.Params["target_id"] = v
}

// OwnerID parameter
func (b *VideoGetAlbumsByVideoBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoID parameter
func (b *VideoGetAlbumsByVideoBulder) VideoID(v int) {
	b.Params["video_id"] = v
}

// Extended parameter
func (b *VideoGetAlbumsByVideoBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// VideoGetCommentsBulder builder
//
// Returns a list of comments on a video.
//
// https://vk.com/dev/video.getComments
type VideoGetCommentsBulder struct {
	api.Params
}

// NewVideoGetCommentsBulder func
func NewVideoGetCommentsBulder() *VideoGetCommentsBulder {
	return &VideoGetCommentsBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoGetCommentsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoID Video ID.
func (b *VideoGetCommentsBulder) VideoID(v int) {
	b.Params["video_id"] = v
}

// NeedLikes '1' — to return an additional 'likes' field
func (b *VideoGetCommentsBulder) NeedLikes(v bool) {
	b.Params["need_likes"] = v
}

// StartCommentID parameter
func (b *VideoGetCommentsBulder) StartCommentID(v int) {
	b.Params["start_comment_id"] = v
}

// Offset Offset needed to return a specific subset of comments.
func (b *VideoGetCommentsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of comments to return.
func (b *VideoGetCommentsBulder) Count(v int) {
	b.Params["count"] = v
}

// Sort Sort order: 'asc' — oldest comment first, 'desc' — newest comment first
func (b *VideoGetCommentsBulder) Sort(v string) {
	b.Params["sort"] = v
}

// Extended parameter
func (b *VideoGetCommentsBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields parameter
func (b *VideoGetCommentsBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// VideoRemoveFromAlbumBulder builder
//
// https://vk.com/dev/video.removeFromAlbum
type VideoRemoveFromAlbumBulder struct {
	api.Params
}

// NewVideoRemoveFromAlbumBulder func
func NewVideoRemoveFromAlbumBulder() *VideoRemoveFromAlbumBulder {
	return &VideoRemoveFromAlbumBulder{api.Params{}}
}

// TargetID parameter
func (b *VideoRemoveFromAlbumBulder) TargetID(v int) {
	b.Params["target_id"] = v
}

// AlbumID parameter
func (b *VideoRemoveFromAlbumBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// AlbumIDs parameter
func (b *VideoRemoveFromAlbumBulder) AlbumIDs(v []int) {
	b.Params["album_ids"] = v
}

// OwnerID parameter
func (b *VideoRemoveFromAlbumBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoID parameter
func (b *VideoRemoveFromAlbumBulder) VideoID(v int) {
	b.Params["video_id"] = v
}

// VideoReorderAlbumsBulder builder
//
// Reorders the album in the list of user video albums.
//
// https://vk.com/dev/video.reorderAlbums
type VideoReorderAlbumsBulder struct {
	api.Params
}

// NewVideoReorderAlbumsBulder func
func NewVideoReorderAlbumsBulder() *VideoReorderAlbumsBulder {
	return &VideoReorderAlbumsBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the albums..
func (b *VideoReorderAlbumsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumID Album ID.
func (b *VideoReorderAlbumsBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// Before ID of the album before which the album in question shall be placed.
func (b *VideoReorderAlbumsBulder) Before(v int) {
	b.Params["before"] = v
}

// After ID of the album after which the album in question shall be placed.
func (b *VideoReorderAlbumsBulder) After(v int) {
	b.Params["after"] = v
}

// VideoReorderVideosBulder builder
//
// Reorders the video in the video album.
//
// https://vk.com/dev/video.reorderVideos
type VideoReorderVideosBulder struct {
	api.Params
}

// NewVideoReorderVideosBulder func
func NewVideoReorderVideosBulder() *VideoReorderVideosBulder {
	return &VideoReorderVideosBulder{api.Params{}}
}

// TargetID ID of the user or community that owns the album with videos.
func (b *VideoReorderVideosBulder) TargetID(v int) {
	b.Params["target_id"] = v
}

// AlbumID ID of the video album.
func (b *VideoReorderVideosBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoReorderVideosBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoID ID of the video.
func (b *VideoReorderVideosBulder) VideoID(v int) {
	b.Params["video_id"] = v
}

// BeforeOwnerID ID of the user or community that owns the video before which the video in question shall be placed.
func (b *VideoReorderVideosBulder) BeforeOwnerID(v int) {
	b.Params["before_owner_id"] = v
}

// BeforeVideoID ID of the video before which the video in question shall be placed.
func (b *VideoReorderVideosBulder) BeforeVideoID(v int) {
	b.Params["before_video_id"] = v
}

// AfterOwnerID ID of the user or community that owns the video after which the photo in question shall be placed.
func (b *VideoReorderVideosBulder) AfterOwnerID(v int) {
	b.Params["after_owner_id"] = v
}

// AfterVideoID ID of the video after which the photo in question shall be placed.
func (b *VideoReorderVideosBulder) AfterVideoID(v int) {
	b.Params["after_video_id"] = v
}

// VideoReportBulder builder
//
// Reports (submits a complaint about) a video.
//
// https://vk.com/dev/video.report
type VideoReportBulder struct {
	api.Params
}

// NewVideoReportBulder func
func NewVideoReportBulder() *VideoReportBulder {
	return &VideoReportBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoReportBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoID Video ID.
func (b *VideoReportBulder) VideoID(v int) {
	b.Params["video_id"] = v
}

// Reason Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
func (b *VideoReportBulder) Reason(v int) {
	b.Params["reason"] = v
}

// Comment Comment describing the complaint.
func (b *VideoReportBulder) Comment(v string) {
	b.Params["comment"] = v
}

// SearchQuery (If the video was found in search results.) Search query string.
func (b *VideoReportBulder) SearchQuery(v string) {
	b.Params["search_query"] = v
}

// VideoReportCommentBulder builder
//
// Reports (submits a complaint about) a comment on a video.
//
// https://vk.com/dev/video.reportComment
type VideoReportCommentBulder struct {
	api.Params
}

// NewVideoReportCommentBulder func
func NewVideoReportCommentBulder() *VideoReportCommentBulder {
	return &VideoReportCommentBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoReportCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID ID of the comment being reported.
func (b *VideoReportCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// Reason Reason for the complaint: , 0 – spam , 1 – child pornography , 2 – extremism , 3 – violence , 4 – drug propaganda , 5 – adult material , 6 – insult, abuse
func (b *VideoReportCommentBulder) Reason(v int) {
	b.Params["reason"] = v
}

// VideoRestoreBulder builder
//
// Restores a previously deleted video.
//
// https://vk.com/dev/video.restore
type VideoRestoreBulder struct {
	api.Params
}

// NewVideoRestoreBulder func
func NewVideoRestoreBulder() *VideoRestoreBulder {
	return &VideoRestoreBulder{api.Params{}}
}

// VideoID Video ID.
func (b *VideoRestoreBulder) VideoID(v int) {
	b.Params["video_id"] = v
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoRestoreBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoRestoreCommentBulder builder
//
// Restores a previously deleted comment on a video.
//
// https://vk.com/dev/video.restoreComment
type VideoRestoreCommentBulder struct {
	api.Params
}

// NewVideoRestoreCommentBulder func
func NewVideoRestoreCommentBulder() *VideoRestoreCommentBulder {
	return &VideoRestoreCommentBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoRestoreCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID ID of the deleted comment.
func (b *VideoRestoreCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// VideoSaveBulder builder
//
// Returns a server address (required for upload) and video data.
//
// https://vk.com/dev/video.save
type VideoSaveBulder struct {
	api.Params
}

// NewVideoSaveBulder func
func NewVideoSaveBulder() *VideoSaveBulder {
	return &VideoSaveBulder{api.Params{}}
}

// Name Name of the video.
func (b *VideoSaveBulder) Name(v string) {
	b.Params["name"] = v
}

// Description Description of the video.
func (b *VideoSaveBulder) Description(v string) {
	b.Params["description"] = v
}

// IsPrivate '1' — to designate the video as private (send it via a private message), the video will not appear on the user's video list and will not be available by ID for other users, '0' — not to designate the video as private
func (b *VideoSaveBulder) IsPrivate(v bool) {
	b.Params["is_private"] = v
}

// Wallpost '1' — to post the saved video on a user's wall, '0' — not to post the saved video on a user's wall
func (b *VideoSaveBulder) Wallpost(v bool) {
	b.Params["wallpost"] = v
}

// Link URL for embedding the video from an external website.
func (b *VideoSaveBulder) Link(v string) {
	b.Params["link"] = v
}

// GroupID ID of the community in which the video will be saved. By default, the current user's page.
func (b *VideoSaveBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// AlbumID ID of the album to which the saved video will be added.
func (b *VideoSaveBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// PrivacyView parameter
func (b *VideoSaveBulder) PrivacyView(v []string) {
	b.Params["privacy_view"] = v
}

// PrivacyComment parameter
func (b *VideoSaveBulder) PrivacyComment(v []string) {
	b.Params["privacy_comment"] = v
}

// NoComments parameter
func (b *VideoSaveBulder) NoComments(v bool) {
	b.Params["no_comments"] = v
}

// Repeat '1' — to repeat the playback of the video, '0' — to play the video once,
func (b *VideoSaveBulder) Repeat(v bool) {
	b.Params["repeat"] = v
}

// Compression parameter
func (b *VideoSaveBulder) Compression(v bool) {
	b.Params["compression"] = v
}

// VideoSearchBulder builder
//
// Returns a list of videos under the set search criterion.
//
// https://vk.com/dev/video.search
type VideoSearchBulder struct {
	api.Params
}

// NewVideoSearchBulder func
func NewVideoSearchBulder() *VideoSearchBulder {
	return &VideoSearchBulder{api.Params{}}
}

// Q Search query string (e.g., 'The Beatles').
func (b *VideoSearchBulder) Q(v string) {
	b.Params["q"] = v
}

// Sort Sort order: '1' — by duration, '2' — by relevance, '0' — by date added
func (b *VideoSearchBulder) Sort(v int) {
	b.Params["sort"] = v
}

// Hd If not null, only searches for high-definition videos.
func (b *VideoSearchBulder) Hd(v int) {
	b.Params["hd"] = v
}

// Adult '1' — to disable the Safe Search filter, '0' — to enable the Safe Search filter
func (b *VideoSearchBulder) Adult(v bool) {
	b.Params["adult"] = v
}

// Filters Filters to apply: 'youtube' — return YouTube videos only, 'vimeo' — return Vimeo videos only, 'short' — return short videos only, 'long' — return long videos only
func (b *VideoSearchBulder) Filters(v []string) {
	b.Params["filters"] = v
}

// SearchOwn parameter
func (b *VideoSearchBulder) SearchOwn(v bool) {
	b.Params["search_own"] = v
}

// Offset Offset needed to return a specific subset of videos.
func (b *VideoSearchBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Longer parameter
func (b *VideoSearchBulder) Longer(v int) {
	b.Params["longer"] = v
}

// Shorter parameter
func (b *VideoSearchBulder) Shorter(v int) {
	b.Params["shorter"] = v
}

// Count Number of videos to return.
func (b *VideoSearchBulder) Count(v int) {
	b.Params["count"] = v
}

// Extended parameter
func (b *VideoSearchBulder) Extended(v bool) {
	b.Params["extended"] = v
}
