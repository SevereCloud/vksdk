package api

import (
	"os"
	"testing"
)

func TestVK_WallSearch(t *testing.T) {
	serviceToken := os.Getenv("SERVICE_TOKEN")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WallSearchResponse
		wantVkErr Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallSearch empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vk.WallSearch(tt.argParams)
			// if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
			// 	t.Errorf("VK.WallSearch() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			// }
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallSearch() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallSearchExtended(t *testing.T) {
	serviceToken := os.Getenv("SERVICE_TOKEN")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WallSearchExtendedResponse
		wantVkErr Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallSearchExtended empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vk.WallSearchExtended(tt.argParams)
			// if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
			// 	t.Errorf("VK.WallSearchExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			// }
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallSearchExtended() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallGetReposts(t *testing.T) {
	serviceToken := os.Getenv("SERVICE_TOKEN")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WallGetRepostsResponse
		wantVkErr Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallGetReposts empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vk.WallGetReposts(tt.argParams)
			// if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
			// 	t.Errorf("VK.WallGetReposts() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			// }
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallGetReposts() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallGetCommentsExtended(t *testing.T) {
	serviceToken := os.Getenv("SERVICE_TOKEN")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WallGetCommentsExtendedResponse
		wantVkErr Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallGetCommentsExtended empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vk.WallGetCommentsExtended(tt.argParams)
			// if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
			// 	t.Errorf("VK.WallGetCommentsExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			// }
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallGetCommentsExtended() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallGetComments(t *testing.T) {
	serviceToken := os.Getenv("SERVICE_TOKEN")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WallGetCommentsResponse
		wantVkErr Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallGetComments empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vk.WallGetComments(tt.argParams)
			// if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
			// 	t.Errorf("VK.WallGetComments() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			// }
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallGetComments() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallGet(t *testing.T) {
	serviceToken := os.Getenv("SERVICE_TOKEN")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WallGetResponse
		wantVkErr Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallGet empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vk.WallGet(tt.argParams)
			// if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
			// 	t.Errorf("VK.WallGet() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			// }
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallGet() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallGetExtended(t *testing.T) {
	serviceToken := os.Getenv("SERVICE_TOKEN")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WallGetExtendedResponse
		wantVkErr Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallGetExtended empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vk.WallGetExtended(tt.argParams)
			// if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
			// 	t.Errorf("VK.WallGetExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			// }
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallGetExtended() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallGetByID(t *testing.T) {
	serviceToken := os.Getenv("SERVICE_TOKEN")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WallGetByIDResponse
		wantVkErr Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallGetByID empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vk.WallGetByID(tt.argParams)
			// if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
			// 	t.Errorf("VK.WallGetByID() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			// }
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallGetByID() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallGetByIDExtended(t *testing.T) {
	serviceToken := os.Getenv("SERVICE_TOKEN")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WallGetByIDExtendedResponse
		wantVkErr Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallGetByIDExtended empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vk.WallGetByIDExtended(tt.argParams)
			// if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
			// 	t.Errorf("VK.WallGetByIDExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			// }
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallGetByIDExtended() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}
