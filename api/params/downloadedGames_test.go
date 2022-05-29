package params_test

import (
	"testing"

	"github.com/Derad6709/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestDownloadedGamesGetPaidStatusBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDownloadedGamesGetPaidStatusBuilder()

	b.UserID(1)

	assert.Equal(t, b.Params["user_id"], 1)
}
