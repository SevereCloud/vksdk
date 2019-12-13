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
	assert.NotEmpty(t, topic)

	res, err := vkUser.BoardCloseTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.BoardOpenTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.BoardEditTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
		"title":    "Test topic edited",
		"text":     "Test edited",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.BoardFixTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.BoardUnfixTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	comment, err := vkUser.BoardCreateComment(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
		"message":  "topic comment",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, comment)

	res, err = vkUser.BoardEditComment(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"topic_id":   strconv.Itoa(topic),
		"comment_id": strconv.Itoa(comment),
		"message":    "topic comment",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.BoardDeleteComment(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"topic_id":   strconv.Itoa(topic),
		"comment_id": strconv.Itoa(comment),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.BoardRestoreComment(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"topic_id":   strconv.Itoa(topic),
		"comment_id": strconv.Itoa(comment),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.BoardDeleteTopic(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"topic_id": strconv.Itoa(topic),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_BoardGetComments(t *testing.T) {
	needUserToken(t)

	params := map[string]string{
		"group_id":   "1",
		"topic_id":   "21972169",
		"need_likes": "1",
	}

	res, err := vkUser.BoardGetComments(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].FromID)
		assert.NotEmpty(t, res.Items[0].Date)
		assert.NotEmpty(t, res.Items[0].Text)
		assert.NotEmpty(t, res.Items[0].Likes)
		assert.NotEmpty(t, res.Items[0].Likes.CanLike)
	}

	res2, err := vkUser.BoardGetCommentsExtended(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, res2.Count)

	if assert.NotEmpty(t, res2.Items) {
		assert.NotEmpty(t, res2.Items[0].ID)
		assert.NotEmpty(t, res2.Items[0].FromID)
		assert.NotEmpty(t, res2.Items[0].Date)
		assert.NotEmpty(t, res2.Items[0].Text)
		assert.NotEmpty(t, res2.Items[0].Likes)
		// assert.NotEmpty(t, res.Items[0].CanEdit)
		assert.NotEmpty(t, res.Items[0].Likes.CanLike)
	}

	if assert.NotEmpty(t, res2.Profiles) {
		assert.NotEmpty(t, res2.Profiles[0].ID)
		assert.NotEmpty(t, res2.Profiles[0].FirstName)
		assert.NotEmpty(t, res2.Profiles[0].LastName)
		assert.NotEmpty(t, res2.Profiles[0].Sex)
		assert.NotEmpty(t, res2.Profiles[0].ScreenName)
		assert.NotEmpty(t, res2.Profiles[0].Photo50)
		assert.NotEmpty(t, res2.Profiles[0].Photo100)
		assert.NotEmpty(t, res2.Profiles[0].OnlineInfo)
	}
}

func TestVK_BoardGetTopics(t *testing.T) {
	needUserToken(t)

	params := map[string]string{
		"group_id":  "1",
		"topic_ids": "21972169",
	}

	res, err := vkUser.BoardGetTopics(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Title)
		assert.NotEmpty(t, res.Items[0].Created)
		assert.NotEmpty(t, res.Items[0].CreatedBy)
		assert.NotEmpty(t, res.Items[0].Updated)
		assert.NotEmpty(t, res.Items[0].UpdatedBy)
		// assert.NotEmpty(t, res.Items[0].IsClosed)
		// assert.NotEmpty(t, res.Items[0].IsFixed)
		assert.NotEmpty(t, res.Items[0].Comments)
	}

	assert.NotEmpty(t, res.DefaultOrder)
	// assert.NotEmpty(t, res.CanAddTopics)

	res2, err := vkUser.BoardGetTopicsExtended(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, res2.Count)

	if assert.NotEmpty(t, res2.Items) {
		assert.NotEmpty(t, res2.Items[0].ID)
		assert.NotEmpty(t, res2.Items[0].Title)
		assert.NotEmpty(t, res2.Items[0].Created)
		assert.NotEmpty(t, res2.Items[0].CreatedBy)
		assert.NotEmpty(t, res2.Items[0].Updated)
		assert.NotEmpty(t, res2.Items[0].UpdatedBy)
		// assert.NotEmpty(t, res2.Items[0].IsClosed)
		// assert.NotEmpty(t, res2.Items[0].IsFixed)
		assert.NotEmpty(t, res2.Items[0].Comments)
	}

	assert.NotEmpty(t, res2.DefaultOrder)
	// assert.NotEmpty(t, res2.CanAddTopics)
	if assert.NotEmpty(t, res2.Profiles) {
		assert.NotEmpty(t, res2.Profiles[0].ID)
		assert.NotEmpty(t, res2.Profiles[0].FirstName)
		assert.NotEmpty(t, res2.Profiles[0].LastName)
		assert.NotEmpty(t, res2.Profiles[0].Sex)
		assert.NotEmpty(t, res2.Profiles[0].ScreenName)
		assert.NotEmpty(t, res2.Profiles[0].Photo50)
		assert.NotEmpty(t, res2.Profiles[0].Photo100)
		assert.NotEmpty(t, res2.Profiles[0].OnlineInfo)
	}
}
