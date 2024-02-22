package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestStorageGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStorageGetBuilder()

	b.Key("text")
	b.Keys([]string{"text"})
	b.UserID(1)
	b.Global(true)

	assert.Equal(t, "text", b.Params["key"])
	assert.Equal(t, []string{"text"}, b.Params["keys"])
	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, true, b.Params["global"])
}

func TestStorageGetKeysBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStorageGetKeysBuilder()

	b.UserID(1)
	b.Global(true)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, true, b.Params["global"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestStorageSetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStorageSetBuilder()

	b.Key("text")
	b.Value("text")
	b.UserID(1)
	b.Global(true)

	assert.Equal(t, "text", b.Params["key"])
	assert.Equal(t, "text", b.Params["value"])
	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, true, b.Params["global"])
}
