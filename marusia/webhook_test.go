package marusia_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SevereCloud/vksdk/v3/marusia"
	"github.com/stretchr/testify/assert"
)

type responseSession struct {
	SessionID string `json:"session_id"`
	MessageID int    `json:"message_id"`
	UserID    string `json:"user_id"`
}

type response struct {
	Response marusia.Response `json:"response"` // Данные для ответа.
	Session  responseSession  `json:"session"`  // Данные о сессии.
	Version  string           `json:"version"`  // Версия протокола.
}

func TestWebhook(t *testing.T) {
	t.Parallel()

	wh := marusia.NewWebhook()
	wh.EnableDebuging()

	f := func(r marusia.Request, wantResp response) {
		t.Helper()

		raw, err := json.Marshal(&r)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/webhook", bytes.NewBuffer(raw))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json; encoding=utf-8")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(wh.HandleFunc)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json; encoding=utf-8", rr.Header().Get("Content-Type"))

		var resp response

		err = json.NewDecoder(rr.Body).Decode(&resp)
		assert.NoError(t, err)

		assert.Equal(t, wantResp, resp)
	}

	wh.OnEvent(func(r marusia.Request) (resp marusia.Response) {
		assert.Equal(t, "command", r.Request.Command)

		resp.Text = "text"
		resp.TTS = "tts"

		return
	})
	f(
		marusia.Request{
			Request: marusia.RequestIn{
				Command: "command",
			},
		},
		response{
			Response: marusia.Response{
				Text: "text",
				TTS:  "tts",
			},
			Version: marusia.Version,
		},
	)
}

func TestWebhookBadContentType(t *testing.T) {
	t.Parallel()

	wh := marusia.NewWebhook()

	req, err := http.NewRequest(http.MethodPost, "/webhook", bytes.NewBufferString("test"))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "text/plain; encoding=utf-8")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(wh.HandleFunc)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestWebhookBadJSON(t *testing.T) {
	t.Parallel()

	wh := marusia.NewWebhook()

	req, err := http.NewRequest(http.MethodPost, "/webhook", bytes.NewBufferString("[]"))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json; encoding=utf-8")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(wh.HandleFunc)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestWebhookCors(t *testing.T) {
	t.Parallel()

	wh := marusia.NewWebhook()
	wh.EnableDebuging()

	req, err := http.NewRequest(http.MethodOptions, "/webhook", bytes.NewBufferString("[]"))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(wh.HandleFunc)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, "https://skill-debugger.marusia.mail.ru", rr.Header().Get("Access-Control-Allow-Origin"))
}
