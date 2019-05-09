package object

// SearchHint struct
type SearchHint struct {

	// Object description
	Description string `json:"description"`

	// Information whether the object has been found globally
	Global  int          `json:"global,omitempty"`
	Group   *GroupsGroup `json:"group,omitempty"`
	Profile *UsersUser   `json:"profile,omitempty"`
	Section string       `json:"section"`
	Type    string       `json:"type"`
}
