package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// GroupsAddAddressBuilder builder
//
// https://vk.com/dev/groups.addAddress
type GroupsAddAddressBuilder struct {
	api.Params
}

// NewGroupsAddAddressBuilder func
func NewGroupsAddAddressBuilder() *GroupsAddAddressBuilder {
	return &GroupsAddAddressBuilder{api.Params{}}
}

// GroupID parameter
func (b *GroupsAddAddressBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Title parameter
func (b *GroupsAddAddressBuilder) Title(v string) {
	b.Params["title"] = v
}

// Address parameter
func (b *GroupsAddAddressBuilder) Address(v string) {
	b.Params["address"] = v
}

// AdditionalAddress parameter
func (b *GroupsAddAddressBuilder) AdditionalAddress(v string) {
	b.Params["additional_address"] = v
}

// CountryID parameter
func (b *GroupsAddAddressBuilder) CountryID(v int) {
	b.Params["country_id"] = v
}

// CityID parameter
func (b *GroupsAddAddressBuilder) CityID(v int) {
	b.Params["city_id"] = v
}

// MetroID parameter
func (b *GroupsAddAddressBuilder) MetroID(v int) {
	b.Params["metro_id"] = v
}

// Latitude parameter
func (b *GroupsAddAddressBuilder) Latitude(v float64) {
	b.Params["latitude"] = v
}

// Longitude parameter
func (b *GroupsAddAddressBuilder) Longitude(v float64) {
	b.Params["longitude"] = v
}

// Phone parameter
func (b *GroupsAddAddressBuilder) Phone(v string) {
	b.Params["phone"] = v
}

// WorkInfoStatus parameter
func (b *GroupsAddAddressBuilder) WorkInfoStatus(v string) {
	b.Params["work_info_status"] = v
}

// Timetable parameter
func (b *GroupsAddAddressBuilder) Timetable(v string) {
	b.Params["timetable"] = v
}

// IsMainAddress parameter
func (b *GroupsAddAddressBuilder) IsMainAddress(v bool) {
	b.Params["is_main_address"] = v
}

// GroupsAddCallbackServerBuilder builder
//
// https://vk.com/dev/groups.addCallbackServer
type GroupsAddCallbackServerBuilder struct {
	api.Params
}

// NewGroupsAddCallbackServerBuilder func
func NewGroupsAddCallbackServerBuilder() *GroupsAddCallbackServerBuilder {
	return &GroupsAddCallbackServerBuilder{api.Params{}}
}

// GroupID parameter
func (b *GroupsAddCallbackServerBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// URL parameter
func (b *GroupsAddCallbackServerBuilder) URL(v string) {
	b.Params["url"] = v
}

// Title parameter
func (b *GroupsAddCallbackServerBuilder) Title(v string) {
	b.Params["title"] = v
}

// SecretKey parameter
func (b *GroupsAddCallbackServerBuilder) SecretKey(v string) {
	b.Params["secret_key"] = v
}

// GroupsAddLinkBuilder builder
//
// Allows to add a link to the community.
//
// https://vk.com/dev/groups.addLink
type GroupsAddLinkBuilder struct {
	api.Params
}

// NewGroupsAddLinkBuilder func
func NewGroupsAddLinkBuilder() *GroupsAddLinkBuilder {
	return &GroupsAddLinkBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsAddLinkBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Link Link URL.
func (b *GroupsAddLinkBuilder) Link(v string) {
	b.Params["link"] = v
}

// Text Description text for the link.
func (b *GroupsAddLinkBuilder) Text(v string) {
	b.Params["text"] = v
}

// GroupsApproveRequestBuilder builder
//
// Allows to approve join request to the community.
//
// https://vk.com/dev/groups.approveRequest
type GroupsApproveRequestBuilder struct {
	api.Params
}

// NewGroupsApproveRequestBuilder func
func NewGroupsApproveRequestBuilder() *GroupsApproveRequestBuilder {
	return &GroupsApproveRequestBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsApproveRequestBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID User ID.
func (b *GroupsApproveRequestBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupsBanBuilder builder
//
// https://vk.com/dev/groups.ban
type GroupsBanBuilder struct {
	api.Params
}

// NewGroupsBanBuilder func
func NewGroupsBanBuilder() *GroupsBanBuilder {
	return &GroupsBanBuilder{api.Params{}}
}

// GroupID parameter
func (b *GroupsBanBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// OwnerID parameter
func (b *GroupsBanBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// EndDate parameter
func (b *GroupsBanBuilder) EndDate(v int) {
	b.Params["end_date"] = v
}

// Reason parameter
func (b *GroupsBanBuilder) Reason(v int) {
	b.Params["reason"] = v
}

// Comment parameter
func (b *GroupsBanBuilder) Comment(v string) {
	b.Params["comment"] = v
}

// CommentVisible parameter
func (b *GroupsBanBuilder) CommentVisible(v bool) {
	b.Params["comment_visible"] = v
}

// GroupsCreateBuilder builder
//
// Creates a new community.
//
// https://vk.com/dev/groups.create
type GroupsCreateBuilder struct {
	api.Params
}

// NewGroupsCreateBuilder func
func NewGroupsCreateBuilder() *GroupsCreateBuilder {
	return &GroupsCreateBuilder{api.Params{}}
}

// Title Community title.
func (b *GroupsCreateBuilder) Title(v string) {
	b.Params["title"] = v
}

// Description Community description (ignored for 'type' = 'public').
func (b *GroupsCreateBuilder) Description(v string) {
	b.Params["description"] = v
}

// Type Community type. Possible values:
//
// * group – group;
//
// * eent' – event;
//
// * pblic' – public page
func (b *GroupsCreateBuilder) Type(v string) {
	b.Params["type"] = v
}

// PublicCategory Category ID (for 'type' = 'public' only).
func (b *GroupsCreateBuilder) PublicCategory(v int) {
	b.Params["public_category"] = v
}

// Subtype Public page subtype. Possible values:
//
// * 1 – place or small business;
//
// * 2 – company, organization or website;
//
// * 3 – famous person or group of people;
//
// * 4 – product or work of art.
func (b *GroupsCreateBuilder) Subtype(v int) {
	b.Params["subtype"] = v
}

// GroupsDeleteCallbackServerBuilder builder
//
// https://vk.com/dev/groups.deleteCallbackServer
type GroupsDeleteCallbackServerBuilder struct {
	api.Params
}

// NewGroupsDeleteCallbackServerBuilder func
func NewGroupsDeleteCallbackServerBuilder() *GroupsDeleteCallbackServerBuilder {
	return &GroupsDeleteCallbackServerBuilder{api.Params{}}
}

// GroupID parameter
func (b *GroupsDeleteCallbackServerBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// ServerID parameter
func (b *GroupsDeleteCallbackServerBuilder) ServerID(v int) {
	b.Params["server_id"] = v
}

// GroupsDeleteLinkBuilder builder
//
// Allows to delete a link from the community.
//
// https://vk.com/dev/groups.deleteLink
type GroupsDeleteLinkBuilder struct {
	api.Params
}

// NewGroupsDeleteLinkBuilder func
func NewGroupsDeleteLinkBuilder() *GroupsDeleteLinkBuilder {
	return &GroupsDeleteLinkBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsDeleteLinkBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// LinkID Link ID.
func (b *GroupsDeleteLinkBuilder) LinkID(v int) {
	b.Params["link_id"] = v
}

// GroupsDisableOnlineBuilder builder
//
// https://vk.com/dev/groups.disableOnline
type GroupsDisableOnlineBuilder struct {
	api.Params
}

// NewGroupsDisableOnlineBuilder func
func NewGroupsDisableOnlineBuilder() *GroupsDisableOnlineBuilder {
	return &GroupsDisableOnlineBuilder{api.Params{}}
}

// GroupID parameter
func (b *GroupsDisableOnlineBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// GroupsEditBuilder builder
//
// Edits a community.
//
// https://vk.com/dev/groups.edit
type GroupsEditBuilder struct {
	api.Params
}

// NewGroupsEditBuilder func
func NewGroupsEditBuilder() *GroupsEditBuilder {
	return &GroupsEditBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsEditBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Title Community title.
func (b *GroupsEditBuilder) Title(v string) {
	b.Params["title"] = v
}

// Description Community description.
func (b *GroupsEditBuilder) Description(v string) {
	b.Params["description"] = v
}

// ScreenName Community screen name.
func (b *GroupsEditBuilder) ScreenName(v string) {
	b.Params["screen_name"] = v
}

// Access Community type. Possible values:
//
// * 0 – open;
//
// * 1 – closed;
//
// * 2 – private.
func (b *GroupsEditBuilder) Access(v int) {
	b.Params["access"] = v
}

// Website Website that will be displayed in the community information field.
func (b *GroupsEditBuilder) Website(v string) {
	b.Params["website"] = v
}

// Subject Community subject. Possible values:
//
// * 1 – auto/moto,
//
// * 2 – activity holidays,
//
// * 3 – business,
//
// * 4 – pets,
//
// * 5 – health,
//
// * 6 – dating and communication,
//
// * 7 – games,
//
// * 8 – IT (computers and software),
//
// * 9 – cinema,
//
// * 10 – beauty and fashion,
//
// * 11 – cooking,
//
// * 12 – art and culture,
//
// * 13 – literature,
//
// * 14 – mobile services and internet,
//
// * 15 – music,
//
// * 16 – science and technology,
//
// * 17 – real estate,
//
// * 18 – news and media,
//
// * 19 – security,
//
// * 20 – education,
//
// * 21 – home and renovations,
//
// * 22 – politics,
//
// * 23 – food,
//
// * 24 – industry,
//
// * 25 – travel,
//
// * 26 – work,
//
// * 27 – entertainment,
//
// * 28 – religion,
//
// * 29 – family,
//
// * 30 – sports,
//
// * 31 – insurance,
//
// * 32 – television,
//
// * 33 – goods and services,
//
// * 34 – hobbies,
//
// * 35 – finance,
//
// * 36 – photo,
//
// * 37 – esoterics,
//
// * 38 – electronics and appliances,
//
// * 39 – erotic,
//
// * 40 – humor,
//
// * 41 – society, humanities,
//
// * 42 – design and graphics.
func (b *GroupsEditBuilder) Subject(v string) {
	b.Params["subject"] = v
}

// Email Organizer email (for events).
func (b *GroupsEditBuilder) Email(v string) {
	b.Params["email"] = v
}

// Phone Organizer phone number (for events).
func (b *GroupsEditBuilder) Phone(v string) {
	b.Params["phone"] = v
}

// Rss RSS feed address for import (available only to communities with special permission.
// Contact vk.com/support to get it.
func (b *GroupsEditBuilder) Rss(v string) {
	b.Params["rss"] = v
}

// EventStartDate Event start date in Unixtime format.
func (b *GroupsEditBuilder) EventStartDate(v int) {
	b.Params["event_start_date"] = v
}

// EventFinishDate Event finish date in Unixtime format.
func (b *GroupsEditBuilder) EventFinishDate(v int) {
	b.Params["event_finish_date"] = v
}

// EventGroupID Organizer community ID (for events only).
func (b *GroupsEditBuilder) EventGroupID(v int) {
	b.Params["event_group_id"] = v
}

// PublicCategory Public page category ID.
func (b *GroupsEditBuilder) PublicCategory(v int) {
	b.Params["public_category"] = v
}

// PublicSubcategory Public page subcategory ID.
func (b *GroupsEditBuilder) PublicSubcategory(v int) {
	b.Params["public_subcategory"] = v
}

// PublicDate Founding date of a company or organization owning the community in "dd.mm.YYYY" format.
func (b *GroupsEditBuilder) PublicDate(v string) {
	b.Params["public_date"] = v
}

// Wall Wall settings. Possible values:
//
// * 0 – disabled;
//
// * 1 – open;
//
// * 2 – limited (groups and events only);
//
// * 3 – closed (groups and events only).
func (b *GroupsEditBuilder) Wall(v int) {
	b.Params["wall"] = v
}

// Topics Board topics settings. Possbile values:
//
// * 0 – disabled;
//
// * 1 – open;
//
// * 2 – limited (for groups and events only).
func (b *GroupsEditBuilder) Topics(v int) {
	b.Params["topics"] = v
}

// Photos Photos settings. Possible values:
//
// * 0 – disabled;
//
// * 1 – open;
//
// * 2 – limited (for groups and events only).
func (b *GroupsEditBuilder) Photos(v int) {
	b.Params["photos"] = v
}

// Video Video settings. Possible values:
//
// * 0 – disabled;
//
// * 1 – open;
//
// * 2 – limited (for groups and events only).
func (b *GroupsEditBuilder) Video(v int) {
	b.Params["video"] = v
}

// Audio Audio settings. Possible values:
//
// * 0 – disabled;
//
// * 1 – open;
//
// * 2 – limited (for groups and events only).
func (b *GroupsEditBuilder) Audio(v int) {
	b.Params["audio"] = v
}

// Links Links settings (for public pages only). Possible values:
//
// * 0 – disabled;
//
// * 1 – enabled.
func (b *GroupsEditBuilder) Links(v bool) {
	b.Params["links"] = v
}

// Events Events settings (for public pages only). Possible values:
//
// * 0 – disabled;
//
// * 1 – enabled.
func (b *GroupsEditBuilder) Events(v bool) {
	b.Params["events"] = v
}

// Places Places settings (for public pages only). Possible values:
//
// * 0 – disabled;
//
// * 1 – enabled.
func (b *GroupsEditBuilder) Places(v bool) {
	b.Params["places"] = v
}

// Contacts Contacts settings (for public pages only). Possible values:
//
// * 0 – disabled;
//
// * 1 – enabled.
func (b *GroupsEditBuilder) Contacts(v bool) {
	b.Params["contacts"] = v
}

// Docs Documents settings. Possible values:
//
// * 0 – disabled;
//
// * 1 – open;
//
// * 2 – limited (for groups and events only).
func (b *GroupsEditBuilder) Docs(v int) {
	b.Params["docs"] = v
}

// Wiki Wiki pages settings. Possible values:
//
// * 0 – disabled;
//
// * 1 – open;
//
// * 2 – limited (for groups and events only).
func (b *GroupsEditBuilder) Wiki(v int) {
	b.Params["wiki"] = v
}

// Messages Community messages. Possible values:
//
// * 0 — disabled;
//
// * 1 — enabled.
func (b *GroupsEditBuilder) Messages(v bool) {
	b.Params["messages"] = v
}

// Articles parameter
func (b *GroupsEditBuilder) Articles(v bool) {
	b.Params["articles"] = v
}

// Addresses parameter
func (b *GroupsEditBuilder) Addresses(v bool) {
	b.Params["addresses"] = v
}

// AgeLimits Community age limits. Possible values:
//
// * 1 — no limits;
//
// * 2 — 16+;
//
// * 3 — 18+.
func (b *GroupsEditBuilder) AgeLimits(v int) {
	b.Params["age_limits"] = v
}

// Market Market settings. Possible values:
//
// * 0 – disabled;
//
// * 1 – enabled.
func (b *GroupsEditBuilder) Market(v bool) {
	b.Params["market"] = v
}

// MarketComments market comments settings. Possible values:
//
// * 0 – disabled;
//
// * 1 – enabled.
func (b *GroupsEditBuilder) MarketComments(v bool) {
	b.Params["market_comments"] = v
}

// MarketCountry Market delivery countries.
func (b *GroupsEditBuilder) MarketCountry(v []int) {
	b.Params["market_country"] = v
}

// MarketCity Market delivery cities (if only one country is specified).
func (b *GroupsEditBuilder) MarketCity(v []int) {
	b.Params["market_city"] = v
}

// MarketCurrency Market currency settings. Possbile values:
//
// * 643 – Russian rubles,
//
// * 980 – Ukrainian hryvnia,
//
// * 398 – Kazakh tenge,
//
// * 978 – Euro,
//
// * 840 – US dollars
func (b *GroupsEditBuilder) MarketCurrency(v int) {
	b.Params["market_currency"] = v
}

// MarketContact Seller contact for market. Set '0' for community messages.
func (b *GroupsEditBuilder) MarketContact(v int) {
	b.Params["market_contact"] = v
}

// MarketWiki ID of a wiki page with market description.
func (b *GroupsEditBuilder) MarketWiki(v int) {
	b.Params["market_wiki"] = v
}

// ObsceneFilter Obscene expressions filter in comments. Possible values:
//
// * 0 – disabled;
//
// * 1 – enabled.
func (b *GroupsEditBuilder) ObsceneFilter(v bool) {
	b.Params["obscene_filter"] = v
}

// ObsceneStopwords Stopwords filter in comments. Possible values:
//
// * 0 – disabled;
//
// * 1 – enabled.
func (b *GroupsEditBuilder) ObsceneStopwords(v bool) {
	b.Params["obscene_stopwords"] = v
}

// ObsceneWords Keywords for stopwords filter.
func (b *GroupsEditBuilder) ObsceneWords(v []string) {
	b.Params["obscene_words"] = v
}

// MainSection parameter
func (b *GroupsEditBuilder) MainSection(v int) {
	b.Params["main_section"] = v
}

// SecondarySection parameter
func (b *GroupsEditBuilder) SecondarySection(v int) {
	b.Params["secondary_section"] = v
}

// Country Country of the community.
func (b *GroupsEditBuilder) Country(v int) {
	b.Params["country"] = v
}

// City City of the community.
func (b *GroupsEditBuilder) City(v int) {
	b.Params["city"] = v
}

// GroupsEditAddressBuilder builder
//
// https://vk.com/dev/groups.editAddress
type GroupsEditAddressBuilder struct {
	api.Params
}

// NewGroupsEditAddressBuilder func
func NewGroupsEditAddressBuilder() *GroupsEditAddressBuilder {
	return &GroupsEditAddressBuilder{api.Params{}}
}

// GroupID parameter
func (b *GroupsEditAddressBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// AddressID parameter
func (b *GroupsEditAddressBuilder) AddressID(v int) {
	b.Params["address_id"] = v
}

// Title parameter
func (b *GroupsEditAddressBuilder) Title(v string) {
	b.Params["title"] = v
}

// Address parameter
func (b *GroupsEditAddressBuilder) Address(v string) {
	b.Params["address"] = v
}

// AdditionalAddress parameter
func (b *GroupsEditAddressBuilder) AdditionalAddress(v string) {
	b.Params["additional_address"] = v
}

// CountryID parameter
func (b *GroupsEditAddressBuilder) CountryID(v int) {
	b.Params["country_id"] = v
}

// CityID parameter
func (b *GroupsEditAddressBuilder) CityID(v int) {
	b.Params["city_id"] = v
}

// MetroID parameter
func (b *GroupsEditAddressBuilder) MetroID(v int) {
	b.Params["metro_id"] = v
}

// Latitude parameter
func (b *GroupsEditAddressBuilder) Latitude(v float64) {
	b.Params["latitude"] = v
}

// Longitude parameter
func (b *GroupsEditAddressBuilder) Longitude(v float64) {
	b.Params["longitude"] = v
}

// Phone parameter
func (b *GroupsEditAddressBuilder) Phone(v string) {
	b.Params["phone"] = v
}

// WorkInfoStatus parameter
func (b *GroupsEditAddressBuilder) WorkInfoStatus(v string) {
	b.Params["work_info_status"] = v
}

// Timetable parameter
func (b *GroupsEditAddressBuilder) Timetable(v string) {
	b.Params["timetable"] = v
}

// IsMainAddress parameter
func (b *GroupsEditAddressBuilder) IsMainAddress(v bool) {
	b.Params["is_main_address"] = v
}

// GroupsEditCallbackServerBuilder builder
//
// https://vk.com/dev/groups.editCallbackServer
type GroupsEditCallbackServerBuilder struct {
	api.Params
}

// NewGroupsEditCallbackServerBuilder func
func NewGroupsEditCallbackServerBuilder() *GroupsEditCallbackServerBuilder {
	return &GroupsEditCallbackServerBuilder{api.Params{}}
}

// GroupID parameter
func (b *GroupsEditCallbackServerBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// ServerID parameter
func (b *GroupsEditCallbackServerBuilder) ServerID(v int) {
	b.Params["server_id"] = v
}

// URL parameter
func (b *GroupsEditCallbackServerBuilder) URL(v string) {
	b.Params["url"] = v
}

// Title parameter
func (b *GroupsEditCallbackServerBuilder) Title(v string) {
	b.Params["title"] = v
}

// SecretKey parameter
func (b *GroupsEditCallbackServerBuilder) SecretKey(v string) {
	b.Params["secret_key"] = v
}

// GroupsEditLinkBuilder builder
//
// Allows to edit a link in the community.
//
// https://vk.com/dev/groups.editLink
type GroupsEditLinkBuilder struct {
	api.Params
}

// NewGroupsEditLinkBuilder func
func NewGroupsEditLinkBuilder() *GroupsEditLinkBuilder {
	return &GroupsEditLinkBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsEditLinkBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// LinkID Link ID.
func (b *GroupsEditLinkBuilder) LinkID(v int) {
	b.Params["link_id"] = v
}

// Text New description text for the link.
func (b *GroupsEditLinkBuilder) Text(v string) {
	b.Params["text"] = v
}

// GroupsEditManagerBuilder builder
//
// Allows to add, remove or edit the community manager.
//
// https://vk.com/dev/groups.editManager
type GroupsEditManagerBuilder struct {
	api.Params
}

// NewGroupsEditManagerBuilder func
func NewGroupsEditManagerBuilder() *GroupsEditManagerBuilder {
	return &GroupsEditManagerBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsEditManagerBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID User ID.
func (b *GroupsEditManagerBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// Role Manager role. Possible values:
//
// * moderator;
//
// * eitor';
//
// * aministrator'.
func (b *GroupsEditManagerBuilder) Role(v string) {
	b.Params["role"] = v
}

// IsContact '1' — to show the manager in Contacts block of the community.
func (b *GroupsEditManagerBuilder) IsContact(v bool) {
	b.Params["is_contact"] = v
}

// ContactPosition Position to show in Contacts block.
func (b *GroupsEditManagerBuilder) ContactPosition(v string) {
	b.Params["contact_position"] = v
}

// ContactPhone Contact phone.
func (b *GroupsEditManagerBuilder) ContactPhone(v string) {
	b.Params["contact_phone"] = v
}

// ContactEmail Contact e-mail.
func (b *GroupsEditManagerBuilder) ContactEmail(v string) {
	b.Params["contact_email"] = v
}

// GroupsEnableOnlineBuilder builder
//
// https://vk.com/dev/groups.enableOnline
type GroupsEnableOnlineBuilder struct {
	api.Params
}

// NewGroupsEnableOnlineBuilder func
func NewGroupsEnableOnlineBuilder() *GroupsEnableOnlineBuilder {
	return &GroupsEnableOnlineBuilder{api.Params{}}
}

// GroupID parameter
func (b *GroupsEnableOnlineBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// GroupsGetBuilder builder
//
// Returns a list of the communities to which a user belongs.
//
// https://vk.com/dev/groups.get
type GroupsGetBuilder struct {
	api.Params
}

// NewGroupsGetBuilder func
func NewGroupsGetBuilder() *GroupsGetBuilder {
	return &GroupsGetBuilder{api.Params{}}
}

// UserID User ID.
func (b *GroupsGetBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// Extended parameter
//
// * 1 — to return complete information about a user's communities,
//
// * 0 — to return a list of community IDs without any additional fields (default),
func (b *GroupsGetBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Filter Types of communities to return:
//
// 'admin' — to return communities administered by the user
//
// 'editor' — to return communities where the user is an administrator or editor,
//
// 'moder' — to return communities where the user is an administrator, editor, or moderator,
//
// 'groups' — to return only groups,
//
// 'publics' — to return only public pages,
//
// 'events' — to return only events
func (b *GroupsGetBuilder) Filter(v []string) {
	b.Params["filter"] = v
}

// Fields Profile fields to return.
func (b *GroupsGetBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// Offset Offset needed to return a specific subset of communities.
func (b *GroupsGetBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of communities to return.
func (b *GroupsGetBuilder) Count(v int) {
	b.Params["count"] = v
}

// GroupsGetAddressesBuilder builder
//
// Returns a list of community addresses.
//
// https://vk.com/dev/groups.getAddresses
type GroupsGetAddressesBuilder struct {
	api.Params
}

// NewGroupsGetAddressesBuilder func
func NewGroupsGetAddressesBuilder() *GroupsGetAddressesBuilder {
	return &GroupsGetAddressesBuilder{api.Params{}}
}

// GroupID ID or screen name of the community.
func (b *GroupsGetAddressesBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// AddressIDs parameter
func (b *GroupsGetAddressesBuilder) AddressIDs(v []int) {
	b.Params["address_ids"] = v
}

// Latitude Latitude of  the user geo position.
func (b *GroupsGetAddressesBuilder) Latitude(v float64) {
	b.Params["latitude"] = v
}

// Longitude Longitude of the user geo position.
func (b *GroupsGetAddressesBuilder) Longitude(v float64) {
	b.Params["longitude"] = v
}

// Offset Offset needed to return a specific subset of community addresses.
func (b *GroupsGetAddressesBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of community addresses to return.
func (b *GroupsGetAddressesBuilder) Count(v int) {
	b.Params["count"] = v
}

// Fields Address fields
func (b *GroupsGetAddressesBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupsGetBannedBuilder builder
//
// Returns a list of users on a community blacklist.
//
// https://vk.com/dev/groups.getBanned
type GroupsGetBannedBuilder struct {
	api.Params
}

// NewGroupsGetBannedBuilder func
func NewGroupsGetBannedBuilder() *GroupsGetBannedBuilder {
	return &GroupsGetBannedBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsGetBannedBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Offset Offset needed to return a specific subset of users.
func (b *GroupsGetBannedBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of users to return.
func (b *GroupsGetBannedBuilder) Count(v int) {
	b.Params["count"] = v
}

// Fields parameter
func (b *GroupsGetBannedBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// OwnerID parameter
func (b *GroupsGetBannedBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// GroupsGetByIDBuilder builder
//
// Returns information about communities by their IDs.
//
// https://vk.com/dev/groups.getById
type GroupsGetByIDBuilder struct {
	api.Params
}

// NewGroupsGetByIDBuilder func
func NewGroupsGetByIDBuilder() *GroupsGetByIDBuilder {
	return &GroupsGetByIDBuilder{api.Params{}}
}

// GroupIDs IDs or screen names of communities.
func (b *GroupsGetByIDBuilder) GroupIDs(v []string) {
	b.Params["group_ids"] = v
}

// GroupID ID or screen name of the community.
func (b *GroupsGetByIDBuilder) GroupID(v string) {
	b.Params["group_id"] = v
}

// Fields Group fields to return.
func (b *GroupsGetByIDBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupsGetCallbackConfirmationCodeBuilder builder
//
// Returns Callback API confirmation code for the community.
//
// https://vk.com/dev/groups.getCallbackConfirmationCode
type GroupsGetCallbackConfirmationCodeBuilder struct {
	api.Params
}

// NewGroupsGetCallbackConfirmationCodeBuilder func
func NewGroupsGetCallbackConfirmationCodeBuilder() *GroupsGetCallbackConfirmationCodeBuilder {
	return &GroupsGetCallbackConfirmationCodeBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsGetCallbackConfirmationCodeBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// GroupsGetCallbackServersBuilder builder
//
// https://vk.com/dev/groups.getCallbackServers
type GroupsGetCallbackServersBuilder struct {
	api.Params
}

// NewGroupsGetCallbackServersBuilder func
func NewGroupsGetCallbackServersBuilder() *GroupsGetCallbackServersBuilder {
	return &GroupsGetCallbackServersBuilder{api.Params{}}
}

// GroupID parameter
func (b *GroupsGetCallbackServersBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// ServerIDs parameter
func (b *GroupsGetCallbackServersBuilder) ServerIDs(v []int) {
	b.Params["server_ids"] = v
}

// GroupsGetCallbackSettingsBuilder builder
//
// Returns [vk.com/dev/callback_api|Callback API] notifications settings.
//
// https://vk.com/dev/groups.getCallbackSettings
type GroupsGetCallbackSettingsBuilder struct {
	api.Params
}

// NewGroupsGetCallbackSettingsBuilder func
func NewGroupsGetCallbackSettingsBuilder() *GroupsGetCallbackSettingsBuilder {
	return &GroupsGetCallbackSettingsBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsGetCallbackSettingsBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// ServerID Server ID.
func (b *GroupsGetCallbackSettingsBuilder) ServerID(v int) {
	b.Params["server_id"] = v
}

// GroupsGetCatalogBuilder builder
//
// Returns communities list for a catalog category.
//
// https://vk.com/dev/groups.getCatalog
type GroupsGetCatalogBuilder struct {
	api.Params
}

// NewGroupsGetCatalogBuilder func
func NewGroupsGetCatalogBuilder() *GroupsGetCatalogBuilder {
	return &GroupsGetCatalogBuilder{api.Params{}}
}

// CategoryID Category id received from [vk.com/dev/groups.getCatalogInfo|groups.getCatalogInfo].
func (b *GroupsGetCatalogBuilder) CategoryID(v int) {
	b.Params["category_id"] = v
}

// SubcategoryID Subcategory id received from [vk.com/dev/groups.getCatalogInfo|groups.getCatalogInfo].
func (b *GroupsGetCatalogBuilder) SubcategoryID(v int) {
	b.Params["subcategory_id"] = v
}

// GroupsGetCatalogInfoBuilder builder
//
// Returns categories list for communities catalog
//
// https://vk.com/dev/groups.getCatalogInfo
type GroupsGetCatalogInfoBuilder struct {
	api.Params
}

// NewGroupsGetCatalogInfoBuilder func
func NewGroupsGetCatalogInfoBuilder() *GroupsGetCatalogInfoBuilder {
	return &GroupsGetCatalogInfoBuilder{api.Params{}}
}

// Extended 1 – to return communities count and three communities for preview. By default: 0.
func (b *GroupsGetCatalogInfoBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Subcategories 1 – to return subcategories info. By default: 0.
func (b *GroupsGetCatalogInfoBuilder) Subcategories(v bool) {
	b.Params["subcategories"] = v
}

// GroupsGetInvitedUsersBuilder builder
//
// Returns invited users list of a community
//
// https://vk.com/dev/groups.getInvitedUsers
type GroupsGetInvitedUsersBuilder struct {
	api.Params
}

// NewGroupsGetInvitedUsersBuilder func
func NewGroupsGetInvitedUsersBuilder() *GroupsGetInvitedUsersBuilder {
	return &GroupsGetInvitedUsersBuilder{api.Params{}}
}

// GroupID Group ID to return invited users for.
func (b *GroupsGetInvitedUsersBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *GroupsGetInvitedUsersBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of results to return.
func (b *GroupsGetInvitedUsersBuilder) Count(v int) {
	b.Params["count"] = v
}

// Fields List of additional fields to be returned. Available values: 'sex, bdate, city, country, photo_50,
// photo_100, photo_200_orig, photo_200, photo_400_orig, photo_max, photo_max_orig, online, online_mobile,
// lists, domain, has_mobile, contacts, connections, site, education, universities, schools, can_post,
// can_see_all_posts, can_see_audio, can_write_private_message, status, last_seen, common_count,
// relation, relatives, counters'.
func (b *GroupsGetInvitedUsersBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase Case for declension of user name and surname. Possible values:
//
// * nom — nominative (default);
//
// * gn' — genitive;
//
// * dt' — dative;
//
// * ac' — accusative,
//
// * ins — instrumental;
//
// * al' — prepositional.
func (b *GroupsGetInvitedUsersBuilder) NameCase(v string) {
	b.Params["name_case"] = v
}

// GroupsGetInvitesBuilder builder
//
// Returns a list of invitations to join communities and events.
//
// https://vk.com/dev/groups.getInvites
type GroupsGetInvitesBuilder struct {
	api.Params
}

// NewGroupsGetInvitesBuilder func
func NewGroupsGetInvitesBuilder() *GroupsGetInvitesBuilder {
	return &GroupsGetInvitesBuilder{api.Params{}}
}

// Offset Offset needed to return a specific subset of invitations.
func (b *GroupsGetInvitesBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of invitations to return.
func (b *GroupsGetInvitesBuilder) Count(v int) {
	b.Params["count"] = v
}

// Extended '1' — to return additional [vk.com/dev/fields_groups|fields] for communities..
func (b *GroupsGetInvitesBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// GroupsGetLongPollServerBuilder builder
//
// Returns the data needed to query a Long Poll server for events
//
// https://vk.com/dev/groups.getLongPollServer
type GroupsGetLongPollServerBuilder struct {
	api.Params
}

// NewGroupsGetLongPollServerBuilder func
func NewGroupsGetLongPollServerBuilder() *GroupsGetLongPollServerBuilder {
	return &GroupsGetLongPollServerBuilder{api.Params{}}
}

// GroupID Community ID
func (b *GroupsGetLongPollServerBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// GroupsGetLongPollSettingsBuilder builder
//
// Returns Long Poll notification settings
//
// https://vk.com/dev/groups.getLongPollSettings
type GroupsGetLongPollSettingsBuilder struct {
	api.Params
}

// NewGroupsGetLongPollSettingsBuilder func
func NewGroupsGetLongPollSettingsBuilder() *GroupsGetLongPollSettingsBuilder {
	return &GroupsGetLongPollSettingsBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsGetLongPollSettingsBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// GroupsGetMembersBuilder builder
//
// Returns a list of community members.
//
// https://vk.com/dev/groups.getMembers
type GroupsGetMembersBuilder struct {
	api.Params
}

// NewGroupsGetMembersBuilder func
func NewGroupsGetMembersBuilder() *GroupsGetMembersBuilder {
	return &GroupsGetMembersBuilder{api.Params{}}
}

// GroupID ID or screen name of the community.
func (b *GroupsGetMembersBuilder) GroupID(v string) {
	b.Params["group_id"] = v
}

// Sort Sort order. Available values: 'id_asc', 'id_desc', 'time_asc', 'time_desc'. 'time_asc' and 'time_desc'
// are availavle only if the method is called by the group's 'moderator'.
func (b *GroupsGetMembersBuilder) Sort(v string) {
	b.Params["sort"] = v
}

// Offset Offset needed to return a specific subset of community members.
func (b *GroupsGetMembersBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of community members to return.
func (b *GroupsGetMembersBuilder) Count(v int) {
	b.Params["count"] = v
}

// Fields List of additional fields to be returned. Available values: 'sex, bdate, city, country, photo_50, photo_100,
// photo_200_orig, photo_200, photo_400_orig, photo_max, photo_max_orig, online, online_mobile, lists, domain,
// has_mobile, contacts, connections, site, education, universities, schools, can_post, can_see_all_posts,
// can_see_audio, can_write_private_message, status, last_seen, common_count, relation, relatives, counters'.
func (b *GroupsGetMembersBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// Filter *'friends' – only friends in this community will be returned;
//
// * usure' – only those who pressed 'I may attend' will be returned (if it's an event).
func (b *GroupsGetMembersBuilder) Filter(v string) {
	b.Params["filter"] = v
}

// GroupsGetRequestsBuilder builder
//
// Returns a list of requests to the community.
//
// https://vk.com/dev/groups.getRequests
type GroupsGetRequestsBuilder struct {
	api.Params
}

// NewGroupsGetRequestsBuilder func
func NewGroupsGetRequestsBuilder() *GroupsGetRequestsBuilder {
	return &GroupsGetRequestsBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsGetRequestsBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *GroupsGetRequestsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of results to return.
func (b *GroupsGetRequestsBuilder) Count(v int) {
	b.Params["count"] = v
}

// Fields Profile fields to return.
func (b *GroupsGetRequestsBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupsGetSettingsBuilder builder
//
// Returns community settings.
//
// https://vk.com/dev/groups.getSettings
type GroupsGetSettingsBuilder struct {
	api.Params
}

// NewGroupsGetSettingsBuilder func
func NewGroupsGetSettingsBuilder() *GroupsGetSettingsBuilder {
	return &GroupsGetSettingsBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsGetSettingsBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// GroupsInviteBuilder builder
//
// Allows to invite friends to the community.
//
// https://vk.com/dev/groups.invite
type GroupsInviteBuilder struct {
	api.Params
}

// NewGroupsInviteBuilder func
func NewGroupsInviteBuilder() *GroupsInviteBuilder {
	return &GroupsInviteBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsInviteBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID User ID.
func (b *GroupsInviteBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupsIsMemberBuilder builder
//
// Returns information specifying whether a user is a member of a community.
//
// https://vk.com/dev/groups.isMember
type GroupsIsMemberBuilder struct {
	api.Params
}

// NewGroupsIsMemberBuilder func
func NewGroupsIsMemberBuilder() *GroupsIsMemberBuilder {
	return &GroupsIsMemberBuilder{api.Params{}}
}

// GroupID ID or screen name of the community.
func (b *GroupsIsMemberBuilder) GroupID(v string) {
	b.Params["group_id"] = v
}

// UserID User ID.
func (b *GroupsIsMemberBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// UserIDs User IDs.
func (b *GroupsIsMemberBuilder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// Extended '1' — to return an extended response with additional fields. By default: '0'.
func (b *GroupsIsMemberBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// GroupsJoinBuilder builder
//
// With this method you can join the group or public page, and also confirm your participation in an event.
//
// https://vk.com/dev/groups.join
type GroupsJoinBuilder struct {
	api.Params
}

// NewGroupsJoinBuilder func
func NewGroupsJoinBuilder() *GroupsJoinBuilder {
	return &GroupsJoinBuilder{api.Params{}}
}

// GroupID ID or screen name of the community.
func (b *GroupsJoinBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// NotSure Optional parameter which is taken into account when 'gid' belongs to the event:
//
// * 1 — Perhaps I will attend;
//
// * 0 — I will be there for sure (default).
func (b *GroupsJoinBuilder) NotSure(v string) {
	b.Params["not_sure"] = v
}

// GroupsLeaveBuilder builder
//
// With this method you can leave a group, public page, or event.
//
// https://vk.com/dev/groups.leave
type GroupsLeaveBuilder struct {
	api.Params
}

// NewGroupsLeaveBuilder func
func NewGroupsLeaveBuilder() *GroupsLeaveBuilder {
	return &GroupsLeaveBuilder{api.Params{}}
}

// GroupID ID or screen name of the community.
func (b *GroupsLeaveBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// GroupsRemoveUserBuilder builder
//
// Removes a user from the community.
//
// https://vk.com/dev/groups.removeUser
type GroupsRemoveUserBuilder struct {
	api.Params
}

// NewGroupsRemoveUserBuilder func
func NewGroupsRemoveUserBuilder() *GroupsRemoveUserBuilder {
	return &GroupsRemoveUserBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsRemoveUserBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID User ID.
func (b *GroupsRemoveUserBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupsReorderLinkBuilder builder
//
// Allows to reorder links in the community.
//
// https://vk.com/dev/groups.reorderLink
type GroupsReorderLinkBuilder struct {
	api.Params
}

// NewGroupsReorderLinkBuilder func
func NewGroupsReorderLinkBuilder() *GroupsReorderLinkBuilder {
	return &GroupsReorderLinkBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsReorderLinkBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// LinkID Link ID.
func (b *GroupsReorderLinkBuilder) LinkID(v int) {
	b.Params["link_id"] = v
}

// After ID of the link after which to place the link with 'link_id'.
func (b *GroupsReorderLinkBuilder) After(v int) {
	b.Params["after"] = v
}

// GroupsSearchBuilder builder
//
// Returns a list of communities matching the search criteria.
//
// https://vk.com/dev/groups.search
type GroupsSearchBuilder struct {
	api.Params
}

// NewGroupsSearchBuilder func
func NewGroupsSearchBuilder() *GroupsSearchBuilder {
	return &GroupsSearchBuilder{api.Params{}}
}

// Q Search query string.
func (b *GroupsSearchBuilder) Q(v string) {
	b.Params["q"] = v
}

// Type Community type. Possible values: 'group, page, event.'
func (b *GroupsSearchBuilder) Type(v string) {
	b.Params["type"] = v
}

// CountryID Country ID.
func (b *GroupsSearchBuilder) CountryID(v int) {
	b.Params["country_id"] = v
}

// CityID City ID. If this parameter is transmitted, country_id is ignored.
func (b *GroupsSearchBuilder) CityID(v int) {
	b.Params["city_id"] = v
}

// Future '1' — to return only upcoming events. Works with the 'type' = 'event' only.
func (b *GroupsSearchBuilder) Future(v bool) {
	b.Params["future"] = v
}

// Market '1' — to return communities with enabled market only.
func (b *GroupsSearchBuilder) Market(v bool) {
	b.Params["market"] = v
}

// Sort Sort order. Possible values:
//
// * 0 — default sorting (similar the full version of the site);
//
// * 1 — by growth speed;
//
// * 2— by the "day attendance/members number" ratio;
//
// * 3 — by the "Likes number/members number" ratio;
//
// * 4 — by the "comments number/members number" ratio;
//
// * 5 — by the "boards entries number/members number" ratio.
func (b *GroupsSearchBuilder) Sort(v int) {
	b.Params["sort"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *GroupsSearchBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of communities to return.
// Note that you can not receive more than first thousand of results, regardless of 'count' and 'offset' values.
func (b *GroupsSearchBuilder) Count(v int) {
	b.Params["count"] = v
}

// GroupsSetCallbackSettingsBuilder builder
//
// Allow to set notifications settings for group.
//
// https://vk.com/dev/groups.setCallbackSettings
type GroupsSetCallbackSettingsBuilder struct {
	api.Params
}

// NewGroupsSetCallbackSettingsBuilder func
func NewGroupsSetCallbackSettingsBuilder() *GroupsSetCallbackSettingsBuilder {
	return &GroupsSetCallbackSettingsBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsSetCallbackSettingsBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// ServerID Server ID.
func (b *GroupsSetCallbackSettingsBuilder) ServerID(v int) {
	b.Params["server_id"] = v
}

// APIVersion parameter
func (b *GroupsSetCallbackSettingsBuilder) APIVersion(v string) {
	b.Params["api_version"] = v
}

// MessageNew A new incoming message has been received ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) MessageNew(v bool) {
	b.Params["message_new"] = v
}

// MessageReply A new outcoming message has been received ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) MessageReply(v bool) {
	b.Params["message_reply"] = v
}

// MessageAllow Allowed messages notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) MessageAllow(v bool) {
	b.Params["message_allow"] = v
}

// MessageEdit parameter
func (b *GroupsSetCallbackSettingsBuilder) MessageEdit(v bool) {
	b.Params["message_edit"] = v
}

// MessageDeny Denied messages notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) MessageDeny(v bool) {
	b.Params["message_deny"] = v
}

// MessageTypingState parameter
func (b *GroupsSetCallbackSettingsBuilder) MessageTypingState(v bool) {
	b.Params["message_typing_state"] = v
}

// PhotoNew New photos notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) PhotoNew(v bool) {
	b.Params["photo_new"] = v
}

// AudioNew New audios notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) AudioNew(v bool) {
	b.Params["audio_new"] = v
}

// VideoNew New videos notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) VideoNew(v bool) {
	b.Params["video_new"] = v
}

// WallReplyNew New wall replies notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) WallReplyNew(v bool) {
	b.Params["wall_reply_new"] = v
}

// WallReplyEdit Wall replies edited notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) WallReplyEdit(v bool) {
	b.Params["wall_reply_edit"] = v
}

// WallReplyDelete A wall comment has been deleted ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) WallReplyDelete(v bool) {
	b.Params["wall_reply_delete"] = v
}

// WallReplyRestore A wall comment has been restored ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) WallReplyRestore(v bool) {
	b.Params["wall_reply_restore"] = v
}

// WallPostNew New wall posts notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) WallPostNew(v bool) {
	b.Params["wall_post_new"] = v
}

// WallRepost New wall posts notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) WallRepost(v bool) {
	b.Params["wall_repost"] = v
}

// BoardPostNew New board posts notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) BoardPostNew(v bool) {
	b.Params["board_post_new"] = v
}

// BoardPostEdit Board posts edited notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) BoardPostEdit(v bool) {
	b.Params["board_post_edit"] = v
}

// BoardPostRestore Board posts restored notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) BoardPostRestore(v bool) {
	b.Params["board_post_restore"] = v
}

// BoardPostDelete Board posts deleted notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) BoardPostDelete(v bool) {
	b.Params["board_post_delete"] = v
}

// PhotoCommentNew New comment to photo notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) PhotoCommentNew(v bool) {
	b.Params["photo_comment_new"] = v
}

// PhotoCommentEdit A photo comment has been edited ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) PhotoCommentEdit(v bool) {
	b.Params["photo_comment_edit"] = v
}

// PhotoCommentDelete A photo comment has been deleted ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) PhotoCommentDelete(v bool) {
	b.Params["photo_comment_delete"] = v
}

// PhotoCommentRestore A photo comment has been restored ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) PhotoCommentRestore(v bool) {
	b.Params["photo_comment_restore"] = v
}

// VideoCommentNew New comment to video notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) VideoCommentNew(v bool) {
	b.Params["video_comment_new"] = v
}

// VideoCommentEdit A video comment has been edited ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) VideoCommentEdit(v bool) {
	b.Params["video_comment_edit"] = v
}

// VideoCommentDelete A video comment has been deleted ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) VideoCommentDelete(v bool) {
	b.Params["video_comment_delete"] = v
}

// VideoCommentRestore A video comment has been restored ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) VideoCommentRestore(v bool) {
	b.Params["video_comment_restore"] = v
}

// MarketCommentNew New comment to market item notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) MarketCommentNew(v bool) {
	b.Params["market_comment_new"] = v
}

// MarketCommentEdit A market comment has been edited ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) MarketCommentEdit(v bool) {
	b.Params["market_comment_edit"] = v
}

// MarketCommentDelete A market comment has been deleted ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) MarketCommentDelete(v bool) {
	b.Params["market_comment_delete"] = v
}

// MarketCommentRestore A market comment has been restored ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) MarketCommentRestore(v bool) {
	b.Params["market_comment_restore"] = v
}

// PollVoteNew A vote in a public poll has been added ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) PollVoteNew(v bool) {
	b.Params["poll_vote_new"] = v
}

// GroupJoin Joined community notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) GroupJoin(v bool) {
	b.Params["group_join"] = v
}

// GroupLeave Left community notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBuilder) GroupLeave(v bool) {
	b.Params["group_leave"] = v
}

// GroupChangeSettings parameter
func (b *GroupsSetCallbackSettingsBuilder) GroupChangeSettings(v bool) {
	b.Params["group_change_settings"] = v
}

// GroupChangePhoto parameter
func (b *GroupsSetCallbackSettingsBuilder) GroupChangePhoto(v bool) {
	b.Params["group_change_photo"] = v
}

// GroupOfficersEdit parameter
func (b *GroupsSetCallbackSettingsBuilder) GroupOfficersEdit(v bool) {
	b.Params["group_officers_edit"] = v
}

// UserBlock User added to community blacklist
func (b *GroupsSetCallbackSettingsBuilder) UserBlock(v bool) {
	b.Params["user_block"] = v
}

// UserUnblock User removed from community blacklist
func (b *GroupsSetCallbackSettingsBuilder) UserUnblock(v bool) {
	b.Params["user_unblock"] = v
}

// LeadFormsNew New form in lead forms
func (b *GroupsSetCallbackSettingsBuilder) LeadFormsNew(v bool) {
	b.Params["lead_forms_new"] = v
}

// GroupsSetLongPollSettingsBuilder builder
//
// Sets Long Poll notification settings
//
// https://vk.com/dev/groups.setLongPollSettings
type GroupsSetLongPollSettingsBuilder struct {
	api.Params
}

// NewGroupsSetLongPollSettingsBuilder func
func NewGroupsSetLongPollSettingsBuilder() *GroupsSetLongPollSettingsBuilder {
	return &GroupsSetLongPollSettingsBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsSetLongPollSettingsBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Enabled Sets whether Long Poll is enabled ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) Enabled(v bool) {
	b.Params["enabled"] = v
}

// APIVersion parameter
func (b *GroupsSetLongPollSettingsBuilder) APIVersion(v string) {
	b.Params["api_version"] = v
}

// MessageNew A new incoming message has been received ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) MessageNew(v bool) {
	b.Params["message_new"] = v
}

// MessageReply A new outcoming message has been received ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) MessageReply(v bool) {
	b.Params["message_reply"] = v
}

// MessageAllow Allowed messages notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) MessageAllow(v bool) {
	b.Params["message_allow"] = v
}

// MessageDeny Denied messages notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) MessageDeny(v bool) {
	b.Params["message_deny"] = v
}

// MessageEdit A message has been edited ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) MessageEdit(v bool) {
	b.Params["message_edit"] = v
}

// MessageTypingState parameter
func (b *GroupsSetLongPollSettingsBuilder) MessageTypingState(v bool) {
	b.Params["message_typing_state"] = v
}

// PhotoNew New photos notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) PhotoNew(v bool) {
	b.Params["photo_new"] = v
}

// AudioNew New audios notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) AudioNew(v bool) {
	b.Params["audio_new"] = v
}

// VideoNew New videos notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) VideoNew(v bool) {
	b.Params["video_new"] = v
}

// WallReplyNew New wall replies notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) WallReplyNew(v bool) {
	b.Params["wall_reply_new"] = v
}

// WallReplyEdit Wall replies edited notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) WallReplyEdit(v bool) {
	b.Params["wall_reply_edit"] = v
}

// WallReplyDelete A wall comment has been deleted ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) WallReplyDelete(v bool) {
	b.Params["wall_reply_delete"] = v
}

// WallReplyRestore A wall comment has been restored ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) WallReplyRestore(v bool) {
	b.Params["wall_reply_restore"] = v
}

// WallPostNew New wall posts notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) WallPostNew(v bool) {
	b.Params["wall_post_new"] = v
}

// WallRepost New wall posts notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) WallRepost(v bool) {
	b.Params["wall_repost"] = v
}

// BoardPostNew New board posts notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) BoardPostNew(v bool) {
	b.Params["board_post_new"] = v
}

// BoardPostEdit Board posts edited notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) BoardPostEdit(v bool) {
	b.Params["board_post_edit"] = v
}

// BoardPostRestore Board posts restored notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) BoardPostRestore(v bool) {
	b.Params["board_post_restore"] = v
}

// BoardPostDelete Board posts deleted notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) BoardPostDelete(v bool) {
	b.Params["board_post_delete"] = v
}

// PhotoCommentNew New comment to photo notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) PhotoCommentNew(v bool) {
	b.Params["photo_comment_new"] = v
}

// PhotoCommentEdit A photo comment has been edited ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) PhotoCommentEdit(v bool) {
	b.Params["photo_comment_edit"] = v
}

// PhotoCommentDelete A photo comment has been deleted ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) PhotoCommentDelete(v bool) {
	b.Params["photo_comment_delete"] = v
}

// PhotoCommentRestore A photo comment has been restored ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) PhotoCommentRestore(v bool) {
	b.Params["photo_comment_restore"] = v
}

// VideoCommentNew New comment to video notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) VideoCommentNew(v bool) {
	b.Params["video_comment_new"] = v
}

// VideoCommentEdit A video comment has been edited ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) VideoCommentEdit(v bool) {
	b.Params["video_comment_edit"] = v
}

// VideoCommentDelete A video comment has been deleted ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) VideoCommentDelete(v bool) {
	b.Params["video_comment_delete"] = v
}

