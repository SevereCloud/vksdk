package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/vmihailenco/msgpack/v5/msgpcode"
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

func TestAdsResponse_DecodeMsgpack(t *testing.T) {
	t.Parallel()

	f := func(data []byte, expected api.AdsAddOfficeUsersItem, wantErr string) {
		var r api.AdsAddOfficeUsersItem

		err := msgpack.Unmarshal(data, &r)
		if err != nil || wantErr != "" {
			assert.EqualError(t, err, wantErr)
		}

		assert.Equal(t, expected, r)
	}

	f([]byte{msgpcode.False}, api.AdsAddOfficeUsersItem{OK: false}, "")
	f([]byte{msgpcode.True}, api.AdsAddOfficeUsersItem{OK: true}, "")
	f(
		[]byte{
			0x82, 0xAA, 0x65, 0x72, 0x72, 0x6F, 0x72, 0x5F, 0x63, 0x6F, 0x64,
			0x65, 0x64, 0xAA, 0x65, 0x72, 0x72, 0x6F, 0x72, 0x5F, 0x64, 0x65,
			0x73, 0x63, 0xD9, 0x48, 0x4F, 0x6E, 0x65, 0x20, 0x6F, 0x66, 0x20,
			0x74, 0x68, 0x65, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6D, 0x65, 0x74,
			0x65, 0x72, 0x73, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
			0x65, 0x64, 0x20, 0x77, 0x61, 0x73, 0x20, 0x6D, 0x69, 0x73, 0x73,
			0x69, 0x6E, 0x67, 0x20, 0x6F, 0x72, 0x20, 0x69, 0x6E, 0x76, 0x61,
			0x6C, 0x69, 0x64, 0x3A, 0x20, 0x64, 0x61, 0x74, 0x61, 0x5B, 0x31,
			0x5D, 0x5B, 0x75, 0x73, 0x65, 0x72, 0x5F, 0x69, 0x64, 0x5D,
		},
		api.AdsAddOfficeUsersItem{
			OK: false,
			Error: api.AdsError{
				Code: 100,
				Desc: "One of the parameters specified was missing or invalid: data[1][user_id]",
			},
		},
		"",
	)
	f(nil, api.AdsAddOfficeUsersItem{}, "EOF")
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
