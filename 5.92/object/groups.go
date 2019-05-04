package object // import "github.com/SevereCloud/vksdk/5.92/object"

// GroupsGroup struct
type GroupsGroup struct {
	AdminLevel   int    `json:"admin_level"`
	Deactivated  string `json:"deactivated"`
	ID           int    `json:"id"`
	IsAdmin      int    `json:"is_admin"`
	IsClosed     int    `json:"is_closed"`
	IsMember     int    `json:"is_member"`
	Name         string `json:"name"`
	Photo100     string `json:"photo_100"`
	Photo200     string `json:"photo_200"`
	Photo50      string `json:"photo_50"`
	ScreenName   string `json:"screen_name"`
	Type         string `json:"type"`
	IsAdvertiser int    `json:"is_advertiser"`
}

type groupsBanInfo struct {
	AdminID        int    `json:"admin_id"`
	Comment        string `json:"comment"`
	Date           int    `json:"date"`
	EndDate        int    `json:"end_date"`
	Reason         int    `json:"reason"`
	CommentVisible bool   `json:"comment_visible"`
}

// GroupsCallbackSettings struct
type GroupsCallbackSettings struct {
	APIVersion string               `json:"api_version"`
	Events     GroupsLongPollEvents `json:"events"`
}

type groupsContactsItem struct {
	Desc   string `json:"desc"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	UserID int    `json:"user_id"`
}

type groupsCountersGroup struct {
	Albums int `json:"albums"`
	Audios int `json:"audios"`
	Docs   int `json:"docs"`
	Market int `json:"market"`
	Photos int `json:"photos"`
	Topics int `json:"topics"`
	Videos int `json:"videos"`
}

type groupsCover struct {
	Enabled int         `json:"enabled"`
	Images  []baseImage `json:"images"`
}

type groupsGroupBanInfo struct {
	Comment string `json:"comment"`
	EndDate int    `json:"end_date"`
}

// GroupsGroupCategory struct
type GroupsGroupCategory struct {
	ID            int                  `json:"id"`
	Name          string               `json:"name"`
	Subcategories []baseObjectWithName `json:"subcategories"`
}

// GroupsGroupCategoryFull struct
type GroupsGroupCategoryFull struct {
	ID            int                   `json:"id"`
	Name          string                `json:"name"`
	PageCount     int                   `json:"page_count"`
	PagePreviews  []GroupsGroup         `json:"page_previews"`
	Subcategories []GroupsGroupCategory `json:"subcategories"`
}

type groupsGroupCategoryType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GroupsGroupLink struct
type GroupsGroupLink struct {
	Desc            string `json:"desc"`
	EditTitle       int    `json:"edit_title"`
	Name            string `json:"name"`
	ID              int    `json:"id"`
	ImageProcessing int    `json:"image_processing"`
	URL             string `json:"url"`
}

type groupsGroupPublicCategoryList struct {
	ID           int                       `json:"id"`
	Name         string                    `json:"name"`
	SubtypesList []groupsGroupCategoryType `json:"subtypes_list"`
}

// GroupsGroupSettings struct
type GroupsGroupSettings struct {
	Access             int                             `json:"access"`
	Address            string                          `json:"address"`
	Audio              int                             `json:"audio"`
	Description        string                          `json:"description"`
	Docs               int                             `json:"docs"`
	ObsceneFilter      int                             `json:"obscene_filter"`
	ObsceneStopwords   int                             `json:"obscene_stopwords"`
	ObsceneWords       string                          `json:"obscene_words"`
	Photos             int                             `json:"photos"`
	Place              placesPlaceMin                  `json:"place"`
	PublicCategory     int                             `json:"public_category"`
	PublicCategoryList []groupsGroupPublicCategoryList `json:"public_category_list"`
	PublicSubcategory  int                             `json:"public_subcategory"`
	Rss                string                          `json:"rss"`
	Subject            int                             `json:"subject"`
	SubjectList        []groupsSubjectItem             `json:"subject_list"`
	Title              string                          `json:"title"`
	Topics             int                             `json:"topics"`
	Video              int                             `json:"video"`
	Wall               int                             `json:"wall"`
	Website            string                          `json:"website"`
	Wiki               int                             `json:"wiki"`
	CountryID          int                             `json:"country_id"`
	CityID             int                             `json:"city_id"`
	Messages           int                             `json:"messages"`
	Articles           int                             `json:"articles"`
	Events             int                             `json:"events"`
	AgeLimits          int                             `json:"age_limits"`
	Market             struct {
		Enabled         int   `json:"enabled"`
		CommentsEnabled int   `json:"comments_enabled"`
		CountryIds      []int `json:"country_ids"`
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

type groupsGroupXtrInvitedBy struct {
	AdminLevel int    `json:"admin_level"`
	ID         string `json:"id"`
	InvitedBy  int    `json:"invited_by"`
	IsAdmin    int    `json:"is_admin"`
	IsClosed   int    `json:"is_closed"`
	IsMember   int    `json:"is_member"`
	Name       string `json:"name"`
	Photo100   string `json:"photo_100"`
	Photo200   string `json:"photo_200"`
	Photo50    string `json:"photo_50"`
	ScreenName string `json:"screen_name"`
	Type       string `json:"type"`
}

type groupsGroupsArray struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

type groupsLinksItem struct {
	Desc      string `json:"desc"`
	EditTitle int    `json:"edit_title"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Photo100  string `json:"photo_100"`
	Photo50   string `json:"photo_50"`
	URL       string `json:"url"`
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
	Key    string `json:"key"`
	Server string `json:"server"`
	Ts     string `json:"ts"`
}

// GroupsLongPollSettings struct
type GroupsLongPollSettings struct {
	APIVersion string               `json:"api_version"`
	Events     GroupsLongPollEvents `json:"events"`
	IsEnabled  bool                 `json:"is_enabled"`
}

type groupsMarketInfo struct {
	ContactID    int            `json:"contact_id"`
	Currency     marketCurrency `json:"currency"`
	CurrencyText string         `json:"currency_text"`
	Enabled      int            `json:"enabled"`
	MainAlbumID  int            `json:"main_album_id"`
	PriceMax     int            `json:"price_max"`
	PriceMin     int            `json:"price_min"`
}

type groupsMemberRole struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}

type groupsMemberStatus struct {
	Member int `json:"member"`
	UserID int `json:"user_id"`
}

type groupsMemberStatusFull struct {
	Invitation int `json:"invitation"`
	Member     int `json:"member"`
	Request    int `json:"request"`
	UserID     int `json:"user_id"`
}

// GroupsOnlineStatus struct
type GroupsOnlineStatus struct {
	Minutes int    `json:"minutes"`
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
