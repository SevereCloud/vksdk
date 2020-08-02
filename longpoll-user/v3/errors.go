package wrapper // import "github.com/SevereCloud/vksdk/longpoll-user/v3"

import "fmt"

// tooShortArray type.
type tooShortArray struct {
	structName string
	least      int
	got        int
}

// Error returns the message of a tooShortArray.
func (e tooShortArray) Error() string {
	return fmt.Sprintf(
		"wrapper: cannot parse array into Go struct %s: expected at least %d element(s), got %d",
		e.structName,
		e.least,
		e.got,
	)
}

// expectedSlice type.
type expectedSlice struct {
	v interface{}
}

// Error returns the message of a expectedSlice.
func (e expectedSlice) Error() string {
	return fmt.Sprintf(
		"wrapper: expected a slice, got %T",
		e.v,
	)
}

// failedCast type.
type failedCast struct {
	v interface{}
}

// Error returns the message of a failedCast.
func (e failedCast) Error() string {
	return fmt.Sprintf(
		"wrapper: cast failed, value type: %T",
		e.v,
	)
}
