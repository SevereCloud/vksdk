package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestCallsStartBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewCallsStartBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
}

func TestCallsForceFinishBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewCallsForceFinishBuilder()

	b.CallID("text")

	assert.Equal(t, "text", b.Params["call_id"])
}
