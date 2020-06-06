package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// TODO: AdsAddOfficeUsersResponse struct.
// type AdsAddOfficeUsersResponse struct{}

// TODO: AdsAddOfficeUsers ...
//
// https://vk.com/dev/ads.addOfficeUsers
// func (vk *VK) AdsAddOfficeUsers(params Params) (response AdsAddOfficeUsersResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.addOfficeUsers", params, &response)
// 	return
// }

// TODO: AdsCheckLinkResponse struct.
// type AdsCheckLinkResponse struct{}

// TODO: AdsCheckLink ...
//
// https://vk.com/dev/ads.checkLink
// func (vk *VK) AdsCheckLink(params Params) (response AdsCheckLinkResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.checkLink", params, &response)
// 	return
// }

// TODO: AdsCreateAdsResponse struct.
// type AdsCreateAdsResponse struct{}

// TODO: AdsCreateAds ...
//
// https://vk.com/dev/ads.createAds
// func (vk *VK) AdsCreateAds(params Params) (response AdsCreateAdsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.createAds", params, &response)
// 	return
// }

// TODO: AdsCreateCampaignsResponse struct.
// type AdsCreateCampaignsResponse struct{}

// TODO: AdsCreateCampaigns ...
//
// https://vk.com/dev/ads.createCampaigns
// func (vk *VK) AdsCreateCampaigns(params Params) (response AdsCreateCampaignsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.createCampaigns", params, &response)
// 	return
// }

// TODO: AdsCreateClientsResponse struct.
// type AdsCreateClientsResponse struct{}

// TODO: AdsCreateClients ...
//
// https://vk.com/dev/ads.createClients
// func (vk *VK) AdsCreateClients(params Params) (response AdsCreateClientsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.createClients", params, &response)
// 	return
// }

// TODO: AdsCreateLookalikeRequestResponse struct.
// type AdsCreateLookalikeRequestResponse struct{}

// TODO: AdsCreateLookalikeRequest ...
//
// https://vk.com/dev/ads.createLookalikeRequest
// func (vk *VK) AdsCreateLookalikeRequest(params Params) (response AdsCreateLookalikeRequestResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.createLookalikeRequest", params, &response)
// 	return
// }

// TODO: AdsCreateTargetGroupResponse struct.
// type AdsCreateTargetGroupResponse struct{}

// TODO: AdsCreateTargetGroup ...
//
// https://vk.com/dev/ads.createTargetGroup
// func (vk *VK) AdsCreateTargetGroup(params Params) (response AdsCreateTargetGroupResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.createTargetGroup", params, &response)
// 	return
// }

// TODO: AdsCreateTargetPixelResponse struct.
// type AdsCreateTargetPixelResponse struct{}

// TODO: AdsCreateTargetPixel ...
//
// https://vk.com/dev/ads.createTargetPixel
// func (vk *VK) AdsCreateTargetPixel(params Params) (response AdsCreateTargetPixelResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.createTargetPixel", params, &response)
// 	return
// }

// TODO: AdsDeleteAdsResponse struct.
// type AdsDeleteAdsResponse struct{}

// TODO: AdsDeleteAds ...
//
// https://vk.com/dev/ads.deleteAds
// func (vk *VK) AdsDeleteAds(params Params) (response AdsDeleteAdsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.deleteAds", params, &response)
// 	return
// }

// TODO: AdsDeleteCampaignsResponse struct.
// type AdsDeleteCampaignsResponse struct{}

// TODO: AdsDeleteCampaigns ...
//
// https://vk.com/dev/ads.deleteCampaigns
// func (vk *VK) AdsDeleteCampaigns(params Params) (response AdsDeleteCampaignsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.deleteCampaigns", params, &response)
// 	return
// }

// TODO: AdsDeleteClientsResponse struct.
// type AdsDeleteClientsResponse struct{}

// TODO: AdsDeleteClients ...
//
// https://vk.com/dev/ads.deleteClients
// func (vk *VK) AdsDeleteClients(params Params) (response AdsDeleteClientsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.deleteClients", params, &response)
// 	return
// }

