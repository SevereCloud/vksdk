package api // import "github.com/severecloud/vksdk/5.92/api"

import (
	"encoding/json"

	"github.com/severecloud/vksdk/5.92/object"
)

// LeadFormsCreateResponse struct
type LeadFormsCreateResponse struct {
	FormID int    `json:"form_id"`
	URL    string `json:"url"`
}

// LeadFormsCreate leadForms.create
// https://vk.com/dev/leadForms.create
func (vk VK) LeadFormsCreate(params map[string]string) (response LeadFormsCreateResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("leadForms.create", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// LeadFormsDeleteResponse struct
type LeadFormsDeleteResponse struct {
	FormID int `json:"form_id"`
}

// LeadFormsDelete leadForms.delete
// https://vk.com/dev/leadForms.delete
func (vk VK) LeadFormsDelete(params map[string]string) (response LeadFormsDeleteResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("leadForms.delete", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// LeadFormsGetResponse struct
type LeadFormsGetResponse object.LeadFormsForm

// LeadFormsGet leadForms.get
// https://vk.com/dev/leadForms.get
func (vk VK) LeadFormsGet(params map[string]string) (response LeadFormsGetResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("leadForms.get", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// LeadFormsGetLeadsResponse struct
type LeadFormsGetLeadsResponse struct {
	Leads []object.LeadFormsLead `json:"leads"`
}

// LeadFormsGetLeads leadForms.getLeads
// https://vk.com/dev/leadForms.getLeads
func (vk VK) LeadFormsGetLeads(params map[string]string) (response LeadFormsGetLeadsResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("leadForms.getLeads", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// LeadFormsGetUploadURLResponse struct
type LeadFormsGetUploadURLResponse string

// LeadFormsGetUploadURL leadForms.getUploadURL
// https://vk.com/dev/leadForms.getUploadURL
func (vk VK) LeadFormsGetUploadURL() (response LeadFormsGetUploadURLResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("leadForms.getUploadURL", map[string]string{})
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// LeadFormsListResponse struct
type LeadFormsListResponse []object.LeadFormsForm

// LeadFormsList leadForms.list
// https://vk.com/dev/leadForms.list
func (vk VK) LeadFormsList(params map[string]string) (response LeadFormsListResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("leadForms.list", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// LeadFormsUpdateResponse struct
type LeadFormsUpdateResponse struct {
	FormID int    `json:"form_id"`
	URL    string `json:"url"`
}

// LeadFormsUpdate leadForms.update
// https://vk.com/dev/leadForms.update
func (vk VK) LeadFormsUpdate(params map[string]string) (response LeadFormsUpdateResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("leadForms.update", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}
