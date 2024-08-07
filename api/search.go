package api // import "github.com/SevereCloud/vksdk/v3/api"

import "github.com/SevereCloud/vksdk/v3/object"

// SearchGetHintsResponse struct.
type SearchGetHintsResponse struct {
	Count int                 `json:"count"`
	Items []object.SearchHint `json:"items"`
}

// SearchGetHints allows the programmer to do a quick search for any substring.
//
// https://dev.vk.com/method/search.getHints
func (vk *VK) SearchGetHints(params Params) (response SearchGetHintsResponse, err error) {
	err = vk.RequestUnmarshal("search.getHints", &response, params)
	return
}
