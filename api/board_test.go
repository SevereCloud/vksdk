package api

import (
	"strconv"
	"testing"
)

func TestVK_BoardAddTopic(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	topic, err := vkUser.BoardAddTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"title":    "Test topic",
		"text":     "Test",
	})
	if err != nil {
		t.Errorf("VK.BoardAddTopic() err = %v", err)
	}

	_, err = vkUser.BoardCloseTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	if err != nil {
		t.Errorf("VK.BoardCloseTopic() err = %v", err)
	}

	_, err = vkUser.BoardOpenTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	if err != nil {
		t.Errorf("VK.BoardOpenTopic() err = %v", err)
	}

	_, err = vkUser.BoardEditTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
		"title":    "Test topic edited",
		"text":     "Test edited",
	})
	if err != nil {
		t.Errorf("VK.BoardEditTopic() err = %v", err)
	}

	_, err = vkUser.BoardFixTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	if err != nil {
		t.Errorf("VK.BoardFixTopic() err = %v", err)
	}

	_, err = vkUser.BoardUnfixTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	if err != nil {
		t.Errorf("VK.BoardUnfixTopic() err = %v", err)
	}

	comment, err := vkUser.BoardCreateComment(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
		"message":  "topic comment",
	})
	if err != nil {
		t.Errorf("VK.BoardCreateComment() err = %v", err)
	}

	_, err = vkUser.BoardEditComment(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"topic_id":   strconv.Itoa(topic),
		"comment_id": strconv.Itoa(comment),
		"message":    "topic comment",
	})
	if err != nil {
		t.Errorf("VK.BoardEditComment() err = %v", err)
	}

	_, err = vkUser.BoardDeleteComment(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"topic_id":   strconv.Itoa(topic),
		"comment_id": strconv.Itoa(comment),
	})
	if err != nil {
		t.Errorf("VK.BoardDeleteComment() err = %v", err)
	}

	_, err = vkUser.BoardRestoreComment(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"topic_id":   strconv.Itoa(topic),
		"comment_id": strconv.Itoa(comment),
	})
	if err != nil {
		t.Errorf("VK.BoardRestoreComment() err = %v", err)
	}

	_, err = vkUser.BoardDeleteTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	if err != nil {
		t.Errorf("VK.BoardDeleteTopic() err = %v", err)
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

	_, err := vkUser.BoardGetComments(params)
	if err != nil {
		t.Errorf("VK.BoardGetComments() err = %v", err)
	}

	_, err = vkUser.BoardGetCommentsExtended(params)
	if err != nil {
		t.Errorf("VK.BoardGetCommentsExtended() err = %v", err)
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

	_, err := vkUser.BoardGetTopics(params)
	if err != nil {
		t.Errorf("VK.BoardGetTopics() err = %v", err)
	}

	_, err = vkUser.BoardGetTopicsExtended(params)
	if err != nil {
		t.Errorf("VK.BoardGetTopicsExtended() err = %v", err)
	}
}
