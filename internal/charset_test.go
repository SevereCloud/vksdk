package internal_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/internal"
	"github.com/stretchr/testify/assert"
)

func TestCharsetReader(t *testing.T) {
	t.Parallel()

	f := func(charset string, wantErr string) {
		_, err := internal.CharsetReader(charset, nil)
		if err != nil || wantErr != "" {
			assert.EqualError(t, err, wantErr)
		}
	}

	f("test", "unknown charset: test")
	f("windows-1251", "")
	f("WINDOWS-1251", "")
}
