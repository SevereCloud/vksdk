package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestMarusiaSavePictureBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarusiaSavePictureBuilder()

	b.Server(1)
	b.Photo("text")
	b.Hash("text")

	assert.Equal(t, b.Params["server"], 1)
	assert.Equal(t, b.Params["photo"], "text")
	assert.Equal(t, b.Params["hash"], "text")
}

func TestMarusiaDeletePictureBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarusiaDeletePictureBuilder()

	b.ID(1)

	assert.Equal(t, b.Params["id"], 1)
}
