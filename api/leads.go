package api // import "github.com/SevereCloud/vksdk/v3/api"

import (
	"github.com/SevereCloud/vksdk/v3/object"
)

// LeadsCheckUserResponse struct.
type LeadsCheckUserResponse object.LeadsChecked

// LeadsCheckUser checks if the user can start the lead.
//
// https://dev.vk.com/method/leads.checkUser
func (vk *VK) LeadsCheckUser(params Params) (response LeadsCheckUserResponse, err error) {
	err = vk.RequestUnmarshal("leads.checkUser", &response, params)
	return
}

// LeadsCompleteResponse struct.
type LeadsCompleteResponse object.LeadsComplete

// LeadsComplete completes the lead started by user.
//
// https://dev.vk.com/method/leads.complete
func (vk *VK) LeadsComplete(params Params) (response LeadsCompleteResponse, err error) {
	err = vk.RequestUnmarshal("leads.complete", &response, params)
	return
}

// LeadsGetStatsResponse struct.
type LeadsGetStatsResponse object.LeadsLead

// LeadsGetStats returns lead stats data.
//
// https://dev.vk.com/method/leads.getStats
func (vk *VK) LeadsGetStats(params Params) (response LeadsGetStatsResponse, err error) {
	err = vk.RequestUnmarshal("leads.getStats", &response, params)
	return
}

// LeadsGetUsersResponse struct.
type LeadsGetUsersResponse object.LeadsEntry

// LeadsGetUsers returns a list of last user actions for the offer.
//
// https://dev.vk.com/method/leads.getUsers
func (vk *VK) LeadsGetUsers(params Params) (response LeadsGetUsersResponse, err error) {
	err = vk.RequestUnmarshal("leads.getUsers", &response, params)
	return
}

// LeadsMetricHitResponse struct.
type LeadsMetricHitResponse struct {
	Result       object.BaseBoolInt `json:"result"`        // Information whether request has been processed successfully
	RedirectLink string             `json:"redirect_link"` // Redirect link
}

// LeadsMetricHit counts the metric event.
//
// https://dev.vk.com/method/leads.metricHit
func (vk *VK) LeadsMetricHit(params Params) (response LeadsMetricHitResponse, err error) {
	err = vk.RequestUnmarshal("leads.metricHit", &response, params)
	return
}

// LeadsStartResponse struct.
type LeadsStartResponse object.LeadsStart

// LeadsStart creates new session for the user passing the offer.
//
// https://dev.vk.com/method/leads.start
func (vk *VK) LeadsStart(params Params) (response LeadsStartResponse, err error) {
	err = vk.RequestUnmarshal("leads.start", &response, params)
	return
}
