/*
Package internal unimportable
*/
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

// ContextClient return *http.Client
func ContextClient(ctx context.Context) *http.Client {
	if ctx != nil {
		if hc, ok := ctx.Value(HTTPClient).(*http.Client); ok {
			return hc
		}
	}

	return http.DefaultClient
}

// DoRequest sends an HTTP request and returns an HTTP response.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func DoRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	client := ContextClient(ctx)
	resp, err := client.Do(req.WithContext(ctx))
	// If we got an error, and the context has been canceled,
	// the context's error is probably more useful.
	if err != nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}
	}

	return resp, err
}
