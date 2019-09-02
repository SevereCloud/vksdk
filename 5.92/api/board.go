package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// BoardAddTopic creates a new topic on a community's discussion board.
//
// https://vk.com/dev/board.addTopic
func (vk *VK) BoardAddTopic(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("board.addTopic", params, &response, &vkErr)
	return
}

// BoardCloseTopic closes a topic on a community's discussion board so that comments cannot be posted.
//
// https://vk.com/dev/board.closeTopic
func (vk *VK) BoardCloseTopic(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("board.closeTopic", params, &response, &vkErr)
	return
}

// BoardCreateComment adds a comment on a topic on a community's discussion board.
//
// https://vk.com/dev/board.createComment
func (vk *VK) BoardCreateComment(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("board.createComment", params, &response, &vkErr)
	return
}

// BoardDeleteComment deletes a comment on a topic on a community's discussion board.
//
// https://vk.com/dev/board.deleteComment
func (vk *VK) BoardDeleteComment(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("board.deleteComment", params, &response, &vkErr)
	return
}

// BoardDeleteTopic deletes a topic from a community's discussion board.
//
// https://vk.com/dev/board.deleteTopic
func (vk *VK) BoardDeleteTopic(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("board.deleteTopic", params, &response, &vkErr)
	return
}

// BoardEditComment edits a comment on a topic on a community's discussion board.
//
// https://vk.com/dev/board.editComment
func (vk *VK) BoardEditComment(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("board.editComment", params, &response, &vkErr)
	return
}

// BoardEditTopic edits the title of a topic on a community's discussion board.
//
// https://vk.com/dev/board.editTopic
func (vk *VK) BoardEditTopic(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("board.editTopic", params, &response, &vkErr)
	return
}

// BoardFixTopic pins a topic (fixes its place) to the top of a community's discussion board.
//
// https://vk.com/dev/board.fixTopic
func (vk *VK) BoardFixTopic(params map[string]string) (response int, vkErr Error) {
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
	DefaultOrder float64             `json:"default_order"` // BUG(VK): default_order int https://vk.com/bug136682
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
	DefaultOrder float64              `json:"default_order"` // BUG(VK): default_order int https://vk.com/bug136682
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

// BoardOpenTopic re-opens a previously closed topic on a community's discussion board.
//
// https://vk.com/dev/board.openTopic
func (vk *VK) BoardOpenTopic(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("board.openTopic", params, &response, &vkErr)
	return
}

// BoardRestoreComment restores a comment deleted from a topic on a community's discussion board.
//
// https://vk.com/dev/board.restoreComment
func (vk *VK) BoardRestoreComment(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("board.restoreComment", params, &response, &vkErr)
	return
}

// BoardUnfixTopic unpins a pinned topic from the top of a community's discussion board.
//
// https://vk.com/dev/board.unfixTopic
func (vk *VK) BoardUnfixTopic(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("board.unfixTopic", params, &response, &vkErr)
	return
}