// VideoCommentRestore A video comment has been restored ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) VideoCommentRestore(v bool) {
	b.Params["video_comment_restore"] = v
}

// MarketCommentNew New comment to market item notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) MarketCommentNew(v bool) {
	b.Params["market_comment_new"] = v
}

// MarketCommentEdit A market comment has been edited ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) MarketCommentEdit(v bool) {
	b.Params["market_comment_edit"] = v
}

// MarketCommentDelete A market comment has been deleted ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) MarketCommentDelete(v bool) {
	b.Params["market_comment_delete"] = v
}

// MarketCommentRestore A market comment has been restored ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) MarketCommentRestore(v bool) {
	b.Params["market_comment_restore"] = v
}

// PollVoteNew A vote in a public poll has been added ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) PollVoteNew(v bool) {
	b.Params["poll_vote_new"] = v
}

// GroupJoin Joined community notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) GroupJoin(v bool) {
	b.Params["group_join"] = v
}

// GroupLeave Left community notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBuilder) GroupLeave(v bool) {
	b.Params["group_leave"] = v
}

// GroupChangeSettings parameter
func (b *GroupsSetLongPollSettingsBuilder) GroupChangeSettings(v bool) {
	b.Params["group_change_settings"] = v
}

// GroupChangePhoto parameter
func (b *GroupsSetLongPollSettingsBuilder) GroupChangePhoto(v bool) {
	b.Params["group_change_photo"] = v
}

// GroupOfficersEdit parameter
func (b *GroupsSetLongPollSettingsBuilder) GroupOfficersEdit(v bool) {
	b.Params["group_officers_edit"] = v
}

// UserBlock User added to community blacklist
func (b *GroupsSetLongPollSettingsBuilder) UserBlock(v bool) {
	b.Params["user_block"] = v
}

// UserUnblock User removed from community blacklist
func (b *GroupsSetLongPollSettingsBuilder) UserUnblock(v bool) {
	b.Params["user_unblock"] = v
}

// GroupsUnbanBuilder builder
//
// https://vk.com/dev/groups.unban
type GroupsUnbanBuilder struct {
	api.Params
}

// NewGroupsUnbanBuilder func
func NewGroupsUnbanBuilder() *GroupsUnbanBuilder {
	return &GroupsUnbanBuilder{api.Params{}}
}

// GroupID parameter
func (b *GroupsUnbanBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// OwnerID parameter
func (b *GroupsUnbanBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}
