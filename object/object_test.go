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