// AdsDeleteTargetGroup deletes target group.
//
// https://vk.com/dev/ads.deleteTargetGroup
func (vk *VK) AdsDeleteTargetGroup(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("ads.deleteTargetGroup", params, &response)
	return
}

// AdsDeleteTargetPixel deletes target pixel.
//
// https://vk.com/dev/ads.deleteTargetPixel
func (vk *VK) AdsDeleteTargetPixel(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("ads.deleteTargetPixel", params, &response)
	return
}

// AdsGetAccountsResponse struct.
type AdsGetAccountsResponse []object.AdsAccount

// AdsGetAccounts returns a list of advertising accounts.
//
// https://vk.com/dev/ads.getAccounts
func (vk *VK) AdsGetAccounts(params Params) (response AdsGetAccountsResponse, err error) {
	err = vk.RequestUnmarshal("ads.getAccounts", params, &response)
	return
}

// TODO: AdsGetAdsResponse struct.
// type AdsGetAdsResponse struct{}

// AdsGetAds TODO: ...
//
// https://vk.com/dev/ads.getAds
// func (vk *VK) AdsGetAds(params Params) (response AdsGetAdsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getAds", params, &response)
// 	return
// }

// TODO: AdsGetAdsLayoutResponse struct.
// type AdsGetAdsLayoutResponse struct{}

// TODO: AdsGetAdsLayout ...
//
// https://vk.com/dev/ads.getAdsLayout
// func (vk *VK) AdsGetAdsLayout(params Params) (response AdsGetAdsLayoutResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getAdsLayout", params, &response)
// 	return
// }

// TODO: AdsGetAdsTargetingResponse struct.
// type AdsGetAdsTargetingResponse struct{}

// TODO: AdsGetAdsTargeting ...
//
// https://vk.com/dev/ads.getAdsTargeting
// func (vk *VK) AdsGetAdsTargeting(params Params) (response AdsGetAdsTargetingResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getAdsTargeting", params, &response)
// 	return
// }

// TODO: AdsGetBudgetResponse struct.
// type AdsGetBudgetResponse struct{}

// TODO: AdsGetBudget ...
//
// https://vk.com/dev/ads.getBudget
// func (vk *VK) AdsGetBudget(params Params) (response AdsGetBudgetResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getBudget", params, &response)
// 	return
// }

// TODO: AdsGetCampaignsResponse struct.
// type AdsGetCampaignsResponse struct{}

// TODO: AdsGetCampaigns ...
//
// https://vk.com/dev/ads.getCampaigns
// func (vk *VK) AdsGetCampaigns(params Params) (response AdsGetCampaignsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getCampaigns", params, &response)
// 	return
// }

// TODO: AdsGetCategoriesResponse struct.
// type AdsGetCategoriesResponse struct{}

// TODO: AdsGetCategories ...
//
// https://vk.com/dev/ads.getCategories
// func (vk *VK) AdsGetCategories(params Params) (response AdsGetCategoriesResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getCategories", params, &response)
// 	return
// }

// TODO: AdsGetClientsResponse struct.
// type AdsGetClientsResponse struct{}

// TODO: AdsGetClients ...
//
// https://vk.com/dev/ads.getClients
// func (vk *VK) AdsGetClients(params Params) (response AdsGetClientsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getClients", params, &response)
// 	return
// }

// TODO: AdsGetDemographicsResponse struct.
// type AdsGetDemographicsResponse struct{}

// TODO: AdsGetDemographics ...
//
// https://vk.com/dev/ads.getDemographics
// func (vk *VK) AdsGetDemographics(params Params) (response AdsGetDemographicsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getDemographics", params, &response)
// 	return
// }

// TODO: AdsGetFloodStatsResponse struct.
// type AdsGetFloodStatsResponse struct{}

// TODO: AdsGetFloodStats ...
//
// https://vk.com/dev/ads.getFloodStats
// func (vk *VK) AdsGetFloodStats(params Params) (response AdsGetFloodStatsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getFloodStats", params, &response)
// 	return
// }

