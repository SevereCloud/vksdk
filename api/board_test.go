package api

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_BoardAddTopic(t *testing.T) {
	needUserToken(t)

	topic, err := vkUser.BoardAddTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"title":    "Test topic",
		"text":     "Test",
	})
	assert.NoError(t, err)

	_, err = vkUser.BoardCloseTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	assert.NoError(t, err)

	_, err = vkUser.BoardOpenTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	assert.NoError(t, err)

	_, err = vkUser.BoardEditTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
		"title":    "Test topic edited",
		"text":     "Test edited",
	})
	assert.NoError(t, err)

	_, err = vkUser.BoardFixTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	assert.NoError(t, err)

	_, err = vkUser.BoardUnfixTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	assert.NoError(t, err)

	comment, err := vkUser.BoardCreateComment(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
		"message":  "topic comment",
	})
	assert.NoError(t, err)

	_, err = vkUser.BoardEditComment(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"topic_id":   strconv.Itoa(topic),
		"comment_id": strconv.Itoa(comment),
		"message":    "topic comment",
	})
	assert.NoError(t, err)

	_, err = vkUser.BoardDeleteComment(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"topic_id":   strconv.Itoa(topic),
		"comment_id": strconv.Itoa(comment),
	})
	assert.NoError(t, err)

	_, err = vkUser.BoardRestoreComment(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"topic_id":   strconv.Itoa(topic),
		"comment_id": strconv.Itoa(comment),
	})
	assert.NoError(t, err)

	_, err = vkUser.BoardDeleteTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	assert.NoError(t, err)
}

func TestVK_BoardGetComments(t *testing.T) {
	needUserToken(t)

	params := map[string]string{
		"group_id": "1",
		"topic_id": "21972169",
	}

	_, err := vkUser.BoardGetComments(params)
	assert.NoError(t, err)

	_, err = vkUser.BoardGetCommentsExtended(params)
	assert.NoError(t, err)
}

func TestVK_BoardGetTopics(t *testing.T) {
	needUserToken(t)

	params := map[string]string{
		"group_id": "1",
		"topic_id": "21972169",
	}

	_, err := vkUser.BoardGetTopics(params)
	assert.NoError(t, err)

	_, err = vkUser.BoardGetTopicsExtended(params)
	assert.NoError(t, err)
}
