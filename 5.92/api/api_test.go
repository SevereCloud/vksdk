package api

import (
	"os"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name   string
		args   args
		wantVk VK
	}{
		{
			name: "init test",
			args: args{
				token: "test",
			},
			wantVk: VK{
				AccessToken: "test",
				Version:     "5.92",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVk := Init(tt.args.token); !reflect.DeepEqual(gotVk, tt.wantVk) {
				t.Errorf("Init() = %v, want %v", gotVk, tt.wantVk)
			}
		})
	}
}

func TestVK_Request(t *testing.T) {
	type fields struct {
		AccessToken  string
		Version      string
		ProxyAddress string
	}
	type args struct {
		method string
		params map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
		want1  Error
	}{
		{
			name: "Request proxy error",
			fields: fields{
				ProxyAddress: "error",
			},
			want1: Error{
				Code: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vk := VK{
				AccessToken:  tt.fields.AccessToken,
				Version:      tt.fields.Version,
				ProxyAddress: tt.fields.ProxyAddress,
			}
			got, got1 := vk.Request(tt.args.method, tt.args.params)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VK.Request() got = %v, want %v", got, tt.want)
			}
			if got1.Code != tt.want1.Code {
				t.Errorf("VK.Request() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestVK_Execute(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)

	type args struct {
		Code string
	}
	tests := []struct {
		name         string
		argCode      string
		wantResponse []byte
		wantVkErr    Error
	}{
		{
			name:         "execute test",
			argCode:      `return 1;`,
			wantResponse: []byte(`1`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vk.Execute(tt.argCode)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.Execute() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if !reflect.DeepEqual(gotVkErr, tt.wantVkErr) {
				t.Errorf("VK.Execute() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}
