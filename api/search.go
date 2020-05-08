package api // import "github.com/SevereCloud/vksdk/api"

import "github.com/SevereCloud/vksdk/object"

// SearchGetHintsResponse struct.
type SearchGetHintsResponse struct {
	Count int                 `json:"count"`
	Items []object.SearchHint `json:"items"`
}

// SearchGetHints allows the programmer to do a quick search for any substring.
//
// https://vk.com/dev/search.getHints
func (vk *VK) SearchGetHints(params Params) (response SearchGetHintsResponse, err error) {
	err = vk.RequestUnmarshal("search.getHints", params, &response)
	return
}
