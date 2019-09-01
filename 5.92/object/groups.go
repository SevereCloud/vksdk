package object // import "github.com/SevereCloud/vksdk/5.92/object"

import (
	"fmt"
)

// GroupsAddress struct
type GroupsAddress struct {
	AdditionalAddress string                 `json:"additional_address"` // Additional address to the place (6 floor, left door)
	Address           string                 `json:"address"`            // String address to the place (Nevsky, 28)
	CityID            int                    `json:"city_id"`            // City id of address
	CountryID         int                    `json:"country_id"`         // Country id of address
	Distance          int                    `json:"distance"`           // Distance from the point
	ID                int                    `json:"id"`                 // Address id
	Latitude          float64                `json:"latitude"`           // Address latitude
	Longitude         float64                `json:"longitude"`          // Address longitude
	MetroStationID    int                    `json:"metro_station_id"`   // Metro id of address
	Phone             string                 `json:"phone"`              // Address phone
	TimeOffset        int                    `json:"time_offset"`        // Time offset int minutes from utc time
	Timetable         groupsAddressTimetable `json:"timetable"`          // Week timetable for the address
	Title             string                 `json:"title"`              // Title of the place (Zinger, etc)
	WorkInfoStatus    string                 `json:"work_info_status"`   // Status of information about timetable
}

// groupsAddressTimetable Timetable for a week
type groupsAddressTimetable struct {
	Fri groupsAddressTimetableDay `json:"fri"` // Timetable for friday
	Mon groupsAddressTimetableDay `json:"mon"` // Timetable for monday
	Sat groupsAddressTimetableDay `json:"sat"` // Timetable for saturday
	Sun groupsAddressTimetableDay `json:"sun"` // Timetable for sunday
	Thu groupsAddressTimetableDay `json:"thu"` // Timetable for thursday
	Tue groupsAddressTimetableDay `json:"tue"` // Timetable for tuesday
	Wed groupsAddressTimetableDay `json:"wed"` // Timetable for wednesday
}

// groupsAddressTimetableDay Timetable for one day
type groupsAddressTimetableDay struct {
	BreakCloseTime int `json:"break_close_time"` // Close time of the break in minutes
	BreakOpenTime  int `json:"break_open_time"`  // Start time of the break in minutes
	CloseTime      int `json:"close_time"`       // Close time in minutes
	OpenTime       int `json:"open_time"`        // Open time in minutes
}

// GroupsAddressesInfo struct
type GroupsAddressesInfo struct {
	IsEnabled     bool `json:"is_enabled"`      // Information whether addresses is enabled
	MainAddressID int  `json:"main_address_id"` // Main address id for group
}

