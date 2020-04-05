package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/stretchr/testify/assert"
)

func TestVK_WidgetsGetComments(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.WidgetsGetComments(api.Params{
		"widget_api_id": 6862945,
		"url":           "http://irsay.ru/irsay_duine/",
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Posts) {
		post := res.Posts[0]
		assert.NotEmpty(t, post.ID)
		assert.NotEmpty(t, post.FromID)
		assert.NotEmpty(t, post.ToID)
		assert.NotEmpty(t, post.Date)
		assert.NotEmpty(t, post.PostType)
		assert.NotEmpty(t, post.Text)
		assert.NotEmpty(t, post.PostSource)
		// assert.NotEmpty(t, post.Likes)
		// assert.NotEmpty(t, post.Reposts)
		// assert.NotEmpty(t, post.Views)
		// assert.NotEmpty(t, post.IsFavorite)
		assert.NotEmpty(t, post.Comments)
	}
}

func TestVK_WidgetsGetPages(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.WidgetsGetPages(api.Params{})
	noError(t, err)
	assert.NotEmpty(t, res)
}
