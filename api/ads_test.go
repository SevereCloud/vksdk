package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/stretchr/testify/assert"
)

func TestAdsResponse_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(data []byte, expected api.AdsAddOfficeUsersItem) {
		var r api.AdsAddOfficeUsersItem

		r.UnmarshalJSON(data)

		assert.Equal(t, expected, r)
	}

	f([]byte("false"), api.AdsAddOfficeUsersItem{OK: false})
	f([]byte("true"), api.AdsAddOfficeUsersItem{OK: true})
	f(
		[]byte(`{"error_code": 100,"error_desc": "One of the parameters specified was missing or invalid: data[1][user_id]"}`),
		api.AdsAddOfficeUsersItem{
			OK: false,
			Error: api.AdsError{
				Code: 100,
				Desc: "One of the parameters specified was missing or invalid: data[1][user_id]",
			},
		},
	)
}

func TestVK_AdsGetAccounts(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.AdsGetAccounts(nil)
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

func TestVK_AdsCheckLink(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.AdsCheckLink(api.Params{
		"account_id": needAccountID(t),
		"link_type":  "post",
		"link_url":   "https://vk.com/wall-1_337125",
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}
