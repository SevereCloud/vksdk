package api

import (
	"os"
	"testing"
)

func TestVK_AuthCheckPhone(t *testing.T) {
	serviceToken := os.Getenv("SERVICE_TOKEN")
	clientSecret := os.Getenv("CLIENT_SECRET")
	clientID := os.Getenv("CLIENT_ID")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

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
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vk.AuthCheckPhone(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.AuthCheckPhone() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.AuthCheckPhone() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}
