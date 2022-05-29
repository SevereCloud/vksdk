package params_test

import (
	"testing"

	"github.com/Derad6709/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestGroupsAddAddressBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsAddAddressBuilder()

	b.GroupID(1)
	b.Title("text")
	b.Address("text")
	b.AdditionalAddress("text")
	b.CountryID(1)
	b.CityID(1)
	b.MetroID(1)
	b.Latitude(1.1)
	b.Longitude(1.1)
	b.Phone("text")
	b.WorkInfoStatus("text")
	b.Timetable("text")
	b.IsMainAddress(true)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["address"], "text")
	assert.Equal(t, b.Params["additional_address"], "text")
	assert.Equal(t, b.Params["country_id"], 1)
	assert.Equal(t, b.Params["city_id"], 1)
	assert.Equal(t, b.Params["metro_id"], 1)
	assert.Equal(t, b.Params["latitude"], 1.1)
	assert.Equal(t, b.Params["longitude"], 1.1)
	assert.Equal(t, b.Params["phone"], "text")
	assert.Equal(t, b.Params["work_info_status"], "text")
	assert.Equal(t, b.Params["timetable"], "text")
	assert.Equal(t, b.Params["is_main_address"], true)
}

func TestGroupsAddCallbackServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsAddCallbackServerBuilder()

	b.GroupID(1)
	b.URL("text")
	b.Title("text")
	b.SecretKey("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["url"], "text")
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["secret_key"], "text")
}

func TestGroupsAddLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsAddLinkBuilder()

	b.GroupID(1)
	b.Link("text")
	b.Text("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["link"], "text")
	assert.Equal(t, b.Params["text"], "text")
}

func TestGroupsApproveRequestBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsApproveRequestBuilder()

	b.GroupID(1)
	b.UserID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
}

func TestGroupsBanBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsBanBuilder()

	b.GroupID(1)
	b.OwnerID(1)
	b.EndDate(1)
	b.Reason(1)
	b.Comment("text")
	b.CommentVisible(true)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["end_date"], 1)
	assert.Equal(t, b.Params["reason"], 1)
	assert.Equal(t, b.Params["comment"], "text")
	assert.Equal(t, b.Params["comment_visible"], true)
}

func TestGroupsCreateBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsCreateBuilder()

	b.Title("text")
	b.Description("text")
	b.Type("text")
	b.PublicCategory(1)
	b.PublicSubcategory(1)
	b.Subtype(1)

	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["description"], "text")
	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["public_category"], 1)
	assert.Equal(t, b.Params["public_subcategory"], 1)
	assert.Equal(t, b.Params["subtype"], 1)
}

func TestGroupsDeleteCallbackServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsDeleteCallbackServerBuilder()

	b.GroupID(1)
	b.ServerID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["server_id"], 1)
}

func TestGroupsDeleteLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsDeleteLinkBuilder()

	b.GroupID(1)
	b.LinkID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["link_id"], 1)
}

func TestGroupsDisableOnlineBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsDisableOnlineBuilder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestGroupsEditBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsEditBuilder()

	b.GroupID(1)
	b.Title("text")
	b.Description("text")
	b.ScreenName("text")
	b.Access(1)
	b.Website("text")
	b.Subject("text")
	b.Email("text")
	b.Phone("text")
	b.Rss("text")
	b.EventStartDate(1)
	b.EventFinishDate(1)
	b.EventGroupID(1)
	b.PublicCategory(1)
	b.PublicSubcategory(1)
	b.PublicDate("text")
	b.Wall(1)
	b.Topics(1)
	b.Photos(1)
	b.Video(1)
	b.Audio(1)
	b.Links(true)
	b.Events(true)
	b.Places(true)
	b.Contacts(true)
	b.Docs(1)
	b.Wiki(1)
	b.Messages(true)
	b.Articles(true)
	b.Addresses(true)
	b.AgeLimits(1)
	b.Market(true)
	b.MarketComments(true)
	b.MarketCountry([]int{1})
	b.MarketCity([]int{1})
	b.MarketCurrency(1)
	b.MarketContact(1)
	b.MarketWiki(1)
	b.ObsceneFilter(true)
	b.ObsceneStopwords(true)
	b.ObsceneWords([]string{"text"})
	b.MainSection(1)
	b.SecondarySection(1)
	b.Country(1)
	b.City(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["description"], "text")
	assert.Equal(t, b.Params["screen_name"], "text")
	assert.Equal(t, b.Params["access"], 1)
	assert.Equal(t, b.Params["website"], "text")
	assert.Equal(t, b.Params["subject"], "text")
	assert.Equal(t, b.Params["email"], "text")
	assert.Equal(t, b.Params["phone"], "text")
	assert.Equal(t, b.Params["rss"], "text")
	assert.Equal(t, b.Params["event_start_date"], 1)
	assert.Equal(t, b.Params["event_finish_date"], 1)
	assert.Equal(t, b.Params["event_group_id"], 1)
	assert.Equal(t, b.Params["public_category"], 1)
	assert.Equal(t, b.Params["public_subcategory"], 1)
	assert.Equal(t, b.Params["public_date"], "text")
	assert.Equal(t, b.Params["wall"], 1)
	assert.Equal(t, b.Params["topics"], 1)
	assert.Equal(t, b.Params["photos"], 1)
	assert.Equal(t, b.Params["video"], 1)
	assert.Equal(t, b.Params["audio"], 1)
	assert.Equal(t, b.Params["links"], true)
	assert.Equal(t, b.Params["events"], true)
	assert.Equal(t, b.Params["places"], true)
	assert.Equal(t, b.Params["contacts"], true)
	assert.Equal(t, b.Params["docs"], 1)
	assert.Equal(t, b.Params["wiki"], 1)
	assert.Equal(t, b.Params["messages"], true)
	assert.Equal(t, b.Params["articles"], true)
	assert.Equal(t, b.Params["addresses"], true)
	assert.Equal(t, b.Params["age_limits"], 1)
	assert.Equal(t, b.Params["market"], true)
	assert.Equal(t, b.Params["market_comments"], true)
	assert.Equal(t, b.Params["market_country"], []int{1})
	assert.Equal(t, b.Params["market_city"], []int{1})
	assert.Equal(t, b.Params["market_currency"], 1)
	assert.Equal(t, b.Params["market_contact"], 1)
	assert.Equal(t, b.Params["market_wiki"], 1)
	assert.Equal(t, b.Params["obscene_filter"], true)
	assert.Equal(t, b.Params["obscene_stopwords"], true)
	assert.Equal(t, b.Params["obscene_words"], []string{"text"})
	assert.Equal(t, b.Params["main_section"], 1)
	assert.Equal(t, b.Params["secondary_section"], 1)
	assert.Equal(t, b.Params["country"], 1)
	assert.Equal(t, b.Params["city"], 1)
}

func TestGroupsEditAddressBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsEditAddressBuilder()

	b.GroupID(1)
	b.AddressID(1)
	b.Title("text")
	b.Address("text")
	b.AdditionalAddress("text")
	b.CountryID(1)
	b.CityID(1)
	b.MetroID(1)
	b.Latitude(1.1)
	b.Longitude(1.1)
	b.Phone("text")
	b.WorkInfoStatus("text")
	b.Timetable("text")
	b.IsMainAddress(true)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["address_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["address"], "text")
	assert.Equal(t, b.Params["additional_address"], "text")
	assert.Equal(t, b.Params["country_id"], 1)
	assert.Equal(t, b.Params["city_id"], 1)
	assert.Equal(t, b.Params["metro_id"], 1)
	assert.Equal(t, b.Params["latitude"], 1.1)
	assert.Equal(t, b.Params["longitude"], 1.1)
	assert.Equal(t, b.Params["phone"], "text")
	assert.Equal(t, b.Params["work_info_status"], "text")
	assert.Equal(t, b.Params["timetable"], "text")
	assert.Equal(t, b.Params["is_main_address"], true)
}

func TestGroupsEditCallbackServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsEditCallbackServerBuilder()

	b.GroupID(1)
	b.ServerID(1)
	b.URL("text")
	b.Title("text")
	b.SecretKey("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["server_id"], 1)
	assert.Equal(t, b.Params["url"], "text")
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["secret_key"], "text")
}

func TestGroupsEditLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsEditLinkBuilder()

	b.GroupID(1)
	b.LinkID(1)
	b.Text("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["link_id"], 1)
	assert.Equal(t, b.Params["text"], "text")
}

func TestGroupsEditManagerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsEditManagerBuilder()

	b.GroupID(1)
	b.UserID(1)
	b.Role("text")
	b.IsContact(true)
	b.ContactPosition("text")
	b.ContactPhone("text")
	b.ContactEmail("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["role"], "text")
	assert.Equal(t, b.Params["is_contact"], true)
	assert.Equal(t, b.Params["contact_position"], "text")
	assert.Equal(t, b.Params["contact_phone"], "text")
	assert.Equal(t, b.Params["contact_email"], "text")
}

func TestGroupsEnableOnlineBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsEnableOnlineBuilder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestGroupsGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetBuilder()

	b.UserID(1)
	b.Extended(true)
	b.Filter([]string{"test"})
	b.Fields([]string{"test"})
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["filter"], []string{"test"})
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestGroupsGetAddressesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetAddressesBuilder()

	b.GroupID(1)
	b.AddressIDs([]int{1})
	b.Latitude(1.1)
	b.Longitude(1.1)
	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["address_ids"], []int{1})
	assert.Equal(t, b.Params["latitude"], 1.1)
	assert.Equal(t, b.Params["longitude"], 1.1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestGroupsGetBannedBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetBannedBuilder()

	b.GroupID(1)
	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"test"})
	b.OwnerID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["owner_id"], 1)
}

func TestGroupsGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetByIDBuilder()

	b.GroupIDs([]string{"text"})
	b.GroupID("text")
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["group_ids"], []string{"text"})
	assert.Equal(t, b.Params["group_id"], "text")
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestGroupsGetCallbackConfirmationCodeBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetCallbackConfirmationCodeBuilder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestGroupsGetCallbackServersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetCallbackServersBuilder()

	b.GroupID(1)
	b.ServerIDs([]int{1})

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["server_ids"], []int{1})
}

func TestGroupsGetCallbackSettingsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetCallbackSettingsBuilder()

	b.GroupID(1)
	b.ServerID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["server_id"], 1)
}

func TestGroupsGetCatalogBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetCatalogBuilder()

	b.CategoryID(1)
	b.SubcategoryID(1)

	assert.Equal(t, b.Params["category_id"], 1)
	assert.Equal(t, b.Params["subcategory_id"], 1)
}

func TestGroupsGetCatalogInfoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetCatalogInfoBuilder()

	b.Extended(true)
	b.Subcategories(true)

	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["subcategories"], true)
}

func TestGroupsGetInvitedUsersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetInvitedUsersBuilder()

	b.GroupID(1)
	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"test"})
	b.NameCase("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["name_case"], "text")
}

func TestGroupsGetInvitesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetInvitesBuilder()

	b.Offset(1)
	b.Count(1)
	b.Extended(true)

	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["extended"], true)
}

func TestGroupsGetLongPollServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetLongPollServerBuilder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestGroupsGetLongPollSettingsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetLongPollSettingsBuilder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestGroupsGetMembersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetMembersBuilder()

	b.GroupID("text")
	b.Sort("text")
	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"test"})
	b.Filter("text")

	assert.Equal(t, b.Params["group_id"], "text")
	assert.Equal(t, b.Params["sort"], "text")
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["filter"], "text")
}

func TestGroupsGetRequestsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetRequestsBuilder()

	b.GroupID(1)
	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestGroupsGetSettingsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetSettingsBuilder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestGroupsGetTagListBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetTagListBuilder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestGroupsInviteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsInviteBuilder()

	b.GroupID(1)
	b.UserID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
}

func TestGroupsIsMemberBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsIsMemberBuilder()

	b.GroupID("text")
	b.UserID(1)
	b.UserIDs([]int{1})
	b.Extended(true)

	assert.Equal(t, b.Params["group_id"], "text")
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["extended"], true)
}

func TestGroupsJoinBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsJoinBuilder()

	b.GroupID(1)
	b.NotSure("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["not_sure"], "text")
}

func TestGroupsLeaveBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsLeaveBuilder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestGroupsRemoveUserBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsRemoveUserBuilder()

	b.GroupID(1)
	b.UserID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
}

func TestGroupsReorderLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsReorderLinkBuilder()

	b.GroupID(1)
	b.LinkID(1)
	b.After(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["link_id"], 1)
	assert.Equal(t, b.Params["after"], 1)
}

func TestGroupsSearchBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsSearchBuilder()

	b.Q("text")
	b.Type("text")
	b.CountryID(1)
	b.CityID(1)
	b.Future(true)
	b.Market(true)
	b.Sort(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["country_id"], 1)
	assert.Equal(t, b.Params["city_id"], 1)
	assert.Equal(t, b.Params["future"], true)
	assert.Equal(t, b.Params["market"], true)
	assert.Equal(t, b.Params["sort"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestGroupsSetCallbackSettingsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsSetCallbackSettingsBuilder()

	b.GroupID(1)
	b.ServerID(1)
	b.APIVersion("text")
	b.MessageNew(true)
	b.MessageReply(true)
	b.MessageAllow(true)
	b.MessageEdit(true)
	b.MessageDeny(true)
	b.MessageTypingState(true)
	b.PhotoNew(true)
	b.AudioNew(true)
	b.VideoNew(true)
	b.WallReplyNew(true)
	b.WallReplyEdit(true)
	b.WallReplyDelete(true)
	b.WallReplyRestore(true)
	b.WallPostNew(true)
	b.WallRepost(true)
	b.BoardPostNew(true)
	b.BoardPostEdit(true)
	b.BoardPostRestore(true)
	b.BoardPostDelete(true)
	b.PhotoCommentNew(true)
	b.PhotoCommentEdit(true)
	b.PhotoCommentDelete(true)
	b.PhotoCommentRestore(true)
	b.VideoCommentNew(true)
	b.VideoCommentEdit(true)
	b.VideoCommentDelete(true)
	b.VideoCommentRestore(true)
	b.MarketCommentNew(true)
	b.MarketCommentEdit(true)
	b.MarketCommentDelete(true)
	b.MarketCommentRestore(true)
	b.MarketOrderNew(true)
	b.MarketOrderEdit(true)
	b.PollVoteNew(true)
	b.GroupJoin(true)
	b.GroupLeave(true)
	b.GroupChangeSettings(true)
	b.GroupChangePhoto(true)
	b.GroupOfficersEdit(true)
	b.UserBlock(true)
	b.UserUnblock(true)
	b.LeadFormsNew(true)
	b.LikeAdd(true)
	b.LikeRemove(true)
	b.DonutSubscriptionCreate(true)
	b.DonutSubscriptionProlonged(true)
	b.DonutSubscriptionExpired(true)
	b.DonutSubscriptionCancelled(true)
	b.DonutSubscriptionPriceChanged(true)
	b.DonutMoneyWithdraw(true)
	b.DonutMoneyWithdrawError(true)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["server_id"], 1)
	assert.Equal(t, b.Params["api_version"], "text")
	assert.Equal(t, b.Params["message_new"], true)
	assert.Equal(t, b.Params["message_reply"], true)
	assert.Equal(t, b.Params["message_allow"], true)
	assert.Equal(t, b.Params["message_edit"], true)
	assert.Equal(t, b.Params["message_deny"], true)
	assert.Equal(t, b.Params["message_typing_state"], true)
	assert.Equal(t, b.Params["photo_new"], true)
	assert.Equal(t, b.Params["audio_new"], true)
	assert.Equal(t, b.Params["video_new"], true)
	assert.Equal(t, b.Params["wall_reply_new"], true)
	assert.Equal(t, b.Params["wall_reply_edit"], true)
	assert.Equal(t, b.Params["wall_reply_delete"], true)
	assert.Equal(t, b.Params["wall_reply_restore"], true)
	assert.Equal(t, b.Params["wall_post_new"], true)
	assert.Equal(t, b.Params["wall_repost"], true)
	assert.Equal(t, b.Params["board_post_new"], true)
	assert.Equal(t, b.Params["board_post_edit"], true)
	assert.Equal(t, b.Params["board_post_restore"], true)
	assert.Equal(t, b.Params["board_post_delete"], true)
	assert.Equal(t, b.Params["photo_comment_new"], true)
	assert.Equal(t, b.Params["photo_comment_edit"], true)
	assert.Equal(t, b.Params["photo_comment_delete"], true)
	assert.Equal(t, b.Params["photo_comment_restore"], true)
	assert.Equal(t, b.Params["video_comment_new"], true)
	assert.Equal(t, b.Params["video_comment_edit"], true)
	assert.Equal(t, b.Params["video_comment_delete"], true)
	assert.Equal(t, b.Params["video_comment_restore"], true)
	assert.Equal(t, b.Params["market_comment_new"], true)
	assert.Equal(t, b.Params["market_comment_edit"], true)
	assert.Equal(t, b.Params["market_comment_delete"], true)
	assert.Equal(t, b.Params["market_comment_restore"], true)
	assert.Equal(t, b.Params["market_order_new"], true)
	assert.Equal(t, b.Params["market_order_edit"], true)
	assert.Equal(t, b.Params["poll_vote_new"], true)
	assert.Equal(t, b.Params["group_join"], true)
	assert.Equal(t, b.Params["group_leave"], true)
	assert.Equal(t, b.Params["group_change_settings"], true)
	assert.Equal(t, b.Params["group_change_photo"], true)
	assert.Equal(t, b.Params["group_officers_edit"], true)
	assert.Equal(t, b.Params["user_block"], true)
	assert.Equal(t, b.Params["user_unblock"], true)
	assert.Equal(t, b.Params["lead_forms_new"], true)
	assert.Equal(t, b.Params["like_add"], true)
	assert.Equal(t, b.Params["like_remove"], true)
	assert.Equal(t, b.Params["donut_subscription_create"], true)
	assert.Equal(t, b.Params["donut_subscription_prolonged"], true)
	assert.Equal(t, b.Params["donut_subscription_expired"], true)
	assert.Equal(t, b.Params["donut_subscription_cancelled"], true)
	assert.Equal(t, b.Params["donut_subscription_price_changed"], true)
	assert.Equal(t, b.Params["donut_money_withdraw"], true)
	assert.Equal(t, b.Params["donut_money_withdraw_error"], true)
}

func TestGroupsSetLongPollSettingsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsSetLongPollSettingsBuilder()

	b.GroupID(1)
	b.Enabled(true)
	b.APIVersion("text")
	b.MessageNew(true)
	b.MessageReply(true)
	b.MessageAllow(true)
	b.MessageDeny(true)
	b.MessageEdit(true)
	b.MessageTypingState(true)
	b.PhotoNew(true)
	b.AudioNew(true)
	b.VideoNew(true)
	b.WallReplyNew(true)
	b.WallReplyEdit(true)
	b.WallReplyDelete(true)
	b.WallReplyRestore(true)
	b.WallPostNew(true)
	b.WallRepost(true)
	b.BoardPostNew(true)
	b.BoardPostEdit(true)
	b.BoardPostRestore(true)
	b.BoardPostDelete(true)
	b.PhotoCommentNew(true)
	b.PhotoCommentEdit(true)
	b.PhotoCommentDelete(true)
	b.PhotoCommentRestore(true)
	b.VideoCommentNew(true)
	b.VideoCommentEdit(true)
	b.VideoCommentDelete(true)
	b.VideoCommentRestore(true)
	b.MarketCommentNew(true)
	b.MarketCommentEdit(true)
	b.MarketCommentDelete(true)
	b.MarketOrderNew(true)
	b.MarketOrderEdit(true)
	b.MarketCommentRestore(true)
	b.PollVoteNew(true)
	b.GroupJoin(true)
	b.GroupLeave(true)
	b.GroupChangeSettings(true)
	b.GroupChangePhoto(true)
	b.GroupOfficersEdit(true)
	b.UserBlock(true)
	b.UserUnblock(true)
	b.LikeAdd(true)
	b.LikeRemove(true)
	b.DonutSubscriptionCreate(true)
	b.DonutSubscriptionProlonged(true)
	b.DonutSubscriptionExpired(true)
	b.DonutSubscriptionCancelled(true)
	b.DonutSubscriptionPriceChanged(true)
	b.DonutMoneyWithdraw(true)
	b.DonutMoneyWithdrawError(true)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["enabled"], true)
	assert.Equal(t, b.Params["api_version"], "text")
	assert.Equal(t, b.Params["message_new"], true)
	assert.Equal(t, b.Params["message_reply"], true)
	assert.Equal(t, b.Params["message_allow"], true)
	assert.Equal(t, b.Params["message_deny"], true)
	assert.Equal(t, b.Params["message_edit"], true)
	assert.Equal(t, b.Params["message_typing_state"], true)
	assert.Equal(t, b.Params["photo_new"], true)
	assert.Equal(t, b.Params["audio_new"], true)
	assert.Equal(t, b.Params["video_new"], true)
	assert.Equal(t, b.Params["wall_reply_new"], true)
	assert.Equal(t, b.Params["wall_reply_edit"], true)
	assert.Equal(t, b.Params["wall_reply_delete"], true)
	assert.Equal(t, b.Params["wall_reply_restore"], true)
	assert.Equal(t, b.Params["wall_post_new"], true)
	assert.Equal(t, b.Params["wall_repost"], true)
	assert.Equal(t, b.Params["board_post_new"], true)
	assert.Equal(t, b.Params["board_post_edit"], true)
	assert.Equal(t, b.Params["board_post_restore"], true)
	assert.Equal(t, b.Params["board_post_delete"], true)
	assert.Equal(t, b.Params["photo_comment_new"], true)
	assert.Equal(t, b.Params["photo_comment_edit"], true)
	assert.Equal(t, b.Params["photo_comment_delete"], true)
	assert.Equal(t, b.Params["photo_comment_restore"], true)
	assert.Equal(t, b.Params["video_comment_new"], true)
	assert.Equal(t, b.Params["video_comment_edit"], true)
	assert.Equal(t, b.Params["video_comment_delete"], true)
	assert.Equal(t, b.Params["video_comment_restore"], true)
	assert.Equal(t, b.Params["market_comment_new"], true)
	assert.Equal(t, b.Params["market_comment_edit"], true)
	assert.Equal(t, b.Params["market_comment_delete"], true)
	assert.Equal(t, b.Params["market_comment_restore"], true)
	assert.Equal(t, b.Params["market_order_new"], true)
	assert.Equal(t, b.Params["market_order_edit"], true)
	assert.Equal(t, b.Params["poll_vote_new"], true)
	assert.Equal(t, b.Params["group_join"], true)
	assert.Equal(t, b.Params["group_leave"], true)
	assert.Equal(t, b.Params["group_change_settings"], true)
	assert.Equal(t, b.Params["group_change_photo"], true)
	assert.Equal(t, b.Params["group_officers_edit"], true)
	assert.Equal(t, b.Params["user_block"], true)
	assert.Equal(t, b.Params["user_unblock"], true)
	assert.Equal(t, b.Params["like_add"], true)
	assert.Equal(t, b.Params["like_remove"], true)
	assert.Equal(t, b.Params["donut_subscription_create"], true)
	assert.Equal(t, b.Params["donut_subscription_prolonged"], true)
	assert.Equal(t, b.Params["donut_subscription_expired"], true)
	assert.Equal(t, b.Params["donut_subscription_cancelled"], true)
	assert.Equal(t, b.Params["donut_subscription_price_changed"], true)
	assert.Equal(t, b.Params["donut_money_withdraw"], true)
	assert.Equal(t, b.Params["donut_money_withdraw_error"], true)
}

func TestGroupsSetUserNoteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsSetUserNoteBuilder()

	b.GroupID(1)
	b.UserID(1)
	b.Note("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["note"], "text")
}

func TestGroupsTagAddBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsTagAddBuilder()

	b.GroupID(1)
	b.TagName("text")
	b.TagColor("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["tag_name"], "text")
	assert.Equal(t, b.Params["tag_color"], "text")
}

func TestGroupsTagBindBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsTagBindBuilder()

	b.GroupID(1)
	b.TagID(1)
	b.UserID(1)
	b.Act("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["tag_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["act"], "text")

	b.Bind()
	assert.Equal(t, b.Params["act"], "bind")

	b.Unbind()
	assert.Equal(t, b.Params["act"], "unbind")
}

func TestGroupsTagDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsTagDeleteBuilder()

	b.GroupID(1)
	b.TagID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["tag_id"], 1)
}

func TestGroupsTagUpdateBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsTagUpdateBuilder()

	b.GroupID(1)
	b.TagID(1)
	b.TagName("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["tag_id"], 1)
	assert.Equal(t, b.Params["tag_name"], "text")
}

func TestGroupsToggleMarketBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsToggleMarketBuilder()

	b.GroupID(1)
	b.State("none")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["state"], params.GroupMarketState("none"))
}

func TestGroupsUnbanBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsUnbanBuilder()

	b.GroupID(1)
	b.OwnerID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
}
