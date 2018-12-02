package vksdk

import "encoding/json"

// baseBoolInt base bool (integer, [0,1])
type baseBoolInt int

// baseObject base object
type baseObject struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// baseObjectCount base object count
type baseObjectCount struct {
	Count int `json:"count,omitempty"`
}

// baseObjectWithName base object with name
type baseObjectWithName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type baseRequestParam struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// User type
// FIXME add "role": "editor", "permissions": ["ads"]...
type User struct {
	ID          int        `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Bdate       string     `json:"bdate,omitempty"`
	PhotoID     string     `json:"photo_id,omitempty"`
	City        baseObject `json:"city,omitempty"`
	Online      int        `json:"online,omitempty"`
	Lists       []int      `json:"lists,omitempty"`
	Deactivated string     `json:"deactivated,omitempty"`
	Sex         int        `json:"sex,omitempty"`
}

type button []messagesKeyboardButton

type messagesKeyboard struct {
	Buttons []button `json:"buttons"`
	OneTime bool     `json:"one_time"`
}

type messagesKeyboardButton struct {
	Action messagesKeyboardButtonAction `json:"action"`
	Color  string                       `json:"color"`
}

type messagesKeyboardButtonAction struct {
	Label   string `json:"label,omitempty"`
	Payload string `json:"payload,omitempty"`
	Type    string `json:"type"`
}

// ------------------------------------------------------------//
// Далее сгенерированные объекты. После проверки переносить выше
// ------------------------------------------------------------//

// Callback struct
type Callback struct {
	Type    string          `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int             `json:"group_id"`
}

type accountAccountCounters struct {
	AppRequests        int `json:"app_requests,omitempty"`
	Events             int `json:"events,omitempty"`
	Friends            int `json:"friends,omitempty"`
	FriendsSuggestions int `json:"friends_suggestions,omitempty"`
	Gifts              int `json:"gifts,omitempty"`
	Groups             int `json:"groups,omitempty"`
	Messages           int `json:"messages,omitempty"`
	Notifications      int `json:"notifications,omitempty"`
	Photos             int `json:"photos,omitempty"`
	Videos             int `json:"videos,omitempty"`
}

type accountInfo struct {
	Country         string      `json:"country,omitempty"`
	HTTPSRequired   baseBoolInt `json:"https_required,omitempty"`
	Intro           baseBoolInt `json:"intro,omitempty"`
	Lang            int         `json:"lang,omitempty"`
	NoWallReplies   baseBoolInt `json:"no_wall_replies,omitempty"`
	OwnPostsDefault baseBoolInt `json:"own_posts_default,omitempty"`
	FaRequired      baseBoolInt `json:"2fa_required,omitempty"`
}

type accountLookupResult struct {
	Found []accountUserXtrContact `json:"found,omitempty"`
	Other []accountOtherContact   `json:"other,omitempty"`
}

type accountNameRequest struct {
	FirstName string                   `json:"first_name,omitempty"`
	ID        int                      `json:"id,omitempty"`
	LastName  string                   `json:"last_name,omitempty"`
	Status    accountNameRequestStatus `json:"status,omitempty"`
}

// Request status
type accountNameRequestStatus string

type accountOffer struct {
	Description      string `json:"description,omitempty"`
	ID               int    `json:"id,omitempty"`
	Img              string `json:"img,omitempty"`
	Instruction      string `json:"instruction,omitempty"`
	InstructionHTML  string `json:"instruction_html,omitempty"`
	Price            int    `json:"price,omitempty"`
	ShortDescription string `json:"short_description,omitempty"`
	Tag              string `json:"tag,omitempty"`
	Title            string `json:"title,omitempty"`
}

// Settings parameters
type accountOnoffOptions string

type accountOtherContact struct {
	CommonCount int    `json:"common_count,omitempty"`
	Contact     string `json:"contact,omitempty"`
}

type accountPushConversations struct {
	Count int                            `json:"count,omitempty"`
	Items []accountPushConversationsItem `json:"items,omitempty"`
}

type accountPushConversationsItem struct {
	DisabledUntil int         `json:"disabled_until,omitempty"`
	PeerID        int         `json:"peer_id,omitempty"`
	Sound         baseBoolInt `json:"sound,omitempty"`
}

type accountPushParams struct {
	AppRequest     []accountOnoffOptions       `json:"app_request,omitempty"`
	Birthday       []accountOnoffOptions       `json:"birthday,omitempty"`
	Chat           []accountPushParamsMode     `json:"chat,omitempty"`
	Comment        []accountPushParamsSettings `json:"comment,omitempty"`
	EventSoon      []accountOnoffOptions       `json:"event_soon,omitempty"`
	Friend         []accountOnoffOptions       `json:"friend,omitempty"`
	FriendAccepted []accountOnoffOptions       `json:"friend_accepted,omitempty"`
	FriendFound    []accountOnoffOptions       `json:"friend_found,omitempty"`
	GroupAccepted  []accountOnoffOptions       `json:"group_accepted,omitempty"`
	GroupInvite    []accountOnoffOptions       `json:"group_invite,omitempty"`
	Like           []accountPushParamsSettings `json:"like,omitempty"`
	Mention        []accountPushParamsSettings `json:"mention,omitempty"`
	Msg            []accountPushParamsMode     `json:"msg,omitempty"`
	NewPost        []accountOnoffOptions       `json:"new_post,omitempty"`
	PhotosTag      []accountPushParamsSettings `json:"photos_tag,omitempty"`
	Reply          []accountOnoffOptions       `json:"reply,omitempty"`
	Repost         []accountPushParamsSettings `json:"repost,omitempty"`
	SdkOpen        []accountOnoffOptions       `json:"sdk_open,omitempty"`
	WallPost       []accountOnoffOptions       `json:"wall_post,omitempty"`
	WallPublish    []accountOnoffOptions       `json:"wall_publish,omitempty"`
}

// Settings parameters
type accountPushParamsMode string

// Settings parameters
type accountPushParamsSettings string

type accountPushSettings struct {
	Conversations accountPushConversations `json:"conversations,omitempty"`
	Disabled      baseBoolInt              `json:"disabled,omitempty"`
	DisabledUntil int                      `json:"disabled_until,omitempty"`
	Settings      accountPushParams        `json:"settings,omitempty"`
}

type accountUserSettings struct {
	Bdate            string             `json:"bdate,omitempty"`
	BdateVisibility  int                `json:"bdate_visibility,omitempty"`
	City             baseObject         `json:"city,omitempty"`
	Country          baseCountry        `json:"country,omitempty"`
	FirstName        string             `json:"first_name,omitempty"`
	HomeTown         string             `json:"home_town,omitempty"`
	LastName         string             `json:"last_name,omitempty"`
	MaidenName       string             `json:"maiden_name,omitempty"`
	NameRequest      accountNameRequest `json:"name_request,omitempty"`
	Phone            string             `json:"phone,omitempty"`
	Relation         int                `json:"relation,omitempty"`
	RelationPartner  usersUserMin       `json:"relation_partner,omitempty"`
	RelationPending  int                `json:"relation_pending,omitempty"`
	RelationRequests []usersUserMin     `json:"relation_requests,omitempty"`
	ScreenName       string             `json:"screen_name,omitempty"`
	Sex              baseSex            `json:"sex,omitempty"`
	Status           string             `json:"status,omitempty"`
}

type accountUserXtrContact struct {
}

// Current user's role
type adsAccessRole string

type adsAccesses struct {
	ClientID string        `json:"client_id,omitempty"`
	Role     adsAccessRole `json:"role,omitempty"`
}

type adsAccount struct {
	AccessRole    adsAccessRole  `json:"access_role"`
	AccountID     int            `json:"account_id"`
	AccountStatus baseBoolInt    `json:"account_status"`
	AccountType   adsAccountType `json:"account_type"`
}

// Account type
type adsAccountType string

//type adsAd struct {
//	AdFormat              int `json:"ad_format"`
//	AdPlatform            string `json:"ad_platform,omitempty"`
//	AllLimit              int           `json:"all_limit"`
//	Approved              adsAdApproved `json:"approved"`
//	CampaignID            int           `json:"campaign_id"`
//	Category1ID           int           `json:"category1_id,omitempty"`
//	Category2ID           int           `json:"category2_id,omitempty"`
//	CostType              adsAdCostType `json:"cost_type"`
//	Cpc                   int           `json:"cpc,omitempty"`
//	Cpm                   int           `json:"cpm,omitempty"`
//	DisclaimerMedical     baseBoolInt   `json:"disclaimer_medical,omitempty"`
//	DisclaimerSpecialist  baseBoolInt   `json:"disclaimer_specialist,omitempty"`
//	DisclaimerSupplements baseBoolInt   `json:"disclaimer_supplements,omitempty"`
//	ID                    int           `json:"id"`
//	ImpressionsLimit      int           `json:"impressions_limit,omitempty"`
//	ImpressionsLimited    baseBoolInt   `json:"impressions_limited,omitempty"`
//	Name                  string        `json:"name"`
//	Status                adsAdStatus   `json:"status"`
//	Video                 baseBoolInt   `json:"video,omitempty"`
//}

// Review status
type adsAdApproved int

// Cost type
type adsAdCostType int

type adsAdLayout struct {
	AdFormat    int                 `json:"ad_format"`
	CampaignID  int                 `json:"campaign_id"`
	CostType    adsAdLayoutCostType `json:"cost_type"`
	Description string              `json:"description"`
	ID          int                 `json:"id"`
	ImageSrc    int                 `json:"image_src"`
	ImageSrc2x  int                 `json:"image_src_2x,omitempty"`
	LinkDomain  string              `json:"link_domain,omitempty"`
	LinkURL     string              `json:"link_url"`
	PreviewLink string              `json:"preview_link,omitempty"`
	Title       string              `json:"title"`
	Video       baseBoolInt         `json:"video,omitempty"`
}

// Cost type
type adsAdLayoutCostType int

// Ad atatus
type adsAdStatus int

type adsCampaign struct {
	AllLimit  string            `json:"all_limit"`
	DayLimit  string            `json:"day_limit"`
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	StartTime int               `json:"start_time"`
	Status    adsCampaignStatus `json:"status"`
	StopTime  int               `json:"stop_time"`
	Type      adsCampaignType   `json:"type"`
}

// Campaign status
type adsCampaignStatus int

// Campaign type
type adsCampaignType string

type adsCategory struct {
	ID            int                  `json:"id"`
	Name          string               `json:"name"`
	Subcategories []baseObjectWithName `json:"subcategories,omitempty"`
}

