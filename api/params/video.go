package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// VideoAddBuilder builder.
//
// Adds a video to a user or community page.
//
// https://vk.com/dev/video.add
type VideoAddBuilder struct {
	api.Params
}

// NewVideoAddBuilder func.
func NewVideoAddBuilder() *VideoAddBuilder {
	return &VideoAddBuilder{api.Params{}}
}

// TargetID identifier of a user or community to add a video to. Use a negative value to designate a community ID.
func (b *VideoAddBuilder) TargetID(v int) *VideoAddBuilder {
	b.Params["target_id"] = v
	return b
}

// VideoID parameter.
func (b *VideoAddBuilder) VideoID(v int) *VideoAddBuilder {
	b.Params["video_id"] = v
	return b
}

// OwnerID ID of the user or community that owns the video. Use a negative value to designate a community ID.
func (b *VideoAddBuilder) OwnerID(v int) *VideoAddBuilder {
	b.Params["owner_id"] = v
	return b
}

// VideoAddAlbumBuilder builder.
//
// Creates an empty album for videos.
//
// https://vk.com/dev/video.addAlbum
type VideoAddAlbumBuilder struct {
	api.Params
}

// NewVideoAddAlbumBuilder func.
func NewVideoAddAlbumBuilder() *VideoAddAlbumBuilder {
	return &VideoAddAlbumBuilder{api.Params{}}
}

// GroupID community ID (if the album will be created in a community).
func (b *VideoAddAlbumBuilder) GroupID(v int) *VideoAddAlbumBuilder {
	b.Params["group_id"] = v
	return b
}

