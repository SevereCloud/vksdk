package api_test

import (
	"log"
	"os"
	"reflect"
	"sync"
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/api/errors"
	"github.com/SevereCloud/vksdk/object"
	"github.com/stretchr/testify/assert"
)

func noError(t *testing.T, err error) bool {
	t.Helper()

	switch errors.GetType(err) {
	case errors.TooMany:
		t.Skip("Too many requests per second")
	case errors.Server:
		t.Skip("Internal server error")
	case errors.Permission:
		t.Skip("Permission to perform this action is denied")
	case errors.Captcha:
		t.Skip("Captcha needed")
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

func needChatID(t *testing.T) int {
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

var vkGroup, vkService, vkUser *api.VK // nolint:gochecknoglobals
var vkUserID, vkGroupID, vkChatID int  // nolint:gochecknoglobals

func TestMain(m *testing.M) {
	vkGroup = api.Init(os.Getenv("GROUP_TOKEN"))
	if vkGroup.AccessToken != "" {
		group, err := vkGroup.GroupsGetByID(api.Params{})
		if err != nil {
			log.Fatalf("GROUP_TOKEN bad: %v", err)
		}

		vkGroupID = group[0].ID
	}

	vkService = api.Init(os.Getenv("SERVICE_TOKEN"))
	vkService.Limit = 3
	vkUser = api.Init(os.Getenv("USER_TOKEN"))
	vkUser.Limit = 3

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
	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	vk := api.Init(groupToken)

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
}

func TestVK_Execute_error(t *testing.T) {
	needGroupToken(t)

	var response int

	err := vkGroup.Execute(`API.users.get({user_id:-1});return 1;`, &response)
	assert.Error(t, err)
	assert.Equal(t, 1, response)
}

func TestVK_Execute_object(t *testing.T) {
	needGroupToken(t)

	var response struct {
		Text string `json:"text"`
	}

	err := vkGroup.Execute(`return {text: "hello"};`, &response)
	assert.NoError(t, err)
	assert.Equal(t, "hello", response.Text)
}

func TestVK_RequestUnmarshal(t *testing.T) {
	needGroupToken(t)

	var testObj string

	err := vkGroup.RequestUnmarshal("t/t", api.Params{}, testObj)
	if err == nil {
		t.Errorf("VK.RequestUnmarshal() error = %v", err)
	}
}

type renamedBool bool

func Test_FmtValue(t *testing.T) {
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
	needUserToken(t)

	_, err := vkUser.CaptchaForce(api.Params{})
	if errors.GetType(err) != errors.Captcha {
		t.Errorf("VK.CaptchaForce() err=%v, want 14", err)
	}
}
