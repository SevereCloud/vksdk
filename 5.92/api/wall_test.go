package api

import (
	"reflect"
	"testing"

	"github.com/SevereCloud/vksdk/5.92/object"
)

func TestVK_WallSearch(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

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
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vkService.WallSearch(tt.argParams)
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
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

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
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vkService.WallSearchExtended(tt.argParams)
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
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

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
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vkService.WallGetReposts(tt.argParams)
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
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

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
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vkService.WallGetCommentsExtended(tt.argParams)
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
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

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
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vkService.WallGetComments(tt.argParams)
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
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WallGetResponse
		wantVkErr Error
	}{
		{
			name: "WallGet",
			argParams: map[string]string{
				"owner_id": "-86529522",
				"count":    "1",
			},
		},
		{
			name:      "WallGet empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vkService.WallGet(tt.argParams)
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
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WallGetExtendedResponse
		wantVkErr Error
	}{
		{
			name: "WallGetExtended",
			argParams: map[string]string{
				"owner_id": "-86529522",
				"count":    "1",
			},
		},
		{
			name:      "WallGetExtended empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vkService.WallGetExtended(tt.argParams)
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
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WallGetByIDResponse
		wantVkErr Error
	}{
		{
			name: "WallGetByID",
			argParams: map[string]string{
				"posts": "-86529522_204382",
			},
		},
		{
			name:      "WallGetByID empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vkService.WallGetByID(tt.argParams)
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
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WallGetByIDExtendedResponse
		wantVkErr Error
	}{
		{
			name: "WallGetByIDExtended",
			argParams: map[string]string{
				"posts": "-86529522_204382",
			},
		},
		{
			name:      "WallGetByIDExtended empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vkService.WallGetByIDExtended(tt.argParams)
			// if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
			// 	t.Errorf("VK.WallGetByIDExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			// }
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallGetByIDExtended() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallCloseComments(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WallCloseCommentsResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallCloseComments empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.WallCloseComments(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.WallCloseComments() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallCloseComments() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallCreateComment(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WallCreateCommentResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallCreateComment empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.WallCreateComment(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.WallCreateComment() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallCreateComment() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallDelete(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WallDeleteResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallDelete empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.WallDelete(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.WallDelete() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallDelete() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallDeleteComment(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WallDeleteCommentResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallDeleteComment empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.WallDeleteComment(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.WallDeleteComment() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallDeleteComment() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallEdit(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WallEditResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallEdit empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.WallEdit(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.WallEdit() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallEdit() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallEditAdsStealth(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WallEditAdsStealthResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallEditAdsStealth empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.WallEditAdsStealth(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.WallEditAdsStealth() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallEditAdsStealth() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallEditComment(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WallEditCommentResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallEditComment empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.WallEditComment(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.WallEditComment() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallEditComment() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallGetComment(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WallGetCommentResponse
		wantVkErr Error
	}{
		{
			name: "WallGetComment",
			argParams: map[string]string{
				"owner_id":   "66559",
				"comment_id": "73674",
			},
		},
		{
			name:      "WallGetComment empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vkUser.WallGetComment(tt.argParams)
			// if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
			// 	t.Errorf("VK.WallGetComment() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			// }
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallGetComment() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallGetCommentExtended(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WallGetCommentExtendedResponse
		wantVkErr Error
	}{
		{
			name: "WallGetCommentExtended",
			argParams: map[string]string{
				"owner_id":   "66559",
				"comment_id": "73674",
			},
		},
		{
			name:      "WallGetCommentExtended empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vkUser.WallGetCommentExtended(tt.argParams)
			// if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
			// 	t.Errorf("VK.WallGetCommentExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			// }
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallGetCommentExtended() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallOpenComments(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WallOpenCommentsResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallOpenComments empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.WallOpenComments(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.WallOpenComments() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallOpenComments() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallPin(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WallPinResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallPin empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.WallPin(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.WallPin() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallPin() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallPost(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse WallPostResponse
		wantVkErr Error
	}{
		{
			name: "WallPost empty",
			argParams: map[string]string{
				"message": "Test post from github.com SevereCloud/vksdk",
			},
		},
		{
			name:      "WallPost empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vkUser.WallPost(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallPost() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallPostAdsStealth(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WallPostAdsStealthResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallPostAdsStealth empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.WallPostAdsStealth(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.WallPostAdsStealth() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallPostAdsStealth() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallReportComment(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WallReportCommentResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallReportComment empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.WallReportComment(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.WallReportComment() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallReportComment() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallReportPost(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WallReportPostResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallReportPost empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.WallReportPost(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.WallReportPost() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallReportPost() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallRepost(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WallRepostResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallRepost empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.WallRepost(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.WallRepost() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallRepost() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallRestore(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WallRestoreResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallRestore empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.WallRestore(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.WallRestore() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallRestore() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallRestoreComment(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WallRestoreCommentResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallRestoreComment empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.WallRestoreComment(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.WallRestoreComment() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallRestoreComment() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_WallUnpin(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse WallUnpinResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "WallUnpin empty",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.WallUnpin(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.WallUnpin() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.WallUnpin() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}
