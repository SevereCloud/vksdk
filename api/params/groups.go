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
func (b *GroupsAddAddressBuilder) GroupID(v int) *GroupsAddAddressBuilder {
	b.Params["group_id"] = v
	return b
}

// Title parameter
func (b *GroupsAddAddressBuilder) Title(v string) *GroupsAddAddressBuilder {
	b.Params["title"] = v
	return b
}

// Address parameter
func (b *GroupsAddAddressBuilder) Address(v string) *GroupsAddAddressBuilder {
	b.Params["address"] = v
	return b
}

// AdditionalAddress parameter
func (b *GroupsAddAddressBuilder) AdditionalAddress(v string) *GroupsAddAddressBuilder {
	b.Params["additional_address"] = v
	return b
}

// CountryID parameter
func (b *GroupsAddAddressBuilder) CountryID(v int) *GroupsAddAddressBuilder {
	b.Params["country_id"] = v
	return b
}

// CityID parameter
func (b *GroupsAddAddressBuilder) CityID(v int) *GroupsAddAddressBuilder {
	b.Params["city_id"] = v
	return b
}

// MetroID parameter
func (b *GroupsAddAddressBuilder) MetroID(v int) *GroupsAddAddressBuilder {
	b.Params["metro_id"] = v
	return b
}

// Latitude parameter
func (b *GroupsAddAddressBuilder) Latitude(v float64) *GroupsAddAddressBuilder {
	b.Params["latitude"] = v
	return b
}

// Longitude parameter
func (b *GroupsAddAddressBuilder) Longitude(v float64) *GroupsAddAddressBuilder {
	b.Params["longitude"] = v
	return b
}

// Phone parameter
func (b *GroupsAddAddressBuilder) Phone(v string) *GroupsAddAddressBuilder {
	b.Params["phone"] = v
	return b
}

// WorkInfoStatus parameter
func (b *GroupsAddAddressBuilder) WorkInfoStatus(v string) *GroupsAddAddressBuilder {
	b.Params["work_info_status"] = v
	return b
}

// Timetable parameter
func (b *GroupsAddAddressBuilder) Timetable(v string) *GroupsAddAddressBuilder {
	b.Params["timetable"] = v
	return b
}

