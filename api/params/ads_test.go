package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestAdsAddOfficeUsersBuilder(t *testing.T) {
	b := params.NewAdsAddOfficeUsersBuilder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["data"], "text")
}

func TestAdsCheckLinkBuilder(t *testing.T) {
	b := params.NewAdsCheckLinkBuilder()

	b.AccountID(1)
	b.LinkType("text")
	b.LinkURL("text")
	b.CampaignID(1)

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["link_type"], "text")
	assert.Equal(t, b.Params["link_url"], "text")
	assert.Equal(t, b.Params["campaign_id"], 1)
}

func TestAdsCreateAdsBuilder(t *testing.T) {
	b := params.NewAdsCreateAdsBuilder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["data"], "text")
}

func TestAdsCreateCampaignsBuilder(t *testing.T) {
	b := params.NewAdsCreateCampaignsBuilder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["data"], "text")
}

func TestAdsCreateClientsBuilder(t *testing.T) {
	b := params.NewAdsCreateClientsBuilder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["data"], "text")
}

func TestAdsCreateTargetGroupBuilder(t *testing.T) {
	b := params.NewAdsCreateTargetGroupBuilder()

	b.AccountID(1)
	b.ClientID(1)
	b.Name("text")
	b.Lifetime(1)
	b.TargetPixelID(1)
	b.TargetPixelRules("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["client_id"], 1)
	assert.Equal(t, b.Params["name"], "text")
	assert.Equal(t, b.Params["lifetime"], 1)
	assert.Equal(t, b.Params["target_pixel_id"], 1)
	assert.Equal(t, b.Params["target_pixel_rules"], "text")
}

func TestAdsDeleteAdsBuilder(t *testing.T) {
	b := params.NewAdsDeleteAdsBuilder()

	b.AccountID(1)
	b.IDs("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ids"], "text")
}

func TestAdsDeleteCampaignsBuilder(t *testing.T) {
	b := params.NewAdsDeleteCampaignsBuilder()

	b.AccountID(1)
	b.IDs("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ids"], "text")
}

func TestAdsDeleteClientsBuilder(t *testing.T) {
	b := params.NewAdsDeleteClientsBuilder()

	b.AccountID(1)
	b.IDs("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ids"], "text")
}

func TestAdsDeleteTargetGroupBuilder(t *testing.T) {
	b := params.NewAdsDeleteTargetGroupBuilder()

	b.AccountID(1)
	b.ClientID(1)
	b.TargetGroupID(1)

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["client_id"], 1)
	assert.Equal(t, b.Params["target_group_id"], 1)
}

func TestAdsGetAdsBuilder(t *testing.T) {
	b := params.NewAdsGetAdsBuilder()

	b.AccountID(1)
	b.AdIDs("text")
	b.CampaignIDs("text")
	b.ClientID(1)
	b.IncludeDeleted(true)
	b.Limit(1)
	b.Offset(1)

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ad_ids"], "text")
	assert.Equal(t, b.Params["campaign_ids"], "text")
	assert.Equal(t, b.Params["client_id"], 1)
	assert.Equal(t, b.Params["include_deleted"], true)
	assert.Equal(t, b.Params["limit"], 1)
	assert.Equal(t, b.Params["offset"], 1)
}

func TestAdsGetAdsLayoutBuilder(t *testing.T) {
	b := params.NewAdsGetAdsLayoutBuilder()

	b.AccountID(1)
	b.AdIDs("text")
	b.CampaignIDs("text")
	b.ClientID(1)
	b.IncludeDeleted(true)
	b.Limit(1)
	b.Offset(1)

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ad_ids"], "text")
	assert.Equal(t, b.Params["campaign_ids"], "text")
	assert.Equal(t, b.Params["client_id"], 1)
	assert.Equal(t, b.Params["include_deleted"], true)
	assert.Equal(t, b.Params["limit"], 1)
	assert.Equal(t, b.Params["offset"], 1)
}

