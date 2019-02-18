package api // import "github.com/severecloud/vksdk/api"

import "encoding/json"

// WallCloseCommentsResponse struct
type WallCloseCommentsResponse struct{}

// TODO: wall.closeComments

// WallCreateCommentResponse struct
type WallCreateCommentResponse struct {
	CommentID    int   `json:"comment_id"`
	ParentsStack []int `json:"parents_stack"`
}

// WallCreateComment Adds a comment to a post on a user wall or community wall.
func (vk VK) WallCreateComment(params map[string]string) (response WallCreateCommentResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("wall.createComment", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// WallDeleteResponse struct
type WallDeleteResponse struct{}

// TODO: wall.delete Deletes a post from a user wall or community wall.

// WallDeleteCommentResponse struct
type WallDeleteCommentResponse struct{}

// TODO: wall.deleteComment Deletes a comment on a post on a user wall or community wall.

// WallEditResponse struct
type WallEditResponse struct{}

// TODO: wall.edit Edits a post on a user wall or community wall.

// WallEditAdsStealthResponse struct
type WallEditAdsStealthResponse struct{}

// TODO: wall.editAdsStealth Allows to edit hidden post.

// WallEditCommentResponse struct
type WallEditCommentResponse struct{}

// TODO: wall.editComment Edits a comment on a user wall or community wall.

// WallGetResponse struct
type WallGetResponse struct{}

// TODO: wall.get Returns a list of posts on a user wall or community wall.

// WallGetByIDResponse struct
type WallGetByIDResponse struct{}

// TODO: wall.getById Returns a list of posts from user or community walls by their IDs.

// WallGetCommentResponse struct
type WallGetCommentResponse struct{}

// TODO: wall.getComment Allows to obtain wall comment info.

// WallGetCommentsResponse struct
type WallGetCommentsResponse struct{}

// TODO: wall.getComments Returns a list of comments on a post on a user wall or community wall.

// WallGetRepostsResponse struct
type WallGetRepostsResponse struct{}

// TODO: wall.getReposts Returns information about reposts of a post on user wall or community wall.

// WallOpenCommentsResponse struct
type WallOpenCommentsResponse struct{}

// TODO: wall.openComments

// WallPinResponse struct
type WallPinResponse struct{}

// TODO: wall.pin Pins the post on wall.

// WallPostResponse struct
type WallPostResponse struct {
	PostID int `json:"post_id"`
}

// WallPost Adds a new post on a user wall or community wall.Can also be used to publish suggested or scheduled posts.
func (vk VK) WallPost(params map[string]string) (response WallPostResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("wall.post", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// WallPostAdsStealthResponse struct
type WallPostAdsStealthResponse struct{}

// TODO: wall.postAdsStealth Allows to create hidden post which will not be shown on the community's wall and can be used for creating an ad with type "Community post".

// WallReportCommentResponse struct
type WallReportCommentResponse struct{}

// TODO: wall.reportComment Reports (submits a complaint about) a comment on a post on a user wall or community wall.

// WallReportPostResponse struct
type WallReportPostResponse struct{}

// TODO: wall.reportPost Reports (submits a complaint about) a post on a user wall or community wall.

// WallRepostResponse struct
type WallRepostResponse struct{}

// TODO: wall.repost Reposts ( copies) an object to a user wall or community wall.

// WallRestoreResponse struct
type WallRestoreResponse struct{}

// TODO: wall.restore Restores a post deleted from a user wall or community wall.

// WallRestoreCommentResponse struct
type WallRestoreCommentResponse struct{}

// TODO: wall.restoreComment Restores a comment deleted from a user wall or community wall.

// WallSearchResponse struct
type WallSearchResponse struct{}

// TODO: wall.search Allows to search posts on user or community walls.

// WallUnpinResponse struct
type WallUnpinResponse struct{}

// TODO: wall.unpin Unpins the post on wall.
