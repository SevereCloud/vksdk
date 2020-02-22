package streaming_test

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gorilla/websocket"

	"github.com/stretchr/testify/assert"

	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/streaming"
)

var stream *streaming.Streaming // nolint:gochecknoglobals

func needServiceToken(t *testing.T) {
	t.Helper()

	if stream == nil {
		t.Skip("SERVICE_TOKEN empty or bad")
	}
}

func TestMain(m *testing.M) {
	vkService := api.Init(os.Getenv("SERVICE_TOKEN"))
	vkService.Limit = 3

	stream, _ = streaming.Init(vkService)

	runTests := m.Run()
	os.Exit(runTests)
}

func TestStreaming_AllRule(t *testing.T) {
	needServiceToken(t)

	err := stream.AddRule("abcd", "test")
	assert.NoError(t, err)

	err = stream.UpdateRules([]streaming.Rule{{Tag: "test", Value: "test"}})
	assert.NoError(t, err)
}

func TestStreaming_AddRule(t *testing.T) {
	needServiceToken(t)

	err := stream.AddRule("test", "_")
	assert.Error(t, err)
}

func TestStreaming_DeleteRule(t *testing.T) {
	needServiceToken(t)

	err := stream.DeleteRule("_")
	assert.Error(t, err)
}

func TestStreaming_UpdateRules(t *testing.T) {
	needServiceToken(t)

	err := stream.UpdateRules([]streaming.Rule{{Tag: "_", Value: "_"}})
	assert.Error(t, err)
}

func TestStreaming_NoHost(t *testing.T) {
	s := &streaming.Streaming{
		Client: http.DefaultClient,
	}

	_, err := s.GetRules()
	assert.Error(t, err)

	err = s.AddRule("test", "_")
	assert.Error(t, err)

	err = s.DeleteRule("test")
	assert.Error(t, err)

	err = s.UpdateRules([]streaming.Rule{{Tag: "test", Value: "ujhbkf"}})
	assert.Error(t, err)
}

func TestStreaming_BadHost(t *testing.T) {
	s := &streaming.Streaming{
		Endpoint: "a b.com",
		Client:   http.DefaultClient,
	}

	_, err := s.GetRules()
	assert.Error(t, err)

	err = s.AddRule("test", "_")
	assert.Error(t, err)

	err = s.DeleteRule("test")
	assert.Error(t, err)

	err = s.UpdateRules([]streaming.Rule{{Tag: "test", Value: "ujhbkf"}})
	assert.Error(t, err)
}

func TestStreaming_NotVKStreamingHost(t *testing.T) {
	s := &streaming.Streaming{
		Endpoint: "vk.com",
		Client:   http.DefaultClient,
	}

	_, err := s.GetRules()
	assert.Error(t, err)
}

func TestStreaming_OnEvent(t *testing.T) {
	s := &streaming.Streaming{}
	s.OnEvent(func(e streaming.Event) {})
}

func TestStreaming_Run(t *testing.T) {
	needServiceToken(t)

	var skip bool

	good := make(chan bool, 1)
	exit := make(chan bool, 1)

	err := stream.UpdateRules([]streaming.Rule{{Tag: "test", Value: "привет"}})
	if !assert.NoError(t, err) {
		return
	}

	stream.OnEvent(func(e streaming.Event) {
		// TODO: assert
		good <- true
	})

	go func() {
		err = stream.Run()
		assert.NoError(t, err)
		exit <- true
	}()

	select {
	case <-good:
	case <-time.After(30 * time.Second):
		skip = true
	}

	stream.Shutdown()

	select {
	case <-exit:
	case <-time.After(5 * time.Second):
	}

	_ = stream.UpdateRules([]streaming.Rule{})

	if skip {
		t.Error("Timeout")
	}
}

func TestStreaming_Run_BadStreamID(t *testing.T) {
	needServiceToken(t)

	stream.StreamID = -1

	err := stream.Run()
	assert.Error(t, err)

	stream.StreamID = 0
}

func TestStreaming_Run_Bad(t *testing.T) {
	s := &streaming.Streaming{
		Endpoint: "a b.com",
		Client:   http.DefaultClient,
		Dialer:   websocket.DefaultDialer,
	}

	err := s.Run()
	assert.Error(t, err)
}

func TestInit(t *testing.T) {
	_, err := streaming.Init(api.Init(""))
	assert.Error(t, err)
}