type adsClient struct {
	AllLimit string `json:"all_limit"`
	DayLimit string `json:"day_limit"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
}

type adsCriteria struct {
	AgeFrom              int                `json:"age_from,omitempty"`
	AgeTo                int                `json:"age_to,omitempty"`
	Apps                 string             `json:"apps,omitempty"`
	AppsNot              string             `json:"apps_not,omitempty"`
	Birthday             int                `json:"birthday,omitempty"`
	Cities               string             `json:"cities,omitempty"`
	CitiesNot            string             `json:"cities_not,omitempty"`
	Country              int                `json:"country,omitempty"`
	Districts            string             `json:"districts,omitempty"`
	Groups               string             `json:"groups,omitempty"`
	InterestCategories   string             `json:"interest_categories,omitempty"`
	Interests            string             `json:"interests,omitempty"`
	Paying               baseBoolInt        `json:"paying,omitempty"`
	Positions            string             `json:"positions,omitempty"`
	Religions            string             `json:"religions,omitempty"`
	RetargetingGroups    string             `json:"retargeting_groups,omitempty"`
	RetargetingGroupsNot string             `json:"retargeting_groups_not,omitempty"`
	SchoolFrom           int                `json:"school_from,omitempty"`
	SchoolTo             int                `json:"school_to,omitempty"`
	Schools              string             `json:"schools,omitempty"`
	Sex                  adsCriteriaSex     `json:"sex,omitempty"`
	Stations             string             `json:"stations,omitempty"`
	Statuses             string             `json:"statuses,omitempty"`
	Streets              string             `json:"streets,omitempty"`
	Travellers           basePropertyExists `json:"travellers,omitempty"`
	UniFrom              int                `json:"uni_from,omitempty"`
	UniTo                int                `json:"uni_to,omitempty"`
	UserBrowsers         string             `json:"user_browsers,omitempty"`
	UserDevices          string             `json:"user_devices,omitempty"`
	UserOs               string             `json:"user_os,omitempty"`
}

// Sex
type adsCriteriaSex int

type adsDemoStats struct {
	ID    int                `json:"id,omitempty"`
	Stats adsDemostatsFormat `json:"stats,omitempty"`
	Type  adsObjectType      `json:"type,omitempty"`
}

type adsDemostatsFormat struct {
	Age     []adsStatsAge    `json:"age,omitempty"`
	Cities  []adsStatsCities `json:"cities,omitempty"`
	Day     string           `json:"day,omitempty"`
	Month   string           `json:"month,omitempty"`
	Overall int              `json:"overall,omitempty"`
	Sex     []adsStatsSex    `json:"sex,omitempty"`
	SexAge  []adsStatsSexAge `json:"sex_age,omitempty"`
}

type adsFloodStats struct {
	Left    int `json:"left"`
	Refresh int `json:"refresh"`
}

type adsLinkStatus struct {
	Description string `json:"description"`
	RedirectURL string `json:"redirect_url"`
	Status      string `json:"status"`
}

// Object type
type adsObjectType string

type adsParagraphs struct {
	Paragraph string `json:"paragraph,omitempty"`
}

type adsPostStats struct {
}

type adsRejectReason struct {
	Comment string     `json:"comment,omitempty"`
	Rules   []adsRules `json:"rules,omitempty"`
}

type adsRules struct {
	Paragraphs []adsParagraphs `json:"paragraphs,omitempty"`
	Title      string          `json:"title,omitempty"`
}

type adsStats struct {
	ID    int            `json:"id,omitempty"`
	Stats adsStatsFormat `json:"stats,omitempty"`
	Type  adsObjectType  `json:"type,omitempty"`
}

type adsStatsAge struct {
	ClicksRate      float64 `json:"clicks_rate,omitempty"`
	ImpressionsRate float64 `json:"impressions_rate,omitempty"`
	Value           string  `json:"value,omitempty"`
}

type adsStatsCities struct {
	ClicksRate      float64 `json:"clicks_rate,omitempty"`
	ImpressionsRate float64 `json:"impressions_rate,omitempty"`
	Name            string  `json:"name,omitempty"`
	Value           int     `json:"value,omitempty"`
}

type adsStatsFormat struct {
	Clicks          int    `json:"clicks,omitempty"`
	Day             string `json:"day,omitempty"`
	Impressions     int    `json:"impressions,omitempty"`
	JoinRate        int    `json:"join_rate,omitempty"`
	Month           string `json:"month,omitempty"`
	Overall         int    `json:"overall,omitempty"`
	Reach           int    `json:"reach,omitempty"`
	Spent           int    `json:"spent,omitempty"`
	VideoClicksSite int    `json:"video_clicks_site,omitempty"`
	VideoViews      int    `json:"video_views,omitempty"`
	VideoViewsFull  int    `json:"video_views_full,omitempty"`
	VideoViewsHalf  int    `json:"video_views_half,omitempty"`
}

type adsStatsSex struct {
	ClicksRate      float64          `json:"clicks_rate,omitempty"`
	ImpressionsRate float64          `json:"impressions_rate,omitempty"`
	Value           adsStatsSexValue `json:"value,omitempty"`
}

type adsStatsSexAge struct {
	ClicksRate      float64 `json:"clicks_rate,omitempty"`
	ImpressionsRate float64 `json:"impressions_rate,omitempty"`
	Value           string  `json:"value,omitempty"`
}

// Sex
type adsStatsSexValue string

type adsTargSettings struct {
}

type adsTargStats struct {
	AudienceCount  int     `json:"audience_count"`
	RecommendedCpc float64 `json:"recommended_cpc,omitempty"`
	RecommendedCpm float64 `json:"recommended_cpm,omitempty"`
}

type adsTargSuggestions struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type adsTargSuggestionsCities struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Parent string `json:"parent,omitempty"`
}

type adsTargSuggestionsRegions struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type adsTargSuggestionsSchools struct {
	Desc   string                        `json:"desc,omitempty"`
	ID     int                           `json:"id,omitempty"`
	Name   string                        `json:"name,omitempty"`
	Parent string                        `json:"parent,omitempty"`
	Type   adsTargSuggestionsSchoolsType `json:"type,omitempty"`
}

// School type
type adsTargSuggestionsSchoolsType string

type adsTargetGroup struct {
	AudienceCount int    `json:"audience_count,omitempty"`
	Domain        string `json:"domain,omitempty"`
	ID            int    `json:"id,omitempty"`
	Lifetime      int    `json:"lifetime,omitempty"`
	Name          string `json:"name,omitempty"`
	Pixel         string `json:"pixel,omitempty"`
}

type adsUsers struct {
	Accesses []adsAccesses `json:"accesses"`
	UserID   int           `json:"user_id"`
}

type appsApp struct {
	AuthorGroup     int                    `json:"author_group,omitempty"`
	AuthorID        int                    `json:"author_id,omitempty"`
	AuthorURL       string                 `json:"author_url,omitempty"`
	Banner1120      string                 `json:"banner_1120,omitempty"`
	Banner560       string                 `json:"banner_560,omitempty"`
	CatalogPosition int                    `json:"catalog_position,omitempty"`
	Description     string                 `json:"description,omitempty"`
	Genre           string                 `json:"genre,omitempty"`
	GenreID         int                    `json:"genre_id,omitempty"`
	ID              int                    `json:"id"`
	Icon139         string                 `json:"icon_139,omitempty"`
	Icon150         string                 `json:"icon_150,omitempty"`
	Icon278         string                 `json:"icon_278,omitempty"`
	Icon75          string                 `json:"icon_75,omitempty"`
	International   int                    `json:"international,omitempty"`
	IsInCatalog     int                    `json:"is_in_catalog,omitempty"`
	LeaderboardType appsAppLeaderboardType `json:"leaderboard_type,omitempty"`
	MembersCount    int                    `json:"members_count,omitempty"`
	PlatformID      int                    `json:"platform_id,omitempty"`
	PublishedDate   int                    `json:"published_date,omitempty"`
	ScreenName      string                 `json:"screen_name,omitempty"`
	Screenshots     []photosPhoto          `json:"screenshots,omitempty"`
	Section         string                 `json:"section,omitempty"`
	Title           string                 `json:"title"`
	Type            appsAppType            `json:"type"`
}

// Leaderboard type
type appsAppLeaderboardType int

// Application type
type appsAppType string

type appsLeaderboard struct {
	Level  int `json:"level,omitempty"`
	Points int `json:"points,omitempty"`
	Score  int `json:"score,omitempty"`
	UserID int `json:"user_id"`
}

type audioAudio struct {
	AccessKey string `json:"access_key,omitempty"`
	Artist    string `json:"artist"`
	ID        int    `json:"id"`
	OwnerID   int    `json:"owner_id"`
	Title     string `json:"title"`
	URL       string `json:"url,omitempty"`
}

type audioAudioFull struct {
}

type audioAudioUploadResponse struct {
	Audio    string `json:"audio,omitempty"`
	Hash     string `json:"hash,omitempty"`
	Redirect string `json:"redirect,omitempty"`
	Server   int    `json:"server,omitempty"`
}

type audioLyrics struct {
	LyricsID int    `json:"lyrics_id"`
	Text     string `json:"text"`
}

type baseCity struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type baseCommentsInfo struct {
	CanPost       baseBoolInt `json:"can_post,omitempty"`
	Count         int         `json:"count,omitempty"`
	GroupsCanPost baseBoolInt `json:"groups_can_post,omitempty"`
}

type baseCountry struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type baseError struct {
	ErrorCode     int                `json:"error_code,omitempty"`
	ErrorMsg      string             `json:"error_msg,omitempty"`
	RequestParams []baseRequestParam `json:"request_params,omitempty"`
}

type baseGeo struct {
	Coordinates baseGeoCoordinates `json:"coordinates,omitempty"`
	Place       basePlace          `json:"place,omitempty"`
	Showmap     int                `json:"showmap,omitempty"`
	Type        string             `json:"type,omitempty"`
}

type baseGeoCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type baseImage struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type baseLikes struct {
	Count     int         `json:"count,omitempty"`
	UserLikes baseBoolInt `json:"user_likes,omitempty"`
}

type baseLikesInfo struct {
	CanLike    baseBoolInt `json:"can_like"`
	CanPublish baseBoolInt `json:"can_publish,omitempty"`
	Count      int         `json:"count"`
	UserLikes  int         `json:"user_likes"`
}

type baseLink struct {
	Application baseLinkApplication `json:"application,omitempty"`
	Button      baseLinkButton      `json:"button,omitempty"`
	Caption     string              `json:"caption,omitempty"`
	Description string              `json:"description,omitempty"`
	Photo       photosPhoto         `json:"photo,omitempty"`
	PreviewPage string              `json:"preview_page,omitempty"`
	PreviewURL  string              `json:"preview_url,omitempty"`
	Product     baseLinkProduct     `json:"product,omitempty"`
	Rating      baseLinkRating      `json:"rating,omitempty"`
	Title       string              `json:"title,omitempty"`
	URL         string              `json:"url"`
}

type baseLinkApplication struct {
	AppID float64                  `json:"app_id,omitempty"`
	Store baseLinkApplicationStore `json:"store,omitempty"`
}

type baseLinkApplicationStore struct {
	ID   float64 `json:"id,omitempty"`
	Name string  `json:"name,omitempty"`
}

type baseLinkButton struct {
	Action baseLinkButtonAction `json:"action,omitempty"`
	Title  string               `json:"title,omitempty"`
}

type baseLinkButtonAction struct {
	Type baseLinkButtonActionType `json:"type,omitempty"`
	URL  string                   `json:"url,omitempty"`
}

// Action type
type baseLinkButtonActionType string

type baseLinkProduct struct {
	Price marketPrice `json:"price,omitempty"`
}

type baseLinkRating struct {
	ReviewsCount int     `json:"reviews_count,omitempty"`
	Stars        float64 `json:"stars,omitempty"`
}

// Returns 1 if request has been processed successfully
type baseOkResponse int

type basePlace struct {
	Address   string  `json:"address,omitempty"`
	Checkins  int     `json:"checkins,omitempty"`
	City      string  `json:"city,omitempty"`
	Country   string  `json:"country,omitempty"`
	Created   int     `json:"created,omitempty"`
	ID        int     `json:"id,omitempty"`
	Icon      string  `json:"icon,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Title     string  `json:"title,omitempty"`
	Type      string  `json:"type,omitempty"`
}

type basePropertyExists int

type baseRepostsInfo struct {
	Count        int `json:"count,omitempty"`
	UserReposted int `json:"user_reposted,omitempty"`
}

type baseSex int

type baseSticker struct {
	Images               []baseImage `json:"images,omitempty"`
	ImagesWithBackground []baseImage `json:"images_with_background,omitempty"`
	ProductID            int         `json:"product_id,omitempty"`
	StickerID            int         `json:"sticker_id,omitempty"`
}

type baseUploadServer struct {
	UploadURL string `json:"upload_url"`
}

type baseUserID struct {
	UserID int `json:"user_id,omitempty"`
}

// Sort type
type boardDefaultOrder int

type boardTopic struct {
	Comments  int         `json:"comments,omitempty"`
	Created   int         `json:"created,omitempty"`
	CreatedBy int         `json:"created_by,omitempty"`
	ID        int         `json:"id,omitempty"`
	IsClosed  baseBoolInt `json:"is_closed,omitempty"`
	IsFixed   baseBoolInt `json:"is_fixed,omitempty"`
	Title     string      `json:"title,omitempty"`
	Updated   int         `json:"updated,omitempty"`
	UpdatedBy int         `json:"updated_by,omitempty"`
}

type boardTopicComment struct {
	Attachments []wallCommentAttachment `json:"attachments,omitempty"`
	Date        int                     `json:"date"`
	FromID      int                     `json:"from_id"`
	ID          int                     `json:"id"`
	RealOffset  int                     `json:"real_offset,omitempty"`
	Text        string                  `json:"text"`
}

type boardTopicPoll struct {
	AnswerID int           `json:"answer_id"`
	Answers  []pollsAnswer `json:"answers"`
	Created  int           `json:"created"`
	IsClosed baseBoolInt   `json:"is_closed,omitempty"`
	OwnerID  int           `json:"owner_id"`
	PollID   int           `json:"poll_id"`
	Question string        `json:"question"`
	Votes    string        `json:"votes"`
}

// User ID
type commonFriend int

type databaseCity struct {
}

type databaseFaculty struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type databaseRegion struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type databaseSchool struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type databaseUniversity struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type docsDoc struct {
	AccessKey string         `json:"access_key,omitempty"`
	Date      int            `json:"date"`
	Ext       string         `json:"ext"`
	ID        int            `json:"id"`
	OwnerID   int            `json:"owner_id"`
	Preview   docsDocPreview `json:"preview,omitempty"`
	Size      int            `json:"size"`
	Title     string         `json:"title"`
	Type      int            `json:"type"`
	URL       string         `json:"url,omitempty"`
}

type docsDocPreview struct {
	Photo docsDocPreviewPhoto `json:"photo,omitempty"`
	Video docsDocPreviewVideo `json:"video,omitempty"`
}

type docsDocPreviewPhoto struct {
	Sizes []photosPhotoSizes `json:"sizes,omitempty"`
}

type docsDocPreviewVideo struct {
	Filesize int    `json:"filesize"`
	Height   int    `json:"height"`
	Src      string `json:"src"`
	Width    int    `json:"width"`
}

