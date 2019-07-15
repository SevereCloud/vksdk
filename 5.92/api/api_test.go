package api

import (
	"log"
	"os"
	"testing"
)

var vkGroup, vkService, vkUser VK // nolint:gochecknoglobals
var vkUserID, vkGroupID int       // nolint:gochecknoglobals

func TestMain(m *testing.M) {
	vkGroup = Init(os.Getenv("GROUP_TOKEN"))
	if vkGroup.AccessToken != "" {
		group, vkErr := vkGroup.GroupsGetByID(map[string]string{})
		if vkErr.Code != 0 {
			log.Fatal(vkErr)
		}
		vkGroupID = group[0].ID
	}
	vkService = Init(os.Getenv("SERVICE_TOKEN"))
	vkUser = Init(os.Getenv("USER_TOKEN"))
	vkUser.Limit = 3
	if vkUser.AccessToken != "" {
		user, vkErr := vkUser.UsersGet(map[string]string{})
		if vkErr.Code != 0 {
			log.Fatal(vkErr)
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
	vk := Init(groupToken)

	t.Run("Request 403 error", func(t *testing.T) {
		_, vkErr := vk.Request("", map[string]string{})
		if vkErr.Code != -1 {
			t.Errorf("VK.Request() got1 = %v, want -1", vkErr)
		}
	})
	vk.MethodURL = ""
	t.Run("Client error", func(t *testing.T) {
		_, vkErr := vk.Request("test", map[string]string{"test": "test"})
		if vkErr.Code != -1 {
			t.Errorf("VK.Request() got1 = %v, want -1", vkErr)
		}
	})
}

func TestVK_RequestLimit(t *testing.T) {
	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(groupToken)
	vk.Limit = 2

	t.Run("vk.Limit", func(t *testing.T) {
		go vk.UsersGet(map[string]string{})
		for i := 0; i < 2; i++ {
			vk.UsersGet(map[string]string{})
		}
	})

}

func TestVK_Execute(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	t.Run("Execute test", func(t *testing.T) {
		var response int
		var vkErr Error
		vkGroup.Execute(`return 1;`, &response, &vkErr)
		if vkErr.Code != 0 {
			t.Errorf("VK.Execute() gotVkErr = %v, want 0", vkErr)
		}
		if response != 1 {
			t.Error("Execute response error")
		}
	})
}

func TestVK_RequestUnmarshal(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	var testObj string
	type args struct {
		method string
		params map[string]string
		obj    interface{}
		vkErr  *Error
	}
	tests := []struct {
		name      string
		args      args
		wantVkErr Error
	}{
		{
			name: "execute error",
			args: args{
				method: "execute",
				params: map[string]string{"code": "return 1;"},
				obj:    &testObj,
				vkErr:  &Error{},
			},
			wantVkErr: Error{Code: -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vkGroup.RequestUnmarshal(tt.args.method, tt.args.params, tt.args.obj, tt.args.vkErr)
			if tt.args.vkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.Execute() gotVkErr = %v, want %v", tt.args.vkErr, tt.wantVkErr)
			}
		})
	}
}
