package callback // import "github.com/SevereCloud/vksdk/callback"

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SevereCloud/vksdk/events"
	"github.com/SevereCloud/vksdk/object"
)

// Callback struct SecretKeys [GroupID]SecretKey
type Callback struct {
	ConfirmationKeys map[int]string
	ConfirmationKey  string
	SecretKeys       map[int]string
	SecretKey        string
	events.FuncList
}

// HandleFunc handler
func (cb Callback) HandleFunc(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var e object.GroupEvent
	if err := decoder.Decode(&e); err != nil {
		log.Printf("Callback.HandleFunc: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)

		return
	}

	if cb.SecretKeys[e.GroupID] != "" || cb.SecretKey != "" {
		if e.Secret != cb.SecretKeys[e.GroupID] && e.Secret != cb.SecretKey {
			http.Error(w, "Bad Secret", http.StatusForbidden)
			return
		}
	}

	if e.Type == "confirmation" {
		if cb.ConfirmationKeys[e.GroupID] != "" {
			fmt.Fprintf(w, cb.ConfirmationKeys[e.GroupID])
		} else {
			fmt.Fprintf(w, cb.ConfirmationKey)
		}

		return
	}

	if err := cb.Handler(e); err != nil {
		log.Printf("Callback.HandleFunc: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)

		return
	}

	fmt.Fprintf(w, "ok")
}