// GroupsGroup struct
type GroupsGroup struct {
	AdminLevel   int    `json:"admin_level"`
	Deactivated  string `json:"deactivated"`   // Information whether community is banned
	FinishDate   int    `json:"finish_date"`   // Finish date in Unixtime format
	ID           int    `json:"id"`            // Community ID
	IsAdmin      int    `json:"is_admin"`      // Information whether current user is administrator
	IsAdvertiser int    `json:"is_advertiser"` // Information whether current user is advertiser
	IsClosed     int    `json:"is_closed"`
	IsMember     int    `json:"is_member"`   // Information whether current user is member
	Name         string `json:"name"`        // Community name
	Photo100     string `json:"photo_100"`   // URL of square photo of the community with 100 pixels in width
	Photo200     string `json:"photo_200"`   // URL of square photo of the community with 200 pixels in width
	Photo50      string `json:"photo_50"`    // URL of square photo of the community with 50 pixels in width
	ScreenName   string `json:"screen_name"` // Domain of the community page
	StartDate    int    `json:"start_date"`  // Start date in Unixtime format
	Type         string `json:"type"`

	Market               GroupsMarketInfo     `json:"market"`
	MemberStatus         int                  `json:"member_status"` // Current user's member status
	IsFavorite           int                  `json:"is_favorite"`   // Information whether community is in faves
	IsSubscribed         int                  `json:"is_subscribed"` // Information whether current user is subscribed
	City                 BaseObject           `json:"city"`
	Country              BaseCountry          `json:"country"`
	Verified             int                  `json:"verified"`      // Information whether community is verified
	Description          string               `json:"description"`   // Community description
	WikiPage             string               `json:"wiki_page"`     // Community's main wiki page title
	MembersCount         int                  `json:"members_count"` // Community members number
	Counters             groupsCountersGroup  `json:"counters"`
	Cover                groupsCover          `json:"cover"`
	CanPost              int                  `json:"can_post"`          // Information whether current user can post on community's wall
	CanSeeAllPosts       int                  `json:"can_see_all_posts"` // Information whether current user can see all posts on community's wall
	Activity             string               `json:"activity"`          // Type of group, start date of event or category of public page
	FixedPost            int                  `json:"fixed_post"`        // Fixed post ID
	CanCreateTopic       int                  `json:"can_create_topic"`  // Information whether current user can create topic
	CanUploadVideo       int                  `json:"can_upload_video"`  // Information whether current user can upload video
	HasPhoto             int                  `json:"has_photo"`         // Information whether community has photo
	Status               string               `json:"status"`            // Community status
	MainAlbumID          int                  `json:"main_album_id"`     // Community's main photo album ID
	Links                []GroupsLinksItem    `json:"links"`
	Contacts             []groupsContactsItem `json:"contacts"`
	Site                 string               `json:"site"` // Community's website
	MainSection          int                  `json:"main_section"`
	Trending             int                  `json:"trending"`               // Information whether the community has a fire pictogram.
	CanMessage           int                  `json:"can_message"`            // Information whether current user can send a message to community
	IsMessagesBlocked    int                  `json:"is_messages_blocked"`    // Information whether community can send a message to current user
	CanSendNotify        int                  `json:"can_send_notify"`        // Information whether community can send notifications by phone number to current user
	OnlineStatus         GroupsOnlineStatus   `json:"online_status"`          // Status of replies in community messages
	AgeLimits            int                  `json:"age_limits"`             // Information whether age limit
	BanInfo              groupsGroupBanInfo   `json:"ban_info"`               // User ban info
	Addresses            GroupsAddressesInfo  `json:"addresses"`              // Info about addresses in groups
	IsSubscribedPodcasts bool                 `json:"is_subscribed_podcasts"` // Information whether current user is subscribed to podcasts
	CanSubscribePodcasts bool                 `json:"can_subscribe_podcasts"` // Owner in whitelist or not
	CanSubscribePosts    bool                 `json:"can_subscribe_posts"`    // Can subscribe to wall
}

func (group GroupsGroup) ToMention() string {
	return fmt.Sprintf("[club%d|%s]", group.ID, group.Name)
}

type groupsBanInfo struct {
	AdminID        int    `json:"admin_id"` // Administrator ID
	Comment        string `json:"comment"`  // Comment for a ban
	Date           int    `json:"date"`     // Date when user has been added to blacklist in Unixtime
	EndDate        int    `json:"end_date"` // Date when user will be removed from blacklist in Unixtime
	Reason         int    `json:"reason"`
	CommentVisible bool   `json:"comment_visible"`
}

// GroupsCallbackServer struct
type GroupsCallbackServer struct {
	CreatorID int    `json:"creator_id"`
	ID        int    `json:"id"`
	SecretKey string `json:"secret_key"`
	Status    string `json:"status"`
	Title     string `json:"title"`
	URL       string `json:"url"`
}

// GroupsCallbackSettings struct
type GroupsCallbackSettings struct {
	APIVersion string               `json:"api_version"` // API version used for the events
	Events     GroupsLongPollEvents `json:"events"`
}

type groupsContactsItem struct {
	Desc   string `json:"desc"`    // Contact description
	Email  string `json:"email"`   // Contact email
	Phone  string `json:"phone"`   // Contact phone
	UserID int    `json:"user_id"` // User ID
}

