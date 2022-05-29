package params_test

import (
	"testing"

	"github.com/Derad6709/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestGiftsGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewGiftsGetBuilder()

	b.UserID(1)
	b.Count(1)
	b.Offset(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
}
