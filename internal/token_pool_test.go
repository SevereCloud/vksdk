package internal

import (
	"testing"
)

func TestTokenPool(t *testing.T) {
	t.Parallel()

	tokens := []string{
		"1", "2", "3", "4", "5",
	}

	p := NewTokenPool(tokens...)

	for i := 0; i < len(tokens)*2; i++ {
		tokenIdx := i % len(tokens)

		expect := tokens[tokenIdx]
		actual := p.Get()

		if actual != expect {
			t.Errorf("%s and %s should be equal", actual, expect)
		}
	}
}
