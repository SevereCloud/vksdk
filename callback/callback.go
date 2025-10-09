/*
Package callback implements Callback API.

See more https://dev.vk.com/ru/api/callback/getting-started
*/
package callback // import "github.com/SevereCloud/vksdk/v3/callback"

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/SevereCloud/vksdk/v3/events"
	"github.com/SevereCloud/vksdk/v3/internal"
)

// Callback struct SecretKeys [GroupID]SecretKey.
type Callback struct {
	*events.FuncList

	ConfirmationKeys map[int]string
	ConfirmationKey  string
	SecretKeys       map[int]string
	SecretKey        string
	Title            string

	// ErrorLog specifies a logger for errors unexpected
	// behavior from handlers.
	// By default, [slog.Default] is used.
	ErrorLog *slog.Logger
}

// NewCallback return *Callback.
func NewCallback() *Callback {
	cb := &Callback{
		Title:            "vksdk",
		ConfirmationKeys: make(map[int]string),
		SecretKeys:       make(map[int]string),
		ErrorLog:         slog.Default(),
		FuncList:         events.NewFuncList(),
	}

	return cb
}

func (cb *Callback) confirmationKey(groupID int) string {
	if v := cb.ConfirmationKeys[groupID]; v != "" {
		return v
	}

	return cb.ConfirmationKey
}

var groupEventPool = sync.Pool{ //nolint:gochecknoglobals
	New: func() any {
		return &events.GroupEvent{Object: make(json.RawMessage, 0, 1024)}
	},
}

// HandleFunc handler.
func (cb *Callback) HandleFunc(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	e := groupEventPool.Get().(*events.GroupEvent)

	defer func() {
		*e = events.GroupEvent{Object: e.Object[:0]}
		groupEventPool.Put(e)
	}()

	err := json.NewDecoder(r.Body).Decode(e)
	if err != nil {
		cb.ErrorLog.ErrorContext(ctx, "failed to json decode request body", "error", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)

		return
	}

	secretKey, ok := cb.SecretKeys[e.GroupID]
	if !ok {
		secretKey = cb.SecretKey
	}

	if secretKey != "" && e.Secret != secretKey {
		cb.ErrorLog.WarnContext(ctx, "wrong secret key", "group_id", e.GroupID)
		http.Error(w, "Bad Secret", http.StatusForbidden)

		return
	}

	if e.Type == events.EventConfirmation {
		_, err := fmt.Fprint(w, cb.confirmationKey(e.GroupID))
		if err != nil {
			cb.ErrorLog.ErrorContext(ctx, "failed to write confirmation key", "error", err)
		}

		return
	}

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

	err = cb.Handler(ctx, *e)
	if err != nil {
		cb.ErrorLog.ErrorContext(ctx, "failed to handle event", "error", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)

		return
	}

	if remove {
		cb.response(ctx, w, "remove")

		return
	}

	if code != 0 {
		w.Header().Set("Retry-After", date.Format(http.TimeFormat)) // RFC 7231, 7.1.3
		http.Error(w, http.StatusText(code), code)

		return
	}

	cb.response(ctx, w, "ok")
}

func (cb *Callback) response(ctx context.Context, w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	_, err := w.Write([]byte(data))
	if err != nil {
		slog.ErrorContext(ctx, "failed to write http response",
			slog.Any("error", err),
			slog.String("data", data),
		)
	}
}
