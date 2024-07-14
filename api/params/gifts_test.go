package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/stretchr/testify/assert"
)

func TestGiftsGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGiftsGetBuilder()

	b.UserID(1)
	b.Count(1)
	b.Offset(1)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["offset"])
}
