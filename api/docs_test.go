package api

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_DocsAdd(t *testing.T) {
	needUserToken(t)

	doc, err := vkUser.DocsAdd(map[string]string{
		"owner_id": "-166562603",
		"doc_id":   "483194018",
	})
	assert.NoError(t, err)

	_, err = vkUser.DocsEdit(map[string]string{
		"owner_id": strconv.Itoa(vkUserID),
		"doc_id":   strconv.Itoa(doc),
		"title":    "test_title",
	})
	assert.NoError(t, err)

	_, err = vkUser.DocsDelete(map[string]string{
		"owner_id": strconv.Itoa(vkUserID),
		"doc_id":   strconv.Itoa(doc),
	})
	assert.NoError(t, err)
}

func TestVK_DocsGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.DocsGet(map[string]string{
		"owner_id": "-166562603",
	})
	assert.NoError(t, err)
}

func TestVK_DocsGetByID(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.DocsGetByID(map[string]string{
		"docs": "2314852_165123053",
	})
	assert.NoError(t, err)
}

func TestVK_DocsGetTypes(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.DocsGetTypes(map[string]string{})
	assert.NoError(t, err)
}

func TestVK_DocsGetUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.DocsGetUploadServer(map[string]string{})
	assert.NoError(t, err)
}

func TestVK_DocsGetMessagesUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.DocsGetMessagesUploadServer(map[string]string{})
	assert.NoError(t, err)
}

func TestVK_DocsGetWallUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.DocsGetWallUploadServer(map[string]string{})
	assert.NoError(t, err)
}

// TODO: TestVK_DocsSave

func TestVK_DocsSearch(t *testing.T) {
	needGroupToken(t)

	_, err := vkGroup.DocsSearch(map[string]string{
		"q": "golang",
	})
	assert.NoError(t, err)
}
