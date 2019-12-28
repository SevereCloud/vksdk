package callback

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCallback_HandleFunc(t *testing.T) {
	type fields struct {
		ConfirmationKeys map[int]string
		ConfirmationKey  string
		SecretKeys       map[int]string
		SecretKey        string
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
