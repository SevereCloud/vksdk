package api

import (
	"reflect"
	"testing"
)

func TestVK_StoriesGet(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGet error",
		// 	argParams: map[string]string{},
		// 	wantErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.StoriesGet(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.StoriesGet() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGet() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetExtended(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetExtended error",
		// 	argParams: map[string]string{},
		// 	wantErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.StoriesGetExtended(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.StoriesGetExtended() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetBanned(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetBanned error",
		// 	argParams: map[string]string{},
		// 	wantErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.StoriesGetBanned(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.StoriesGetBanned() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetBanned() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetBannedExtended(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetBannedExtended error",
		// 	argParams: map[string]string{},
		// 	wantErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.StoriesGetBannedExtended(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.StoriesGetBannedExtended() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetBannedExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetByID(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetByID error",
		// 	argParams: map[string]string{},
		// 	wantErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.StoriesGetByID(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.StoriesGetByID() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetByID() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetByIDExtended(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetByIDExtended error",
		// 	argParams: map[string]string{},
		// 	wantErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.StoriesGetByIDExtended(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.StoriesGetByIDExtended() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetByIDExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetPhotoUploadServer(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetPhotoUploadServer error",
		// 	argParams: map[string]string{},
		// 	wantErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.StoriesGetPhotoUploadServer(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.StoriesGetPhotoUploadServer() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetPhotoUploadServer() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetReplies(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetReplies error",
		// 	argParams: map[string]string{},
		// 	wantErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.StoriesGetReplies(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.StoriesGetReplies() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetReplies() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetRepliesExtended(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetRepliesExtended error",
		// 	argParams: map[string]string{},
		// 	wantErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.StoriesGetRepliesExtended(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.StoriesGetRepliesExtended() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetRepliesExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetStats(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetStats error",
		// 	argParams: map[string]string{},
		// 	wantErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.StoriesGetStats(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.StoriesGetStats() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetStats() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetVideoUploadServer(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetVideoUploadServer error",
		// 	argParams: map[string]string{},
		// 	wantErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.StoriesGetVideoUploadServer(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.StoriesGetVideoUploadServer() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetVideoUploadServer() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetViewers(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetViewers error",
		// 	argParams: map[string]string{},
		// 	wantErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.StoriesGetViewers(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.StoriesGetViewers() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetViewers() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetViewersExtended(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetViewersExtended error",
		// 	argParams: map[string]string{},
		// 	wantErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.StoriesGetViewersExtended(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.StoriesGetViewersExtended() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetViewersExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesHideAllReplies(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name      string
		argParams map[string]string
		wantErr   Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesHideAllReplies error",
		// 	argParams: map[string]string{},
		// 	wantErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := vkGroup.StoriesHideAllReplies(tt.argParams); !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("VK.StoriesHideAllReplies() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_StoriesHideReply(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name      string
		argParams map[string]string
		wantErr   Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesHideReply error",
		// 	argParams: map[string]string{},
		// 	wantErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := vkGroup.StoriesHideReply(tt.argParams); !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("VK.StoriesHideReply() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
