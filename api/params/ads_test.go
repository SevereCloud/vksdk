package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestAdsAddOfficeUsersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsAddOfficeUsersBuilder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["data"])
}

func TestAdsCheckLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsCheckLinkBuilder()

	b.AccountID(1)
	b.LinkType("text")
	b.LinkURL("text")
	b.CampaignID(1)

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["link_type"])
	assert.Equal(t, "text", b.Params["link_url"])
	assert.Equal(t, 1, b.Params["campaign_id"])
}

func TestAdsCreateAdsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsCreateAdsBuilder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["data"])
}

func TestAdsCreateCampaignsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsCreateCampaignsBuilder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["data"])
}

func TestAdsCreateClientsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsCreateClientsBuilder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["data"])
}

func TestAdsCreateTargetGroupBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsCreateTargetGroupBuilder()

	b.AccountID(1)
	b.ClientID(1)
	b.Name("text")
	b.Lifetime(1)
	b.TargetPixelID(1)
	b.TargetPixelRules("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, 1, b.Params["client_id"])
	assert.Equal(t, "text", b.Params["name"])
	assert.Equal(t, 1, b.Params["lifetime"])
	assert.Equal(t, 1, b.Params["target_pixel_id"])
	assert.Equal(t, "text", b.Params["target_pixel_rules"])
}

func TestAdsDeleteAdsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsDeleteAdsBuilder()

	b.AccountID(1)
	b.IDs("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["ids"])
}

func TestAdsDeleteCampaignsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsDeleteCampaignsBuilder()

	b.AccountID(1)
	b.IDs("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["ids"])
}

func TestAdsDeleteClientsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsDeleteClientsBuilder()

	b.AccountID(1)
	b.IDs("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["ids"])
}

func TestAdsDeleteTargetGroupBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsDeleteTargetGroupBuilder()

	b.AccountID(1)
	b.ClientID(1)
	b.TargetGroupID(1)

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, 1, b.Params["client_id"])
	assert.Equal(t, 1, b.Params["target_group_id"])
}

func TestAdsGetAdsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetAdsBuilder()

	b.AccountID(1)
	b.AdIDs("text")
	b.CampaignIDs("text")
	b.ClientID(1)
	b.IncludeDeleted(true)
	b.Limit(1)
	b.Offset(1)

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["ad_ids"])
	assert.Equal(t, "text", b.Params["campaign_ids"])
	assert.Equal(t, 1, b.Params["client_id"])
	assert.Equal(t, true, b.Params["include_deleted"])
	assert.Equal(t, 1, b.Params["limit"])
	assert.Equal(t, 1, b.Params["offset"])
}

func TestAdsGetAdsLayoutBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetAdsLayoutBuilder()

	b.AccountID(1)
	b.AdIDs("text")
	b.CampaignIDs("text")
	b.ClientID(1)
	b.IncludeDeleted(true)
	b.Limit(1)
	b.Offset(1)

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["ad_ids"])
	assert.Equal(t, "text", b.Params["campaign_ids"])
	assert.Equal(t, 1, b.Params["client_id"])
	assert.Equal(t, true, b.Params["include_deleted"])
	assert.Equal(t, 1, b.Params["limit"])
	assert.Equal(t, 1, b.Params["offset"])
}

func TestAdsGetAdsTargetingBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetAdsTargetingBuilder()

	b.AccountID(1)
	b.AdIDs("text")
	b.CampaignIDs("text")
	b.ClientID(1)
	b.IncludeDeleted(true)
	b.Limit(1)
	b.Offset(1)
	b.OnlyDeleted(true)

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["ad_ids"])
	assert.Equal(t, "text", b.Params["campaign_ids"])
	assert.Equal(t, 1, b.Params["client_id"])
	assert.Equal(t, true, b.Params["include_deleted"])
	assert.Equal(t, 1, b.Params["limit"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, true, b.Params["only_deleted"])
}

func TestAdsGetBudgetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetBudgetBuilder()

	b.AccountID(1)

	assert.Equal(t, 1, b.Params["account_id"])
}

func TestAdsGetCampaignsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetCampaignsBuilder()

	b.AccountID(1)
	b.ClientID(1)
	b.IncludeDeleted(true)
	b.CampaignIDs("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, 1, b.Params["client_id"])
	assert.Equal(t, true, b.Params["include_deleted"])
	assert.Equal(t, "text", b.Params["campaign_ids"])
}

func TestAdsGetCategoriesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetCategoriesBuilder()

	b.Lang("text")

	assert.Equal(t, "text", b.Params["lang"])
}

func TestAdsGetClientsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetClientsBuilder()

	b.AccountID(1)

	assert.Equal(t, 1, b.Params["account_id"])
}

func TestAdsGetDemographicsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetDemographicsBuilder()

	b.AccountID(1)
	b.IDsType("text")
	b.IDs("text")
	b.Period("text")
	b.DateFrom("text")
	b.DateTo("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["ids_type"])
	assert.Equal(t, "text", b.Params["ids"])
	assert.Equal(t, "text", b.Params["period"])
	assert.Equal(t, "text", b.Params["date_from"])
	assert.Equal(t, "text", b.Params["date_to"])
}

func TestAdsGetFloodStatsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetFloodStatsBuilder()

	b.AccountID(1)

	assert.Equal(t, 1, b.Params["account_id"])
}

func TestAdsGetMusiciansBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetMusiciansBuilder()

	b.ArtistName("text")

	assert.Equal(t, "text", b.Params["artist_name"])
}

func TestAdsGetOfficeUsersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetOfficeUsersBuilder()

	b.AccountID(1)

	assert.Equal(t, 1, b.Params["account_id"])
}

func TestAdsGetPostsReachBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetPostsReachBuilder()

	b.AccountID(1)
	b.IDsType("text")
	b.IDs("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["ids_type"])
	assert.Equal(t, "text", b.Params["ids"])
}

func TestAdsGetRejectionReasonBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetRejectionReasonBuilder()

	b.AccountID(1)
	b.AdID(1)

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, 1, b.Params["ad_id"])
}

func TestAdsGetStatisticsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetStatisticsBuilder()

	b.AccountID(1)
	b.IDsType("text")
	b.IDs("text")
	b.Period("text")
	b.DateFrom("text")
	b.DateTo("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["ids_type"])
	assert.Equal(t, "text", b.Params["ids"])
	assert.Equal(t, "text", b.Params["period"])
	assert.Equal(t, "text", b.Params["date_from"])
	assert.Equal(t, "text", b.Params["date_to"])
}

func TestAdsGetSuggestionsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetSuggestionsBuilder()

	b.Section("text")
	b.IDs("text")
	b.Q("text")
	b.Country(1)
	b.Cities("text")
	b.Lang("text")

	assert.Equal(t, "text", b.Params["section"])
	assert.Equal(t, "text", b.Params["ids"])
	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, 1, b.Params["country"])
	assert.Equal(t, "text", b.Params["cities"])
	assert.Equal(t, "text", b.Params["lang"])
}

func TestAdsGetTargetGroupsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetTargetGroupsBuilder()

	b.AccountID(1)
	b.ClientID(1)
	b.Extended(true)

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, 1, b.Params["client_id"])
	assert.Equal(t, true, b.Params["extended"])
}

func TestAdsGetTargetingStatsBuilder(t *testing.T) {
	t.Parallel()

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

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, 1, b.Params["client_id"])
	assert.Equal(t, "text", b.Params["criteria"])
	assert.Equal(t, 1, b.Params["ad_id"])
	assert.Equal(t, 1, b.Params["ad_format"])
	assert.Equal(t, "text", b.Params["ad_platform"])
	assert.Equal(t, "text", b.Params["ad_platform_no_wall"])
	assert.Equal(t, "text", b.Params["ad_platform_no_ad_network"])
	assert.Equal(t, "text", b.Params["link_url"])
	assert.Equal(t, "text", b.Params["link_domain"])
}

func TestAdsGetUploadURLBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsGetUploadURLBuilder()

	b.AdFormat(1)
	b.Icon(1)

	assert.Equal(t, 1, b.Params["ad_format"])
	assert.Equal(t, 1, b.Params["icon"])
}

func TestAdsImportTargetContactsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsImportTargetContactsBuilder()

	b.AccountID(1)
	b.ClientID(1)
	b.TargetGroupID(1)
	b.Contacts("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, 1, b.Params["client_id"])
	assert.Equal(t, 1, b.Params["target_group_id"])
	assert.Equal(t, "text", b.Params["contacts"])
}

func TestAdsRemoveOfficeUsersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsRemoveOfficeUsersBuilder()

	b.AccountID(1)
	b.IDs("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["ids"])
}

func TestAdsUpdateAdsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsUpdateAdsBuilder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["data"])
}

func TestAdsUpdateCampaignsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsUpdateCampaignsBuilder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["data"])
}

func TestAdsUpdateClientsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsUpdateClientsBuilder()

	b.AccountID(1)
	b.Data("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, "text", b.Params["data"])
}

func TestAdsUpdateTargetGroupBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAdsUpdateTargetGroupBuilder()

	b.AccountID(1)
	b.ClientID(1)
	b.TargetGroupID(1)
	b.Name("text")
	b.Domain("text")
	b.Lifetime(1)
	b.TargetPixelID(1)
	b.TargetPixelRules("text")

	assert.Equal(t, 1, b.Params["account_id"])
	assert.Equal(t, 1, b.Params["client_id"])
	assert.Equal(t, 1, b.Params["target_group_id"])
	assert.Equal(t, "text", b.Params["name"])
	assert.Equal(t, "text", b.Params["domain"])
	assert.Equal(t, 1, b.Params["lifetime"])
	assert.Equal(t, 1, b.Params["target_pixel_id"])
	assert.Equal(t, "text", b.Params["target_pixel_rules"])
}
