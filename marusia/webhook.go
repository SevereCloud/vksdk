package marusia // import "github.com/Derad6709/vksdk/v2/marusia"

import (
	"encoding/json"
	"mime"
	"net/http"
)

// Version версия протокола.
const Version = "1.0"

const debugURL = "https://skill-debugger.marusia.mail.ru"

// responseSession данные о сессии.
type responseSession struct {
	SessionID   string      `json:"session_id"`
	MessageID   int         `json:"message_id"`
	UserID      string      `json:"user_id"`
	SkillID     string      `json:"skill_id"`
	New         bool        `json:"new"`
	User        User        `json:"user"`
	Application Application `json:"application"`
}

// response структура ответа серверу.
type response struct {
	Response Response        `json:"response"` // Данные для ответа.
	Session  responseSession `json:"session"`  // Данные о сессии.
	Version  string          `json:"version"`  // Версия протокола.
}

// Webhook структура.
type Webhook struct {
	event    func(r Request) Response
	debuging bool
}

// NewWebhook возвращает новый Webhook.
func NewWebhook() *Webhook {
	return &Webhook{}
}

// OnEvent обработчик скилла.
//
// Таймаут ожидания ответа — 5 секунд, после чего сервер Маруси завершит
// сессию.
func (wh *Webhook) OnEvent(f func(r Request) Response) {
	wh.event = f
}

// EnableDebuging включает CORS для https://skill-debugger.marusia.mail.ru
func (wh *Webhook) EnableDebuging() {
	wh.debuging = true
}

// HandleFunc обработчик http запросов.
func (wh *Webhook) HandleFunc(w http.ResponseWriter, r *http.Request) {
	// Проброс CORS-заголовков
	if r.Method == http.MethodOptions {
		if wh.debuging {
			w.Header().Set("Access-Control-Allow-Origin", debugURL)
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
		}

		w.WriteHeader(http.StatusOK)

		return
	}

	mediatype, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if mediatype != "application/json" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	resp := wh.event(req)

	fullResponse := response{
		Response: resp,
		Session: responseSession{
			SessionID:   req.Session.SessionID,
			MessageID:   req.Session.MessageID,
			UserID:      req.Session.UserID,
			SkillID:     req.Session.SkillID,
			New:         req.Session.New,
			User:        req.Session.User,
			Application: req.Session.Application,
		},
		Version: Version,
	}

	if wh.debuging {
		w.Header().Set("Access-Control-Allow-Origin", debugURL)
	}

	// Возвращаем данные
	w.Header().Add("Content-Type", "application/json; encoding=utf-8")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(fullResponse)
}
