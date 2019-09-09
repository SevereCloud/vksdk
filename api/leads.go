package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// LeadsCheckUserResponse struct
type LeadsCheckUserResponse object.LeadsChecked

// LeadsCheckUser Checks if the user can start the lead.
//
// https://vk.com/dev/leads.checkUser
func (vk *VK) LeadsCheckUser(params map[string]string) (response LeadsCheckUserResponse, vkErr Error) {
	vk.RequestUnmarshal("leads.checkUser", params, &response, &vkErr)
	return
}

// LeadsCompleteResponse struct
type LeadsCompleteResponse object.LeadsComplete

// LeadsComplete Completes the lead started by user.
//
// https://vk.com/dev/leads.complete
func (vk *VK) LeadsComplete(params map[string]string) (response LeadsCompleteResponse, vkErr Error) {
	vk.RequestUnmarshal("leads.complete", params, &response, &vkErr)
	return
}

// LeadsGetStatsResponse struct
type LeadsGetStatsResponse object.LeadsLead

// LeadsGetStats Returns lead stats data.
//
// https://vk.com/dev/leads.getStats
func (vk *VK) LeadsGetStats(params map[string]string) (response LeadsGetStatsResponse, vkErr Error) {
	vk.RequestUnmarshal("leads.getStats", params, &response, &vkErr)
	return
}

// LeadsGetUsersResponse struct
type LeadsGetUsersResponse object.LeadsEntry

// LeadsGetUsers Returns a list of last user actions for the offer.
//
// https://vk.com/dev/leads.getUsers
func (vk *VK) LeadsGetUsers(params map[string]string) (response LeadsGetUsersResponse, vkErr Error) {
	vk.RequestUnmarshal("leads.getUsers", params, &response, &vkErr)
	return
}

// LeadsMetricHitResponse struct
type LeadsMetricHitResponse struct {
	Result       bool   `json:"result"`        // Information whether request has been processed successfully
	RedirectLink string `json:"redirect_link"` // Redirect link
}

// LeadsMetricHit Counts the metric event.
//
// https://vk.com/dev/leads.metricHit
func (vk *VK) LeadsMetricHit(params map[string]string) (response LeadsMetricHitResponse, vkErr Error) {
	vk.RequestUnmarshal("leads.metricHit", params, &response, &vkErr)
	return
}

// LeadsStartResponse struct
type LeadsStartResponse object.LeadsStart

// LeadsStart Creates new session for the user passing the offer.
//
// https://vk.com/dev/leads.start
func (vk *VK) LeadsStart(params map[string]string) (response LeadsStartResponse, vkErr Error) {
	vk.RequestUnmarshal("leads.start", params, &response, &vkErr)
	return
}
