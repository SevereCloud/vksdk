package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_DocsAdd(t *testing.T) {
	needUserToken(t)

	doc, err := vkUser.DocsAdd(Params{
		"owner_id": -166562603,
		"doc_id":   483194018,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, doc)

	res, err := vkUser.DocsEdit(Params{
		"owner_id": vkUserID,
		"doc_id":   doc,
		"title":    "test_title",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.DocsDelete(Params{
		"owner_id": vkUserID,
		"doc_id":   doc,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_DocsGet(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.DocsGet(Params{
		"owner_id": -166562603,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Title)
		assert.NotEmpty(t, res.Items[0].OwnerID)
		assert.NotEmpty(t, res.Items[0].Size)
		assert.NotEmpty(t, res.Items[0].Ext)
		assert.NotEmpty(t, res.Items[0].URL)
		assert.NotEmpty(t, res.Items[0].Date)
		assert.NotEmpty(t, res.Items[0].Type)
	}
}

func TestVK_DocsGetByID(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.DocsGetByID(Params{
		"docs": "2314852_165123053",
	})
	assert.NoError(t, err)

	if assert.NotEmpty(t, res) {
		assert.NotEmpty(t, res[0].ID)
		assert.NotEmpty(t, res[0].Title)
		assert.NotEmpty(t, res[0].OwnerID)
		assert.NotEmpty(t, res[0].Size)
		assert.NotEmpty(t, res[0].Ext)
		assert.NotEmpty(t, res[0].URL)
		assert.NotEmpty(t, res[0].Date)
		assert.NotEmpty(t, res[0].Type)

		if assert.NotEmpty(t, res[0].Preview) {
			if assert.NotEmpty(t, res[0].Preview.Photo.Sizes) {
				size := res[0].Preview.Photo.Sizes[0]
				assert.NotEmpty(t, size.Src)
				assert.NotEmpty(t, size.Width)
				assert.NotEmpty(t, size.Height)
				assert.NotEmpty(t, size.Type)
			}

			assert.NotEmpty(t, res[0].Preview.Video.Src)
			assert.NotEmpty(t, res[0].Preview.Video.Width)
			assert.NotEmpty(t, res[0].Preview.Video.Height)
			assert.NotEmpty(t, res[0].Preview.Video.FileSize)
		}
	}
}

func TestVK_DocsGetTypes(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.DocsGetTypes(Params{})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Name)
		assert.NotEmpty(t, res.Items[0].Count)
	}
}

func TestVK_DocsGetUploadServer(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.DocsGetUploadServer(Params{})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.UploadURL)
}

func TestVK_DocsGetMessagesUploadServer(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.DocsGetMessagesUploadServer(Params{})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.UploadURL)
}

func TestVK_DocsGetWallUploadServer(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.DocsGetWallUploadServer(Params{})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.UploadURL)
}

// TODO: TestVK_DocsSave

func TestVK_DocsSearch(t *testing.T) {
	needGroupToken(t)

	res, err := vkGroup.DocsSearch(Params{
		"q": "golang",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].OwnerID)
		assert.NotEmpty(t, res.Items[0].Title)
		assert.NotEmpty(t, res.Items[0].Size)
		assert.NotEmpty(t, res.Items[0].Ext)
		assert.NotEmpty(t, res.Items[0].URL)
		assert.NotEmpty(t, res.Items[0].Date)
		assert.NotEmpty(t, res.Items[0].Type)
	}
}