func TestAdsGetAdsTargetingBuilder(t *testing.T) {
	b := params.NewAdsGetAdsTargetingBuilder()

	b.AccountID(1)
	b.AdIDs("text")
	b.CampaignIDs("text")
	b.ClientID(1)
	b.IncludeDeleted(true)
	b.Limit(1)
	b.Offset(1)

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ad_ids"], "text")
	assert.Equal(t, b.Params["campaign_ids"], "text")
	assert.Equal(t, b.Params["client_id"], 1)
	assert.Equal(t, b.Params["include_deleted"], true)
	assert.Equal(t, b.Params["limit"], 1)
	assert.Equal(t, b.Params["offset"], 1)
}

func TestAdsGetBudgetBuilder(t *testing.T) {
	b := params.NewAdsGetBudgetBuilder()

	b.AccountID(1)

	assert.Equal(t, b.Params["account_id"], 1)
}

func TestAdsGetCampaignsBuilder(t *testing.T) {
	b := params.NewAdsGetCampaignsBuilder()

	b.AccountID(1)
	b.ClientID(1)
	b.IncludeDeleted(true)
	b.CampaignIDs("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["client_id"], 1)
	assert.Equal(t, b.Params["include_deleted"], true)
	assert.Equal(t, b.Params["campaign_ids"], "text")
}

func TestAdsGetCategoriesBuilder(t *testing.T) {
	b := params.NewAdsGetCategoriesBuilder()

	b.Lang("text")

	assert.Equal(t, b.Params["lang"], "text")
}

func TestAdsGetClientsBuilder(t *testing.T) {
	b := params.NewAdsGetClientsBuilder()

	b.AccountID(1)

	assert.Equal(t, b.Params["account_id"], 1)
}

func TestAdsGetDemographicsBuilder(t *testing.T) {
	b := params.NewAdsGetDemographicsBuilder()

	b.AccountID(1)
	b.IDsType("text")
	b.IDs("text")
	b.Period("text")
	b.DateFrom("text")
	b.DateTo("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ids_type"], "text")
	assert.Equal(t, b.Params["ids"], "text")
	assert.Equal(t, b.Params["period"], "text")
	assert.Equal(t, b.Params["date_from"], "text")
	assert.Equal(t, b.Params["date_to"], "text")
}

func TestAdsGetFloodStatsBuilder(t *testing.T) {
	b := params.NewAdsGetFloodStatsBuilder()

	b.AccountID(1)

	assert.Equal(t, b.Params["account_id"], 1)
}

func TestAdsGetOfficeUsersBuilder(t *testing.T) {
	b := params.NewAdsGetOfficeUsersBuilder()

	b.AccountID(1)

	assert.Equal(t, b.Params["account_id"], 1)
}

func TestAdsGetPostsReachBuilder(t *testing.T) {
	b := params.NewAdsGetPostsReachBuilder()

	b.AccountID(1)
	b.IDsType("text")
	b.IDs("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ids_type"], "text")
	assert.Equal(t, b.Params["ids"], "text")
}

func TestAdsGetRejectionReasonBuilder(t *testing.T) {
	b := params.NewAdsGetRejectionReasonBuilder()

	b.AccountID(1)
	b.AdID(1)

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ad_id"], 1)
}

func TestAdsGetStatisticsBuilder(t *testing.T) {
	b := params.NewAdsGetStatisticsBuilder()

	b.AccountID(1)
	b.IDsType("text")
	b.IDs("text")
	b.Period("text")
	b.DateFrom("text")
	b.DateTo("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ids_type"], "text")
	assert.Equal(t, b.Params["ids"], "text")
	assert.Equal(t, b.Params["period"], "text")
	assert.Equal(t, b.Params["date_from"], "text")
	assert.Equal(t, b.Params["date_to"], "text")
}