// IsMainAddress parameter
func (b *GroupsAddAddressBuilder) IsMainAddress(v bool) *GroupsAddAddressBuilder {
	b.Params["is_main_address"] = v
	return b
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
func (b *GroupsAddCallbackServerBuilder) GroupID(v int) *GroupsAddCallbackServerBuilder {
	b.Params["group_id"] = v
	return b
}

// URL parameter
func (b *GroupsAddCallbackServerBuilder) URL(v string) *GroupsAddCallbackServerBuilder {
	b.Params["url"] = v
	return b
}

// Title parameter
func (b *GroupsAddCallbackServerBuilder) Title(v string) *GroupsAddCallbackServerBuilder {
	b.Params["title"] = v
	return b
}

// SecretKey parameter
func (b *GroupsAddCallbackServerBuilder) SecretKey(v string) *GroupsAddCallbackServerBuilder {
	b.Params["secret_key"] = v
	return b
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
func (b *GroupsAddLinkBuilder) GroupID(v int) *GroupsAddLinkBuilder {
	b.Params["group_id"] = v
	return b
}

// Link Link URL.
func (b *GroupsAddLinkBuilder) Link(v string) *GroupsAddLinkBuilder {
	b.Params["link"] = v
	return b
}

// Text Description text for the link.
func (b *GroupsAddLinkBuilder) Text(v string) *GroupsAddLinkBuilder {
	b.Params["text"] = v
	return b
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
func (b *GroupsApproveRequestBuilder) GroupID(v int) *GroupsApproveRequestBuilder {
	b.Params["group_id"] = v
	return b
}

// UserID User ID.
func (b *GroupsApproveRequestBuilder) UserID(v int) *GroupsApproveRequestBuilder {
	b.Params["user_id"] = v
	return b
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
func (b *GroupsBanBuilder) GroupID(v int) *GroupsBanBuilder {
	b.Params["group_id"] = v
	return b
}

// OwnerID parameter
func (b *GroupsBanBuilder) OwnerID(v int) *GroupsBanBuilder {
	b.Params["owner_id"] = v
	return b
}

// EndDate parameter
func (b *GroupsBanBuilder) EndDate(v int) *GroupsBanBuilder {
	b.Params["end_date"] = v
	return b
}

// Reason parameter
func (b *GroupsBanBuilder) Reason(v int) *GroupsBanBuilder {
	b.Params["reason"] = v
	return b
}

// Comment parameter
func (b *GroupsBanBuilder) Comment(v string) *GroupsBanBuilder {
	b.Params["comment"] = v
	return b
}

// CommentVisible parameter
func (b *GroupsBanBuilder) CommentVisible(v bool) *GroupsBanBuilder {
	b.Params["comment_visible"] = v
	return b
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
func (b *GroupsCreateBuilder) Title(v string) *GroupsCreateBuilder {
	b.Params["title"] = v
	return b
}

// Description Community description (ignored for 'type' = 'public').
func (b *GroupsCreateBuilder) Description(v string) *GroupsCreateBuilder {
	b.Params["description"] = v
	return b
}

// Type Community type. Possible values:
//
// * group – group
//
// * event – event
//
// * public – public page
func (b *GroupsCreateBuilder) Type(v string) *GroupsCreateBuilder {
	b.Params["type"] = v
	return b
}

// PublicCategory Category ID (for 'type' = 'public' only).
func (b *GroupsCreateBuilder) PublicCategory(v int) *GroupsCreateBuilder {
	b.Params["public_category"] = v
	return b
}

// Subtype Public page subtype. Possible values:
//
// * 1 – place or small business
//
// * 2 – company, organization or website
//
// * 3 – famous person or group of people
//
// * 4 – product or work of art.
func (b *GroupsCreateBuilder) Subtype(v int) *GroupsCreateBuilder {
	b.Params["subtype"] = v
	return b
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
func (b *GroupsDeleteCallbackServerBuilder) GroupID(v int) *GroupsDeleteCallbackServerBuilder {
	b.Params["group_id"] = v
	return b
}

// ServerID parameter
func (b *GroupsDeleteCallbackServerBuilder) ServerID(v int) *GroupsDeleteCallbackServerBuilder {
	b.Params["server_id"] = v
	return b
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
func (b *GroupsDeleteLinkBuilder) GroupID(v int) *GroupsDeleteLinkBuilder {
	b.Params["group_id"] = v
	return b
}

// LinkID Link ID.
func (b *GroupsDeleteLinkBuilder) LinkID(v int) *GroupsDeleteLinkBuilder {
	b.Params["link_id"] = v
	return b
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
func (b *GroupsDisableOnlineBuilder) GroupID(v int) *GroupsDisableOnlineBuilder {
	b.Params["group_id"] = v
	return b
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
func (b *GroupsEditBuilder) GroupID(v int) *GroupsEditBuilder {
	b.Params["group_id"] = v
	return b
}

// Title Community title.
func (b *GroupsEditBuilder) Title(v string) *GroupsEditBuilder {
	b.Params["title"] = v
	return b
}

// Description Community description.
func (b *GroupsEditBuilder) Description(v string) *GroupsEditBuilder {
	b.Params["description"] = v
	return b
}

// ScreenName Community screen name.
func (b *GroupsEditBuilder) ScreenName(v string) *GroupsEditBuilder {
	b.Params["screen_name"] = v
	return b
}

// Access Community type. Possible values:
//
// * 0 – open
//
// * 1 – closed
//
// * 2 – private.
func (b *GroupsEditBuilder) Access(v int) *GroupsEditBuilder {
	b.Params["access"] = v
	return b
}

// Website Website that will be displayed in the community information field.
func (b *GroupsEditBuilder) Website(v string) *GroupsEditBuilder {
	b.Params["website"] = v
	return b
}

// Subject Community subject. Possible values:
//
// * 1 – auto/moto
//
// * 2 – activity holidays
//
// * 3 – business
//
// * 4 – pets
//
// * 5 – health
//
// * 6 – dating and communication
//
// * 7 – games
//
// * 8 – IT (computers and software)
//
// * 9 – cinema
//
// * 10 – beauty and fashion
//
// * 11 – cooking
//
// * 12 – art and culture
//
// * 13 – literature
//
// * 14 – mobile services and internet
//
// * 15 – music
//
// * 16 – science and technology
//
// * 17 – real estate
//
// * 18 – news and media
//
// * 19 – security
//
// * 20 – education
//
// * 21 – home and renovations
//
// * 22 – politics
//
// * 23 – food
//
// * 24 – industry
//
// * 25 – travel
//
// * 26 – work
//
// * 27 – entertainment
//
// * 28 – religion
//
// * 29 – family
//
// * 30 – sports
//
// * 31 – insurance
//
// * 32 – television
//
// * 33 – goods and services
//
// * 34 – hobbies
//
// * 35 – finance
//
// * 36 – photo
//
// * 37 – esoterics
//
// * 38 – electronics and appliances
//
// * 39 – erotic
//
// * 40 – humor
//
// * 41 – society, humanities
//
// * 42 – design and graphics.
func (b *GroupsEditBuilder) Subject(v string) *GroupsEditBuilder {
	b.Params["subject"] = v
	return b
}

// Email Organizer email (for events).
func (b *GroupsEditBuilder) Email(v string) *GroupsEditBuilder {
	b.Params["email"] = v
	return b
}

// Phone Organizer phone number (for events).
func (b *GroupsEditBuilder) Phone(v string) *GroupsEditBuilder {
	b.Params["phone"] = v
	return b
}

// Rss RSS feed address for import available only to communities with special permission.
// Contact vk.com/support to get it.
func (b *GroupsEditBuilder) Rss(v string) *GroupsEditBuilder {
	b.Params["rss"] = v
	return b
}

// EventStartDate Event start date in Unixtime format.
func (b *GroupsEditBuilder) EventStartDate(v int) *GroupsEditBuilder {
	b.Params["event_start_date"] = v
	return b
}

// EventFinishDate Event finish date in Unixtime format.
func (b *GroupsEditBuilder) EventFinishDate(v int) *GroupsEditBuilder {
	b.Params["event_finish_date"] = v
	return b
}

// EventGroupID Organizer community ID (for events only).
func (b *GroupsEditBuilder) EventGroupID(v int) *GroupsEditBuilder {
	b.Params["event_group_id"] = v
	return b
}

// PublicCategory Public page category ID.
func (b *GroupsEditBuilder) PublicCategory(v int) *GroupsEditBuilder {
	b.Params["public_category"] = v
	return b
}

// PublicSubcategory Public page subcategory ID.
func (b *GroupsEditBuilder) PublicSubcategory(v int) *GroupsEditBuilder {
	b.Params["public_subcategory"] = v
	return b
}

// PublicDate Founding date of a company or organization owning the community in dd.mm.YYYY format.
func (b *GroupsEditBuilder) PublicDate(v string) *GroupsEditBuilder {
	b.Params["public_date"] = v
	return b
}

// Wall Wall settings. Possible values:
//
// * 0 – disabled
//
// * 1 – open
//
// * 2 – limited (groups and events only)
//
// * 3 – closed (groups and events only).
func (b *GroupsEditBuilder) Wall(v int) *GroupsEditBuilder {
	b.Params["wall"] = v
	return b
}

// Topics Board topics settings. Possbile values:
//
// * 0 – disabled
//
// * 1 – open
//
// * 2 – limited (for groups and events only).
func (b *GroupsEditBuilder) Topics(v int) *GroupsEditBuilder {
	b.Params["topics"] = v
	return b
}

// Photos Photos settings. Possible values:
//
// * 0 – disabled
//
// * 1 – open
//
// * 2 – limited (for groups and events only).
func (b *GroupsEditBuilder) Photos(v int) *GroupsEditBuilder {
	b.Params["photos"] = v
	return b
}

// Video Video settings. Possible values:
//
// * 0 – disabled
//
// * 1 – open
//
// * 2 – limited (for groups and events only).
func (b *GroupsEditBuilder) Video(v int) *GroupsEditBuilder {
	b.Params["video"] = v
	return b
}

// Audio Audio settings. Possible values:
//
// * 0 – disabled
//
// * 1 – open
//
// * 2 – limited (for groups and events only).
func (b *GroupsEditBuilder) Audio(v int) *GroupsEditBuilder {
	b.Params["audio"] = v
	return b
}

// Links Links settings (for public pages only). Possible values:
//
// * 0 – disabled
//
// * 1 – enabled.
func (b *GroupsEditBuilder) Links(v bool) *GroupsEditBuilder {
	b.Params["links"] = v
	return b
}

// Events Events settings (for public pages only). Possible values:
//
// * 0 – disabled
//
// * 1 – enabled.
func (b *GroupsEditBuilder) Events(v bool) *GroupsEditBuilder {
	b.Params["events"] = v
	return b
}

// Places Places settings (for public pages only). Possible values:
//
// * 0 – disabled
//
// * 1 – enabled.
func (b *GroupsEditBuilder) Places(v bool) *GroupsEditBuilder {
	b.Params["places"] = v
	return b
}

// Contacts Contacts settings (for public pages only). Possible values:
//
// * 0 – disabled
//
// * 1 – enabled.
func (b *GroupsEditBuilder) Contacts(v bool) *GroupsEditBuilder {
	b.Params["contacts"] = v
	return b
}

// Docs Documents settings. Possible values:
//
// * 0 – disabled
//
// * 1 – open
//
// * 2 – limited (for groups and events only).
func (b *GroupsEditBuilder) Docs(v int) *GroupsEditBuilder {
	b.Params["docs"] = v
	return b
}

// Wiki Wiki pages settings. Possible values:
//
// * 0 – disabled
//
// * 1 – open
//
// * 2 – limited (for groups and events only).
func (b *GroupsEditBuilder) Wiki(v int) *GroupsEditBuilder {
	b.Params["wiki"] = v
	return b
}

// Messages Community messages. Possible values:
//
// * 0 — disabled
//
// * 1 — enabled.
func (b *GroupsEditBuilder) Messages(v bool) *GroupsEditBuilder {
	b.Params["messages"] = v
	return b
}

// Articles parameter
func (b *GroupsEditBuilder) Articles(v bool) *GroupsEditBuilder {
	b.Params["articles"] = v
	return b
}

// Addresses parameter
func (b *GroupsEditBuilder) Addresses(v bool) *GroupsEditBuilder {
	b.Params["addresses"] = v
	return b
}

// AgeLimits Community age limits. Possible values:
//
// * 1 — no limits
//
// * 2 — 16+
//
// * 3 — 18+.
func (b *GroupsEditBuilder) AgeLimits(v int) *GroupsEditBuilder {
	b.Params["age_limits"] = v
	return b
}

// Market Market settings. Possible values:
//
// * 0 – disabled
//
// * 1 – enabled.
func (b *GroupsEditBuilder) Market(v bool) *GroupsEditBuilder {
	b.Params["market"] = v
	return b
}

// MarketComments market comments settings. Possible values:
//
// * 0 – disabled
//
// * 1 – enabled.
func (b *GroupsEditBuilder) MarketComments(v bool) *GroupsEditBuilder {
	b.Params["market_comments"] = v
	return b
}

// MarketCountry Market delivery countries.
func (b *GroupsEditBuilder) MarketCountry(v []int) *GroupsEditBuilder {
	b.Params["market_country"] = v
	return b
}

// MarketCity Market delivery cities (if only one country is specified).
func (b *GroupsEditBuilder) MarketCity(v []int) *GroupsEditBuilder {
	b.Params["market_city"] = v
	return b
}

// MarketCurrency Market currency settings. Possbile values:
//
// * 643 – Russian rubles
//
// * 980 – Ukrainian hryvnia
//
// * 398 – Kazakh tenge
//
// * 978 – Euro
//
// * 840 – US dollars
func (b *GroupsEditBuilder) MarketCurrency(v int) *GroupsEditBuilder {
	b.Params["market_currency"] = v
	return b
}

// MarketContact Seller contact for market.
// Set '0' for community messages.
func (b *GroupsEditBuilder) MarketContact(v int) *GroupsEditBuilder {
	b.Params["market_contact"] = v
	return b
}

// MarketWiki ID of a wiki page with market description.
func (b *GroupsEditBuilder) MarketWiki(v int) *GroupsEditBuilder {
	b.Params["market_wiki"] = v
	return b
}

// ObsceneFilter Obscene expressions filter in comments. Possible values:
//
// * 0 – disabled
//
// * 1 – enabled.
func (b *GroupsEditBuilder) ObsceneFilter(v bool) *GroupsEditBuilder {
	b.Params["obscene_filter"] = v
	return b
}

// ObsceneStopwords Stopwords filter in comments. Possible values:
//
// * 0 – disabled
//
// * 1 – enabled.
func (b *GroupsEditBuilder) ObsceneStopwords(v bool) *GroupsEditBuilder {
	b.Params["obscene_stopwords"] = v
	return b
}

// ObsceneWords Keywords for stopwords filter.
func (b *GroupsEditBuilder) ObsceneWords(v []string) *GroupsEditBuilder {
	b.Params["obscene_words"] = v
	return b
}

// MainSection parameter
func (b *GroupsEditBuilder) MainSection(v int) *GroupsEditBuilder {
	b.Params["main_section"] = v
	return b
}

// SecondarySection parameter
func (b *GroupsEditBuilder) SecondarySection(v int) *GroupsEditBuilder {
	b.Params["secondary_section"] = v
	return b
}

// Country Country of the community.
func (b *GroupsEditBuilder) Country(v int) *GroupsEditBuilder {
	b.Params["country"] = v
	return b
}

// City City of the community.
func (b *GroupsEditBuilder) City(v int) *GroupsEditBuilder {
	b.Params["city"] = v
	return b
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
func (b *GroupsEditAddressBuilder) GroupID(v int) *GroupsEditAddressBuilder {
	b.Params["group_id"] = v
	return b
}

// AddressID parameter
func (b *GroupsEditAddressBuilder) AddressID(v int) *GroupsEditAddressBuilder {
	b.Params["address_id"] = v
	return b
}

// Title parameter
func (b *GroupsEditAddressBuilder) Title(v string) *GroupsEditAddressBuilder {
	b.Params["title"] = v
	return b
}

// Address parameter
func (b *GroupsEditAddressBuilder) Address(v string) *GroupsEditAddressBuilder {
	b.Params["address"] = v
	return b
}

// AdditionalAddress parameter
func (b *GroupsEditAddressBuilder) AdditionalAddress(v string) *GroupsEditAddressBuilder {
	b.Params["additional_address"] = v
	return b
}

// CountryID parameter
func (b *GroupsEditAddressBuilder) CountryID(v int) *GroupsEditAddressBuilder {
	b.Params["country_id"] = v
	return b
}

// CityID parameter
func (b *GroupsEditAddressBuilder) CityID(v int) *GroupsEditAddressBuilder {
	b.Params["city_id"] = v
	return b
}

// MetroID parameter
func (b *GroupsEditAddressBuilder) MetroID(v int) *GroupsEditAddressBuilder {
	b.Params["metro_id"] = v
	return b
}

// Latitude parameter
func (b *GroupsEditAddressBuilder) Latitude(v float64) *GroupsEditAddressBuilder {
	b.Params["latitude"] = v
	return b
}

// Longitude parameter
func (b *GroupsEditAddressBuilder) Longitude(v float64) *GroupsEditAddressBuilder {
	b.Params["longitude"] = v
	return b
}

// Phone parameter
func (b *GroupsEditAddressBuilder) Phone(v string) *GroupsEditAddressBuilder {
	b.Params["phone"] = v
	return b
}

// WorkInfoStatus parameter
func (b *GroupsEditAddressBuilder) WorkInfoStatus(v string) *GroupsEditAddressBuilder {
	b.Params["work_info_status"] = v
	return b
}

// Timetable parameter
func (b *GroupsEditAddressBuilder) Timetable(v string) *GroupsEditAddressBuilder {
	b.Params["timetable"] = v
	return b
}

// IsMainAddress parameter
func (b *GroupsEditAddressBuilder) IsMainAddress(v bool) *GroupsEditAddressBuilder {
	b.Params["is_main_address"] = v
	return b
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
func (b *GroupsEditCallbackServerBuilder) GroupID(v int) *GroupsEditCallbackServerBuilder {
	b.Params["group_id"] = v
	return b
}

// ServerID parameter
func (b *GroupsEditCallbackServerBuilder) ServerID(v int) *GroupsEditCallbackServerBuilder {
	b.Params["server_id"] = v
	return b
}

// URL parameter
func (b *GroupsEditCallbackServerBuilder) URL(v string) *GroupsEditCallbackServerBuilder {
	b.Params["url"] = v
	return b
}

// Title parameter
func (b *GroupsEditCallbackServerBuilder) Title(v string) *GroupsEditCallbackServerBuilder {
	b.Params["title"] = v
	return b
}

// SecretKey parameter
func (b *GroupsEditCallbackServerBuilder) SecretKey(v string) *GroupsEditCallbackServerBuilder {
	b.Params["secret_key"] = v
	return b
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
func (b *GroupsEditLinkBuilder) GroupID(v int) *GroupsEditLinkBuilder {
	b.Params["group_id"] = v
	return b
}

// LinkID Link ID.
func (b *GroupsEditLinkBuilder) LinkID(v int) *GroupsEditLinkBuilder {
	b.Params["link_id"] = v
	return b
}

// Text New description text for the link.
func (b *GroupsEditLinkBuilder) Text(v string) *GroupsEditLinkBuilder {
	b.Params["text"] = v
	return b
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
func (b *GroupsEditManagerBuilder) GroupID(v int) *GroupsEditManagerBuilder {
	b.Params["group_id"] = v
	return b
}

// UserID User ID.
func (b *GroupsEditManagerBuilder) UserID(v int) *GroupsEditManagerBuilder {
	b.Params["user_id"] = v
	return b
}

// Role Manager role. Possible values:
//
// * moderator
//
// * editor
//
// * administrator.
func (b *GroupsEditManagerBuilder) Role(v string) *GroupsEditManagerBuilder {
	b.Params["role"] = v
	return b
}

// IsContact parameter
//
// * '1' — to show the manager in Contacts block of the community.
func (b *GroupsEditManagerBuilder) IsContact(v bool) *GroupsEditManagerBuilder {
	b.Params["is_contact"] = v
	return b
}

// ContactPosition Position to show in Contacts block.
func (b *GroupsEditManagerBuilder) ContactPosition(v string) *GroupsEditManagerBuilder {
	b.Params["contact_position"] = v
	return b
}

// ContactPhone Contact phone.
func (b *GroupsEditManagerBuilder) ContactPhone(v string) *GroupsEditManagerBuilder {
	b.Params["contact_phone"] = v
	return b
}

// ContactEmail Contact e-mail.
func (b *GroupsEditManagerBuilder) ContactEmail(v string) *GroupsEditManagerBuilder {
	b.Params["contact_email"] = v
	return b
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
func (b *GroupsEnableOnlineBuilder) GroupID(v int) *GroupsEnableOnlineBuilder {
	b.Params["group_id"] = v
	return b
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
func (b *GroupsGetBuilder) UserID(v int) *GroupsGetBuilder {
	b.Params["user_id"] = v
	return b
}

// Extended parameter
//
// * '1' — to return complete information about a user's communities
//
// * '0' — to return a list of community IDs without any additional fields (default),
func (b *GroupsGetBuilder) Extended(v bool) *GroupsGetBuilder {
	b.Params["extended"] = v
	return b
}

// Filter Types of communities to return:
//
// * 'admin' — to return communities administered by the user
//
// * 'editor' — to return communities where the user is an administrator or editor
//
// * 'moder' — to return communities where the user is an administrator, editor, or moderator
//
// * 'groups' — to return only groups
//
// * 'publics' — to return only public pages
//
// * 'events' — to return only events
func (b *GroupsGetBuilder) Filter(v []string) *GroupsGetBuilder {
	b.Params["filter"] = v
	return b
}

// Fields Profile fields to return.
func (b *GroupsGetBuilder) Fields(v []string) *GroupsGetBuilder {
	b.Params["fields"] = v
	return b
}

// Offset Offset needed to return a specific subset of communities.
func (b *GroupsGetBuilder) Offset(v int) *GroupsGetBuilder {
	b.Params["offset"] = v
	return b
}

// Count Number of communities to return.
func (b *GroupsGetBuilder) Count(v int) *GroupsGetBuilder {
	b.Params["count"] = v
	return b
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
func (b *GroupsGetAddressesBuilder) GroupID(v int) *GroupsGetAddressesBuilder {
	b.Params["group_id"] = v
	return b
}

// AddressIDs parameter
func (b *GroupsGetAddressesBuilder) AddressIDs(v []int) *GroupsGetAddressesBuilder {
	b.Params["address_ids"] = v
	return b
}

// Latitude Latitude of  the user geo position.
func (b *GroupsGetAddressesBuilder) Latitude(v float64) *GroupsGetAddressesBuilder {
	b.Params["latitude"] = v
	return b
}

// Longitude Longitude of the user geo position.
func (b *GroupsGetAddressesBuilder) Longitude(v float64) *GroupsGetAddressesBuilder {
	b.Params["longitude"] = v
	return b
}

// Offset Offset needed to return a specific subset of community addresses.
func (b *GroupsGetAddressesBuilder) Offset(v int) *GroupsGetAddressesBuilder {
	b.Params["offset"] = v
	return b
}

// Count Number of community addresses to return.
func (b *GroupsGetAddressesBuilder) Count(v int) *GroupsGetAddressesBuilder {
	b.Params["count"] = v
	return b
}

// Fields Address fields
func (b *GroupsGetAddressesBuilder) Fields(v []string) *GroupsGetAddressesBuilder {
	b.Params["fields"] = v
	return b
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
func (b *GroupsGetBannedBuilder) GroupID(v int) *GroupsGetBannedBuilder {
	b.Params["group_id"] = v
	return b
}

// Offset Offset needed to return a specific subset of users.
func (b *GroupsGetBannedBuilder) Offset(v int) *GroupsGetBannedBuilder {
	b.Params["offset"] = v
	return b
}

// Count Number of users to return.
func (b *GroupsGetBannedBuilder) Count(v int) *GroupsGetBannedBuilder {
	b.Params["count"] = v
	return b
}

// Fields parameter
func (b *GroupsGetBannedBuilder) Fields(v []string) *GroupsGetBannedBuilder {
	b.Params["fields"] = v
	return b
}

// OwnerID parameter
func (b *GroupsGetBannedBuilder) OwnerID(v int) *GroupsGetBannedBuilder {
	b.Params["owner_id"] = v
	return b
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
func (b *GroupsGetByIDBuilder) GroupIDs(v []string) *GroupsGetByIDBuilder {
	b.Params["group_ids"] = v
	return b
}

// GroupID ID or screen name of the community.
func (b *GroupsGetByIDBuilder) GroupID(v string) *GroupsGetByIDBuilder {
	b.Params["group_id"] = v
	return b
}

// Fields Group fields to return.
func (b *GroupsGetByIDBuilder) Fields(v []string) *GroupsGetByIDBuilder {
	b.Params["fields"] = v
	return b
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
func (b *GroupsGetCallbackConfirmationCodeBuilder) GroupID(v int) *GroupsGetCallbackConfirmationCodeBuilder {
	b.Params["group_id"] = v
	return b
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
func (b *GroupsGetCallbackServersBuilder) GroupID(v int) *GroupsGetCallbackServersBuilder {
	b.Params["group_id"] = v
	return b
}

// ServerIDs parameter
func (b *GroupsGetCallbackServersBuilder) ServerIDs(v []int) *GroupsGetCallbackServersBuilder {
	b.Params["server_ids"] = v
	return b
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
func (b *GroupsGetCallbackSettingsBuilder) GroupID(v int) *GroupsGetCallbackSettingsBuilder {
	b.Params["group_id"] = v
	return b
}

// ServerID Server ID.
func (b *GroupsGetCallbackSettingsBuilder) ServerID(v int) *GroupsGetCallbackSettingsBuilder {
	b.Params["server_id"] = v
	return b
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
func (b *GroupsGetCatalogBuilder) CategoryID(v int) *GroupsGetCatalogBuilder {
	b.Params["category_id"] = v
	return b
}

// SubcategoryID Subcategory id received from [vk.com/dev/groups.getCatalogInfo|groups.getCatalogInfo].
func (b *GroupsGetCatalogBuilder) SubcategoryID(v int) *GroupsGetCatalogBuilder {
	b.Params["subcategory_id"] = v
	return b
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

// Extended 1 – to return communities count and three communities for preview.
//
// By default: 0.
func (b *GroupsGetCatalogInfoBuilder) Extended(v bool) *GroupsGetCatalogInfoBuilder {
	b.Params["extended"] = v
	return b
}

// Subcategories 1 – to return subcategories info.
//
// By default: 0.
func (b *GroupsGetCatalogInfoBuilder) Subcategories(v bool) *GroupsGetCatalogInfoBuilder {
	b.Params["subcategories"] = v
	return b
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
func (b *GroupsGetInvitedUsersBuilder) GroupID(v int) *GroupsGetInvitedUsersBuilder {
	b.Params["group_id"] = v
	return b
}

// Offset Offset needed to return a specific subset of results.
func (b *GroupsGetInvitedUsersBuilder) Offset(v int) *GroupsGetInvitedUsersBuilder {
	b.Params["offset"] = v
	return b
}

// Count Number of results to return.
func (b *GroupsGetInvitedUsersBuilder) Count(v int) *GroupsGetInvitedUsersBuilder {
	b.Params["count"] = v
	return b
}

// Fields List of additional fields to be returned. Available values
//
// * sex
//
// * bdate
//
// * city
//
// * country
//
// * photo_50
//
// * photo_100
//
// * photo_200_orig
//
// * photo_200
//
// * photo_400_orig
//
// * photo_max
//
// * photo_max_orig
//
// * online
//
// * online_mobile
//
// * lists
//
// * domain
//
// * has_mobile
//
// * contacts
//
// * connections
//
// * site
//
// * education
//
// * universities
//
// * schools
//
// * can_post
//
// * can_see_all_posts
//
// * can_see_audio
//
// * can_write_private_message
//
// * status
//
// * last_seen
//
// * common_count
//
// * relation
//
// * relatives
//
// * counters.
func (b *GroupsGetInvitedUsersBuilder) Fields(v []string) *GroupsGetInvitedUsersBuilder {
	b.Params["fields"] = v
	return b
}

// NameCase Case for declension of user name and surname. Possible values:
//
// * nom — nominative (default)
//
// * gen — genitive
//
// * dat — dative
//
// * acc — accusative
//
// * ins — instrumental
//
// * abl — prepositional.
func (b *GroupsGetInvitedUsersBuilder) NameCase(v string) *GroupsGetInvitedUsersBuilder {
	b.Params["name_case"] = v
	return b
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
func (b *GroupsGetInvitesBuilder) Offset(v int) *GroupsGetInvitesBuilder {
	b.Params["offset"] = v
	return b
}

// Count Number of invitations to return.
func (b *GroupsGetInvitesBuilder) Count(v int) *GroupsGetInvitesBuilder {
	b.Params["count"] = v
	return b
}

// Extended parameter
//
// * '1' — to return additional [vk.com/dev/fields_groups|fields] for communities..
func (b *GroupsGetInvitesBuilder) Extended(v bool) *GroupsGetInvitesBuilder {
	b.Params["extended"] = v
	return b
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
func (b *GroupsGetLongPollServerBuilder) GroupID(v int) *GroupsGetLongPollServerBuilder {
	b.Params["group_id"] = v
	return b
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
func (b *GroupsGetLongPollSettingsBuilder) GroupID(v int) *GroupsGetLongPollSettingsBuilder {
	b.Params["group_id"] = v
	return b
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
func (b *GroupsGetMembersBuilder) GroupID(v string) *GroupsGetMembersBuilder {
	b.Params["group_id"] = v
	return b
}

// Sort Sort order.
// Available values
//
// *  'id_asc'
//
// *  'id_desc'
//
// *  'time_asc'
//
// *  'time_desc'.
// 'time_asc' and 'time_desc' are availavle only if the method is called by the group's 'moderator'.
func (b *GroupsGetMembersBuilder) Sort(v string) *GroupsGetMembersBuilder {
	b.Params["sort"] = v
	return b
}

// Offset Offset needed to return a specific subset of community members.
func (b *GroupsGetMembersBuilder) Offset(v int) *GroupsGetMembersBuilder {
	b.Params["offset"] = v
	return b
}

// Count Number of community members to return.
func (b *GroupsGetMembersBuilder) Count(v int) *GroupsGetMembersBuilder {
	b.Params["count"] = v
	return b
}

// Fields List of additional fields to be returned. Available values
//
// * sex
//
// * bdate
//
// * city
//
// * country
//
// * photo_50
//
// * photo_100
//
// * photo_200_orig
//
// * photo_200
//
// * photo_400_orig
//
// * photo_max
//
// * photo_max_orig
//
// * online
//
// * online_mobile
//
// * lists
//
// * domain
//
// * has_mobile
//
// * contacts
//
// * connections
//
// * site
//
// * education
//
// * universities
//
// * schools
//
// * can_post
//
// * can_see_all_posts
//
// * can_see_audio
//
// * can_write_private_message
//
// * status
//
// * last_seen
//
// * common_count
//
// * relation
//
// * relatives
//
// * counters.
func (b *GroupsGetMembersBuilder) Fields(v []string) *GroupsGetMembersBuilder {
	b.Params["fields"] = v
	return b
}

// Filter parameter
//
// * friends – only friends in this community will be returned
//
// * unsure – only those who pressed I may attend will be returned (if its an event).
func (b *GroupsGetMembersBuilder) Filter(v string) *GroupsGetMembersBuilder {
	b.Params["filter"] = v
	return b
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
func (b *GroupsGetRequestsBuilder) GroupID(v int) *GroupsGetRequestsBuilder {
	b.Params["group_id"] = v
	return b
}

// Offset Offset needed to return a specific subset of results.
func (b *GroupsGetRequestsBuilder) Offset(v int) *GroupsGetRequestsBuilder {
	b.Params["offset"] = v
	return b
}

// Count Number of results to return.
func (b *GroupsGetRequestsBuilder) Count(v int) *GroupsGetRequestsBuilder {
	b.Params["count"] = v
	return b
}

// Fields Profile fields to return.
func (b *GroupsGetRequestsBuilder) Fields(v []string) *GroupsGetRequestsBuilder {
	b.Params["fields"] = v
	return b
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
func (b *GroupsGetSettingsBuilder) GroupID(v int) *GroupsGetSettingsBuilder {
	b.Params["group_id"] = v
	return b
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
func (b *GroupsInviteBuilder) GroupID(v int) *GroupsInviteBuilder {
	b.Params["group_id"] = v
	return b
}

// UserID User ID.
func (b *GroupsInviteBuilder) UserID(v int) *GroupsInviteBuilder {
	b.Params["user_id"] = v
	return b
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
func (b *GroupsIsMemberBuilder) GroupID(v string) *GroupsIsMemberBuilder {
	b.Params["group_id"] = v
	return b
}

// UserID User ID.
func (b *GroupsIsMemberBuilder) UserID(v int) *GroupsIsMemberBuilder {
	b.Params["user_id"] = v
	return b
}

// UserIDs User IDs.
func (b *GroupsIsMemberBuilder) UserIDs(v []int) *GroupsIsMemberBuilder {
	b.Params["user_ids"] = v
	return b
}

// Extended parameter
//
// * '1' — to return an extended response with additional fields.
//
// By default: '0'.
func (b *GroupsIsMemberBuilder) Extended(v bool) *GroupsIsMemberBuilder {
	b.Params["extended"] = v
	return b
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
func (b *GroupsJoinBuilder) GroupID(v int) *GroupsJoinBuilder {
	b.Params["group_id"] = v
	return b
}

// NotSure Optional parameter which is taken into account when 'gid' belongs to the event:
//
// * '1' — Perhaps I will attend
//
// * '0' — I will be there for sure (default)
func (b *GroupsJoinBuilder) NotSure(v string) *GroupsJoinBuilder {
	b.Params["not_sure"] = v
	return b
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
func (b *GroupsLeaveBuilder) GroupID(v int) *GroupsLeaveBuilder {
	b.Params["group_id"] = v
	return b
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
func (b *GroupsRemoveUserBuilder) GroupID(v int) *GroupsRemoveUserBuilder {
	b.Params["group_id"] = v
	return b
}

// UserID User ID.
func (b *GroupsRemoveUserBuilder) UserID(v int) *GroupsRemoveUserBuilder {
	b.Params["user_id"] = v
	return b
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
func (b *GroupsReorderLinkBuilder) GroupID(v int) *GroupsReorderLinkBuilder {
	b.Params["group_id"] = v
	return b
}

// LinkID Link ID.
func (b *GroupsReorderLinkBuilder) LinkID(v int) *GroupsReorderLinkBuilder {
	b.Params["link_id"] = v
	return b
}

// After ID of the link after which to place the link with 'link_id'.
func (b *GroupsReorderLinkBuilder) After(v int) *GroupsReorderLinkBuilder {
	b.Params["after"] = v
	return b
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
func (b *GroupsSearchBuilder) Q(v string) *GroupsSearchBuilder {
	b.Params["q"] = v
	return b
}

// Type Community type.
// Possible values: 'group, page, event.'
func (b *GroupsSearchBuilder) Type(v string) *GroupsSearchBuilder {
	b.Params["type"] = v
	return b
}

// CountryID Country ID.
func (b *GroupsSearchBuilder) CountryID(v int) *GroupsSearchBuilder {
	b.Params["country_id"] = v
	return b
}

// CityID City ID.
// If this parameter is transmitted, country_id is ignored.
func (b *GroupsSearchBuilder) CityID(v int) *GroupsSearchBuilder {
	b.Params["city_id"] = v
	return b
}

// Future parameter
//
// * '1' — to return only upcoming events.
// Works with the 'type' = 'event' only.
func (b *GroupsSearchBuilder) Future(v bool) *GroupsSearchBuilder {
	b.Params["future"] = v
	return b
}

// Market parameter
//
// * '1' — to return communities with enabled market only.
func (b *GroupsSearchBuilder) Market(v bool) *GroupsSearchBuilder {
	b.Params["market"] = v
	return b
}

// Sort Sort order. Possible values:
//
// * 0 — default sorting (similar the full version of the site)
//
// * 1 — by growth speed
//
// * 2— by the day attendance/members number ratio
//
// * 3 — by the Likes number/members number ratio
//
// * 4 — by the comments number/members number ratio
//
// * 5 — by the boards entries number/members number ratio.
func (b *GroupsSearchBuilder) Sort(v int) *GroupsSearchBuilder {
	b.Params["sort"] = v
	return b
}

// Offset Offset needed to return a specific subset of results.
func (b *GroupsSearchBuilder) Offset(v int) *GroupsSearchBuilder {
	b.Params["offset"] = v
	return b
}

// Count Number of communities to return.
//
// Note that you can not receive more than first thousand of results, regardless of 'count' and 'offset' values.
func (b *GroupsSearchBuilder) Count(v int) *GroupsSearchBuilder {
	b.Params["count"] = v
	return b
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
func (b *GroupsSetCallbackSettingsBuilder) GroupID(v int) *GroupsSetCallbackSettingsBuilder {
	b.Params["group_id"] = v
	return b
}

// ServerID Server ID.
func (b *GroupsSetCallbackSettingsBuilder) ServerID(v int) *GroupsSetCallbackSettingsBuilder {
	b.Params["server_id"] = v
	return b
}

// APIVersion parameter
func (b *GroupsSetCallbackSettingsBuilder) APIVersion(v string) *GroupsSetCallbackSettingsBuilder {
	b.Params["api_version"] = v
	return b
}

// MessageNew A new incoming message has been received
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) MessageNew(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["message_new"] = v
	return b
}

// MessageReply A new outcoming message has been received
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) MessageReply(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["message_reply"] = v
	return b
}

// MessageAllow Allowed messages notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) MessageAllow(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["message_allow"] = v
	return b
}

// MessageEdit parameter
func (b *GroupsSetCallbackSettingsBuilder) MessageEdit(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["message_edit"] = v
	return b
}

// MessageDeny Denied messages notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) MessageDeny(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["message_deny"] = v
	return b
}

// MessageTypingState parameter
func (b *GroupsSetCallbackSettingsBuilder) MessageTypingState(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["message_typing_state"] = v
	return b
}

// PhotoNew New photos notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) PhotoNew(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["photo_new"] = v
	return b
}

// AudioNew New audios notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) AudioNew(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["audio_new"] = v
	return b
}

// VideoNew New videos notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) VideoNew(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["video_new"] = v
	return b
}

// WallReplyNew New wall replies notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) WallReplyNew(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["wall_reply_new"] = v
	return b
}

// WallReplyEdit Wall replies edited notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) WallReplyEdit(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["wall_reply_edit"] = v
	return b
}

// WallReplyDelete A wall comment has been deleted
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) WallReplyDelete(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["wall_reply_delete"] = v
	return b
}

// WallReplyRestore A wall comment has been restored
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) WallReplyRestore(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["wall_reply_restore"] = v
	return b
}

// WallPostNew New wall posts notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) WallPostNew(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["wall_post_new"] = v
	return b
}

// WallRepost New wall posts notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) WallRepost(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["wall_repost"] = v
	return b
}

// BoardPostNew New board posts notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) BoardPostNew(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["board_post_new"] = v
	return b
}

// BoardPostEdit Board posts edited notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) BoardPostEdit(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["board_post_edit"] = v
	return b
}

// BoardPostRestore Board posts restored notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) BoardPostRestore(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["board_post_restore"] = v
	return b
}

// BoardPostDelete Board posts deleted notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) BoardPostDelete(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["board_post_delete"] = v
	return b
}

// PhotoCommentNew New comment to photo notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) PhotoCommentNew(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["photo_comment_new"] = v
	return b
}

// PhotoCommentEdit A photo comment has been edited
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) PhotoCommentEdit(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["photo_comment_edit"] = v
	return b
}

// PhotoCommentDelete A photo comment has been deleted
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) PhotoCommentDelete(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["photo_comment_delete"] = v
	return b
}

// PhotoCommentRestore A photo comment has been restored
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) PhotoCommentRestore(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["photo_comment_restore"] = v
	return b
}

// VideoCommentNew New comment to video notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) VideoCommentNew(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["video_comment_new"] = v
	return b
}

// VideoCommentEdit A video comment has been edited
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) VideoCommentEdit(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["video_comment_edit"] = v
	return b
}

// VideoCommentDelete A video comment has been deleted
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) VideoCommentDelete(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["video_comment_delete"] = v
	return b
}

// VideoCommentRestore A video comment has been restored
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) VideoCommentRestore(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["video_comment_restore"] = v
	return b
}

// MarketCommentNew New comment to market item notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) MarketCommentNew(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["market_comment_new"] = v
	return b
}

// MarketCommentEdit A market comment has been edited
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) MarketCommentEdit(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["market_comment_edit"] = v
	return b
}

// MarketCommentDelete A market comment has been deleted
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) MarketCommentDelete(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["market_comment_delete"] = v
	return b
}

// MarketCommentRestore A market comment has been restored
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) MarketCommentRestore(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["market_comment_restore"] = v
	return b
}

// PollVoteNew A vote in a public poll has been added
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) PollVoteNew(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["poll_vote_new"] = v
	return b
}

// GroupJoin Joined community notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) GroupJoin(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["group_join"] = v
	return b
}

// GroupLeave Left community notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetCallbackSettingsBuilder) GroupLeave(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["group_leave"] = v
	return b
}

// GroupChangeSettings parameter
func (b *GroupsSetCallbackSettingsBuilder) GroupChangeSettings(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["group_change_settings"] = v
	return b
}

// GroupChangePhoto parameter
func (b *GroupsSetCallbackSettingsBuilder) GroupChangePhoto(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["group_change_photo"] = v
	return b
}

// GroupOfficersEdit parameter
func (b *GroupsSetCallbackSettingsBuilder) GroupOfficersEdit(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["group_officers_edit"] = v
	return b
}

// UserBlock User added to community blacklist
func (b *GroupsSetCallbackSettingsBuilder) UserBlock(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["user_block"] = v
	return b
}

// UserUnblock User removed from community blacklist
func (b *GroupsSetCallbackSettingsBuilder) UserUnblock(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["user_unblock"] = v
	return b
}

// LeadFormsNew New form in lead forms
func (b *GroupsSetCallbackSettingsBuilder) LeadFormsNew(v bool) *GroupsSetCallbackSettingsBuilder {
	b.Params["lead_forms_new"] = v
	return b
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
func (b *GroupsSetLongPollSettingsBuilder) GroupID(v int) *GroupsSetLongPollSettingsBuilder {
	b.Params["group_id"] = v
	return b
}

// Enabled Sets whether Long Poll is enabled
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) Enabled(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["enabled"] = v
	return b
}

// APIVersion parameter
func (b *GroupsSetLongPollSettingsBuilder) APIVersion(v string) *GroupsSetLongPollSettingsBuilder {
	b.Params["api_version"] = v
	return b
}

// MessageNew A new incoming message has been received
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) MessageNew(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["message_new"] = v
	return b
}

// MessageReply A new outcoming message has been received
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) MessageReply(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["message_reply"] = v
	return b
}

// MessageAllow Allowed messages notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) MessageAllow(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["message_allow"] = v
	return b
}

// MessageDeny Denied messages notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) MessageDeny(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["message_deny"] = v
	return b
}

// MessageEdit A message has been edited
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) MessageEdit(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["message_edit"] = v
	return b
}

// MessageTypingState parameter
func (b *GroupsSetLongPollSettingsBuilder) MessageTypingState(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["message_typing_state"] = v
	return b
}

// PhotoNew New photos notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) PhotoNew(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["photo_new"] = v
	return b
}

