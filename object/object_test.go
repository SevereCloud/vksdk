/*
Package object contains objects for VK.

See more https://vk.com/dev/objects
*/
package object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseBoolInt_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(data []byte, want BaseBoolInt, wantErr string) {
		t.Helper()

		var b BaseBoolInt

		err := b.UnmarshalJSON(data)
		if err != nil {
			assert.EqualError(t, err, wantErr)
		}

		assert.Equal(t, want, b)
	}

	f([]byte("1"), true, "")
	f([]byte("true"), true, "")
	f([]byte("0"), false, "")
	f([]byte("false"), false, "")
	f([]byte("null"), false, "json: cannot unmarshal ? into Go value of type BaseBoolInt")
}

func TestBaseImage_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(data []byte, want BaseImage, wantErr string) {
		t.Helper()

		var b BaseImage

		err := b.UnmarshalJSON(data)
		if err != nil {
			assert.EqualError(t, err, wantErr)
		}

		assert.Equal(t, want, b)
	}

	f([]byte("{}"), BaseImage{}, "")
	f([]byte(`{
		"src": "https://pp.vk.me/c633825/v633825034/7369/wbsAsrooqfA.jpg",
		"width": 130,
		"height": 87,
		"type": "m"
		}`),
		BaseImage{
			URL:    "https://pp.vk.me/c633825/v633825034/7369/wbsAsrooqfA.jpg",
			Width:  130,
			Height: 87,
			Type:   "m",
		},
		"",
	)
	f([]byte(`{
		"url": "https://pp.vk.me/c633825/v633825034/7369/wbsAsrooqfA.jpg",
		"width": 130,
		"height": 87,
		"type": "m"
		}`),
		BaseImage{
			URL:    "https://pp.vk.me/c633825/v633825034/7369/wbsAsrooqfA.jpg",
			Width:  130,
			Height: 87,
			Type:   "m",
		},
		"",
	)
	f([]byte("null"), BaseImage{}, "json: cannot unmarshal ? into Go value of type BaseImage")
}
