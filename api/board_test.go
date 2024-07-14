package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_BoardAddTopic(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	topic, err := vkUser.BoardAddTopic(api.Params{
		"group_id": vkGroupID,
		"title":    "Test topic",
		"text":     "Test",
	})
	noError(t, err)
	assert.NotEmpty(t, topic)

	res, err := vkUser.BoardCloseTopic(api.Params{
		"group_id": vkGroupID,
		"topic_id": topic,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.BoardOpenTopic(api.Params{
		"group_id": vkGroupID,
		"topic_id": topic,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.BoardEditTopic(api.Params{
		"group_id": vkGroupID,
		"topic_id": topic,
		"title":    "Test topic edited",
		"text":     "Test edited",
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.BoardFixTopic(api.Params{
		"group_id": vkGroupID,
		"topic_id": topic,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.BoardUnfixTopic(api.Params{
		"group_id": vkGroupID,
		"topic_id": topic,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	comment, err := vkUser.BoardCreateComment(api.Params{
		"group_id": vkGroupID,
		"topic_id": topic,
		"message":  "topic comment",
	})
	noError(t, err)
	assert.NotEmpty(t, comment)

	res, err = vkUser.BoardEditComment(api.Params{
		"group_id":   vkGroupID,
		"topic_id":   topic,
		"comment_id": comment,
		"message":    "topic comment",
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.BoardDeleteComment(api.Params{
		"group_id":   vkGroupID,
		"topic_id":   topic,
		"comment_id": comment,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.BoardRestoreComment(api.Params{
		"group_id":   vkGroupID,
		"topic_id":   topic,
		"comment_id": comment,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.BoardDeleteTopic(api.Params{
		"group_id": vkGroupID,
		"topic_id": topic,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_BoardGetComments(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	params := api.Params{
		"group_id":   1,
		"topic_id":   21972169,
		"need_likes": true,
	}

	res, err := vkUser.BoardGetComments(params)
	noError(t, err)
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
	noError(t, err)
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
		// BUG(VK): deactivated users not return ScreenName
		// assert.NotEmpty(t, res2.Profiles[0].ScreenName)
		assert.NotEmpty(t, res2.Profiles[0].Photo50)
		assert.NotEmpty(t, res2.Profiles[0].Photo100)
		assert.NotEmpty(t, res2.Profiles[0].OnlineInfo)
	}
}

func TestVK_BoardGetTopics(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	params := api.Params{
		"group_id":  1,
		"topic_ids": 21972169,
	}

	res, err := vkUser.BoardGetTopics(params)
	noError(t, err)
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
	noError(t, err)
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
