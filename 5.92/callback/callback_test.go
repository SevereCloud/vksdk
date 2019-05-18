package callback

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SevereCloud/vksdk/5.92/handler"
	"github.com/SevereCloud/vksdk/5.92/object"
)

func TestCallback_Handler(t *testing.T) { // nolint:gocyclo
	cb := &Callback{}
	t.Run("MessageNew", func(t *testing.T) {
		cb.MessageNew(func(obj object.MessageNewObject, groupID int) {})
		if len(cb.funcList.MessageNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MessageReply", func(t *testing.T) {
		cb.MessageReply(func(obj object.MessageReplyObject, groupID int) {})
		if len(cb.funcList.MessageReply) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MessageEdit", func(t *testing.T) {
		cb.MessageEdit(func(obj object.MessageEditObject, groupID int) {})
		if len(cb.funcList.MessageEdit) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MessageAllow", func(t *testing.T) {
		cb.MessageAllow(func(obj object.MessageAllowObject, groupID int) {})
		if len(cb.funcList.MessageAllow) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MessageDeny", func(t *testing.T) {
		cb.MessageDeny(func(obj object.MessageDenyObject, groupID int) {})
		if len(cb.funcList.MessageDeny) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MessageTypingState", func(t *testing.T) {
		cb.MessageTypingState(func(obj object.MessageTypingStateObject, groupID int) {})
		if len(cb.funcList.MessageTypingState) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("PhotoNew", func(t *testing.T) {
		cb.PhotoNew(func(obj object.PhotoNewObject, groupID int) {})
		if len(cb.funcList.PhotoNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("PhotoCommentNew", func(t *testing.T) {
		cb.PhotoCommentNew(func(obj object.PhotoCommentNewObject, groupID int) {})
		if len(cb.funcList.PhotoCommentNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("PhotoCommentEdit", func(t *testing.T) {
		cb.PhotoCommentEdit(func(obj object.PhotoCommentEditObject, groupID int) {})
		if len(cb.funcList.PhotoCommentEdit) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("PhotoCommentRestore", func(t *testing.T) {
		cb.PhotoCommentRestore(func(obj object.PhotoCommentRestoreObject, groupID int) {})
		if len(cb.funcList.PhotoCommentRestore) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("PhotoCommentDelete", func(t *testing.T) {
		cb.PhotoCommentDelete(func(obj object.PhotoCommentDeleteObject, groupID int) {})
		if len(cb.funcList.PhotoCommentDelete) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("AudioNew", func(t *testing.T) {
		cb.AudioNew(func(obj object.AudioNewObject, groupID int) {})
		if len(cb.funcList.AudioNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("VideoNew", func(t *testing.T) {
		cb.VideoNew(func(obj object.VideoNewObject, groupID int) {})
		if len(cb.funcList.VideoNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("VideoCommentNew", func(t *testing.T) {
		cb.VideoCommentNew(func(obj object.VideoCommentNewObject, groupID int) {})
		if len(cb.funcList.VideoCommentNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("VideoCommentEdit", func(t *testing.T) {
		cb.VideoCommentEdit(func(obj object.VideoCommentEditObject, groupID int) {})
		if len(cb.funcList.VideoCommentEdit) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("VideoCommentRestore", func(t *testing.T) {
		cb.VideoCommentRestore(func(obj object.VideoCommentRestoreObject, groupID int) {})
		if len(cb.funcList.VideoCommentRestore) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("VideoCommentDelete", func(t *testing.T) {
		cb.VideoCommentDelete(func(obj object.VideoCommentDeleteObject, groupID int) {})
		if len(cb.funcList.VideoCommentDelete) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("WallPostNew", func(t *testing.T) {
		cb.WallPostNew(func(obj object.WallPostNewObject, groupID int) {})
		if len(cb.funcList.WallPostNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("WallRepost", func(t *testing.T) {
		cb.WallRepost(func(obj object.WallRepostObject, groupID int) {})
		if len(cb.funcList.WallRepost) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("WallReplyNew", func(t *testing.T) {
		cb.WallReplyNew(func(obj object.WallReplyNewObject, groupID int) {})
		if len(cb.funcList.WallReplyNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("WallReplyEdit", func(t *testing.T) {
		cb.WallReplyEdit(func(obj object.WallReplyEditObject, groupID int) {})
		if len(cb.funcList.WallReplyEdit) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("WallReplyRestore", func(t *testing.T) {
		cb.WallReplyRestore(func(obj object.WallReplyRestoreObject, groupID int) {})
		if len(cb.funcList.WallReplyRestore) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("WallReplyDelete", func(t *testing.T) {
		cb.WallReplyDelete(func(obj object.WallReplyDeleteObject, groupID int) {})
		if len(cb.funcList.WallReplyDelete) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("BoardPostNew", func(t *testing.T) {
		cb.BoardPostNew(func(obj object.BoardPostNewObject, groupID int) {})
		if len(cb.funcList.BoardPostNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("BoardPostEdit", func(t *testing.T) {
		cb.BoardPostEdit(func(obj object.BoardPostEditObject, groupID int) {})
		if len(cb.funcList.BoardPostEdit) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("BoardPostRestore", func(t *testing.T) {
		cb.BoardPostRestore(func(obj object.BoardPostRestoreObject, groupID int) {})
		if len(cb.funcList.BoardPostRestore) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("BoardPostDelete", func(t *testing.T) {
		cb.BoardPostDelete(func(obj object.BoardPostDeleteObject, groupID int) {})
		if len(cb.funcList.BoardPostDelete) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MarketCommentNew", func(t *testing.T) {
		cb.MarketCommentNew(func(obj object.MarketCommentNewObject, groupID int) {})
		if len(cb.funcList.MarketCommentNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MarketCommentEdit", func(t *testing.T) {
		cb.MarketCommentEdit(func(obj object.MarketCommentEditObject, groupID int) {})
		if len(cb.funcList.MarketCommentEdit) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MarketCommentRestore", func(t *testing.T) {
		cb.MarketCommentRestore(func(obj object.MarketCommentRestoreObject, groupID int) {})
		if len(cb.funcList.MarketCommentRestore) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MarketCommentDelete", func(t *testing.T) {
		cb.MarketCommentDelete(func(obj object.MarketCommentDeleteObject, groupID int) {})
		if len(cb.funcList.MarketCommentDelete) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("GroupLeave", func(t *testing.T) {
		cb.GroupLeave(func(obj object.GroupLeaveObject, groupID int) {})
		if len(cb.funcList.GroupLeave) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("GroupJoin", func(t *testing.T) {
		cb.GroupJoin(func(obj object.GroupJoinObject, groupID int) {})
		if len(cb.funcList.GroupJoin) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("UserBlock", func(t *testing.T) {
		cb.UserBlock(func(obj object.UserBlockObject, groupID int) {})
		if len(cb.funcList.UserBlock) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("UserUnblock", func(t *testing.T) {
		cb.UserUnblock(func(obj object.UserUnblockObject, groupID int) {})
		if len(cb.funcList.UserUnblock) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("PollVoteNew", func(t *testing.T) {
		cb.PollVoteNew(func(obj object.PollVoteNewObject, groupID int) {})
		if len(cb.funcList.PollVoteNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("GroupOfficersEdit", func(t *testing.T) {
		cb.GroupOfficersEdit(func(obj object.GroupOfficersEditObject, groupID int) {})
		if len(cb.funcList.GroupOfficersEdit) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("GroupChangeSettings", func(t *testing.T) {
		cb.GroupChangeSettings(func(obj object.GroupChangeSettingsObject, groupID int) {})
		if len(cb.funcList.GroupChangeSettings) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("GroupChangePhoto", func(t *testing.T) {
		cb.GroupChangePhoto(func(obj object.GroupChangePhotoObject, groupID int) {})
		if len(cb.funcList.GroupChangePhoto) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("VkpayTransaction", func(t *testing.T) {
		cb.VkpayTransaction(func(obj object.VkpayTransactionObject, groupID int) {})
		if len(cb.funcList.VkpayTransaction) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("LeadFormsNew", func(t *testing.T) {
		cb.LeadFormsNew(func(obj object.LeadFormsNewObject, groupID int) {})
		if len(cb.funcList.LeadFormsNew) != 1 {
			t.Error("Want len = 1")
		}
	})
}

func TestCallback_HandleFunc(t *testing.T) {
	type fields struct {
		ConfirmationKeys map[int]string
		ConfirmationKey  string
		SecretKeys       map[int]string
		SecretKey        string
		funcList         handler.FuncList
	}
	tests := []struct {
		name     string
		fields   fields
		body     string
		expected string
	}{
		{
			name: "check Decode error",
		},
		// ConfirmationKey test
		{
			name: "check ConfirmationKey",
			fields: fields{
				ConfirmationKey: "confirmation_123456",
			},
			body:     `{"type": "confirmation", "group_id": 123456}`,
			expected: "confirmation_123456",
		},
		{
			name: "check ConfirmationKeys",
			fields: fields{
				ConfirmationKeys: map[int]string{
					123456: "confirmation_123456",
					654321: "confirmation_1654321",
				},
			},
			body:     `{"type": "confirmation", "group_id": 123456}`,
			expected: "confirmation_123456",
		},
		{
			name: "check ConfirmationKeys bad not found",
			fields: fields{
				ConfirmationKeys: map[int]string{
					654321: "confirmation_1654321",
				},
			},
			body:     `{"type": "confirmation", "group_id": 123456}`,
			expected: "",
		},
		{
			name: "check ConfirmationKeys with true key",
			fields: fields{
				ConfirmationKey: "confirmation_1654321",
				ConfirmationKeys: map[int]string{
					123456: "confirmation_123456",
				},
			},
			body:     `{"type": "confirmation", "group_id": 123456}`,
			expected: "confirmation_123456",
		},
		{
			name: "check ConfirmationKey with true key",
			fields: fields{
				ConfirmationKey: "confirmation_123456",
				ConfirmationKeys: map[int]string{
					654321: "confirmation_1654321",
				},
			},
			body:     `{"type": "confirmation", "group_id": 123456}`,
			expected: "confirmation_123456",
		},
		// SecretKeys test
		{
			name: "check SecretKey",
			fields: fields{
				ConfirmationKey: "confirmation_123456",
				SecretKey:       "secret_123456",
			},
			body:     `{"type": "confirmation", "group_id": 123456, "secret": "secret_123456"}`,
			expected: "confirmation_123456",
		},
		{
			name: "check SecretKey bad",
			fields: fields{
				ConfirmationKey: "confirmation_123456",
				SecretKey:       "secret_654321",
			},
			body:     `{"type": "confirmation", "group_id": 123456, "secret": "secret_123456"}`,
			expected: "bad secret",
		},
		{
			name: "check SecretKeys",
			fields: fields{
				ConfirmationKey: "confirmation_123456",
				SecretKeys: map[int]string{
					123456: "secret_123456",
				},
			},
			body:     `{"type": "confirmation", "group_id": 123456, "secret": "secret_123456"}`,
			expected: "confirmation_123456",
		},
		{
			name: "check SecretKeys bad",
			fields: fields{
				ConfirmationKey: "confirmation_123456",
				SecretKeys: map[int]string{
					123456: "secret_654321",
					654321: "secret_123456",
				},
			},
			body:     `{"type": "confirmation", "group_id": 123456, "secret": "secret_123456"}`,
			expected: "bad secret",
		},
		{
			name: "check SecretKeys not found",
			fields: fields{
				ConfirmationKey: "confirmation_123456",
				SecretKeys: map[int]string{
					654321: "secret_123456",
				},
			},
			body:     `{"type": "confirmation", "group_id": 123456, "secret": "secret_123456"}`,
			expected: "confirmation_123456",
		},
		{
			name: "check SecretKey with true",
			fields: fields{
				ConfirmationKey: "confirmation_123456",
				SecretKey:       "secret_123456",
				SecretKeys: map[int]string{
					654321: "secret_654321",
				},
			},
			body:     `{"type": "confirmation", "group_id": 123456, "secret": "secret_123456"}`,
			expected: "confirmation_123456",
		},
		{
			name: "check SecretKeys with true",
			fields: fields{
				ConfirmationKey: "confirmation_123456",
				SecretKey:       "secret_654321",
				SecretKeys: map[int]string{
					123456: "secret_123456",
				},
			},
			body:     `{"type": "confirmation", "group_id": 123456, "secret": "secret_123456"}`,
			expected: "confirmation_123456",
		},
		{
			name:     "check bad message_new",
			fields:   fields{},
			body:     `{"type": "message_new", "object": 1}`,
			expected: "",
		},
		{
			name:   "check good message_new",
			fields: fields{},
			body: `{"type": "message_new", "object": 
			{"date":1558100846,"from_id":194250225,"id":0,"out":0,
			"peer_id":2000000001,"text":"","conversation_message_id":147826,
			"action":{"type":"chat_kick_user","member_id":194250225},
			"fwd_messages":[],"important":false,"random_id":0,"attachments":[],
			"is_hidden":false}}`,
			expected: "ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := Callback{
				ConfirmationKeys: tt.fields.ConfirmationKeys,
				ConfirmationKey:  tt.fields.ConfirmationKey,
				SecretKeys:       tt.fields.SecretKeys,
				SecretKey:        tt.fields.SecretKey,
				funcList:         tt.fields.funcList,
			}

			jsonStr := []byte(tt.body)

			req, err := http.NewRequest("POST", "/callback", bytes.NewBuffer(jsonStr))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(cb.HandleFunc)

			handler.ServeHTTP(rr, req)

			if rr.Body.String() != tt.expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expected)
			}
		})
	}
}
