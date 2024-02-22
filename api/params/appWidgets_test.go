package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestAppWidgetsUpdateBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppWidgetsUpdateBuilder()

	b.Code("text")
	b.Type("text")

	assert.Equal(t, "text", b.Params["code"])
	assert.Equal(t, "text", b.Params["type"])
}
