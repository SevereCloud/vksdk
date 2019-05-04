package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"encoding/json"

	"github.com/SevereCloud/vksdk/5.92/object"
)

// BoardAddTopicResponse struct
type BoardAddTopicResponse int

// BoardAddTopic creates a new topic on a community's discussion board.
// https://vk.com/dev/board.addTopic
func (vk VK) BoardAddTopic(params map[string]string) (response BoardAddTopicResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("board.addTopic", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// BoardCloseTopic closes a topic on a community's discussion board so that comments cannot be posted.
// https://vk.com/dev/board.closeTopic
func (vk VK) BoardCloseTopic(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("board.closeTopic", params)

	return
}

// BoardCreateCommentResponse struct
type BoardCreateCommentResponse int

// BoardCreateComment adds a comment on a topic on a community's discussion board.
// https://vk.com/dev/board.createComment
func (vk VK) BoardCreateComment(params map[string]string) (response BoardCreateCommentResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("board.createComment", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// BoardDeleteComment deletes a comment on a topic on a community's discussion board.
// https://vk.com/dev/board.deleteComment
func (vk VK) BoardDeleteComment(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("board.deleteComment", params)

	return
}

// BoardDeleteTopic deletes a topic from a community's discussion board.
// https://vk.com/dev/board.deleteTopic
func (vk VK) BoardDeleteTopic(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("board.deleteTopic", params)

	return
}

// BoardEditComment edits a comment on a topic on a community's discussion board.
// https://vk.com/dev/board.editComment
func (vk VK) BoardEditComment(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("board.editComment", params)

	return
}

// BoardEditTopic edits the title of a topic on a community's discussion board.
// https://vk.com/dev/board.editTopic
func (vk VK) BoardEditTopic(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("board.editTopic", params)

	return
}

// BoardFixTopic pins a topic (fixes its place) to the top of a community's discussion board.
// https://vk.com/dev/board.fixTopic
func (vk VK) BoardFixTopic(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("board.fixTopic", params)

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
// https://vk.com/dev/board.getComments
func (vk VK) BoardGetComments(params map[string]string) (response BoardGetCommentsResponse, vkErr Error) {
	params["extended"] = "0"
	rawResponse, vkErr := vk.Request("board.getComments", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

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
// https://vk.com/dev/board.getComments
func (vk VK) BoardGetCommentsExtended(params map[string]string) (response BoardGetCommentsExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	rawResponse, vkErr := vk.Request("board.getComments", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

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
// https://vk.com/dev/board.getTopics
func (vk VK) BoardGetTopics(params map[string]string) (response BoardGetTopicsResponse, vkErr Error) {
	params["extended"] = "0"
	rawResponse, vkErr := vk.Request("board.getTopics", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

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
// https://vk.com/dev/board.getTopics
func (vk VK) BoardGetTopicsExtended(params map[string]string) (response BoardGetTopicsExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	rawResponse, vkErr := vk.Request("board.getTopics", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// BoardOpenTopic re-opens a previously closed topic on a community's discussion board.
// https://vk.com/dev/board.openTopic
func (vk VK) BoardOpenTopic(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("board.openTopic", params)

	return
}

// BoardRestoreComment restores a comment deleted from a topic on a community's discussion board.
// https://vk.com/dev/board.restoreComment
func (vk VK) BoardRestoreComment(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("board.restoreComment", params)

	return
}

// BoardUnfixTopicR unpins a pinned topic from the top of a community's discussion board.
// https://vk.com/dev/board.unfixTopic
func (vk VK) BoardUnfixTopicR(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("board.unfixTopic", params)

	return
}
