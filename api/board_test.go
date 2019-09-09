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
		t.Errorf("VK.BoardAddTopic() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.BoardCloseTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardCloseTopic() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.BoardOpenTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardOpenTopic() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.BoardEditTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
		"title":    "Test topic edited",
		"text":     "Test edited",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardEditTopic() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.BoardFixTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardFixTopic() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.BoardUnfixTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardUnfixTopic() gotVkErr = %v", gotVkErr)
	}

	comment, gotVkErr := vkUser.BoardCreateComment(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
		"message":  "topic comment",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardCreateComment() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.BoardEditComment(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"topic_id":   strconv.Itoa(topic),
		"comment_id": strconv.Itoa(comment),
		"message":    "topic comment",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardEditComment() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.BoardDeleteComment(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"topic_id":   strconv.Itoa(topic),
		"comment_id": strconv.Itoa(comment),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardDeleteComment() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.BoardRestoreComment(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"topic_id":   strconv.Itoa(topic),
		"comment_id": strconv.Itoa(comment),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardRestoreComment() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.BoardDeleteTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardDeleteTopic() gotVkErr = %v", gotVkErr)
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
		t.Errorf("VK.BoardGetComments() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.BoardGetCommentsExtended(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardGetCommentsExtended() gotVkErr = %v", gotVkErr)
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
		t.Errorf("VK.BoardGetTopics() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.BoardGetTopicsExtended(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.BoardGetTopicsExtended() gotVkErr = %v", gotVkErr)
	}
}
