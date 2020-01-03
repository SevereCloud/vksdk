package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// GroupsAddAddressBulder builder
//
// https://vk.com/dev/groups.addAddress
type GroupsAddAddressBulder struct {
	api.Params
}

// NewGroupsAddAddressBulder func
func NewGroupsAddAddressBulder() *GroupsAddAddressBulder {
	return &GroupsAddAddressBulder{api.Params{}}
}

// GroupID parameter
func (b *GroupsAddAddressBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Title parameter
func (b *GroupsAddAddressBulder) Title(v string) {
	b.Params["title"] = v
}

// Address parameter
func (b *GroupsAddAddressBulder) Address(v string) {
	b.Params["address"] = v
}

// AdditionalAddress parameter
func (b *GroupsAddAddressBulder) AdditionalAddress(v string) {
	b.Params["additional_address"] = v
}

// CountryID parameter
func (b *GroupsAddAddressBulder) CountryID(v int) {
	b.Params["country_id"] = v
}

// CityID parameter
func (b *GroupsAddAddressBulder) CityID(v int) {
	b.Params["city_id"] = v
}

// MetroID parameter
func (b *GroupsAddAddressBulder) MetroID(v int) {
	b.Params["metro_id"] = v
}

// Latitude parameter
func (b *GroupsAddAddressBulder) Latitude(v float64) {
	b.Params["latitude"] = v
}

// Longitude parameter
func (b *GroupsAddAddressBulder) Longitude(v float64) {
	b.Params["longitude"] = v
}

// Phone parameter
func (b *GroupsAddAddressBulder) Phone(v string) {
	b.Params["phone"] = v
}

// WorkInfoStatus parameter
func (b *GroupsAddAddressBulder) WorkInfoStatus(v string) {
	b.Params["work_info_status"] = v
}

// Timetable parameter
func (b *GroupsAddAddressBulder) Timetable(v string) {
	b.Params["timetable"] = v
}

// IsMainAddress parameter
func (b *GroupsAddAddressBulder) IsMainAddress(v bool) {
	b.Params["is_main_address"] = v
}

// GroupsAddCallbackServerBulder builder
//
// https://vk.com/dev/groups.addCallbackServer
type GroupsAddCallbackServerBulder struct {
	api.Params
}

// NewGroupsAddCallbackServerBulder func
func NewGroupsAddCallbackServerBulder() *GroupsAddCallbackServerBulder {
	return &GroupsAddCallbackServerBulder{api.Params{}}
}

// GroupID parameter
func (b *GroupsAddCallbackServerBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// URL parameter
func (b *GroupsAddCallbackServerBulder) URL(v string) {
	b.Params["url"] = v
}

// Title parameter
func (b *GroupsAddCallbackServerBulder) Title(v string) {
	b.Params["title"] = v
}

// SecretKey parameter
func (b *GroupsAddCallbackServerBulder) SecretKey(v string) {
	b.Params["secret_key"] = v
}

// GroupsAddLinkBulder builder
//
// Allows to add a link to the community.
//
// https://vk.com/dev/groups.addLink
type GroupsAddLinkBulder struct {
	api.Params
}

// NewGroupsAddLinkBulder func
func NewGroupsAddLinkBulder() *GroupsAddLinkBulder {
	return &GroupsAddLinkBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsAddLinkBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Link Link URL.
func (b *GroupsAddLinkBulder) Link(v string) {
	b.Params["link"] = v
}

// Text Description text for the link.
func (b *GroupsAddLinkBulder) Text(v string) {
	b.Params["text"] = v
}

// GroupsApproveRequestBulder builder
//
// Allows to approve join request to the community.
//
// https://vk.com/dev/groups.approveRequest
type GroupsApproveRequestBulder struct {
	api.Params
}

// NewGroupsApproveRequestBulder func
func NewGroupsApproveRequestBulder() *GroupsApproveRequestBulder {
	return &GroupsApproveRequestBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsApproveRequestBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID User ID.
func (b *GroupsApproveRequestBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupsBanBulder builder
//
// https://vk.com/dev/groups.ban
type GroupsBanBulder struct {
	api.Params
}

// NewGroupsBanBulder func
func NewGroupsBanBulder() *GroupsBanBulder {
	return &GroupsBanBulder{api.Params{}}
}

// GroupID parameter
func (b *GroupsBanBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// OwnerID parameter
func (b *GroupsBanBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// EndDate parameter
func (b *GroupsBanBulder) EndDate(v int) {
	b.Params["end_date"] = v
}

// Reason parameter
func (b *GroupsBanBulder) Reason(v int) {
	b.Params["reason"] = v
}

// Comment parameter
func (b *GroupsBanBulder) Comment(v string) {
	b.Params["comment"] = v
}

// CommentVisible parameter
func (b *GroupsBanBulder) CommentVisible(v bool) {
	b.Params["comment_visible"] = v
}

// GroupsCreateBulder builder
//
// Creates a new community.
//
// https://vk.com/dev/groups.create
type GroupsCreateBulder struct {
	api.Params
}

// NewGroupsCreateBulder func
func NewGroupsCreateBulder() *GroupsCreateBulder {
	return &GroupsCreateBulder{api.Params{}}
}

// Title Community title.
func (b *GroupsCreateBulder) Title(v string) {
	b.Params["title"] = v
}

// Description Community description (ignored for 'type' = 'public').
func (b *GroupsCreateBulder) Description(v string) {
	b.Params["description"] = v
}

// Type Community type. Possible values: *'group' – group,, *'event' – event,, *'public' – public page
func (b *GroupsCreateBulder) Type(v string) {
	b.Params["type"] = v
}

// PublicCategory Category ID (for 'type' = 'public' only).
func (b *GroupsCreateBulder) PublicCategory(v int) {
	b.Params["public_category"] = v
}

// Subtype Public page subtype. Possible values: *'1' – place or small business,, *'2' – company, organization or website,, *'3' – famous person or group of people,, *'4' – product or work of art.
func (b *GroupsCreateBulder) Subtype(v int) {
	b.Params["subtype"] = v
}

// GroupsDeleteCallbackServerBulder builder
//
// https://vk.com/dev/groups.deleteCallbackServer
type GroupsDeleteCallbackServerBulder struct {
	api.Params
}

// NewGroupsDeleteCallbackServerBulder func
func NewGroupsDeleteCallbackServerBulder() *GroupsDeleteCallbackServerBulder {
	return &GroupsDeleteCallbackServerBulder{api.Params{}}
}

// GroupID parameter
func (b *GroupsDeleteCallbackServerBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// ServerID parameter
func (b *GroupsDeleteCallbackServerBulder) ServerID(v int) {
	b.Params["server_id"] = v
}

// GroupsDeleteLinkBulder builder
//
// Allows to delete a link from the community.
//
// https://vk.com/dev/groups.deleteLink
type GroupsDeleteLinkBulder struct {
	api.Params
}

// NewGroupsDeleteLinkBulder func
func NewGroupsDeleteLinkBulder() *GroupsDeleteLinkBulder {
	return &GroupsDeleteLinkBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsDeleteLinkBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// LinkID Link ID.
func (b *GroupsDeleteLinkBulder) LinkID(v int) {
	b.Params["link_id"] = v
}

// GroupsDisableOnlineBulder builder
//
// https://vk.com/dev/groups.disableOnline
type GroupsDisableOnlineBulder struct {
	api.Params
}

// NewGroupsDisableOnlineBulder func
func NewGroupsDisableOnlineBulder() *GroupsDisableOnlineBulder {
	return &GroupsDisableOnlineBulder{api.Params{}}
}

// GroupID parameter
func (b *GroupsDisableOnlineBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// GroupsEditBulder builder
//
// Edits a community.
//
// https://vk.com/dev/groups.edit
type GroupsEditBulder struct {
	api.Params
}

// NewGroupsEditBulder func
func NewGroupsEditBulder() *GroupsEditBulder {
	return &GroupsEditBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsEditBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Title Community title.
func (b *GroupsEditBulder) Title(v string) {
	b.Params["title"] = v
}

// Description Community description.
func (b *GroupsEditBulder) Description(v string) {
	b.Params["description"] = v
}

// ScreenName Community screen name.
func (b *GroupsEditBulder) ScreenName(v string) {
	b.Params["screen_name"] = v
}

// Access Community type. Possible values: *'0' – open,, *'1' – closed,, *'2' – private.
func (b *GroupsEditBulder) Access(v int) {
	b.Params["access"] = v
}

// Website Website that will be displayed in the community information field.
func (b *GroupsEditBulder) Website(v string) {
	b.Params["website"] = v
}

// Subject Community subject. Possible values: , *'1' – auto/moto,, *'2' – activity holidays,, *'3' – business,, *'4' – pets,, *'5' – health,, *'6' – dating and communication, , *'7' – games,, *'8' – IT (computers and software),, *'9' – cinema,, *'10' – beauty and fashion,, *'11' – cooking,, *'12' – art and culture,, *'13' – literature,, *'14' – mobile services and internet,, *'15' – music,, *'16' – science and technology,, *'17' – real estate,, *'18' – news and media,, *'19' – security,, *'20' – education,, *'21' – home and renovations,, *'22' – politics,, *'23' – food,, *'24' – industry,, *'25' – travel,, *'26' – work,, *'27' – entertainment,, *'28' – religion,, *'29' – family,, *'30' – sports,, *'31' – insurance,, *'32' – television,, *'33' – goods and services,, *'34' – hobbies,, *'35' – finance,, *'36' – photo,, *'37' – esoterics,, *'38' – electronics and appliances,, *'39' – erotic,, *'40' – humor,, *'41' – society, humanities,, *'42' – design and graphics.
func (b *GroupsEditBulder) Subject(v string) {
	b.Params["subject"] = v
}

// Email Organizer email (for events).
func (b *GroupsEditBulder) Email(v string) {
	b.Params["email"] = v
}

// Phone Organizer phone number (for events).
func (b *GroupsEditBulder) Phone(v string) {
	b.Params["phone"] = v
}

// Rss RSS feed address for import (available only to communities with special permission. Contact vk.com/support to get it.
func (b *GroupsEditBulder) Rss(v string) {
	b.Params["rss"] = v
}

// EventStartDate Event start date in Unixtime format.
func (b *GroupsEditBulder) EventStartDate(v int) {
	b.Params["event_start_date"] = v
}

// EventFinishDate Event finish date in Unixtime format.
func (b *GroupsEditBulder) EventFinishDate(v int) {
	b.Params["event_finish_date"] = v
}

// EventGroupID Organizer community ID (for events only).
func (b *GroupsEditBulder) EventGroupID(v int) {
	b.Params["event_group_id"] = v
}

// PublicCategory Public page category ID.
func (b *GroupsEditBulder) PublicCategory(v int) {
	b.Params["public_category"] = v
}

// PublicSubcategory Public page subcategory ID.
func (b *GroupsEditBulder) PublicSubcategory(v int) {
	b.Params["public_subcategory"] = v
}

// PublicDate Founding date of a company or organization owning the community in "dd.mm.YYYY" format.
func (b *GroupsEditBulder) PublicDate(v string) {
	b.Params["public_date"] = v
}

// Wall Wall settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (groups and events only),, *'3' – closed (groups and events only).
func (b *GroupsEditBulder) Wall(v int) {
	b.Params["wall"] = v
}

// Topics Board topics settings. Possbile values: , *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
func (b *GroupsEditBulder) Topics(v int) {
	b.Params["topics"] = v
}

// Photos Photos settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
func (b *GroupsEditBulder) Photos(v int) {
	b.Params["photos"] = v
}

// Video Video settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
func (b *GroupsEditBulder) Video(v int) {
	b.Params["video"] = v
}

// Audio Audio settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
func (b *GroupsEditBulder) Audio(v int) {
	b.Params["audio"] = v
}

// Links Links settings (for public pages only). Possible values: *'0' – disabled,, *'1' – enabled.
func (b *GroupsEditBulder) Links(v bool) {
	b.Params["links"] = v
}

// Events Events settings (for public pages only). Possible values: *'0' – disabled,, *'1' – enabled.
func (b *GroupsEditBulder) Events(v bool) {
	b.Params["events"] = v
}

// Places Places settings (for public pages only). Possible values: *'0' – disabled,, *'1' – enabled.
func (b *GroupsEditBulder) Places(v bool) {
	b.Params["places"] = v
}

// Contacts Contacts settings (for public pages only). Possible values: *'0' – disabled,, *'1' – enabled.
func (b *GroupsEditBulder) Contacts(v bool) {
	b.Params["contacts"] = v
}

// Docs Documents settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
func (b *GroupsEditBulder) Docs(v int) {
	b.Params["docs"] = v
}

// Wiki Wiki pages settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
func (b *GroupsEditBulder) Wiki(v int) {
	b.Params["wiki"] = v
}

// Messages Community messages. Possible values: *'0' — disabled,, *'1' — enabled.
func (b *GroupsEditBulder) Messages(v bool) {
	b.Params["messages"] = v
}

// Articles parameter
func (b *GroupsEditBulder) Articles(v bool) {
	b.Params["articles"] = v
}

// Addresses parameter
func (b *GroupsEditBulder) Addresses(v bool) {
	b.Params["addresses"] = v
}

// AgeLimits Community age limits. Possible values: *'1' — no limits,, *'2' — 16+,, *'3' — 18+.
func (b *GroupsEditBulder) AgeLimits(v int) {
	b.Params["age_limits"] = v
}

// Market Market settings. Possible values: *'0' – disabled,, *'1' – enabled.
func (b *GroupsEditBulder) Market(v bool) {
	b.Params["market"] = v
}

// MarketComments market comments settings. Possible values: *'0' – disabled,, *'1' – enabled.
func (b *GroupsEditBulder) MarketComments(v bool) {
	b.Params["market_comments"] = v
}

// MarketCountry Market delivery countries.
func (b *GroupsEditBulder) MarketCountry(v []int) {
	b.Params["market_country"] = v
}

// MarketCity Market delivery cities (if only one country is specified).
func (b *GroupsEditBulder) MarketCity(v []int) {
	b.Params["market_city"] = v
}

// MarketCurrency Market currency settings. Possbile values: , *'643' – Russian rubles,, *'980' – Ukrainian hryvnia,, *'398' – Kazakh tenge,, *'978' – Euro,, *'840' – US dollars
func (b *GroupsEditBulder) MarketCurrency(v int) {
	b.Params["market_currency"] = v
}

// MarketContact Seller contact for market. Set '0' for community messages.
func (b *GroupsEditBulder) MarketContact(v int) {
	b.Params["market_contact"] = v
}

// MarketWiki ID of a wiki page with market description.
func (b *GroupsEditBulder) MarketWiki(v int) {
	b.Params["market_wiki"] = v
}

// ObsceneFilter Obscene expressions filter in comments. Possible values: , *'0' – disabled,, *'1' – enabled.
func (b *GroupsEditBulder) ObsceneFilter(v bool) {
	b.Params["obscene_filter"] = v
}

// ObsceneStopwords Stopwords filter in comments. Possible values: , *'0' – disabled,, *'1' – enabled.
func (b *GroupsEditBulder) ObsceneStopwords(v bool) {
	b.Params["obscene_stopwords"] = v
}

// ObsceneWords Keywords for stopwords filter.
func (b *GroupsEditBulder) ObsceneWords(v []string) {
	b.Params["obscene_words"] = v
}

// MainSection parameter
func (b *GroupsEditBulder) MainSection(v int) {
	b.Params["main_section"] = v
}

// SecondarySection parameter
func (b *GroupsEditBulder) SecondarySection(v int) {
	b.Params["secondary_section"] = v
}

// Country Country of the community.
func (b *GroupsEditBulder) Country(v int) {
	b.Params["country"] = v
}

// City City of the community.
func (b *GroupsEditBulder) City(v int) {
	b.Params["city"] = v
}

// GroupsEditAddressBulder builder
//
// https://vk.com/dev/groups.editAddress
type GroupsEditAddressBulder struct {
	api.Params
}

// NewGroupsEditAddressBulder func
func NewGroupsEditAddressBulder() *GroupsEditAddressBulder {
	return &GroupsEditAddressBulder{api.Params{}}
}

// GroupID parameter
func (b *GroupsEditAddressBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// AddressID parameter
func (b *GroupsEditAddressBulder) AddressID(v int) {
	b.Params["address_id"] = v
}

// Title parameter
func (b *GroupsEditAddressBulder) Title(v string) {
	b.Params["title"] = v
}

// Address parameter
func (b *GroupsEditAddressBulder) Address(v string) {
	b.Params["address"] = v
}

// AdditionalAddress parameter
func (b *GroupsEditAddressBulder) AdditionalAddress(v string) {
	b.Params["additional_address"] = v
}

// CountryID parameter
func (b *GroupsEditAddressBulder) CountryID(v int) {
	b.Params["country_id"] = v
}

// CityID parameter
func (b *GroupsEditAddressBulder) CityID(v int) {
	b.Params["city_id"] = v
}

// MetroID parameter
func (b *GroupsEditAddressBulder) MetroID(v int) {
	b.Params["metro_id"] = v
}

// Latitude parameter
func (b *GroupsEditAddressBulder) Latitude(v float64) {
	b.Params["latitude"] = v
}

// Longitude parameter
func (b *GroupsEditAddressBulder) Longitude(v float64) {
	b.Params["longitude"] = v
}

// Phone parameter
func (b *GroupsEditAddressBulder) Phone(v string) {
	b.Params["phone"] = v
}

// WorkInfoStatus parameter
func (b *GroupsEditAddressBulder) WorkInfoStatus(v string) {
	b.Params["work_info_status"] = v
}

// Timetable parameter
func (b *GroupsEditAddressBulder) Timetable(v string) {
	b.Params["timetable"] = v
}

// IsMainAddress parameter
func (b *GroupsEditAddressBulder) IsMainAddress(v bool) {
	b.Params["is_main_address"] = v
}

// GroupsEditCallbackServerBulder builder
//
// https://vk.com/dev/groups.editCallbackServer
type GroupsEditCallbackServerBulder struct {
	api.Params
}

// NewGroupsEditCallbackServerBulder func
func NewGroupsEditCallbackServerBulder() *GroupsEditCallbackServerBulder {
	return &GroupsEditCallbackServerBulder{api.Params{}}
}

// GroupID parameter
func (b *GroupsEditCallbackServerBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// ServerID parameter
func (b *GroupsEditCallbackServerBulder) ServerID(v int) {
	b.Params["server_id"] = v
}

// URL parameter
func (b *GroupsEditCallbackServerBulder) URL(v string) {
	b.Params["url"] = v
}

// Title parameter
func (b *GroupsEditCallbackServerBulder) Title(v string) {
	b.Params["title"] = v
}

// SecretKey parameter
func (b *GroupsEditCallbackServerBulder) SecretKey(v string) {
	b.Params["secret_key"] = v
}

// GroupsEditLinkBulder builder
//
// Allows to edit a link in the community.
//
// https://vk.com/dev/groups.editLink
type GroupsEditLinkBulder struct {
	api.Params
}

// NewGroupsEditLinkBulder func
func NewGroupsEditLinkBulder() *GroupsEditLinkBulder {
	return &GroupsEditLinkBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsEditLinkBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// LinkID Link ID.
func (b *GroupsEditLinkBulder) LinkID(v int) {
	b.Params["link_id"] = v
}

// Text New description text for the link.
func (b *GroupsEditLinkBulder) Text(v string) {
	b.Params["text"] = v
}

// GroupsEditManagerBulder builder
//
// Allows to add, remove or edit the community manager.
//
// https://vk.com/dev/groups.editManager
type GroupsEditManagerBulder struct {
	api.Params
}

// NewGroupsEditManagerBulder func
func NewGroupsEditManagerBulder() *GroupsEditManagerBulder {
	return &GroupsEditManagerBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsEditManagerBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID User ID.
func (b *GroupsEditManagerBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Role Manager role. Possible values: *'moderator',, *'editor',, *'administrator'.
func (b *GroupsEditManagerBulder) Role(v string) {
	b.Params["role"] = v
}

// IsContact '1' — to show the manager in Contacts block of the community.
func (b *GroupsEditManagerBulder) IsContact(v bool) {
	b.Params["is_contact"] = v
}

// ContactPosition Position to show in Contacts block.
func (b *GroupsEditManagerBulder) ContactPosition(v string) {
	b.Params["contact_position"] = v
}

// ContactPhone Contact phone.
func (b *GroupsEditManagerBulder) ContactPhone(v string) {
	b.Params["contact_phone"] = v
}

// ContactEmail Contact e-mail.
func (b *GroupsEditManagerBulder) ContactEmail(v string) {
	b.Params["contact_email"] = v
}

// GroupsEnableOnlineBulder builder
//
// https://vk.com/dev/groups.enableOnline
type GroupsEnableOnlineBulder struct {
	api.Params
}

// NewGroupsEnableOnlineBulder func
func NewGroupsEnableOnlineBulder() *GroupsEnableOnlineBulder {
	return &GroupsEnableOnlineBulder{api.Params{}}
}

// GroupID parameter
func (b *GroupsEnableOnlineBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// GroupsGetBulder builder
//
// Returns a list of the communities to which a user belongs.
//
// https://vk.com/dev/groups.get
type GroupsGetBulder struct {
	api.Params
}

// NewGroupsGetBulder func
func NewGroupsGetBulder() *GroupsGetBulder {
	return &GroupsGetBulder{api.Params{}}
}

// UserID User ID.
func (b *GroupsGetBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Extended '1' — to return complete information about a user's communities, '0' — to return a list of community IDs without any additional fields (default),
func (b *GroupsGetBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Filter Types of communities to return: 'admin' — to return communities administered by the user , 'editor' — to return communities where the user is an administrator or editor, 'moder' — to return communities where the user is an administrator, editor, or moderator, 'groups' — to return only groups, 'publics' — to return only public pages, 'events' — to return only events
func (b *GroupsGetBulder) Filter(v []string) {
	b.Params["filter"] = v
}

// Fields Profile fields to return.
func (b *GroupsGetBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// Offset Offset needed to return a specific subset of communities.
func (b *GroupsGetBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of communities to return.
func (b *GroupsGetBulder) Count(v int) {
	b.Params["count"] = v
}

// GroupsGetAddressesBulder builder
//
// Returns a list of community addresses.
//
// https://vk.com/dev/groups.getAddresses
type GroupsGetAddressesBulder struct {
	api.Params
}

// NewGroupsGetAddressesBulder func
func NewGroupsGetAddressesBulder() *GroupsGetAddressesBulder {
	return &GroupsGetAddressesBulder{api.Params{}}
}

// GroupID ID or screen name of the community.
func (b *GroupsGetAddressesBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// AddressIDs parameter
func (b *GroupsGetAddressesBulder) AddressIDs(v []int) {
	b.Params["address_ids"] = v
}

// Latitude Latitude of  the user geo position.
func (b *GroupsGetAddressesBulder) Latitude(v float64) {
	b.Params["latitude"] = v
}

// Longitude Longitude of the user geo position.
func (b *GroupsGetAddressesBulder) Longitude(v float64) {
	b.Params["longitude"] = v
}

// Offset Offset needed to return a specific subset of community addresses.
func (b *GroupsGetAddressesBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of community addresses to return.
func (b *GroupsGetAddressesBulder) Count(v int) {
	b.Params["count"] = v
}

// Fields Address fields
func (b *GroupsGetAddressesBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupsGetBannedBulder builder
//
// Returns a list of users on a community blacklist.
//
// https://vk.com/dev/groups.getBanned
type GroupsGetBannedBulder struct {
	api.Params
}

// NewGroupsGetBannedBulder func
func NewGroupsGetBannedBulder() *GroupsGetBannedBulder {
	return &GroupsGetBannedBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsGetBannedBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Offset Offset needed to return a specific subset of users.
func (b *GroupsGetBannedBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of users to return.
func (b *GroupsGetBannedBulder) Count(v int) {
	b.Params["count"] = v
}

// Fields parameter
func (b *GroupsGetBannedBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// OwnerID parameter
func (b *GroupsGetBannedBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// GroupsGetByIDBulder builder
//
// Returns information about communities by their IDs.
//
// https://vk.com/dev/groups.getById
type GroupsGetByIDBulder struct {
	api.Params
}

// NewGroupsGetByIDBulder func
func NewGroupsGetByIDBulder() *GroupsGetByIDBulder {
	return &GroupsGetByIDBulder{api.Params{}}
}

// GroupIDs IDs or screen names of communities.
func (b *GroupsGetByIDBulder) GroupIDs(v []string) {
	b.Params["group_ids"] = v
}

// GroupID ID or screen name of the community.
func (b *GroupsGetByIDBulder) GroupID(v string) {
	b.Params["group_id"] = v
}

// Fields Group fields to return.
func (b *GroupsGetByIDBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupsGetCallbackConfirmationCodeBulder builder
//
// Returns Callback API confirmation code for the community.
//
// https://vk.com/dev/groups.getCallbackConfirmationCode
type GroupsGetCallbackConfirmationCodeBulder struct {
	api.Params
}

// NewGroupsGetCallbackConfirmationCodeBulder func
func NewGroupsGetCallbackConfirmationCodeBulder() *GroupsGetCallbackConfirmationCodeBulder {
	return &GroupsGetCallbackConfirmationCodeBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsGetCallbackConfirmationCodeBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// GroupsGetCallbackServersBulder builder
//
// https://vk.com/dev/groups.getCallbackServers
type GroupsGetCallbackServersBulder struct {
	api.Params
}

// NewGroupsGetCallbackServersBulder func
func NewGroupsGetCallbackServersBulder() *GroupsGetCallbackServersBulder {
	return &GroupsGetCallbackServersBulder{api.Params{}}
}

// GroupID parameter
func (b *GroupsGetCallbackServersBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// ServerIDs parameter
func (b *GroupsGetCallbackServersBulder) ServerIDs(v []int) {
	b.Params["server_ids"] = v
}

// GroupsGetCallbackSettingsBulder builder
//
// Returns [vk.com/dev/callback_api|Callback API] notifications settings.
//
// https://vk.com/dev/groups.getCallbackSettings
type GroupsGetCallbackSettingsBulder struct {
	api.Params
}

// NewGroupsGetCallbackSettingsBulder func
func NewGroupsGetCallbackSettingsBulder() *GroupsGetCallbackSettingsBulder {
	return &GroupsGetCallbackSettingsBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsGetCallbackSettingsBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// ServerID Server ID.
func (b *GroupsGetCallbackSettingsBulder) ServerID(v int) {
	b.Params["server_id"] = v
}

// GroupsGetCatalogBulder builder
//
// Returns communities list for a catalog category.
//
// https://vk.com/dev/groups.getCatalog
type GroupsGetCatalogBulder struct {
	api.Params
}

// NewGroupsGetCatalogBulder func
func NewGroupsGetCatalogBulder() *GroupsGetCatalogBulder {
	return &GroupsGetCatalogBulder{api.Params{}}
}

// CategoryID Category id received from [vk.com/dev/groups.getCatalogInfo|groups.getCatalogInfo].
func (b *GroupsGetCatalogBulder) CategoryID(v int) {
	b.Params["category_id"] = v
}

// SubcategoryID Subcategory id received from [vk.com/dev/groups.getCatalogInfo|groups.getCatalogInfo].
func (b *GroupsGetCatalogBulder) SubcategoryID(v int) {
	b.Params["subcategory_id"] = v
}

// GroupsGetCatalogInfoBulder builder
//
// Returns categories list for communities catalog
//
// https://vk.com/dev/groups.getCatalogInfo
type GroupsGetCatalogInfoBulder struct {
	api.Params
}

// NewGroupsGetCatalogInfoBulder func
func NewGroupsGetCatalogInfoBulder() *GroupsGetCatalogInfoBulder {
	return &GroupsGetCatalogInfoBulder{api.Params{}}
}

// Extended 1 – to return communities count and three communities for preview. By default: 0.
func (b *GroupsGetCatalogInfoBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Subcategories 1 – to return subcategories info. By default: 0.
func (b *GroupsGetCatalogInfoBulder) Subcategories(v bool) {
	b.Params["subcategories"] = v
}

// GroupsGetInvitedUsersBulder builder
//
// Returns invited users list of a community
//
// https://vk.com/dev/groups.getInvitedUsers
type GroupsGetInvitedUsersBulder struct {
	api.Params
}

// NewGroupsGetInvitedUsersBulder func
func NewGroupsGetInvitedUsersBulder() *GroupsGetInvitedUsersBulder {
	return &GroupsGetInvitedUsersBulder{api.Params{}}
}

// GroupID Group ID to return invited users for.
func (b *GroupsGetInvitedUsersBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *GroupsGetInvitedUsersBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of results to return.
func (b *GroupsGetInvitedUsersBulder) Count(v int) {
	b.Params["count"] = v
}

// Fields List of additional fields to be returned. Available values: 'sex, bdate, city, country, photo_50, photo_100, photo_200_orig, photo_200, photo_400_orig, photo_max, photo_max_orig, online, online_mobile, lists, domain, has_mobile, contacts, connections, site, education, universities, schools, can_post, can_see_all_posts, can_see_audio, can_write_private_message, status, last_seen, common_count, relation, relatives, counters'.
func (b *GroupsGetInvitedUsersBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase Case for declension of user name and surname. Possible values: *'nom' — nominative (default),, *'gen' — genitive,, *'dat' — dative,, *'acc' — accusative, , *'ins' — instrumental,, *'abl' — prepositional.
func (b *GroupsGetInvitedUsersBulder) NameCase(v string) {
	b.Params["name_case"] = v
}

// GroupsGetInvitesBulder builder
//
// Returns a list of invitations to join communities and events.
//
// https://vk.com/dev/groups.getInvites
type GroupsGetInvitesBulder struct {
	api.Params
}

// NewGroupsGetInvitesBulder func
func NewGroupsGetInvitesBulder() *GroupsGetInvitesBulder {
	return &GroupsGetInvitesBulder{api.Params{}}
}

// Offset Offset needed to return a specific subset of invitations.
func (b *GroupsGetInvitesBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of invitations to return.
func (b *GroupsGetInvitesBulder) Count(v int) {
	b.Params["count"] = v
}

// Extended '1' — to return additional [vk.com/dev/fields_groups|fields] for communities..
func (b *GroupsGetInvitesBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// GroupsGetLongPollServerBulder builder
//
// Returns the data needed to query a Long Poll server for events
//
// https://vk.com/dev/groups.getLongPollServer
type GroupsGetLongPollServerBulder struct {
	api.Params
}

// NewGroupsGetLongPollServerBulder func
func NewGroupsGetLongPollServerBulder() *GroupsGetLongPollServerBulder {
	return &GroupsGetLongPollServerBulder{api.Params{}}
}

// GroupID Community ID
func (b *GroupsGetLongPollServerBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// GroupsGetLongPollSettingsBulder builder
//
// Returns Long Poll notification settings
//
// https://vk.com/dev/groups.getLongPollSettings
type GroupsGetLongPollSettingsBulder struct {
	api.Params
}

// NewGroupsGetLongPollSettingsBulder func
func NewGroupsGetLongPollSettingsBulder() *GroupsGetLongPollSettingsBulder {
	return &GroupsGetLongPollSettingsBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsGetLongPollSettingsBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// GroupsGetMembersBulder builder
//
// Returns a list of community members.
//
// https://vk.com/dev/groups.getMembers
type GroupsGetMembersBulder struct {
	api.Params
}

// NewGroupsGetMembersBulder func
func NewGroupsGetMembersBulder() *GroupsGetMembersBulder {
	return &GroupsGetMembersBulder{api.Params{}}
}

// GroupID ID or screen name of the community.
func (b *GroupsGetMembersBulder) GroupID(v string) {
	b.Params["group_id"] = v
}

// Sort Sort order. Available values: 'id_asc', 'id_desc', 'time_asc', 'time_desc'. 'time_asc' and 'time_desc' are availavle only if the method is called by the group's 'moderator'.
func (b *GroupsGetMembersBulder) Sort(v string) {
	b.Params["sort"] = v
}

// Offset Offset needed to return a specific subset of community members.
func (b *GroupsGetMembersBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of community members to return.
func (b *GroupsGetMembersBulder) Count(v int) {
	b.Params["count"] = v
}

// Fields List of additional fields to be returned. Available values: 'sex, bdate, city, country, photo_50, photo_100, photo_200_orig, photo_200, photo_400_orig, photo_max, photo_max_orig, online, online_mobile, lists, domain, has_mobile, contacts, connections, site, education, universities, schools, can_post, can_see_all_posts, can_see_audio, can_write_private_message, status, last_seen, common_count, relation, relatives, counters'.
func (b *GroupsGetMembersBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// Filter *'friends' – only friends in this community will be returned,, *'unsure' – only those who pressed 'I may attend' will be returned (if it's an event).
func (b *GroupsGetMembersBulder) Filter(v string) {
	b.Params["filter"] = v
}

// GroupsGetRequestsBulder builder
//
// Returns a list of requests to the community.
//
// https://vk.com/dev/groups.getRequests
type GroupsGetRequestsBulder struct {
	api.Params
}

// NewGroupsGetRequestsBulder func
func NewGroupsGetRequestsBulder() *GroupsGetRequestsBulder {
	return &GroupsGetRequestsBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsGetRequestsBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *GroupsGetRequestsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of results to return.
func (b *GroupsGetRequestsBulder) Count(v int) {
	b.Params["count"] = v
}

// Fields Profile fields to return.
func (b *GroupsGetRequestsBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupsGetSettingsBulder builder
//
// Returns community settings.
//
// https://vk.com/dev/groups.getSettings
type GroupsGetSettingsBulder struct {
	api.Params
}

// NewGroupsGetSettingsBulder func
func NewGroupsGetSettingsBulder() *GroupsGetSettingsBulder {
	return &GroupsGetSettingsBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsGetSettingsBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// GroupsInviteBulder builder
//
// Allows to invite friends to the community.
//
// https://vk.com/dev/groups.invite
type GroupsInviteBulder struct {
	api.Params
}

// NewGroupsInviteBulder func
func NewGroupsInviteBulder() *GroupsInviteBulder {
	return &GroupsInviteBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsInviteBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID User ID.
func (b *GroupsInviteBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupsIsMemberBulder builder
//
// Returns information specifying whether a user is a member of a community.
//
// https://vk.com/dev/groups.isMember
type GroupsIsMemberBulder struct {
	api.Params
}

// NewGroupsIsMemberBulder func
func NewGroupsIsMemberBulder() *GroupsIsMemberBulder {
	return &GroupsIsMemberBulder{api.Params{}}
}

// GroupID ID or screen name of the community.
func (b *GroupsIsMemberBulder) GroupID(v string) {
	b.Params["group_id"] = v
}

// UserID User ID.
func (b *GroupsIsMemberBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// UserIDs User IDs.
func (b *GroupsIsMemberBulder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// Extended '1' — to return an extended response with additional fields. By default: '0'.
func (b *GroupsIsMemberBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// GroupsJoinBulder builder
//
// With this method you can join the group or public page, and also confirm your participation in an event.
//
// https://vk.com/dev/groups.join
type GroupsJoinBulder struct {
	api.Params
}

// NewGroupsJoinBulder func
func NewGroupsJoinBulder() *GroupsJoinBulder {
	return &GroupsJoinBulder{api.Params{}}
}

// GroupID ID or screen name of the community.
func (b *GroupsJoinBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// NotSure Optional parameter which is taken into account when 'gid' belongs to the event: '1' — Perhaps I will attend, '0' — I will be there for sure (default), ,
func (b *GroupsJoinBulder) NotSure(v string) {
	b.Params["not_sure"] = v
}

// GroupsLeaveBulder builder
//
// With this method you can leave a group, public page, or event.
//
// https://vk.com/dev/groups.leave
type GroupsLeaveBulder struct {
	api.Params
}

// NewGroupsLeaveBulder func
func NewGroupsLeaveBulder() *GroupsLeaveBulder {
	return &GroupsLeaveBulder{api.Params{}}
}

// GroupID ID or screen name of the community.
func (b *GroupsLeaveBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// GroupsRemoveUserBulder builder
//
// Removes a user from the community.
//
// https://vk.com/dev/groups.removeUser
type GroupsRemoveUserBulder struct {
	api.Params
}

// NewGroupsRemoveUserBulder func
func NewGroupsRemoveUserBulder() *GroupsRemoveUserBulder {
	return &GroupsRemoveUserBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsRemoveUserBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID User ID.
func (b *GroupsRemoveUserBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupsReorderLinkBulder builder
//
// Allows to reorder links in the community.
//
// https://vk.com/dev/groups.reorderLink
type GroupsReorderLinkBulder struct {
	api.Params
}

// NewGroupsReorderLinkBulder func
func NewGroupsReorderLinkBulder() *GroupsReorderLinkBulder {
	return &GroupsReorderLinkBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsReorderLinkBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// LinkID Link ID.
func (b *GroupsReorderLinkBulder) LinkID(v int) {
	b.Params["link_id"] = v
}

// After ID of the link after which to place the link with 'link_id'.
func (b *GroupsReorderLinkBulder) After(v int) {
	b.Params["after"] = v
}

// GroupsSearchBulder builder
//
// Returns a list of communities matching the search criteria.
//
// https://vk.com/dev/groups.search
type GroupsSearchBulder struct {
	api.Params
}

// NewGroupsSearchBulder func
func NewGroupsSearchBulder() *GroupsSearchBulder {
	return &GroupsSearchBulder{api.Params{}}
}

// Q Search query string.
func (b *GroupsSearchBulder) Q(v string) {
	b.Params["q"] = v
}

// Type Community type. Possible values: 'group, page, event.'
func (b *GroupsSearchBulder) Type(v string) {
	b.Params["type"] = v
}

// CountryID Country ID.
func (b *GroupsSearchBulder) CountryID(v int) {
	b.Params["country_id"] = v
}

// CityID City ID. If this parameter is transmitted, country_id is ignored.
func (b *GroupsSearchBulder) CityID(v int) {
	b.Params["city_id"] = v
}

// Future '1' — to return only upcoming events. Works with the 'type' = 'event' only.
func (b *GroupsSearchBulder) Future(v bool) {
	b.Params["future"] = v
}

// Market '1' — to return communities with enabled market only.
func (b *GroupsSearchBulder) Market(v bool) {
	b.Params["market"] = v
}

// Sort Sort order. Possible values: *'0' — default sorting (similar the full version of the site),, *'1' — by growth speed,, *'2'— by the "day attendance/members number" ratio,, *'3' — by the "Likes number/members number" ratio,, *'4' — by the "comments number/members number" ratio,, *'5' — by the "boards entries number/members number" ratio.
func (b *GroupsSearchBulder) Sort(v int) {
	b.Params["sort"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *GroupsSearchBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of communities to return. "Note that you can not receive more than first thousand of results, regardless of 'count' and 'offset' values."
func (b *GroupsSearchBulder) Count(v int) {
	b.Params["count"] = v
}

// GroupsSetCallbackSettingsBulder builder
//
// Allow to set notifications settings for group.
//
// https://vk.com/dev/groups.setCallbackSettings
type GroupsSetCallbackSettingsBulder struct {
	api.Params
}

// NewGroupsSetCallbackSettingsBulder func
func NewGroupsSetCallbackSettingsBulder() *GroupsSetCallbackSettingsBulder {
	return &GroupsSetCallbackSettingsBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsSetCallbackSettingsBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// ServerID Server ID.
func (b *GroupsSetCallbackSettingsBulder) ServerID(v int) {
	b.Params["server_id"] = v
}

// APIVersion parameter
func (b *GroupsSetCallbackSettingsBulder) APIVersion(v string) {
	b.Params["api_version"] = v
}

// MessageNew A new incoming message has been received ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) MessageNew(v bool) {
	b.Params["message_new"] = v
}

// MessageReply A new outcoming message has been received ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) MessageReply(v bool) {
	b.Params["message_reply"] = v
}

// MessageAllow Allowed messages notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) MessageAllow(v bool) {
	b.Params["message_allow"] = v
}

// MessageEdit parameter
func (b *GroupsSetCallbackSettingsBulder) MessageEdit(v bool) {
	b.Params["message_edit"] = v
}

// MessageDeny Denied messages notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) MessageDeny(v bool) {
	b.Params["message_deny"] = v
}

// MessageTypingState parameter
func (b *GroupsSetCallbackSettingsBulder) MessageTypingState(v bool) {
	b.Params["message_typing_state"] = v
}

// PhotoNew New photos notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) PhotoNew(v bool) {
	b.Params["photo_new"] = v
}

// AudioNew New audios notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) AudioNew(v bool) {
	b.Params["audio_new"] = v
}

// VideoNew New videos notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) VideoNew(v bool) {
	b.Params["video_new"] = v
}

// WallReplyNew New wall replies notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) WallReplyNew(v bool) {
	b.Params["wall_reply_new"] = v
}

// WallReplyEdit Wall replies edited notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) WallReplyEdit(v bool) {
	b.Params["wall_reply_edit"] = v
}

// WallReplyDelete A wall comment has been deleted ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) WallReplyDelete(v bool) {
	b.Params["wall_reply_delete"] = v
}

// WallReplyRestore A wall comment has been restored ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) WallReplyRestore(v bool) {
	b.Params["wall_reply_restore"] = v
}

// WallPostNew New wall posts notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) WallPostNew(v bool) {
	b.Params["wall_post_new"] = v
}

// WallRepost New wall posts notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) WallRepost(v bool) {
	b.Params["wall_repost"] = v
}

// BoardPostNew New board posts notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) BoardPostNew(v bool) {
	b.Params["board_post_new"] = v
}

// BoardPostEdit Board posts edited notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) BoardPostEdit(v bool) {
	b.Params["board_post_edit"] = v
}

// BoardPostRestore Board posts restored notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) BoardPostRestore(v bool) {
	b.Params["board_post_restore"] = v
}

// BoardPostDelete Board posts deleted notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) BoardPostDelete(v bool) {
	b.Params["board_post_delete"] = v
}

// PhotoCommentNew New comment to photo notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) PhotoCommentNew(v bool) {
	b.Params["photo_comment_new"] = v
}

// PhotoCommentEdit A photo comment has been edited ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) PhotoCommentEdit(v bool) {
	b.Params["photo_comment_edit"] = v
}

// PhotoCommentDelete A photo comment has been deleted ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) PhotoCommentDelete(v bool) {
	b.Params["photo_comment_delete"] = v
}

// PhotoCommentRestore A photo comment has been restored ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) PhotoCommentRestore(v bool) {
	b.Params["photo_comment_restore"] = v
}

// VideoCommentNew New comment to video notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) VideoCommentNew(v bool) {
	b.Params["video_comment_new"] = v
}

// VideoCommentEdit A video comment has been edited ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) VideoCommentEdit(v bool) {
	b.Params["video_comment_edit"] = v
}

// VideoCommentDelete A video comment has been deleted ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) VideoCommentDelete(v bool) {
	b.Params["video_comment_delete"] = v
}

// VideoCommentRestore A video comment has been restored ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) VideoCommentRestore(v bool) {
	b.Params["video_comment_restore"] = v
}

// MarketCommentNew New comment to market item notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) MarketCommentNew(v bool) {
	b.Params["market_comment_new"] = v
}

// MarketCommentEdit A market comment has been edited ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) MarketCommentEdit(v bool) {
	b.Params["market_comment_edit"] = v
}

// MarketCommentDelete A market comment has been deleted ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) MarketCommentDelete(v bool) {
	b.Params["market_comment_delete"] = v
}

// MarketCommentRestore A market comment has been restored ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) MarketCommentRestore(v bool) {
	b.Params["market_comment_restore"] = v
}

// PollVoteNew A vote in a public poll has been added ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) PollVoteNew(v bool) {
	b.Params["poll_vote_new"] = v
}

// GroupJoin Joined community notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) GroupJoin(v bool) {
	b.Params["group_join"] = v
}

// GroupLeave Left community notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetCallbackSettingsBulder) GroupLeave(v bool) {
	b.Params["group_leave"] = v
}

// GroupChangeSettings parameter
func (b *GroupsSetCallbackSettingsBulder) GroupChangeSettings(v bool) {
	b.Params["group_change_settings"] = v
}

// GroupChangePhoto parameter
func (b *GroupsSetCallbackSettingsBulder) GroupChangePhoto(v bool) {
	b.Params["group_change_photo"] = v
}

// GroupOfficersEdit parameter
func (b *GroupsSetCallbackSettingsBulder) GroupOfficersEdit(v bool) {
	b.Params["group_officers_edit"] = v
}

// UserBlock User added to community blacklist
func (b *GroupsSetCallbackSettingsBulder) UserBlock(v bool) {
	b.Params["user_block"] = v
}

// UserUnblock User removed from community blacklist
func (b *GroupsSetCallbackSettingsBulder) UserUnblock(v bool) {
	b.Params["user_unblock"] = v
}

// LeadFormsNew New form in lead forms
func (b *GroupsSetCallbackSettingsBulder) LeadFormsNew(v bool) {
	b.Params["lead_forms_new"] = v
}

// GroupsSetLongPollSettingsBulder builder
//
// Sets Long Poll notification settings
//
// https://vk.com/dev/groups.setLongPollSettings
type GroupsSetLongPollSettingsBulder struct {
	api.Params
}

// NewGroupsSetLongPollSettingsBulder func
func NewGroupsSetLongPollSettingsBulder() *GroupsSetLongPollSettingsBulder {
	return &GroupsSetLongPollSettingsBulder{api.Params{}}
}

// GroupID Community ID.
func (b *GroupsSetLongPollSettingsBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Enabled Sets whether Long Poll is enabled ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) Enabled(v bool) {
	b.Params["enabled"] = v
}

// APIVersion parameter
func (b *GroupsSetLongPollSettingsBulder) APIVersion(v string) {
	b.Params["api_version"] = v
}

// MessageNew A new incoming message has been received ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) MessageNew(v bool) {
	b.Params["message_new"] = v
}

// MessageReply A new outcoming message has been received ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) MessageReply(v bool) {
	b.Params["message_reply"] = v
}

// MessageAllow Allowed messages notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) MessageAllow(v bool) {
	b.Params["message_allow"] = v
}

// MessageDeny Denied messages notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) MessageDeny(v bool) {
	b.Params["message_deny"] = v
}

// MessageEdit A message has been edited ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) MessageEdit(v bool) {
	b.Params["message_edit"] = v
}

// MessageTypingState parameter
func (b *GroupsSetLongPollSettingsBulder) MessageTypingState(v bool) {
	b.Params["message_typing_state"] = v
}

// PhotoNew New photos notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) PhotoNew(v bool) {
	b.Params["photo_new"] = v
}

// AudioNew New audios notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) AudioNew(v bool) {
	b.Params["audio_new"] = v
}

// VideoNew New videos notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) VideoNew(v bool) {
	b.Params["video_new"] = v
}

// WallReplyNew New wall replies notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) WallReplyNew(v bool) {
	b.Params["wall_reply_new"] = v
}

// WallReplyEdit Wall replies edited notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) WallReplyEdit(v bool) {
	b.Params["wall_reply_edit"] = v
}

// WallReplyDelete A wall comment has been deleted ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) WallReplyDelete(v bool) {
	b.Params["wall_reply_delete"] = v
}

// WallReplyRestore A wall comment has been restored ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) WallReplyRestore(v bool) {
	b.Params["wall_reply_restore"] = v
}

// WallPostNew New wall posts notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) WallPostNew(v bool) {
	b.Params["wall_post_new"] = v
}

// WallRepost New wall posts notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) WallRepost(v bool) {
	b.Params["wall_repost"] = v
}

// BoardPostNew New board posts notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) BoardPostNew(v bool) {
	b.Params["board_post_new"] = v
}

// BoardPostEdit Board posts edited notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) BoardPostEdit(v bool) {
	b.Params["board_post_edit"] = v
}

// BoardPostRestore Board posts restored notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) BoardPostRestore(v bool) {
	b.Params["board_post_restore"] = v
}

// BoardPostDelete Board posts deleted notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) BoardPostDelete(v bool) {
	b.Params["board_post_delete"] = v
}

// PhotoCommentNew New comment to photo notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) PhotoCommentNew(v bool) {
	b.Params["photo_comment_new"] = v
}

// PhotoCommentEdit A photo comment has been edited ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) PhotoCommentEdit(v bool) {
	b.Params["photo_comment_edit"] = v
}

// PhotoCommentDelete A photo comment has been deleted ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) PhotoCommentDelete(v bool) {
	b.Params["photo_comment_delete"] = v
}

// PhotoCommentRestore A photo comment has been restored ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) PhotoCommentRestore(v bool) {
	b.Params["photo_comment_restore"] = v
}

// VideoCommentNew New comment to video notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) VideoCommentNew(v bool) {
	b.Params["video_comment_new"] = v
}

// VideoCommentEdit A video comment has been edited ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) VideoCommentEdit(v bool) {
	b.Params["video_comment_edit"] = v
}

// VideoCommentDelete A video comment has been deleted ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) VideoCommentDelete(v bool) {
	b.Params["video_comment_delete"] = v
}

