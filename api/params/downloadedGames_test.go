package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/stretchr/testify/assert"
)

func TestDownloadedGamesGetPaidStatusBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDownloadedGamesGetPaidStatusBuilder()

	b.UserID(1)

	assert.Equal(t, 1, b.Params["user_id"])
}
