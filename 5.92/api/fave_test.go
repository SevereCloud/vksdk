package api

import (
	"reflect"
	"testing"
	"time"
)

func TestVK_FaveAddArticle(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveAddArticleResponse
		wantVkErr    Error
	}{
		{
			name: "add article",
			argParams: map[string]string{
				"url": "https://vk.com/@vkappsdev-vk-apps-kak-popast-v-katalog",
			},
			wantResponse: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveAddArticle(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveAddArticle() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveAddArticle() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveAddLink(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveAddLinkResponse
		wantVkErr    Error
	}{
		{
			name: "add link",
			argParams: map[string]string{
				"link": "https://ya.ru",
			},
			wantResponse: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveAddLink(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveAddLink() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveAddLink() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveAddPage(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveAddPageResponse
		wantVkErr    Error
	}{
		{
			name: "add durov",
			argParams: map[string]string{
				"user_id": "1",
			},
			wantResponse: 1,
		},
		{
			name: "add api",
			argParams: map[string]string{
				"group_id": "1",
			},
			wantResponse: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveAddPage(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveAddPage() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveAddPage() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveAddPost(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveAddPostResponse
		wantVkErr    Error
	}{
		{
			name: "add wall-28551727_5713",
			argParams: map[string]string{
				"owner_id": "-28551727",
				"id":       "5713",
			},
			wantResponse: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveAddPost(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveAddPost() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveAddPost() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveAddProduct(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveAddProductResponse
		wantVkErr    Error
	}{
		{
			name: "add product",
			argParams: map[string]string{
				"owner_id": "-169097025",
				"id":       "3398864",
			},
			wantResponse: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveAddProduct(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveAddProduct() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveAddProduct() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveAddTag(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse FaveAddTagResponse
		wantVkErr Error
	}{
		{
			name: "add tag",
			argParams: map[string]string{
				"name": "Важное",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			_, gotVkErr := vkUser.FaveAddTag(tt.argParams)
			// if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
			// 	t.Errorf("VK.FaveAddTag() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			// }
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveAddTag() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveAddVideo(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveAddVideoResponse
		wantVkErr    Error
	}{
		{
			name: "add video-84585194_456239018",
			argParams: map[string]string{
				"owner_id": "-84585194",
				"id":       "456239018",
			},
			wantResponse: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveAddVideo(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveAddVideo() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveAddVideo() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveEditTag(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveEditTagResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveEditTag(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveEditTag() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveEditTag() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveGet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveGetResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveGet(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.FaveGet() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveGet() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveGetExtended(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveGetExtendedResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveGetExtended(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.FaveGetExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveGetExtended() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveGetPages(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveGetPagesResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveGetPages(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.FaveGetPages() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveGetPages() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveGetTags(t *testing.T) {
	tests := []struct {
		name string
		// wantResponse FaveGetTagsResponse
		wantVkErr Error
	}{
		{
			name: "empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			_, gotVkErr := vkUser.FaveGetTags()
			// if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
			// 	t.Errorf("VK.FaveGetTags() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			// }
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveGetTags() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveMarkSeen(t *testing.T) {
	tests := []struct {
		name         string
		wantResponse FaveMarkSeenResponse
		wantVkErr    Error
	}{
		{
			name:         "empty",
			wantResponse: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveMarkSeen()
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveMarkSeen() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveMarkSeen() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveRemoveArticle(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveRemoveArticleResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveRemoveArticle(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveRemoveArticle() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveRemoveArticle() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveRemoveLink(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveRemoveLinkResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveRemoveLink(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveRemoveLink() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveRemoveLink() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveRemovePage(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveRemovePageResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveRemovePage(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveRemovePage() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveRemovePage() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveRemovePost(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveRemovePostResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveRemovePost(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveRemovePost() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveRemovePost() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveRemoveProduct(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveRemoveProductResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveRemoveProduct(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveRemoveProduct() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveRemoveProduct() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveRemoveTag(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveRemoveTagResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveRemoveTag(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveRemoveTag() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveRemoveTag() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveRemoveVideo(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveRemoveVideoResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveRemoveVideo(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveRemoveVideo() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveRemoveVideo() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveReorderTags(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveReorderTagsResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveReorderTags(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveReorderTags() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveReorderTags() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveSetPageTags(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveSetPageTagsResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveSetPageTags(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveSetPageTags() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveSetPageTags() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveSetTags(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveSetTagsResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveSetTags(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveSetTags() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveSetTags() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_FaveTrackPageInteraction(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse FaveTrackPageInteractionResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(500 * time.Millisecond)
			gotResponse, gotVkErr := vkUser.FaveTrackPageInteraction(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.FaveTrackPageInteraction() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.FaveTrackPageInteraction() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}
