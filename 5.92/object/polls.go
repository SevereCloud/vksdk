package object // import "github.com/SevereCloud/vksdk/5.92/object"

import (
	"fmt"
)

type pollsAnswer struct {
	ID    int     `json:"id"`
	Rate  float64 `json:"rate"`
	Text  string  `json:"text"`
	Votes int     `json:"votes"`
}

// PollsPoll struct
type PollsPoll struct {
	AnswerID   int             `json:"answer_id"` // Current user's answer ID
	Answers    []pollsAnswer   `json:"answers"`
	Created    int             `json:"created"`  // Date when poll has been created in Unixtime
	ID         int             `json:"id"`       // Poll ID
	OwnerID    int             `json:"owner_id"` // Poll owner's ID
	Question   string          `json:"question"` // Poll question
	Votes      int             `json:"votes"`    // Votes number
	AnswerIDs  []int           `json:"answer_ids"`
	EndDate    int             `json:"end_date"`
	Anonymous  bool            `json:"anonymous"` // Information whether the pole is anonymous
	Closed     bool            `json:"closed"`
	IsBoard    bool            `json:"is_board"`
	CanEdit    bool            `json:"can_edit"`
	CanVote    bool            `json:"can_vote"`
	CanReport  bool            `json:"can_report"`
	CanShare   bool            `json:"can_share"`
	Multiple   bool            `json:"multiple"`
	Photo      PhotosPhoto     `json:"photo"`
	AuthorID   int             `json:"author_id"`
	Background PollsBackground `json:"background"`
	Friends    []pollsFriend   `json:"friends"`
	Profiles   []UsersUser     `json:"profiles"`
	Groups     []GroupsGroup   `json:"groups"`
}

func (poll PollsPoll) ToAttachment() string {
	return fmt.Sprintf("poll%d_%d", poll.OwnerID, poll.ID)
}

type pollsFriend struct {
	ID int `json:"id"`
}

// PollsVoters struct
type PollsVoters struct {
	AnswerID int              `json:"answer_id"` // Answer ID
	Users    pollsVotersUsers `json:"users"`
}

type pollsVotersUsers struct {
	Count int   `json:"count"` // Votes number
	Items []int `json:"items"`
}

// PollsVotersFields struct
type PollsVotersFields struct {
	AnswerID int                    `json:"answer_id"` // Answer ID
	Users    pollsVotersUsersFields `json:"users"`
}

type pollsVotersUsersFields struct {
	Count int         `json:"count"` // Votes number
	Items []UsersUser `json:"items"`
}

// PollsBackground struct
type PollsBackground struct {
	Type   string `json:"type"`
	Angle  int    `json:"angle"`
	Color  string `json:"color"`
	Points []struct {
		Position float64 `json:"position"`
		Color    string  `json:"color"`
	} `json:"points"`
	ID int `json:"id"`
}

// PollsPhoto struct
type PollsPhoto struct {
	ID     int           `json:"id"`
	Color  string        `json:"color"`
	Images []PhotosImage `json:"images"`
}
