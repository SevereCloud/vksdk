package api

import (
	"os"
	"reflect"
	"testing"
)

func TestVK_WidgetsGetComments(t *testing.T) {
	serviceToken := os.Getenv("SERVICE_TOKEN")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WidgetsGetCommentsResponse
		wantVkErr    Error
	}{
		{
			name:      "widgets.getComments error",
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vk.WidgetsGetComments(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WidgetsGetComments() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.WidgetsGetComments() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_WidgetsGetPages(t *testing.T) {
	serviceToken := os.Getenv("SERVICE_TOKEN")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WidgetsGetPagesResponse
		wantVkErr    Error
	}{
		{
			name:      "widgets.getPages error",
			wantVkErr: Error{Code: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vk.WidgetsGetPages(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WidgetsGetPages() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			// if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
			// 	t.Errorf("VK.WidgetsGetPages() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			// }
		})
	}
}
