package api_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/stretchr/testify/assert"
)

const sleepTime = time.Second

func noError(t *testing.T, err error) bool {
	t.Helper()

	var e *api.Error
	if errors.As(err, &e) {
		switch e.Code {
		case api.ErrServer:
			t.Skip("Internal server error")
		case api.ErrPermission:
			t.Skip("Permission to perform this action is denied")
		case api.ErrCaptcha:
			t.Skip("Captcha needed")
		case api.ErrFlood:
			t.Skip("Flood control")
		default:
			s := "\n"
			s += fmt.Sprintf("code: %d\n", e.Code)
			s += fmt.Sprintf("text: %s\n", e.Text)
			s += fmt.Sprintf("message: %s\n", e.Message)

			if e.RedirectURI != "" {
				s += fmt.Sprintf("redirect_uri: %s\n", e.RedirectURI)
			}

			s += "params:\n"

			for _, param := range e.RequestParams {
				s += fmt.Sprintf("\t%s: %s\n", param.Key, param.Value)
			}

			t.Log(s)
		}
	} else if err != nil {
		t.Log(fmt.Sprintf("\n%#v", err))
	}

	return assert.NoError(t, err)
}

func noErrorOrFail(t *testing.T, err error) {
	t.Helper()

	if !noError(t, err) {
		t.FailNow()
	}
}

func needUserToken(t *testing.T) {
	t.Helper()

	if vkUser == nil {
		t.Skip("USER_TOKEN empty")
	}
}

func needGroupToken(t *testing.T) {
	t.Helper()

	if vkGroup == nil {
		t.Skip("GROUP_TOKEN empty")
	}
}

func needServiceToken(t *testing.T) {
	t.Helper()

	if vkService == nil {
		t.Skip("SERVICE_TOKEN empty")
	}
}

func needWidgetToken(t *testing.T) {
	t.Helper()

	if vkWidget == nil {
		t.Skip("WIDGET_TOKEN empty")
	}
}

func needAccountID(t *testing.T) int {
	t.Helper()

	if vkAccountID == 0 {
		t.Skip("ACCOUNT_ID empty")
	}

	return vkAccountID
}

func needChatID(t *testing.T) int {
	t.Helper()

	mux.Lock()
	defer mux.Unlock()

	needUserToken(t)

	if vkChatID == 0 {
		var err error

		vkChatID, err = vkUser.MessagesCreateChat(api.Params{
			"title": "TestChat",
		})
		if err != nil {
			t.Skip("Get chatID", err)
		}
	}

	return vkChatID
}

func needMarusiaToken(t *testing.T) {
	t.Helper()

	if vkMarusia == nil {
		t.Skip("MARUSIA_TOKEN empty")
	}
}

var (
	vkGroup, vkService, vkUser, vkWidget, vkMarusia *api.VK    // nolint:gochecknoglobals
	vkUserID, vkGroupID, vkChatID, vkAccountID      int        // nolint:gochecknoglobals
	mux                                             sync.Mutex // nolint:gochecknoglobals
)

func TestMain(m *testing.M) {
	if token := os.Getenv("GROUP_TOKEN"); token != "" {
		vkGroup = api.NewVK(token)

		group, err := vkGroup.GroupsGetByID(nil)
		if err != nil {
			log.Fatalf("GROUP_TOKEN bad: %v", err)
		}

		vkGroupID = group[0].ID
	}

	if token := os.Getenv("WIDGET_TOKEN"); token != "" {
		vkWidget = api.NewVK(token)
	}

	if token := os.Getenv("MARUSIA_TOKEN"); token != "" {
		vkMarusia = api.NewVK(token)
	}

	if token := os.Getenv("SERVICE_TOKEN"); token != "" {
		vkService = api.NewVK(token)
		vkService.Limit = api.LimitUserToken
	}

	if token := os.Getenv("USER_TOKEN"); token != "" {
		vkUser = api.NewVK(token)
		vkUser.Limit = api.LimitUserToken

		user, err := vkUser.UsersGet(nil)
		if err != nil {
			log.Fatalf("USER_TOKEN bad: %v", err)
		}

		vkUserID = user[0].ID
	}

	vkAccountID, _ = strconv.Atoi(os.Getenv("ACCOUNT_ID"))

	runTests := m.Run()
	os.Exit(runTests)
}

func TestVK_Request(t *testing.T) {
	t.Parallel()

	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	vk := api.NewVK(groupToken)

	_, err := vk.Request("", nil)
	if err == nil {
		t.Errorf("VK.Request() got1 = %v, want -1", err)
	}

	vk.MethodURL = ""

	_, err = vk.Request("test", api.Params{"test": "test"})
	if err == nil {
		t.Errorf("VK.Request() got1 = %v, want -1", err)
	}
}

func TestVK_RequestLimit(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func() {
			_, err := vkUser.UsersGet(nil)
			assert.NoError(t, err)

			wg.Done()
		}()
	}

	wg.Wait()

	vkUser.Limit = 3
}

func TestVK_InvalidContentType(t *testing.T) {
	t.Parallel()

	var testObj string

	vk := api.NewVK("")
	vk.MethodURL = "https://api.vk.com"

	err := vk.RequestUnmarshal("", testObj, nil)
	if err == nil || err.Error() != "api: invalid content-type" {
		t.Errorf("VK.RequestUnmarshal() error = %v", err)
	}
}

type renamedBool bool

func Test_FmtValue(t *testing.T) {
	t.Parallel()

	f := func(value interface{}, want string) {
		t.Helper()

		got := api.FmtValue(value, 0)
		assert.Equal(t, got, want)
	}

	f(nil, "")
	f(reflect.ValueOf(nil), "")
	f(reflect.Value{}, "")

	f(true, "1")
	f(false, "0")
	f(renamedBool(true), "1")

	f(123, "123")
	f(1.1, "1.1")
	f("abc", "abc")

	// Attachment
	photo := object.PhotosPhoto{
		OwnerID: 321,
		ID:      123,
	}
	f(photo, "photo321_123")

	// Keyboard
	keyboard := object.NewMessagesKeyboard(true)
	f(keyboard, keyboard.ToJSON())

	// Slice
	intSlice := []int{1, 2, 3}
	f(intSlice, "1,2,3")
	f([]object.PhotosPhoto{photo, photo}, "photo321_123,photo321_123")

	// Pointer
	f(&intSlice, "1,2,3")
	f(&photo, "photo321_123")
}

func TestParams_methods(t *testing.T) {
	t.Parallel()

	p := api.Params{}

	p.Lang(1)
	p.TestMode(true)
	p.CaptchaSID("text")
	p.CaptchaKey("text")
	p.Confirm(true)

	assert.Equal(t, p["lang"], 1)
	assert.Equal(t, p["test_mode"], true)
	assert.Equal(t, p["captcha_sid"], "text")
	assert.Equal(t, p["captcha_key"], "text")
	assert.Equal(t, p["confirm"], true)
}

func TestContext(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	p := api.Params{}
	p.WithContext(ctx)

	_, err := vkUser.UsersGet(p)
	assert.EqualError(t, err, "Post \"https://api.vk.com/method/users.get\": context deadline exceeded")
}
