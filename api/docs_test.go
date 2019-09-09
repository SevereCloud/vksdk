package api

import (
	"strconv"
	"testing"
)

func TestVK_DocsAdd(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	doc, gotVkErr := vkUser.DocsAdd(map[string]string{
		"owner_id": "-166562603",
		"doc_id":   "483194018",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DocsAdd() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.DocsEdit(map[string]string{
		"owner_id": strconv.Itoa(vkUserID),
		"doc_id":   strconv.Itoa(doc),
		"title":    "test_title",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DocsEdit() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.DocsDelete(map[string]string{
		"owner_id": strconv.Itoa(vkUserID),
		"doc_id":   strconv.Itoa(doc),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DocsDelete() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DocsGet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.DocsGet(map[string]string{
		"owner_id": "-166562603",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DocsGet() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DocsGetByID(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.DocsGetByID(map[string]string{
		"docs": "2314852_165123053",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DocsGetByID() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DocsGetTypes(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.DocsGetTypes(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DocsGetTypes() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DocsGetUploadServer(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.DocsGetUploadServer(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DocsGetUploadServer() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DocsGetMessagesUploadServer(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.DocsGetMessagesUploadServer(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DocsGetMessagesUploadServer() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DocsGetWallUploadServer(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.DocsGetWallUploadServer(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DocsGetWallUploadServer() gotVkErr = %v", gotVkErr)
	}
}

// TODO: TestVK_DocsSave

func TestVK_DocsSearch(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	_, gotVkErr := vkGroup.DocsSearch(map[string]string{
		"q": "golang",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DocsSearch() gotVkErr = %v", gotVkErr)
	}
}
