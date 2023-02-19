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

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestCallsForceFinishBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewCallsForceFinishBuilder()

	b.CallID("text")

	assert.Equal(t, b.Params["call_id"], "text")
}
