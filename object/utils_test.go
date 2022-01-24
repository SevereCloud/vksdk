package object_test

import (
	"encoding/json"
	"testing"

	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/vmihailenco/msgpack/v5/msgpcode"
)

func TestUtilsDomainResolved_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(data []byte, want object.UtilsDomainResolved) {
		t.Helper()

		var resolved object.UtilsDomainResolved

		err := json.Unmarshal(data, &resolved)
		assert.NoError(t, err)
		assert.Equal(t, want, resolved)
	}

	f([]byte("[]"), object.UtilsDomainResolved{})
	f([]byte(`{"object_id":1}`), object.UtilsDomainResolved{ObjectID: 1})

	var resolved object.UtilsDomainResolved
	err := json.Unmarshal([]byte("0"), &resolved)
	assert.Error(t, err)
}

func TestUtilsDomainResolved_DecodeMsgpack(t *testing.T) {
	t.Parallel()

	f := func(data []byte, want object.UtilsDomainResolved, wantErr string) {
		t.Helper()

		var resolved object.UtilsDomainResolved

		err := msgpack.Unmarshal(data, &resolved)
		if err != nil || wantErr != "" {
			assert.EqualError(t, err, wantErr)
		}

		assert.Equal(t, want, resolved)
	}

	f([]byte{msgpcode.FixedArrayLow}, object.UtilsDomainResolved{}, "")
	f([]byte{
		0x81, 0xA9, 'o', 'b', 'j', 'e', 'c', 't', '_', 'i', 'd', 0x01,
	}, object.UtilsDomainResolved{ObjectID: 1}, "")
	f([]byte("\xc2"), object.UtilsDomainResolved{}, "msgpack: unexpected code=c2 decoding map length")
	f(nil, object.UtilsDomainResolved{}, "EOF")
}
