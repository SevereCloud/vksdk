package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestAdsAddOfficeUsersBulder(t *testing.T) {
	b := params.NewAdsAddOfficeUsersBulder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["data"], "text")
}

func TestAdsCheckLinkBulder(t *testing.T) {
	b := params.NewAdsCheckLinkBulder()

	b.AccountID(1)
	b.LinkType("text")
	b.LinkURL("text")
	b.CampaignID(1)

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["link_type"], "text")
	assert.Equal(t, b.Params["link_url"], "text")
	assert.Equal(t, b.Params["campaign_id"], 1)
}

func TestAdsCreateAdsBulder(t *testing.T) {
	b := params.NewAdsCreateAdsBulder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["data"], "text")
}

func TestAdsCreateCampaignsBulder(t *testing.T) {
	b := params.NewAdsCreateCampaignsBulder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["data"], "text")
}

func TestAdsCreateClientsBulder(t *testing.T) {
	b := params.NewAdsCreateClientsBulder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["data"], "text")
}

func TestAdsCreateTargetGroupBulder(t *testing.T) {
	b := params.NewAdsCreateTargetGroupBulder()

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

func TestAdsDeleteAdsBulder(t *testing.T) {
	b := params.NewAdsDeleteAdsBulder()

	b.AccountID(1)
	b.IDs("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ids"], "text")
}

func TestAdsDeleteCampaignsBulder(t *testing.T) {
	b := params.NewAdsDeleteCampaignsBulder()

	b.AccountID(1)
	b.IDs("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ids"], "text")
}

func TestAdsDeleteClientsBulder(t *testing.T) {
	b := params.NewAdsDeleteClientsBulder()

	b.AccountID(1)
	b.IDs("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ids"], "text")
}

func TestAdsDeleteTargetGroupBulder(t *testing.T) {
	b := params.NewAdsDeleteTargetGroupBulder()

	b.AccountID(1)
	b.ClientID(1)
	b.TargetGroupID(1)

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["client_id"], 1)
	assert.Equal(t, b.Params["target_group_id"], 1)
}

func TestAdsGetAdsBulder(t *testing.T) {
	b := params.NewAdsGetAdsBulder()

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

func TestAdsGetAdsLayoutBulder(t *testing.T) {
	b := params.NewAdsGetAdsLayoutBulder()

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

func TestAdsGetAdsTargetingBulder(t *testing.T) {
	b := params.NewAdsGetAdsTargetingBulder()

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

func TestAdsGetBudgetBulder(t *testing.T) {
	b := params.NewAdsGetBudgetBulder()

	b.AccountID(1)

	assert.Equal(t, b.Params["account_id"], 1)
}

func TestAdsGetCampaignsBulder(t *testing.T) {
	b := params.NewAdsGetCampaignsBulder()

	b.AccountID(1)
	b.ClientID(1)
	b.IncludeDeleted(true)
	b.CampaignIDs("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["client_id"], 1)
	assert.Equal(t, b.Params["include_deleted"], true)
	assert.Equal(t, b.Params["campaign_ids"], "text")
}

func TestAdsGetCategoriesBulder(t *testing.T) {
	b := params.NewAdsGetCategoriesBulder()

	b.Lang("text")

	assert.Equal(t, b.Params["lang"], "text")
}

func TestAdsGetClientsBulder(t *testing.T) {
	b := params.NewAdsGetClientsBulder()

	b.AccountID(1)

	assert.Equal(t, b.Params["account_id"], 1)
}

func TestAdsGetDemographicsBulder(t *testing.T) {
	b := params.NewAdsGetDemographicsBulder()

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

func TestAdsGetFloodStatsBulder(t *testing.T) {
	b := params.NewAdsGetFloodStatsBulder()

	b.AccountID(1)

	assert.Equal(t, b.Params["account_id"], 1)
}

func TestAdsGetOfficeUsersBulder(t *testing.T) {
	b := params.NewAdsGetOfficeUsersBulder()

	b.AccountID(1)

	assert.Equal(t, b.Params["account_id"], 1)
}

func TestAdsGetPostsReachBulder(t *testing.T) {
	b := params.NewAdsGetPostsReachBulder()

	b.AccountID(1)
	b.IDsType("text")
	b.IDs("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ids_type"], "text")
	assert.Equal(t, b.Params["ids"], "text")
}

func TestAdsGetRejectionReasonBulder(t *testing.T) {
	b := params.NewAdsGetRejectionReasonBulder()

	b.AccountID(1)
	b.AdID(1)

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ad_id"], 1)
}

func TestAdsGetStatisticsBulder(t *testing.T) {
	b := params.NewAdsGetStatisticsBulder()

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

func TestAdsGetSuggestionsBulder(t *testing.T) {
	b := params.NewAdsGetSuggestionsBulder()

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

func TestAdsGetTargetGroupsBulder(t *testing.T) {
	b := params.NewAdsGetTargetGroupsBulder()

	b.AccountID(1)
	b.ClientID(1)
	b.Extended(true)

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["client_id"], 1)
	assert.Equal(t, b.Params["extended"], true)
}

func TestAdsGetTargetingStatsBulder(t *testing.T) {
	b := params.NewAdsGetTargetingStatsBulder()

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

func TestAdsGetUploadURLBulder(t *testing.T) {
	b := params.NewAdsGetUploadURLBulder()

	b.AdFormat(1)
	b.Icon(1)

	assert.Equal(t, b.Params["ad_format"], 1)
	assert.Equal(t, b.Params["icon"], 1)
}

func TestAdsImportTargetContactsBulder(t *testing.T) {
	b := params.NewAdsImportTargetContactsBulder()

	b.AccountID(1)
	b.ClientID(1)
	b.TargetGroupID(1)
	b.Contacts("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["client_id"], 1)
	assert.Equal(t, b.Params["target_group_id"], 1)
	assert.Equal(t, b.Params["contacts"], "text")
}

func TestAdsRemoveOfficeUsersBulder(t *testing.T) {
	b := params.NewAdsRemoveOfficeUsersBulder()

	b.AccountID(1)
	b.IDs("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["ids"], "text")
}

func TestAdsUpdateAdsBulder(t *testing.T) {
	b := params.NewAdsUpdateAdsBulder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["data"], "text")
}

func TestAdsUpdateCampaignsBulder(t *testing.T) {
	b := params.NewAdsUpdateCampaignsBulder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["data"], "text")
}

func TestAdsUpdateClientsBulder(t *testing.T) {
	b := params.NewAdsUpdateClientsBulder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, b.Params["account_id"], 1)
	assert.Equal(t, b.Params["data"], "text")
}

func TestAdsUpdateTargetGroupBulder(t *testing.T) {
	b := params.NewAdsUpdateTargetGroupBulder()

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
