package longpoll

import (
	"errors"
	"fmt"
)

// ErrNotValidVersion not valid version.
var ErrNotValidVersion = errors.New("longpoll: not valid version")

// Failed struct.
type Failed struct {
	Code int
}

// Error returns the message of a Failed.
func (e Failed) Error() string {
	return fmt.Sprintf(
		"longpoll: failed code %d",
		e.Code,
	)
}
