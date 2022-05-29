/*
Package callback implements Callback API.

See more https://vk.com/dev/callback_api
*/
package callback // import "github.com/Derad6709/vksdk/v2/callback"

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Derad6709/vksdk/v2/events"
	"github.com/Derad6709/vksdk/v2/internal"
)

// Callback struct SecretKeys [GroupID]SecretKey.
type Callback struct {
	ConfirmationKeys map[int]string
	ConfirmationKey  string
	SecretKeys       map[int]string
	SecretKey        string
	Title            string

	// ErrorLog specifies an optional logger for errors unexpected
	// behavior from handlers.
	// If nil, logging is done via the log package's standard logger.
	ErrorLog *log.Logger

	events.FuncList
}

// NewCallback return *Callback.
func NewCallback() *Callback {
	cb := &Callback{
		Title:            "vksdk",
		ConfirmationKeys: make(map[int]string),
		SecretKeys:       make(map[int]string),
		FuncList:         *events.NewFuncList(),
	}

	return cb
}

// HandleFunc handler.
func (cb *Callback) HandleFunc(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var e events.GroupEvent
	if err := decoder.Decode(&e); err != nil {
		cb.logf("callback: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)

		return
	}

	secretKey, ok := cb.SecretKeys[e.GroupID]
	if !ok {
		secretKey = cb.SecretKey
	}

	if secretKey != "" && e.Secret != secretKey {
		cb.logf("callback: bad secret %d", e.GroupID)
		http.Error(w, "Bad Secret", http.StatusForbidden)

		return
	}

	if e.Type == events.EventConfirmation {
		if cb.ConfirmationKeys[e.GroupID] != "" {
			fmt.Fprintf(w, cb.ConfirmationKeys[e.GroupID])
		} else {
			fmt.Fprintf(w, cb.ConfirmationKey)
		}

		return
	}

	ctx := context.Background()

	retryCounter, _ := strconv.Atoi(r.Header.Get("X-Retry-Counter"))
	ctx = context.WithValue(ctx, internal.CallbackRetryCounterKey, retryCounter)

	var (
		code   int
		date   time.Time
		remove bool
	)

	retryAfter := func(c int, d time.Time) {
		code = c
		date = d
	}
	ctx = context.WithValue(ctx, internal.CallbackRetryAfterKey, retryAfter)

	removeFunc := func() {
		remove = true
	}
	ctx = context.WithValue(ctx, internal.CallbackRemove, removeFunc)

	if err := cb.Handler(ctx, e); err != nil {
		cb.logf("callback: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)

		return
	}

	if remove {
		_, _ = w.Write([]byte("remove"))

		return
	}

	if code != 0 {
		w.Header().Set("Retry-After", date.Format(http.TimeFormat)) // RFC 7231, 7.1.3
		http.Error(w, http.StatusText(code), code)

		return
	}

	_, _ = w.Write([]byte("ok"))
}

func (cb *Callback) logf(format string, args ...interface{}) {
	if cb.ErrorLog != nil {
		cb.ErrorLog.Printf(format, args...)
	} else {
		log.Printf(format, args...)
	}
}
