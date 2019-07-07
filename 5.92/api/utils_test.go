package api

import (
	"reflect"
	"testing"

	"github.com/SevereCloud/vksdk/5.92/object"
)

func TestVK_UtilsCheckLink(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse UtilsCheckLinkResponse
		wantVkErr    Error
	}{
		{
			name:      "utils.checkLink google.ru",
			argParams: map[string]string{"url": "http://google.ru"},
			wantResponse: UtilsCheckLinkResponse{
				Status: "not_banned",
				Link:   "http://google.ru",
			},
		},
		{
			name:      "utils.checkLink error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkService.UtilsCheckLink(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.UtilsCheckLink() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.UtilsCheckLink() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_UtilsDeleteFromLastShortened(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		// TODO: Add test cases.
		{
			name:      "utils.deleteFromLastShortened error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vkUser.UtilsDeleteFromLastShortened(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.UtilsDeleteFromLastShortened() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_UtilsGetLastShortenedLinks(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse UtilsGetLastShortenedLinksResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "utils.getLastShortenedLinks error",
			argParams: map[string]string{"offset": "-1"},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.UtilsGetLastShortenedLinks(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.UtilsGetLastShortenedLinks() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.UtilsGetLastShortenedLinks() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_UtilsGetLinkStats(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse UtilsGetLinkStatsResponse
		wantVkErr    Error
	}{
		{
			name: "utils.getLinkStats 8TDuIz",
			argParams: map[string]string{
				"key":             "8TDuIz",
				"intervals_count": "1",
			},
			wantResponse: UtilsGetLinkStatsResponse{
				Key: "8TDuIz",
			},
		},
		{
			name:      "utils.getLinkStats error",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkService.UtilsGetLinkStats(tt.argParams)
			if gotResponse.Key != tt.wantResponse.Key {
				t.Errorf("VK.UtilsGetLinkStats() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.UtilsGetLinkStats() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_UtilsGetLinkStatsExtended(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse UtilsGetLinkStatsExtendedResponse
		wantVkErr    Error
	}{
		{
			name: "utils.getLinkStatsExtended 8TDuIz",
			argParams: map[string]string{
				"key":             "8TDuIz",
				"intervals_count": "1",
			},
			wantResponse: UtilsGetLinkStatsExtendedResponse{
				Key: "8TDuIz",
			},
		},
		{
			name:      "utils.getLinkStatsExtended error",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkService.UtilsGetLinkStatsExtended(tt.argParams)
			if gotResponse.Key != tt.wantResponse.Key {
				t.Errorf("VK.UtilsGetLinkStatsExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.UtilsGetLinkStatsExtended() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_UtilsGetServerTime(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	gotResponse, vkErr := vkService.UtilsGetServerTime(map[string]string{})
	if vkErr.Code != 0 {
		t.Errorf("VK.UtilsGetServerTime() vkErr.Code = %v, want 0", vkErr)
	}
	if gotResponse < 1 {
		t.Errorf("VK.UtilsGetServerTime() gotResponse = %v, want >0", gotResponse)
	}
}
func TestVK_UtilsGetServerTimeError(t *testing.T) {
	vk := Init("")
	_, vkErr := vk.UtilsGetServerTime(map[string]string{})
	if vkErr.Code != 5 {
		t.Errorf("VK.UtilsGetServerTime() error bad %d %s", vkErr.Code, vkErr.Message)
	}
}

func TestVK_UtilsGetShortLink(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse UtilsGetShortLinkResponse
		wantVkErr    Error
	}{
		{
			name:      "utils.getShortLink google.ru",
			argParams: map[string]string{"url": "http://google.ru"},
			wantResponse: UtilsGetShortLinkResponse{
				ShortURL: "https://vk.cc/8TDuIz",
				URL:      "http://google.ru",
				Key:      "8TDuIz",
			},
		},
		{
			name:      "utils.getShortLink error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkService.UtilsGetShortLink(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.UtilsGetShortLink() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.UtilsGetShortLink() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_UtilsResolveScreenName(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse UtilsResolveScreenNameResponse
		wantVkErr    Error
	}{
		{
			name:      "utils.resolveScreenName durov",
			argParams: map[string]string{"screen_name": "durov"},
			wantResponse: UtilsResolveScreenNameResponse{
				Type:     "user",
				ObjectID: 1,
			},
		},
		{
			name:         "utils.resolveScreenName not found",
			argParams:    map[string]string{"screen_name": "z"},
			wantResponse: UtilsResolveScreenNameResponse{},
		},
		{
			name:      "utils.resolveScreenName error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkService.UtilsResolveScreenName(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.UtilsResolveScreenName() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.UtilsResolveScreenName() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}
