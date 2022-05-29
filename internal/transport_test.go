package internal_test

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Derad6709/vksdk/v2/internal"
	"github.com/stretchr/testify/assert"
)

func TestContextUserAgent(t *testing.T) {
	t.Parallel()

	f := func(ctx context.Context, wantUserAgent string) {
		userAgent := internal.ContextUserAgent(ctx)
		assert.Equal(t, wantUserAgent, userAgent)
	}

	f(context.Background(), internal.UserAgent)

	ctx := context.Background()
	userAgent := "agent"
	ctx = context.WithValue(ctx, internal.UserAgentKey, userAgent)
	f(ctx, userAgent)
}

func TestContextClient(t *testing.T) {
	t.Parallel()

	f := func(ctx context.Context, wantClient *http.Client) {
		client := internal.ContextClient(ctx)
		assert.Equal(t, wantClient, client)
	}

	f(context.Background(), http.DefaultClient)

	ctx := context.Background()
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, internal.HTTPClientKey, httpClient)
	f(ctx, httpClient)
}

func TestGo17Context(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ok")
	}))
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatalf("unexpected error in http.NewRequests: %v", err)
	}

	ctx := context.Background()

	resp, err := internal.DoRequest(ctx, req)
	if resp == nil || err != nil {
		t.Fatalf("error received from client: %v %v", err, resp)
	}

	resp.Body.Close()
}

const (
	requestDuration = 100 * time.Millisecond
	requestBody     = "ok"
)

func okHandler(w http.ResponseWriter, _ *http.Request) {
	time.Sleep(requestDuration)
	io.WriteString(w, requestBody)
}

func TestNoTimeout(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(okHandler))
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatalf("unexpected error in http.NewRequests: %v", err)
	}

	ctx := context.Background()

	resp, err := internal.DoRequest(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	slurp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(slurp) != requestBody {
		t.Errorf("body = %q; want %q", slurp, requestBody)
	}
}

func TestCancelBeforeHeaders(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())

	blockServer := make(chan struct{})
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cancel()
		<-blockServer

		io.WriteString(w, requestBody)
	}))

	defer ts.Close()
	defer close(blockServer)

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatalf("unexpected error in http.NewRequests: %v", err)
	}

	resp, err := internal.DoRequest(ctx, req)
	if err == nil {
		resp.Body.Close()

		t.Fatal("Get returned unexpected nil error")
	}

	if !errors.Is(err, context.Canceled) {
		t.Errorf("err = %v; want %v", err, context.Canceled)
	}
}
