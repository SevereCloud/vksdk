package api

import (
	"reflect"
	"testing"
)

func TestVK_WidgetsGetComments(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WidgetsGetCommentsResponse
		wantErr      bool
	}{
		// TODO: add test
		{
			name:    "widgets.getComments error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkService.WidgetsGetComments(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.WidgetsGetComments() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.WidgetsGetComments() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_WidgetsGetPages(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WidgetsGetPagesResponse
		wantErr bool
	}{
		// TODO: add test
		{
			name:    "widgets.getPages error",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vkService.WidgetsGetPages(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.WidgetsGetPages() err = %v, want %v", err, tt.wantErr)
			}
			// if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
			// 	t.Errorf("VK.WidgetsGetPages() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			// }
		})
	}
}