type docsDocTypes struct {
	Count int    `json:"count,omitempty"`
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type docsDocUploadResponse struct {
	File string `json:"file,omitempty"`
}

type faveFavesLink struct {
	Description string `json:"description,omitempty"`
	ID          int    `json:"id,omitempty"`
	Photo100    string `json:"photo_100,omitempty"`
	Photo200    string `json:"photo_200,omitempty"`
	Photo50     string `json:"photo_50,omitempty"`
	Title       string `json:"title,omitempty"`
	URL         string `json:"url,omitempty"`
}

type friendsFriendStatus struct {
	FriendStatus   friendsFriendStatusStatus `json:"friend_status"`
	ReadState      baseBoolInt               `json:"read_state,omitempty"`
	RequestMessage string                    `json:"request_message,omitempty"`
	Sign           string                    `json:"sign,omitempty"`
	UserID         int                       `json:"user_id"`
}

// Friend status with the user
type friendsFriendStatusStatus int

type friendsFriendsList struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type friendsMutualFriend struct {
	CommonCount   int            `json:"common_count,omitempty"`
	CommonFriends []commonFriend `json:"common_friends,omitempty"`
	ID            int            `json:"id,omitempty"`
}

type friendsRequests struct {
	From   string                `json:"from,omitempty"`
	Mutual friendsRequestsMutual `json:"mutual,omitempty"`
	UserID int                   `json:"user_id,omitempty"`
}

type friendsRequestsMutual struct {
	Count int                         `json:"count,omitempty"`
	Users []friendsRequestsMutualUser `json:"users,omitempty"`
}

// User ID
type friendsRequestsMutualUser int

type friendsRequestsXtrMessage struct {
	From    string                `json:"from,omitempty"`
	Message string                `json:"message,omitempty"`
	Mutual  friendsRequestsMutual `json:"mutual,omitempty"`
	UserID  int                   `json:"user_id,omitempty"`
}

type friendsUserXtrLists struct {
}

type friendsUserXtrPhone struct {
}

type giftsGift struct {
	Date     int              `json:"date,omitempty"`
	FromID   int              `json:"from_id,omitempty"`
	Gift     giftsLayout      `json:"gift,omitempty"`
	GiftHash string           `json:"gift_hash,omitempty"`
	ID       int              `json:"id,omitempty"`
	Message  string           `json:"message,omitempty"`
	Privacy  giftsGiftPrivacy `json:"privacy,omitempty"`
}

// Gift privacy
type giftsGiftPrivacy int

type giftsLayout struct {
	ID       int    `json:"id,omitempty"`
	Thumb256 string `json:"thumb_256,omitempty"`
	Thumb48  string `json:"thumb_48,omitempty"`
	Thumb96  string `json:"thumb_96,omitempty"`
}

type groupsBanInfo struct {
	AdminID int                 `json:"admin_id,omitempty"`
	Comment string              `json:"comment,omitempty"`
	Date    int                 `json:"date,omitempty"`
	EndDate int                 `json:"end_date,omitempty"`
	Reason  groupsBanInfoReason `json:"reason,omitempty"`
}

// Ban reason
type groupsBanInfoReason int

type groupsCallbackSettings struct {
	APIVersion string               `json:"api_version,omitempty"`
	Events     groupsLongPollEvents `json:"events,omitempty"`
}

type groupsContactsItem struct {
	Desc   string `json:"desc,omitempty"`
	Email  string `json:"email,omitempty"`
	Phone  string `json:"phone,omitempty"`
	UserID int    `json:"user_id,omitempty"`
}

type groupsCountersGroup struct {
	Albums int `json:"albums,omitempty"`
	Audios int `json:"audios,omitempty"`
	Docs   int `json:"docs,omitempty"`
	Market int `json:"market,omitempty"`
	Photos int `json:"photos,omitempty"`
	Topics int `json:"topics,omitempty"`
	Videos int `json:"videos,omitempty"`
}

type groupsCover struct {
	Enabled baseBoolInt `json:"enabled"`
	Images  []baseImage `json:"images,omitempty"`
}

type groupsGroup struct {
	AdminLevel  groupsGroupAdminLevel `json:"admin_level,omitempty"`
	Deactivated string                `json:"deactivated,omitempty"`
	ID          int                   `json:"id,omitempty"`
	IsAdmin     baseBoolInt           `json:"is_admin,omitempty"`
	IsClosed    groupsGroupIsClosed   `json:"is_closed,omitempty"`
	IsMember    baseBoolInt           `json:"is_member,omitempty"`
	Name        string                `json:"name,omitempty"`
	Photo100    string                `json:"photo_100,omitempty"`
	Photo200    string                `json:"photo_200,omitempty"`
	Photo50     string                `json:"photo_50,omitempty"`
	ScreenName  string                `json:"screen_name,omitempty"`
	Type        groupsGroupType       `json:"type,omitempty"`
}

// Level of current user's credentials as manager
type groupsGroupAdminLevel int

type groupsGroupBanInfo struct {
	Comment string `json:"comment,omitempty"`
	EndDate int    `json:"end_date,omitempty"`
}

type groupsGroupCategory struct {
	ID            int                  `json:"id"`
	Name          string               `json:"name"`
	Subcategories []baseObjectWithName `json:"subcategories,omitempty"`
}

type groupsGroupCategoryFull struct {
	ID            int                   `json:"id"`
	Name          string                `json:"name"`
	PageCount     int                   `json:"page_count"`
	PagePreviews  []groupsGroup         `json:"page_previews"`
	Subcategories []groupsGroupCategory `json:"subcategories,omitempty"`
}

type groupsGroupCategoryType struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type groupsGroupFull struct {
}

type groupsGroupFullAgeLimits int

// Main section of community
type groupsGroupFullMainSection int

type groupsGroupFullMemberStatus int

// Information whether community is closed
type groupsGroupIsClosed int

type groupsGroupLink struct {
	Desc            string      `json:"desc,omitempty"`
	EditTitle       baseBoolInt `json:"edit_title,omitempty"`
	ID              int         `json:"id,omitempty"`
	ImageProcessing baseBoolInt `json:"image_processing,omitempty"`
	URL             string      `json:"url,omitempty"`
}

type groupsGroupPublicCategoryList struct {
	ID           int                       `json:"id,omitempty"`
	Name         string                    `json:"name,omitempty"`
	SubtypesList []groupsGroupCategoryType `json:"subtypes_list,omitempty"`
}

type groupsGroupSettings struct {
	Access             int                             `json:"access,omitempty"`
	Address            string                          `json:"address,omitempty"`
	Audio              int                             `json:"audio,omitempty"`
	Description        string                          `json:"description,omitempty"`
	Docs               int                             `json:"docs,omitempty"`
	ObsceneFilter      baseBoolInt                     `json:"obscene_filter,omitempty"`
	ObsceneStopwords   baseBoolInt                     `json:"obscene_stopwords,omitempty"`
	ObsceneWords       string                          `json:"obscene_words,omitempty"`
	Photos             int                             `json:"photos,omitempty"`
	Place              placesPlaceMin                  `json:"place,omitempty"`
	PublicCategory     int                             `json:"public_category,omitempty"`
	PublicCategoryList []groupsGroupPublicCategoryList `json:"public_category_list,omitempty"`
	PublicSubcategory  int                             `json:"public_subcategory,omitempty"`
	Rss                string                          `json:"rss,omitempty"`
	Subject            int                             `json:"subject,omitempty"`
	SubjectList        []groupsSubjectItem             `json:"subject_list,omitempty"`
	Title              string                          `json:"title,omitempty"`
	Topics             int                             `json:"topics,omitempty"`
	Video              int                             `json:"video,omitempty"`
	Wall               int                             `json:"wall,omitempty"`
	Website            string                          `json:"website,omitempty"`
	Wiki               int                             `json:"wiki,omitempty"`
}

// Community type
type groupsGroupType string

type groupsGroupXtrInvitedBy struct {
	AdminLevel groupsGroupXtrInvitedByAdminLevel `json:"admin_level,omitempty"`
	ID         string                            `json:"id,omitempty"`
	InvitedBy  int                               `json:"invited_by,omitempty"`
	IsAdmin    baseBoolInt                       `json:"is_admin,omitempty"`
	IsClosed   baseBoolInt                       `json:"is_closed,omitempty"`
	IsMember   baseBoolInt                       `json:"is_member,omitempty"`
	Name       string                            `json:"name,omitempty"`
	Photo100   string                            `json:"photo_100,omitempty"`
	Photo200   string                            `json:"photo_200,omitempty"`
	Photo50    string                            `json:"photo_50,omitempty"`
	ScreenName string                            `json:"screen_name,omitempty"`
	Type       groupsGroupXtrInvitedByType       `json:"type,omitempty"`
}

// Level of current user's credentials as manager
type groupsGroupXtrInvitedByAdminLevel int

// Community type
type groupsGroupXtrInvitedByType string

type groupsGroupsArray struct {
	Count int                     `json:"count"`
	Items []groupsGroupsArrayItem `json:"items"`
}

// Community ID
type groupsGroupsArrayItem int

type groupsLinksItem struct {
	Desc      string      `json:"desc,omitempty"`
	EditTitle baseBoolInt `json:"edit_title,omitempty"`
	ID        int         `json:"id,omitempty"`
	Name      string      `json:"name,omitempty"`
	Photo100  string      `json:"photo_100,omitempty"`
	Photo50   string      `json:"photo_50,omitempty"`
	URL       string      `json:"url,omitempty"`
}

type groupsLongPollEvents struct {
	AudioNew             baseBoolInt `json:"audio_new"`
	BoardPostDelete      baseBoolInt `json:"board_post_delete"`
	BoardPostEdit        baseBoolInt `json:"board_post_edit"`
	BoardPostNew         baseBoolInt `json:"board_post_new"`
	BoardPostRestore     baseBoolInt `json:"board_post_restore"`
	GroupChangePhoto     baseBoolInt `json:"group_change_photo"`
	GroupChangeSettings  baseBoolInt `json:"group_change_settings"`
	GroupJoin            baseBoolInt `json:"group_join"`
	GroupLeave           baseBoolInt `json:"group_leave"`
	GroupOfficersEdit    baseBoolInt `json:"group_officers_edit"`
	LeadFormsNew         baseBoolInt `json:"lead_forms_new,omitempty"`
	MarketCommentDelete  baseBoolInt `json:"market_comment_delete"`
	MarketCommentEdit    baseBoolInt `json:"market_comment_edit"`
	MarketCommentNew     baseBoolInt `json:"market_comment_new"`
	MarketCommentRestore baseBoolInt `json:"market_comment_restore"`
	MessageAllow         baseBoolInt `json:"message_allow"`
	MessageDeny          baseBoolInt `json:"message_deny"`
	MessageNew           baseBoolInt `json:"message_new"`
	MessageReply         baseBoolInt `json:"message_reply"`
	MessageTypingState   baseBoolInt `json:"message_typing_state"`
	MessagesEdit         baseBoolInt `json:"messages_edit"`
	PhotoCommentDelete   baseBoolInt `json:"photo_comment_delete"`
	PhotoCommentEdit     baseBoolInt `json:"photo_comment_edit"`
	PhotoCommentNew      baseBoolInt `json:"photo_comment_new"`
	PhotoCommentRestore  baseBoolInt `json:"photo_comment_restore"`
	PhotoNew             baseBoolInt `json:"photo_new"`
	PollVoteNew          baseBoolInt `json:"poll_vote_new"`
	UserBlock            baseBoolInt `json:"user_block"`
	UserUnblock          baseBoolInt `json:"user_unblock"`
	VideoCommentDelete   baseBoolInt `json:"video_comment_delete"`
	VideoCommentEdit     baseBoolInt `json:"video_comment_edit"`
	VideoCommentNew      baseBoolInt `json:"video_comment_new"`
	VideoCommentRestore  baseBoolInt `json:"video_comment_restore"`
	VideoNew             baseBoolInt `json:"video_new"`
	WallPostNew          baseBoolInt `json:"wall_post_new"`
	WallReplyDelete      baseBoolInt `json:"wall_reply_delete"`
	WallReplyEdit        baseBoolInt `json:"wall_reply_edit"`
	WallReplyNew         baseBoolInt `json:"wall_reply_new"`
	WallReplyRestore     baseBoolInt `json:"wall_reply_restore"`
	WallRepost           baseBoolInt `json:"wall_repost"`
}

type groupsLongPollServer struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	Ts     int    `json:"ts"`
}

type groupsLongPollSettings struct {
	APIVersion string               `json:"api_version,omitempty"`
	Events     groupsLongPollEvents `json:"events"`
	IsEnabled  bool                 `json:"is_enabled"`
}

type groupsMarketInfo struct {
	ContactID    int            `json:"contact_id,omitempty"`
	Currency     marketCurrency `json:"currency,omitempty"`
	CurrencyText string         `json:"currency_text,omitempty"`
	Enabled      baseBoolInt    `json:"enabled,omitempty"`
	MainAlbumID  int            `json:"main_album_id,omitempty"`
	PriceMax     int            `json:"price_max,omitempty"`
	PriceMin     int            `json:"price_min,omitempty"`
}

