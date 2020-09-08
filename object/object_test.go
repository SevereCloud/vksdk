package object_test

import (
	"reflect"
	"testing"

	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/stretchr/testify/assert"
)

func TestBaseBoolInt_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(data []byte, want object.BaseBoolInt, wantErr string) {
		t.Helper()

		var b object.BaseBoolInt

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
	f([]byte("null"), false, "json: cannot unmarshal null into Go value of type *object.BaseBoolInt")
}

func TestBaseImage_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(data []byte, want object.BaseImage, wantErr string) {
		t.Helper()

		var b object.BaseImage

		err := b.UnmarshalJSON(data)
		if err != nil {
			assert.EqualError(t, err, wantErr)
		}

		assert.Equal(t, want, b)
	}

	f([]byte("{}"), object.BaseImage{}, "")
	f([]byte(`{
		"src": "https://pp.vk.me/c633825/v633825034/7369/wbsAsrooqfA.jpg",
		"width": 130,
		"height": 87,
		"type": "m"
		}`),
		object.BaseImage{
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
		object.BaseImage{
			URL:    "https://pp.vk.me/c633825/v633825034/7369/wbsAsrooqfA.jpg",
			Width:  130,
			Height: 87,
			Type:   "m",
		},
		"",
	)
	f([]byte("null"), object.BaseImage{}, "json: cannot unmarshal ? into Go value of type BaseImage")
}

func TestBaseSticker_MaxSize(t *testing.T) {
	t.Parallel()

	f := func(photo object.BaseSticker, want object.BaseImage) {
		t.Helper()

		if got := photo.MaxSize(); !reflect.DeepEqual(got, want) {
			t.Errorf("BaseSticker.MaxSize() = %v, want %v", got, want)
		}
	}

	f(object.BaseSticker{}, object.BaseImage{})
	f(object.BaseSticker{
		Images: []object.BaseImage{
			{
				Width: 10, Height: 20,
			},
			{
				Width: 100, Height: 200,
			},
		},
	}, object.BaseImage{
		Width: 100, Height: 200,
	})
}

func TestBaseSticker_MinSize(t *testing.T) {
	t.Parallel()

	f := func(photo object.BaseSticker, want object.BaseImage) {
		t.Helper()

		if got := photo.MinSize(); !reflect.DeepEqual(got, want) {
			t.Errorf("BaseSticker.MinSize() = %v, want %v", got, want)
		}
	}

	f(object.BaseSticker{}, object.BaseImage{})
	f(object.BaseSticker{
		Images: []object.BaseImage{
			{
				Width: 10, Height: 20,
			},
			{
				Width: 100, Height: 200,
			},
		},
	}, object.BaseImage{
		Width: 10, Height: 20,
	})
}

func TestBaseSticker_MaxSizeBackground(t *testing.T) {
	t.Parallel()

	f := func(photo object.BaseSticker, want object.BaseImage) {
		t.Helper()

		if got := photo.MaxSizeBackground(); !reflect.DeepEqual(got, want) {
			t.Errorf("BaseSticker.MaxSizeBackground() = %v, want %v", got, want)
		}
	}

	f(object.BaseSticker{}, object.BaseImage{})
	f(object.BaseSticker{
		ImagesWithBackground: []object.BaseImage{
			{
				Width: 10, Height: 20,
			},
			{
				Width: 100, Height: 200,
			},
		},
	}, object.BaseImage{
		Width: 100, Height: 200,
	})
}

func TestBaseSticker_MinSizeBackground(t *testing.T) {
	t.Parallel()

	f := func(photo object.BaseSticker, want object.BaseImage) {
		t.Helper()

		if got := photo.MinSizeBackground(); !reflect.DeepEqual(got, want) {
			t.Errorf("BaseSticker.MinSizeBackground() = %v, want %v", got, want)
		}
	}

	f(object.BaseSticker{}, object.BaseImage{})
	f(object.BaseSticker{
		ImagesWithBackground: []object.BaseImage{
			{
				Width: 10, Height: 20,
			},
			{
				Width: 100, Height: 200,
			},
		},
	}, object.BaseImage{
		Width: 10, Height: 20,
	})
}
