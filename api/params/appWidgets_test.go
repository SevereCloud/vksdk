package params_test

import (
	"testing"

	"github.com/Derad6709/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestAppWidgetsUpdateBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppWidgetsUpdateBuilder()

	b.Code("text")
	b.Type("text")

	assert.Equal(t, b.Params["code"], "text")
	assert.Equal(t, b.Params["type"], "text")
}