type groupsMemberRole struct {
	ID   int                    `json:"id,omitempty"`
	Role groupsMemberRoleStatus `json:"role,omitempty"`
}

// User's credentials as community admin
type groupsMemberRoleStatus string

type groupsMemberStatus struct {
	Member baseBoolInt `json:"member"`
	UserID int         `json:"user_id"`
}

type groupsMemberStatusFull struct {
	Invitation baseBoolInt `json:"invitation,omitempty"`
	Member     baseBoolInt `json:"member"`
	Request    baseBoolInt `json:"request,omitempty"`
	UserID     int         `json:"user_id"`
}

// Online status of group
type groupsOnlineStatus struct {
	Minutes int                    `json:"minutes,omitempty"`
	Status  groupsOnlineStatusType `json:"status"`
}

// Type of online status of group
type groupsOnlineStatusType string

type groupsOwnerXtrBanInfo struct {
	BanInfo groupsBanInfo             `json:"ban_info,omitempty"`
	Group   groupsGroup               `json:"group,omitempty"`
	Profile usersUser                 `json:"profile,omitempty"`
	Type    groupsOwnerXtrBanInfoType `json:"type,omitempty"`
}

// Owner type
type groupsOwnerXtrBanInfoType string

// User's credentials as community admin
type groupsRoleOptions string

type groupsSubjectItem struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type groupsTokenPermissionSetting struct {
	Name    string `json:"name"`
	Setting int    `json:"setting"`
}

type groupsTokenPermissions struct {
	Mask        int                            `json:"mask"`
	Permissions []groupsTokenPermissionSetting `json:"permissions,omitempty"`
}

type groupsUserXtrRole struct {
}

// User's languages
type lang string

type leadsChecked struct {
	Reason    string             `json:"reason,omitempty"`
	Result    leadsCheckedResult `json:"result,omitempty"`
	Sid       string             `json:"sid,omitempty"`
	StartLink string             `json:"start_link,omitempty"`
}

// Information whether user can start the lead
type leadsCheckedResult string

type leadsComplete struct {
	Cost     int            `json:"cost,omitempty"`
	Limit    int            `json:"limit,omitempty"`
	Spent    int            `json:"spent,omitempty"`
	Success  baseOkResponse `json:"success,omitempty"`
	TestMode baseBoolInt    `json:"test_mode,omitempty"`
}

type leadsEntry struct {
	Aid       int         `json:"aid,omitempty"`
	Comment   string      `json:"comment,omitempty"`
	Date      int         `json:"date,omitempty"`
	Sid       string      `json:"sid,omitempty"`
	StartDate int         `json:"start_date,omitempty"`
	Status    int         `json:"status,omitempty"`
	TestMode  baseBoolInt `json:"test_mode,omitempty"`
	UID       int         `json:"uid,omitempty"`
}

type leadsLead struct {
	Completed   int           `json:"completed,omitempty"`
	Cost        int           `json:"cost,omitempty"`
	Days        leadsLeadDays `json:"days,omitempty"`
	Impressions int           `json:"impressions,omitempty"`
	Limit       int           `json:"limit,omitempty"`
	Spent       int           `json:"spent,omitempty"`
	Started     int           `json:"started,omitempty"`
}

type leadsLeadDays struct {
	Completed   int `json:"completed,omitempty"`
	Impressions int `json:"impressions,omitempty"`
	Spent       int `json:"spent,omitempty"`
	Started     int `json:"started,omitempty"`
}

type leadsStart struct {
	TestMode baseBoolInt `json:"test_mode,omitempty"`
	VkSid    string      `json:"vk_sid,omitempty"`
}

type marketCurrency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type marketMarketAlbum struct {
	Count       int         `json:"count"`
	ID          int         `json:"id"`
	OwnerID     int         `json:"owner_id"`
	Photo       photosPhoto `json:"photo,omitempty"`
	Title       string      `json:"title"`
	UpdatedTime int         `json:"updated_time"`
}

type marketMarketCategory struct {
	ID      int           `json:"id"`
	Name    string        `json:"name"`
	Section marketSection `json:"section"`
}

type marketMarketItem struct {
	Availability marketMarketItemAvailability `json:"availability"`
	Category     marketMarketCategory         `json:"category"`
	Date         int                          `json:"date"`
	Description  string                       `json:"description"`
	ID           int                          `json:"id"`
	OwnerID      int                          `json:"owner_id"`
	Price        marketPrice                  `json:"price"`
	ThumbPhoto   string                       `json:"thumb_photo"`
	Title        string                       `json:"title"`
}

// Information whether the item is available
type marketMarketItemAvailability int

type marketMarketItemFull struct {
}

type marketPrice struct {
	Amount   string         `json:"amount,omitempty"`
	Currency marketCurrency `json:"currency,omitempty"`
	Text     string         `json:"text,omitempty"`
}

type marketSection struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type messagesChat struct {
	AdminID      int                      `json:"admin_id"`
	ID           int                      `json:"id"`
	Kicked       baseBoolInt              `json:"kicked,omitempty"`
	Left         baseBoolInt              `json:"left,omitempty"`
	Photo100     string                   `json:"photo_100,omitempty"`
	Photo200     string                   `json:"photo_200,omitempty"`
	Photo50      string                   `json:"photo_50,omitempty"`
	PushSettings messagesChatPushSettings `json:"push_settings,omitempty"`
	Title        string                   `json:"title,omitempty"`
	Type         string                   `json:"type"`
	Users        []messagesChatUser       `json:"users"`
}

type messagesChatFull struct {
	AdminID      int                        `json:"admin_id"`
	ID           int                        `json:"id"`
	Kicked       baseBoolInt                `json:"kicked,omitempty"`
	Left         baseBoolInt                `json:"left,omitempty"`
	Photo100     string                     `json:"photo_100,omitempty"`
	Photo200     string                     `json:"photo_200,omitempty"`
	Photo50      string                     `json:"photo_50,omitempty"`
	PushSettings messagesChatPushSettings   `json:"push_settings,omitempty"`
	Title        string                     `json:"title,omitempty"`
	Type         string                     `json:"type"`
	Users        []messagesUserXtrInvitedBy `json:"users"`
}

type messagesChatPushSettings struct {
	DisabledUntil int         `json:"disabled_until,omitempty"`
	Sound         baseBoolInt `json:"sound,omitempty"`
}

type messagesChatSettingsPhoto struct {
	Photo100 string `json:"photo_100,omitempty"`
	Photo200 string `json:"photo_200,omitempty"`
	Photo50  string `json:"photo_50,omitempty"`
}

type messagesChatSettingsState string

// User ID
type messagesChatUser int

type messagesConversation struct {
	CanWrite     messagesConversationCanWrite     `json:"can_write,omitempty"`
	ChatSettings messagesConversationChatSettings `json:"chat_settings,omitempty"`
	Important    bool                             `json:"important,omitempty"`
	InRead       int                              `json:"in_read,omitempty"`
	OutRead      int                              `json:"out_read,omitempty"`
	Peer         messagesConversationPeer         `json:"peer,omitempty"`
	PushSettings messagesConversationPushSettings `json:"push_settings,omitempty"`
	Unanswered   bool                             `json:"unanswered,omitempty"`
	UnreadCount  int                              `json:"unread_count,omitempty"`
}

type messagesConversationCanWrite struct {
	Allowed bool `json:"allowed,omitempty"`
	Reason  int  `json:"reason,omitempty"`
}

type messagesConversationChatSettings struct {
	MembersCount  int                       `json:"members_count,omitempty"`
	Photo         messagesChatSettingsPhoto `json:"photo,omitempty"`
	PinnedMessage messagesPinnedMessage     `json:"pinned_message,omitempty"`
	State         messagesChatSettingsState `json:"state,omitempty"`
	Title         string                    `json:"title,omitempty"`
}

type messagesConversationPeer struct {
	ID      int    `json:"id,omitempty"`
	LocalID int    `json:"local_id,omitempty"`
	Type    string `json:"type,omitempty"`
}

type messagesConversationPushSettings struct {
	DisabledForever bool `json:"disabled_forever,omitempty"`
	DisabledUntil   int  `json:"disabled_until,omitempty"`
	NoSound         bool `json:"no_sound,omitempty"`
}

type messagesConversationWithMessage struct {
	Conversation messagesConversation `json:"conversation,omitempty"`
	LastMessage  messagesMessage      `json:"last_message,omitempty"`
}

type messagesDialog struct {
	Important  baseBoolInt     `json:"important,omitempty"`
	InRead     int             `json:"in_read,omitempty"`
	Message    messagesMessage `json:"message,omitempty"`
	OutRead    int             `json:"out_read,omitempty"`
	Unanswered baseBoolInt     `json:"unanswered,omitempty"`
	Unread     int             `json:"unread,omitempty"`
}

type messagesHistoryAttachment struct {
	Attachment messagesHistoryMessageAttachment `json:"attachment"`
	MessageID  int                              `json:"message_id,omitempty"`
}

type messagesHistoryMessageAttachment struct {
	Audio  audioAudioFull                       `json:"audio,omitempty"`
	Doc    docsDoc                              `json:"doc,omitempty"`
	Link   baseLink                             `json:"link,omitempty"`
	Market baseLink                             `json:"market,omitempty"`
	Photo  photosPhoto                          `json:"photo,omitempty"`
	Share  baseLink                             `json:"share,omitempty"`
	Type   messagesHistoryMessageAttachmentType `json:"type"`
	Video  videoVideo                           `json:"video,omitempty"`
	Wall   baseLink                             `json:"wall,omitempty"`
}

// Attachments type
type messagesHistoryMessageAttachmentType string

type messagesLastActivity struct {
	Online baseBoolInt `json:"online"`
	Time   int         `json:"time"`
}

type messagesLongpollMessages struct {
	Count int               `json:"count,omitempty"`
	Items []messagesMessage `json:"items,omitempty"`
}

type messagesLongpollParams struct {
	Key    string `json:"key,omitempty"`
	Pts    int    `json:"pts,omitempty"`
	Server string `json:"server,omitempty"`
	Ts     int    `json:"ts,omitempty"`
}

type messagesMessage struct {
	Action                messagesMessageAction       `json:"action,omitempty"`
	Attachments           []messagesMessageAttachment `json:"attachments,omitempty"`
	ConversationMessageID int                         `json:"conversation_message_id,omitempty"`
	Date                  int                         `json:"date"`
	FromID                int                         `json:"from_id"`
	FwdMessages           []messagesMessage           `json:"fwd_messages,omitempty"`
	Geo                   baseGeo                     `json:"geo,omitempty"`
	ID                    int                         `json:"id"`
	Important             bool                        `json:"important,omitempty"`
	Keyboard              messagesKeyboard            `json:"keyboard,omitempty"`
	Payload               string                      `json:"payload,omitempty"`
	PeerID                int                         `json:"peer_id"`
	RandomID              int                         `json:"random_id,omitempty"`
	Text                  string                      `json:"text"`
	UpdateTime            int                         `json:"update_time,omitempty"`
}

type messagesMessageAction struct {
	ConversationMessageID int                         `json:"conversation_message_id,omitempty"`
	Email                 string                      `json:"email,omitempty"`
	MemberID              int                         `json:"member_id,omitempty"`
	Message               string                      `json:"message,omitempty"`
	Photo                 messagesChatSettingsPhoto   `json:"photo,omitempty"`
	Text                  string                      `json:"text,omitempty"`
	Type                  messagesMessageActionStatus `json:"type"`
}

// Action status
type messagesMessageActionStatus string

type messagesMessageAttachment struct {
	Audio             audioAudioFull                `json:"audio,omitempty"`
	Doc               docsDoc                       `json:"doc,omitempty"`
	Gift              giftsLayout                   `json:"gift,omitempty"`
	Link              baseLink                      `json:"link,omitempty"`
	Market            marketMarketItem              `json:"market,omitempty"`
	MarketMarketAlbum marketMarketAlbum             `json:"market_market_album,omitempty"`
	Photo             photosPhoto                   `json:"photo,omitempty"`
	Sticker           baseSticker                   `json:"sticker,omitempty"`
	Type              messagesMessageAttachmentType `json:"type"`
	Video             videoVideo                    `json:"video,omitempty"`
	Wall              wallWallpostAttached          `json:"wall,omitempty"`
	WallReply         wallWallComment               `json:"wall_reply,omitempty"`
}

// Attachment type
type messagesMessageAttachmentType string

type messagesPinnedMessage struct {
	Attachments           []messagesMessageAttachment `json:"attachments,omitempty"`
	ConversationMessageID int                         `json:"conversation_message_id,omitempty"`
	Date                  int                         `json:"date"`
	FromID                int                         `json:"from_id"`
	FwdMessages           []messagesMessage           `json:"fwd_messages,omitempty"`
	Geo                   baseGeo                     `json:"geo,omitempty"`
	ID                    int                         `json:"id"`
	PeerID                int                         `json:"peer_id"`
	Text                  string                      `json:"text"`
}

