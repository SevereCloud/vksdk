package longpoll

import (
	"os"
	"strconv"
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"
)

func TestLongPoll_Shutdown(t *testing.T) {
	t.Parallel()

	lp := &LongPoll{}

	lp.Shutdown()

	if lp.inShutdown != 1 {
		t.Error("inShutdown != 1")
	}
}

func TestLongPoll_Handler(t *testing.T) {
	t.Parallel()
	// nolint:gocyclo
	lp := &LongPoll{}

	lp.FullResponse(func(resp Response) {})

	if len(lp.funcFullResponseList) != 1 {
		t.Error("Want len = 1")
	}
}

func TestNewLongPoll(t *testing.T) {
	f := func(vk *api.VK, groupID int, wantErr bool) {
		_, err := NewLongPoll(vk, groupID)
		if (err != nil) != wantErr {
			t.Errorf("NewLongPoll() error = %v, wantErr %v", err, wantErr)
			return
		}
	}

	f(api.NewVK(""), 0, true)
	t.Run("groupToken", func(t *testing.T) {
		groupToken := os.Getenv("GROUP_TOKEN")
		if groupToken == "" {
			t.Skip("GROUP_TOKEN empty")
		}

		groupID, _ := strconv.Atoi(os.Getenv("GROUP_ID"))
		f(api.NewVK(groupToken), groupID, false)
	})
}

func TestNewLongPollCommunity(t *testing.T) {
	f := func(vk *api.VK, wantErr bool) {
		_, err := NewLongPollCommunity(vk)
		if (err != nil) != wantErr {
			t.Errorf("NewLongPollCommunity() error = %v, wantErr %v", err, wantErr)
			return
		}
	}

	f(api.NewVK(""), true)
	t.Run("groupToken", func(t *testing.T) {
		groupToken := os.Getenv("GROUP_TOKEN")
		if groupToken == "" {
			t.Skip("GROUP_TOKEN empty")
		}
		f(api.NewVK(groupToken), false)
	})
	t.Run("userToken", func(t *testing.T) {
		userToken := os.Getenv("USER_TOKEN")
		if userToken == "" {
			t.Skip("USER_TOKEN empty")
		}
		f(api.NewVK(userToken), true)
	})
}

func TestLongPoll_checkResponse(t *testing.T) {
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
	for _, tt := range tests {
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

	lp.Server = "http://example.com"

	if err := lp.Run(); err == nil {
		t.Error(err)
	}
}
