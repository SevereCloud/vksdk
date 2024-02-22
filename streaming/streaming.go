/*
Package streaming implements Streaming API.

Note that Streaming API is available in beta testing mode. This could be
changes and amended.

Streaming API transfers data via the WebSocket protocol. You create a stream,
add necessary rules and receive data matching the rules.

With Streaming API you can receive not more than 1% of all public data
matching the rules set. To get access to the extended Streaming API version
that includes 100% data please write to support
https://vk.com/support?act=new_api

VK documentation https://vk.com/dev/streaming_api_docs

# Initialization

This can be used with a service token.

	vk := api.NewVK(serviceToken)
	s, err := streaming.NewStreaming(vk)

You can change client for http and websocket:

	s.Client = ... // default http.DefaultClient,
	s.Dialer = ... // default websocket.DefaultDialer,

# Rules Format

A rule is a set of keywords. If they present in an object text, that object
gets into the stream.

- If the words are listed without double quotes, the search will be
simplified (all word forms, case insensitive).

- To search an exact match (all word forms, case sensitive etc.) each word
should be put inside double quotes.

- Use minus (-) to exclude text with this keyword.

Each rule has a ‘’’value’’’, the rule’s content and a ‘’’tag’’’. With each
object you receive a list of its tags to figure out what rules does it fit.

# Limitations

- the maximum number of rules — 300;

- the maximum number of keywords in a rule — 100;

- the maximum length of a rule (value) in bytes — 4096;

# Methods List

You need the following methods to work with Streaming API:

	s.GetRules() //  gets current rules from a stream
	s.AddRule(tag, value) // adds a new rule to a stream
	s.DeleteRule(tag) // removes a rule from a stream
	s.UpdateRules(rules) // removes all rules and adds a new rule to a stream

# Getting Stream

Handler for event:

	s.OnEvent(func(e streaming.Event) {
		...
	}))

For start streaming:

	if err := s.Run(); err != nil {
		...
	}

For stop streaming:

	s.Shutdown()
*/
package streaming

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"sync/atomic"

	"github.com/gorilla/websocket"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/internal"
)

type response struct {
	Code           int            `json:"code"`                      // reply code.
	Event          Event          `json:"event,omitempty"`           // (for code = 100) event
	Rules          []Rule         `json:"rules,omitempty"`           // (for code = 200) info about rules in the stream.
	ServiceMessage serviceMessage `json:"service_message,omitempty"` // (for code = 300) service message
	Error          Error          `json:"error,omitempty"`           // (for code = 400) error description.
}

const (
	codeEvent = 100
	codeError = 400
)

// serviceMessage describes a service message. It is returned for replies with
// code=300 field.
type serviceMessage struct {
	Message     string `json:"message"`      // Service message text
	ServiceCode int    `json:"service_code"` // Service message code
}

// Rule struct.
//
// Each rule has a "value", the rule’s content and a "tag". With each object
// you receive a list of its tags to figure out what rules does it fit.
type Rule struct {
	Tag   string `json:"tag"`   // rule’s tag
	Value string `json:"value"` // string representation of the rule
}

// Streaming struct.
type Streaming struct {
	Endpoint string // URL for further requests
	Key      string // access key for streaming api
	StreamID int

	Client    *http.Client      // A Client is an HTTP client
	Dialer    *websocket.Dialer // A Dialer contains options for connecting to WebSocket server
	UserAgent string            // UserAgent sent in the request.

	inShutdown int32
	eventFunc  []func(Event)
}

func (s *Streaming) doRequest(req *http.Request) (*response, error) {
	req.Header.Set("User-Agent", s.UserAgent)

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r response

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	if r.Code == codeError {
		return nil, &r.Error
	}

	return &r, nil
}

// GetRules gets current rules from a stream.
func (s *Streaming) GetRules() ([]Rule, error) {
	link := fmt.Sprintf("https://%s/rules/?key=%s", s.Endpoint, s.Key)

	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return nil, err
	}

	r, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}

	return r.Rules, nil
}