type messagesUserXtrInvitedBy struct {
}

type newsfeedItemAudio struct {
	Audio  newsfeedItemAudioAudio `json:"audio,omitempty"`
	PostID int                    `json:"post_id,omitempty"`
}

type newsfeedItemAudioAudio struct {
	Count int              `json:"count,omitempty"`
	Items []audioAudioFull `json:"items,omitempty"`
}

type newsfeedItemFriend struct {
	Friends newsfeedItemFriendFriends `json:"friends,omitempty"`
}

type newsfeedItemFriendFriends struct {
	Count int          `json:"count,omitempty"`
	Items []baseUserID `json:"items,omitempty"`
}

type newsfeedItemNote struct {
	Notes newsfeedItemNoteNotes `json:"notes,omitempty"`
}

type newsfeedItemNoteNotes struct {
	Count int                    `json:"count,omitempty"`
	Items []newsfeedNewsfeedNote `json:"items,omitempty"`
}

type newsfeedItemPhoto struct {
	Photos newsfeedItemPhotoPhotos `json:"photos,omitempty"`
	PostID int                     `json:"post_id,omitempty"`
}

type newsfeedItemPhotoPhotos struct {
	Count int                     `json:"count,omitempty"`
	Items []newsfeedNewsfeedPhoto `json:"items,omitempty"`
}

type newsfeedItemPhotoTag struct {
	PhotoTags newsfeedItemPhotoTagPhotoTags `json:"photo_tags,omitempty"`
	PostID    int                           `json:"post_id,omitempty"`
}

type newsfeedItemPhotoTagPhotoTags struct {
	Count int                     `json:"count,omitempty"`
	Items []newsfeedNewsfeedPhoto `json:"items,omitempty"`
}

type newsfeedItemTopic struct {
	Comments baseCommentsInfo `json:"comments,omitempty"`
	Likes    baseLikesInfo    `json:"likes,omitempty"`
	PostID   int              `json:"post_id"`
	Text     string           `json:"text"`
}

type newsfeedItemVideo struct {
	Video newsfeedItemVideoVideo `json:"video,omitempty"`
}

type newsfeedItemVideoVideo struct {
	Count int          `json:"count,omitempty"`
	Items []videoVideo `json:"items,omitempty"`
}

type newsfeedItemWallpost struct {
	Attachments []wallWallpostAttachment `json:"attachments,omitempty"`
	Comments    baseCommentsInfo         `json:"comments,omitempty"`
	CopyHistory []wallWallpost           `json:"copy_history,omitempty"`
	Geo         baseGeo                  `json:"geo,omitempty"`
	Likes       baseLikesInfo            `json:"likes,omitempty"`
	PostID      int                      `json:"post_id,omitempty"`
	PostSource  wallPostSource           `json:"post_source,omitempty"`
	PostType    newsfeedItemWallpostType `json:"post_type,omitempty"`
	Reposts     baseRepostsInfo          `json:"reposts,omitempty"`
	Text        string                   `json:"text,omitempty"`
}

// Post type
type newsfeedItemWallpostType string

type newsfeedList struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type newsfeedListFull struct {
}

type newsfeedNewsfeedItem struct {
}

// Item type
type newsfeedNewsfeedItemType string

type newsfeedNewsfeedNote struct {
	Comments int    `json:"comments,omitempty"`
	ID       int    `json:"id,omitempty"`
	OwnerID  int    `json:"owner_id,omitempty"`
	Title    string `json:"title,omitempty"`
}

type newsfeedNewsfeedPhoto struct {
}

type notesNote struct {
	CanComment baseBoolInt `json:"can_comment,omitempty"`
	Comments   int         `json:"comments"`
	Date       int         `json:"date"`
	ID         int         `json:"id"`
	OwnerID    int         `json:"owner_id"`
	Text       string      `json:"text,omitempty"`
	TextWiki   string      `json:"text_wiki,omitempty"`
	Title      string      `json:"title"`
	ViewURL    string      `json:"view_url"`
}

type notesNoteComment struct {
	Date    int    `json:"date"`
	ID      int    `json:"id"`
	Message string `json:"message"`
	Nid     int    `json:"nid"`
	Oid     int    `json:"oid"`
	ReplyTo int    `json:"reply_to,omitempty"`
	UID     int    `json:"uid"`
}

type notificationsFeedback struct {
	Attachments []wallWallpostAttachment `json:"attachments,omitempty"`
	FromID      int                      `json:"from_id,omitempty"`
	Geo         baseGeo                  `json:"geo,omitempty"`
	ID          int                      `json:"id,omitempty"`
	Likes       baseLikesInfo            `json:"likes,omitempty"`
	Text        string                   `json:"text,omitempty"`
	ToID        int                      `json:"to_id,omitempty"`
}

type notificationsNotification struct {
	Date     int                             `json:"date,omitempty"`
	Feedback notificationsFeedback           `json:"feedback,omitempty"`
	Parent   notificationsNotificationParent `json:"parent,omitempty"`
	Reply    notificationsReply              `json:"reply,omitempty"`
	Type     string                          `json:"type,omitempty"`
}

type notificationsNotificationParent struct {
}

type notificationsNotificationsComment struct {
	Date    int          `json:"date,omitempty"`
	ID      int          `json:"id,omitempty"`
	OwnerID int          `json:"owner_id,omitempty"`
	Photo   photosPhoto  `json:"photo,omitempty"`
	Post    wallWallpost `json:"post,omitempty"`
	Text    string       `json:"text,omitempty"`
	Topic   boardTopic   `json:"topic,omitempty"`
	Video   videoVideo   `json:"video,omitempty"`
}

type notificationsReply struct {
	Date int `json:"date,omitempty"`
	ID   int `json:"id,omitempty"`
	Text int `json:"text,omitempty"`
}

type oauthError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	RedirectURI      string `json:"redirect_uri,omitempty"`
}

type objects interface{}

type ordersAmount struct {
	Amounts  []ordersAmountItem `json:"amounts,omitempty"`
	Currency string             `json:"currency,omitempty"`
}

type ordersAmountItem struct {
	Amount      int    `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
	Votes       string `json:"votes,omitempty"`
}

type ordersOrder struct {
	Amount              int    `json:"amount,omitempty"`
	AppOrderID          int    `json:"app_order_id,omitempty"`
	CancelTransactionID int    `json:"cancel_transaction_id,omitempty"`
	Date                int    `json:"date,omitempty"`
	ID                  int    `json:"id,omitempty"`
	Item                string `json:"item,omitempty"`
	ReceiverID          int    `json:"receiver_id,omitempty"`
	Status              string `json:"status,omitempty"`
	TransactionID       int    `json:"transaction_id,omitempty"`
	UserID              int    `json:"user_id,omitempty"`
}

type pagesPrivacySettings int

type pagesWikipage struct {
	CreatorID   int                  `json:"creator_id,omitempty"`
	CreatorName int                  `json:"creator_name,omitempty"`
	EditorID    int                  `json:"editor_id,omitempty"`
	EditorName  string               `json:"editor_name,omitempty"`
	GroupID     int                  `json:"group_id"`
	ID          int                  `json:"id"`
	Title       string               `json:"title"`
	Views       int                  `json:"views"`
	WhoCanEdit  pagesPrivacySettings `json:"who_can_edit"`
	WhoCanView  pagesPrivacySettings `json:"who_can_view"`
}

type pagesWikipageFull struct {
	Created                  int                  `json:"created"`
	CreatorID                int                  `json:"creator_id,omitempty"`
	CurrentUserCanEdit       baseBoolInt          `json:"current_user_can_edit,omitempty"`
	CurrentUserCanEditAccess baseBoolInt          `json:"current_user_can_edit_access,omitempty"`
	Edited                   int                  `json:"edited"`
	EditorID                 int                  `json:"editor_id,omitempty"`
	GroupID                  int                  `json:"group_id"`
	HTML                     string               `json:"html,omitempty"`
	ID                       int                  `json:"id"`
	Source                   string               `json:"source,omitempty"`
	Title                    string               `json:"title"`
	ViewURL                  string               `json:"view_url"`
	Views                    int                  `json:"views"`
	WhoCanEdit               pagesPrivacySettings `json:"who_can_edit"`
	WhoCanView               pagesPrivacySettings `json:"who_can_view"`
}

type pagesWikipageVersion struct {
	Edited     int    `json:"edited,omitempty"`
	EditorID   int    `json:"editor_id,omitempty"`
	EditorName string `json:"editor_name,omitempty"`
	ID         int    `json:"id,omitempty"`
	Length     int    `json:"length,omitempty"`
}

type photosCommentXtrPid struct {
	Attachments    []wallCommentAttachment `json:"attachments,omitempty"`
	Date           int                     `json:"date"`
	FromID         int                     `json:"from_id"`
	ID             int                     `json:"id"`
	Likes          baseLikesInfo           `json:"likes,omitempty"`
	Pid            int                     `json:"pid"`
	ReplyToComment int                     `json:"reply_to_comment,omitempty"`
	ReplyToUser    int                     `json:"reply_to_user,omitempty"`
	Text           string                  `json:"text"`
}

type photosImage struct {
	Height int             `json:"height,omitempty"`
	Type   photosImageType `json:"type,omitempty"`
	URL    string          `json:"url,omitempty"`
	Width  int             `json:"width,omitempty"`
}

// Photo's type.
type photosImageType string

// String ID of photo
type photosListItem string

type photosMarketAlbumUploadResponse struct {
	Gid    int    `json:"gid,omitempty"`
	Hash   string `json:"hash,omitempty"`
	Photo  string `json:"photo,omitempty"`
	Server int    `json:"server,omitempty"`
}

type photosMarketUploadResponse struct {
	CropData string `json:"crop_data,omitempty"`
	CropHash string `json:"crop_hash,omitempty"`
	GroupID  int    `json:"group_id,omitempty"`
	Hash     string `json:"hash,omitempty"`
	Photo    string `json:"photo,omitempty"`
	Server   int    `json:"server,omitempty"`
}

type photosMessageUploadResponse struct {
	Hash   string `json:"hash,omitempty"`
	Photo  string `json:"photo,omitempty"`
	Server int    `json:"server,omitempty"`
}

type photosOwnerUploadResponse struct {
	Hash   string `json:"hash,omitempty"`
	Photo  string `json:"photo,omitempty"`
	Server int    `json:"server,omitempty"`
}

type photosPhoto struct {
	AccessKey string        `json:"access_key,omitempty"`
	AlbumID   int           `json:"album_id"`
	Date      int           `json:"date"`
	Height    int           `json:"height,omitempty"`
	ID        int           `json:"id"`
	Images    []photosImage `json:"images,omitempty"`
	Lat       float64       `json:"lat,omitempty"`
	Long      float64       `json:"long,omitempty"`
	OwnerID   int           `json:"owner_id"`
	PostID    int           `json:"post_id,omitempty"`
	Text      string        `json:"text,omitempty"`
	UserID    int           `json:"user_id,omitempty"`
	Width     int           `json:"width,omitempty"`
}

type photosPhotoAlbum struct {
	Created     int         `json:"created"`
	Description string      `json:"description,omitempty"`
	ID          int         `json:"id"`
	OwnerID     int         `json:"owner_id"`
	Size        int         `json:"size"`
	Thumb       photosPhoto `json:"thumb,omitempty"`
	Title       string      `json:"title"`
	Updated     int         `json:"updated"`
}

type photosPhotoAlbumFull struct {
	CanUpload          baseBoolInt                              `json:"can_upload,omitempty"`
	CommentsDisabled   baseBoolInt                              `json:"comments_disabled,omitempty"`
	Created            int                                      `json:"created"`
	Description        string                                   `json:"description,omitempty"`
	ID                 int                                      `json:"id"`
	OwnerID            int                                      `json:"owner_id"`
	PrivacyComment     []photosPhotoAlbumFullPrivacyCommentItem `json:"privacy_comment,omitempty"`
	PrivacyView        []photosPhotoAlbumFullPrivacyViewItem    `json:"privacy_view,omitempty"`
	Size               int                                      `json:"size"`
	Sizes              []photosPhotoSizes                       `json:"sizes,omitempty"`
	ThumbID            int                                      `json:"thumb_id,omitempty"`
	ThumbIsLast        baseBoolInt                              `json:"thumb_is_last,omitempty"`
	ThumbSrc           string                                   `json:"thumb_src,omitempty"`
	Title              string                                   `json:"title"`
	Updated            int                                      `json:"updated"`
	UploadByAdminsOnly baseBoolInt                              `json:"upload_by_admins_only,omitempty"`
}

// Privacy comment
type photosPhotoAlbumFullPrivacyCommentItem string

// Privacy view
type photosPhotoAlbumFullPrivacyViewItem string

type photosPhotoFull struct {
	AccessKey  string          `json:"access_key,omitempty"`
	AlbumID    int             `json:"album_id"`
	CanComment baseBoolInt     `json:"can_comment,omitempty"`
	Comments   baseObjectCount `json:"comments,omitempty"`
	Date       int             `json:"date"`
	Height     int             `json:"height,omitempty"`
	ID         int             `json:"id"`
	Images     []photosImage   `json:"images,omitempty"`
	Lat        float64         `json:"lat,omitempty"`
	Likes      baseLikes       `json:"likes,omitempty"`
	Long       float64         `json:"long,omitempty"`
	OwnerID    int             `json:"owner_id"`
	PostID     int             `json:"post_id,omitempty"`
	Reposts    baseObjectCount `json:"reposts,omitempty"`
	Tags       baseObjectCount `json:"tags,omitempty"`
	Text       string          `json:"text,omitempty"`
	UserID     int             `json:"user_id,omitempty"`
	Width      int             `json:"width,omitempty"`
}

type photosPhotoFullXtrRealOffset struct {
	AccessKey  string             `json:"access_key,omitempty"`
	AlbumID    int                `json:"album_id"`
	CanComment baseBoolInt        `json:"can_comment,omitempty"`
	Comments   baseObjectCount    `json:"comments,omitempty"`
	Date       int                `json:"date"`
	Height     int                `json:"height,omitempty"`
	Hidden     basePropertyExists `json:"hidden,omitempty"`
	ID         int                `json:"id"`
	Lat        float64            `json:"lat,omitempty"`
	Likes      baseLikes          `json:"likes,omitempty"`
	Long       float64            `json:"long,omitempty"`
	OwnerID    int                `json:"owner_id"`
	Photo1280  string             `json:"photo_1280,omitempty"`
	Photo130   string             `json:"photo_130,omitempty"`
	Photo2560  string             `json:"photo_2560,omitempty"`
	Photo604   string             `json:"photo_604,omitempty"`
	Photo75    string             `json:"photo_75,omitempty"`
	Photo807   string             `json:"photo_807,omitempty"`
	PostID     int                `json:"post_id,omitempty"`
	RealOffset int                `json:"real_offset,omitempty"`
	Reposts    baseObjectCount    `json:"reposts,omitempty"`
	Sizes      []photosPhotoSizes `json:"sizes,omitempty"`
	Tags       baseObjectCount    `json:"tags,omitempty"`
	Text       string             `json:"text,omitempty"`
	UserID     int                `json:"user_id,omitempty"`
	Width      int                `json:"width,omitempty"`
}

type photosPhotoSizes struct {
	Height int                  `json:"height"`
	Src    string               `json:"src"`
	Type   photosPhotoSizesType `json:"type"`
	Width  int                  `json:"width"`
}

// Size type
type photosPhotoSizesType string

type photosPhotoTag struct {
	Date       int         `json:"date"`
	ID         int         `json:"id"`
	PlacerID   int         `json:"placer_id"`
	TaggedName string      `json:"tagged_name"`
	UserID     int         `json:"user_id"`
	Viewed     baseBoolInt `json:"viewed"`
	X          float64     `json:"x"`
	X2         float64     `json:"x2"`
	Y          float64     `json:"y"`
	Y2         float64     `json:"y2"`
}

type photosPhotoUpload struct {
	AlbumID   int    `json:"album_id"`
	UploadURL string `json:"upload_url"`
	UserID    int    `json:"user_id"`
}

type photosPhotoUploadResponse struct {
	Aid        int    `json:"aid,omitempty"`
	Hash       string `json:"hash,omitempty"`
	PhotosList string `json:"photos_list,omitempty"`
	Server     int    `json:"server,omitempty"`
}

type photosPhotoXtrRealOffset struct {
	AccessKey  string             `json:"access_key,omitempty"`
	AlbumID    int                `json:"album_id"`
	Date       int                `json:"date"`
	Height     int                `json:"height,omitempty"`
	Hidden     basePropertyExists `json:"hidden,omitempty"`
	ID         int                `json:"id"`
	Lat        float64            `json:"lat,omitempty"`
	Long       float64            `json:"long,omitempty"`
	OwnerID    int                `json:"owner_id"`
	Photo1280  string             `json:"photo_1280,omitempty"`
	Photo130   string             `json:"photo_130,omitempty"`
	Photo2560  string             `json:"photo_2560,omitempty"`
	Photo604   string             `json:"photo_604,omitempty"`
	Photo75    string             `json:"photo_75,omitempty"`
	Photo807   string             `json:"photo_807,omitempty"`
	PostID     int                `json:"post_id,omitempty"`
	RealOffset int                `json:"real_offset,omitempty"`
	Sizes      []photosPhotoSizes `json:"sizes,omitempty"`
	Text       string             `json:"text,omitempty"`
	UserID     int                `json:"user_id,omitempty"`
	Width      int                `json:"width,omitempty"`
}

type photosPhotoXtrTagInfo struct {
	AccessKey  string             `json:"access_key,omitempty"`
	AlbumID    int                `json:"album_id"`
	Date       int                `json:"date"`
	Height     int                `json:"height,omitempty"`
	ID         int                `json:"id"`
	Lat        float64            `json:"lat,omitempty"`
	Long       float64            `json:"long,omitempty"`
	OwnerID    int                `json:"owner_id"`
	Photo1280  string             `json:"photo_1280,omitempty"`
	Photo130   string             `json:"photo_130,omitempty"`
	Photo2560  string             `json:"photo_2560,omitempty"`
	Photo604   string             `json:"photo_604,omitempty"`
	Photo75    string             `json:"photo_75,omitempty"`
	Photo807   string             `json:"photo_807,omitempty"`
	PlacerID   int                `json:"placer_id,omitempty"`
	PostID     int                `json:"post_id,omitempty"`
	Sizes      []photosPhotoSizes `json:"sizes,omitempty"`
	TagCreated int                `json:"tag_created,omitempty"`
	TagID      int                `json:"tag_id,omitempty"`
	Text       string             `json:"text,omitempty"`
	UserID     int                `json:"user_id,omitempty"`
	Width      int                `json:"width,omitempty"`
}

type photosWallUploadResponse struct {
	Hash   string `json:"hash,omitempty"`
	Photo  string `json:"photo,omitempty"`
	Server int    `json:"server,omitempty"`
}

type placesCheckin struct {
	Date         int     `json:"date"`
	Distance     int     `json:"distance,omitempty"`
	ID           int     `json:"id"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	PlaceCity    int     `json:"place_city,omitempty"`
	PlaceCountry int     `json:"place_country,omitempty"`
	PlaceID      int     `json:"place_id,omitempty"`
	PlaceIcon    string  `json:"place_icon,omitempty"`
	PlaceTitle   string  `json:"place_title,omitempty"`
	PlaceType    string  `json:"place_type,omitempty"`
	Text         string  `json:"text,omitempty"`
	UserID       int     `json:"user_id"`
}

