/*
Package callback implements Callback API.

See more https://vk.com/dev/callback_api
*/
package callback // import "github.com/SevereCloud/vksdk/v2/callback"

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SevereCloud/vksdk/v2/events"
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

	if cb.SecretKeys[e.GroupID] != "" || cb.SecretKey != "" {
		if e.Secret != cb.SecretKeys[e.GroupID] && e.Secret != cb.SecretKey {
			cb.logf("callback: bad secret %d", e.GroupID)
			http.Error(w, "Bad Secret", http.StatusForbidden)

			return
		}
	}

	if e.Type == events.EventConfirmation {
		if cb.ConfirmationKeys[e.GroupID] != "" {
			fmt.Fprintf(w, cb.ConfirmationKeys[e.GroupID])
		} else {
			fmt.Fprintf(w, cb.ConfirmationKey)
		}

		return
	}

	if err := cb.Handler(context.TODO(), e); err != nil {
		cb.logf("callback: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)

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
