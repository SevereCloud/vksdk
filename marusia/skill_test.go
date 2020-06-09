package marusia_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SevereCloud/vksdk/marusia"
	"github.com/stretchr/testify/assert"
)

type my struct {
	Name string `json:"name"`
}

func TestInterfaces_IsScreen(t *testing.T) {
	t.Parallel()

	f := func(i marusia.Interfaces, actual bool) {
		t.Helper()

		assert.Equal(t, i.IsScreen(), actual)
	}

	f(marusia.Interfaces{}, false)
	f(marusia.Interfaces{
		Screen: nil,
	}, false)
	f(marusia.Interfaces{
		Screen: &marusia.Screen{},
	}, true)
}

func TestNewBigImage(t *testing.T) {
	t.Parallel()

	f := func(title, description string, imageID int, actual *marusia.Card) {
		t.Helper()

		card := marusia.NewBigImage(title, description, imageID)

		assert.Equal(t, card, actual)
	}

	f("title", "description", 1234,
		&marusia.Card{
			Type:        marusia.BigImage,
			Title:       "title",
			Description: "description",
			ImageID:     1234,
		},
	)
}

// TODO:
// func TestNewItemsList(t *testing.T) {
// 	t.Parallel()

// 	f := func(title, description string, items []marusia.CardItem, actual *marusia.Card) {
// 		t.Helper()

// 		card := marusia.NewItemsList(title, description, items)

// 		assert.Equal(t, card, actual)
// 	}

// 	f("title", "description", []marusia.CardItem{{1}, {2}},
// 		&marusia.Card{
// 			Type:        marusia.ItemsList,
// 			Title:       "title",
// 			Description: "description",
// 			Items:       []marusia.CardItem{{1}, {2}},
// 		},
// 	)
// }

// TODO:
// func TestNewImageList(t *testing.T) {
// 	t.Parallel()

// 	f := func(title, description string, items []int, actual *marusia.Card) {
// 		t.Helper()

// 		card := marusia.NewImageList(title, description, items...)

// 		assert.Equal(t, card, actual)
// 	}

// 	f("title", "description", []int{1, 2},
// 		&marusia.Card{
// 			Type:        marusia.ItemsList,
// 			Title:       "title",
// 			Description: "description",
// 			Items:       []marusia.CardItem{{1}, {2}},
// 		},
// 	)
// }

func TestResponse_AddURL(t *testing.T) {
	t.Parallel()

	f := func(title string, link string, actual marusia.Response) {
		t.Helper()

		resp := marusia.Response{}
		resp.AddURL(title, link)

		assert.Equal(t, resp, actual)
	}

	f("title", "https://vk.com",
		marusia.Response{
			Buttons: []marusia.Button{
				{
					Title: "title",
					URL:   "https://vk.com",
				},
			},
		},
	)
}

func TestResponse_AddButton(t *testing.T) {
	t.Parallel()

	f := func(title string, payload interface{}, actual marusia.Response) {
		t.Helper()

		resp := marusia.Response{}
		resp.AddButton(title, payload)

		assert.Equal(t, resp, actual)
	}

	f(
		"title",
		nil,
		marusia.Response{
			Buttons: []marusia.Button{
				{
					Title: "title",
				},
			},
		},
	)
	f(
		"title",
		my{"test"},
		marusia.Response{
			Buttons: []marusia.Button{
				{
					Title:   "title",
					Payload: my{"test"},
				},
			},
		},
	)
}

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

	f := func(r marusia.Request, wantResp response) {
		t.Helper()

		raw, err := json.Marshal(&r)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/webhook", bytes.NewBuffer(raw))
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

	req, err := http.NewRequest("POST", "/webhook", bytes.NewBuffer([]byte("test")))
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

	req, err := http.NewRequest("POST", "/webhook", bytes.NewBuffer([]byte("[]")))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json; encoding=utf-8")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(wh.HandleFunc)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