type placesPlaceFull struct {
	Address    string  `json:"address,omitempty"`
	Checkins   int     `json:"checkins,omitempty"`
	City       int     `json:"city,omitempty"`
	Country    int     `json:"country,omitempty"`
	Created    int     `json:"created,omitempty"`
	Distance   int     `json:"distance,omitempty"`
	GroupID    int     `json:"group_id,omitempty"`
	GroupPhoto string  `json:"group_photo,omitempty"`
	ID         int     `json:"id,omitempty"`
	Icon       string  `json:"icon,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
	Title      string  `json:"title,omitempty"`
	Type       string  `json:"type,omitempty"`
}

type placesPlaceMin struct {
	Address   string  `json:"address,omitempty"`
	Checkins  int     `json:"checkins,omitempty"`
	City      int     `json:"city,omitempty"`
	Country   int     `json:"country,omitempty"`
	Created   int     `json:"created,omitempty"`
	ID        int     `json:"id,omitempty"`
	Icon      string  `json:"icon,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Title     string  `json:"title,omitempty"`
	Type      string  `json:"type,omitempty"`
}

type placesTypes struct {
	ID    int    `json:"id,omitempty"`
	Icon  string `json:"icon,omitempty"`
	Title string `json:"title,omitempty"`
}

type pollsAnswer struct {
	ID    int     `json:"id"`
	Rate  float64 `json:"rate"`
	Text  string  `json:"text"`
	Votes int     `json:"votes"`
}

type pollsPoll struct {
	Anonymous baseBoolInt   `json:"anonymous"`
	AnswerID  int           `json:"answer_id"`
	Answers   []pollsAnswer `json:"answers"`
	Created   int           `json:"created"`
	ID        int           `json:"id"`
	OwnerID   int           `json:"owner_id"`
	Question  string        `json:"question"`
	Votes     string        `json:"votes"`
}

type pollsVoters struct {
	AnswerID int              `json:"answer_id,omitempty"`
	Users    pollsVotersUsers `json:"users,omitempty"`
}

type pollsVotersUsers struct {
	Count int                    `json:"count,omitempty"`
	Items []pollsVotersUsersItem `json:"items,omitempty"`
}

// User ID
type pollsVotersUsersItem int

type searchHint struct {
	Description string            `json:"description"`
	Global      baseBoolInt       `json:"global,omitempty"`
	Group       groupsGroup       `json:"group,omitempty"`
	Profile     usersUserMin      `json:"profile,omitempty"`
	Section     searchHintSection `json:"section"`
	Type        searchHintType    `json:"type"`
}

// Section title
type searchHintSection string

// Object type
type searchHintType string

type secureLevel struct {
	Level int `json:"level,omitempty"`
	UID   int `json:"uid,omitempty"`
}

type secureSmsNotification struct {
	AppID   int    `json:"app_id,omitempty"`
	Date    int    `json:"date,omitempty"`
	ID      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	UserID  int    `json:"user_id,omitempty"`
}

type secureTokenChecked struct {
	Date    int            `json:"date,omitempty"`
	Expire  int            `json:"expire,omitempty"`
	Success baseOkResponse `json:"success,omitempty"`
	UserID  int            `json:"user_id,omitempty"`
}

type secureTransaction struct {
	Date    int `json:"date,omitempty"`
	ID      int `json:"id,omitempty"`
	UIDFrom int `json:"uid_from,omitempty"`
	UIDTo   int `json:"uid_to,omitempty"`
	Votes   int `json:"votes,omitempty"`
}

// Activity stats
type statsActivity struct {
	Comments     int `json:"comments,omitempty"`
	Copies       int `json:"copies,omitempty"`
	Hidden       int `json:"hidden,omitempty"`
	Likes        int `json:"likes,omitempty"`
	Subscribed   int `json:"subscribed,omitempty"`
	Unsubscribed int `json:"unsubscribed,omitempty"`
}

type statsCity struct {
	Count int    `json:"count,omitempty"`
	Name  string `json:"name,omitempty"`
	Value int    `json:"value,omitempty"`
}

type statsCountry struct {
	Code  string `json:"code,omitempty"`
	Count int    `json:"count,omitempty"`
	Name  string `json:"name,omitempty"`
	Value int    `json:"value,omitempty"`
}

type statsPeriod struct {
	Activity   statsActivity `json:"activity,omitempty"`
	PeriodFrom string        `json:"period_from,omitempty"`
	PeriodTo   string        `json:"period_to,omitempty"`
	Reach      statsReach    `json:"reach,omitempty"`
	Visitors   statsViews    `json:"visitors,omitempty"`
}

// Reach stats
type statsReach struct {
	Age              []statsSexAge  `json:"age,omitempty"`
	Cities           []statsCity    `json:"cities,omitempty"`
	Countries        []statsCountry `json:"countries,omitempty"`
	MobileReach      int            `json:"mobile_reach,omitempty"`
	Reach            int            `json:"reach,omitempty"`
	ReachSubscribers int            `json:"reach_subscribers,omitempty"`
	Sex              []statsSexAge  `json:"sex,omitempty"`
	SexAge           []statsSexAge  `json:"sex_age,omitempty"`
}

type statsSexAge struct {
	Count int    `json:"count,omitempty"`
	Value string `json:"value,omitempty"`
}

// Views stats
type statsViews struct {
	Age         []statsSexAge  `json:"age,omitempty"`
	Cities      []statsCity    `json:"cities,omitempty"`
	Countries   []statsCountry `json:"countries,omitempty"`
	MobileViews int            `json:"mobile_views,omitempty"`
	Sex         []statsSexAge  `json:"sex,omitempty"`
	SexAge      []statsSexAge  `json:"sex_age,omitempty"`
	Views       int            `json:"views,omitempty"`
	Visitors    int            `json:"visitors,omitempty"`
}

type statsWallpostStat struct {
	Hide             int `json:"hide,omitempty"`
	JoinGroup        int `json:"join_group,omitempty"`
	Links            int `json:"links,omitempty"`
	ReachSubscribers int `json:"reach_subscribers,omitempty"`
	ReachTotal       int `json:"reach_total,omitempty"`
	Report           int `json:"report,omitempty"`
	ToGroup          int `json:"to_group,omitempty"`
	Unsubscribe      int `json:"unsubscribe,omitempty"`
}

