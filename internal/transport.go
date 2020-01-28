package internal // import "github.com/SevereCloud/vksdk/internal"

import (
	"context"
	"net/http"
)

// HTTPClient is the context key to use with context's
// WithValue function to associate an *http.Client value with a context.
var HTTPClient ContextKey // nolint:gochecknoglobals

// ContextKey is just an empty struct. It exists so HTTPClient can be
// an immutable public variable with a unique type. It's immutable
// because nobody else can create a ContextKey, being unexported.
type ContextKey struct{}

func ContextClient(ctx context.Context) *http.Client {
	if ctx != nil {
		if hc, ok := ctx.Value(HTTPClient).(*http.Client); ok {
			return hc
		}
	}

	return http.DefaultClient
}
