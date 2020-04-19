package longpoll

import (
	"os"
	"strconv"
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/object"
)

func TestLongpoll_Shutdown(t *testing.T) {
	t.Parallel()

	t.Run("Shutdown", func(t *testing.T) {
		lp := &Longpoll{}
		lp.Shutdown()
		if lp.inShutdown != 1 {
			t.Error("inShutdown != 1")
		}
	})
}

func TestLongpoll_Handler(t *testing.T) {
	t.Parallel()
	// nolint:gocyclo
	lp := &Longpoll{}

	t.Run("FullResponse", func(t *testing.T) {
		lp.FullResponse(func(resp object.LongpollBotResponse) {})
		if len(lp.funcFullResponseList) != 1 {
			t.Error("Want len = 1")
		}
	})
}

func TestNewLongpoll(t *testing.T) {
	f := func(vk *api.VK, groupID int, wantErr bool) {
		_, err := NewLongpoll(vk, groupID)
		if (err != nil) != wantErr {
			t.Errorf("NewLongpoll() error = %v, wantErr %v", err, wantErr)
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

func TestNewLongpollCommunity(t *testing.T) {
	f := func(vk *api.VK, wantErr bool) {
		_, err := NewLongpollCommunity(vk)
		if (err != nil) != wantErr {
			t.Errorf("NewLongpollCommunity() error = %v, wantErr %v", err, wantErr)
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

func TestLongpoll_checkResponse(t *testing.T) {
	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	vk := api.NewVK(groupToken)
	groupID, _ := strconv.Atoi(os.Getenv("GROUP_ID"))
	lp, _ := NewLongpoll(vk, groupID)

	tests := []struct {
		name        string
		argResponse object.LongpollBotResponse
		wantErr     bool
	}{
		{
			name: "ok",
		},
		{
			name:        "failed: 1",
			argResponse: object.LongpollBotResponse{Failed: 1},
		},
		{
			name:        "failed: 2",
			argResponse: object.LongpollBotResponse{Failed: 2},
		},
		{
			name:        "failed: 3",
			argResponse: object.LongpollBotResponse{Failed: 3},
		},
		{
			name:        "failed: 4",
			argResponse: object.LongpollBotResponse{Failed: 4},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := lp.checkResponse(tt.argResponse); (err != nil) != tt.wantErr {
				t.Errorf("Longpoll.checkResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TODO: write TestLongpoll_Run
// func TestLongpoll_Run(t *testing.T) {
// 	groupToken := os.Getenv("GROUP_TOKEN")
// 	if groupToken == "" {
// 		t.Skip("GROUP_TOKEN empty")
// 	}
// 	vk := api.NewVK(groupToken)
// 	groupID, _ := strconv.Atoi(os.Getenv("GROUP_ID"))
// 	lp, _ := NewLongpoll(&vk, groupID)
// 	lp.Wait = 1

// 	t.Run("Run", func(t *testing.T) {
//
// 	})
// }

func TestLongpoll_RunError(t *testing.T) {
	t.Parallel()

	vk := api.NewVK("")
	lp, _ := Init(vk, 0)
	lp.Wait = 1

	t.Run("Run client error", func(t *testing.T) {
		if err := lp.Run(); err == nil {
			t.Error(err)
		}
	})

	lp.Server = "http://example.com"

	t.Run("Run json error", func(t *testing.T) {
		if err := lp.Run(); err == nil {
			t.Error(err)
		}
	})
}
