package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/stretchr/testify/assert"
)

func TestVK_AdsGetAccounts(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.AdsGetAccounts(api.Params{})
	noError(t, err)
}

func TestVK_AdsGetMusicians(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.AdsGetMusicians(api.Params{
		"artist_name": "Лазарев",
	})
	noError(t, err)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Name)
	}
}