type statusStatus struct {
	Audio audioAudioFull `json:"audio,omitempty"`
	Text  string         `json:"text,omitempty"`
}

type storiesReplies struct {
	Count int `json:"count"`
	New   int `json:"new,omitempty"`
}

//type storiesStory struct {
//	AccessKey            string            `json:"access_key,omitempty"`
//	CanComment           baseBoolInt       `json:"can_comment,omitempty"`
//	CanReply             baseBoolInt       `json:"can_reply,omitempty"`
//	CanSee               baseBoolInt       `json:"can_see,omitempty"`
//	CanShare             baseBoolInt       `json:"can_share,omitempty"`
//	Date                 int               `json:"date,omitempty"`
//	ID                   int               `json:"id"`
//	IsDeleted            bool              `json:"is_deleted,omitempty"`
//	IsExpired            bool              `json:"is_expired,omitempty"`
//	Link                 storiesStoryLink  `json:"link,omitempty"`
//	OwnerID              int               `json:"owner_id"`
//	ParentStory          storiesStory      `json:"parent_story,omitempty"`
//	ParentStoryAccessKey string            `json:"parent_story_access_key,omitempty"`
//	ParentStoryID        int               `json:"parent_story_id,omitempty"`
//	ParentStoryOwnerID   int               `json:"parent_story_owner_id,omitempty"`
//	Photo                photosPhoto       `json:"photo,omitempty"`
//	Replies              []storiesReplies  `json:"replies,omitempty"`
//	Seen                 baseBoolInt       `json:"seen,omitempty"`
//	Type                 storiesStoryType  `json:"type,omitempty"`
//	Video                storiesStoryVideo `json:"video,omitempty"`
//	Views                int               `json:"views,omitempty"`
//}

type storiesStoryLink struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

type storiesStoryStats struct {
	Answer      storiesStoryStatsStat `json:"answer"`
	Bans        storiesStoryStatsStat `json:"bans"`
	OpenLink    storiesStoryStatsStat `json:"open_link"`
	Replies     storiesStoryStatsStat `json:"replies"`
	Shares      storiesStoryStatsStat `json:"shares"`
	Subscribers storiesStoryStatsStat `json:"subscribers"`
	Views       storiesStoryStatsStat `json:"views"`
}

type storiesStoryStatsStat struct {
	Count int                    `json:"count,omitempty"`
	State storiesStoryStatsState `json:"state"`
}

// Statistic state
type storiesStoryStatsState string

// Story type.
type storiesStoryType string

type storiesStoryVideo struct {
}

type usersCareer struct {
	CityID    int    `json:"city_id,omitempty"`
	Company   string `json:"company,omitempty"`
	CountryID int    `json:"country_id,omitempty"`
	From      int    `json:"from,omitempty"`
	GroupID   int    `json:"group_id,omitempty"`
	Position  string `json:"position,omitempty"`
	Until     int    `json:"until,omitempty"`
}

type usersCropPhoto struct {
	Crop  usersCropPhotoCrop `json:"crop,omitempty"`
	Photo photosPhoto        `json:"photo,omitempty"`
	Rect  usersCropPhotoRect `json:"rect,omitempty"`
}

type usersCropPhotoCrop struct {
	X  float64 `json:"x,omitempty"`
	X2 float64 `json:"x2,omitempty"`
	Y  float64 `json:"y,omitempty"`
	Y2 float64 `json:"y2,omitempty"`
}

type usersCropPhotoRect struct {
	X  float64 `json:"x,omitempty"`
	X2 float64 `json:"x2,omitempty"`
	Y  float64 `json:"y,omitempty"`
	Y2 float64 `json:"y2,omitempty"`
}

type usersExports struct {
	Facebook    int `json:"facebook,omitempty"`
	Livejournal int `json:"livejournal,omitempty"`
	Twitter     int `json:"twitter,omitempty"`
}

type usersLastSeen struct {
	Platform int `json:"platform,omitempty"`
	Time     int `json:"time,omitempty"`
}

type usersMilitary struct {
	CountryID int    `json:"country_id,omitempty"`
	From      int    `json:"from,omitempty"`
	Unit      string `json:"unit,omitempty"`
	UnitID    int    `json:"unit_id,omitempty"`
	Until     int    `json:"until,omitempty"`
}

type usersOccupation struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type usersPersonal struct {
	Alcohol    int    `json:"alcohol,omitempty"`
	InspiredBy string `json:"inspired_by,omitempty"`
	Langs      []lang `json:"langs,omitempty"`
	LifeMain   int    `json:"life_main,omitempty"`
	PeopleMain int    `json:"people_main,omitempty"`
	Political  int    `json:"political,omitempty"`
	Religion   string `json:"religion,omitempty"`
	Smoking    int    `json:"smoking,omitempty"`
}

type usersRelative struct {
	ID   int    `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

type usersSchool struct {
	City          int    `json:"city,omitempty"`
	Class         string `json:"class,omitempty"`
	Country       int    `json:"country,omitempty"`
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Type          int    `json:"type,omitempty"`
	TypeStr       string `json:"type_str,omitempty"`
	YearFrom      int    `json:"year_from,omitempty"`
	YearGraduated int    `json:"year_graduated,omitempty"`
	YearTo        int    `json:"year_to,omitempty"`
}

type usersUniversity struct {
	Chair           int    `json:"chair,omitempty"`
	ChairName       string `json:"chair_name,omitempty"`
	City            int    `json:"city,omitempty"`
	Country         int    `json:"country,omitempty"`
	EducationForm   string `json:"education_form,omitempty"`
	EducationStatus string `json:"education_status,omitempty"`
	Faculty         int    `json:"faculty,omitempty"`
	FacultyName     string `json:"faculty_name,omitempty"`
	Graduation      int    `json:"graduation,omitempty"`
	ID              int    `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
}

type usersUser struct {
}

type usersUserBroadcast struct {
}

type usersUserCounters struct {
	Albums        int `json:"albums,omitempty"`
	Audios        int `json:"audios,omitempty"`
	Followers     int `json:"followers,omitempty"`
	Friends       int `json:"friends,omitempty"`
	Gifts         int `json:"gifts,omitempty"`
	Groups        int `json:"groups,omitempty"`
	Notes         int `json:"notes,omitempty"`
	OnlineFriends int `json:"online_friends,omitempty"`
	Pages         int `json:"pages,omitempty"`
	Photos        int `json:"photos,omitempty"`
	Subscriptions int `json:"subscriptions,omitempty"`
	UserPhotos    int `json:"user_photos,omitempty"`
	UserVideos    int `json:"user_videos,omitempty"`
	Videos        int `json:"videos,omitempty"`
}

type usersUserFull struct {
}

type usersUserFullXtrType struct {
}

type usersUserLim struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	NameGen string `json:"name_gen,omitempty"`
	Photo   string `json:"photo,omitempty"`
}

type usersUserMin struct {
	Deactivated string `json:"deactivated,omitempty"`
	FirstName   string `json:"first_name"`
	Hidden      int    `json:"hidden,omitempty"`
	ID          int    `json:"id"`
	LastName    string `json:"last_name"`
}

type utilsDomainResolved struct {
	ObjectID int                     `json:"object_id,omitempty"`
	Type     utilsDomainResolvedType `json:"type,omitempty"`
}

// Object type
type utilsDomainResolvedType string

