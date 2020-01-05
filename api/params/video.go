package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// VideoAddBuilder builder
//
// Adds a video to a user or community page.
//
// https://vk.com/dev/video.add
type VideoAddBuilder struct {
	api.Params
}

// NewVideoAddBuilder func
func NewVideoAddBuilder() *VideoAddBuilder {
	return &VideoAddBuilder{api.Params{}}
}

// TargetID identifier of a user or community to add a video to. Use a negative value to designate a community ID.
func (b *VideoAddBuilder) TargetID(v int) {
	b.Params["target_id"] = v
}

// VideoID Video ID.
func (b *VideoAddBuilder) VideoID(v int) {
	b.Params["video_id"] = v
}

// OwnerID ID of the user or community that owns the video. Use a negative value to designate a community ID.
func (b *VideoAddBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoAddAlbumBuilder builder
//
// Creates an empty album for videos.
//
// https://vk.com/dev/video.addAlbum
type VideoAddAlbumBuilder struct {
	api.Params
}

// NewVideoAddAlbumBuilder func
func NewVideoAddAlbumBuilder() *VideoAddAlbumBuilder {
	return &VideoAddAlbumBuilder{api.Params{}}
}

// GroupID Community ID (if the album will be created in a community).
func (b *VideoAddAlbumBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Title Album title.
func (b *VideoAddAlbumBuilder) Title(v string) {
	b.Params["title"] = v
}

// Privacy new access permissions for the album. Possible values:
//
// * 0 – all users,
//
// * 1 – friends only,
//
// * 2 – friends and friends of friends,
//
// * 3 – "only me".
func (b *VideoAddAlbumBuilder) Privacy(v []string) {
	b.Params["privacy"] = v
}

// VideoAddToAlbumBuilder builder
//
// https://vk.com/dev/video.addToAlbum
type VideoAddToAlbumBuilder struct {
	api.Params
}

// NewVideoAddToAlbumBuilder func
func NewVideoAddToAlbumBuilder() *VideoAddToAlbumBuilder {
	return &VideoAddToAlbumBuilder{api.Params{}}
}

// TargetID parameter
func (b *VideoAddToAlbumBuilder) TargetID(v int) {
	b.Params["target_id"] = v
}

// AlbumID parameter
func (b *VideoAddToAlbumBuilder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// AlbumIDs parameter
func (b *VideoAddToAlbumBuilder) AlbumIDs(v []int) {
	b.Params["album_ids"] = v
}

// OwnerID parameter
func (b *VideoAddToAlbumBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoID parameter
func (b *VideoAddToAlbumBuilder) VideoID(v int) {
	b.Params["video_id"] = v
}

// VideoCreateCommentBuilder builder
//
// Adds a new comment on a video.
//
// https://vk.com/dev/video.createComment
type VideoCreateCommentBuilder struct {
	api.Params
}

// NewVideoCreateCommentBuilder func
func NewVideoCreateCommentBuilder() *VideoCreateCommentBuilder {
	return &VideoCreateCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoCreateCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoID Video ID.
func (b *VideoCreateCommentBuilder) VideoID(v int) {
	b.Params["video_id"] = v
}

// Message New comment text.
func (b *VideoCreateCommentBuilder) Message(v string) {
	b.Params["message"] = v
}

// Attachments List of objects attached to the comment, in the following format:
// "<owner_id>_<media_id>,<owner_id>_<media_id>",
// '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document,
// '<owner_id>' — ID of the media attachment owner. '<media_id>' — Media attachment ID.
// Example: "photo100172_166443618,photo66748_265827614"
func (b *VideoCreateCommentBuilder) Attachments(v interface{}) {
	b.Params["attachments"] = v
}

// FromGroup '1' — to post the comment from a community name (only if 'owner_id'<0)
func (b *VideoCreateCommentBuilder) FromGroup(v bool) {
	b.Params["from_group"] = v
}

// ReplyToComment parameter
func (b *VideoCreateCommentBuilder) ReplyToComment(v int) {
	b.Params["reply_to_comment"] = v
}

// StickerID parameter
func (b *VideoCreateCommentBuilder) StickerID(v int) {
	b.Params["sticker_id"] = v
}

// GUID parameter
func (b *VideoCreateCommentBuilder) GUID(v string) {
	b.Params["guid"] = v
}

// VideoDeleteBuilder builder
//
// Deletes a video from a user or community page.
//
// https://vk.com/dev/video.delete
type VideoDeleteBuilder struct {
	api.Params
}

// NewVideoDeleteBuilder func
func NewVideoDeleteBuilder() *VideoDeleteBuilder {
	return &VideoDeleteBuilder{api.Params{}}
}

// VideoID Video ID.
func (b *VideoDeleteBuilder) VideoID(v int) {
	b.Params["video_id"] = v
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoDeleteBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// TargetID parameter
func (b *VideoDeleteBuilder) TargetID(v int) {
	b.Params["target_id"] = v
}

// VideoDeleteAlbumBuilder builder
//
// Deletes a video album.
//
// https://vk.com/dev/video.deleteAlbum
type VideoDeleteAlbumBuilder struct {
	api.Params
}

// NewVideoDeleteAlbumBuilder func
func NewVideoDeleteAlbumBuilder() *VideoDeleteAlbumBuilder {
	return &VideoDeleteAlbumBuilder{api.Params{}}
}

// GroupID Community ID (if the album is owned by a community).
func (b *VideoDeleteAlbumBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// AlbumID Album ID.
func (b *VideoDeleteAlbumBuilder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// VideoDeleteCommentBuilder builder
//
// Deletes a comment on a video.
//
// https://vk.com/dev/video.deleteComment
type VideoDeleteCommentBuilder struct {
	api.Params
}

// NewVideoDeleteCommentBuilder func
func NewVideoDeleteCommentBuilder() *VideoDeleteCommentBuilder {
	return &VideoDeleteCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoDeleteCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID ID of the comment to be deleted.
func (b *VideoDeleteCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// VideoEditBuilder builder
//
// Edits information about a video on a user or community page.
//
// https://vk.com/dev/video.edit
type VideoEditBuilder struct {
	api.Params
}

// NewVideoEditBuilder func
func NewVideoEditBuilder() *VideoEditBuilder {
	return &VideoEditBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoEditBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoID Video ID.
func (b *VideoEditBuilder) VideoID(v int) {
	b.Params["video_id"] = v
}

// Name New video title.
func (b *VideoEditBuilder) Name(v string) {
	b.Params["name"] = v
}

// Desc New video description.
func (b *VideoEditBuilder) Desc(v string) {
	b.Params["desc"] = v
}

// PrivacyView Privacy settings in a [vk.com/dev/privacy_setting|special format].
// Privacy setting is available for videos uploaded to own profile by user.
func (b *VideoEditBuilder) PrivacyView(v []string) {
	b.Params["privacy_view"] = v
}

// PrivacyComment Privacy settings for comments in a [vk.com/dev/privacy_setting|special format].
func (b *VideoEditBuilder) PrivacyComment(v []string) {
	b.Params["privacy_comment"] = v
}

// NoComments Disable comments for the group video.
func (b *VideoEditBuilder) NoComments(v bool) {
	b.Params["no_comments"] = v
}

// Repeat '1' — to repeat the playback of the video, '0' — to play the video once,
func (b *VideoEditBuilder) Repeat(v bool) {
	b.Params["repeat"] = v
}

// VideoEditAlbumBuilder builder
//
// Edits the title of a video album.
//
// https://vk.com/dev/video.editAlbum
type VideoEditAlbumBuilder struct {
	api.Params
}

// NewVideoEditAlbumBuilder func
func NewVideoEditAlbumBuilder() *VideoEditAlbumBuilder {
	return &VideoEditAlbumBuilder{api.Params{}}
}

// GroupID Community ID (if the album edited is owned by a community).
func (b *VideoEditAlbumBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// AlbumID Album ID.
func (b *VideoEditAlbumBuilder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// Title New album title.
func (b *VideoEditAlbumBuilder) Title(v string) {
	b.Params["title"] = v
}

// Privacy new access permissions for the album. Possible values:
//
// * 0 – all users,
//
// * 1 – friends only,
//
// * 2 – friends and friends of friends,
//
// * 3 – "only me".
func (b *VideoEditAlbumBuilder) Privacy(v []string) {
	b.Params["privacy"] = v
}

// VideoEditCommentBuilder builder
//
// Edits the text of a comment on a video.
//
// https://vk.com/dev/video.editComment
type VideoEditCommentBuilder struct {
	api.Params
}

// NewVideoEditCommentBuilder func
func NewVideoEditCommentBuilder() *VideoEditCommentBuilder {
	return &VideoEditCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoEditCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID Comment ID.
func (b *VideoEditCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// Message New comment text.
func (b *VideoEditCommentBuilder) Message(v string) {
	b.Params["message"] = v
}

// Attachments List of objects attached to the comment, in the following format:
// "<owner_id>_<media_id>,<owner_id>_<media_id>",
// '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document,
// '<owner_id>' — ID of the media attachment owner. '<media_id>' — Media attachment ID.
// Example: "photo100172_166443618,photo66748_265827614"
func (b *VideoEditCommentBuilder) Attachments(v interface{}) {
	b.Params["attachments"] = v
}

// VideoGetBuilder builder
//
// Returns detailed information about videos.
//
// https://vk.com/dev/video.get
type VideoGetBuilder struct {
	api.Params
}

// NewVideoGetBuilder func
func NewVideoGetBuilder() *VideoGetBuilder {
	return &VideoGetBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video(s).
func (b *VideoGetBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Videos Video IDs, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>",
// Use a negative value to designate a community ID.
// Example: "-4363_136089719,13245770_137352259"
func (b *VideoGetBuilder) Videos(v []string) {
	b.Params["videos"] = v
}

// AlbumID ID of the album containing the video(s).
func (b *VideoGetBuilder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// Count Number of videos to return.
func (b *VideoGetBuilder) Count(v int) {
	b.Params["count"] = v
}

// Offset Offset needed to return a specific subset of videos.
func (b *VideoGetBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Extended '1' — to return an extended response with additional fields
func (b *VideoGetBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// VideoGetAlbumByIDBuilder builder
//
// Returns video album info
//
// https://vk.com/dev/video.getAlbumById
type VideoGetAlbumByIDBuilder struct {
	api.Params
}

// NewVideoGetAlbumByIDBuilder func
func NewVideoGetAlbumByIDBuilder() *VideoGetAlbumByIDBuilder {
	return &VideoGetAlbumByIDBuilder{api.Params{}}
}

// OwnerID identifier of a user or community to add a video to. Use a negative value to designate a community ID.
func (b *VideoGetAlbumByIDBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumID Album ID.
func (b *VideoGetAlbumByIDBuilder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// VideoGetAlbumsBuilder builder
//
// Returns a list of video albums owned by a user or community.
//
// https://vk.com/dev/video.getAlbums
type VideoGetAlbumsBuilder struct {
	api.Params
}

// NewVideoGetAlbumsBuilder func
func NewVideoGetAlbumsBuilder() *VideoGetAlbumsBuilder {
	return &VideoGetAlbumsBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video album(s).
func (b *VideoGetAlbumsBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Offset Offset needed to return a specific subset of video albums.
func (b *VideoGetAlbumsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of video albums to return.
func (b *VideoGetAlbumsBuilder) Count(v int) {
	b.Params["count"] = v
}

// Extended '1' — to return additional information about album privacy settings for the current user
func (b *VideoGetAlbumsBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// NeedSystem parameter
func (b *VideoGetAlbumsBuilder) NeedSystem(v bool) {
	b.Params["need_system"] = v
}

// VideoGetAlbumsByVideoBuilder builder
//
// https://vk.com/dev/video.getAlbumsByVideo
type VideoGetAlbumsByVideoBuilder struct {
	api.Params
}

// NewVideoGetAlbumsByVideoBuilder func
func NewVideoGetAlbumsByVideoBuilder() *VideoGetAlbumsByVideoBuilder {
	return &VideoGetAlbumsByVideoBuilder{api.Params{}}
}

// TargetID parameter
func (b *VideoGetAlbumsByVideoBuilder) TargetID(v int) {
	b.Params["target_id"] = v
}

// OwnerID parameter
func (b *VideoGetAlbumsByVideoBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoID parameter
func (b *VideoGetAlbumsByVideoBuilder) VideoID(v int) {
	b.Params["video_id"] = v
}

// Extended parameter
func (b *VideoGetAlbumsByVideoBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// VideoGetCommentsBuilder builder
//
// Returns a list of comments on a video.
//
// https://vk.com/dev/video.getComments
type VideoGetCommentsBuilder struct {
	api.Params
}

// NewVideoGetCommentsBuilder func
func NewVideoGetCommentsBuilder() *VideoGetCommentsBuilder {
	return &VideoGetCommentsBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoGetCommentsBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoID Video ID.
func (b *VideoGetCommentsBuilder) VideoID(v int) {
	b.Params["video_id"] = v
}

// NeedLikes '1' — to return an additional 'likes' field
func (b *VideoGetCommentsBuilder) NeedLikes(v bool) {
	b.Params["need_likes"] = v
}

// StartCommentID parameter
func (b *VideoGetCommentsBuilder) StartCommentID(v int) {
	b.Params["start_comment_id"] = v
}

// Offset Offset needed to return a specific subset of comments.
func (b *VideoGetCommentsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of comments to return.
func (b *VideoGetCommentsBuilder) Count(v int) {
	b.Params["count"] = v
}

// Sort Sort order: 'asc' — oldest comment first, 'desc' — newest comment first
func (b *VideoGetCommentsBuilder) Sort(v string) {
	b.Params["sort"] = v
}

// Extended parameter
func (b *VideoGetCommentsBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields parameter
func (b *VideoGetCommentsBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// VideoRemoveFromAlbumBuilder builder
//
// https://vk.com/dev/video.removeFromAlbum
type VideoRemoveFromAlbumBuilder struct {
	api.Params
}

// NewVideoRemoveFromAlbumBuilder func
func NewVideoRemoveFromAlbumBuilder() *VideoRemoveFromAlbumBuilder {
	return &VideoRemoveFromAlbumBuilder{api.Params{}}
}

// TargetID parameter
func (b *VideoRemoveFromAlbumBuilder) TargetID(v int) {
	b.Params["target_id"] = v
}

// AlbumID parameter
func (b *VideoRemoveFromAlbumBuilder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// AlbumIDs parameter
func (b *VideoRemoveFromAlbumBuilder) AlbumIDs(v []int) {
	b.Params["album_ids"] = v
}

// OwnerID parameter
func (b *VideoRemoveFromAlbumBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoID parameter
func (b *VideoRemoveFromAlbumBuilder) VideoID(v int) {
	b.Params["video_id"] = v
}

// VideoReorderAlbumsBuilder builder
//
// Reorders the album in the list of user video albums.
//
// https://vk.com/dev/video.reorderAlbums
type VideoReorderAlbumsBuilder struct {
	api.Params
}

// NewVideoReorderAlbumsBuilder func
func NewVideoReorderAlbumsBuilder() *VideoReorderAlbumsBuilder {
	return &VideoReorderAlbumsBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the albums..
func (b *VideoReorderAlbumsBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumID Album ID.
func (b *VideoReorderAlbumsBuilder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// Before ID of the album before which the album in question shall be placed.
func (b *VideoReorderAlbumsBuilder) Before(v int) {
	b.Params["before"] = v
}

// After ID of the album after which the album in question shall be placed.
func (b *VideoReorderAlbumsBuilder) After(v int) {
	b.Params["after"] = v
}

// VideoReorderVideosBuilder builder
//
// Reorders the video in the video album.
//
// https://vk.com/dev/video.reorderVideos
type VideoReorderVideosBuilder struct {
	api.Params
}

// NewVideoReorderVideosBuilder func
func NewVideoReorderVideosBuilder() *VideoReorderVideosBuilder {
	return &VideoReorderVideosBuilder{api.Params{}}
}

// TargetID ID of the user or community that owns the album with videos.
func (b *VideoReorderVideosBuilder) TargetID(v int) {
	b.Params["target_id"] = v
}

// AlbumID ID of the video album.
func (b *VideoReorderVideosBuilder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoReorderVideosBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoID ID of the video.
func (b *VideoReorderVideosBuilder) VideoID(v int) {
	b.Params["video_id"] = v
}

// BeforeOwnerID ID of the user or community that owns the video before which the video in question shall be placed.
func (b *VideoReorderVideosBuilder) BeforeOwnerID(v int) {
	b.Params["before_owner_id"] = v
}

// BeforeVideoID ID of the video before which the video in question shall be placed.
func (b *VideoReorderVideosBuilder) BeforeVideoID(v int) {
	b.Params["before_video_id"] = v
}

// AfterOwnerID ID of the user or community that owns the video after which the photo in question shall be placed.
func (b *VideoReorderVideosBuilder) AfterOwnerID(v int) {
	b.Params["after_owner_id"] = v
}

// AfterVideoID ID of the video after which the photo in question shall be placed.
func (b *VideoReorderVideosBuilder) AfterVideoID(v int) {
	b.Params["after_video_id"] = v
}

// VideoReportBuilder builder
//
// Reports (submits a complaint about) a video.
//
// https://vk.com/dev/video.report
type VideoReportBuilder struct {
	api.Params
}

// NewVideoReportBuilder func
func NewVideoReportBuilder() *VideoReportBuilder {
	return &VideoReportBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoReportBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoID Video ID.
func (b *VideoReportBuilder) VideoID(v int) {
	b.Params["video_id"] = v
}

// Reason Reason for the complaint:
//
// * 0 – spam;
//
// * 1 – child pornography;
//
// * 2 – extremism;
//
// * 3 – violence;
//
// * 4 – drug propaganda;
//
// * 5 – adult material;
//
// * 6 – insult, abuse.
func (b *VideoReportBuilder) Reason(v int) {
	b.Params["reason"] = v
}

// Comment Comment describing the complaint.
func (b *VideoReportBuilder) Comment(v string) {
	b.Params["comment"] = v
}

// SearchQuery (If the video was found in search results.) Search query string.
func (b *VideoReportBuilder) SearchQuery(v string) {
	b.Params["search_query"] = v
}

// VideoReportCommentBuilder builder
//
// Reports (submits a complaint about) a comment on a video.
//
// https://vk.com/dev/video.reportComment
type VideoReportCommentBuilder struct {
	api.Params
}

// NewVideoReportCommentBuilder func
func NewVideoReportCommentBuilder() *VideoReportCommentBuilder {
	return &VideoReportCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoReportCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID ID of the comment being reported.
func (b *VideoReportCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// Reason Reason for the complaint:
//
// * 0 – spam;
//
// * 1 – child pornography;
//
// * 2 – extremism;
//
// * 3 – violence;
//
// * 4 – drug propaganda;
//
// * 5 – adult material;
//
// * 6 – insult, abuse.
func (b *VideoReportCommentBuilder) Reason(v int) {
	b.Params["reason"] = v
}

// VideoRestoreBuilder builder
//
// Restores a previously deleted video.
//
// https://vk.com/dev/video.restore
type VideoRestoreBuilder struct {
	api.Params
}

// NewVideoRestoreBuilder func
func NewVideoRestoreBuilder() *VideoRestoreBuilder {
	return &VideoRestoreBuilder{api.Params{}}
}

// VideoID Video ID.
func (b *VideoRestoreBuilder) VideoID(v int) {
	b.Params["video_id"] = v
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoRestoreBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// VideoRestoreCommentBuilder builder
//
// Restores a previously deleted comment on a video.
//
// https://vk.com/dev/video.restoreComment
type VideoRestoreCommentBuilder struct {
	api.Params
}

// NewVideoRestoreCommentBuilder func
func NewVideoRestoreCommentBuilder() *VideoRestoreCommentBuilder {
	return &VideoRestoreCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoRestoreCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID ID of the deleted comment.
func (b *VideoRestoreCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// VideoSaveBuilder builder
//
// Returns a server address (required for upload) and video data.
//
// https://vk.com/dev/video.save
type VideoSaveBuilder struct {
	api.Params
}

// NewVideoSaveBuilder func
func NewVideoSaveBuilder() *VideoSaveBuilder {
	return &VideoSaveBuilder{api.Params{}}
}

// Name Name of the video.
func (b *VideoSaveBuilder) Name(v string) {
	b.Params["name"] = v
}

// Description Description of the video.
func (b *VideoSaveBuilder) Description(v string) {
	b.Params["description"] = v
}

// IsPrivate parameter
//
// * 1 — to designate the video as private (send it via a private message), the video will not appear
// on the user's video list and will not be available by ID for other users,
//
// * 0 — not to designate the video as private
func (b *VideoSaveBuilder) IsPrivate(v bool) {
	b.Params["is_private"] = v
}

// Wallpost parameter
//
// * 1 — to post the saved video on a user's wall
// * 0 — not to post the saved video on a user's wall
func (b *VideoSaveBuilder) Wallpost(v bool) {
	b.Params["wallpost"] = v
}

// Link URL for embedding the video from an external website.
func (b *VideoSaveBuilder) Link(v string) {
	b.Params["link"] = v
}

// GroupID ID of the community in which the video will be saved. By default, the current user's page.
func (b *VideoSaveBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// AlbumID ID of the album to which the saved video will be added.
func (b *VideoSaveBuilder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// PrivacyView parameter
func (b *VideoSaveBuilder) PrivacyView(v []string) {
	b.Params["privacy_view"] = v
}

// PrivacyComment parameter
func (b *VideoSaveBuilder) PrivacyComment(v []string) {
	b.Params["privacy_comment"] = v
}

// NoComments parameter
func (b *VideoSaveBuilder) NoComments(v bool) {
	b.Params["no_comments"] = v
}

// Repeat '1' — to repeat the playback of the video, '0' — to play the video once,
func (b *VideoSaveBuilder) Repeat(v bool) {
	b.Params["repeat"] = v
}

// Compression parameter
func (b *VideoSaveBuilder) Compression(v bool) {
	b.Params["compression"] = v
}

// VideoSearchBuilder builder
//
// Returns a list of videos under the set search criterion.
//
// https://vk.com/dev/video.search
type VideoSearchBuilder struct {
	api.Params
}

// NewVideoSearchBuilder func
func NewVideoSearchBuilder() *VideoSearchBuilder {
	return &VideoSearchBuilder{api.Params{}}
}

// Q Search query string (e.g., 'The Beatles').
func (b *VideoSearchBuilder) Q(v string) {
	b.Params["q"] = v
}

// Sort Sort order: '1' — by duration, '2' — by relevance, '0' — by date added
func (b *VideoSearchBuilder) Sort(v int) {
	b.Params["sort"] = v
}

// Hd If not null, only searches for high-definition videos.
func (b *VideoSearchBuilder) Hd(v int) {
	b.Params["hd"] = v
}

// Adult '1' — to disable the Safe Search filter, '0' — to enable the Safe Search filter
func (b *VideoSearchBuilder) Adult(v bool) {
	b.Params["adult"] = v
}

// Filters Filters to apply:
//
// * youtube — return YouTube videos only;
//
// * vimeo — return Vimeo videos only;
//
// * short — return short videos only;
//
// * long — return long videos only.
func (b *VideoSearchBuilder) Filters(v []string) {
	b.Params["filters"] = v
}

// SearchOwn parameter
func (b *VideoSearchBuilder) SearchOwn(v bool) {
	b.Params["search_own"] = v
}

// Offset Offset needed to return a specific subset of videos.
func (b *VideoSearchBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Longer parameter
func (b *VideoSearchBuilder) Longer(v int) {
	b.Params["longer"] = v
}

// Shorter parameter
func (b *VideoSearchBuilder) Shorter(v int) {
	b.Params["shorter"] = v
}

// Count Number of videos to return.
func (b *VideoSearchBuilder) Count(v int) {
	b.Params["count"] = v
}

// Extended parameter
func (b *VideoSearchBuilder) Extended(v bool) {
	b.Params["extended"] = v
}
