package api

import (
	"strconv"
	"testing"
)

func TestVK_DocsAdd(t *testing.T) {
	needUserToken(t)

	doc, err := vkUser.DocsAdd(map[string]string{
		"owner_id": "-166562603",
		"doc_id":   "483194018",
	})
	if err != nil {
		t.Errorf("VK.DocsAdd() err = %v", err)
	}

	_, err = vkUser.DocsEdit(map[string]string{
		"owner_id": strconv.Itoa(vkUserID),
		"doc_id":   strconv.Itoa(doc),
		"title":    "test_title",
	})
	if err != nil {
		t.Errorf("VK.DocsEdit() err = %v", err)
	}

	_, err = vkUser.DocsDelete(map[string]string{
		"owner_id": strconv.Itoa(vkUserID),
		"doc_id":   strconv.Itoa(doc),
	})
	if err != nil {
		t.Errorf("VK.DocsDelete() err = %v", err)
	}
}

func TestVK_DocsGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.DocsGet(map[string]string{
		"owner_id": "-166562603",
	})
	if err != nil {
		t.Errorf("VK.DocsGet() err = %v", err)
	}
}

func TestVK_DocsGetByID(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.DocsGetByID(map[string]string{
		"docs": "2314852_165123053",
	})
	if err != nil {
		t.Errorf("VK.DocsGetByID() err = %v", err)
	}
}

func TestVK_DocsGetTypes(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.DocsGetTypes(map[string]string{})
	if err != nil {
		t.Errorf("VK.DocsGetTypes() err = %v", err)
	}
}

func TestVK_DocsGetUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.DocsGetUploadServer(map[string]string{})
	if err != nil {
		t.Errorf("VK.DocsGetUploadServer() err = %v", err)
	}
}

func TestVK_DocsGetMessagesUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.DocsGetMessagesUploadServer(map[string]string{})
	if err != nil {
		t.Errorf("VK.DocsGetMessagesUploadServer() err = %v", err)
	}
}

func TestVK_DocsGetWallUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.DocsGetWallUploadServer(map[string]string{})
	if err != nil {
		t.Errorf("VK.DocsGetWallUploadServer() err = %v", err)
	}
}

// TODO: TestVK_DocsSave

func TestVK_DocsSearch(t *testing.T) {
	needGroupToken(t)

	_, err := vkGroup.DocsSearch(map[string]string{
		"q": "golang",
	})
	if err != nil {
		t.Errorf("VK.DocsSearch() err = %v", err)
	}
}
