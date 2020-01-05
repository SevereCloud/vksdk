package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// AdsAddOfficeUsersBuilder builder
//
// Adds managers and/or supervisors to advertising account.
//
// https://vk.com/dev/ads.addOfficeUsers
type AdsAddOfficeUsersBuilder struct {
	api.Params
}

// NewAdsAddOfficeUsersBuilder func
func NewAdsAddOfficeUsersBuilder() *AdsAddOfficeUsersBuilder {
	return &AdsAddOfficeUsersBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsAddOfficeUsersBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// Data Serialized JSON array of objects that describe added managers.
// Description of 'user_specification' objects see below.
func (b *AdsAddOfficeUsersBuilder) Data(v string) {
	b.Params["data"] = v
}

// AdsCheckLinkBuilder builder
//
// Allows to check the ad link.
//
// https://vk.com/dev/ads.checkLink
type AdsCheckLinkBuilder struct {
	api.Params
}

// NewAdsCheckLinkBuilder func
func NewAdsCheckLinkBuilder() *AdsCheckLinkBuilder {
	return &AdsCheckLinkBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsCheckLinkBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// LinkType Object type:
//
// * community — community,
//
// * post — community post,
//
// * application — VK application,
//
// * video — video,
//
// * site — external site.
func (b *AdsCheckLinkBuilder) LinkType(v string) {
	b.Params["link_type"] = v
}

// LinkURL Object URL.
func (b *AdsCheckLinkBuilder) LinkURL(v string) {
	b.Params["link_url"] = v
}

// CampaignID Campaign ID
func (b *AdsCheckLinkBuilder) CampaignID(v int) {
	b.Params["campaign_id"] = v
}

// AdsCreateAdsBuilder builder
//
// Creates ads.
//
// https://vk.com/dev/ads.createAds
type AdsCreateAdsBuilder struct {
	api.Params
}

// NewAdsCreateAdsBuilder func
func NewAdsCreateAdsBuilder() *AdsCreateAdsBuilder {
	return &AdsCreateAdsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsCreateAdsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// Data Serialized JSON array of objects that describe created ads.
// Description of 'ad_specification' objects see below.
func (b *AdsCreateAdsBuilder) Data(v string) {
	b.Params["data"] = v
}

// AdsCreateCampaignsBuilder builder
//
// Creates advertising campaigns.
//
// https://vk.com/dev/ads.createCampaigns
type AdsCreateCampaignsBuilder struct {
	api.Params
}

// NewAdsCreateCampaignsBuilder func
func NewAdsCreateCampaignsBuilder() *AdsCreateCampaignsBuilder {
	return &AdsCreateCampaignsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsCreateCampaignsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// Data Serialized JSON array of objects that describe created campaigns.
// Description of 'campaign_specification' objects see below.
func (b *AdsCreateCampaignsBuilder) Data(v string) {
	b.Params["data"] = v
}

// AdsCreateClientsBuilder builder
//
// Creates clients of an advertising agency.
//
// https://vk.com/dev/ads.createClients
type AdsCreateClientsBuilder struct {
	api.Params
}

// NewAdsCreateClientsBuilder func
func NewAdsCreateClientsBuilder() *AdsCreateClientsBuilder {
	return &AdsCreateClientsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsCreateClientsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// Data Serialized JSON array of objects that describe created campaigns.
// Description of 'client_specification' objects see below.
func (b *AdsCreateClientsBuilder) Data(v string) {
	b.Params["data"] = v
}

// AdsCreateTargetGroupBuilder builder
//
// Creates a group to re-target ads for users who visited advertiser's site
// (viewed information about the product, registered, etc.).
//
// https://vk.com/dev/ads.createTargetGroup
type AdsCreateTargetGroupBuilder struct {
	api.Params
}

// NewAdsCreateTargetGroupBuilder func
func NewAdsCreateTargetGroupBuilder() *AdsCreateTargetGroupBuilder {
	return &AdsCreateTargetGroupBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsCreateTargetGroupBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// ClientID Only for advertising agencies.
// ID of the client with the advertising account where the group will be created.
func (b *AdsCreateTargetGroupBuilder) ClientID(v int) {
	b.Params["client_id"] = v
}

// Name Name of the target group — a string up to 64 characters long.
func (b *AdsCreateTargetGroupBuilder) Name(v string) {
	b.Params["name"] = v
}

// Lifetime For groups with auditory created with pixel code only.
// Number of days after that users will be automatically removed from the group.
func (b *AdsCreateTargetGroupBuilder) Lifetime(v int) {
	b.Params["lifetime"] = v
}

// TargetPixelID parameter
func (b *AdsCreateTargetGroupBuilder) TargetPixelID(v int) {
	b.Params["target_pixel_id"] = v
}

// TargetPixelRules parameter
func (b *AdsCreateTargetGroupBuilder) TargetPixelRules(v string) {
	b.Params["target_pixel_rules"] = v
}

// AdsDeleteAdsBuilder builder
//
// Archives ads.
//
// https://vk.com/dev/ads.deleteAds
type AdsDeleteAdsBuilder struct {
	api.Params
}

// NewAdsDeleteAdsBuilder func
func NewAdsDeleteAdsBuilder() *AdsDeleteAdsBuilder {
	return &AdsDeleteAdsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsDeleteAdsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// IDs Serialized JSON array with ad IDs.
func (b *AdsDeleteAdsBuilder) IDs(v string) {
	b.Params["ids"] = v
}

// AdsDeleteCampaignsBuilder builder
//
// Archives advertising campaigns.
//
// https://vk.com/dev/ads.deleteCampaigns
type AdsDeleteCampaignsBuilder struct {
	api.Params
}

// NewAdsDeleteCampaignsBuilder func
func NewAdsDeleteCampaignsBuilder() *AdsDeleteCampaignsBuilder {
	return &AdsDeleteCampaignsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsDeleteCampaignsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// IDs Serialized JSON array with IDs of deleted campaigns.
func (b *AdsDeleteCampaignsBuilder) IDs(v string) {
	b.Params["ids"] = v
}

// AdsDeleteClientsBuilder builder
//
// Archives clients of an advertising agency.
//
// https://vk.com/dev/ads.deleteClients
type AdsDeleteClientsBuilder struct {
	api.Params
}

// NewAdsDeleteClientsBuilder func
func NewAdsDeleteClientsBuilder() *AdsDeleteClientsBuilder {
	return &AdsDeleteClientsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsDeleteClientsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// IDs Serialized JSON array with IDs of deleted clients.
func (b *AdsDeleteClientsBuilder) IDs(v string) {
	b.Params["ids"] = v
}

// AdsDeleteTargetGroupBuilder builder
//
// Deletes a retarget group.
//
// https://vk.com/dev/ads.deleteTargetGroup
type AdsDeleteTargetGroupBuilder struct {
	api.Params
}

// NewAdsDeleteTargetGroupBuilder func
func NewAdsDeleteTargetGroupBuilder() *AdsDeleteTargetGroupBuilder {
	return &AdsDeleteTargetGroupBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsDeleteTargetGroupBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// ClientID Only for advertising agencies.
// ID of the client with the advertising account where the group will be created.
func (b *AdsDeleteTargetGroupBuilder) ClientID(v int) {
	b.Params["client_id"] = v
}

// TargetGroupID Group ID.
func (b *AdsDeleteTargetGroupBuilder) TargetGroupID(v int) {
	b.Params["target_group_id"] = v
}

// AdsGetAdsBuilder builder
//
// Returns number of ads.
//
// https://vk.com/dev/ads.getAds
type AdsGetAdsBuilder struct {
	api.Params
}

// NewAdsGetAdsBuilder func
func NewAdsGetAdsBuilder() *AdsGetAdsBuilder {
	return &AdsGetAdsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetAdsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// AdIDs Filter by ads. Serialized JSON array with ad IDs. If the parameter is null, all ads will be shown.
func (b *AdsGetAdsBuilder) AdIDs(v string) {
	b.Params["ad_ids"] = v
}

// CampaignIDs Filter by advertising campaigns. Serialized JSON array with campaign IDs. If the parameter is null,
// ads of all campaigns will be shown.
func (b *AdsGetAdsBuilder) CampaignIDs(v string) {
	b.Params["campaign_ids"] = v
}

// ClientID 'Available and required for advertising agencies.' ID of the client ads are retrieved from.
func (b *AdsGetAdsBuilder) ClientID(v int) {
	b.Params["client_id"] = v
}

// IncludeDeleted Flag that specifies whether archived ads shall be shown:
//
// * 0 — show only active ads,
//
// * 1 — show all ads.
func (b *AdsGetAdsBuilder) IncludeDeleted(v bool) {
	b.Params["include_deleted"] = v
}

// Limit Limit of number of returned ads. Used only if ad_ids parameter is null, and 'campaign_ids' parameter
// contains ID of only one campaign.
func (b *AdsGetAdsBuilder) Limit(v int) {
	b.Params["limit"] = v
}

// Offset Offset. Used in the same cases as 'limit' parameter.
func (b *AdsGetAdsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// AdsGetAdsLayoutBuilder builder
//
// Returns descriptions of ad layouts.
//
// https://vk.com/dev/ads.getAdsLayout
type AdsGetAdsLayoutBuilder struct {
	api.Params
}

// NewAdsGetAdsLayoutBuilder func
func NewAdsGetAdsLayoutBuilder() *AdsGetAdsLayoutBuilder {
	return &AdsGetAdsLayoutBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetAdsLayoutBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// AdIDs Filter by ads. Serialized JSON array with ad IDs. If the parameter is null, all ads will be shown.
func (b *AdsGetAdsLayoutBuilder) AdIDs(v string) {
	b.Params["ad_ids"] = v
}

// CampaignIDs Filter by advertising campaigns. Serialized JSON array with campaign IDs. If the parameter is null,
// ads of all campaigns will be shown.
func (b *AdsGetAdsLayoutBuilder) CampaignIDs(v string) {
	b.Params["campaign_ids"] = v
}

// ClientID 'For advertising agencies.' ID of the client ads are retrieved from.
func (b *AdsGetAdsLayoutBuilder) ClientID(v int) {
	b.Params["client_id"] = v
}

// IncludeDeleted Flag that specifies whether archived ads shall be shown.
//
// * 0 — show only active ads,
//
// * 1 — show all ads.
func (b *AdsGetAdsLayoutBuilder) IncludeDeleted(v bool) {
	b.Params["include_deleted"] = v
}

// Limit Limit of number of returned ads. Used only if 'ad_ids' parameter is null, and 'campaign_ids' parameter
// contains ID of only one campaign.
func (b *AdsGetAdsLayoutBuilder) Limit(v int) {
	b.Params["limit"] = v
}

// Offset Offset. Used in the same cases as 'limit' parameter.
func (b *AdsGetAdsLayoutBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// AdsGetAdsTargetingBuilder builder
//
// Returns ad targeting parameters.
//
// https://vk.com/dev/ads.getAdsTargeting
type AdsGetAdsTargetingBuilder struct {
	api.Params
}

// NewAdsGetAdsTargetingBuilder func
func NewAdsGetAdsTargetingBuilder() *AdsGetAdsTargetingBuilder {
	return &AdsGetAdsTargetingBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetAdsTargetingBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// AdIDs Filter by ads. Serialized JSON array with ad IDs. If the parameter is null, all ads will be shown.
func (b *AdsGetAdsTargetingBuilder) AdIDs(v string) {
	b.Params["ad_ids"] = v
}

// CampaignIDs Filter by advertising campaigns. Serialized JSON array with campaign IDs. If the parameter is null,
// ads of all campaigns will be shown.
func (b *AdsGetAdsTargetingBuilder) CampaignIDs(v string) {
	b.Params["campaign_ids"] = v
}

// ClientID 'For advertising agencies.' ID of the client ads are retrieved from.
func (b *AdsGetAdsTargetingBuilder) ClientID(v int) {
	b.Params["client_id"] = v
}

// IncludeDeleted flag that specifies whether archived ads shall be shown:
//
// * 0 — show only active ads,
//
// * 1 — show all ads.
func (b *AdsGetAdsTargetingBuilder) IncludeDeleted(v bool) {
	b.Params["include_deleted"] = v
}

// Limit Limit of number of returned ads. Used only if 'ad_ids' parameter is null, and 'campaign_ids' parameter
// contains ID of only one campaign.
func (b *AdsGetAdsTargetingBuilder) Limit(v int) {
	b.Params["limit"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *AdsGetAdsTargetingBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// AdsGetBudgetBuilder builder
//
// Returns current budget of the advertising account.
//
// https://vk.com/dev/ads.getBudget
type AdsGetBudgetBuilder struct {
	api.Params
}

// NewAdsGetBudgetBuilder func
func NewAdsGetBudgetBuilder() *AdsGetBudgetBuilder {
	return &AdsGetBudgetBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetBudgetBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// AdsGetCampaignsBuilder builder
//
// Returns a list of campaigns in an advertising account.
//
// https://vk.com/dev/ads.getCampaigns
type AdsGetCampaignsBuilder struct {
	api.Params
}

// NewAdsGetCampaignsBuilder func
func NewAdsGetCampaignsBuilder() *AdsGetCampaignsBuilder {
	return &AdsGetCampaignsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetCampaignsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// ClientID 'For advertising agencies'. ID of the client advertising campaigns are retrieved from.
func (b *AdsGetCampaignsBuilder) ClientID(v int) {
	b.Params["client_id"] = v
}

// IncludeDeleted Flag that specifies whether archived ads shall be shown.
//
// * 0 — show only active campaigns,
//
// * 1 — show all campaigns.
func (b *AdsGetCampaignsBuilder) IncludeDeleted(v bool) {
	b.Params["include_deleted"] = v
}

// CampaignIDs Filter of advertising campaigns to show. Serialized JSON array with campaign IDs. Only campaigns that
// exist in 'campaign_ids' and belong to the specified advertising account will be shown. If the parameter is null,
// all campaigns will be shown.
func (b *AdsGetCampaignsBuilder) CampaignIDs(v string) {
	b.Params["campaign_ids"] = v
}

// AdsGetCategoriesBuilder builder
//
// Returns a list of possible ad categories.
//
// https://vk.com/dev/ads.getCategories
type AdsGetCategoriesBuilder struct {
	api.Params
}

// NewAdsGetCategoriesBuilder func
func NewAdsGetCategoriesBuilder() *AdsGetCategoriesBuilder {
	return &AdsGetCategoriesBuilder{api.Params{}}
}

// Lang Language. The full list of supported languages is [vk.com/dev/api_requests|here].
func (b *AdsGetCategoriesBuilder) Lang(v string) {
	b.Params["lang"] = v
}

// AdsGetClientsBuilder builder
//
// Returns a list of advertising agency's clients.
//
// https://vk.com/dev/ads.getClients
type AdsGetClientsBuilder struct {
	api.Params
}

// NewAdsGetClientsBuilder func
func NewAdsGetClientsBuilder() *AdsGetClientsBuilder {
	return &AdsGetClientsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetClientsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// AdsGetDemographicsBuilder builder
//
// Returns demographics for ads or campaigns.
//
// https://vk.com/dev/ads.getDemographics
type AdsGetDemographicsBuilder struct {
	api.Params
}

// NewAdsGetDemographicsBuilder func
func NewAdsGetDemographicsBuilder() *AdsGetDemographicsBuilder {
	return &AdsGetDemographicsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetDemographicsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// IDsType Type of requested objects listed in 'ids' parameter:
//
// * ad — ads,
//
// * campaign — campaigns.
func (b *AdsGetDemographicsBuilder) IDsType(v string) {
	b.Params["ids_type"] = v
}

// IDs IDs requested ads or campaigns, separated with a comma, depending on the value set in 'ids_type'.
// Maximum 2000 objects.
func (b *AdsGetDemographicsBuilder) IDs(v string) {
	b.Params["ids"] = v
}

// Period Data grouping by dates:
//
// * day — statistics by days,
//
// * month — statistics by months,
//
// * overall — overall statistics. 'date_from' and 'date_to' parameters set temporary limits.
func (b *AdsGetDemographicsBuilder) Period(v string) {
	b.Params["period"] = v
}

// DateFrom Date to show statistics from. For different value of 'period' different date format is used:
//
// * day: YYYY-MM-DD, example: 2011-09-27 — September 27, 2011
//
// * *0 — day it was created on,
//
// * month: YYYY-MM, example: 2011-09 — September 2011
//
// * *0 — month it was created in,
//
// * overall: 0.
func (b *AdsGetDemographicsBuilder) DateFrom(v string) {
	b.Params["date_from"] = v
}

// DateTo Date to show statistics to. For different value of 'period' different date format is used:
//
// * day: YYYY-MM-DD, example: 2011-09-27 — September 27, 2011
//
// * *0 — current day,
//
// * month: YYYY-MM, example: 2011-09 — September 2011
//
// * *0 — current month,
//
// * overall: 0.
func (b *AdsGetDemographicsBuilder) DateTo(v string) {
	b.Params["date_to"] = v
}

// AdsGetFloodStatsBuilder builder
//
// Returns information about current state of a counter — number of remaining runs of methods and time to the next
// counter nulling in seconds.
//
// https://vk.com/dev/ads.getFloodStats
type AdsGetFloodStatsBuilder struct {
	api.Params
}

// NewAdsGetFloodStatsBuilder func
func NewAdsGetFloodStatsBuilder() *AdsGetFloodStatsBuilder {
	return &AdsGetFloodStatsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetFloodStatsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// AdsGetOfficeUsersBuilder builder
//
// Returns a list of managers and supervisors of advertising account.
//
// https://vk.com/dev/ads.getOfficeUsers
type AdsGetOfficeUsersBuilder struct {
	api.Params
}

// NewAdsGetOfficeUsersBuilder func
func NewAdsGetOfficeUsersBuilder() *AdsGetOfficeUsersBuilder {
	return &AdsGetOfficeUsersBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetOfficeUsersBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// AdsGetPostsReachBuilder builder
//
// Returns detailed statistics of promoted posts reach from campaigns and ads.
//
// https://vk.com/dev/ads.getPostsReach
type AdsGetPostsReachBuilder struct {
	api.Params
}

// NewAdsGetPostsReachBuilder func
func NewAdsGetPostsReachBuilder() *AdsGetPostsReachBuilder {
	return &AdsGetPostsReachBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetPostsReachBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// IDsType Type of requested objects listed in 'ids' parameter:
//
// * ad — ads,
//
// * campaign — campaigns.
func (b *AdsGetPostsReachBuilder) IDsType(v string) {
	b.Params["ids_type"] = v
}

// IDs IDs requested ads or campaigns, separated with a comma, depending on the value set in 'ids_type'.
// Maximum 100 objects.
func (b *AdsGetPostsReachBuilder) IDs(v string) {
	b.Params["ids"] = v
}

// AdsGetRejectionReasonBuilder builder
//
// Returns a reason of ad rejection for pre-moderation.
//
// https://vk.com/dev/ads.getRejectionReason
type AdsGetRejectionReasonBuilder struct {
	api.Params
}

// NewAdsGetRejectionReasonBuilder func
func NewAdsGetRejectionReasonBuilder() *AdsGetRejectionReasonBuilder {
	return &AdsGetRejectionReasonBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetRejectionReasonBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// AdID Ad ID.
func (b *AdsGetRejectionReasonBuilder) AdID(v int) {
	b.Params["ad_id"] = v
}

// AdsGetStatisticsBuilder builder
//
// Returns statistics of performance indicators for ads, campaigns, clients or the whole account.
//
// https://vk.com/dev/ads.getStatistics
type AdsGetStatisticsBuilder struct {
	api.Params
}

// NewAdsGetStatisticsBuilder func
func NewAdsGetStatisticsBuilder() *AdsGetStatisticsBuilder {
	return &AdsGetStatisticsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetStatisticsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// IDsType Type of requested objects listed in 'ids' parameter:
//
// * ad — ads,
//
// * campaign — campaigns,
//
// * client — clients,
//
// * office — account.
func (b *AdsGetStatisticsBuilder) IDsType(v string) {
	b.Params["ids_type"] = v
}

// IDs IDs requested ads, campaigns, clients or account, separated with a comma, depending on the value set in
// 'ids_type'. Maximum 2000 objects.
func (b *AdsGetStatisticsBuilder) IDs(v string) {
	b.Params["ids"] = v
}

// Period Data grouping by dates:
//
// * day — statistics by days,
//
// * month — statistics by months,
//
// * overall — overall statistics. 'date_from' and 'date_to' parameters set temporary limits.
func (b *AdsGetStatisticsBuilder) Period(v string) {
	b.Params["period"] = v
}

// DateFrom Date to show statistics from. For different value of 'period' different date format is used:
//
// * day: YYYY-MM-DD, example: 2011-09-27 — September 27, 2011
//
// * *0 — day it was created on,
//
// * month: YYYY-MM, example: 2011-09 — September 2011
//
// * *0 — month it was created in,
//
// * overall: 0.
func (b *AdsGetStatisticsBuilder) DateFrom(v string) {
	b.Params["date_from"] = v
}

// DateTo Date to show statistics to. For different value of 'period' different date format is used:
//
// * day: YYYY-MM-DD, example: 2011-09-27 — September 27, 2011
//
// * *0 — current day,
//
// * month: YYYY-MM, example: 2011-09 — September 2011
//
// * *0 — current month,
//
// * overall: 0.
func (b *AdsGetStatisticsBuilder) DateTo(v string) {
	b.Params["date_to"] = v
}

// AdsGetSuggestionsBuilder builder
//
// Returns a set of auto-suggestions for various targeting parameters.
//
// https://vk.com/dev/ads.getSuggestions
type AdsGetSuggestionsBuilder struct {
	api.Params
}

// NewAdsGetSuggestionsBuilder func
func NewAdsGetSuggestionsBuilder() *AdsGetSuggestionsBuilder {
	return &AdsGetSuggestionsBuilder{api.Params{}}
}

// Section Section, suggestions are retrieved in. Available values:
//
// * countries — request of a list of countries. If q is not set or blank, a short list of countries is shown.
// Otherwise, a full list of countries is shown.
//
// * regions — requested list of regions. 'country' parameter is required.
//
// * cities — requested list of cities. 'country' parameter is required.
//
// * districts — requested list of districts. 'cities' parameter is required.
//
// * stations — requested list of subway stations. 'cities' parameter is required.
//
// * streets — requested list of streets. 'cities' parameter is required.
//
// * schools — requested list of educational organizations. 'cities' parameter is required.
//
// * interests — requested list of interests.
//
// * positions — requested list of positions (professions).
//
// * group_types — requested list of group types.
//
// * religions — requested list of religious commitments.
//
// * browsers — requested list of browsers and mobile devices.
func (b *AdsGetSuggestionsBuilder) Section(v string) {
	b.Params["section"] = v
}

// IDs Objects IDs separated by commas. If the parameter is passed, 'q, country, cities' should not be passed.
func (b *AdsGetSuggestionsBuilder) IDs(v string) {
	b.Params["ids"] = v
}

// Q Filter-line of the request (for countries, regions, cities, streets, schools, interests, positions).
func (b *AdsGetSuggestionsBuilder) Q(v string) {
	b.Params["q"] = v
}

// Country ID of the country objects are searched in.
func (b *AdsGetSuggestionsBuilder) Country(v int) {
	b.Params["country"] = v
}

// Cities IDs of cities where objects are searched in, separated with a comma.
func (b *AdsGetSuggestionsBuilder) Cities(v string) {
	b.Params["cities"] = v
}

// Lang Language of the returned string values. Supported languages:
//
// * ru — Russian,
//
// * ua — Ukrainian,
//
// * en — English.
func (b *AdsGetSuggestionsBuilder) Lang(v string) {
	b.Params["lang"] = v
}

// AdsGetTargetGroupsBuilder builder
//
// Returns a list of target groups.
//
// https://vk.com/dev/ads.getTargetGroups
type AdsGetTargetGroupsBuilder struct {
	api.Params
}

// NewAdsGetTargetGroupsBuilder func
func NewAdsGetTargetGroupsBuilder() *AdsGetTargetGroupsBuilder {
	return &AdsGetTargetGroupsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetTargetGroupsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// ClientID Only for advertising agencies.
// ID of the client with the advertising account where the group will be created.
func (b *AdsGetTargetGroupsBuilder) ClientID(v int) {
	b.Params["client_id"] = v
}

// Extended parameter
//
// * 1 — to return pixel code.
func (b *AdsGetTargetGroupsBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// AdsGetTargetingStatsBuilder builder
//
// Returns the size of targeting audience, and also recommended values for CPC and CPM.
//
// https://vk.com/dev/ads.getTargetingStats
type AdsGetTargetingStatsBuilder struct {
	api.Params
}

// NewAdsGetTargetingStatsBuilder func
func NewAdsGetTargetingStatsBuilder() *AdsGetTargetingStatsBuilder {
	return &AdsGetTargetingStatsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsGetTargetingStatsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// ClientID parameter
func (b *AdsGetTargetingStatsBuilder) ClientID(v int) {
	b.Params["client_id"] = v
}

// Criteria Serialized JSON object that describes targeting parameters.
// Description of 'criteria' object see below.
func (b *AdsGetTargetingStatsBuilder) Criteria(v string) {
	b.Params["criteria"] = v
}

// AdID ID of an ad which targeting parameters shall be analyzed.
func (b *AdsGetTargetingStatsBuilder) AdID(v int) {
	b.Params["ad_id"] = v
}

// AdFormat Ad format. Possible values:
//
// * 1 — image and text,
//
// * 2 — big image,
//
// * 3 — exclusive format,
//
// * 4 — community, square image,
//
// * 7 — special app format,
//
// * 8 — special community format,
//
// * 9 — post in community,
//
// * 10 — app board.
func (b *AdsGetTargetingStatsBuilder) AdFormat(v int) {
	b.Params["ad_format"] = v
}

// AdPlatform Platforms to use for ad showing. Possible values: (for 'ad_format' = '1')
//
// * 0 — VK and partner sites,
//
// * 1 — VK only. (for 'ad_format' = '9')
//
// * all — all platforms,
//
// * desktop — desktop version,
//
// * mobile — mobile version and apps.
func (b *AdsGetTargetingStatsBuilder) AdPlatform(v string) {
	b.Params["ad_platform"] = v
}

// AdPlatformNoWall parameter
func (b *AdsGetTargetingStatsBuilder) AdPlatformNoWall(v string) {
	b.Params["ad_platform_no_wall"] = v
}

// AdPlatformNoAdNetwork parameter
func (b *AdsGetTargetingStatsBuilder) AdPlatformNoAdNetwork(v string) {
	b.Params["ad_platform_no_ad_network"] = v
}

// LinkURL URL for the advertised object.
func (b *AdsGetTargetingStatsBuilder) LinkURL(v string) {
	b.Params["link_url"] = v
}

// LinkDomain Domain of the advertised object.
func (b *AdsGetTargetingStatsBuilder) LinkDomain(v string) {
	b.Params["link_domain"] = v
}

// AdsGetUploadURLBuilder builder
//
// Returns URL to upload an ad photo to.
//
// https://vk.com/dev/ads.getUploadURL
type AdsGetUploadURLBuilder struct {
	api.Params
}

// NewAdsGetUploadURLBuilder func
func NewAdsGetUploadURLBuilder() *AdsGetUploadURLBuilder {
	return &AdsGetUploadURLBuilder{api.Params{}}
}

// AdFormat Ad format:
//
// * 1 — image and text,
//
// * 2 — big image,
//
// * 3 — exclusive format,
//
// * 4 — community, square image,
//
// * 7 — special app format.
func (b *AdsGetUploadURLBuilder) AdFormat(v int) {
	b.Params["ad_format"] = v
}

// Icon parameter
func (b *AdsGetUploadURLBuilder) Icon(v int) {
	b.Params["icon"] = v
}

// AdsImportTargetContactsBuilder builder
//
// Imports a list of advertiser's contacts to count VK registered users against the target group.
//
// https://vk.com/dev/ads.importTargetContacts
type AdsImportTargetContactsBuilder struct {
	api.Params
}

// NewAdsImportTargetContactsBuilder func
func NewAdsImportTargetContactsBuilder() *AdsImportTargetContactsBuilder {
	return &AdsImportTargetContactsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsImportTargetContactsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// ClientID Only for advertising agencies.
// ID of the client with the advertising account where the group will be created.
func (b *AdsImportTargetContactsBuilder) ClientID(v int) {
	b.Params["client_id"] = v
}

// TargetGroupID Target group ID.
func (b *AdsImportTargetContactsBuilder) TargetGroupID(v int) {
	b.Params["target_group_id"] = v
}

// Contacts List of phone numbers, emails or user IDs separated with a comma.
func (b *AdsImportTargetContactsBuilder) Contacts(v string) {
	b.Params["contacts"] = v
}

// AdsRemoveOfficeUsersBuilder builder
//
// Removes managers and/or supervisors from advertising account.
//
// https://vk.com/dev/ads.removeOfficeUsers
type AdsRemoveOfficeUsersBuilder struct {
	api.Params
}

// NewAdsRemoveOfficeUsersBuilder func
func NewAdsRemoveOfficeUsersBuilder() *AdsRemoveOfficeUsersBuilder {
	return &AdsRemoveOfficeUsersBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsRemoveOfficeUsersBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// IDs Serialized JSON array with IDs of deleted managers.
func (b *AdsRemoveOfficeUsersBuilder) IDs(v string) {
	b.Params["ids"] = v
}

// AdsUpdateAdsBuilder builder
//
// Edits ads.
//
// https://vk.com/dev/ads.updateAds
type AdsUpdateAdsBuilder struct {
	api.Params
}

// NewAdsUpdateAdsBuilder func
func NewAdsUpdateAdsBuilder() *AdsUpdateAdsBuilder {
	return &AdsUpdateAdsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsUpdateAdsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// Data Serialized JSON array of objects that describe changes in ads.
// Description of 'ad_edit_specification' objects see below.
func (b *AdsUpdateAdsBuilder) Data(v string) {
	b.Params["data"] = v
}

// AdsUpdateCampaignsBuilder builder
//
// Edits advertising campaigns.
//
// https://vk.com/dev/ads.updateCampaigns
type AdsUpdateCampaignsBuilder struct {
	api.Params
}

// NewAdsUpdateCampaignsBuilder func
func NewAdsUpdateCampaignsBuilder() *AdsUpdateCampaignsBuilder {
	return &AdsUpdateCampaignsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsUpdateCampaignsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// Data Serialized JSON array of objects that describe changes in campaigns.
// Description of 'campaign_mod' objects see below.
func (b *AdsUpdateCampaignsBuilder) Data(v string) {
	b.Params["data"] = v
}

// AdsUpdateClientsBuilder builder
//
// Edits clients of an advertising agency.
//
// https://vk.com/dev/ads.updateClients
type AdsUpdateClientsBuilder struct {
	api.Params
}

// NewAdsUpdateClientsBuilder func
func NewAdsUpdateClientsBuilder() *AdsUpdateClientsBuilder {
	return &AdsUpdateClientsBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsUpdateClientsBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// Data Serialized JSON array of objects that describe changes in clients.
// Description of 'client_mod' objects see below.
func (b *AdsUpdateClientsBuilder) Data(v string) {
	b.Params["data"] = v
}

// AdsUpdateTargetGroupBuilder builder
//
// Edits a retarget group.
//
// https://vk.com/dev/ads.updateTargetGroup
type AdsUpdateTargetGroupBuilder struct {
	api.Params
}

// NewAdsUpdateTargetGroupBuilder func
func NewAdsUpdateTargetGroupBuilder() *AdsUpdateTargetGroupBuilder {
	return &AdsUpdateTargetGroupBuilder{api.Params{}}
}

// AccountID Advertising account ID.
func (b *AdsUpdateTargetGroupBuilder) AccountID(v int) {
	b.Params["account_id"] = v
}

// ClientID Only for advertising agencies.
// ID of the client with the advertising account where the group will
// be created.
func (b *AdsUpdateTargetGroupBuilder) ClientID(v int) {
	b.Params["client_id"] = v
}

// TargetGroupID Group ID.
func (b *AdsUpdateTargetGroupBuilder) TargetGroupID(v int) {
	b.Params["target_group_id"] = v
}

// Name New name of the target group — a string up to 64 characters long.
func (b *AdsUpdateTargetGroupBuilder) Name(v string) {
	b.Params["name"] = v
}

// Domain Domain of the site where user accounting code will be placed.
func (b *AdsUpdateTargetGroupBuilder) Domain(v string) {
	b.Params["domain"] = v
}

// Lifetime Only for the groups that get audience from sites with user accounting code.
// Time in days when users added to a retarget group will be automatically excluded from it.
//
// * 0 – automatic exclusion is off.
func (b *AdsUpdateTargetGroupBuilder) Lifetime(v int) {
	b.Params["lifetime"] = v
}

// TargetPixelID parameter
func (b *AdsUpdateTargetGroupBuilder) TargetPixelID(v int) {
	b.Params["target_pixel_id"] = v
}

// TargetPixelRules parameter
func (b *AdsUpdateTargetGroupBuilder) TargetPixelRules(v string) {
	b.Params["target_pixel_rules"] = v
}
