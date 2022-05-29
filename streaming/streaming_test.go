package streaming_test

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gorilla/websocket"

	"github.com/stretchr/testify/assert"

	"github.com/Derad6709/vksdk/v2/api"
	"github.com/Derad6709/vksdk/v2/streaming"
)

var stream *streaming.Streaming // nolint:gochecknoglobals

func needServiceToken(t *testing.T) {
	t.Helper()

	t.Parallel()

	if stream == nil {
		t.Skip("SERVICE_TOKEN empty or bad")
	}
}

func TestMain(m *testing.M) {
	vkService := api.NewVK(os.Getenv("SERVICE_TOKEN"))
	vkService.Limit = 3

	stream, _ = streaming.NewStreaming(vkService)

	runTests := m.Run()
	os.Exit(runTests)
}

func TestStreaming_AllRule(t *testing.T) { //nolint:paralleltest
	needServiceToken(t)

	err := stream.AddRule("rule", "test")
	assert.NoError(t, err)

	time.Sleep(time.Second)

	err = stream.UpdateRules([]streaming.Rule{{Tag: "test", Value: "test"}})
	assert.NoError(t, err)
}

func TestStreaming_AddRule(t *testing.T) { //nolint:paralleltest
	needServiceToken(t)

	err := stream.AddRule("test", "_")
	assert.Error(t, err)
}

func TestStreaming_DeleteRule(t *testing.T) { //nolint:paralleltest
	needServiceToken(t)

	err := stream.DeleteRule("_")
	assert.Error(t, err)
}

func TestStreaming_UpdateRules(t *testing.T) { //nolint:paralleltest
	needServiceToken(t)

	err := stream.UpdateRules([]streaming.Rule{{Tag: "_", Value: "_"}})
	assert.Error(t, err)
}

func TestStreaming_NoHost(t *testing.T) {
	t.Parallel()

	s := &streaming.Streaming{
		Client: http.DefaultClient,
	}

	_, err := s.GetRules()
	assert.Error(t, err)

	err = s.AddRule("test", "_")
	assert.Error(t, err)

	err = s.DeleteRule("test")
	assert.Error(t, err)

	err = s.UpdateRules([]streaming.Rule{{Tag: "test", Value: "value"}})
	assert.Error(t, err)
}

func TestStreaming_BadHost(t *testing.T) {
	t.Parallel()

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

	err = s.UpdateRules([]streaming.Rule{{Tag: "test", Value: "value"}})
	assert.Error(t, err)
}

func TestStreaming_NotVKStreamingHost(t *testing.T) {
	t.Parallel()

	s := &streaming.Streaming{
		Endpoint: "vk.com",
		Client:   http.DefaultClient,
	}

	_, err := s.GetRules()
	assert.Error(t, err)
}

func TestStreaming_OnEvent(t *testing.T) {
	t.Parallel()

	s := &streaming.Streaming{}
	s.OnEvent(func(e streaming.Event) {})
}

func TestStreaming_Run(t *testing.T) { //nolint:paralleltest
	needServiceToken(t)

	var skip bool

	good := make(chan bool, 1)
	exit := make(chan bool, 1)

	time.Sleep(3 * time.Second)

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

/*func TestStreaming_Run_BadStreamID(t *testing.T) { //nolint:paralleltest
	needServiceToken(t)

	stream.StreamID = -1

	err := stream.Run()
	assert.Error(t, err)

	stream.StreamID = 0
}*/

func TestStreaming_Run_Bad(t *testing.T) {
	t.Parallel()

	s := &streaming.Streaming{
		Endpoint: "a b.com",
		Client:   http.DefaultClient,
		Dialer:   websocket.DefaultDialer,
	}

	err := s.Run()
	assert.Error(t, err)
}

func TestNewStreaming(t *testing.T) {
	t.Parallel()

	_, err := streaming.NewStreaming(api.NewVK(""))
	assert.Error(t, err)
}
