package object // import "github.com/SevereCloud/vksdk/object"

import (
	"fmt"
)

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
	OnlineApp              int               `json:"online_app"`
	OnlineMobile           int               `json:"online_mobile"`
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
	Career                 []usersCareer     `json:"career"`
	Military               []usersMilitary   `json:"military"`
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
	FoundWith              int               `json:"found_with"` // TODO: check it
	OnlineInfo             UsersOnlineInfo   `json:"online_info"`
	// TODO: education
}

// ToMention return mention
func (user UsersUser) ToMention() string {
	return fmt.Sprintf("[id%d|%s %s]", user.ID, user.FirstName, user.LastName)
}

// UsersOnlineInfo struct
type UsersOnlineInfo struct {
	AppID    int  `json:"app_id"`
	Visible  bool `json:"visible"`
	IsOnline bool `json:"is_online"`
	IsMobile bool `json:"is_mobile"`
}

// UsersUserMin struct
type UsersUserMin struct {
	Deactivated string `json:"deactivated"` // Returns if a profile is deleted or blocked
	FirstName   string `json:"first_name"`  // User first name
	Hidden      int    `json:"hidden"`      // Returns if a profile is hidden.
	ID          int    `json:"id"`          // User ID
	LastName    string `json:"last_name"`   // User last name
}

// ToMention return mention
func (user UsersUserMin) ToMention() string {
	return fmt.Sprintf("[id%d|%s %s]", user.ID, user.FirstName, user.LastName)
}

type usersContacts struct {
	MobilePhone string `json:"mobile_phone"`
	HomePhone   string `json:"home_phone"`
}

type usersCareer struct {
	CityID    int    `json:"city_id"`    // City ID
	Company   string `json:"company"`    // Company name
	CountryID int    `json:"country_id"` // Country ID
	From      int    `json:"from"`       // From year
	GroupID   int    `json:"group_id"`   // Community ID
	ID        int    `json:"id"`         // Career ID
	Position  string `json:"position"`   // Position
	Until     int    `json:"until"`      // Till year
}

type usersCropPhoto struct {
	Crop  usersCropPhotoCrop `json:"crop"`
	Photo PhotosPhoto        `json:"photo"`
	Rect  usersCropPhotoRect `json:"rect"`
}

type usersCropPhotoCrop struct {
	X  float64 `json:"x"`  // Coordinate X of the left upper corner
	X2 float64 `json:"x2"` // Coordinate X of the right lower corner
	Y  float64 `json:"y"`  // Coordinate Y of the left upper corner
	Y2 float64 `json:"y2"` // Coordinate Y of the right lower corner
}

type usersCropPhotoRect struct {
	X  float64 `json:"x"`  // Coordinate X of the left upper corner
	X2 float64 `json:"x2"` // Coordinate X of the right lower corner
	Y  float64 `json:"y"`  // Coordinate Y of the left upper corner
	Y2 float64 `json:"y2"` // Coordinate Y of the right lower corner
}

type usersExports struct {
	Facebook    int `json:"facebook"`
	Livejournal int `json:"livejournal"`
	Twitter     int `json:"twitter"`
}

type usersLastSeen struct {
	Platform int `json:"platform"` // Type of the platform that used for the last authorization
	Time     int `json:"time"`     // Last visit date (in Unix time)
}

type usersMilitary struct {
	CountryID int    `json:"country_id"` // Country ID
	From      int    `json:"from"`       // From year
	ID        int    `json:"id"`         // Military ID
	Unit      string `json:"unit"`       // Unit name
	UnitID    int    `json:"unit_id"`    // Unit ID
	Until     int    `json:"until"`      // Till year
}

type usersOccupation struct {
	// BUG(VK): https://vk.com/bug136108
	ID   float64 `json:"id"`   // ID of school, university, company group
	Name string  `json:"name"` // Name of occupation
	Type string  `json:"type"` // Type of occupation
}

type usersPersonal struct {
	Alcohol    int      `json:"alcohol"`     // User's views on alcohol
	InspiredBy string   `json:"inspired_by"` // User's inspired by
	Langs      []string `json:"langs"`
	LifeMain   int      `json:"life_main"`   // User's personal priority in life
	PeopleMain int      `json:"people_main"` // User's personal priority in people
	Political  int      `json:"political"`   // User's political views
	Religion   string   `json:"religion"`    // User's religion
	Smoking    int      `json:"smoking"`     // User's views on smoking
}

type usersRelative struct {
	BirthDate string `json:"birth_date"` // Date of child birthday (format dd.mm.yyyy)
	ID        int    `json:"id"`         // Relative ID
	Name      string `json:"name"`       // Name of relative
	Type      string `json:"type"`       // Relative type
}

type usersSchool struct {
	City          int    `json:"city"`           // City ID
	Class         string `json:"class"`          // School class letter
	Country       int    `json:"country"`        // Country ID
	ID            string `json:"id"`             // School ID
	Name          string `json:"name"`           // School name
	Type          int    `json:"type"`           // School type ID
	TypeStr       string `json:"type_str"`       // School type name
	YearFrom      int    `json:"year_from"`      // Year the user started to study
	YearGraduated int    `json:"year_graduated"` // Graduation year
	YearTo        int    `json:"year_to"`        // Year the user finished to study
}

type usersUniversity struct {
	Chair           int    `json:"chair"`            // Chair ID
	ChairName       string `json:"chair_name"`       // Chair name
	City            int    `json:"city"`             // City ID
	Country         int    `json:"country"`          // Country ID
	EducationForm   string `json:"education_form"`   // Education form
	EducationStatus string `json:"education_status"` // Education status
	Faculty         int    `json:"faculty"`          // Faculty ID
	FacultyName     string `json:"faculty_name"`     // Faculty name
	Graduation      int    `json:"graduation"`       // Graduation year
	ID              int    `json:"id"`               // University ID
	Name            string `json:"name"`             // University name
}

type usersUserCounters struct {
	Albums        int `json:"albums"`         // Albums number
	Audios        int `json:"audios"`         // Audios number
	Followers     int `json:"followers"`      // Followers number
	Friends       int `json:"friends"`        // Friends number
	Gifts         int `json:"gifts"`          // Gifts number
	Groups        int `json:"groups"`         // Communities number
	Notes         int `json:"notes"`          // Notes number
	OnlineFriends int `json:"online_friends"` // Online friends number
	Pages         int `json:"pages"`          // Public pages number
	Photos        int `json:"photos"`         // Photos number
	Subscriptions int `json:"subscriptions"`  // Subscriptions number
	UserPhotos    int `json:"user_photos"`    // Number of photos with user
	UserVideos    int `json:"user_videos"`    // Number of videos with user
	Videos        int `json:"videos"`         // Videos number
}

// UsersUserLim struct
type UsersUserLim struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	NameGen string `json:"name_gen"`
	Photo   string `json:"photo"`
}
