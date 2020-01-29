package internal_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/SevereCloud/vksdk/internal"
	"github.com/stretchr/testify/assert"
)

func TestContextClient(t *testing.T) {
	f := func(ctx context.Context, wantClient *http.Client) {
		client := internal.ContextClient(ctx)
		assert.Equal(t, wantClient, client)
	}

	f(context.Background(), http.DefaultClient)

	ctx := context.Background()
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, internal.HTTPClient, httpClient)
	f(ctx, httpClient)
}
