package longpoll

import (
	"context"
	"errors"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/events"
	"github.com/stretchr/testify/assert"
)

func TestLongPoll_Shutdown(t *testing.T) {
	t.Parallel()

	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	userToken := os.Getenv("USER_TOKEN")
	if userToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	vk := api.NewVK(groupToken)
	lp, _ := NewLongPollCommunity(vk)
	lp.MessageNew(func(_ context.Context, _ events.MessageNewObject) {
		lp.Shutdown()
	})

	c1 := make(chan string)

	go func() {
		err := lp.Run()
		if err != nil && !errors.Is(err, context.Canceled) {
			t.Error(err)
		}

		c1 <- "one"
	}()

	time.Sleep(time.Millisecond * 300)

	vkUser := api.NewVK(userToken)

	_, err := vkUser.MessagesSend(api.Params{
		"peer_id":   -lp.GroupID,
		"random_id": 0,
		"message":   "lp.Shutdown()",
	})
	if err != nil {
		t.Fatal(err)
	}

	// time.Sleep(time.Millisecond * 300)
	select {
	case <-time.After(time.Second * 3):
		lp.Shutdown()
		t.Fatal("timeout")
	case <-c1:
	}
}

func TestLongPoll_Handler(t *testing.T) {
	t.Parallel()
	//nolint:gocyclo
	lp := &LongPoll{}

	lp.FullResponse(func(_ Response) {})

	if len(lp.funcFullResponseList) != 1 {
		t.Error("Want len = 1")
	}
}

func TestNewLongPoll(t *testing.T) {
	t.Parallel()

	f := func(vk *api.VK, groupID int, wantErr bool) {
		_, err := NewLongPoll(vk, groupID)
		if (err != nil) != wantErr {
			t.Errorf("NewLongPoll() error = %v, wantErr %v", err, wantErr)
			return
		}
	}

	f(api.NewVK(""), 0, true)

	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	groupID, _ := strconv.Atoi(os.Getenv("GROUP_ID"))
	f(api.NewVK(groupToken), groupID, false)
}

func TestNewLongPollCommunity(t *testing.T) {
	t.Parallel()

	f := func(vk *api.VK, wantErr bool) {
		_, err := NewLongPollCommunity(vk)
		if (err != nil) != wantErr {
			t.Errorf("NewLongPollCommunity() error = %v, wantErr %v", err, wantErr)
			return
		}
	}

	f(api.NewVK(""), true)
	t.Run("groupToken", func(t *testing.T) {
		t.Parallel()

		groupToken := os.Getenv("GROUP_TOKEN")
		if groupToken == "" {
			t.Skip("GROUP_TOKEN empty")
		}

		f(api.NewVK(groupToken), false)
	})
	t.Run("userToken", func(t *testing.T) {
		t.Parallel()

		userToken := os.Getenv("USER_TOKEN")
		if userToken == "" {
			t.Skip("USER_TOKEN empty")
		}

		f(api.NewVK(userToken), true)
	})
}

func TestLongPoll_checkResponse(t *testing.T) { //nolint: tparallel
	t.Parallel()

	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	vk := api.NewVK(groupToken)
	groupID, _ := strconv.Atoi(os.Getenv("GROUP_ID"))
	lp, _ := NewLongPoll(vk, groupID)

	tests := []struct {
		name        string
		argResponse Response
		wantErr     bool
	}{
		{
			name: "ok",
		},
		{
			name:        "failed: 1",
			argResponse: Response{Failed: 1},
		},
		{
			name:        "failed: 2",
			argResponse: Response{Failed: 2},
		},
		{
			name:        "failed: 3",
			argResponse: Response{Failed: 3},
		},
		{
			name:        "failed: 4",
			argResponse: Response{Failed: 4},
			wantErr:     true,
		},
	}
	for _, tt := range tests { //nolint: paralleltest
		t.Run(tt.name, func(t *testing.T) {
			if err := lp.checkResponse(tt.argResponse); (err != nil) != tt.wantErr {
				t.Errorf("LongPoll.checkResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TODO: write TestLongPoll_Run
// func TestLongPoll_Run(t *testing.T) {
// 	groupToken := os.Getenv("GROUP_TOKEN")
// 	if groupToken == "" {
// 		t.Skip("GROUP_TOKEN empty")
// 	}
// 	vk := api.NewVK(groupToken)
// 	groupID, _ := strconv.Atoi(os.Getenv("GROUP_ID"))
// 	lp, _ := NewLongPoll(&vk, groupID)
// 	lp.Wait = 1

// 	t.Run("Run", func(t *testing.T) {
//
// 	})
// }

func TestLongPoll_RunError(t *testing.T) {
	t.Parallel()

	vk := api.NewVK("")
	lp, _ := NewLongPoll(vk, 0)
	lp.Wait = 1

	if err := lp.Run(); err == nil {
		t.Error(err)
	}

	if err := lp.RunWithContext(context.Background()); err == nil {
		t.Error(err)
	}

	lp.Server = "http://example.com"

	if err := lp.Run(); err == nil {
		t.Error(err)
	}
}

func TestParseResponse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		response string
		expected Response
	}{
		{
			name:     "failed: 1",
			response: `{"ts":8,"failed":1}`,
			expected: Response{
				Ts:     "8",
				Failed: 1,
			},
		},
		{
			name:     "failed: 2",
			response: `{"failed": 2}`,
			expected: Response{
				Failed: 2,
			},
		},
		{
			name:     "empty updates",
			response: `{"ts":"8","updates":[]}`,
			expected: Response{
				Ts:      "8",
				Updates: []events.GroupEvent{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual, err := parseResponse(strings.NewReader(test.response))
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}
