package params

import "github.com/SevereCloud/vksdk/v2/api"

// DonutGetFriendsBuilder struct.
type DonutGetFriendsBuilder struct {
	api.Params
}

// NewDonutGetFriendsBuilder func.
func NewDonutGetFriendsBuilder() *DonutGetFriendsBuilder {
	return &DonutGetFriendsBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *DonutGetFriendsBuilder) OwnerID(v int) *DonutGetFriendsBuilder {
	b.Params["owner_id"] = v
	return b
}

// Offset parameter.
func (b *DonutGetFriendsBuilder) Offset(v int) *DonutGetFriendsBuilder {
	b.Params["offset"] = v
	return b
}

// Count - number of results to return.
func (b *DonutGetFriendsBuilder) Count(v int) *DonutGetFriendsBuilder {
	b.Params["count"] = v
	return b
}

// Fields parameter.
func (b *DonutGetFriendsBuilder) Fields(v []string) *DonutGetFriendsBuilder {
	b.Params["fields"] = v
	return b
}

// DonutGetSubscriptionBuilder struct.
type DonutGetSubscriptionBuilder struct {
	api.Params
}

// NewDonutGetSubscriptionBuilder func.
func NewDonutGetSubscriptionBuilder() *DonutGetSubscriptionBuilder {
	return &DonutGetSubscriptionBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *DonutGetSubscriptionBuilder) OwnerID(v int) *DonutGetSubscriptionBuilder {
	b.Params["owner_id"] = v
	return b
}

// DonutGetSubscriptionsBuilder struct.
type DonutGetSubscriptionsBuilder struct {
	api.Params
}

// NewDonutGetSubscriptionsBuilder func.
func NewDonutGetSubscriptionsBuilder() *DonutGetSubscriptionsBuilder {
	return &DonutGetSubscriptionsBuilder{api.Params{}}
}

// Fields parameter.
func (b *DonutGetSubscriptionsBuilder) Fields(v []string) *DonutGetSubscriptionsBuilder {
	b.Params["fields"] = v
	return b
}

// Offset parameter.
func (b *DonutGetSubscriptionsBuilder) Offset(v int) *DonutGetSubscriptionsBuilder {
	b.Params["offset"] = v
	return b
}

// Count - number of results to return.
func (b *DonutGetSubscriptionsBuilder) Count(v int) *DonutGetSubscriptionsBuilder {
	b.Params["count"] = v
	return b
}

// DonutIsDonBuilder struct.
type DonutIsDonBuilder struct {
	api.Params
}

// NewDonutIsDonBuilder func.
func NewDonutIsDonBuilder() *DonutIsDonBuilder {
	return &DonutIsDonBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *DonutIsDonBuilder) OwnerID(v int) *DonutIsDonBuilder {
	b.Params["owner_id"] = v
	return b
}