// AudioNew New audios notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) AudioNew(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["audio_new"] = v
	return b
}

// VideoNew New videos notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) VideoNew(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["video_new"] = v
	return b
}

// WallReplyNew New wall replies notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) WallReplyNew(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["wall_reply_new"] = v
	return b
}

// WallReplyEdit Wall replies edited notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) WallReplyEdit(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["wall_reply_edit"] = v
	return b
}

// WallReplyDelete A wall comment has been deleted
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) WallReplyDelete(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["wall_reply_delete"] = v
	return b
}

// WallReplyRestore A wall comment has been restored
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) WallReplyRestore(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["wall_reply_restore"] = v
	return b
}

// WallPostNew New wall posts notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) WallPostNew(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["wall_post_new"] = v
	return b
}

// WallRepost New wall posts notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) WallRepost(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["wall_repost"] = v
	return b
}

// BoardPostNew New board posts notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) BoardPostNew(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["board_post_new"] = v
	return b
}

// BoardPostEdit Board posts edited notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) BoardPostEdit(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["board_post_edit"] = v
	return b
}

// BoardPostRestore Board posts restored notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) BoardPostRestore(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["board_post_restore"] = v
	return b
}

// BoardPostDelete Board posts deleted notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) BoardPostDelete(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["board_post_delete"] = v
	return b
}

