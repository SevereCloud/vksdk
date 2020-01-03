package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestGiftsGetBulder(t *testing.T) {
	b := params.NewGiftsGetBulder()

	b.UserID(1)
	b.Count(1)
	b.Offset(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
}
