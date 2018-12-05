package object

// UsersUser struct
type UsersUser struct {
	ID                     int               `json:"id"`
	FirstName              string            `json:"first_name"`
	LastName               string            `json:"last_name"`
	IsClosed               bool              `json:"is_closed"`
	CanAccessClosed        bool              `json:"can_access_closed"`
	Sex                    int               `json:"sex"`
	Nickname               string            `json:"nickname"`
	Domain                 string            `json:"domain"`
	ScreenName             string            `json:"screen_name"`
	Bdate                  string            `json:"bdate"`
	City                   BaseObject        `json:"city"`
	Country                BaseObject        `json:"country"`
	Photo50                string            `json:"photo_50"`
	Photo100               string            `json:"photo_100"`
	Photo200               string            `json:"photo_200"`
	PhotoMax               string            `json:"photo_max"`
	Photo200Orig           string            `json:"photo_200_orig"`
	Photo400Orig           string            `json:"photo_400_orig"`
	PhotoMaxOrig           string            `json:"photo_max_orig"`
	PhotoID                string            `json:"photo_id"`
	HasPhoto               int               `json:"has_photo"`
	HasMobile              int               `json:"has_mobile"`
	IsFriend               int               `json:"is_friend"`
	FriendStatus           int               `json:"friend_status"`
	Online                 int               `json:"online"`
	CanPost                int               `json:"can_post"`
	CanSeeAllPosts         int               `json:"can_see_all_posts"`
	CanSeeAudio            int               `json:"can_see_audio"`
	CanWritePrivateMessage int               `json:"can_write_private_message"`
	CanSendFriendRequest   int               `json:"can_send_friend_request"`
	Facebook               string            `json:"facebook"`
	FacebookName           string            `json:"facebook_name"`
	Twitter                string            `json:"twitter"`
	Instagram              string            `json:"instagram"`
	Site                   string            `json:"site"`
	Status                 string            `json:"status"`
	StatusAudio            AudioAudioFull    `json:"status_audio"`
	LastSeen               usersLastSeen     `json:"last_seen"`
	CropPhoto              usersCropPhoto    `json:"crop_photo"`
	Verified               int               `json:"verified"`
	FollowersCount         int               `json:"followers_count"`
	Blacklisted            int               `json:"blacklisted"`
	BlacklistedByMe        int               `json:"blacklisted_by_me"`
	IsFavorite             int               `json:"is_favorite"`
	IsHiddenFromFeed       int               `json:"is_hidden_from_feed"`
	CommonCount            int               `json:"common_count"`
	Occupation             usersOccupation   `json:"occupation"`
	Career                 usersCareer       `json:"career"`
	Military               usersMilitary     `json:"military"`
	University             int               `json:"university"`
	UniversityName         string            `json:"university_name"`
	Faculty                int               `json:"faculty"`
	FacultyName            string            `json:"faculty_name"`
	Graduation             int               `json:"graduation"`
	HomeTown               string            `json:"home_town"`
	Relation               int               `json:"relation"`
	Personal               usersPersonal     `json:"personal"`
	Interests              string            `json:"interests"`
	Music                  string            `json:"music"`
	Activities             string            `json:"activities"`
	Movies                 string            `json:"movies"`
	Tv                     string            `json:"tv"`
	Books                  string            `json:"books"`
	Games                  string            `json:"games"`
	Universities           []usersUniversity `json:"universities"`
	Schools                []usersSchool     `json:"schools"`
	About                  string            `json:"about"`
	Relatives              []usersRelative   `json:"relatives"`
	Quotes                 string            `json:"quotes"`
	Lists                  []int             `json:"lists"`
	Deactivated            string            `json:"deactivated"`
	WallDefault            string            `json:"wall_default"`
	Trending               int               `json:"trending"`
	Timezone               int               `json:"timezone"`
	MaidenName             string            `json:"maiden_name"`
	Exports                usersExports      `json:"exports"`
	Counters               usersUserCounters `json:"counters"`
	Contacts               usersContacts     `json:"contacts"`
	// TODO: education
}

