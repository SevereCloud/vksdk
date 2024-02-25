package object_test

import (
	"encoding/json"
	"testing"

	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
)

func TestRawMessage_MarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(msg object.RawMessage, want []byte, errString string) {
		t.Helper()

		b, err := json.Marshal(msg)
		if err != nil || errString != "" {
			assert.EqualError(t, err, errString)
		}

		assert.Equal(t, want, b)
	}

	f(object.RawMessage("1"), []byte("1"), "")
	f(object.RawMessage(nil), []byte("null"), "")
}

func TestRawMessage_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(b []byte, want object.RawMessage, errString string) {
		t.Helper()

		var v object.RawMessage

		err := json.Unmarshal(b, &v)
		if err != nil || errString != "" {
			assert.EqualError(t, err, errString)
		}

		assert.Equal(t, want, v)
	}

	f([]byte("0"), object.RawMessage("0"), "")
	f([]byte("null"), object.RawMessage("null"), "")
	f(nil, object.RawMessage(nil), "unexpected end of JSON input")
}

func TestRawMessage_EncodeMsgpack(t *testing.T) {
	t.Parallel()

	f := func(msg object.RawMessage, want []byte, errString string) {
		t.Helper()

		b, err := msgpack.Marshal(msg)
		if err != nil || errString != "" {
			assert.EqualError(t, err, errString)
		}

		assert.Equal(t, want, b)
	}

	f(object.RawMessage{0x00}, []byte{0x00}, "")
	f(object.RawMessage(nil), []byte{0xc0}, "")
}

func TestRawMessage_DecodeMsgpack(t *testing.T) {
	t.Parallel()

	f := func(b []byte, want object.RawMessage, errString string) {
		t.Helper()

		var v object.RawMessage

		err := msgpack.Unmarshal(b, &v)
		if err != nil || errString != "" {
			assert.EqualError(t, err, errString)
		}

		assert.Equal(t, want, v)
	}

	f([]byte{0x00}, object.RawMessage{0x00}, "")
	f([]byte{0xc0}, object.RawMessage(nil), "")
	f(nil, object.RawMessage(nil), "object.RawMessage: EOF")
}
