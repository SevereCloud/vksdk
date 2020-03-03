package internal_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/internal"
)

func TestCharsetReader(t *testing.T) {
	f := func(charset string, wantErr bool) {
		_, err := internal.CharsetReader(charset, nil)
		if (err != nil) != wantErr {
			t.Errorf("CharsetReader() error = %v, wantErr %v", err, wantErr)
			return
		}
	}

	f("", true)
	f("windows-1251", false)
	f("WINDOWS-1251", false)
}
