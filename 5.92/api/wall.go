package api // import "github.com/SevereCloud/vksdk/5.92/api"

// WallCloseCommentsResponse struct
type WallCloseCommentsResponse struct{}

// TODO: wall.closeComments x
//
// https://vk.com/dev/wall.closeComments

// WallCreateCommentResponse struct
type WallCreateCommentResponse struct {
	CommentID    int   `json:"comment_id"`
	ParentsStack []int `json:"parents_stack"`
}

// WallCreateComment Adds a comment to a post on a user wall or community wall.
//
// https://vk.com/dev/wall.createComment
func (vk *VK) WallCreateComment(params map[string]string) (response WallCreateCommentResponse, vkErr Error) {
	vk.RequestUnmarshal("wall.createComment", params, &response, &vkErr)
	return
}

// WallDeleteResponse struct
type WallDeleteResponse struct{}

// TODO: wall.delete deletes a post from a user wall or community wall.
//
// https://vk.com/dev/wall.delete

// WallDeleteCommentResponse struct
type WallDeleteCommentResponse struct{}

// TODO: wall.deleteComment deletes a comment on a post on a user wall or community wall.
//
// https://vk.com/dev/wall.deleteComment

// WallEditResponse struct
type WallEditResponse struct{}

// TODO: wall.edit edits a post on a user wall or community wall.
//
// https://vk.com/dev/wall.edit

// WallEditAdsStealthResponse struct
type WallEditAdsStealthResponse struct{}

// TODO: wall.editAdsStealth allows to edit hidden post.
//
// https://vk.com/dev/wall.editAdsStealth

// WallEditCommentResponse struct
type WallEditCommentResponse struct{}

// TODO: wall.editComment edits a comment on a user wall or community wall.
//
// https://vk.com/dev/wall.editComment

// WallGetResponse struct
type WallGetResponse struct{}

// TODO: wall.get returns a list of posts on a user wall or community wall.
//
// https://vk.com/dev/wall.get

// WallGetByIDResponse struct
type WallGetByIDResponse struct{}

// TODO: wall.getById returns a list of posts from user or community walls by their IDs.
//
// https://vk.com/dev/wall.getById

// WallGetCommentResponse struct
type WallGetCommentResponse struct{}

// TODO: wall.getComment allows to obtain wall comment info.
//
// https://vk.com/dev/wall.getComment

// WallGetCommentsResponse struct
type WallGetCommentsResponse struct{}

// TODO: wall.getComments returns a list of comments on a post on a user wall or community wall.
//
// https://vk.com/dev/wall.getComments

// WallGetRepostsResponse struct
type WallGetRepostsResponse struct{}

// TODO: wall.getReposts returns information about reposts of a post on user wall or community wall.
//
// https://vk.com/dev/wall.getReposts

// WallOpenCommentsResponse struct
type WallOpenCommentsResponse struct{}

// TODO: wall.openComments x
//
// https://vk.com/dev/wall.openComments

// WallPinResponse struct
type WallPinResponse struct{}

// TODO: wall.pin pins the post on wall.
//
// https://vk.com/dev/wall.pin

// WallPostResponse struct
type WallPostResponse struct {
	PostID int `json:"post_id"`
}

// WallPost adds a new post on a user wall or community wall.Can also be used to publish suggested or scheduled posts.
//
// https://vk.com/dev/wall.post
func (vk *VK) WallPost(params map[string]string) (response WallPostResponse, vkErr Error) {
	vk.RequestUnmarshal("wall.post", params, &response, &vkErr)
	return
}

// WallPostAdsStealthResponse struct
type WallPostAdsStealthResponse struct{}

// TODO: wall.postAdsStealth allows to create hidden post which will
// not be shown on the community's wall and can be used for creating
// an ad with type "Community post".
//
// https://vk.com/dev/wall.postAdsStealth

// WallReportCommentResponse struct
type WallReportCommentResponse struct{}

// TODO: wall.reportComment reports (submits a complaint about) a comment on a post on a user wall or community wall.
//
// https://vk.com/dev/wall.reportComment

// WallReportPostResponse struct
type WallReportPostResponse struct{}

// TODO: wall.reportPost reports (submits a complaint about) a post on a user wall or community wall.
//
// https://vk.com/dev/wall.reportPost

// WallRepostResponse struct
type WallRepostResponse struct{}

// TODO: wall.repost reposts ( copies) an object to a user wall or community wall.
//
// https://vk.com/dev/wall.repost

// WallRestoreResponse struct
type WallRestoreResponse struct{}

// TODO: wall.restore restores a post deleted from a user wall or community wall.
//
// https://vk.com/dev/wall.restore

// WallRestoreCommentResponse struct
type WallRestoreCommentResponse struct{}

// TODO: wall.restoreComment restores a comment deleted from a user wall or community wall.
//
// https://vk.com/dev/wall.restoreComment

// WallSearchResponse struct
type WallSearchResponse struct{}

// TODO: wall.search allows to search posts on user or community walls.
//
// https://vk.com/dev/wall.search

// WallUnpinResponse struct
type WallUnpinResponse struct{}

// TODO: wall.unpin unpins the post on wall.
//
// https://vk.com/dev/wall.unpin
