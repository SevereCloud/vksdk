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
func (vk *VK) LeadFormsCreate(params Params) (response LeadFormsCreateResponse, err error) {
	err = vk.RequestUnmarshal("leadForms.create", params, &response)
	return
}

// LeadFormsDeleteResponse struct
type LeadFormsDeleteResponse struct {
	FormID int `json:"form_id"`
}

// LeadFormsDelete leadForms.delete
//
// https://vk.com/dev/leadForms.delete
func (vk *VK) LeadFormsDelete(params Params) (response LeadFormsDeleteResponse, err error) {
	err = vk.RequestUnmarshal("leadForms.delete", params, &response)
	return
}

// LeadFormsGetResponse struct
type LeadFormsGetResponse object.LeadFormsForm

// LeadFormsGet leadForms.get
//
// https://vk.com/dev/leadForms.get
func (vk *VK) LeadFormsGet(params Params) (response LeadFormsGetResponse, err error) {
	err = vk.RequestUnmarshal("leadForms.get", params, &response)
	return
}

// LeadFormsGetLeadsResponse struct
type LeadFormsGetLeadsResponse struct {
	Leads []object.LeadFormsLead `json:"leads"`
}

// LeadFormsGetLeads leadForms.getLeads
//
// https://vk.com/dev/leadForms.getLeads
func (vk *VK) LeadFormsGetLeads(params Params) (response LeadFormsGetLeadsResponse, err error) {
	err = vk.RequestUnmarshal("leadForms.getLeads", params, &response)
	return
}

// LeadFormsGetUploadURL leadForms.getUploadURL
//
// https://vk.com/dev/leadForms.getUploadURL
func (vk *VK) LeadFormsGetUploadURL(params Params) (response string, err error) {
	err = vk.RequestUnmarshal("leadForms.getUploadURL", params, &response)
	return
}

// LeadFormsListResponse struct
type LeadFormsListResponse []object.LeadFormsForm

// LeadFormsList leadForms.list
//
// https://vk.com/dev/leadForms.list
func (vk *VK) LeadFormsList(params Params) (response LeadFormsListResponse, err error) {
	err = vk.RequestUnmarshal("leadForms.list", params, &response)
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
func (vk *VK) LeadFormsUpdate(params Params) (response LeadFormsUpdateResponse, err error) {
	err = vk.RequestUnmarshal("leadForms.update", params, &response)
	return
}
