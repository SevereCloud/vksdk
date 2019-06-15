package api

import (
	"os"
	"testing"

	"github.com/SevereCloud/vksdk/5.92/object"
)

func TestVK_DocsSearch(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name: "docs.search",
			argParams: map[string]string{
				"q": "golang",
			},
		},
		{
			name:      "docs.search error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.DocsSearch(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.DocsSearch() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if gotVkErr.Code == 0 && (gotResponse.Count == 0) {
				t.Errorf("VK.DocsSearch() gotResponse = %v", gotResponse)
			}
		})
	}
}

func TestVK_DocsGetWallUploadServer(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name: "groups.GetWallUploadServer",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:      "docs.GetWallUploadServer error",
			wantVkErr: Error{Code: object.ErrorAccess},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.DocsGetWallUploadServer(tt.argParams)

			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.DocsGetWallUploadServer() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if gotVkErr.Code == 0 && gotResponse.UploadURL == "" {
				t.Errorf("VK.DocsGetWallUploadServer() gotResponse = %v", gotResponse)
			}
		})
	}
}

func TestVK_DocsGetMessagesUploadServer(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		// FIXME: groups.GetMessagesUploadServer need peer_id
		// {
		// 	name: "groups.GetMessagesUploadServer",
		// 	argParams: map[string]string{
		// 		"group_id": os.Getenv("GROUP_ID"),
		// 		"peer_id":  "",
		// 	},
		// },
		{
			name:      "docs.GetMessagesUploadServer error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.DocsGetMessagesUploadServer(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.DocsGetMessagesUploadServer() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if gotVkErr.Code == 0 && gotResponse.UploadURL == "" {
				t.Errorf("VK.DocsGetMessagesUploadServer() gotResponse = %v", gotResponse)
			}
		})
	}
}
