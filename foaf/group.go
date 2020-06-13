package foaf // import "github.com/SevereCloud/vksdk/foaf"

import (
	"context"
	"net/http"
	"strconv"
)

// GroupType type.
type GroupType string

// Types of GroupType.
const (
	GroupTypeGroup  GroupType = "group"
	GroupTypePublic GroupType = "public"
	GroupTypeEvent  GroupType = "event"
)

// Group - A class of Agents.
//
// http://xmlns.com/foaf/spec/#term_Group
type Group struct {
	GroupType    GroupType `xml:"groupType"`
	Name         string    `xml:"name"`
	URI          []URI     `xml:"URI"`
	Img          Img       `xml:"img"`
	Weblog       Weblog    `xml:"weblog"`
	MembersCount int       `xml:"membersCount"`
	StartDate    Date      `xml:"startDate"`
	FinishDate   Date      `xml:"finishDate"`
}

// GetGroup return Group.
func GetGroup(ctx context.Context, groupID int) (Group, error) {
	req, _ := http.NewRequest("GET", FOAFURL, nil)
	q := req.URL.Query()
	q.Add("id", strconv.Itoa(-groupID))
	req.URL.RawQuery = q.Encode()

	r, err := getFoaf(ctx, req)

	return r.Group, err
}
