/*
Package internal unimportable
*/
package internal // import "github.com/SevereCloud/vksdk/internal"

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

// CharsetReader if non-nil, defines a function to generate
// charset-conversion readers, converting from the provided
// non-UTF-8 charset into UTF-8. If CharsetReader is nil or
// returns an error, parsing stops with an error. One of the
// the CharsetReader's result values must be non-nil.
func CharsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch strings.ToLower(charset) {
	case "windows-1251":
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	default:
		return nil, fmt.Errorf("unknown charset: %s", charset)
	}
}
