package callback

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type callbackEvent struct {
	Type    string          `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int             `json:"group_id"`
}

// Callback struct SecretKeys [GroupID]SecretKey
type Callback struct {
	SecretKeys map[int]string
	SecretKey  string
}

// HandleFunc handler
func (cb *Callback) HandleFunc(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var e callbackEvent
	err := decoder.Decode(&e)
	if err != nil {
		fmt.Fprintf(w, "decode error")
		return
	}
	if e.Type == "confirmation" {
		if cb.SecretKeys[e.GroupID] != "" {
			fmt.Fprintf(w, cb.SecretKeys[e.GroupID])
		} else {
			fmt.Fprintf(w, cb.SecretKey)
		}
		return
	}

	switch e.Type {
	case "message_new":
		// TODO new message received
		var obj MessageNewObject
		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			log.Print(r.Body)
			log.Print(err)
		}
	case "message_reply":
		// TODO new message sent
	case "message_allow":
		// TODO new user consent to messages sending
	case "messages_deny":
		// TODO new user prohibition to messages sending
	case "photo_new":
		// TODO new photo in a community
	case "photo_comment_new":
		// TODO new photo comment
	case "photo_comment_edit":
		// TODO editing a photo comment
	case "photo_comment_restore":
		// TODO restoring a photo comment
	case "photo_comment_delete":
		// TODO deleting a photo comment
	case "audio_new":
		// TODO new audio in a community list
	case "video_new":
		// TODO new video in a community list
	case "video_comment_new":
		// TODO new video comment
	case "video_comment_edit":
		// TODO editing a video comment
	case "video_comment_restore":
		// TODO restoring a video comment
	case "video_comment_delete":
		// TODO deleting a video comment
	case "wall_post_new":
		// TODO adding a post on wall
	case "wall_repost":
		// TODO new repost
	case "wall_reply_new":
		// TODO new comment on wall
	case "wall_reply_edit":
		// TODO editing a comment on wall
	case "wall_reply_restore":
		// TODO restoring a comment on wall
	case "wall_reply_delete":
		// TODO deleting a comment on wall
	case "board_post_new":
		// TODO new comment in a board
	case "board_post_edit":
		// TODO editing a board comment
	case "board_post_restore":
		// TODO restoring a board comment
	case "board_post_delete":
		// TODO deleting a board comment
	case "market_comment_new":
		// TODO new comment to a market item
	case "market_comment_edit":
		// TODO editing a market comment
	case "market_comment_restore":
		// TODO restoring a market comment
	case "market_comment_delete":
		// TODO deleting a market comment
	case "group_leave":
		// TODO member removed from a community
	case "group_join":
		// TODO new member or join request
	case "user_block":
		// TODO new user in the blacklist
	case "user_unblock":
		// TODO user has been deleted from the blacklist
	case "poll_vote_new":
		// TODO new vote in a public poll
	case "group_officers_edit":
		// TODO changes in the administrators list
	case "group_change_settings":
		// TODO changes in a community settings
	case "group_change_photo":
		// TODO changes of community main photo
	case "vkpay_transaction":
		// TODO vkpay_transaction
	}
}