// TODO: AdsGetLookalikeRequestsResponse struct.
// type AdsGetLookalikeRequestsResponse struct{}

// TODO: AdsGetLookalikeRequests ...
//
// https://vk.com/dev/ads.getLookalikeRequests
// func (vk *VK) AdsGetLookalikeRequests(params Params) (response AdsGetLookalikeRequestsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getLookalikeRequests", params, &response)
// 	return
// }

// TODO: AdsGetOfficeUsersResponse struct.
// type AdsGetOfficeUsersResponse struct{}

// TODO: AdsGetOfficeUsers ...
//
// https://vk.com/dev/ads.getOfficeUsers
// func (vk *VK) AdsGetOfficeUsers(params Params) (response AdsGetOfficeUsersResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getOfficeUsers", params, &response)
// 	return
// }

// TODO: AdsGetPostsReachResponse struct.
// type AdsGetPostsReachResponse struct{}

// TODO: AdsGetPostsReach ...
//
// https://vk.com/dev/ads.getPostsReach
// func (vk *VK) AdsGetPostsReach(params Params) (response AdsGetPostsReachResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getPostsReach", params, &response)
// 	return
// }

// TODO: AdsGetRejectionReasonResponse struct.
// type AdsGetRejectionReasonResponse struct{}

// TODO: AdsGetRejectionReason ...
//
// https://vk.com/dev/ads.getRejectionReason
// func (vk *VK) AdsGetRejectionReason(params Params) (response AdsGetRejectionReasonResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getRejectionReason", params, &response)
// 	return
// }

// TODO: AdsGetStatisticsResponse struct.
// type AdsGetStatisticsResponse struct{}

// TODO: AdsGetStatistics ...
//
// https://vk.com/dev/ads.getStatistics
// func (vk *VK) AdsGetStatistics(params Params) (response AdsGetStatisticsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getStatistics", params, &response)
// 	return
// }

// TODO: AdsGetSuggestionsResponse struct.
// type AdsGetSuggestionsResponse struct{}

// TODO: AdsGetSuggestions ...
//
// https://vk.com/dev/ads.getSuggestions
// func (vk *VK) AdsGetSuggestions(params Params) (response AdsGetSuggestionsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getSuggestions", params, &response)
// 	return
// }

// TODO: AdsGetTargetGroupsResponse struct.
// type AdsGetTargetGroupsResponse struct{}

// TODO: AdsGetTargetGroups ...
//
// https://vk.com/dev/ads.getTargetGroups
// func (vk *VK) AdsGetTargetGroups(params Params) (response AdsGetTargetGroupsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getTargetGroups", params, &response)
// 	return
// }

// TODO: AdsGetTargetPixelsResponse struct.
// type AdsGetTargetPixelsResponse struct{}

// TODO: AdsGetTargetPixels ...
//
// https://vk.com/dev/ads.getTargetPixels
// func (vk *VK) AdsGetTargetPixels(params Params) (response AdsGetTargetPixelsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getTargetPixels", params, &response)
// 	return
// }

// TODO: AdsGetTargetingStatsResponse struct.
// type AdsGetTargetingStatsResponse struct{}

// TODO: AdsGetTargetingStats ...
//
// https://vk.com/dev/ads.getTargetingStats
// func (vk *VK) AdsGetTargetingStats(params Params) (response AdsGetTargetingStatsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getTargetingStats", params, &response)
// 	return
// }

// TODO: AdsGetUploadURLResponse struct.
// type AdsGetUploadURLResponse struct{}

// TODO: AdsGetUploadURL ...
//
// https://vk.com/dev/ads.getUploadURL
// func (vk *VK) AdsGetUploadURL(params Params) (response AdsGetUploadURLResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getUploadURL", params, &response)
// 	return
// }

// TODO: AdsGetVideoUploadURLResponse struct.
// type AdsGetVideoUploadURLResponse struct{}