type groupsCountersGroup struct {
	Addresses int `json:"addresses"` // Addresses number
	Albums    int `json:"albums"`    // Photo albums number
	Audios    int `json:"audios"`    // Audios number
	Docs      int `json:"docs"`      // Docs number
	Market    int `json:"market"`    // Market items number
	Photos    int `json:"photos"`    // Photos number
	Topics    int `json:"topics"`    // Topics number
	Videos    int `json:"videos"`    // Videos number
}

type groupsCover struct {
	Enabled int         `json:"enabled"` // Information whether cover is enabled
	Images  []baseImage `json:"images"`
}

type groupsGroupBanInfo struct {
	Comment string `json:"comment"`  // Ban comment
	EndDate int    `json:"end_date"` // End date of ban in Unixtime
}

// GroupsGroupCategory struct
type GroupsGroupCategory struct {
	ID            int                  `json:"id"`   // Category ID
	Name          string               `json:"name"` // Category name
	Subcategories []baseObjectWithName `json:"subcategories"`
}

// GroupsGroupCategoryFull struct
type GroupsGroupCategoryFull struct {
	ID            int                   `json:"id"`         // Category ID
	Name          string                `json:"name"`       // Category name
	PageCount     int                   `json:"page_count"` // Pages number
	PagePreviews  []GroupsGroup         `json:"page_previews"`
	Subcategories []GroupsGroupCategory `json:"subcategories"`
}

type groupsGroupCategoryType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GroupsGroupLink struct
type GroupsGroupLink struct {
	Desc            string `json:"desc"`       // Link description
	EditTitle       int    `json:"edit_title"` // Information whether the title can be edited
	Name            string `json:"name"`
	ID              int    `json:"id"`               // Link ID
	ImageProcessing int    `json:"image_processing"` // Information whether the image on processing
	URL             string `json:"url"`              // Link URL
}

type groupsGroupPublicCategoryList struct {
	ID           int                       `json:"id"`
	Name         string                    `json:"name"`
	SubtypesList []groupsGroupCategoryType `json:"subtypes_list"`
}

