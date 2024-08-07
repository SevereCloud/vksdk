package api // import "github.com/SevereCloud/vksdk/v3/api"

import "github.com/SevereCloud/vksdk/v3/object"

// PrettyCardsCreateResponse struct.
type PrettyCardsCreateResponse struct {
	OwnerID int    `json:"owner_id"` // Owner ID of created pretty card
	CardID  string `json:"card_id"`  // Card ID of created pretty card
}

// PrettyCardsCreate method.
//
// https://dev.vk.com/method/prettyCards.create
func (vk *VK) PrettyCardsCreate(params Params) (response PrettyCardsCreateResponse, err error) {
	err = vk.RequestUnmarshal("prettyCards.create", &response, params)
	return
}

// PrettyCardsDeleteResponse struct.
type PrettyCardsDeleteResponse struct {
	OwnerID int    `json:"owner_id"` // Owner ID of created pretty card
	CardID  string `json:"card_id"`  // Card ID of created pretty card
	Error   string `json:"error"`    // Error reason if error happened
}

// PrettyCardsDelete method.
//
// https://dev.vk.com/method/prettyCards.delete
func (vk *VK) PrettyCardsDelete(params Params) (response PrettyCardsDeleteResponse, err error) {
	err = vk.RequestUnmarshal("prettyCards.delete", &response, params)
	return
}

// PrettyCardsEditResponse struct.
type PrettyCardsEditResponse struct {
	OwnerID int    `json:"owner_id"` // Owner ID of created pretty card
	CardID  string `json:"card_id"`  // Card ID of created pretty card
}

// PrettyCardsEdit method.
//
// https://dev.vk.com/method/prettyCards.edit
func (vk *VK) PrettyCardsEdit(params Params) (response PrettyCardsEditResponse, err error) {
	err = vk.RequestUnmarshal("prettyCards.edit", &response, params)
	return
}

// PrettyCardsGetResponse struct.
type PrettyCardsGetResponse struct {
	Count int                            `json:"count"` // Total number
	Items []object.PrettyCardsPrettyCard `json:"items"`
}

// PrettyCardsGet method.
//
// https://dev.vk.com/method/prettyCards.get
func (vk *VK) PrettyCardsGet(params Params) (response PrettyCardsGetResponse, err error) {
	err = vk.RequestUnmarshal("prettyCards.get", &response, params)
	return
}

// PrettyCardsGetByIDResponse struct.
type PrettyCardsGetByIDResponse []object.PrettyCardsPrettyCard

// PrettyCardsGetByID method.
//
// https://dev.vk.com/method/prettyCards.getById
func (vk *VK) PrettyCardsGetByID(params Params) (response PrettyCardsGetByIDResponse, err error) {
	err = vk.RequestUnmarshal("prettyCards.getById", &response, params)
	return
}

// PrettyCardsGetUploadURL method.
//
// https://dev.vk.com/method/prettyCards.getUploadURL
func (vk *VK) PrettyCardsGetUploadURL(params Params) (response string, err error) {
	err = vk.RequestUnmarshal("prettyCards.getUploadURL", &response, params)
	return
}
