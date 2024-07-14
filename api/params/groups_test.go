package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api/params"
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

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, "text", b.Params["address"])
	assert.Equal(t, "text", b.Params["additional_address"])
	assert.Equal(t, 1, b.Params["country_id"])
	assert.Equal(t, 1, b.Params["city_id"])
	assert.Equal(t, 1, b.Params["metro_id"])
	assert.InEpsilon(t, 1.1, b.Params["latitude"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["longitude"], 0.1)
	assert.Equal(t, "text", b.Params["phone"])
	assert.Equal(t, "text", b.Params["work_info_status"])
	assert.Equal(t, "text", b.Params["timetable"])
	assert.Equal(t, true, b.Params["is_main_address"])
}

func TestGroupsAddCallbackServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsAddCallbackServerBuilder()

	b.GroupID(1)
	b.URL("text")
	b.Title("text")
	b.SecretKey("text")

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, "text", b.Params["url"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, "text", b.Params["secret_key"])
}

func TestGroupsAddLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsAddLinkBuilder()

	b.GroupID(1)
	b.Link("text")
	b.Text("text")

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, "text", b.Params["link"])
	assert.Equal(t, "text", b.Params["text"])
}

func TestGroupsApproveRequestBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsApproveRequestBuilder()

	b.GroupID(1)
	b.UserID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["user_id"])
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

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["end_date"])
	assert.Equal(t, 1, b.Params["reason"])
	assert.Equal(t, "text", b.Params["comment"])
	assert.Equal(t, true, b.Params["comment_visible"])
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

	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, "text", b.Params["description"])
	assert.Equal(t, "text", b.Params["type"])
	assert.Equal(t, 1, b.Params["public_category"])
	assert.Equal(t, 1, b.Params["public_subcategory"])
	assert.Equal(t, 1, b.Params["subtype"])
}

func TestGroupsDeleteCallbackServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsDeleteCallbackServerBuilder()

	b.GroupID(1)
	b.ServerID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["server_id"])
}

func TestGroupsDeleteLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsDeleteLinkBuilder()

	b.GroupID(1)
	b.LinkID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["link_id"])
}

func TestGroupsDisableOnlineBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsDisableOnlineBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
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

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, "text", b.Params["description"])
	assert.Equal(t, "text", b.Params["screen_name"])
	assert.Equal(t, 1, b.Params["access"])
	assert.Equal(t, "text", b.Params["website"])
	assert.Equal(t, "text", b.Params["subject"])
	assert.Equal(t, "text", b.Params["email"])
	assert.Equal(t, "text", b.Params["phone"])
	assert.Equal(t, "text", b.Params["rss"])
	assert.Equal(t, 1, b.Params["event_start_date"])
	assert.Equal(t, 1, b.Params["event_finish_date"])
	assert.Equal(t, 1, b.Params["event_group_id"])
	assert.Equal(t, 1, b.Params["public_category"])
	assert.Equal(t, 1, b.Params["public_subcategory"])
	assert.Equal(t, "text", b.Params["public_date"])
	assert.Equal(t, 1, b.Params["wall"])
	assert.Equal(t, 1, b.Params["topics"])
	assert.Equal(t, 1, b.Params["photos"])
	assert.Equal(t, 1, b.Params["video"])
	assert.Equal(t, 1, b.Params["audio"])
	assert.Equal(t, true, b.Params["links"])
	assert.Equal(t, true, b.Params["events"])
	assert.Equal(t, true, b.Params["places"])
	assert.Equal(t, true, b.Params["contacts"])
	assert.Equal(t, 1, b.Params["docs"])
	assert.Equal(t, 1, b.Params["wiki"])
	assert.Equal(t, true, b.Params["messages"])
	assert.Equal(t, true, b.Params["articles"])
	assert.Equal(t, true, b.Params["addresses"])
	assert.Equal(t, 1, b.Params["age_limits"])
	assert.Equal(t, true, b.Params["market"])
	assert.Equal(t, true, b.Params["market_comments"])
	assert.Equal(t, []int{1}, b.Params["market_country"])
	assert.Equal(t, []int{1}, b.Params["market_city"])
	assert.Equal(t, 1, b.Params["market_currency"])
	assert.Equal(t, 1, b.Params["market_contact"])
	assert.Equal(t, 1, b.Params["market_wiki"])
	assert.Equal(t, true, b.Params["obscene_filter"])
	assert.Equal(t, true, b.Params["obscene_stopwords"])
	assert.Equal(t, []string{"text"}, b.Params["obscene_words"])
	assert.Equal(t, 1, b.Params["main_section"])
	assert.Equal(t, 1, b.Params["secondary_section"])
	assert.Equal(t, 1, b.Params["country"])
	assert.Equal(t, 1, b.Params["city"])
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

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["address_id"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, "text", b.Params["address"])
	assert.Equal(t, "text", b.Params["additional_address"])
	assert.Equal(t, 1, b.Params["country_id"])
	assert.Equal(t, 1, b.Params["city_id"])
	assert.Equal(t, 1, b.Params["metro_id"])
	assert.InEpsilon(t, 1.1, b.Params["latitude"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["longitude"], 0.1)
	assert.Equal(t, "text", b.Params["phone"])
	assert.Equal(t, "text", b.Params["work_info_status"])
	assert.Equal(t, "text", b.Params["timetable"])
	assert.Equal(t, true, b.Params["is_main_address"])
}

func TestGroupsEditCallbackServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsEditCallbackServerBuilder()

	b.GroupID(1)
	b.ServerID(1)
	b.URL("text")
	b.Title("text")
	b.SecretKey("text")

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["server_id"])
	assert.Equal(t, "text", b.Params["url"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, "text", b.Params["secret_key"])
}

func TestGroupsEditLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsEditLinkBuilder()

	b.GroupID(1)
	b.LinkID(1)
	b.Text("text")

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["link_id"])
	assert.Equal(t, "text", b.Params["text"])
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

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, "text", b.Params["role"])
	assert.Equal(t, true, b.Params["is_contact"])
	assert.Equal(t, "text", b.Params["contact_position"])
	assert.Equal(t, "text", b.Params["contact_phone"])
	assert.Equal(t, "text", b.Params["contact_email"])
}

func TestGroupsEnableOnlineBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsEnableOnlineBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
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

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"test"}, b.Params["filter"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
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

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, []int{1}, b.Params["address_ids"])
	assert.InEpsilon(t, 1.1, b.Params["latitude"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["longitude"], 0.1)
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestGroupsGetBannedBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetBannedBuilder()

	b.GroupID(1)
	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"test"})
	b.OwnerID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, 1, b.Params["owner_id"])
}

func TestGroupsGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetByIDBuilder()

	b.GroupIDs([]string{"text"})
	b.GroupID("text")
	b.Fields([]string{"test"})

	assert.Equal(t, []string{"text"}, b.Params["group_ids"])
	assert.Equal(t, "text", b.Params["group_id"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestGroupsGetCallbackConfirmationCodeBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetCallbackConfirmationCodeBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
}

func TestGroupsGetCallbackServersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetCallbackServersBuilder()

	b.GroupID(1)
	b.ServerIDs([]int{1})

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, []int{1}, b.Params["server_ids"])
}

func TestGroupsGetCallbackSettingsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetCallbackSettingsBuilder()

	b.GroupID(1)
	b.ServerID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["server_id"])
}

func TestGroupsGetCatalogBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetCatalogBuilder()

	b.CategoryID(1)
	b.SubcategoryID(1)

	assert.Equal(t, 1, b.Params["category_id"])
	assert.Equal(t, 1, b.Params["subcategory_id"])
}

func TestGroupsGetCatalogInfoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetCatalogInfoBuilder()

	b.Extended(true)
	b.Subcategories(true)

	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, true, b.Params["subcategories"])
}

func TestGroupsGetInvitedUsersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetInvitedUsersBuilder()

	b.GroupID(1)
	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"test"})
	b.NameCase("text")

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, "text", b.Params["name_case"])
}

func TestGroupsGetInvitesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetInvitesBuilder()

	b.Offset(1)
	b.Count(1)
	b.Extended(true)

	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["extended"])
}

func TestGroupsGetLongPollServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetLongPollServerBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
}

func TestGroupsGetLongPollSettingsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetLongPollSettingsBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
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

	assert.Equal(t, "text", b.Params["group_id"])
	assert.Equal(t, "text", b.Params["sort"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, "text", b.Params["filter"])
}

func TestGroupsGetRequestsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetRequestsBuilder()

	b.GroupID(1)
	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"test"})

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestGroupsGetSettingsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetSettingsBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
}

func TestGroupsGetTagListBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsGetTagListBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
}

func TestGroupsInviteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsInviteBuilder()

	b.GroupID(1)
	b.UserID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["user_id"])
}

func TestGroupsIsMemberBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsIsMemberBuilder()

	b.GroupID("text")
	b.UserID(1)
	b.UserIDs([]int{1})
	b.Extended(true)

	assert.Equal(t, "text", b.Params["group_id"])
	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, []int{1}, b.Params["user_ids"])
	assert.Equal(t, true, b.Params["extended"])
}

func TestGroupsJoinBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsJoinBuilder()

	b.GroupID(1)
	b.NotSure("text")

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, "text", b.Params["not_sure"])
}

func TestGroupsLeaveBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsLeaveBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
}

func TestGroupsRemoveUserBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsRemoveUserBuilder()

	b.GroupID(1)
	b.UserID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["user_id"])
}

func TestGroupsReorderLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsReorderLinkBuilder()

	b.GroupID(1)
	b.LinkID(1)
	b.After(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["link_id"])
	assert.Equal(t, 1, b.Params["after"])
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

	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, "text", b.Params["type"])
	assert.Equal(t, 1, b.Params["country_id"])
	assert.Equal(t, 1, b.Params["city_id"])
	assert.Equal(t, true, b.Params["future"])
	assert.Equal(t, true, b.Params["market"])
	assert.Equal(t, 1, b.Params["sort"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
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

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["server_id"])
	assert.Equal(t, "text", b.Params["api_version"])
	assert.Equal(t, true, b.Params["message_new"])
	assert.Equal(t, true, b.Params["message_reply"])
	assert.Equal(t, true, b.Params["message_allow"])
	assert.Equal(t, true, b.Params["message_edit"])
	assert.Equal(t, true, b.Params["message_deny"])
	assert.Equal(t, true, b.Params["message_typing_state"])
	assert.Equal(t, true, b.Params["photo_new"])
	assert.Equal(t, true, b.Params["audio_new"])
	assert.Equal(t, true, b.Params["video_new"])
	assert.Equal(t, true, b.Params["wall_reply_new"])
	assert.Equal(t, true, b.Params["wall_reply_edit"])
	assert.Equal(t, true, b.Params["wall_reply_delete"])
	assert.Equal(t, true, b.Params["wall_reply_restore"])
	assert.Equal(t, true, b.Params["wall_post_new"])
	assert.Equal(t, true, b.Params["wall_repost"])
	assert.Equal(t, true, b.Params["board_post_new"])
	assert.Equal(t, true, b.Params["board_post_edit"])
	assert.Equal(t, true, b.Params["board_post_restore"])
	assert.Equal(t, true, b.Params["board_post_delete"])
	assert.Equal(t, true, b.Params["photo_comment_new"])
	assert.Equal(t, true, b.Params["photo_comment_edit"])
	assert.Equal(t, true, b.Params["photo_comment_delete"])
	assert.Equal(t, true, b.Params["photo_comment_restore"])
	assert.Equal(t, true, b.Params["video_comment_new"])
	assert.Equal(t, true, b.Params["video_comment_edit"])
	assert.Equal(t, true, b.Params["video_comment_delete"])
	assert.Equal(t, true, b.Params["video_comment_restore"])
	assert.Equal(t, true, b.Params["market_comment_new"])
	assert.Equal(t, true, b.Params["market_comment_edit"])
	assert.Equal(t, true, b.Params["market_comment_delete"])
	assert.Equal(t, true, b.Params["market_comment_restore"])
	assert.Equal(t, true, b.Params["market_order_new"])
	assert.Equal(t, true, b.Params["market_order_edit"])
	assert.Equal(t, true, b.Params["poll_vote_new"])
	assert.Equal(t, true, b.Params["group_join"])
	assert.Equal(t, true, b.Params["group_leave"])
	assert.Equal(t, true, b.Params["group_change_settings"])
	assert.Equal(t, true, b.Params["group_change_photo"])
	assert.Equal(t, true, b.Params["group_officers_edit"])
	assert.Equal(t, true, b.Params["user_block"])
	assert.Equal(t, true, b.Params["user_unblock"])
	assert.Equal(t, true, b.Params["lead_forms_new"])
	assert.Equal(t, true, b.Params["like_add"])
	assert.Equal(t, true, b.Params["like_remove"])
	assert.Equal(t, true, b.Params["donut_subscription_create"])
	assert.Equal(t, true, b.Params["donut_subscription_prolonged"])
	assert.Equal(t, true, b.Params["donut_subscription_expired"])
	assert.Equal(t, true, b.Params["donut_subscription_cancelled"])
	assert.Equal(t, true, b.Params["donut_subscription_price_changed"])
	assert.Equal(t, true, b.Params["donut_money_withdraw"])
	assert.Equal(t, true, b.Params["donut_money_withdraw_error"])
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

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, true, b.Params["enabled"])
	assert.Equal(t, "text", b.Params["api_version"])
	assert.Equal(t, true, b.Params["message_new"])
	assert.Equal(t, true, b.Params["message_reply"])
	assert.Equal(t, true, b.Params["message_allow"])
	assert.Equal(t, true, b.Params["message_deny"])
	assert.Equal(t, true, b.Params["message_edit"])
	assert.Equal(t, true, b.Params["message_typing_state"])
	assert.Equal(t, true, b.Params["photo_new"])
	assert.Equal(t, true, b.Params["audio_new"])
	assert.Equal(t, true, b.Params["video_new"])
	assert.Equal(t, true, b.Params["wall_reply_new"])
	assert.Equal(t, true, b.Params["wall_reply_edit"])
	assert.Equal(t, true, b.Params["wall_reply_delete"])
	assert.Equal(t, true, b.Params["wall_reply_restore"])
	assert.Equal(t, true, b.Params["wall_post_new"])
	assert.Equal(t, true, b.Params["wall_repost"])
	assert.Equal(t, true, b.Params["board_post_new"])
	assert.Equal(t, true, b.Params["board_post_edit"])
	assert.Equal(t, true, b.Params["board_post_restore"])
	assert.Equal(t, true, b.Params["board_post_delete"])
	assert.Equal(t, true, b.Params["photo_comment_new"])
	assert.Equal(t, true, b.Params["photo_comment_edit"])
	assert.Equal(t, true, b.Params["photo_comment_delete"])
	assert.Equal(t, true, b.Params["photo_comment_restore"])
	assert.Equal(t, true, b.Params["video_comment_new"])
	assert.Equal(t, true, b.Params["video_comment_edit"])
	assert.Equal(t, true, b.Params["video_comment_delete"])
	assert.Equal(t, true, b.Params["video_comment_restore"])
	assert.Equal(t, true, b.Params["market_comment_new"])
	assert.Equal(t, true, b.Params["market_comment_edit"])
	assert.Equal(t, true, b.Params["market_comment_delete"])
	assert.Equal(t, true, b.Params["market_comment_restore"])
	assert.Equal(t, true, b.Params["market_order_new"])
	assert.Equal(t, true, b.Params["market_order_edit"])
	assert.Equal(t, true, b.Params["poll_vote_new"])
	assert.Equal(t, true, b.Params["group_join"])
	assert.Equal(t, true, b.Params["group_leave"])
	assert.Equal(t, true, b.Params["group_change_settings"])
	assert.Equal(t, true, b.Params["group_change_photo"])
	assert.Equal(t, true, b.Params["group_officers_edit"])
	assert.Equal(t, true, b.Params["user_block"])
	assert.Equal(t, true, b.Params["user_unblock"])
	assert.Equal(t, true, b.Params["like_add"])
	assert.Equal(t, true, b.Params["like_remove"])
	assert.Equal(t, true, b.Params["donut_subscription_create"])
	assert.Equal(t, true, b.Params["donut_subscription_prolonged"])
	assert.Equal(t, true, b.Params["donut_subscription_expired"])
	assert.Equal(t, true, b.Params["donut_subscription_cancelled"])
	assert.Equal(t, true, b.Params["donut_subscription_price_changed"])
	assert.Equal(t, true, b.Params["donut_money_withdraw"])
	assert.Equal(t, true, b.Params["donut_money_withdraw_error"])
}

func TestGroupsSetUserNoteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsSetUserNoteBuilder()

	b.GroupID(1)
	b.UserID(1)
	b.Note("text")

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, "text", b.Params["note"])
}

func TestGroupsTagAddBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsTagAddBuilder()

	b.GroupID(1)
	b.TagName("text")
	b.TagColor("text")

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, "text", b.Params["tag_name"])
	assert.Equal(t, "text", b.Params["tag_color"])
}

func TestGroupsTagBindBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsTagBindBuilder()

	b.GroupID(1)
	b.TagID(1)
	b.UserID(1)
	b.Act("text")

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["tag_id"])
	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, "text", b.Params["act"])

	b.Bind()
	assert.Equal(t, "bind", b.Params["act"])

	b.Unbind()
	assert.Equal(t, "unbind", b.Params["act"])
}

func TestGroupsTagDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsTagDeleteBuilder()

	b.GroupID(1)
	b.TagID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["tag_id"])
}

func TestGroupsTagUpdateBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsTagUpdateBuilder()

	b.GroupID(1)
	b.TagID(1)
	b.TagName("text")

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["tag_id"])
	assert.Equal(t, "text", b.Params["tag_name"])
}

func TestGroupsToggleMarketBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsToggleMarketBuilder()

	b.GroupID(1)
	b.State("none")

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, params.GroupMarketState("none"), b.Params["state"])
}

func TestGroupsUnbanBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGroupsUnbanBuilder()

	b.GroupID(1)
	b.OwnerID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["owner_id"])
}