func TestAdsGetSuggestionsBuilder(t *testing.T) {
	b := params.NewAdsGetSuggestionsBuilder()

	b.Section("text")
	b.IDs("text")
	b.Q("text")
	b.Country(1)
	b.Cities("text")
	b.Lang("text")

	assert.Equal(t, b.Params["section"], "text")
	assert.Equal(t, b.Params["ids"], "text")
	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["country"], 1)
	assert.Equal(t, b.Params["cities"], "text")
	assert.Equal(t, b.Params["lang"], "text")
}

func TestAdsGetTargetGroupsBuilder(t *testing.T) {
	b := params.NewAdsGetTargetGroupsBuilder()

	b.AccountID(1)
	b.ClientID(1)
	b.Extended(true)

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["client_id"], 1)
	assert.Equal(t, b.Params["extended"], true)
}

func TestAdsGetTargetingStatsBuilder(t *testing.T) {
	b := params.NewAdsGetTargetingStatsBuilder()

	b.AccountID(1)
	b.ClientID(1)
	b.Criteria("text")
	b.AdID(1)
	b.AdFormat(1)
	b.AdPlatform("text")
	b.AdPlatformNoWall("text")
	b.AdPlatformNoAdNetwork("text")
	b.LinkURL("text")
	b.LinkDomain("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["client_id"], 1)
	assert.Equal(t, b.Params["criteria"], "text")
	assert.Equal(t, b.Params["ad_id"], 1)
	assert.Equal(t, b.Params["ad_format"], 1)
	assert.Equal(t, b.Params["ad_platform"], "text")
	assert.Equal(t, b.Params["ad_platform_no_wall"], "text")
	assert.Equal(t, b.Params["ad_platform_no_ad_network"], "text")
	assert.Equal(t, b.Params["link_url"], "text")
	assert.Equal(t, b.Params["link_domain"], "text")
}

func TestAdsGetUploadURLBuilder(t *testing.T) {
	b := params.NewAdsGetUploadURLBuilder()

	b.AdFormat(1)
	b.Icon(1)

	assert.Equal(t, b.Params["ad_format"], 1)
	assert.Equal(t, b.Params["icon"], 1)
}

func TestAdsImportTargetContactsBuilder(t *testing.T) {
	b := params.NewAdsImportTargetContactsBuilder()

	b.AccountID(1)
	b.ClientID(1)
	b.TargetGroupID(1)
	b.Contacts("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["client_id"], 1)
	assert.Equal(t, b.Params["target_group_id"], 1)
	assert.Equal(t, b.Params["contacts"], "text")
}

func TestAdsRemoveOfficeUsersBuilder(t *testing.T) {
	b := params.NewAdsRemoveOfficeUsersBuilder()

	b.AccountID(1)
	b.IDs("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ids"], "text")
}

func TestAdsUpdateAdsBuilder(t *testing.T) {
	b := params.NewAdsUpdateAdsBuilder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["data"], "text")
}

func TestAdsUpdateCampaignsBuilder(t *testing.T) {
	b := params.NewAdsUpdateCampaignsBuilder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["data"], "text")
}

func TestAdsUpdateClientsBuilder(t *testing.T) {
	b := params.NewAdsUpdateClientsBuilder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["data"], "text")
}

func TestAdsUpdateTargetGroupBuilder(t *testing.T) {
	b := params.NewAdsUpdateTargetGroupBuilder()

	b.AccountID(1)
	b.ClientID(1)
	b.TargetGroupID(1)
	b.Name("text")
	b.Domain("text")
	b.Lifetime(1)
	b.TargetPixelID(1)
	b.TargetPixelRules("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["client_id"], 1)
	assert.Equal(t, b.Params["target_group_id"], 1)
	assert.Equal(t, b.Params["name"], "text")
	assert.Equal(t, b.Params["domain"], "text")
	assert.Equal(t, b.Params["lifetime"], 1)
	assert.Equal(t, b.Params["target_pixel_id"], 1)
	assert.Equal(t, b.Params["target_pixel_rules"], "text")
}