type utilsLastShortenedLink struct {
	AccessKey string `json:"access_key,omitempty"`
	Key       string `json:"key,omitempty"`
	ShortURL  string `json:"short_url,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
	URL       string `json:"url,omitempty"`
	Views     int    `json:"views,omitempty"`
}

type utilsLinkChecked struct {
	Link   string                 `json:"link,omitempty"`
	Status utilsLinkCheckedStatus `json:"status,omitempty"`
}

// Link status
type utilsLinkCheckedStatus string

type utilsLinkStats struct {
	Key   string       `json:"key,omitempty"`
	Stats []utilsStats `json:"stats,omitempty"`
}

type utilsLinkStatsExtended struct {
	Key   string               `json:"key,omitempty"`
	Stats []utilsStatsExtended `json:"stats,omitempty"`
}

type utilsShortLink struct {
	AccessKey string `json:"access_key,omitempty"`
	Key       string `json:"key,omitempty"`
	ShortURL  string `json:"short_url,omitempty"`
	URL       string `json:"url,omitempty"`
}

type utilsStats struct {
	Timestamp int `json:"timestamp,omitempty"`
	Views     int `json:"views,omitempty"`
}

type utilsStatsCity struct {
	CityID int `json:"city_id,omitempty"`
	Views  int `json:"views,omitempty"`
}

type utilsStatsCountry struct {
	CountryID int `json:"country_id,omitempty"`
	Views     int `json:"views,omitempty"`
}

type utilsStatsExtended struct {
	Cities    []utilsStatsCity    `json:"cities,omitempty"`
	Countries []utilsStatsCountry `json:"countries,omitempty"`
	SexAge    []utilsStatsSexAge  `json:"sex_age,omitempty"`
	Timestamp int                 `json:"timestamp,omitempty"`
	Views     int                 `json:"views,omitempty"`
}

type utilsStatsSexAge struct {
	AgeRange string `json:"age_range,omitempty"`
	Female   int    `json:"female,omitempty"`
	Male     int    `json:"male,omitempty"`
}

type videoCatBlock struct {
	CanHide baseBoolInt       `json:"can_hide"`
	ID      int               `json:"id"`
	Items   []videoCatElement `json:"items"`
	Name    string            `json:"name"`
	Next    string            `json:"next"`
	Type    videoCatBlockView `json:"type,omitempty"`
	View    videoCatBlockView `json:"view"`
}

// Type of view
type videoCatBlockView string

type videoCatElement struct {
	CanAdd      baseBoolInt         `json:"can_add,omitempty"`
	CanEdit     baseBoolInt         `json:"can_edit,omitempty"`
	Comments    int                 `json:"comments,omitempty"`
	Count       int                 `json:"count,omitempty"`
	Date        int                 `json:"date,omitempty"`
	Description string              `json:"description,omitempty"`
	Duration    int                 `json:"duration,omitempty"`
	ID          int                 `json:"id"`
	IsPrivate   baseBoolInt         `json:"is_private,omitempty"`
	OwnerID     int                 `json:"owner_id"`
	Photo130    string              `json:"photo_130,omitempty"`
	Photo160    string              `json:"photo_160,omitempty"`
	Photo320    string              `json:"photo_320,omitempty"`
	Photo640    string              `json:"photo_640,omitempty"`
	Photo800    string              `json:"photo_800,omitempty"`
	Title       string              `json:"title"`
	Type        videoCatElementType `json:"type"`
	UpdatedTime int                 `json:"updated_time,omitempty"`
	Views       int                 `json:"views,omitempty"`
}

// Element type
type videoCatElementType string

type videoSaveResult struct {
	Description string `json:"description,omitempty"`
	OwnerID     int    `json:"owner_id,omitempty"`
	Title       string `json:"title,omitempty"`
	UploadURL   string `json:"upload_url,omitempty"`
	VideoID     int    `json:"video_id,omitempty"`
}

type videoUploadResponse struct {
	Size    int `json:"size,omitempty"`
	VideoID int `json:"video_id,omitempty"`
}

type videoVideo struct {
	AccessKey   string             `json:"access_key,omitempty"`
	AddingDate  int                `json:"adding_date,omitempty"`
	CanAdd      baseBoolInt        `json:"can_add,omitempty"`
	CanEdit     baseBoolInt        `json:"can_edit,omitempty"`
	Comments    int                `json:"comments,omitempty"`
	Date        int                `json:"date,omitempty"`
	Description string             `json:"description,omitempty"`
	Duration    int                `json:"duration,omitempty"`
	Files       videoVideoFiles    `json:"files,omitempty"`
	ID          int                `json:"id,omitempty"`
	Live        basePropertyExists `json:"live,omitempty"`
	OwnerID     int                `json:"owner_id,omitempty"`
	Photo130    string             `json:"photo_130,omitempty"`
	Photo320    string             `json:"photo_320,omitempty"`
	Photo800    string             `json:"photo_800,omitempty"`
	Player      string             `json:"player,omitempty"`
	Processing  basePropertyExists `json:"processing,omitempty"`
	Title       string             `json:"title,omitempty"`
	Views       int                `json:"views,omitempty"`
}

type videoVideoAlbum struct {
	ID      int    `json:"id"`
	OwnerID int    `json:"owner_id"`
	Title   string `json:"title"`
}

type videoVideoAlbumFull struct {
	Count       int    `json:"count"`
	ID          int    `json:"id"`
	IsSystem    int    `json:"is_system,omitempty"`
	OwnerID     int    `json:"owner_id"`
	Photo160    string `json:"photo_160,omitempty"`
	Photo320    string `json:"photo_320,omitempty"`
	Title       string `json:"title"`
	UpdatedTime int    `json:"updated_time"`
}

type videoVideoFiles struct {
	External string `json:"external,omitempty"`
	Mp1080   string `json:"mp_1080,omitempty"`
	Mp240    string `json:"mp_240,omitempty"`
	Mp360    string `json:"mp_360,omitempty"`
	Mp480    string `json:"mp_480,omitempty"`
	Mp720    string `json:"mp_720,omitempty"`
}

type videoVideoFull struct {
	AccessKey      string                             `json:"access_key,omitempty"`
	AddingDate     int                                `json:"adding_date,omitempty"`
	CanAdd         baseBoolInt                        `json:"can_add,omitempty"`
	CanComment     baseBoolInt                        `json:"can_comment,omitempty"`
	CanEdit        baseBoolInt                        `json:"can_edit,omitempty"`
	CanRepost      baseBoolInt                        `json:"can_repost,omitempty"`
	Comments       int                                `json:"comments,omitempty"`
	Date           int                                `json:"date,omitempty"`
	Description    string                             `json:"description,omitempty"`
	Duration       int                                `json:"duration,omitempty"`
	Files          videoVideoFiles                    `json:"files,omitempty"`
	ID             int                                `json:"id,omitempty"`
	Likes          baseLikes                          `json:"likes,omitempty"`
	Live           basePropertyExists                 `json:"live,omitempty"`
	OwnerID        int                                `json:"owner_id,omitempty"`
	Photo130       string                             `json:"photo_130,omitempty"`
	Photo320       string                             `json:"photo_320,omitempty"`
	Photo800       string                             `json:"photo_800,omitempty"`
	Player         string                             `json:"player,omitempty"`
	PrivacyComment []videoVideoFullPrivacyCommentItem `json:"privacy_comment,omitempty"`
	PrivacyView    []videoVideoFullPrivacyViewItem    `json:"privacy_view,omitempty"`
	Processing     basePropertyExists                 `json:"processing,omitempty"`
	Repeat         baseBoolInt                        `json:"repeat,omitempty"`
	Title          string                             `json:"title,omitempty"`
	Views          int                                `json:"views,omitempty"`
}

// Privacy comment
type videoVideoFullPrivacyCommentItem string

// Privacy view
type videoVideoFullPrivacyViewItem string

type videoVideoTag struct {
	Date       int         `json:"date"`
	ID         int         `json:"id"`
	PlacerID   int         `json:"placer_id"`
	TaggedName string      `json:"tagged_name"`
	UserID     int         `json:"user_id"`
	Viewed     baseBoolInt `json:"viewed"`
}

type videoVideoTagInfo struct {
	AccessKey   string             `json:"access_key,omitempty"`
	AddingDate  int                `json:"adding_date,omitempty"`
	CanAdd      baseBoolInt        `json:"can_add,omitempty"`
	CanEdit     baseBoolInt        `json:"can_edit,omitempty"`
	Comments    int                `json:"comments,omitempty"`
	Date        int                `json:"date,omitempty"`
	Description string             `json:"description,omitempty"`
	Duration    int                `json:"duration,omitempty"`
	Files       videoVideoFiles    `json:"files,omitempty"`
	ID          int                `json:"id,omitempty"`
	Live        basePropertyExists `json:"live,omitempty"`
	OwnerID     int                `json:"owner_id,omitempty"`
	Photo130    string             `json:"photo_130,omitempty"`
	Photo320    string             `json:"photo_320,omitempty"`
	Photo800    string             `json:"photo_800,omitempty"`
	PlacerID    int                `json:"placer_id,omitempty"`
	Player      string             `json:"player,omitempty"`
	Processing  basePropertyExists `json:"processing,omitempty"`
	TagCreated  int                `json:"tag_created,omitempty"`
	TagID       int                `json:"tag_id,omitempty"`
	Title       string             `json:"title,omitempty"`
	Views       int                `json:"views,omitempty"`
}

type wallAppPost struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Photo130 string `json:"photo_130,omitempty"`
	Photo604 string `json:"photo_604,omitempty"`
}

type wallAttachedNote struct {
	Comments     int    `json:"comments"`
	Date         int    `json:"date"`
	ID           int    `json:"id"`
	OwnerID      int    `json:"owner_id"`
	ReadComments int    `json:"read_comments"`
	Title        string `json:"title"`
	ViewURL      string `json:"view_url"`
}

type wallCommentAttachment struct {
	Audio             audioAudioFull            `json:"audio,omitempty"`
	Doc               docsDoc                   `json:"doc,omitempty"`
	Link              baseLink                  `json:"link,omitempty"`
	Market            marketMarketItem          `json:"market,omitempty"`
	MarketMarketAlbum marketMarketAlbum         `json:"market_market_album,omitempty"`
	Note              wallAttachedNote          `json:"note,omitempty"`
	Page              pagesWikipageFull         `json:"page,omitempty"`
	Photo             photosPhoto               `json:"photo,omitempty"`
	Sticker           baseSticker               `json:"sticker,omitempty"`
	Type              wallCommentAttachmentType `json:"type"`
	Video             videoVideo                `json:"video,omitempty"`
}

// Attachment type
type wallCommentAttachmentType string

type wallGraffiti struct {
	ID       int    `json:"id,omitempty"`
	OwnerID  int    `json:"owner_id,omitempty"`
	Photo200 string `json:"photo_200,omitempty"`
	Photo586 string `json:"photo_586,omitempty"`
}

type wallPostSource struct {
	Data     string             `json:"data,omitempty"`
	Platform string             `json:"platform,omitempty"`
	Type     wallPostSourceType `json:"type,omitempty"`
	URL      string             `json:"url,omitempty"`
}

// Type of post source
type wallPostSourceType string

// Post type
type wallPostType string

type wallPostedPhoto struct {
	ID       int    `json:"id,omitempty"`
	OwnerID  int    `json:"owner_id,omitempty"`
	Photo130 string `json:"photo_130,omitempty"`
	Photo604 string `json:"photo_604,omitempty"`
}

type wallViews struct {
	Count int `json:"count,omitempty"`
}

type wallWallComment struct {
	Attachments    []wallCommentAttachment `json:"attachments,omitempty"`
	Date           int                     `json:"date"`
	FromID         int                     `json:"from_id"`
	ID             int                     `json:"id"`
	Likes          baseLikesInfo           `json:"likes,omitempty"`
	RealOffset     int                     `json:"real_offset,omitempty"`
	ReplyToComment int                     `json:"reply_to_comment,omitempty"`
	ReplyToUser    int                     `json:"reply_to_user,omitempty"`
	Text           string                  `json:"text"`
}

type wallWallpost struct {
	AccessKey   string                   `json:"access_key,omitempty"`
	Attachments []wallWallpostAttachment `json:"attachments,omitempty"`
	Date        int                      `json:"date,omitempty"`
	FromID      int                      `json:"from_id,omitempty"`
	Geo         baseGeo                  `json:"geo,omitempty"`
	ID          int                      `json:"id,omitempty"`
	OwnerID     int                      `json:"owner_id,omitempty"`
	PostSource  wallPostSource           `json:"post_source,omitempty"`
	PostType    wallPostType             `json:"post_type,omitempty"`
	SignerID    int                      `json:"signer_id,omitempty"`
	Text        string                   `json:"text,omitempty"`
	Views       wallViews                `json:"views,omitempty"`
}

type wallWallpostAttached struct {
	Attachments []wallWallpostAttachment `json:"attachments,omitempty"`
	CanDelete   baseBoolInt              `json:"can_delete,omitempty"`
	Comments    baseCommentsInfo         `json:"comments,omitempty"`
	CopyOwnerID int                      `json:"copy_owner_id,omitempty"`
	CopyPostID  int                      `json:"copy_post_id,omitempty"`
	CopyText    string                   `json:"copy_text,omitempty"`
	Date        int                      `json:"date,omitempty"`
	FromID      int                      `json:"from_id,omitempty"`
	Geo         baseGeo                  `json:"geo,omitempty"`
	ID          int                      `json:"id,omitempty"`
	Likes       baseLikesInfo            `json:"likes,omitempty"`
	PostSource  wallPostSource           `json:"post_source,omitempty"`
	PostType    wallPostType             `json:"post_type,omitempty"`
	Reposts     baseRepostsInfo          `json:"reposts,omitempty"`
	SignerID    int                      `json:"signer_id,omitempty"`
	Text        string                   `json:"text,omitempty"`
	ToID        int                      `json:"to_id,omitempty"`
}

type wallWallpostAttachment struct {
	Album             photosPhotoAlbum           `json:"album,omitempty"`
	App               wallAppPost                `json:"app,omitempty"`
	Audio             audioAudioFull             `json:"audio,omitempty"`
	Doc               docsDoc                    `json:"doc,omitempty"`
	Graffiti          wallGraffiti               `json:"graffiti,omitempty"`
	Link              baseLink                   `json:"link,omitempty"`
	Market            marketMarketItem           `json:"market,omitempty"`
	MarketMarketAlbum marketMarketAlbum          `json:"market_market_album,omitempty"`
	Note              wallAttachedNote           `json:"note,omitempty"`
	Page              pagesWikipageFull          `json:"page,omitempty"`
	Photo             photosPhoto                `json:"photo,omitempty"`
	PhotosList        []photosListItem           `json:"photos_list,omitempty"`
	Poll              pollsPoll                  `json:"poll,omitempty"`
	PostedPhoto       wallPostedPhoto            `json:"posted_photo,omitempty"`
	Type              wallWallpostAttachmentType `json:"type"`
	Video             videoVideo                 `json:"video,omitempty"`
}

// Attachment type
type wallWallpostAttachmentType string

type wallWallpostFull struct {
}

type wallWallpostToID struct {
	Attachments []wallWallpostAttachment `json:"attachments,omitempty"`
	Comments    baseCommentsInfo         `json:"comments,omitempty"`
	CopyOwnerID int                      `json:"copy_owner_id,omitempty"`
	CopyPostID  int                      `json:"copy_post_id,omitempty"`
	Date        int                      `json:"date,omitempty"`
	FromID      int                      `json:"from_id,omitempty"`
	Geo         baseGeo                  `json:"geo,omitempty"`
	ID          int                      `json:"id,omitempty"`
	Likes       baseLikesInfo            `json:"likes,omitempty"`
	PostID      int                      `json:"post_id,omitempty"`
	PostSource  wallPostSource           `json:"post_source,omitempty"`
	PostType    wallPostType             `json:"post_type,omitempty"`
	Reposts     baseRepostsInfo          `json:"reposts,omitempty"`
	SignerID    int                      `json:"signer_id,omitempty"`
	Text        string                   `json:"text,omitempty"`
	ToID        int                      `json:"to_id,omitempty"`
}

type widgetsCommentMedia struct {
	ItemID   int                     `json:"item_id,omitempty"`
	OwnerID  int                     `json:"owner_id,omitempty"`
	ThumbSrc string                  `json:"thumb_src,omitempty"`
	Type     widgetsCommentMediaType `json:"type,omitempty"`
}

// Media type
type widgetsCommentMediaType string

type widgetsCommentReplies struct {
	CanPost baseBoolInt                 `json:"can_post,omitempty"`
	Count   int                         `json:"count,omitempty"`
	Replies []widgetsCommentRepliesItem `json:"replies,omitempty"`
}

type widgetsCommentRepliesItem struct {
	Cid   int                `json:"cid,omitempty"`
	Date  int                `json:"date,omitempty"`
	Likes widgetsWidgetLikes `json:"likes,omitempty"`
	Text  string             `json:"text,omitempty"`
	UID   int                `json:"uid,omitempty"`
	User  usersUserFull      `json:"user,omitempty"`
}

type widgetsWidgetComment struct {
	Attachments []wallCommentAttachment `json:"attachments,omitempty"`
	CanDelete   baseBoolInt             `json:"can_delete,omitempty"`
	Comments    widgetsCommentReplies   `json:"comments,omitempty"`
	Date        int                     `json:"date"`
	FromID      int                     `json:"from_id"`
	ID          int                     `json:"id"`
	Likes       baseLikesInfo           `json:"likes,omitempty"`
	Media       widgetsCommentMedia     `json:"media,omitempty"`
	PostSource  wallPostSource          `json:"post_source,omitempty"`
	PostType    int                     `json:"post_type"`
	Reposts     baseRepostsInfo         `json:"reposts,omitempty"`
	Text        string                  `json:"text"`
	ToID        int                     `json:"to_id"`
	User        usersUserFull           `json:"user,omitempty"`
}

type widgetsWidgetLikes struct {
	Count int `json:"count,omitempty"`
}

type widgetsWidgetPage struct {
	Comments    baseObjectCount `json:"comments,omitempty"`
	Date        int             `json:"date,omitempty"`
	Description string          `json:"description,omitempty"`
	ID          int             `json:"id,omitempty"`
	Likes       baseObjectCount `json:"likes,omitempty"`
	PageID      string          `json:"page_id,omitempty"`
	Photo       string          `json:"photo,omitempty"`
	Title       string          `json:"title,omitempty"`
	URL         string          `json:"url,omitempty"`
}
