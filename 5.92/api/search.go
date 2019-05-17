package api // import "github.com/SevereCloud/vksdk/5.92/api"

import "github.com/SevereCloud/vksdk/5.92/object"

// SearchGetHintsResponse struct
type SearchGetHintsResponse struct {
	Count int                 `json:"count"`
	Items []object.SearchHint `json:"items"`
}

// SearchGetHints allows the programmer to do a quick search for any substring.
//
// https://vk.com/dev/search.getHints
func (vk *VK) SearchGetHints(params map[string]string) (response SearchGetHintsResponse, vkErr Error) {
	vk.RequestUnmarshal("search.getHints", params, &response, &vkErr)
	return
}
