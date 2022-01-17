package object_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/vmihailenco/msgpack/v5/msgpcode"
)

func TestBaseBoolInt_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(data []byte, want object.BaseBoolInt, wantErr string) {
		t.Helper()

		var b object.BaseBoolInt

		err := b.UnmarshalJSON(data)
		if err != nil || wantErr != "" {
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

func TestBaseBoolInt_DecodeMsgpack(t *testing.T) {
	t.Parallel()

	f := func(data []byte, want object.BaseBoolInt, wantErr string) {
		t.Helper()

		var b object.BaseBoolInt

		dec := msgpack.NewDecoder(bytes.NewReader(data))

		err := b.DecodeMsgpack(dec)
		if err != nil || wantErr != "" {
			assert.EqualError(t, err, wantErr)
		}

		assert.Equal(t, want, b)
	}

	ff := func(c string) string {
		return "json: cannot unmarshal " + c + " into Go value of type *object.BaseBoolInt"
	}

	// fixint
	f([]byte("\x00"), false, "")
	f([]byte("\x01"), true, "")

	// uint 8
	f([]byte("\xcc\x00"), false, "")
	f([]byte("\xcc\x01"), true, "")
	f([]byte("\xcc\x02"), false, ff("\xcc\x02"))

	// uint 16
	f([]byte("\xcd\x00\x00"), false, "")
	f([]byte("\xcd\x00\x01"), true, "")
	f([]byte("\xcd\x00\x02"), false, ff("\xcd\x00\x02"))

	// uint 32
	f([]byte("\xce\x00\x00\x00\x00"), false, "")
	f([]byte("\xce\x00\x00\x00\x01"), true, "")
	f([]byte("\xce\x00\x00\x00\x02"), false, ff("\xce\x00\x00\x00\x02"))

	// uint 64
	f([]byte("\xcf\x00\x00\x00\x00\x00\x00\x00\x00"), false, "")
	f([]byte("\xcf\x00\x00\x00\x00\x00\x00\x00\x01"), true, "")
	f([]byte("\xcf\x00\x00\x00\x00\x00\x00\x00\x02"), false, ff("\xcf\x00\x00\x00\x00\x00\x00\x00\x02"))

	// int 8
	f([]byte("\xd0\x00"), false, "")
	f([]byte("\xd0\x01"), true, "")
	f([]byte("\xd0\x02"), false, ff("\xd0\x02"))

	// int 16
	f([]byte("\xd1\x00\x00"), false, "")
	f([]byte("\xd1\x00\x01"), true, "")
	f([]byte("\xd1\x00\x02"), false, ff("\xd1\x00\x02"))

	// int 32
	f([]byte("\xd2\x00\x00\x00\x00"), false, "")
	f([]byte("\xd2\x00\x00\x00\x01"), true, "")
	f([]byte("\xd2\x00\x00\x00\x02"), false, ff("\xd2\x00\x00\x00\x02"))

	// int 64
	f([]byte("\xd3\x00\x00\x00\x00\x00\x00\x00\x00"), false, "")
	f([]byte("\xd3\x00\x00\x00\x00\x00\x00\x00\x01"), true, "")
	f([]byte("\xd3\x00\x00\x00\x00\x00\x00\x00\x02"), false, ff("\xd3\x00\x00\x00\x00\x00\x00\x00\x02"))

	// bool test
	f([]byte{msgpcode.True}, true, "")
	f([]byte{msgpcode.False}, false, "")

	// nil test
	f([]byte{msgpcode.FixedArrayLow}, false, ff("\x90"))

	f(nil, false, "EOF")
}

func TestBaseImage_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(data []byte, want object.BaseImage, wantErr string) {
		t.Helper()

		var b object.BaseImage

		err := b.UnmarshalJSON(data)
		if err != nil || wantErr != "" {
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
	f([]byte("0"), object.BaseImage{}, "json: cannot unmarshal number into Go value of type object.renamedBaseImage")
}

func TestBaseImage_DecodeMsgpack(t *testing.T) {
	t.Parallel()

	f := func(data []byte, want object.BaseImage, wantErr string) {
		t.Helper()

		var b object.BaseImage

		dec := msgpack.NewDecoder(bytes.NewReader(data))

		err := b.DecodeMsgpack(dec)
		if err != nil || wantErr != "" {
			assert.EqualError(t, err, wantErr)
		}

		assert.Equal(t, want, b)
	}

	f([]byte("\x80"), object.BaseImage{}, "")
	f([]byte{
		0x84, 0xA3, 0x75, 0x72, 0x6C, 0xD9, 0x38, 0x68, 0x74, 0x74, 0x70, 0x73,
		0x3A, 0x2F, 0x2F, 0x70, 0x70, 0x2E, 0x76, 0x6B, 0x2E, 0x6D, 0x65, 0x2F,
		0x63, 0x36, 0x33, 0x33, 0x38, 0x32, 0x35, 0x2F, 0x76, 0x36, 0x33, 0x33,
		0x38, 0x32, 0x35, 0x30, 0x33, 0x34, 0x2F, 0x37, 0x33, 0x36, 0x39, 0x2F,
		0x77, 0x62, 0x73, 0x41, 0x73, 0x72, 0x6F, 0x6F, 0x71, 0x66, 0x41, 0x2E,
		0x6A, 0x70, 0x67, 0xA5, 0x77, 0x69, 0x64, 0x74, 0x68, 0xCC, 0x82, 0xA6,
		0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x57, 0xA4, 0x74, 0x79, 0x70, 0x65,
		0xA1, 0x6D,
	},
		object.BaseImage{
			URL:    "https://pp.vk.me/c633825/v633825034/7369/wbsAsrooqfA.jpg",
			Width:  130,
			Height: 87,
			Type:   "m",
		},
		"",
	)
	f([]byte{
		0x84, 0xA3, 0x73, 0x72, 0x63, 0xD9, 0x38, 0x68, 0x74, 0x74, 0x70, 0x73,
		0x3A, 0x2F, 0x2F, 0x70, 0x70, 0x2E, 0x76, 0x6B, 0x2E, 0x6D, 0x65, 0x2F,
		0x63, 0x36, 0x33, 0x33, 0x38, 0x32, 0x35, 0x2F, 0x76, 0x36, 0x33, 0x33,
		0x38, 0x32, 0x35, 0x30, 0x33, 0x34, 0x2F, 0x37, 0x33, 0x36, 0x39, 0x2F,
		0x77, 0x62, 0x73, 0x41, 0x73, 0x72, 0x6F, 0x6F, 0x71, 0x66, 0x41, 0x2E,
		0x6A, 0x70, 0x67, 0xA5, 0x77, 0x69, 0x64, 0x74, 0x68, 0xCC, 0x82, 0xA6,
		0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x57, 0xA4, 0x74, 0x79, 0x70, 0x65,
		0xA1, 0x6D,
	},
		object.BaseImage{
			URL:    "https://pp.vk.me/c633825/v633825034/7369/wbsAsrooqfA.jpg",
			Width:  130,
			Height: 87,
			Type:   "m",
		},
		"",
	)
	f([]byte("xc0"), object.BaseImage{}, "msgpack: unexpected code=78 decoding map length")
	f(nil, object.BaseImage{}, "EOF")
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
