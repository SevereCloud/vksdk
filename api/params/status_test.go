package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/stretchr/testify/assert"
)

func TestStatusGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStatusGetBuilder()

	b.UserID(1)
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestStatusSetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStatusSetBuilder()

	b.Text("text")
	b.GroupID(1)

	assert.Equal(t, "text", b.Params["text"])
	assert.Equal(t, 1, b.Params["group_id"])
}
