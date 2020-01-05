package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// StatusGetBuilder builder
//
// Returns data required to show the status of a user or community.
//
// https://vk.com/dev/status.get
type StatusGetBuilder struct {
	api.Params
}

// NewStatusGetBuilder func
func NewStatusGetBuilder() *StatusGetBuilder {
	return &StatusGetBuilder{api.Params{}}
}

// UserID User ID or community ID. Use a negative value to designate a community ID.
func (b *StatusGetBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupID parameter
func (b *StatusGetBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// StatusSetBuilder builder
//
// Sets a new status for the current user.
//
// https://vk.com/dev/status.set
type StatusSetBuilder struct {
	api.Params
}

// NewStatusSetBuilder func
func NewStatusSetBuilder() *StatusSetBuilder {
	return &StatusSetBuilder{api.Params{}}
}

// Text Text of the new status.
func (b *StatusSetBuilder) Text(v string) {
	b.Params["text"] = v
}

// GroupID Identifier of a community to set a status in. If left blank the status is set to current user.
func (b *StatusSetBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}
