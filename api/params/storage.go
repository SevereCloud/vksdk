package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// StorageGetBuilder builder.
//
// Returns a value of variable with the name set by key parameter.
//
// https://vk.com/dev/storage.get
type StorageGetBuilder struct {
	api.Params
}

// NewStorageGetBuilder func.
func NewStorageGetBuilder() *StorageGetBuilder {
	return &StorageGetBuilder{api.Params{}}
}

// Key parameter.
func (b *StorageGetBuilder) Key(v string) *StorageGetBuilder {
	b.Params["key"] = v
	return b
}

// Keys parameter.
func (b *StorageGetBuilder) Keys(v []string) *StorageGetBuilder {
	b.Params["keys"] = v
	return b
}

// UserID parameter.
func (b *StorageGetBuilder) UserID(v int) *StorageGetBuilder {
	b.Params["user_id"] = v
	return b
}

// Global parameter.
func (b *StorageGetBuilder) Global(v bool) *StorageGetBuilder {
	b.Params["global"] = v
	return b
}

// StorageGetKeysBuilder builder.
//
// Returns the names of all variables.
//
// https://vk.com/dev/storage.getKeys
type StorageGetKeysBuilder struct {
	api.Params
}

// NewStorageGetKeysBuilder func.
func NewStorageGetKeysBuilder() *StorageGetKeysBuilder {
	return &StorageGetKeysBuilder{api.Params{}}
}

// UserID user id, whose variables names are returned if they were requested with a server method.
func (b *StorageGetKeysBuilder) UserID(v int) *StorageGetKeysBuilder {
	b.Params["user_id"] = v
	return b
}

// Global parameter.
func (b *StorageGetKeysBuilder) Global(v bool) *StorageGetKeysBuilder {
	b.Params["global"] = v
	return b
}

// Offset parameter.
func (b *StorageGetKeysBuilder) Offset(v int) *StorageGetKeysBuilder {
	b.Params["offset"] = v
	return b
}

// Count amount of variable names the info needs to be collected from.
func (b *StorageGetKeysBuilder) Count(v int) *StorageGetKeysBuilder {
	b.Params["count"] = v
	return b
}

// StorageSetBuilder builder.
//
// Saves a value of variable with the name set by 'key' parameter.
//
// https://vk.com/dev/storage.set
type StorageSetBuilder struct {
	api.Params
}

// NewStorageSetBuilder func.
func NewStorageSetBuilder() *StorageSetBuilder {
	return &StorageSetBuilder{api.Params{}}
}

// Key parameter.
func (b *StorageSetBuilder) Key(v string) *StorageSetBuilder {
	b.Params["key"] = v
	return b
}

// Value parameter.
func (b *StorageSetBuilder) Value(v string) *StorageSetBuilder {
	b.Params["value"] = v
	return b
}

// UserID parameter.
func (b *StorageSetBuilder) UserID(v int) *StorageSetBuilder {
	b.Params["user_id"] = v
	return b
}

// Global parameter.
func (b *StorageSetBuilder) Global(v bool) *StorageSetBuilder {
	b.Params["global"] = v
	return b
}
