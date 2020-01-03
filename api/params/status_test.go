package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestStatusGetBulder(t *testing.T) {
	b := params.NewStatusGetBulder()

	b.UserID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestStatusSetBulder(t *testing.T) {
	b := params.NewStatusSetBulder()

	b.Text("text")
	b.GroupID(1)

	assert.Equal(t, b.Params["text"], "text")
	assert.Equal(t, b.Params["group_id"], 1)
}
