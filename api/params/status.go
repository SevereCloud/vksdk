package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// StatusGetBulder builder
//
// Returns data required to show the status of a user or community.
//
// https://vk.com/dev/status.get
type StatusGetBulder struct {
	api.Params
}

// NewStatusGetBulder func
func NewStatusGetBulder() *StatusGetBulder {
	return &StatusGetBulder{api.Params{}}
}

// UserID User ID or community ID. Use a negative value to designate a community ID.
func (b *StatusGetBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupID parameter
func (b *StatusGetBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// StatusSetBulder builder
//
// Sets a new status for the current user.
//
// https://vk.com/dev/status.set
type StatusSetBulder struct {
	api.Params
}

// NewStatusSetBulder func
func NewStatusSetBulder() *StatusSetBulder {
	return &StatusSetBulder{api.Params{}}
}

// Text Text of the new status.
func (b *StatusSetBulder) Text(v string) {
	b.Params["text"] = v
}

// GroupID Identifier of a community to set a status in. If left blank the status is set to current user.
func (b *StatusSetBulder) GroupID(v int) {
	b.Params["group_id"] = v
}
