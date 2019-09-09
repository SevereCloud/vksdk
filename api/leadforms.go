package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// LeadFormsCreateResponse struct
type LeadFormsCreateResponse struct {
	FormID int    `json:"form_id"`
	URL    string `json:"url"`
}

// LeadFormsCreate leadForms.create
//
// https://vk.com/dev/leadForms.create
func (vk *VK) LeadFormsCreate(params map[string]string) (response LeadFormsCreateResponse, vkErr Error) {
	vk.RequestUnmarshal("leadForms.create", params, &response, &vkErr)
	return
}

// LeadFormsDeleteResponse struct
type LeadFormsDeleteResponse struct {
	FormID int `json:"form_id"`
}

// LeadFormsDelete leadForms.delete
//
// https://vk.com/dev/leadForms.delete
func (vk *VK) LeadFormsDelete(params map[string]string) (response LeadFormsDeleteResponse, vkErr Error) {
	vk.RequestUnmarshal("leadForms.delete", params, &response, &vkErr)
	return
}

// LeadFormsGetResponse struct
type LeadFormsGetResponse object.LeadFormsForm

// LeadFormsGet leadForms.get
//
// https://vk.com/dev/leadForms.get
func (vk *VK) LeadFormsGet(params map[string]string) (response LeadFormsGetResponse, vkErr Error) {
	vk.RequestUnmarshal("leadForms.get", params, &response, &vkErr)
	return
}

// LeadFormsGetLeadsResponse struct
type LeadFormsGetLeadsResponse struct {
	Leads []object.LeadFormsLead `json:"leads"`
}

// LeadFormsGetLeads leadForms.getLeads
//
// https://vk.com/dev/leadForms.getLeads
func (vk *VK) LeadFormsGetLeads(params map[string]string) (response LeadFormsGetLeadsResponse, vkErr Error) {
	vk.RequestUnmarshal("leadForms.getLeads", params, &response, &vkErr)
	return
}

// LeadFormsGetUploadURL leadForms.getUploadURL
//
// https://vk.com/dev/leadForms.getUploadURL
func (vk *VK) LeadFormsGetUploadURL(params map[string]string) (response string, vkErr Error) {
	vk.RequestUnmarshal("leadForms.getUploadURL", params, &response, &vkErr)
	return
}

// LeadFormsListResponse struct
type LeadFormsListResponse []object.LeadFormsForm

// LeadFormsList leadForms.list
//
// https://vk.com/dev/leadForms.list
func (vk *VK) LeadFormsList(params map[string]string) (response LeadFormsListResponse, vkErr Error) {
	vk.RequestUnmarshal("leadForms.list", params, &response, &vkErr)
	return
}

// LeadFormsUpdateResponse struct
type LeadFormsUpdateResponse struct {
	FormID int    `json:"form_id"`
	URL    string `json:"url"`
}

// LeadFormsUpdate leadForms.update
//
// https://vk.com/dev/leadForms.update
func (vk *VK) LeadFormsUpdate(params map[string]string) (response LeadFormsUpdateResponse, vkErr Error) {
	vk.RequestUnmarshal("leadForms.update", params, &response, &vkErr)
	return
}
