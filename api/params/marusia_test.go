package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/stretchr/testify/assert"
)

func TestMarusiaSavePictureBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarusiaSavePictureBuilder()

	b.Server(1)
	b.Photo("text")
	b.Hash("text")

	assert.Equal(t, 1, b.Params["server"])
	assert.Equal(t, "text", b.Params["photo"])
	assert.Equal(t, "text", b.Params["hash"])
}

func TestMarusiaDeletePictureBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarusiaDeletePictureBuilder()

	b.ID(1)

	assert.Equal(t, 1, b.Params["id"])
}

func TestMarusiaCreateAudioBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarusiaCreateAudioBuilder()

	b.AudioMeta("text")

	assert.Equal(t, "text", b.Params["audio_meta"])
}

func TestMarusiaDeleteAudioBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarusiaDeleteAudioBuilder()

	b.ID(1)

	assert.Equal(t, 1, b.Params["id"])
}
