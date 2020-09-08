package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// StatusGetBuilder builder.
//
// Returns data required to show the status of a user or community.
//
// https://vk.com/dev/status.get
type StatusGetBuilder struct {
	api.Params
}

// NewStatusGetBuilder func.
func NewStatusGetBuilder() *StatusGetBuilder {
	return &StatusGetBuilder{api.Params{}}
}

// UserID user ID or community ID. Use a negative value to designate a community ID.
func (b *StatusGetBuilder) UserID(v int) *StatusGetBuilder {
	b.Params["user_id"] = v
	return b
}

// GroupID parameter.
func (b *StatusGetBuilder) GroupID(v int) *StatusGetBuilder {
	b.Params["group_id"] = v
	return b
}

// StatusSetBuilder builder.
//
// Sets a new status for the current user.
//
// https://vk.com/dev/status.set
type StatusSetBuilder struct {
	api.Params
}

// NewStatusSetBuilder func.
func NewStatusSetBuilder() *StatusSetBuilder {
	return &StatusSetBuilder{api.Params{}}
}

// Text of the new status.
func (b *StatusSetBuilder) Text(v string) *StatusSetBuilder {
	b.Params["text"] = v
	return b
}

// GroupID identifier of a community to set a status in. If left blank the status is set to current user.
func (b *StatusSetBuilder) GroupID(v int) *StatusSetBuilder {
	b.Params["group_id"] = v
	return b
}
