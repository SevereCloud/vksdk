package api

import (
	"os"
	"testing"
)

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
		_, vkErr := vk.Request("", map[string]string{})
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
	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(groupToken)

	t.Run("Execute test", func(t *testing.T) {
		var response int
		var vkErr Error
		vk.Execute(`return 1;`, &response, &vkErr)
		if vkErr.Code != 0 {
			t.Errorf("VK.Execute() gotVkErr = %v, want 0", vkErr)
		}
		if response != 1 {
			t.Error("Execute response error")
		}
	})
}

func TestVK_RequestUnmarshal(t *testing.T) {
	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(groupToken)

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
			vk.RequestUnmarshal(tt.args.method, tt.args.params, tt.args.obj, tt.args.vkErr)
			if tt.args.vkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.Execute() gotVkErr = %v, want %v", tt.args.vkErr, tt.wantVkErr)
			}
		})
	}
}