// AddRule adds a new rule to a stream.
//
// - the maximum length of a rule's tag (tag) in bytes — 256.
func (s *Streaming) AddRule(tag, value string) error {
	link := fmt.Sprintf("https://%s/rules/?key=%s", s.Endpoint, s.Key)

	type content struct {
		Rule Rule `json:"rule"`
	}

	body := &content{
		Rule: Rule{
			Tag:   tag,
			Value: value,
		},
	}

	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest(http.MethodPost, link, buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = s.doRequest(req)

	return err
}

// DeleteRule removes a rule from a stream.
func (s *Streaming) DeleteRule(tag string) error {
	link := fmt.Sprintf("https://%s/rules/?key=%s", s.Endpoint, s.Key)

	type content struct {
		Tag string `json:"tag"`
	}

	body := &content{
		Tag: tag,
	}

	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest(http.MethodDelete, link, buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = s.doRequest(req)

	return err
}

// UpdateRules removes all rules and adds a new rule to a stream.
func (s *Streaming) UpdateRules(rules []Rule) error {
	oldRules, err := s.GetRules()
	if err != nil {
		return err
	}

	for _, rule := range oldRules {
		err := s.DeleteRule(rule.Tag)
		if err != nil {
			return err
		}
	}

	for _, rule := range rules {
		err := s.AddRule(rule.Tag, rule.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Streaming) handlerWebsocket(r response) error {
	switch r.Code {
	case codeEvent:
		for _, f := range s.eventFunc {
			f(r.Event)
		}
	case codeError:
		return &r.Error
	}

	return nil
}

// OnEvent event handler.
func (s *Streaming) OnEvent(f func(Event)) {
	s.eventFunc = append(s.eventFunc, f)
}

// Run starting stream.
func (s *Streaming) Run() error {
	atomic.StoreInt32(&s.inShutdown, 0)

	u := url.URL{
		Scheme: "wss",
		Host:   s.Endpoint,
		Path:   "/stream",
	}
	q := u.Query()
	q.Set("key", s.Key)
	q.Set("stream_id", strconv.Itoa(s.StreamID))
	u.RawQuery = q.Encode()

	requestHeader := make(http.Header)
	requestHeader.Add("User-Agent", s.UserAgent)

	c, wsResp, err := s.Dialer.Dial(u.String(), requestHeader)
	if err != nil {
		if errors.Is(err, websocket.ErrBadHandshake) && wsResp != nil {
			var r response

			err = json.NewDecoder(wsResp.Body).Decode(&r)
			if err != nil {
				return err
			}

			return s.handlerWebsocket(r)
		}

		return err
	}

	defer wsResp.Body.Close()
	defer c.Close()

	c.SetPingHandler(nil)

	for atomic.LoadInt32(&s.inShutdown) == 0 {
		var r response

		_, message, err := c.ReadMessage()
		if err != nil {
			return err
		}

		err = json.Unmarshal(message, &r)
		if err != nil {
			return err
		}

		err = s.handlerWebsocket(r)
		if err != nil {
			return err
		}
	}

	err = c.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	)

	return err
}

// Shutdown gracefully shuts down the stream.
func (s *Streaming) Shutdown() {
	atomic.StoreInt32(&s.inShutdown, 1)
}

// NewStreaming returns a new Streaming.
//
// This can be called with a service token.
//
// The Streaming will use the http.DefaultClient.
// This means that if the http.DefaultClient is modified by other components
// of your application the modifications will be picked up by the SDK as well.
func NewStreaming(vk *api.VK) (*Streaming, error) {
	resp, err := vk.StreamingGetServerURL(nil)
	if err != nil {
		return nil, err
	}

	s := &Streaming{
		Endpoint:  resp.Endpoint,
		Key:       resp.Key,
		StreamID:  0,
		Client:    http.DefaultClient,
		Dialer:    websocket.DefaultDialer,
		UserAgent: internal.UserAgent,
	}

	return s, nil
}