// VideoCommentRestore A video comment has been restored ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) VideoCommentRestore(v bool) {
	b.Params["video_comment_restore"] = v
}

// MarketCommentNew New comment to market item notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) MarketCommentNew(v bool) {
	b.Params["market_comment_new"] = v
}

// MarketCommentEdit A market comment has been edited ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) MarketCommentEdit(v bool) {
	b.Params["market_comment_edit"] = v
}

// MarketCommentDelete A market comment has been deleted ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) MarketCommentDelete(v bool) {
	b.Params["market_comment_delete"] = v
}

// MarketCommentRestore A market comment has been restored ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) MarketCommentRestore(v bool) {
	b.Params["market_comment_restore"] = v
}

// PollVoteNew A vote in a public poll has been added ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) PollVoteNew(v bool) {
	b.Params["poll_vote_new"] = v
}

// GroupJoin Joined community notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) GroupJoin(v bool) {
	b.Params["group_join"] = v
}

// GroupLeave Left community notifications ('0' — disabled, '1' — enabled).
func (b *GroupsSetLongPollSettingsBulder) GroupLeave(v bool) {
	b.Params["group_leave"] = v
}

// GroupChangeSettings parameter
func (b *GroupsSetLongPollSettingsBulder) GroupChangeSettings(v bool) {
	b.Params["group_change_settings"] = v
}

// GroupChangePhoto parameter
func (b *GroupsSetLongPollSettingsBulder) GroupChangePhoto(v bool) {
	b.Params["group_change_photo"] = v
}

// GroupOfficersEdit parameter
func (b *GroupsSetLongPollSettingsBulder) GroupOfficersEdit(v bool) {
	b.Params["group_officers_edit"] = v
}

// UserBlock User added to community blacklist
func (b *GroupsSetLongPollSettingsBulder) UserBlock(v bool) {
	b.Params["user_block"] = v
}

// UserUnblock User removed from community blacklist
func (b *GroupsSetLongPollSettingsBulder) UserUnblock(v bool) {
	b.Params["user_unblock"] = v
}

// GroupsUnbanBulder builder
//
// https://vk.com/dev/groups.unban
type GroupsUnbanBulder struct {
	api.Params
}

// NewGroupsUnbanBulder func
func NewGroupsUnbanBulder() *GroupsUnbanBulder {
	return &GroupsUnbanBulder{api.Params{}}
}

// GroupID parameter
func (b *GroupsUnbanBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// OwnerID parameter
func (b *GroupsUnbanBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}
