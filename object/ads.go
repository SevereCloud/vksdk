package object

type adsAccesses struct {
	ClientID string `json:"client_id"`
	Role     string `json:"role"`
}

type adsAccount struct {
	AccessRole    string `json:"access_role"`
	AccountID     int    `json:"account_id"`
	AccountStatus int    `json:"account_status"`
	AccountType   string `json:"account_type"`
}

type adsAdLayout struct {
	AdFormat    int    `json:"ad_format"`
	CampaignID  int    `json:"campaign_id"`
	CostType    int    `json:"cost_type"`
	Description string `json:"description"`
	ID          int    `json:"id"`
	ImageSrc    int    `json:"image_src"`
	ImageSrc2x  int    `json:"image_src_2x"`
	LinkDomain  string `json:"link_domain"`
	LinkURL     string `json:"link_url"`
	PreviewLink string `json:"preview_link"`
	Title       string `json:"title"`
	Video       int    `json:"video"`
}

type adsCampaign struct {
	AllLimit  string `json:"all_limit"`
	DayLimit  string `json:"day_limit"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	StartTime int    `json:"start_time"`
	Status    int    `json:"status"`
	StopTime  int    `json:"stop_time"`
	Type      string `json:"type"`
}

type adsCategory struct {
	ID            int                  `json:"id"`
	Name          string               `json:"name"`
	Subcategories []baseObjectWithName `json:"subcategories"`
}

type adsClient struct {
	AllLimit string `json:"all_limit"`
	DayLimit string `json:"day_limit"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
}

type adsCriteria struct {
	AgeFrom              int    `json:"age_from"`
	AgeTo                int    `json:"age_to"`
	Apps                 string `json:"apps"`
	AppsNot              string `json:"apps_not"`
	Birthday             int    `json:"birthday"`
	Cities               string `json:"cities"`
	CitiesNot            string `json:"cities_not"`
	Country              int    `json:"country"`
	Districts            string `json:"districts"`
	Groups               string `json:"groups"`
	InterestCategories   string `json:"interest_categories"`
	Interests            string `json:"interests"`
	Paying               int    `json:"paying"`
	Positions            string `json:"positions"`
	Religions            string `json:"religions"`
	RetargetingGroups    string `json:"retargeting_groups"`
	RetargetingGroupsNot string `json:"retargeting_groups_not"`
	SchoolFrom           int    `json:"school_from"`
	SchoolTo             int    `json:"school_to"`
	Schools              string `json:"schools"`
	Sex                  int    `json:"sex"`
	Stations             string `json:"stations"`
	Statuses             string `json:"statuses"`
	Streets              string `json:"streets"`
	Travellers           int    `json:"travellers"`
	UniFrom              int    `json:"uni_from"`
	UniTo                int    `json:"uni_to"`
	UserBrowsers         string `json:"user_browsers"`
	UserDevices          string `json:"user_devices"`
	UserOs               string `json:"user_os"`
}

type adsDemoStats struct {
	ID    int                `json:"id"`
	Stats adsDemostatsFormat `json:"stats"`
	Type  string             `json:"type"`
}

type adsDemostatsFormat struct {
	Age     []adsStatsAge    `json:"age"`
	Cities  []adsStatsCities `json:"cities"`
	Day     string           `json:"day"`
	Month   string           `json:"month"`
	Overall int              `json:"overall"`
	Sex     []adsStatsSex    `json:"sex"`
	SexAge  []adsStatsSexAge `json:"sex_age"`
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

type adsParagraphs struct {
	Paragraph string `json:"paragraph"`
}

type adsRejectReason struct {
	Comment string     `json:"comment"`
	Rules   []adsRules `json:"rules"`
}

type adsRules struct {
	Paragraphs []adsParagraphs `json:"paragraphs"`
	Title      string          `json:"title"`
}

type adsStats struct {
	ID    int            `json:"id"`
	Stats adsStatsFormat `json:"stats"`
	Type  string         `json:"type"`
}

type adsStatsAge struct {
	ClicksRate      float64 `json:"clicks_rate"`
	ImpressionsRate float64 `json:"impressions_rate"`
	Value           string  `json:"value"`
}

type adsStatsCities struct {
	ClicksRate      float64 `json:"clicks_rate"`
	ImpressionsRate float64 `json:"impressions_rate"`
	Name            string  `json:"name"`
	Value           int     `json:"value"`
}

type adsStatsFormat struct {
	Clicks          int    `json:"clicks"`
	Day             string `json:"day"`
	Impressions     int    `json:"impressions"`
	JoinRate        int    `json:"join_rate"`
	Month           string `json:"month"`
	Overall         int    `json:"overall"`
	Reach           int    `json:"reach"`
	Spent           int    `json:"spent"`
	VideoClicksSite int    `json:"video_clicks_site"`
	VideoViews      int    `json:"video_views"`
	VideoViewsFull  int    `json:"video_views_full"`
	VideoViewsHalf  int    `json:"video_views_half"`
}

type adsStatsSex struct {
	ClicksRate      float64 `json:"clicks_rate"`
	ImpressionsRate float64 `json:"impressions_rate"`
	Value           string  `json:"value"`
}

type adsStatsSexAge struct {
	ClicksRate      float64 `json:"clicks_rate"`
	ImpressionsRate float64 `json:"impressions_rate"`
	Value           string  `json:"value"`
}

type adsTargSettings struct {
}

type adsTargStats struct {
	AudienceCount  int     `json:"audience_count"`
	RecommendedCpc float64 `json:"recommended_cpc"`
	RecommendedCpm float64 `json:"recommended_cpm"`
}

type adsTargSuggestions struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type adsTargSuggestionsCities struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Parent string `json:"parent"`
}

type adsTargSuggestionsRegions struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type adsTargSuggestionsSchools struct {
	Desc   string `json:"desc"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Parent string `json:"parent"`
	Type   string `json:"type"`
}

type adsTargetGroup struct {
	AudienceCount int    `json:"audience_count"`
	Domain        string `json:"domain"`
	ID            int    `json:"id"`
	Lifetime      int    `json:"lifetime"`
	Name          string `json:"name"`
	Pixel         string `json:"pixel"`
}

type adsUsers struct {
	Accesses []adsAccesses `json:"accesses"`
	UserID   int           `json:"user_id"`
}
