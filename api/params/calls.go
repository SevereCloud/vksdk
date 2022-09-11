package params

import "github.com/SevereCloud/vksdk/v2/api"

// CallsStartBuilder builder.
//
// https://vk.com/dev/calls.start
type CallsStartBuilder struct {
	api.Params
}

// NewCallsStartBuilder func.
func NewCallsStartBuilder() *CallsStartBuilder {
	return &CallsStartBuilder{api.Params{}}
}

// GroupID parameter.
func (b *CallsStartBuilder) GroupID(v int) *CallsStartBuilder {
	b.Params["group_id"] = v
	return b
}

// CallsForceFinishBuilder builder.
//
// https://vk.com/dev/calls.forceFinish
type CallsForceFinishBuilder struct {
	api.Params
}

// NewCallsForceFinishBuilder func.
func NewCallsForceFinishBuilder() *CallsForceFinishBuilder {
	return &CallsForceFinishBuilder{api.Params{}}
}

// CallID parameter.
func (b *CallsForceFinishBuilder) CallID(v string) *CallsForceFinishBuilder {
	b.Params["call_id"] = v
	return b
}
