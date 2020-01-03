package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// AdsAddOfficeUsersBulder builder
//
// Adds managers and/or supervisors to advertising account.
//
// https://vk.com/dev/ads.addOfficeUsers
type AdsAddOfficeUsersBulder struct {
	api.Params
}

// NewAdsAddOfficeUsersBulder func
func NewAdsAddOfficeUsersBulder() *AdsAddOfficeUsersBulder {
	return &AdsAddOfficeUsersBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsAddOfficeUsersBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// Data Serialized JSON array of objects that describe added managers. Description of 'user_specification' objects see below.
func (b *AdsAddOfficeUsersBulder) Data(v string) {
	b.Params["data"] = v
}

// AdsCheckLinkBulder builder
//
// Allows to check the ad link.
//
// https://vk.com/dev/ads.checkLink
type AdsCheckLinkBulder struct {
	api.Params
}

// NewAdsCheckLinkBulder func
func NewAdsCheckLinkBulder() *AdsCheckLinkBulder {
	return &AdsCheckLinkBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsCheckLinkBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// LinkType Object type: *'community' — community,, *'post' — community post,, *'application' — VK application,, *'video' — video,, *'site' — external site.
func (b *AdsCheckLinkBulder) LinkType(v string) {
	b.Params["link_type"] = v
}

// LinkURL Object URL.
func (b *AdsCheckLinkBulder) LinkURL(v string) {
	b.Params["link_url"] = v
}

// CampaignID Campaign ID
func (b *AdsCheckLinkBulder) CampaignID(v int) {
	b.Params["campaign_id"] = v
}

// AdsCreateAdsBulder builder
//
// Creates ads.
//
// https://vk.com/dev/ads.createAds
type AdsCreateAdsBulder struct {
	api.Params
}

// NewAdsCreateAdsBulder func
func NewAdsCreateAdsBulder() *AdsCreateAdsBulder {
	return &AdsCreateAdsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsCreateAdsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// Data Serialized JSON array of objects that describe created ads. Description of 'ad_specification' objects see below.
func (b *AdsCreateAdsBulder) Data(v string) {
	b.Params["data"] = v
}

// AdsCreateCampaignsBulder builder
//
// Creates advertising campaigns.
//
// https://vk.com/dev/ads.createCampaigns
type AdsCreateCampaignsBulder struct {
	api.Params
}

// NewAdsCreateCampaignsBulder func
func NewAdsCreateCampaignsBulder() *AdsCreateCampaignsBulder {
	return &AdsCreateCampaignsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsCreateCampaignsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// Data Serialized JSON array of objects that describe created campaigns. Description of 'campaign_specification' objects see below.
func (b *AdsCreateCampaignsBulder) Data(v string) {
	b.Params["data"] = v
}

// AdsCreateClientsBulder builder
//
// Creates clients of an advertising agency.
//
// https://vk.com/dev/ads.createClients
type AdsCreateClientsBulder struct {
	api.Params
}

// NewAdsCreateClientsBulder func
func NewAdsCreateClientsBulder() *AdsCreateClientsBulder {
	return &AdsCreateClientsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsCreateClientsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// Data Serialized JSON array of objects that describe created campaigns. Description of 'client_specification' objects see below.
func (b *AdsCreateClientsBulder) Data(v string) {
	b.Params["data"] = v
}

// AdsCreateTargetGroupBulder builder
//
// Creates a group to re-target ads for users who visited advertiser's site (viewed information about the product, registered, etc.).
//
// https://vk.com/dev/ads.createTargetGroup
type AdsCreateTargetGroupBulder struct {
	api.Params
}

// NewAdsCreateTargetGroupBulder func
func NewAdsCreateTargetGroupBulder() *AdsCreateTargetGroupBulder {
	return &AdsCreateTargetGroupBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsCreateTargetGroupBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// ClientID 'Only for advertising agencies.', ID of the client with the advertising account where the group will be created.
func (b *AdsCreateTargetGroupBulder) ClientID(v int) {
	b.Params["client_id"] = v
}

// Name Name of the target group — a string up to 64 characters long.
func (b *AdsCreateTargetGroupBulder) Name(v string) {
	b.Params["name"] = v
}

// Lifetime 'For groups with auditory created with pixel code only.', , Number of days after that users will be automatically removed from the group.
func (b *AdsCreateTargetGroupBulder) Lifetime(v int) {
	b.Params["lifetime"] = v
}

// TargetPixelID parameter
func (b *AdsCreateTargetGroupBulder) TargetPixelID(v int) {
	b.Params["target_pixel_id"] = v
}

// TargetPixelRules parameter
func (b *AdsCreateTargetGroupBulder) TargetPixelRules(v string) {
	b.Params["target_pixel_rules"] = v
}

// AdsDeleteAdsBulder builder
//
// Archives ads.
//
// https://vk.com/dev/ads.deleteAds
type AdsDeleteAdsBulder struct {
	api.Params
}

// NewAdsDeleteAdsBulder func
func NewAdsDeleteAdsBulder() *AdsDeleteAdsBulder {
	return &AdsDeleteAdsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsDeleteAdsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// IDs Serialized JSON array with ad IDs.
func (b *AdsDeleteAdsBulder) IDs(v string) {
	b.Params["ids"] = v
}

// AdsDeleteCampaignsBulder builder
//
// Archives advertising campaigns.
//
// https://vk.com/dev/ads.deleteCampaigns
type AdsDeleteCampaignsBulder struct {
	api.Params
}

// NewAdsDeleteCampaignsBulder func
func NewAdsDeleteCampaignsBulder() *AdsDeleteCampaignsBulder {
	return &AdsDeleteCampaignsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsDeleteCampaignsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// IDs Serialized JSON array with IDs of deleted campaigns.
func (b *AdsDeleteCampaignsBulder) IDs(v string) {
	b.Params["ids"] = v
}

// AdsDeleteClientsBulder builder
//
// Archives clients of an advertising agency.
//
// https://vk.com/dev/ads.deleteClients
type AdsDeleteClientsBulder struct {
	api.Params
}

// NewAdsDeleteClientsBulder func
func NewAdsDeleteClientsBulder() *AdsDeleteClientsBulder {
	return &AdsDeleteClientsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsDeleteClientsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// IDs Serialized JSON array with IDs of deleted clients.
func (b *AdsDeleteClientsBulder) IDs(v string) {
	b.Params["ids"] = v
}

// AdsDeleteTargetGroupBulder builder
//
// Deletes a retarget group.
//
// https://vk.com/dev/ads.deleteTargetGroup
type AdsDeleteTargetGroupBulder struct {
	api.Params
}

// NewAdsDeleteTargetGroupBulder func
func NewAdsDeleteTargetGroupBulder() *AdsDeleteTargetGroupBulder {
	return &AdsDeleteTargetGroupBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsDeleteTargetGroupBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// ClientID 'Only for advertising agencies.' , ID of the client with the advertising account where the group will be created.
func (b *AdsDeleteTargetGroupBulder) ClientID(v int) {
	b.Params["client_id"] = v
}

// TargetGroupID Group ID.
func (b *AdsDeleteTargetGroupBulder) TargetGroupID(v int) {
	b.Params["target_group_id"] = v
}

// AdsGetAdsBulder builder
//
// Returns number of ads.
//
// https://vk.com/dev/ads.getAds
type AdsGetAdsBulder struct {
	api.Params
}

// NewAdsGetAdsBulder func
func NewAdsGetAdsBulder() *AdsGetAdsBulder {
	return &AdsGetAdsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetAdsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// AdIDs Filter by ads. Serialized JSON array with ad IDs. If the parameter is null, all ads will be shown.
func (b *AdsGetAdsBulder) AdIDs(v string) {
	b.Params["ad_ids"] = v
}

// CampaignIDs Filter by advertising campaigns. Serialized JSON array with campaign IDs. If the parameter is null, ads of all campaigns will be shown.
func (b *AdsGetAdsBulder) CampaignIDs(v string) {
	b.Params["campaign_ids"] = v
}

// ClientID 'Available and required for advertising agencies.' ID of the client ads are retrieved from.
func (b *AdsGetAdsBulder) ClientID(v int) {
	b.Params["client_id"] = v
}

// IncludeDeleted Flag that specifies whether archived ads shall be shown: *0 — show only active ads,, *1 — show all ads.
func (b *AdsGetAdsBulder) IncludeDeleted(v bool) {
	b.Params["include_deleted"] = v
}

// Limit Limit of number of returned ads. Used only if ad_ids parameter is null, and 'campaign_ids' parameter contains ID of only one campaign.
func (b *AdsGetAdsBulder) Limit(v int) {
	b.Params["limit"] = v
}

// Offset Offset. Used in the same cases as 'limit' parameter.
func (b *AdsGetAdsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// AdsGetAdsLayoutBulder builder
//
// Returns descriptions of ad layouts.
//
// https://vk.com/dev/ads.getAdsLayout
type AdsGetAdsLayoutBulder struct {
	api.Params
}

// NewAdsGetAdsLayoutBulder func
func NewAdsGetAdsLayoutBulder() *AdsGetAdsLayoutBulder {
	return &AdsGetAdsLayoutBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetAdsLayoutBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// AdIDs Filter by ads. Serialized JSON array with ad IDs. If the parameter is null, all ads will be shown.
func (b *AdsGetAdsLayoutBulder) AdIDs(v string) {
	b.Params["ad_ids"] = v
}

// CampaignIDs Filter by advertising campaigns. Serialized JSON array with campaign IDs. If the parameter is null, ads of all campaigns will be shown.
func (b *AdsGetAdsLayoutBulder) CampaignIDs(v string) {
	b.Params["campaign_ids"] = v
}

// ClientID 'For advertising agencies.' ID of the client ads are retrieved from.
func (b *AdsGetAdsLayoutBulder) ClientID(v int) {
	b.Params["client_id"] = v
}

// IncludeDeleted Flag that specifies whether archived ads shall be shown. *0 — show only active ads,, *1 — show all ads.
func (b *AdsGetAdsLayoutBulder) IncludeDeleted(v bool) {
	b.Params["include_deleted"] = v
}

// Limit Limit of number of returned ads. Used only if 'ad_ids' parameter is null, and 'campaign_ids' parameter contains ID of only one campaign.
func (b *AdsGetAdsLayoutBulder) Limit(v int) {
	b.Params["limit"] = v
}

// Offset Offset. Used in the same cases as 'limit' parameter.
func (b *AdsGetAdsLayoutBulder) Offset(v int) {
	b.Params["offset"] = v
}

// AdsGetAdsTargetingBulder builder
//
// Returns ad targeting parameters.
//
// https://vk.com/dev/ads.getAdsTargeting
type AdsGetAdsTargetingBulder struct {
	api.Params
}

// NewAdsGetAdsTargetingBulder func
func NewAdsGetAdsTargetingBulder() *AdsGetAdsTargetingBulder {
	return &AdsGetAdsTargetingBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetAdsTargetingBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// AdIDs Filter by ads. Serialized JSON array with ad IDs. If the parameter is null, all ads will be shown.
func (b *AdsGetAdsTargetingBulder) AdIDs(v string) {
	b.Params["ad_ids"] = v
}

// CampaignIDs Filter by advertising campaigns. Serialized JSON array with campaign IDs. If the parameter is null, ads of all campaigns will be shown.
func (b *AdsGetAdsTargetingBulder) CampaignIDs(v string) {
	b.Params["campaign_ids"] = v
}

// ClientID 'For advertising agencies.' ID of the client ads are retrieved from.
func (b *AdsGetAdsTargetingBulder) ClientID(v int) {
	b.Params["client_id"] = v
}

// IncludeDeleted flag that specifies whether archived ads shall be shown: *0 — show only active ads,, *1 — show all ads.
func (b *AdsGetAdsTargetingBulder) IncludeDeleted(v bool) {
	b.Params["include_deleted"] = v
}

// Limit Limit of number of returned ads. Used only if 'ad_ids' parameter is null, and 'campaign_ids' parameter contains ID of only one campaign.
func (b *AdsGetAdsTargetingBulder) Limit(v int) {
	b.Params["limit"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *AdsGetAdsTargetingBulder) Offset(v int) {
	b.Params["offset"] = v
}

// AdsGetBudgetBulder builder
//
// Returns current budget of the advertising account.
//
// https://vk.com/dev/ads.getBudget
type AdsGetBudgetBulder struct {
	api.Params
}

// NewAdsGetBudgetBulder func
func NewAdsGetBudgetBulder() *AdsGetBudgetBulder {
	return &AdsGetBudgetBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetBudgetBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// AdsGetCampaignsBulder builder
//
// Returns a list of campaigns in an advertising account.
//
// https://vk.com/dev/ads.getCampaigns
type AdsGetCampaignsBulder struct {
	api.Params
}

// NewAdsGetCampaignsBulder func
func NewAdsGetCampaignsBulder() *AdsGetCampaignsBulder {
	return &AdsGetCampaignsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetCampaignsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// ClientID 'For advertising agencies'. ID of the client advertising campaigns are retrieved from.
func (b *AdsGetCampaignsBulder) ClientID(v int) {
	b.Params["client_id"] = v
}

// IncludeDeleted Flag that specifies whether archived ads shall be shown. *0 — show only active campaigns,, *1 — show all campaigns.
func (b *AdsGetCampaignsBulder) IncludeDeleted(v bool) {
	b.Params["include_deleted"] = v
}

// CampaignIDs Filter of advertising campaigns to show. Serialized JSON array with campaign IDs. Only campaigns that exist in 'campaign_ids' and belong to the specified advertising account will be shown. If the parameter is null, all campaigns will be shown.
func (b *AdsGetCampaignsBulder) CampaignIDs(v string) {
	b.Params["campaign_ids"] = v
}

// AdsGetCategoriesBulder builder
//
// Returns a list of possible ad categories.
//
// https://vk.com/dev/ads.getCategories
type AdsGetCategoriesBulder struct {
	api.Params
}

// NewAdsGetCategoriesBulder func
func NewAdsGetCategoriesBulder() *AdsGetCategoriesBulder {
	return &AdsGetCategoriesBulder{api.Params{}}
}

// Lang Language. The full list of supported languages is [vk.com/dev/api_requests|here].
func (b *AdsGetCategoriesBulder) Lang(v string) {
	b.Params["lang"] = v
}

// AdsGetClientsBulder builder
//
// Returns a list of advertising agency's clients.
//
// https://vk.com/dev/ads.getClients
type AdsGetClientsBulder struct {
	api.Params
}

// NewAdsGetClientsBulder func
func NewAdsGetClientsBulder() *AdsGetClientsBulder {
	return &AdsGetClientsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetClientsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// AdsGetDemographicsBulder builder
//
// Returns demographics for ads or campaigns.
//
// https://vk.com/dev/ads.getDemographics
type AdsGetDemographicsBulder struct {
	api.Params
}

// NewAdsGetDemographicsBulder func
func NewAdsGetDemographicsBulder() *AdsGetDemographicsBulder {
	return &AdsGetDemographicsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetDemographicsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// IDsType Type of requested objects listed in 'ids' parameter: *ad — ads,, *campaign — campaigns.
func (b *AdsGetDemographicsBulder) IDsType(v string) {
	b.Params["ids_type"] = v
}

// IDs IDs requested ads or campaigns, separated with a comma, depending on the value set in 'ids_type'. Maximum 2000 objects.
func (b *AdsGetDemographicsBulder) IDs(v string) {
	b.Params["ids"] = v
}

// Period Data grouping by dates: *day — statistics by days,, *month — statistics by months,, *overall — overall statistics. 'date_from' and 'date_to' parameters set temporary limits.
func (b *AdsGetDemographicsBulder) Period(v string) {
	b.Params["period"] = v
}

// DateFrom Date to show statistics from. For different value of 'period' different date format is used: *day: YYYY-MM-DD, example: 2011-09-27 — September 27, 2011, **0 — day it was created on,, *month: YYYY-MM, example: 2011-09 — September 2011, **0 — month it was created in,, *overall: 0.
func (b *AdsGetDemographicsBulder) DateFrom(v string) {
	b.Params["date_from"] = v
}

// DateTo Date to show statistics to. For different value of 'period' different date format is used: *day: YYYY-MM-DD, example: 2011-09-27 — September 27, 2011, **0 — current day,, *month: YYYY-MM, example: 2011-09 — September 2011, **0 — current month,, *overall: 0.
func (b *AdsGetDemographicsBulder) DateTo(v string) {
	b.Params["date_to"] = v
}

// AdsGetFloodStatsBulder builder
//
// Returns information about current state of a counter — number of remaining runs of methods and time to the next counter nulling in seconds.
//
// https://vk.com/dev/ads.getFloodStats
type AdsGetFloodStatsBulder struct {
	api.Params
}

// NewAdsGetFloodStatsBulder func
func NewAdsGetFloodStatsBulder() *AdsGetFloodStatsBulder {
	return &AdsGetFloodStatsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetFloodStatsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// AdsGetOfficeUsersBulder builder
//
// Returns a list of managers and supervisors of advertising account.
//
// https://vk.com/dev/ads.getOfficeUsers
type AdsGetOfficeUsersBulder struct {
	api.Params
}

// NewAdsGetOfficeUsersBulder func
func NewAdsGetOfficeUsersBulder() *AdsGetOfficeUsersBulder {
	return &AdsGetOfficeUsersBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetOfficeUsersBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// AdsGetPostsReachBulder builder
//
// Returns detailed statistics of promoted posts reach from campaigns and ads.
//
// https://vk.com/dev/ads.getPostsReach
type AdsGetPostsReachBulder struct {
	api.Params
}

// NewAdsGetPostsReachBulder func
func NewAdsGetPostsReachBulder() *AdsGetPostsReachBulder {
	return &AdsGetPostsReachBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetPostsReachBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// IDsType Type of requested objects listed in 'ids' parameter: *ad — ads,, *campaign — campaigns.
func (b *AdsGetPostsReachBulder) IDsType(v string) {
	b.Params["ids_type"] = v
}

// IDs IDs requested ads or campaigns, separated with a comma, depending on the value set in 'ids_type'. Maximum 100 objects.
func (b *AdsGetPostsReachBulder) IDs(v string) {
	b.Params["ids"] = v
}

// AdsGetRejectionReasonBulder builder
//
// Returns a reason of ad rejection for pre-moderation.
//
// https://vk.com/dev/ads.getRejectionReason
type AdsGetRejectionReasonBulder struct {
	api.Params
}

// NewAdsGetRejectionReasonBulder func
func NewAdsGetRejectionReasonBulder() *AdsGetRejectionReasonBulder {
	return &AdsGetRejectionReasonBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetRejectionReasonBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// AdID Ad ID.
func (b *AdsGetRejectionReasonBulder) AdID(v int) {
	b.Params["ad_id"] = v
}

// AdsGetStatisticsBulder builder
//
// Returns statistics of performance indicators for ads, campaigns, clients or the whole account.
//
// https://vk.com/dev/ads.getStatistics
type AdsGetStatisticsBulder struct {
	api.Params
}

// NewAdsGetStatisticsBulder func
func NewAdsGetStatisticsBulder() *AdsGetStatisticsBulder {
	return &AdsGetStatisticsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetStatisticsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// IDsType Type of requested objects listed in 'ids' parameter: *ad — ads,, *campaign — campaigns,, *client — clients,, *office — account.
func (b *AdsGetStatisticsBulder) IDsType(v string) {
	b.Params["ids_type"] = v
}

// IDs IDs requested ads, campaigns, clients or account, separated with a comma, depending on the value set in 'ids_type'. Maximum 2000 objects.
func (b *AdsGetStatisticsBulder) IDs(v string) {
	b.Params["ids"] = v
}

// Period Data grouping by dates: *day — statistics by days,, *month — statistics by months,, *overall — overall statistics. 'date_from' and 'date_to' parameters set temporary limits.
func (b *AdsGetStatisticsBulder) Period(v string) {
	b.Params["period"] = v
}

// DateFrom Date to show statistics from. For different value of 'period' different date format is used: *day: YYYY-MM-DD, example: 2011-09-27 — September 27, 2011, **0 — day it was created on,, *month: YYYY-MM, example: 2011-09 — September 2011, **0 — month it was created in,, *overall: 0.
func (b *AdsGetStatisticsBulder) DateFrom(v string) {
	b.Params["date_from"] = v
}

// DateTo Date to show statistics to. For different value of 'period' different date format is used: *day: YYYY-MM-DD, example: 2011-09-27 — September 27, 2011, **0 — current day,, *month: YYYY-MM, example: 2011-09 — September 2011, **0 — current month,, *overall: 0.
func (b *AdsGetStatisticsBulder) DateTo(v string) {
	b.Params["date_to"] = v
}

// AdsGetSuggestionsBulder builder
//
// Returns a set of auto-suggestions for various targeting parameters.
//
// https://vk.com/dev/ads.getSuggestions
type AdsGetSuggestionsBulder struct {
	api.Params
}

// NewAdsGetSuggestionsBulder func
func NewAdsGetSuggestionsBulder() *AdsGetSuggestionsBulder {
	return &AdsGetSuggestionsBulder{api.Params{}}
}

// Section Section, suggestions are retrieved in. Available values: *countries — request of a list of countries. If q is not set or blank, a short list of countries is shown. Otherwise, a full list of countries is shown. *regions — requested list of regions. 'country' parameter is required. *cities — requested list of cities. 'country' parameter is required. *districts — requested list of districts. 'cities' parameter is required. *stations — requested list of subway stations. 'cities' parameter is required. *streets — requested list of streets. 'cities' parameter is required. *schools — requested list of educational organizations. 'cities' parameter is required. *interests — requested list of interests. *positions — requested list of positions (professions). *group_types — requested list of group types. *religions — requested list of religious commitments. *browsers — requested list of browsers and mobile devices.
func (b *AdsGetSuggestionsBulder) Section(v string) {
	b.Params["section"] = v
}

// IDs Objects IDs separated by commas. If the parameter is passed, 'q, country, cities' should not be passed.
func (b *AdsGetSuggestionsBulder) IDs(v string) {
	b.Params["ids"] = v
}

// Q Filter-line of the request (for countries, regions, cities, streets, schools, interests, positions).
func (b *AdsGetSuggestionsBulder) Q(v string) {
	b.Params["q"] = v
}

// Country ID of the country objects are searched in.
func (b *AdsGetSuggestionsBulder) Country(v int) {
	b.Params["country"] = v
}

// Cities IDs of cities where objects are searched in, separated with a comma.
func (b *AdsGetSuggestionsBulder) Cities(v string) {
	b.Params["cities"] = v
}

// Lang Language of the returned string values. Supported languages: *ru — Russian,, *ua — Ukrainian,, *en — English.
func (b *AdsGetSuggestionsBulder) Lang(v string) {
	b.Params["lang"] = v
}

// AdsGetTargetGroupsBulder builder
//
// Returns a list of target groups.
//
// https://vk.com/dev/ads.getTargetGroups
type AdsGetTargetGroupsBulder struct {
	api.Params
}

// NewAdsGetTargetGroupsBulder func
func NewAdsGetTargetGroupsBulder() *AdsGetTargetGroupsBulder {
	return &AdsGetTargetGroupsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetTargetGroupsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// ClientID 'Only for advertising agencies.', ID of the client with the advertising account where the group will be created.
func (b *AdsGetTargetGroupsBulder) ClientID(v int) {
	b.Params["client_id"] = v
}

// Extended '1' — to return pixel code.
func (b *AdsGetTargetGroupsBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// AdsGetTargetingStatsBulder builder
//
// Returns the size of targeting audience, and also recommended values for CPC and CPM.
//
// https://vk.com/dev/ads.getTargetingStats
type AdsGetTargetingStatsBulder struct {
	api.Params
}

// NewAdsGetTargetingStatsBulder func
func NewAdsGetTargetingStatsBulder() *AdsGetTargetingStatsBulder {
	return &AdsGetTargetingStatsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetTargetingStatsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// ClientID parameter
func (b *AdsGetTargetingStatsBulder) ClientID(v int) {
	b.Params["client_id"] = v
}

// Criteria Serialized JSON object that describes targeting parameters. Description of 'criteria' object see below.
func (b *AdsGetTargetingStatsBulder) Criteria(v string) {
	b.Params["criteria"] = v
}

// AdID ID of an ad which targeting parameters shall be analyzed.
func (b *AdsGetTargetingStatsBulder) AdID(v int) {
	b.Params["ad_id"] = v
}

// AdFormat Ad format. Possible values: *'1' — image and text,, *'2' — big image,, *'3' — exclusive format,, *'4' — community, square image,, *'7' — special app format,, *'8' — special community format,, *'9' — post in community,, *'10' — app board.
func (b *AdsGetTargetingStatsBulder) AdFormat(v int) {
	b.Params["ad_format"] = v
}

// AdPlatform Platforms to use for ad showing. Possible values: (for 'ad_format' = '1'), *'0' — VK and partner sites,, *'1' — VK only. (for 'ad_format' = '9'), *'all' — all platforms,, *'desktop' — desktop version,, *'mobile' — mobile version and apps.
func (b *AdsGetTargetingStatsBulder) AdPlatform(v string) {
	b.Params["ad_platform"] = v
}

// AdPlatformNoWall parameter
func (b *AdsGetTargetingStatsBulder) AdPlatformNoWall(v string) {
	b.Params["ad_platform_no_wall"] = v
}

// AdPlatformNoAdNetwork parameter
func (b *AdsGetTargetingStatsBulder) AdPlatformNoAdNetwork(v string) {
	b.Params["ad_platform_no_ad_network"] = v
}

// LinkURL URL for the advertised object.
func (b *AdsGetTargetingStatsBulder) LinkURL(v string) {
	b.Params["link_url"] = v
}

// LinkDomain Domain of the advertised object.
func (b *AdsGetTargetingStatsBulder) LinkDomain(v string) {
	b.Params["link_domain"] = v
}

// AdsGetUploadURLBulder builder
//
// Returns URL to upload an ad photo to.
//
// https://vk.com/dev/ads.getUploadURL
type AdsGetUploadURLBulder struct {
	api.Params
}

// NewAdsGetUploadURLBulder func
func NewAdsGetUploadURLBulder() *AdsGetUploadURLBulder {
	return &AdsGetUploadURLBulder{api.Params{}}
}

// AdFormat Ad format: *1 — image and text,, *2 — big image,, *3 — exclusive format,, *4 — community, square image,, *7 — special app format.
func (b *AdsGetUploadURLBulder) AdFormat(v int) {
	b.Params["ad_format"] = v
}

// Icon parameter
func (b *AdsGetUploadURLBulder) Icon(v int) {
	b.Params["icon"] = v
}

// AdsImportTargetContactsBulder builder
//
// Imports a list of advertiser's contacts to count VK registered users against the target group.
//
// https://vk.com/dev/ads.importTargetContacts
type AdsImportTargetContactsBulder struct {
	api.Params
}

// NewAdsImportTargetContactsBulder func
func NewAdsImportTargetContactsBulder() *AdsImportTargetContactsBulder {
	return &AdsImportTargetContactsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsImportTargetContactsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// ClientID 'Only for advertising agencies.' , ID of the client with the advertising account where the group will be created.
func (b *AdsImportTargetContactsBulder) ClientID(v int) {
	b.Params["client_id"] = v
}

// TargetGroupID Target group ID.
func (b *AdsImportTargetContactsBulder) TargetGroupID(v int) {
	b.Params["target_group_id"] = v
}

// Contacts List of phone numbers, emails or user IDs separated with a comma.
func (b *AdsImportTargetContactsBulder) Contacts(v string) {
	b.Params["contacts"] = v
}

// AdsRemoveOfficeUsersBulder builder
//
// Removes managers and/or supervisors from advertising account.
//
// https://vk.com/dev/ads.removeOfficeUsers
type AdsRemoveOfficeUsersBulder struct {
	api.Params
}

// NewAdsRemoveOfficeUsersBulder func
func NewAdsRemoveOfficeUsersBulder() *AdsRemoveOfficeUsersBulder {
	return &AdsRemoveOfficeUsersBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsRemoveOfficeUsersBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// IDs Serialized JSON array with IDs of deleted managers.
func (b *AdsRemoveOfficeUsersBulder) IDs(v string) {
	b.Params["ids"] = v
}

// AdsUpdateAdsBulder builder
//
// Edits ads.
//
// https://vk.com/dev/ads.updateAds
type AdsUpdateAdsBulder struct {
	api.Params
}

// NewAdsUpdateAdsBulder func
func NewAdsUpdateAdsBulder() *AdsUpdateAdsBulder {
	return &AdsUpdateAdsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsUpdateAdsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// Data Serialized JSON array of objects that describe changes in ads. Description of 'ad_edit_specification' objects see below.
func (b *AdsUpdateAdsBulder) Data(v string) {
	b.Params["data"] = v
}

// AdsUpdateCampaignsBulder builder
//
// Edits advertising campaigns.
//
// https://vk.com/dev/ads.updateCampaigns
type AdsUpdateCampaignsBulder struct {
	api.Params
}

// NewAdsUpdateCampaignsBulder func
func NewAdsUpdateCampaignsBulder() *AdsUpdateCampaignsBulder {
	return &AdsUpdateCampaignsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsUpdateCampaignsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// Data Serialized JSON array of objects that describe changes in campaigns. Description of 'campaign_mod' objects see below.
func (b *AdsUpdateCampaignsBulder) Data(v string) {
	b.Params["data"] = v
}

// AdsUpdateClientsBulder builder
//
// Edits clients of an advertising agency.
//
// https://vk.com/dev/ads.updateClients
type AdsUpdateClientsBulder struct {
	api.Params
}

// NewAdsUpdateClientsBulder func
func NewAdsUpdateClientsBulder() *AdsUpdateClientsBulder {
	return &AdsUpdateClientsBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsUpdateClientsBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// Data Serialized JSON array of objects that describe changes in clients. Description of 'client_mod' objects see below.
func (b *AdsUpdateClientsBulder) Data(v string) {
	b.Params["data"] = v
}

// AdsUpdateTargetGroupBulder builder
//
// Edits a retarget group.
//
// https://vk.com/dev/ads.updateTargetGroup
type AdsUpdateTargetGroupBulder struct {
	api.Params
}

// NewAdsUpdateTargetGroupBulder func
func NewAdsUpdateTargetGroupBulder() *AdsUpdateTargetGroupBulder {
	return &AdsUpdateTargetGroupBulder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsUpdateTargetGroupBulder) AccountID(v int) {
	b.Params["account_id"] = v
}

// ClientID 'Only for advertising agencies.' , ID of the client with the advertising account where the group will be created.
func (b *AdsUpdateTargetGroupBulder) ClientID(v int) {
	b.Params["client_id"] = v
}

// TargetGroupID Group ID.
func (b *AdsUpdateTargetGroupBulder) TargetGroupID(v int) {
	b.Params["target_group_id"] = v
}

// Name New name of the target group — a string up to 64 characters long.
func (b *AdsUpdateTargetGroupBulder) Name(v string) {
	b.Params["name"] = v
}

// Domain Domain of the site where user accounting code will be placed.
func (b *AdsUpdateTargetGroupBulder) Domain(v string) {
	b.Params["domain"] = v
}

// Lifetime 'Only for the groups that get audience from sites with user accounting code.', Time in days when users added to a retarget group will be automatically excluded from it. '0' – automatic exclusion is off.
func (b *AdsUpdateTargetGroupBulder) Lifetime(v int) {
	b.Params["lifetime"] = v
}

// TargetPixelID parameter
func (b *AdsUpdateTargetGroupBulder) TargetPixelID(v int) {
	b.Params["target_pixel_id"] = v
}

// TargetPixelRules parameter
func (b *AdsUpdateTargetGroupBulder) TargetPixelRules(v string) {
	b.Params["target_pixel_rules"] = v
}
