package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_GiftsGet(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.GiftsGet(api.Params{
		"user_id": vkUserID,
		"count":   20,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		for _, gift := range res.Items {
			assert.NotEmpty(t, gift.ID)
			// assert.NotEmpty(t, gift.FromID)
			// assert.NotEmpty(t, gift.Message)
			// assert.NotEmpty(t, gift.GiftHash)
			assert.NotEmpty(t, gift.Date)

			if assert.NotEmpty(t, gift.Gift) {
				assert.NotEmpty(t, gift.Gift.ID)
				assert.NotEmpty(t, gift.Gift.Thumb256)
				assert.NotEmpty(t, gift.Gift.Thumb96)
				assert.NotEmpty(t, gift.Gift.Thumb48)
			}
		}
	}
}

func TestVK_GiftsGetCatalog(t *testing.T) {
	t.Parallel()

	t.Skip("Access denied: method allowed only for official app")
	needUserToken(t)

	// NOTE: Access denied: method allowed only for official app
	_, err := vkUser.GiftsGetCatalog(nil)
	if e, ok := err.(*api.Error); ok && e.Code == api.ErrAccess {
		t.Errorf("VK.GiftsGetCatalog() err = %v", err)
	}
}