// PhotoCommentNew New comment to photo notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) PhotoCommentNew(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["photo_comment_new"] = v
	return b
}

// PhotoCommentEdit A photo comment has been edited
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) PhotoCommentEdit(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["photo_comment_edit"] = v
	return b
}

// PhotoCommentDelete A photo comment has been deleted
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) PhotoCommentDelete(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["photo_comment_delete"] = v
	return b
}

// PhotoCommentRestore A photo comment has been restored
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) PhotoCommentRestore(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["photo_comment_restore"] = v
	return b
}

// VideoCommentNew New comment to video notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) VideoCommentNew(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["video_comment_new"] = v
	return b
}

// VideoCommentEdit A video comment has been edited
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) VideoCommentEdit(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["video_comment_edit"] = v
	return b
}

// VideoCommentDelete A video comment has been deleted
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) VideoCommentDelete(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["video_comment_delete"] = v
	return b
}

// VideoCommentRestore A video comment has been restored
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) VideoCommentRestore(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["video_comment_restore"] = v
	return b
}

// MarketCommentNew New comment to market item notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) MarketCommentNew(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["market_comment_new"] = v
	return b
}

// MarketCommentEdit A market comment has been edited
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) MarketCommentEdit(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["market_comment_edit"] = v
	return b
}

// MarketCommentDelete A market comment has been deleted
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) MarketCommentDelete(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["market_comment_delete"] = v
	return b
}

