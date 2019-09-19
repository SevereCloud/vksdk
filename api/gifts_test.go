package api

import (
	"testing"

	"github.com/SevereCloud/vksdk/errors"
	"github.com/stretchr/testify/assert"
)

func TestVK_GiftsGet(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.GiftsGet(map[string]string{
		"user_id": "1",
		"count":   "20",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)
	if assert.NotEmpty(t, res.Items) {
		for _, gift := range res.Items {
			assert.NotEmpty(t, gift.ID)
			// assert.NotEmpty(t, gift.FromID)
			// assert.NotEmpty(t, gift.Message)
			assert.NotEmpty(t, gift.Date)
			if assert.NotEmpty(t, gift.Gift) {
				assert.NotEmpty(t, gift.Gift.ID)
				assert.NotEmpty(t, gift.Gift.Thumb256)
				assert.NotEmpty(t, gift.Gift.Thumb96)
				assert.NotEmpty(t, gift.Gift.Thumb48)
			}
			// assert.NotEmpty(t, gift.GiftHash)
		}
	}
}

func TestVK_GiftsGetCatalog(t *testing.T) {
	t.Skip("Access denied: method allowed only for official app")
	needUserToken(t)

	_, err := vkUser.GiftsGetCatalog(map[string]string{})
	// NOTE: Access denied: method allowed only for official app
	if err != nil && errors.GetType(err) != errors.Access {
		t.Errorf("VK.GiftsGetCatalog() err = %v", err)
	}
}