// TODO: AdsGetVideoUploadURL ...
//
// https://vk.com/dev/ads.getVideoUploadURL
// func (vk *VK) AdsGetVideoUploadURL(params Params) (response AdsGetVideoUploadURLResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.getVideoUploadURL", params, &response)
// 	return
// }

// TODO: AdsImportTargetContactsResponse struct.
// type AdsImportTargetContactsResponse struct{}

// TODO: AdsImportTargetContacts ...
//
// https://vk.com/dev/ads.importTargetContacts
// func (vk *VK) AdsImportTargetContacts(params Params) (response AdsImportTargetContactsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.importTargetContacts", params, &response)
// 	return
// }

// TODO: AdsRemoveOfficeUsersResponse struct.
// type AdsRemoveOfficeUsersResponse struct{}

// TODO: AdsRemoveOfficeUsers ...
//
// https://vk.com/dev/ads.removeOfficeUsers
// func (vk *VK) AdsRemoveOfficeUsers(params Params) (response AdsRemoveOfficeUsersResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.removeOfficeUsers", params, &response)
// 	return
// }

// AdsRemoveTargetContacts accepts the request to exclude the advertiser's
// contacts from the retargeting audience.
//
// The maximum allowed number of contacts to be excluded by a single
// request is 1000.
//
// Contacts are excluded within a few hours of the request.
//
// https://vk.com/dev/ads.removeTargetContacts
func (vk *VK) AdsRemoveTargetContacts(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("ads.removeTargetContacts", params, &response)
	return
}

// TODO: AdsSaveLookalikeRequestResultResponse struct.
// type AdsSaveLookalikeRequestResultResponse struct{}

// TODO: AdsSaveLookalikeRequestResult ...
//
// https://vk.com/dev/ads.saveLookalikeRequestResult
// func (vk *VK) AdsSaveLookalikeRequestResult(params Params) (response AdsSaveLookalikeRequestResultResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.saveLookalikeRequestResult", params, &response)
// 	return
// }

// TODO: AdsShareTargetGroupResponse struct.
// type AdsShareTargetGroupResponse struct{}

// TODO: AdsShareTargetGroup ...
//
// https://vk.com/dev/ads.shareTargetGroup
// func (vk *VK) AdsShareTargetGroup(params Params) (response AdsShareTargetGroupResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.shareTargetGroup", params, &response)
// 	return
// }

// TODO: AdsUpdateAdsResponse struct.
// type AdsUpdateAdsResponse struct{}

// TODO: AdsUpdateAds ...
//
// https://vk.com/dev/ads.updateAds
// func (vk *VK) AdsUpdateAds(params Params) (response AdsUpdateAdsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.updateAds", params, &response)
// 	return
// }

// TODO: AdsUpdateCampaignsResponse struct.
// type AdsUpdateCampaignsResponse struct{}

// TODO: AdsUpdateCampaigns ...
//
// https://vk.com/dev/ads.updateCampaigns
// func (vk *VK) AdsUpdateCampaigns(params Params) (response AdsUpdateCampaignsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.updateCampaigns", params, &response)
// 	return
// }

// TODO: AdsUpdateClientsResponse struct.
// type AdsUpdateClientsResponse struct{}

// TODO: AdsUpdateClients ...
//
// https://vk.com/dev/ads.updateClients
// func (vk *VK) AdsUpdateClients(params Params) (response AdsUpdateClientsResponse, err error) {
// 	err = vk.RequestUnmarshal("ads.updateClients", params, &response)
// 	return
// }

// AdsUpdateTargetGroup edits target group.
//
// https://vk.com/dev/ads.updateTargetGroup
func (vk *VK) AdsUpdateTargetGroup(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("ads.updateTargetGroup", params, &response)
	return
}

// AdsUpdateTargetPixel edits target pixel.
//
// https://vk.com/dev/ads.updateTargetPixel
func (vk *VK) AdsUpdateTargetPixel(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("ads. updateTargetPixel", params, &response)
	return
}
