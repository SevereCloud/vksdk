package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestStorageGetBulder(t *testing.T) {
	b := params.NewStorageGetBulder()

	b.Key("text")
	b.Keys([]string{"text"})
	b.UserID(1)
	b.Global(true)

	assert.Equal(t, b.Params["key"], "text")
	assert.Equal(t, b.Params["keys"], []string{"text"})
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["global"], true)
}

func TestStorageGetKeysBulder(t *testing.T) {
	b := params.NewStorageGetKeysBulder()

	b.UserID(1)
	b.Global(true)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["global"], true)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestStorageSetBulder(t *testing.T) {
	b := params.NewStorageSetBulder()

	b.Key("text")
	b.Value("text")
	b.UserID(1)
	b.Global(true)

	assert.Equal(t, b.Params["key"], "text")
	assert.Equal(t, b.Params["value"], "text")
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["global"], true)
}
