package api_test

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"sync"
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/object"
	"github.com/stretchr/testify/assert"
)

func noError(t *testing.T, err error) bool {
	t.Helper()

	if e, ok := err.(*api.Error); ok {
		switch e.Code {
		case api.ErrServer:
			t.Skip("Internal server error")
		case api.ErrPermission:
			t.Skip("Permission to perform this action is denied")
		case api.ErrCaptcha:
			t.Skip("Captcha needed")
		default:
			s := "\n"
			s += fmt.Sprintf("code: %d\n", e.Code)
			s += fmt.Sprintf("text: %s\n", e.Text)
			s += fmt.Sprintf("message: %s\n", e.Message)
			s += "params:\n"

			for _, param := range e.RequestParams {
				s += fmt.Sprintf("\t%s: %s\n", param.Key, param.Value)
			}

			t.Log(s)
		}
	} else {
		t.Log(fmt.Sprintf("\n%#v", err))
	}

	return assert.NoError(t, err)
}

func needUserToken(t *testing.T) {
	t.Helper()

	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}
}

func needGroupToken(t *testing.T) {
	t.Helper()

	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
}

func needServiceToken(t *testing.T) {
	t.Helper()

	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
}

func needWidgetToken(t *testing.T) {
	t.Helper()

	if vkWidget.AccessToken == "" {
		t.Skip("WIDGET_TOKEN empty")
	}
}

func needChatID(t *testing.T) int {
	mux.Lock()
	defer mux.Unlock()

	t.Helper()
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

var (
	vkGroup, vkService, vkUser, vkWidget *api.VK    // nolint:gochecknoglobals
	vkUserID, vkGroupID, vkChatID        int        // nolint:gochecknoglobals
	mux                                  sync.Mutex // nolint:gochecknoglobals
)

func TestMain(m *testing.M) {
	vkGroup = api.NewVK(os.Getenv("GROUP_TOKEN"))
	if vkGroup.AccessToken != "" {
		group, err := vkGroup.GroupsGetByID(api.Params{})
		if err != nil {
			log.Fatalf("GROUP_TOKEN bad: %v", err)
		}

		vkGroupID = group[0].ID
	}

	vkWidget = api.NewVK(os.Getenv("WIDGET_TOKEN"))
	vkService = api.NewVK(os.Getenv("SERVICE_TOKEN"))
	vkService.Limit = api.LimitUserToken
	vkUser = api.NewVK(os.Getenv("USER_TOKEN"))
	vkUser.Limit = api.LimitUserToken

	if vkUser.AccessToken != "" {
		user, err := vkUser.UsersGet(api.Params{})
		if err != nil {
			log.Fatalf("USER_TOKEN bad: %v", err)
		}

		vkUserID = user[0].ID
	}

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

	t.Run("Request 403 error", func(t *testing.T) {
		_, err := vk.Request("", api.Params{})
		if err == nil {
			t.Errorf("VK.Request() got1 = %v, want -1", err)
		}
	})

	vk.MethodURL = ""

	t.Run("Client error", func(t *testing.T) {
		_, err := vk.Request("test", api.Params{"test": "test"})
		if err == nil {
			t.Errorf("VK.Request() got1 = %v, want -1", err)
		}
	})
}

func TestVK_RequestLimit(t *testing.T) {
	needUserToken(t)

	vkUser.Limit = 4

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func() {
			_, err := vkUser.UsersGet(api.Params{})
			assert.NoError(t, err)

			wg.Done()
		}()
	}

	wg.Wait()

	vkUser.Limit = 3
}

func TestVK_Execute_error(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	var response int

	err := vkGroup.Execute(`API.users.get({user_id:-1});return 1;`, &response)
	assert.Error(t, err)
	assert.Equal(t, 1, response)
}

func TestVK_Execute_object(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	var response struct {
		Text string `json:"text"`
	}

	err := vkGroup.Execute(`return {text: "hello"};`, &response)
	assert.NoError(t, err)
	assert.Equal(t, "hello", response.Text)
}

func TestVK_ExecuteWithArgs_error(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	var response int

	err := vkGroup.ExecuteWithArgs(`API.users.get({user_id: parseInt(Args.user_id)});return 1;`, api.Params{"user_id": -1}, &response)
	assert.Error(t, err)
	assert.Equal(t, 1, response)
}

func TestVK_ExecuteWithArgs_object(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	var response struct {
		Text string `json:"text"`
	}

	err := vkGroup.ExecuteWithArgs(`return {text: Args.text};`, api.Params{"text": "hello"}, &response)
	assert.NoError(t, err)
	assert.Equal(t, "hello", response.Text)
}

func TestVK_InvalidContentType(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	var testObj string

	err := vkGroup.RequestUnmarshal("t/t", api.Params{}, testObj)
	if err == nil || err.Error() != "invalid content-type" {
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

func TestVK_CaptchaForce(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.CaptchaForce(api.Params{})

	if e, ok := err.(*api.Error); err != nil || !ok || e.Code != api.ErrCaptcha {
		t.Errorf("VK.CaptchaForce() err=%v, want 14", err)
	}
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
