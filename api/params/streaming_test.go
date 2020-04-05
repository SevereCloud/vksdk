package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestStreamingSetSettingsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStreamingSetSettingsBuilder()

	b.MonthlyTier("text")

	assert.Equal(t, b.Params["monthly_tier"], "text")
}