// MarketCommentRestore A market comment has been restored
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) MarketCommentRestore(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["market_comment_restore"] = v
	return b
}

// PollVoteNew A vote in a public poll has been added
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) PollVoteNew(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["poll_vote_new"] = v
	return b
}

// GroupJoin Joined community notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) GroupJoin(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["group_join"] = v
	return b
}

// GroupLeave Left community notifications
//
// * '0' — disabled
//
// * '1' — enabled.
func (b *GroupsSetLongPollSettingsBuilder) GroupLeave(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["group_leave"] = v
	return b
}

// GroupChangeSettings parameter
func (b *GroupsSetLongPollSettingsBuilder) GroupChangeSettings(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["group_change_settings"] = v
	return b
}

// GroupChangePhoto parameter
func (b *GroupsSetLongPollSettingsBuilder) GroupChangePhoto(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["group_change_photo"] = v
	return b
}

// GroupOfficersEdit parameter
func (b *GroupsSetLongPollSettingsBuilder) GroupOfficersEdit(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["group_officers_edit"] = v
	return b
}

// UserBlock User added to community blacklist
func (b *GroupsSetLongPollSettingsBuilder) UserBlock(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["user_block"] = v
	return b
}

// UserUnblock User removed from community blacklist
func (b *GroupsSetLongPollSettingsBuilder) UserUnblock(v bool) *GroupsSetLongPollSettingsBuilder {
	b.Params["user_unblock"] = v
	return b
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
func (b *GroupsUnbanBuilder) GroupID(v int) *GroupsUnbanBuilder {
	b.Params["group_id"] = v
	return b
}

// OwnerID parameter
func (b *GroupsUnbanBuilder) OwnerID(v int) *GroupsUnbanBuilder {
	b.Params["owner_id"] = v
	return b
}
