package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// StorageGetBulder builder
//
// Returns a value of variable with the name set by key parameter.
//
// https://vk.com/dev/storage.get
type StorageGetBulder struct {
	api.Params
}

// NewStorageGetBulder func
func NewStorageGetBulder() *StorageGetBulder {
	return &StorageGetBulder{api.Params{}}
}

// Key parameter
func (b *StorageGetBulder) Key(v string) {
	b.Params["key"] = v
}

// Keys parameter
func (b *StorageGetBulder) Keys(v []string) {
	b.Params["keys"] = v
}

// UserID parameter
func (b *StorageGetBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Global parameter
func (b *StorageGetBulder) Global(v bool) {
	b.Params["global"] = v
}

// StorageGetKeysBulder builder
//
// Returns the names of all variables.
//
// https://vk.com/dev/storage.getKeys
type StorageGetKeysBulder struct {
	api.Params
}

// NewStorageGetKeysBulder func
func NewStorageGetKeysBulder() *StorageGetKeysBulder {
	return &StorageGetKeysBulder{api.Params{}}
}

// UserID user id, whose variables names are returned if they were requested with a server method.
func (b *StorageGetKeysBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Global parameter
func (b *StorageGetKeysBulder) Global(v bool) {
	b.Params["global"] = v
}

// Offset parameter
func (b *StorageGetKeysBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count amount of variable names the info needs to be collected from.
func (b *StorageGetKeysBulder) Count(v int) {
	b.Params["count"] = v
}

// StorageSetBulder builder
//
// Saves a value of variable with the name set by 'key' parameter.
//
// https://vk.com/dev/storage.set
type StorageSetBulder struct {
	api.Params
}

// NewStorageSetBulder func
func NewStorageSetBulder() *StorageSetBulder {
	return &StorageSetBulder{api.Params{}}
}

// Key parameter
func (b *StorageSetBulder) Key(v string) {
	b.Params["key"] = v
}

// Value parameter
func (b *StorageSetBulder) Value(v string) {
	b.Params["value"] = v
}

// UserID parameter
func (b *StorageSetBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Global parameter
func (b *StorageSetBulder) Global(v bool) {
	b.Params["global"] = v
}
