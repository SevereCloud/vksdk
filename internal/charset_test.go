package internal_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/SevereCloud/vksdk/v3/internal"
	"github.com/stretchr/testify/assert"
)

func TestXMLSanitizerReader(t *testing.T) {
	t.Parallel()

	r := bytes.NewBuffer([]byte{0x09, 0x0B})
	newR := internal.XMLSanitizerReader(r)

	b, err := io.ReadAll(newR)
	assert.NoError(t, err)
	assert.Equal(t, []byte{0x09}, b)
}

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