// Title album.
func (b *VideoAddAlbumBuilder) Title(v string) *VideoAddAlbumBuilder {
	b.Params["title"] = v
	return b
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
func (b *VideoAddAlbumBuilder) Privacy(v []string) *VideoAddAlbumBuilder {
	b.Params["privacy"] = v
	return b
}

// VideoAddToAlbumBuilder builder.
//
// https://vk.com/dev/video.addToAlbum
type VideoAddToAlbumBuilder struct {
	api.Params
}

// NewVideoAddToAlbumBuilder func.
func NewVideoAddToAlbumBuilder() *VideoAddToAlbumBuilder {
	return &VideoAddToAlbumBuilder{api.Params{}}
}

// TargetID parameter.
func (b *VideoAddToAlbumBuilder) TargetID(v int) *VideoAddToAlbumBuilder {
	b.Params["target_id"] = v
	return b
}

// AlbumID parameter.
func (b *VideoAddToAlbumBuilder) AlbumID(v int) *VideoAddToAlbumBuilder {
	b.Params["album_id"] = v
	return b
}

// AlbumIDs parameter.
func (b *VideoAddToAlbumBuilder) AlbumIDs(v []int) *VideoAddToAlbumBuilder {
	b.Params["album_ids"] = v
	return b
}

// OwnerID parameter.
func (b *VideoAddToAlbumBuilder) OwnerID(v int) *VideoAddToAlbumBuilder {
	b.Params["owner_id"] = v
	return b
}

// VideoID parameter.
func (b *VideoAddToAlbumBuilder) VideoID(v int) *VideoAddToAlbumBuilder {
	b.Params["video_id"] = v
	return b
}

// VideoCreateCommentBuilder builder.
//
// Adds a new comment on a video.
//
// https://vk.com/dev/video.createComment
type VideoCreateCommentBuilder struct {
	api.Params
}

// NewVideoCreateCommentBuilder func.
func NewVideoCreateCommentBuilder() *VideoCreateCommentBuilder {
	return &VideoCreateCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoCreateCommentBuilder) OwnerID(v int) *VideoCreateCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// VideoID parameter.
func (b *VideoCreateCommentBuilder) VideoID(v int) *VideoCreateCommentBuilder {
	b.Params["video_id"] = v
	return b
}

// Message new comment text.
func (b *VideoCreateCommentBuilder) Message(v string) *VideoCreateCommentBuilder {
	b.Params["message"] = v
	return b
}

// Attachments list of objects attached to the comment, in the following format:
// "<owner_id>_<media_id>,<owner_id>_<media_id>",
// ” — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document,
// '<owner_id>' — ID of the media attachment owner. '<media_id>' — Media attachment ID.
// Example: "photo100172_166443618,photo66748_265827614".
func (b *VideoCreateCommentBuilder) Attachments(v interface{}) *VideoCreateCommentBuilder {
	b.Params["attachments"] = v
	return b
}

// FromGroup '1' — to post the comment from a community name (only if 'owner_id'<0).
func (b *VideoCreateCommentBuilder) FromGroup(v bool) *VideoCreateCommentBuilder {
	b.Params["from_group"] = v
	return b
}

// ReplyToComment parameter.
func (b *VideoCreateCommentBuilder) ReplyToComment(v int) *VideoCreateCommentBuilder {
	b.Params["reply_to_comment"] = v
	return b
}

// StickerID parameter.
func (b *VideoCreateCommentBuilder) StickerID(v int) *VideoCreateCommentBuilder {
	b.Params["sticker_id"] = v
	return b
}

// GUID parameter.
func (b *VideoCreateCommentBuilder) GUID(v string) *VideoCreateCommentBuilder {
	b.Params["guid"] = v
	return b
}

// VideoDeleteBuilder builder.
//
// Deletes a video from a user or community page.
//
// https://vk.com/dev/video.delete
type VideoDeleteBuilder struct {
	api.Params
}

// NewVideoDeleteBuilder func.
func NewVideoDeleteBuilder() *VideoDeleteBuilder {
	return &VideoDeleteBuilder{api.Params{}}
}

// VideoID parameter.
func (b *VideoDeleteBuilder) VideoID(v int) *VideoDeleteBuilder {
	b.Params["video_id"] = v
	return b
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoDeleteBuilder) OwnerID(v int) *VideoDeleteBuilder {
	b.Params["owner_id"] = v
	return b
}

// TargetID parameter.
func (b *VideoDeleteBuilder) TargetID(v int) *VideoDeleteBuilder {
	b.Params["target_id"] = v
	return b
}

// VideoDeleteAlbumBuilder builder.
//
// Deletes a video album.
//
// https://vk.com/dev/video.deleteAlbum
type VideoDeleteAlbumBuilder struct {
	api.Params
}

// NewVideoDeleteAlbumBuilder func.
func NewVideoDeleteAlbumBuilder() *VideoDeleteAlbumBuilder {
	return &VideoDeleteAlbumBuilder{api.Params{}}
}

// GroupID community ID (if the album is owned by a community).
func (b *VideoDeleteAlbumBuilder) GroupID(v int) *VideoDeleteAlbumBuilder {
	b.Params["group_id"] = v
	return b
}

// AlbumID parameter.
func (b *VideoDeleteAlbumBuilder) AlbumID(v int) *VideoDeleteAlbumBuilder {
	b.Params["album_id"] = v
	return b
}

// VideoDeleteCommentBuilder builder.
//
// Deletes a comment on a video.
//
// https://vk.com/dev/video.deleteComment
type VideoDeleteCommentBuilder struct {
	api.Params
}

// NewVideoDeleteCommentBuilder func.
func NewVideoDeleteCommentBuilder() *VideoDeleteCommentBuilder {
	return &VideoDeleteCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoDeleteCommentBuilder) OwnerID(v int) *VideoDeleteCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// CommentID ID of the comment to be deleted.
func (b *VideoDeleteCommentBuilder) CommentID(v int) *VideoDeleteCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// VideoEditBuilder builder.
//
// Edits information about a video on a user or community page.
//
// https://vk.com/dev/video.edit
type VideoEditBuilder struct {
	api.Params
}

// NewVideoEditBuilder func.
func NewVideoEditBuilder() *VideoEditBuilder {
	return &VideoEditBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoEditBuilder) OwnerID(v int) *VideoEditBuilder {
	b.Params["owner_id"] = v
	return b
}

// VideoID parameter.
func (b *VideoEditBuilder) VideoID(v int) *VideoEditBuilder {
	b.Params["video_id"] = v
	return b
}

// Name new video title.
func (b *VideoEditBuilder) Name(v string) *VideoEditBuilder {
	b.Params["name"] = v
	return b
}

// Desc new video description.
func (b *VideoEditBuilder) Desc(v string) *VideoEditBuilder {
	b.Params["desc"] = v
	return b
}

// PrivacyView privacy settings in a [vk.com/dev/privacy_setting|special format].
// Privacy setting is available for videos uploaded to own profile by user.
func (b *VideoEditBuilder) PrivacyView(v []string) *VideoEditBuilder {
	b.Params["privacy_view"] = v
	return b
}

// PrivacyComment privacy settings for comments in a [vk.com/dev/privacy_setting|special format].
func (b *VideoEditBuilder) PrivacyComment(v []string) *VideoEditBuilder {
	b.Params["privacy_comment"] = v
	return b
}

// NoComments disable comments for the group video.
func (b *VideoEditBuilder) NoComments(v bool) *VideoEditBuilder {
	b.Params["no_comments"] = v
	return b
}

// Repeat '1' — to repeat the playback of the video, '0' — to play the video once.
func (b *VideoEditBuilder) Repeat(v bool) *VideoEditBuilder {
	b.Params["repeat"] = v
	return b
}

// VideoEditAlbumBuilder builder.
//
// Edits the title of a video album.
//
// https://vk.com/dev/video.editAlbum
type VideoEditAlbumBuilder struct {
	api.Params
}

// NewVideoEditAlbumBuilder func.
func NewVideoEditAlbumBuilder() *VideoEditAlbumBuilder {
	return &VideoEditAlbumBuilder{api.Params{}}
}

// GroupID community ID (if the album edited is owned by a community).
func (b *VideoEditAlbumBuilder) GroupID(v int) *VideoEditAlbumBuilder {
	b.Params["group_id"] = v
	return b
}

// AlbumID parameter.
func (b *VideoEditAlbumBuilder) AlbumID(v int) *VideoEditAlbumBuilder {
	b.Params["album_id"] = v
	return b
}

// Title new album title.
func (b *VideoEditAlbumBuilder) Title(v string) *VideoEditAlbumBuilder {
	b.Params["title"] = v
	return b
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
func (b *VideoEditAlbumBuilder) Privacy(v []string) *VideoEditAlbumBuilder {
	b.Params["privacy"] = v
	return b
}

// VideoEditCommentBuilder builder.
//
// Edits the text of a comment on a video.
//
// https://vk.com/dev/video.editComment
type VideoEditCommentBuilder struct {
	api.Params
}

// NewVideoEditCommentBuilder func.
func NewVideoEditCommentBuilder() *VideoEditCommentBuilder {
	return &VideoEditCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoEditCommentBuilder) OwnerID(v int) *VideoEditCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// CommentID parameter.
func (b *VideoEditCommentBuilder) CommentID(v int) *VideoEditCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// Message new comment text.
func (b *VideoEditCommentBuilder) Message(v string) *VideoEditCommentBuilder {
	b.Params["message"] = v
	return b
}

// Attachments list of objects attached to the comment, in the following format:
// "<owner_id>_<media_id>,<owner_id>_<media_id>",
// ” — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document,
// '<owner_id>' — ID of the media attachment owner. '<media_id>' — Media attachment ID.
// Example: "photo100172_166443618,photo66748_265827614".
func (b *VideoEditCommentBuilder) Attachments(v interface{}) *VideoEditCommentBuilder {
	b.Params["attachments"] = v
	return b
}

// VideoGetBuilder builder.
//
// Returns detailed information about videos.
//
// https://vk.com/dev/video.get
type VideoGetBuilder struct {
	api.Params
}

// NewVideoGetBuilder func.
func NewVideoGetBuilder() *VideoGetBuilder {
	return &VideoGetBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video(s).
func (b *VideoGetBuilder) OwnerID(v int) *VideoGetBuilder {
	b.Params["owner_id"] = v
	return b
}

// Videos IDs, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>",
// Use a negative value to designate a community ID.
// Example: "-4363_136089719,13245770_137352259".
func (b *VideoGetBuilder) Videos(v []string) *VideoGetBuilder {
	b.Params["videos"] = v
	return b
}

// AlbumID ID of the album containing the video(s).
func (b *VideoGetBuilder) AlbumID(v int) *VideoGetBuilder {
	b.Params["album_id"] = v
	return b
}

// Count number of videos to return.
func (b *VideoGetBuilder) Count(v int) *VideoGetBuilder {
	b.Params["count"] = v
	return b
}

// Offset needed to return a specific subset of videos.
func (b *VideoGetBuilder) Offset(v int) *VideoGetBuilder {
	b.Params["offset"] = v
	return b
}

// Extended '1' — to return an extended response with additional fields.
func (b *VideoGetBuilder) Extended(v bool) *VideoGetBuilder {
	b.Params["extended"] = v
	return b
}

// VideoGetAlbumByIDBuilder builder.
//
// Returns video album info.
//
// https://vk.com/dev/video.getAlbumById
type VideoGetAlbumByIDBuilder struct {
	api.Params
}

// NewVideoGetAlbumByIDBuilder func.
func NewVideoGetAlbumByIDBuilder() *VideoGetAlbumByIDBuilder {
	return &VideoGetAlbumByIDBuilder{api.Params{}}
}

// OwnerID identifier of a user or community to add a video to. Use a negative value to designate a community ID.
func (b *VideoGetAlbumByIDBuilder) OwnerID(v int) *VideoGetAlbumByIDBuilder {
	b.Params["owner_id"] = v
	return b
}

// AlbumID parameter.
func (b *VideoGetAlbumByIDBuilder) AlbumID(v int) *VideoGetAlbumByIDBuilder {
	b.Params["album_id"] = v
	return b
}

// VideoGetAlbumsBuilder builder.
//
// Returns a list of video albums owned by a user or community.
//
// https://vk.com/dev/video.getAlbums
type VideoGetAlbumsBuilder struct {
	api.Params
}

// NewVideoGetAlbumsBuilder func.
func NewVideoGetAlbumsBuilder() *VideoGetAlbumsBuilder {
	return &VideoGetAlbumsBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video album(s).
func (b *VideoGetAlbumsBuilder) OwnerID(v int) *VideoGetAlbumsBuilder {
	b.Params["owner_id"] = v
	return b
}

// Offset needed to return a specific subset of video albums.
func (b *VideoGetAlbumsBuilder) Offset(v int) *VideoGetAlbumsBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of video albums to return.
func (b *VideoGetAlbumsBuilder) Count(v int) *VideoGetAlbumsBuilder {
	b.Params["count"] = v
	return b
}

// Extended '1' — to return additional information about album privacy settings for the current user.
func (b *VideoGetAlbumsBuilder) Extended(v bool) *VideoGetAlbumsBuilder {
	b.Params["extended"] = v
	return b
}

// NeedSystem parameter.
func (b *VideoGetAlbumsBuilder) NeedSystem(v bool) *VideoGetAlbumsBuilder {
	b.Params["need_system"] = v
	return b
}

// VideoGetAlbumsByVideoBuilder builder.
//
// https://vk.com/dev/video.getAlbumsByVideo
type VideoGetAlbumsByVideoBuilder struct {
	api.Params
}

// NewVideoGetAlbumsByVideoBuilder func.
func NewVideoGetAlbumsByVideoBuilder() *VideoGetAlbumsByVideoBuilder {
	return &VideoGetAlbumsByVideoBuilder{api.Params{}}
}

// TargetID parameter.
func (b *VideoGetAlbumsByVideoBuilder) TargetID(v int) *VideoGetAlbumsByVideoBuilder {
	b.Params["target_id"] = v
	return b
}

// OwnerID parameter.
func (b *VideoGetAlbumsByVideoBuilder) OwnerID(v int) *VideoGetAlbumsByVideoBuilder {
	b.Params["owner_id"] = v
	return b
}

// VideoID parameter.
func (b *VideoGetAlbumsByVideoBuilder) VideoID(v int) *VideoGetAlbumsByVideoBuilder {
	b.Params["video_id"] = v
	return b
}

// Extended parameter.
func (b *VideoGetAlbumsByVideoBuilder) Extended(v bool) *VideoGetAlbumsByVideoBuilder {
	b.Params["extended"] = v
	return b
}

// VideoGetCommentsBuilder builder.
//
// Returns a list of comments on a video.
//
// https://vk.com/dev/video.getComments
type VideoGetCommentsBuilder struct {
	api.Params
}

// NewVideoGetCommentsBuilder func.
func NewVideoGetCommentsBuilder() *VideoGetCommentsBuilder {
	return &VideoGetCommentsBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoGetCommentsBuilder) OwnerID(v int) *VideoGetCommentsBuilder {
	b.Params["owner_id"] = v
	return b
}

// VideoID parameter.
func (b *VideoGetCommentsBuilder) VideoID(v int) *VideoGetCommentsBuilder {
	b.Params["video_id"] = v
	return b
}

// NeedLikes '1' — to return an additional 'likes' field.
func (b *VideoGetCommentsBuilder) NeedLikes(v bool) *VideoGetCommentsBuilder {
	b.Params["need_likes"] = v
	return b
}

// StartCommentID parameter.
func (b *VideoGetCommentsBuilder) StartCommentID(v int) *VideoGetCommentsBuilder {
	b.Params["start_comment_id"] = v
	return b
}

// Offset needed to return a specific subset of comments.
func (b *VideoGetCommentsBuilder) Offset(v int) *VideoGetCommentsBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of comments to return.
func (b *VideoGetCommentsBuilder) Count(v int) *VideoGetCommentsBuilder {
	b.Params["count"] = v
	return b
}

// Sort order: 'asc' — oldest comment first, 'desc' — newest comment first.
func (b *VideoGetCommentsBuilder) Sort(v string) *VideoGetCommentsBuilder {
	b.Params["sort"] = v
	return b
}

// Extended parameter.
func (b *VideoGetCommentsBuilder) Extended(v bool) *VideoGetCommentsBuilder {
	b.Params["extended"] = v
	return b
}

// Fields parameter.
func (b *VideoGetCommentsBuilder) Fields(v []string) *VideoGetCommentsBuilder {
	b.Params["fields"] = v
	return b
}

// VideoRemoveFromAlbumBuilder builder.
//
// https://vk.com/dev/video.removeFromAlbum
type VideoRemoveFromAlbumBuilder struct {
	api.Params
}

// NewVideoRemoveFromAlbumBuilder func.
func NewVideoRemoveFromAlbumBuilder() *VideoRemoveFromAlbumBuilder {
	return &VideoRemoveFromAlbumBuilder{api.Params{}}
}

// TargetID parameter.
func (b *VideoRemoveFromAlbumBuilder) TargetID(v int) *VideoRemoveFromAlbumBuilder {
	b.Params["target_id"] = v
	return b
}

// AlbumID parameter.
func (b *VideoRemoveFromAlbumBuilder) AlbumID(v int) *VideoRemoveFromAlbumBuilder {
	b.Params["album_id"] = v
	return b
}

// AlbumIDs parameter.
func (b *VideoRemoveFromAlbumBuilder) AlbumIDs(v []int) *VideoRemoveFromAlbumBuilder {
	b.Params["album_ids"] = v
	return b
}

// OwnerID parameter.
func (b *VideoRemoveFromAlbumBuilder) OwnerID(v int) *VideoRemoveFromAlbumBuilder {
	b.Params["owner_id"] = v
	return b
}

// VideoID parameter.
func (b *VideoRemoveFromAlbumBuilder) VideoID(v int) *VideoRemoveFromAlbumBuilder {
	b.Params["video_id"] = v
	return b
}

// VideoReorderAlbumsBuilder builder.
//
// Reorders the album in the list of user video albums.
//
// https://vk.com/dev/video.reorderAlbums
type VideoReorderAlbumsBuilder struct {
	api.Params
}

// NewVideoReorderAlbumsBuilder func.
func NewVideoReorderAlbumsBuilder() *VideoReorderAlbumsBuilder {
	return &VideoReorderAlbumsBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the albums..
func (b *VideoReorderAlbumsBuilder) OwnerID(v int) *VideoReorderAlbumsBuilder {
	b.Params["owner_id"] = v
	return b
}

// AlbumID parameter.
func (b *VideoReorderAlbumsBuilder) AlbumID(v int) *VideoReorderAlbumsBuilder {
	b.Params["album_id"] = v
	return b
}

// Before ID of the album before which the album in question shall be placed.
func (b *VideoReorderAlbumsBuilder) Before(v int) *VideoReorderAlbumsBuilder {
	b.Params["before"] = v
	return b
}

// After ID of the album after which the album in question shall be placed.
func (b *VideoReorderAlbumsBuilder) After(v int) *VideoReorderAlbumsBuilder {
	b.Params["after"] = v
	return b
}

// VideoReorderVideosBuilder builder.
//
// Reorders the video in the video album.
//
// https://vk.com/dev/video.reorderVideos
type VideoReorderVideosBuilder struct {
	api.Params
}

// NewVideoReorderVideosBuilder func.
func NewVideoReorderVideosBuilder() *VideoReorderVideosBuilder {
	return &VideoReorderVideosBuilder{api.Params{}}
}

// TargetID ID of the user or community that owns the album with videos.
func (b *VideoReorderVideosBuilder) TargetID(v int) *VideoReorderVideosBuilder {
	b.Params["target_id"] = v
	return b
}

// AlbumID ID of the video album.
func (b *VideoReorderVideosBuilder) AlbumID(v int) *VideoReorderVideosBuilder {
	b.Params["album_id"] = v
	return b
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoReorderVideosBuilder) OwnerID(v int) *VideoReorderVideosBuilder {
	b.Params["owner_id"] = v
	return b
}

// VideoID ID of the video.
func (b *VideoReorderVideosBuilder) VideoID(v int) *VideoReorderVideosBuilder {
	b.Params["video_id"] = v
	return b
}

// BeforeOwnerID ID of the user or community that owns the video before which the video in question shall be placed.
func (b *VideoReorderVideosBuilder) BeforeOwnerID(v int) *VideoReorderVideosBuilder {
	b.Params["before_owner_id"] = v
	return b
}

// BeforeVideoID ID of the video before which the video in question shall be placed.
func (b *VideoReorderVideosBuilder) BeforeVideoID(v int) *VideoReorderVideosBuilder {
	b.Params["before_video_id"] = v
	return b
}

// AfterOwnerID ID of the user or community that owns the video after which the photo in question shall be placed.
func (b *VideoReorderVideosBuilder) AfterOwnerID(v int) *VideoReorderVideosBuilder {
	b.Params["after_owner_id"] = v
	return b
}

// AfterVideoID ID of the video after which the photo in question shall be placed.
func (b *VideoReorderVideosBuilder) AfterVideoID(v int) *VideoReorderVideosBuilder {
	b.Params["after_video_id"] = v
	return b
}

// VideoReportBuilder builder.
//
// Reports (submits a complaint about) a video.
//
// https://vk.com/dev/video.report
type VideoReportBuilder struct {
	api.Params
}

// NewVideoReportBuilder func.
func NewVideoReportBuilder() *VideoReportBuilder {
	return &VideoReportBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoReportBuilder) OwnerID(v int) *VideoReportBuilder {
	b.Params["owner_id"] = v
	return b
}

// VideoID parameter.
func (b *VideoReportBuilder) VideoID(v int) *VideoReportBuilder {
	b.Params["video_id"] = v
	return b
}

// Reason for the complaint:
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
func (b *VideoReportBuilder) Reason(v int) *VideoReportBuilder {
	b.Params["reason"] = v
	return b
}

// Comment describing the complaint.
func (b *VideoReportBuilder) Comment(v string) *VideoReportBuilder {
	b.Params["comment"] = v
	return b
}

// SearchQuery (If the video was found in search results.) Search query string.
func (b *VideoReportBuilder) SearchQuery(v string) *VideoReportBuilder {
	b.Params["search_query"] = v
	return b
}

// VideoReportCommentBuilder builder.
//
// Reports (submits a complaint about) a comment on a video.
//
// https://vk.com/dev/video.reportComment
type VideoReportCommentBuilder struct {
	api.Params
}

// NewVideoReportCommentBuilder func.
func NewVideoReportCommentBuilder() *VideoReportCommentBuilder {
	return &VideoReportCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoReportCommentBuilder) OwnerID(v int) *VideoReportCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// CommentID ID of the comment being reported.
func (b *VideoReportCommentBuilder) CommentID(v int) *VideoReportCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// Reason for the complaint:
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
func (b *VideoReportCommentBuilder) Reason(v int) *VideoReportCommentBuilder {
	b.Params["reason"] = v
	return b
}

// VideoRestoreBuilder builder.
//
// Restores a previously deleted video.
//
// https://vk.com/dev/video.restore
type VideoRestoreBuilder struct {
	api.Params
}

// NewVideoRestoreBuilder func.
func NewVideoRestoreBuilder() *VideoRestoreBuilder {
	return &VideoRestoreBuilder{api.Params{}}
}

// VideoID parameter.
func (b *VideoRestoreBuilder) VideoID(v int) *VideoRestoreBuilder {
	b.Params["video_id"] = v
	return b
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoRestoreBuilder) OwnerID(v int) *VideoRestoreBuilder {
	b.Params["owner_id"] = v
	return b
}

// VideoRestoreCommentBuilder builder.
//
// Restores a previously deleted comment on a video.
//
// https://vk.com/dev/video.restoreComment
type VideoRestoreCommentBuilder struct {
	api.Params
}

// NewVideoRestoreCommentBuilder func.
func NewVideoRestoreCommentBuilder() *VideoRestoreCommentBuilder {
	return &VideoRestoreCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the video.
func (b *VideoRestoreCommentBuilder) OwnerID(v int) *VideoRestoreCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// CommentID ID of the deleted comment.
func (b *VideoRestoreCommentBuilder) CommentID(v int) *VideoRestoreCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// VideoSaveBuilder builder.
//
// Returns a server address (required for upload) and video data.
//
// https://vk.com/dev/video.save
type VideoSaveBuilder struct {
	api.Params
}

// NewVideoSaveBuilder func.
func NewVideoSaveBuilder() *VideoSaveBuilder {
	return &VideoSaveBuilder{api.Params{}}
}

// Name of the video.
func (b *VideoSaveBuilder) Name(v string) *VideoSaveBuilder {
	b.Params["name"] = v
	return b
}

// Description of the video.
func (b *VideoSaveBuilder) Description(v string) *VideoSaveBuilder {
	b.Params["description"] = v
	return b
}

// IsUnitedVideoUpload parameter.
func (b *VideoSaveBuilder) IsUnitedVideoUpload(v bool) *VideoSaveBuilder {
	b.Params["is_united_video_upload"] = v
	return b
}

// IsPrivate parameter.
//
// * 1 — to designate the video as private (send it via a private message), the video will not appear
// on the user's video list and will not be available by ID for other users,
//
// * 0 — not to designate the video as private.
func (b *VideoSaveBuilder) IsPrivate(v bool) *VideoSaveBuilder {
	b.Params["is_private"] = v
	return b
}

// Wallpost parameter.
//
// * 1 — to post the saved video on a user's wall
// * 0 — not to post the saved video on a user's wall.
func (b *VideoSaveBuilder) Wallpost(v bool) *VideoSaveBuilder {
	b.Params["wallpost"] = v
	return b
}

// Link URL for embedding the video from an external website.
func (b *VideoSaveBuilder) Link(v string) *VideoSaveBuilder {
	b.Params["link"] = v
	return b
}

// GroupID ID of the community in which the video will be saved. By default, the current user's page.
func (b *VideoSaveBuilder) GroupID(v int) *VideoSaveBuilder {
	b.Params["group_id"] = v
	return b
}

// AlbumID ID of the album to which the saved video will be added.
func (b *VideoSaveBuilder) AlbumID(v int) *VideoSaveBuilder {
	b.Params["album_id"] = v
	return b
}

// PrivacyView parameter.
func (b *VideoSaveBuilder) PrivacyView(v []string) *VideoSaveBuilder {
	b.Params["privacy_view"] = v
	return b
}

// PrivacyComment parameter.
func (b *VideoSaveBuilder) PrivacyComment(v []string) *VideoSaveBuilder {
	b.Params["privacy_comment"] = v
	return b
}

// NoComments parameter.
func (b *VideoSaveBuilder) NoComments(v bool) *VideoSaveBuilder {
	b.Params["no_comments"] = v
	return b
}

// Repeat '1' — to repeat the playback of the video, '0' — to play the video once.
func (b *VideoSaveBuilder) Repeat(v bool) *VideoSaveBuilder {
	b.Params["repeat"] = v
	return b
}

// Compression parameter.
func (b *VideoSaveBuilder) Compression(v bool) *VideoSaveBuilder {
	b.Params["compression"] = v
	return b
}

// VideoSearchBuilder builder.
//
// Returns a list of videos under the set search criterion.
//
// https://vk.com/dev/video.search
type VideoSearchBuilder struct {
	api.Params
}

// NewVideoSearchBuilder func.
func NewVideoSearchBuilder() *VideoSearchBuilder {
	return &VideoSearchBuilder{api.Params{}}
}

// Q search query string (e.g., 'The Beatles').
func (b *VideoSearchBuilder) Q(v string) *VideoSearchBuilder {
	b.Params["q"] = v
	return b
}

// Sort order: '1' — by duration, '2' — by relevance, '0' — by date added.
func (b *VideoSearchBuilder) Sort(v int) *VideoSearchBuilder {
	b.Params["sort"] = v
	return b
}

// HD if not null, only searches for high-definition videos.
func (b *VideoSearchBuilder) HD(v int) *VideoSearchBuilder {
	b.Params["hd"] = v
	return b
}

// Adult '1' — to disable the Safe Search filter, '0' — to enable the Safe Search filter.
func (b *VideoSearchBuilder) Adult(v bool) *VideoSearchBuilder {
	b.Params["adult"] = v
	return b
}

// Filters to apply:
//
// * youtube — return YouTube videos only;
//
// * vimeo — return Vimeo videos only;
//
// * short — return short videos only;
//
// * long — return long videos only.
func (b *VideoSearchBuilder) Filters(v []string) *VideoSearchBuilder {
	b.Params["filters"] = v
	return b
}

// SearchOwn parameter.
func (b *VideoSearchBuilder) SearchOwn(v bool) *VideoSearchBuilder {
	b.Params["search_own"] = v
	return b
}

// Offset needed to return a specific subset of videos.
func (b *VideoSearchBuilder) Offset(v int) *VideoSearchBuilder {
	b.Params["offset"] = v
	return b
}

// Longer parameter.
func (b *VideoSearchBuilder) Longer(v int) *VideoSearchBuilder {
	b.Params["longer"] = v
	return b
}

// Shorter parameter.
func (b *VideoSearchBuilder) Shorter(v int) *VideoSearchBuilder {
	b.Params["shorter"] = v
	return b
}

// Count number of videos to return.
func (b *VideoSearchBuilder) Count(v int) *VideoSearchBuilder {
	b.Params["count"] = v
	return b
}

// Extended parameter.
func (b *VideoSearchBuilder) Extended(v bool) *VideoSearchBuilder {
	b.Params["extended"] = v
	return b
}

// Legal parameter.
func (b *VideoSearchBuilder) Legal(v bool) *VideoSearchBuilder {
	b.Params["legal"] = v
	return b
}

// LegalOwner parameter.
func (b *VideoSearchBuilder) LegalOwner(v int) *VideoSearchBuilder {
	b.Params["legal_owner"] = v
	return b
}

// VideoStartStreamingBuilder builder.
//
// https://vk.com/dev/video.startStreaming
type VideoStartStreamingBuilder struct {
	api.Params
}

// NewVideoStartStreamingBuilder func.
func NewVideoStartStreamingBuilder() *VideoStartStreamingBuilder {
	return &VideoStartStreamingBuilder{api.Params{}}
}

// VideoID parameter.
func (b *VideoStartStreamingBuilder) VideoID(v int) *VideoStartStreamingBuilder {
	b.Params["video_id"] = v
	return b
}

// Name parameter.
func (b *VideoStartStreamingBuilder) Name(v string) *VideoStartStreamingBuilder {
	b.Params["name"] = v
	return b
}

// Description parameter.
func (b *VideoStartStreamingBuilder) Description(v string) *VideoStartStreamingBuilder {
	b.Params["description"] = v
	return b
}

// Wallpost parameter.
func (b *VideoStartStreamingBuilder) Wallpost(v bool) *VideoStartStreamingBuilder {
	b.Params["wallpost"] = v
	return b
}

// GroupID parameter.
func (b *VideoStartStreamingBuilder) GroupID(v int) *VideoStartStreamingBuilder {
	b.Params["group_id"] = v
	return b
}

// PrivacyView parameter.
func (b *VideoStartStreamingBuilder) PrivacyView(v []string) *VideoStartStreamingBuilder {
	b.Params["privacy_view"] = v
	return b
}

// PrivacyComment parameter.
func (b *VideoStartStreamingBuilder) PrivacyComment(v []string) *VideoStartStreamingBuilder {
	b.Params["privacy_comment"] = v
	return b
}

// NoComments parameter.
func (b *VideoStartStreamingBuilder) NoComments(v bool) *VideoStartStreamingBuilder {
	b.Params["no_comments"] = v
	return b
}

// CategoryID parameter.
func (b *VideoStartStreamingBuilder) CategoryID(v int) *VideoStartStreamingBuilder {
	b.Params["category_id"] = v
	return b
}

// Publish parameter.
func (b *VideoStartStreamingBuilder) Publish(v bool) *VideoStartStreamingBuilder {
	b.Params["publish"] = v
	return b
}

// VideoStopStreamingBuilder builder.
//
// https://vk.com/dev/video.stopStreaming
type VideoStopStreamingBuilder struct {
	api.Params
}

// NewVideoStopStreamingBuilder func.
func NewVideoStopStreamingBuilder() *VideoStopStreamingBuilder {
	return &VideoStopStreamingBuilder{api.Params{}}
}

// GroupID parameter.
func (b *VideoStopStreamingBuilder) GroupID(v int) *VideoStopStreamingBuilder {
	b.Params["group_id"] = v
	return b
}

// OwnerID parameter.
func (b *VideoStopStreamingBuilder) OwnerID(v int) *VideoStopStreamingBuilder {
	b.Params["owner_id"] = v
	return b
}

// VideoID parameter.
func (b *VideoStopStreamingBuilder) VideoID(v int) *VideoStopStreamingBuilder {
	b.Params["video_id"] = v
	return b
}
