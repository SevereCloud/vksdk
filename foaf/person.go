package foaf // import "github.com/Derad6709/vksdk/v2/foaf"

import (
	"context"
	"net/http"
	"strconv"
)

// Access type.
type Access string

// Types of Access.
const (
	AccessAllowed    Access = "allowed"
	AccessDisallowed Access = "disallowed"
)

// ProfileState is profile state.
type ProfileState string

// Types of ProfileState.
const (
	ProfileStateDeleted     ProfileState = "deleted"
	ProfileStateVerified    ProfileState = "verified"
	ProfileStateActive      ProfileState = "active"
	ProfileStateBanned      ProfileState = "banned"
	ProfileStateDeactivated ProfileState = "deactivated"
)

// Gender - the gender of this Agent (typically but not necessarily 'male' or 'female').
type Gender string

// Types of Gender.
const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

// Person - A person.
//
// http://xmlns.com/foaf/spec/#term_Person
type Person struct {
	PublicAccess      Access       `xml:"publicAccess"`
	ProfileState      ProfileState `xml:"profileState"`
	URI               []URI        `xml:"URI"`
	FirstName         string       `xml:"firstName"`
	SecondName        string       `xml:"secondName"`
	Name              string       `xml:"name"`
	Weblog            Weblog       `xml:"weblog"`
	Gender            Gender       `xml:"gender"`
	Created           Date         `xml:"created"`
	LastLoggedIn      Date         `xml:"lastLoggedIn"`
	Modified          Date         `xml:"modified"`
	SubscribersCount  int          `xml:"subscribersCount"`
	FriendsCount      int          `xml:"friendsCount"`
	SubscribedToCount int          `xml:"subscribedToCount"`
	Birthday          string       `xml:"birthday"`
	DateOfBirth       string       `xml:"dateOfBirth"`
	Img               Img          `xml:"img"`
	FamilyStatus      string       `xml:"familyStatus"`
	SkypeID           string       `xml:"skypeID"` // A Skype ID
	Homepage          string       `xml:"homepage"`
	Phone             []Phone      `xml:"phone"`
	ExternalProfile   []External   `xml:"externalProfile"`
	Interest          string       `xml:"interest"`
	LocationOfBirth   Location     `xml:"locationOfBirth"`
	Location          Location     `xml:"location"`
	Edu               []Edu        `xml:"edu"`
	Job               []Job        `xml:"job"`
}

// External struct.
type External struct {
	Status   string `xml:"status,attr"`
	Resource string `xml:"resource,attr"`
}

// Phone struct.
type Phone struct {
	Primary string `xml:"primary,attr"`
}

// Edu struct.
type Edu struct {
	School     School     `xml:"school"`
	University University `xml:"university"`
}

// School struct.
type School struct {
	Title      string   `xml:"title,attr"`
	DateStart  string   `xml:"dateStart,attr"`
	DateFinish string   `xml:"dateFinish,attr"`
	About      string   `xml:"about,attr"`
	Location   Location `xml:"location"`
}

// University struct.
type University struct {
	Title        string   `xml:"title,attr"`
	ShortCaption string   `xml:"shortCaption,attr"`
	DateFinish   string   `xml:"dateFinish,attr"`
	About        string   `xml:"about,attr"`
	Location     Location `xml:"location"`
}

// Job struct.
type Job struct {
	WorkPlace WorkPlace `xml:"workPlace"`
	Military  Military  `xml:"military"`
}

// Military struct.
type Military struct {
	Title     string   `xml:"title,attr"`
	DateStart string   `xml:"dateStart,attr"`
	About     string   `xml:"about,attr"`
	Location  Location `xml:"location"`
}

// WorkPlace struct.
type WorkPlace struct {
	Title     string   `xml:"title,attr"`
	About     string   `xml:"about,attr"`
	DateStart string   `xml:"dateStart,attr"`
	Position  string   `xml:"position,attr"`
	Location  Location `xml:"location"`
}

// Location struct.
type Location struct {
	Country string `xml:"country,attr"`
	City    string `xml:"city,attr"`
}

// GetPerson return Person.
func GetPerson(ctx context.Context, userID int) (Person, error) {
	req, _ := http.NewRequest("GET", BaseURL, nil)
	q := req.URL.Query()
	q.Add("id", strconv.Itoa(userID))
	req.URL.RawQuery = q.Encode()

	r, err := getFoaf(ctx, req)

	return r.Person, err
}
