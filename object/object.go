package object // import "github.com/SevereCloud/vksdk/object"

import "encoding/json"

// Attachment interface
type Attachment interface {
	ToAttachment() string
}

// JSONObject interface
type JSONObject interface {
	ToJSON() string
}

// BaseCountry struct
type BaseCountry struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// BaseObject struct
type BaseObject struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// BaseObjectCount struct
type BaseObjectCount struct {
	Count int `json:"count"`
}

// BaseObjectWithName struct
type BaseObjectWithName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// BaseRequestParam struct
type BaseRequestParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// GroupEvent struct
type GroupEvent struct {
	Type    string          `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int             `json:"group_id"`
	EventID string          `json:"event_id"`
	Secret  string          `json:"secret"`
}

// LongpollResponse struct
type LongpollResponse struct {
	Ts      int             `json:"ts"`
	Updates [][]interface{} `json:"updates"`
	Failed  int             `json:"failed"`
}

// LongpollBotResponse struct
type LongpollBotResponse struct {
	Ts      string       `json:"ts"`
	Updates []GroupEvent `json:"updates"`
	Failed  int          `json:"failed"`
}

// BaseCommentsInfo struct
type BaseCommentsInfo struct {
	CanPost       int  `json:"can_post"`
	Count         int  `json:"count"`
	GroupsCanPost bool `json:"groups_can_post"`
	CanClose      bool `json:"can_close"`
	CanOpen       bool `json:"can_open"`
}

// BaseGeo struct
type BaseGeo struct {
	Coordinates string    `json:"coordinates"`
	Place       BasePlace `json:"place"`
	Showmap     int       `json:"showmap"`
	Type        string    `json:"type"`
}

// BaseImage struct
type BaseImage struct {
	Height float64 `json:"height"`
	URL    string  `json:"url"`
	Width  float64 `json:"width"`
	Type   string  `json:"type"`
}

// BaseLikes struct
type BaseLikes struct {
	Count     int `json:"count"`
	UserLikes int `json:"user_likes"`
}

// BaseLikesInfo struct
type BaseLikesInfo struct {
	CanLike int `json:"can_like"`
	// BUG(VK): https://github.com/SevereCloud/vksdk/issues/55
	// CanPublish int `json:"can_publish"`
	Count     int `json:"count"`
	UserLikes int `json:"user_likes"`
}

// BaseLink struct
type BaseLink struct {
	Application BaseLinkApplication `json:"application"`
	Button      BaseLinkButton      `json:"button"`
	Caption     string              `json:"caption"`
	Description string              `json:"description"`
	Photo       PhotosPhoto         `json:"photo"`
	PreviewPage string              `json:"preview_page"`
	PreviewURL  string              `json:"preview_url"`
	Product     BaseLinkProduct     `json:"product"`
	Rating      BaseLinkRating      `json:"rating"`
	Title       string              `json:"title"`
	URL         string              `json:"url"`
}

// BaseLinkApplication struct
type BaseLinkApplication struct {
	AppID float64                  `json:"app_id"`
	Store BaseLinkApplicationStore `json:"store"`
}

// BaseLinkApplicationStore struct
type BaseLinkApplicationStore struct {
	ID   float64 `json:"id"`
	Name string  `json:"name"`
}

// BaseLinkButton struct
type BaseLinkButton struct {
	Action BaseLinkButtonAction `json:"action"`
	Title  string               `json:"title"`
}

// BaseLinkButtonAction struct
type BaseLinkButtonAction struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// BaseLinkProduct struct
type BaseLinkProduct struct {
	Price       MarketPrice `json:"price"`
	Merchant    string      `json:"merchant"`
	OrdersCount int         `json:"orders_count"`
}

// BaseLinkRating struct
type BaseLinkRating struct {
	ReviewsCount int     `json:"reviews_count"`
	Stars        float64 `json:"stars"`
}

// BasePlace struct
type BasePlace struct {
	Address        string             `json:"address"`
	Checkins       int                `json:"checkins"`
	City           interface{}        `json:"city"` // BUG(VK): https://github.com/VKCOM/vk-api-schema/issues/143
	Country        interface{}        `json:"country"`
	Created        int                `json:"created"`
	ID             int                `json:"id"`
	Icon           string             `json:"icon"`
	Latitude       float64            `json:"latitude"`
	Longitude      float64            `json:"longitude"`
	Title          string             `json:"title"`
	Type           string             `json:"type"`
	IsDeleted      bool               `json:"is_deleted"`
	TotalCheckins  int                `json:"total_checkins"`
	Updated        int                `json:"updated"`
	CategoryObject BaseCategoryObject `json:"category_object"`
}

// BaseCategoryObject struct
type BaseCategoryObject struct {
	ID    int         `json:"id"`
	Title string      `json:"title"`
	Icons []BaseImage `json:"icons"`
}

// BaseRepostsInfo struct
type BaseRepostsInfo struct {
	Count        int `json:"count"`
	UserReposted int `json:"user_reposted"`
}

// BaseSticker struct
type BaseSticker struct {
	Images               []BaseImage `json:"images"`
	ImagesWithBackground []BaseImage `json:"images_with_background"`
	ProductID            int         `json:"product_id"`
	StickerID            int         `json:"sticker_id"`
}

// BaseUserID struct
type BaseUserID struct {
	UserID int `json:"user_id"`
}

// EventsEventAttach struct
type EventsEventAttach struct {
	Address      string `json:"address,omitempty"`       // address of event
	ButtonText   string `json:"button_text"`             // text of attach
	Friends      []int  `json:"friends"`                 // array of friends ids
	ID           int    `json:"id"`                      // event ID
	IsFavorite   bool   `json:"is_favorite"`             // is favorite
	MemberStatus int    `json:"member_status,omitempty"` // Current user's member status
	Text         string `json:"text"`                    // text of attach
	Time         int    `json:"time,omitempty"`          // event start time
}

// OauthError struct
type OauthError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	RedirectURI      string `json:"redirect_uri"`
}

// Article struct
type Article struct {
	ID            int         `json:"id"`
	OwnerID       int         `json:"owner_id"`
	OwnerName     string      `json:"owner_name"`
	OwnerPhoto    string      `json:"owner_photo"`
	State         string      `json:"state"`
	CanReport     bool        `json:"can_report"`
	IsFavorite    bool        `json:"is_favorite"`
	Title         string      `json:"title"`
	Subtitle      string      `json:"subtitle"`
	Views         int         `json:"views"`
	Shares        int         `json:"shares"`
	URL           string      `json:"url"`
	ViewURL       string      `json:"view_url"`
	AccessKey     string      `json:"access_key"`
	PublishedDate int         `json:"published_date"`
	Photo         PhotosPhoto `json:"photo"`
}

// Error struct
type Error struct {
	Code          int                `json:"error_code"`
	Message       string             `json:"error_msg"`
	Text          string             `json:"error_text"`
	CaptchaSID    string             `json:"captcha_sid"`
	CaptchaImg    string             `json:"captcha_img"`
	RequestParams []BaseRequestParam `json:"request_params"`
}

// ExtendedResponse struct
type ExtendedResponse struct {
	Profiles []UsersUser   `json:"profiles"`
	Groups   []GroupsGroup `json:"groups"`
}

// ClientInfo struct
type ClientInfo struct {
	ButtonActions  []string `json:"button_actions"`
	Keyboard       bool     `json:"keyboard"`
	InlineKeyboard bool     `json:"inline_keyboard"`
	Carousel       bool     `json:"carousel"`
	LangID         int      `json:"lang_id"`
}

// Language code
const (
	LangRU               = 0   // Русский
	LangUK               = 1   // Українська
	LangBE               = 2   // Беларуская (тарашкевiца)
	LangEN               = 3   // English
	LangES               = 4   // Español
	LangFI               = 5   // Suomi
	LangDE               = 6   // Deutsch
	LangIT               = 7   // Italiano
	LangBG               = 8   // Български
	LangHR               = 9   // Hrvatski
	LangHU               = 10  // Magyar
	LangSR               = 11  // Српски
	LangPT               = 12  // Português
	LangEL               = 14  // Ελληνικά
	LangPL               = 15  // Polski
	LangFR               = 16  // Français
	LangKO               = 17  // 한국어
	LangZH               = 18  // 汉语
	LangLT               = 19  // Lietuvių
	LangJA               = 20  // 日本語
	LangCS               = 21  // Čeština
	LangET               = 22  // Eesti
	LangTT               = 50  // Татарча
	LangBA               = 51  // Башҡортса
	LangCV               = 52  // Чăвашла
	LangSK               = 53  // Slovenčina
	LangRO               = 54  // Română
	LangNO               = 55  // Norsk
	LangLV               = 56  // Latviešu
	LangAZ               = 57  // Azərbaycan dili
	LangHY               = 58  // Հայերեն
	LangSQ               = 59  // Shqip
	LangSV               = 60  // Svenska
	LangNL               = 61  // Nederlands
	LangTK               = 62  // Türkmen
	LangKA               = 63  // ქართული
	LangDA               = 64  // Dansk
	LangUZ               = 65  // O‘zbek
	LangMO               = 66  // Moldovenească
	LangBUA              = 67  // Буряад
	LangTH               = 68  // ภาษาไทย
	LangID               = 69  // Bahasa Indonesia
	LangTG               = 70  // Тоҷикӣ
	LangSL               = 71  // Slovenščina
	LangBS               = 72  // Bosanski
	LangPTBR             = 73  // Português brasileiro
	LangFA               = 74  // فارسی
	LangVI               = 75  // Tiếng Việt
	LangHI               = 76  // हिन्दी
	LangSI               = 77  // සිංහල
	LangBN               = 78  // বাংলা
	LangTL               = 79  // Tagalog
	LangMN               = 80  // Монгол
	LangMY               = 81  // ဗမာစာ
	LangTR               = 82  // Türkçe
	LangNE               = 83  // नेपाली
	LangUR               = 85  // اردو
	LangKY               = 87  // Кыргыз тили
	LangPA               = 90  // پنجابی
	LangOS               = 91  // Ирон
	LangKN               = 94  // ಕನ್ನಡ
	LangSW               = 95  // Kiswahili
	LangKK               = 97  // Қазақша
	LangAR               = 98  // العربية
	LangHE               = 99  // עברית
	LangPreRevolutionary = 100 // Дореволюцiонный
	LangMYV              = 101 // Эрзянь кель
	LangKDB              = 102 // Адыгэбзэ
	LangSAH              = 105 // Саха тыла
	LangADY              = 106 // Адыгабзэ
	LangUDM              = 107 // Удмурт
	LangCHM              = 108 // Марий йылме
	LangBE2              = 114 // Беларуская
	LangLEZ              = 118 // Лезги чІал
	LangTW               = 119 // 臺灣話
	LangKUM              = 236 // Къумукъ тил
	LangMVL              = 270 // Mirandés
	LangSLA              = 298 // Русинськый
	LangKRL              = 379 // Karjalan kieli
	LangTYV              = 344 // Тыва дыл
	LangXAL              = 357 // Хальмг келн
	LangTLY              = 373 // Tolışə zıvon
	LangKV               = 375 // Коми кыв
	LangUKClassic        = 452 // Українська (клясична)
	LangUKGalitska       = 454 // Українська (Галицка)
	LangKAB              = 457 // Taqbaylit
	LangEO               = 555 // Esperanto
	LangLA               = 666 // Lingua Latina
	LangSoviet           = 777 // Советский
)

// Button action type
const (
	ButtonText     = "text"
	ButtonVKPay    = "vkpay"
	ButtonVKApp    = "open_app"
	ButtonLocation = "location"
	ButtonOpenLink = "open_link"
)
