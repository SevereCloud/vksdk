package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// LeadFormsCreateResponse struct
type LeadFormsCreateResponse struct {
	FormID int    `json:"form_id"`
	URL    string `json:"url"`
}

// LeadFormsCreate leadForms.create
// https://vk.com/dev/leadForms.create
func (vk *VK) LeadFormsCreate(params map[string]string) (response LeadFormsCreateResponse, vkErr Error) {
	vk.requestU("leadForms.create", params, &response, &vkErr)
	return
}

// LeadFormsDeleteResponse struct
type LeadFormsDeleteResponse struct {
	FormID int `json:"form_id"`
}

// LeadFormsDelete leadForms.delete
// https://vk.com/dev/leadForms.delete
func (vk *VK) LeadFormsDelete(params map[string]string) (response LeadFormsDeleteResponse, vkErr Error) {
	vk.requestU("leadForms.delete", params, &response, &vkErr)
	return
}

// LeadFormsGetResponse struct
type LeadFormsGetResponse object.LeadFormsForm

// LeadFormsGet leadForms.get
// https://vk.com/dev/leadForms.get
func (vk *VK) LeadFormsGet(params map[string]string) (response LeadFormsGetResponse, vkErr Error) {
	vk.requestU("leadForms.get", params, &response, &vkErr)
	return
}

// LeadFormsGetLeadsResponse struct
type LeadFormsGetLeadsResponse struct {
	Leads []object.LeadFormsLead `json:"leads"`
}

// LeadFormsGetLeads leadForms.getLeads
// https://vk.com/dev/leadForms.getLeads
func (vk *VK) LeadFormsGetLeads(params map[string]string) (response LeadFormsGetLeadsResponse, vkErr Error) {
	vk.requestU("leadForms.getLeads", params, &response, &vkErr)
	return
}

// LeadFormsGetUploadURLResponse struct
type LeadFormsGetUploadURLResponse string

// LeadFormsGetUploadURL leadForms.getUploadURL
// https://vk.com/dev/leadForms.getUploadURL
func (vk *VK) LeadFormsGetUploadURL() (response LeadFormsGetUploadURLResponse, vkErr Error) {
	vk.requestU("leadForms.getUploadURL", map[string]string{}, &response, &vkErr)
	return
}

// LeadFormsListResponse struct
type LeadFormsListResponse []object.LeadFormsForm

// LeadFormsList leadForms.list
// https://vk.com/dev/leadForms.list
func (vk *VK) LeadFormsList(params map[string]string) (response LeadFormsListResponse, vkErr Error) {
	vk.requestU("leadForms.list", params, &response, &vkErr)
	return
}

// LeadFormsUpdateResponse struct
type LeadFormsUpdateResponse struct {
	FormID int    `json:"form_id"`
	URL    string `json:"url"`
}

// LeadFormsUpdate leadForms.update
// https://vk.com/dev/leadForms.update
func (vk *VK) LeadFormsUpdate(params map[string]string) (response LeadFormsUpdateResponse, vkErr Error) {
	vk.requestU("leadForms.update", params, &response, &vkErr)
	return
}
