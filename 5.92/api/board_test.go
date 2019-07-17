package api

import (
	"strconv"
	"testing"
)

func TestVK_BoardAddTopic(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	topic, gotVkErr := vkUser.BoardAddTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"title":    "Test topic",
		"text":     "Test",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardAddTopic() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.BoardCloseTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(int(topic)),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardCloseTopic() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.BoardOpenTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(int(topic)),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardOpenTopic() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.BoardEditTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(int(topic)),
		"title":    "Test topic edited",
		"text":     "Test edited",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardEditTopic() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.BoardFixTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(int(topic)),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardFixTopic() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.BoardUnfixTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(int(topic)),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardUnfixTopic() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	comment, gotVkErr := vkUser.BoardCreateComment(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(int(topic)),
		"message":  "topic comment",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardCreateComment() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.BoardEditComment(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"topic_id":   strconv.Itoa(int(topic)),
		"comment_id": strconv.Itoa(int(comment)),
		"message":    "topic comment",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardEditComment() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.BoardDeleteComment(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"topic_id":   strconv.Itoa(int(topic)),
		"comment_id": strconv.Itoa(int(comment)),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardDeleteComment() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.BoardRestoreComment(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"topic_id":   strconv.Itoa(int(topic)),
		"comment_id": strconv.Itoa(int(comment)),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardRestoreComment() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.BoardDeleteTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(int(topic)),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardDeleteTopic() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_BoardGetComments(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	params := map[string]string{
		"group_id": "1",
		"topic_id": "21972169",
	}

	_, gotVkErr := vkUser.BoardGetComments(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardGetComments() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.BoardGetCommentsExtended(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardGetCommentsExtended() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_BoardGetTopics(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	params := map[string]string{
		"group_id": "1",
		"topic_id": "21972169",
	}

	_, gotVkErr := vkUser.BoardGetTopics(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardGetTopics() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.BoardGetTopicsExtended(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardGetTopicsExtended() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}
