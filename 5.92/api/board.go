package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// BoardAddTopicResponse struct
type BoardAddTopicResponse int

// BoardAddTopic creates a new topic on a community's discussion board.
//
// https://vk.com/dev/board.addTopic
func (vk *VK) BoardAddTopic(params map[string]string) (response BoardAddTopicResponse, vkErr Error) {
	vk.RequestUnmarshal("board.addTopic", params, &response, &vkErr)
	return
}

// BoardCloseTopicResponse struct
type BoardCloseTopicResponse int

// BoardCloseTopic closes a topic on a community's discussion board so that comments cannot be posted.
//
// https://vk.com/dev/board.closeTopic
func (vk *VK) BoardCloseTopic(params map[string]string) (response BoardCloseTopicResponse, vkErr Error) {
	vk.RequestUnmarshal("board.closeTopic", params, &response, &vkErr)
	return
}

// BoardCreateCommentResponse struct
type BoardCreateCommentResponse int

// BoardCreateComment adds a comment on a topic on a community's discussion board.
//
// https://vk.com/dev/board.createComment
func (vk *VK) BoardCreateComment(params map[string]string) (response BoardCreateCommentResponse, vkErr Error) {
	vk.RequestUnmarshal("board.createComment", params, &response, &vkErr)
	return
}

// BoardDeleteCommentResponse struct
type BoardDeleteCommentResponse int

// BoardDeleteComment deletes a comment on a topic on a community's discussion board.
//
// https://vk.com/dev/board.deleteComment
func (vk *VK) BoardDeleteComment(params map[string]string) (response BoardDeleteCommentResponse, vkErr Error) {
	vk.RequestUnmarshal("board.deleteComment", params, &response, &vkErr)
	return
}

// BoardDeleteTopicResponse struct
type BoardDeleteTopicResponse int

// BoardDeleteTopic deletes a topic from a community's discussion board.
//
// https://vk.com/dev/board.deleteTopic
func (vk *VK) BoardDeleteTopic(params map[string]string) (response BoardDeleteTopicResponse, vkErr Error) {
	vk.RequestUnmarshal("board.deleteTopic", params, &response, &vkErr)
	return
}

// BoardEditCommentResponse struct
type BoardEditCommentResponse int

// BoardEditComment edits a comment on a topic on a community's discussion board.
//
// https://vk.com/dev/board.editComment
func (vk *VK) BoardEditComment(params map[string]string) (response BoardEditCommentResponse, vkErr Error) {
	vk.RequestUnmarshal("board.editComment", params, &response, &vkErr)
	return
}

// BoardEditTopicResponse struct
type BoardEditTopicResponse int

// BoardEditTopic edits the title of a topic on a community's discussion board.
//
// https://vk.com/dev/board.editTopic
func (vk *VK) BoardEditTopic(params map[string]string) (response BoardEditTopicResponse, vkErr Error) {
	vk.RequestUnmarshal("board.editTopic", params, &response, &vkErr)
	return
}

// BoardFixTopicResponse struct
type BoardFixTopicResponse int

// BoardFixTopic pins a topic (fixes its place) to the top of a community's discussion board.
//
// https://vk.com/dev/board.fixTopic
func (vk *VK) BoardFixTopic(params map[string]string) (response BoardFixTopicResponse, vkErr Error) {
	vk.RequestUnmarshal("board.fixTopic", params, &response, &vkErr)
	return
}

// BoardGetCommentsResponse struct
type BoardGetCommentsResponse struct {
	Count      int                        `json:"count"`
	Items      []object.BoardTopicComment `json:"items"`
	Poll       object.BoardTopicPoll      `json:"poll"`
	RealOffset int                        `json:"real_offset"`
}

// BoardGetComments returns a list of comments on a topic on a community's discussion board.
//
// extended=0
//
// https://vk.com/dev/board.getComments
func (vk *VK) BoardGetComments(params map[string]string) (response BoardGetCommentsResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("board.getComments", params, &response, &vkErr)
	return
}

// BoardGetCommentsExtendedResponse struct
type BoardGetCommentsExtendedResponse struct {
	Count      int                        `json:"count"`
	Items      []object.BoardTopicComment `json:"items"`
	Poll       object.BoardTopicPoll      `json:"poll"`
	RealOffset int                        `json:"real_offset"`
	Profiles   []object.UsersUser         `json:"profiles"`
	Groups     []object.GroupsGroup       `json:"groups"`
}

// BoardGetCommentsExtended returns a list of comments on a topic on a community's discussion board.
//
// extended=1
//
// https://vk.com/dev/board.getComments
func (vk *VK) BoardGetCommentsExtended(params map[string]string) (response BoardGetCommentsExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("board.getComments", params, &response, &vkErr)
	return
}

// BoardGetTopicsResponse struct
type BoardGetTopicsResponse struct {
	Count        int                 `json:"count"`
	Items        []object.BoardTopic `json:"items"`
	DefaultOrder int                 `json:"default_order"`
	CanAddTopics int                 `json:"can_add_topics"`
}

// BoardGetTopics returns a list of topics on a community's discussion board.
//
// extended=0
//
// https://vk.com/dev/board.getTopics
func (vk *VK) BoardGetTopics(params map[string]string) (response BoardGetTopicsResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("board.getTopics", params, &response, &vkErr)
	return
}

// BoardGetTopicsExtendedResponse struct
type BoardGetTopicsExtendedResponse struct {
	Count        int                  `json:"count"`
	Items        []object.BoardTopic  `json:"items"`
	DefaultOrder int                  `json:"default_order"`
	CanAddTopics int                  `json:"can_add_topics"`
	Profiles     []object.UsersUser   `json:"profiles"`
	Groups       []object.GroupsGroup `json:"groups"`
}

// BoardGetTopicsExtended returns a list of topics on a community's discussion board.
//
// extended=1
//
// https://vk.com/dev/board.getTopics
func (vk *VK) BoardGetTopicsExtended(params map[string]string) (response BoardGetTopicsExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("board.getTopics", params, &response, &vkErr)
	return
}

// BoardOpenTopicResponse struct
type BoardOpenTopicResponse int

// BoardOpenTopic re-opens a previously closed topic on a community's discussion board.
//
// https://vk.com/dev/board.openTopic
func (vk *VK) BoardOpenTopic(params map[string]string) (response BoardOpenTopicResponse, vkErr Error) {
	vk.RequestUnmarshal("board.openTopic", params, &response, &vkErr)
	return
}

// BoardRestoreCommentResponse struct
type BoardRestoreCommentResponse int

// BoardRestoreComment restores a comment deleted from a topic on a community's discussion board.
//
// https://vk.com/dev/board.restoreComment
func (vk *VK) BoardRestoreComment(params map[string]string) (response BoardRestoreCommentResponse, vkErr Error) {
	vk.RequestUnmarshal("board.restoreComment", params, &response, &vkErr)
	return
}

// BoardUnfixTopicRResponse struct
type BoardUnfixTopicRResponse int

// BoardUnfixTopicR unpins a pinned topic from the top of a community's discussion board.
//
// https://vk.com/dev/board.unfixTopic
func (vk *VK) BoardUnfixTopicR(params map[string]string) (response BoardUnfixTopicRResponse, vkErr Error) {
	vk.RequestUnmarshal("board.unfixTopic", params, &response, &vkErr)
	return
}
