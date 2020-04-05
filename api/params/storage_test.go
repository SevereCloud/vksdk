package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestStorageGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStorageGetBuilder()

	b.Key("text")
	b.Keys([]string{"text"})
	b.UserID(1)
	b.Global(true)

	assert.Equal(t, b.Params["key"], "text")
	assert.Equal(t, b.Params["keys"], []string{"text"})
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["global"], true)
}

func TestStorageGetKeysBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStorageGetKeysBuilder()

	b.UserID(1)
	b.Global(true)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["global"], true)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestStorageSetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStorageSetBuilder()

	b.Key("text")
	b.Value("text")
	b.UserID(1)
	b.Global(true)

	assert.Equal(t, b.Params["key"], "text")
	assert.Equal(t, b.Params["value"], "text")
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["global"], true)
}
