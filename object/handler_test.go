package object_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/object"
	"github.com/stretchr/testify/assert"
)

func TestMessageNewObject_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(b []byte, want object.MessageNewObject, wantErr bool) {
		var obj object.MessageNewObject

		err := obj.UnmarshalJSON(b)
		if (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}

		assert.Equal(t, obj, want)
	}

	f(
		[]byte(""),
		object.MessageNewObject{},
		true,
	)
	f(
		[]byte(`{"id":1}`),
		object.MessageNewObject{Message: object.MessagesMessage{ID: 1}},
		false,
	)
	f(
		[]byte(`{"message":{"id":1},"client_info":{}}`),
		object.MessageNewObject{Message: object.MessagesMessage{ID: 1}},
		false,
	)
}
