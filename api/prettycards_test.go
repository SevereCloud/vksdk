package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"
)

func TestVK_PrettyCardsCreate(t *testing.T) {
	// TODO: write test
}

func TestVK_PrettyCardsDelete(t *testing.T) {
	// TODO: write test
}

func TestVK_PrettyCardsEdit(t *testing.T) {
	// TODO: write test
}

func TestVK_PrettyCardsGet(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	_, err := vkUser.PrettyCardsGet(api.Params{
		"owner_id": -vkGroupID,
	})
	noError(t, err)
}

func TestVK_PrettyCardsGetByID(t *testing.T) {
	// TODO: write test
}

func TestVK_PrettyCardsGetUploadURL(t *testing.T) {
	// TODO: write test
}