// GroupsGroupSettings struct
type GroupsGroupSettings struct {
	Access             int                             `json:"access"`            // Community access settings
	Address            string                          `json:"address"`           // Community's page domain
	Audio              int                             `json:"audio"`             // Audio settings
	Description        string                          `json:"description"`       // Community description
	Docs               int                             `json:"docs"`              // Docs settings
	ObsceneFilter      int                             `json:"obscene_filter"`    // Information whether the obscene filter is enabled
	ObsceneStopwords   int                             `json:"obscene_stopwords"` // Information whether the stopwords filter is enabled
	ObsceneWords       string                          `json:"obscene_words"`     // The list of stop words
	Photos             int                             `json:"photos"`            // Photos settings
	PublicCategory     int                             `json:"public_category"`   // Information about the group category
	PublicCategoryList []groupsGroupPublicCategoryList `json:"public_category_list"`
	PublicSubcategory  int                             `json:"public_subcategory"` // Information about the group subcategory
	Rss                string                          `json:"rss"`                // URL of the RSS feed
	Subject            int                             `json:"subject"`            // Community subject ID
	SubjectList        []groupsSubjectItem             `json:"subject_list"`
	Title              string                          `json:"title"`   // Community title
	Topics             int                             `json:"topics"`  // Topics settings
	Video              int                             `json:"video"`   // Video settings
	Wall               int                             `json:"wall"`    // Wall settings
	Website            string                          `json:"website"` // Community website
	Wiki               int                             `json:"wiki"`    // Wiki settings
	CountryID          int                             `json:"country_id"`
	CityID             int                             `json:"city_id"`
	Messages           int                             `json:"messages"`
	Articles           int                             `json:"articles"`
	Events             int                             `json:"events"`
	AgeLimits          int                             `json:"age_limits"`
	Market             struct {
		Enabled         int   `json:"enabled"`
		CommentsEnabled int   `json:"comments_enabled"`
		CountryIDs      []int `json:"country_ids"`
		ContactID       int   `json:"contact_id"`
		Currency        struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"currency"`
	} `json:"market"`
	SectionsList     [][]interface{} `json:"sections_list"`
	MainSection      int             `json:"main_section"`
	SecondarySection int             `json:"secondary_section"`
	ActionButton     struct {
		ActionType string        `json:"action_type"`
		Target     []interface{} `json:"target"`
		Title      string        `json:"title"`
	} `json:"action_button"`
	LiveCovers struct {
		IsEnabled bool `json:"is_enabled"`
	} `json:"live_covers"`
}

// GroupsGroupXtrInvitedBy struct
type GroupsGroupXtrInvitedBy struct {
	AdminLevel   int    `json:"admin_level"`
	ID           int    `json:"id"`          // Community ID
	InvitedBy    int    `json:"invited_by"`  // Inviter ID
	IsAdmin      int    `json:"is_admin"`    // Information whether current user is manager
	IsClosed     int    `json:"is_closed"`   // Information whether community is closed
	IsMember     int    `json:"is_member"`   // Information whether current user is member
	Name         string `json:"name"`        // Community name
	Photo100     string `json:"photo_100"`   // URL of square photo of the community with 100 pixels in width
	Photo200     string `json:"photo_200"`   // URL of square photo of the community with 200 pixels in width
	Photo50      string `json:"photo_50"`    // URL of square photo of the community with 50 pixels in width
	ScreenName   string `json:"screen_name"` // Domain of the community page
	Type         string `json:"type"`
	IsAdvertiser int    `json:"is_advertiser"` // Information whether current user is advertiser
}

func (group GroupsGroupXtrInvitedBy) ToMention() string {
	return fmt.Sprintf("[club%d|%s]", group.ID, group.Name)
}

// GroupsLinksItem struct
type GroupsLinksItem struct {
	Desc      string `json:"desc"`       // Link description
	EditTitle int    `json:"edit_title"` // Information whether the link title can be edited
	ID        int    `json:"id"`         // Link ID
	Name      string `json:"name"`       // Link title
	Photo100  string `json:"photo_100"`  // URL of square image of the link with 100 pixels in width
	Photo50   string `json:"photo_50"`   // URL of square image of the link with 50 pixels in width
	URL       string `json:"url"`        // Link URL
}

// GroupsLongPollEvents struct
type GroupsLongPollEvents struct {
	MessageNew           int `json:"message_new"`
	MessageReply         int `json:"message_reply"`
	PhotoNew             int `json:"photo_new"`
	AudioNew             int `json:"audio_new"`
	VideoNew             int `json:"video_new"`
	WallReplyNew         int `json:"wall_reply_new"`
	WallReplyEdit        int `json:"wall_reply_edit"`
	WallReplyDelete      int `json:"wall_reply_delete"`
	WallReplyRestore     int `json:"wall_reply_restore"`
	WallPostNew          int `json:"wall_post_new"`
	BoardPostNew         int `json:"board_post_new"`
	BoardPostEdit        int `json:"board_post_edit"`
	BoardPostRestore     int `json:"board_post_restore"`
	BoardPostDelete      int `json:"board_post_delete"`
	PhotoCommentNew      int `json:"photo_comment_new"`
	PhotoCommentEdit     int `json:"photo_comment_edit"`
	PhotoCommentDelete   int `json:"photo_comment_delete"`
	PhotoCommentRestore  int `json:"photo_comment_restore"`
	VideoCommentNew      int `json:"video_comment_new"`
	VideoCommentEdit     int `json:"video_comment_edit"`
	VideoCommentDelete   int `json:"video_comment_delete"`
	VideoCommentRestore  int `json:"video_comment_restore"`
	MarketCommentNew     int `json:"market_comment_new"`
	MarketCommentEdit    int `json:"market_comment_edit"`
	MarketCommentDelete  int `json:"market_comment_delete"`
	MarketCommentRestore int `json:"market_comment_restore"`
	PollVoteNew          int `json:"poll_vote_new"`
	GroupJoin            int `json:"group_join"`
	GroupLeave           int `json:"group_leave"`
	GroupChangeSettings  int `json:"group_change_settings"`
	GroupChangePhoto     int `json:"group_change_photo"`
	GroupOfficersEdit    int `json:"group_officers_edit"`
	MessageAllow         int `json:"message_allow"`
	MessageDeny          int `json:"message_deny"`
	WallRepost           int `json:"wall_repost"`
	UserBlock            int `json:"user_block"`
	UserUnblock          int `json:"user_unblock"`
	MessageEdit          int `json:"message_edit"`
	MessagesEdit         int `json:"messages_edit"` // BUG(VK): https://vk.com/bugtracker?act=show&id=86762
	MessageTypingState   int `json:"message_typing_state"`
	LeadFormsNew         int `json:"lead_forms_new"`
	LikeAdd              int `json:"like_add"`
	LikeRemove           int `json:"like_remove"`
	VkpayTransaction     int `json:"vkpay_transaction"`
	AppPayload           int `json:"app_payload"`
	MessageRead          int `json:"message_read"`
}

// GroupsLongPollServer struct
type GroupsLongPollServer struct {
	Key    string `json:"key"`    // Long Poll key
	Server string `json:"server"` // Long Poll server address
	Ts     string `json:"ts"`     // Number of the last event
}

// GroupsLongPollSettings struct
type GroupsLongPollSettings struct {
	APIVersion string               `json:"api_version"` // API version used for the events
	Events     GroupsLongPollEvents `json:"events"`
	IsEnabled  bool                 `json:"is_enabled"` // Shows whether Long Poll is enabled
}

// GroupsMarketInfo struct
type GroupsMarketInfo struct {
	ContactID    int            `json:"contact_id"` // Contact person ID
	Currency     marketCurrency `json:"currency"`
	CurrencyText string         `json:"currency_text"` // Currency name
	Enabled      int            `json:"enabled"`       // Information whether the market is enabled
	MainAlbumID  int            `json:"main_album_id"` // Main market album ID
	PriceMax     int            `json:"price_max"`     // Maximum price
	PriceMin     int            `json:"price_min"`     // Minimum price
}

// GroupsMemberRole struct
type GroupsMemberRole struct {
	ID   int    `json:"id"` // User ID
	Role string `json:"role"`
}

// GroupsMemberStatus struct
type GroupsMemberStatus struct {
	Member      int      `json:"member"`  // Information whether user is a member of the group
	UserID      int      `json:"user_id"` // User ID
	Permissions []string `json:"permissions"`
}

// GroupsMemberStatusFull struct
type GroupsMemberStatusFull struct {
	Invitation int `json:"invitation"` // Information whether user has been invited to the group
	Member     int `json:"member"`     // Information whether user is a member of the group
	Request    int `json:"request"`    // Information whether user has send request to the group
	CanInvite  int `json:"can_invite"` // Information whether user can be invite
	CanRecall  int `json:"can_recall"` // Information whether user's invite to the group can be recalled
	UserID     int `json:"user_id"`    // User ID
}

// GroupsOnlineStatus struct
type GroupsOnlineStatus struct {
	Minutes int    `json:"minutes"` // Estimated time of answer (for status = answer_mark)
	Status  string `json:"status"`
}

// GroupsOwnerXtrBanInfo struct
type GroupsOwnerXtrBanInfo struct {
	BanInfo groupsBanInfo `json:"ban_info"`
	Group   GroupsGroup   `json:"group"`
	Profile UsersUser     `json:"profile"`
	Type    string        `json:"type"`
}

type groupsSubjectItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type groupsTokenPermissionSetting struct {
	Name    string `json:"name"`
	Setting int    `json:"setting"`
}

// GroupsTokenPermissions struct
type GroupsTokenPermissions struct {
	Mask        int                            `json:"mask"`
	Permissions []groupsTokenPermissionSetting `json:"permissions"`
}
