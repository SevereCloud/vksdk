package api

import (
	"os"
	"testing"

	"github.com/SevereCloud/vksdk/5.92/object"
)

func TestVK_AuthCheckPhone(t *testing.T) {
	clientSecret := os.Getenv("CLIENT_SECRET")
	clientID := os.Getenv("CLIENT_ID")
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse AuthCheckResponse
		wantVkErr    Error
	}{
		{
			name: "groups.getById error",
			argParams: map[string]string{
				"phone":         "+79523071234",
				"client_id":     clientID,
				"client_secret": clientSecret,
			},
			wantResponse: 1,
		},
		{
			name:      "auth.checkPhone error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkService.AuthCheckPhone(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.AuthCheckPhone() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.AuthCheckPhone() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}