// UsersUserMin struct
type UsersUserMin struct {
	Deactivated string `json:"deactivated"`
	FirstName   string `json:"first_name"`
	Hidden      int    `json:"hidden"`
	ID          int    `json:"id"`
	LastName    string `json:"last_name"`
}

type usersContacts struct {
	MobilePhone string `json:"mobile_phone"`
	HomePhone   string `json:"home_phone"`
}

type usersCareer struct {
	CityID    int    `json:"city_id"`
	Company   string `json:"company"`
	CountryID int    `json:"country_id"`
	From      int    `json:"from"`
	GroupID   int    `json:"group_id"`
	Position  string `json:"position"`
	Until     int    `json:"until"`
}

type usersCropPhoto struct {
	Crop  usersCropPhotoCrop `json:"crop"`
	Photo PhotosPhoto        `json:"photo"`
	Rect  usersCropPhotoRect `json:"rect"`
}

type usersCropPhotoCrop struct {
	X  float64 `json:"x"`
	X2 float64 `json:"x2"`
	Y  float64 `json:"y"`
	Y2 float64 `json:"y2"`
}

type usersCropPhotoRect struct {
	X  float64 `json:"x"`
	X2 float64 `json:"x2"`
	Y  float64 `json:"y"`
	Y2 float64 `json:"y2"`
}

type usersExports struct {
	Facebook    int `json:"facebook"`
	Livejournal int `json:"livejournal"`
	Twitter     int `json:"twitter"`
}

type usersLastSeen struct {
	Platform int `json:"platform"`
	Time     int `json:"time"`
}

type usersMilitary struct {
	CountryID int    `json:"country_id"`
	From      int    `json:"from"`
	Unit      string `json:"unit"`
	UnitID    int    `json:"unit_id"`
	Until     int    `json:"until"`
}

type usersOccupation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type usersPersonal struct {
	Alcohol    int      `json:"alcohol"`
	InspiredBy string   `json:"inspired_by"`
	Langs      []string `json:"langs"`
	LifeMain   int      `json:"life_main"`
	PeopleMain int      `json:"people_main"`
	Political  int      `json:"political"`
	Religion   string   `json:"religion"`
	Smoking    int      `json:"smoking"`
}

type usersRelative struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type usersSchool struct {
	City          int    `json:"city"`
	Class         string `json:"class"`
	Country       int    `json:"country"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	Type          int    `json:"type"`
	TypeStr       string `json:"type_str"`
	YearFrom      int    `json:"year_from"`
	YearGraduated int    `json:"year_graduated"`
	YearTo        int    `json:"year_to"`
}

type usersUniversity struct {
	Chair           int    `json:"chair"`
	ChairName       string `json:"chair_name"`
	City            int    `json:"city"`
	Country         int    `json:"country"`
	EducationForm   string `json:"education_form"`
	EducationStatus string `json:"education_status"`
	Faculty         int    `json:"faculty"`
	FacultyName     string `json:"faculty_name"`
	Graduation      int    `json:"graduation"`
	ID              int    `json:"id"`
	Name            string `json:"name"`
}

type usersUserCounters struct {
	Albums        int `json:"albums"`
	Audios        int `json:"audios"`
	Followers     int `json:"followers"`
	Friends       int `json:"friends"`
	Gifts         int `json:"gifts"`
	Groups        int `json:"groups"`
	Notes         int `json:"notes"`
	OnlineFriends int `json:"online_friends"`
	Pages         int `json:"pages"`
	Photos        int `json:"photos"`
	Subscriptions int `json:"subscriptions"`
	UserPhotos    int `json:"user_photos"`
	UserVideos    int `json:"user_videos"`
	Videos        int `json:"videos"`
}

type usersUserLim struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	NameGen string `json:"name_gen"`
	Photo   string `json:"photo"`
}
