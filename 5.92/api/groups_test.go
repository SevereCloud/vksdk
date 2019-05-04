package api

import (
	"os"
	"reflect"
	"testing"
)

func TestVK_GroupsAddCallbackServer(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddCallbackServerResponse
		wantVkErr    Error
	}{
		{
			name:      "groups.addCallbackServer error",
			wantVkErr: Error{Code: 100},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vk.GroupsAddCallbackServer(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsAddCallbackServer() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsAddCallbackServer() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsDeleteAddress(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name:      "groups.deleteAddress error",
			wantVkErr: Error{Code: 100},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVkErr := vk.GroupsDeleteAddress(tt.argParams); gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsDeleteAddress() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsDeleteCallbackServer(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name:      "groups.deleteCallbackServer error",
			wantVkErr: Error{Code: 100},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVkErr := vk.GroupsDeleteCallbackServer(tt.argParams); gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsDeleteCallbackServer() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsDisableOnline(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name: "groups.disableOnline",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:      "groups.disableOnline error",
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVkErr := vk.GroupsDisableOnline(tt.argParams); gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsDisableOnline() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsEditCallbackServer(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name:      "groups.editCallbackServer error",
			wantVkErr: Error{Code: 100},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVkErr := vk.GroupsEditCallbackServer(tt.argParams); gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsEditCallbackServer() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}
func TestVK_GroupsEnableOnline(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name: "groups.enableOnline",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:      "groups.enableOnline error",
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVkErr := vk.GroupsEnableOnline(tt.argParams); gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsEnableOnline() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsGetBanned(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetBannedResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vk.GroupsGetBanned(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetBanned() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetBanned() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsGetByID(t *testing.T) {
	userToken := os.Getenv("SERVICE_TOKEN")
	if userToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(userToken)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetByIDResponse
		wantVkErr    Error
	}{
		{
			name: "groups.getById",
			argParams: map[string]string{
				"group_ids": "1",
			},
		},
		{
			name: "groups.getById error",
			argParams: map[string]string{
				"group_ids": "0",
			},
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vk.GroupsGetByID(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetByID() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if gotVkErr.Code == 0 && (gotResponse[0].ID != 1) {
				t.Errorf("VK.GroupsGetByID() gotResponse = %v", gotResponse)
			}
		})
	}
}

func TestVK_GroupsGetCallbackConfirmationCode(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetCallbackConfirmationCodeResponse
		wantVkErr    Error
	}{
		{
			name: "groups.getCallbackConfirmationCode",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:      "groups.getCallbackConfirmationCode error",
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vk.GroupsGetCallbackConfirmationCode(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetCallbackConfirmationCode() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if gotVkErr.Code == 0 && (gotResponse.Code == "") {
				t.Errorf("VK.GroupsGetCallbackConfirmationCode() gotResponse = %v", gotResponse)
			}
		})
	}
}

func TestVK_GroupsGetCallbackServers(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name: "groups.getCallbackSettings",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:      "groups.getCallbackSettings error",
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vk.GroupsGetCallbackServers(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetCallbackServers() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if gotVkErr.Code == 0 && (gotResponse.Count != len(gotResponse.Items)) {
				t.Errorf("VK.GroupsGetCallbackServers() gotResponse = %v", gotResponse)
			}
		})
	}
}

func TestVK_GroupsGetCallbackSettings(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)
	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetCallbackSettingsResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vk.GroupsGetCallbackSettings(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetCallbackSettings() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetCallbackSettings() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsGetLongPollServer(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name: "groups.getLongPollServer enabled",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:      "groups.getLongPollServer error",
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vk.GroupsGetLongPollServer(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetLongPollServer() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if gotVkErr.Code == 0 && (gotResponse.Key == "" || gotResponse.Server == "" || gotResponse.Ts == "") {
				t.Errorf("VK.GroupsGetLongPollServer() gotResponse = %v", gotResponse)
			}
		})
	}
}

func TestVK_GroupsGetLongPollSettings(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name: "groups.getLongPollSettings enabled",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:      "groups.getLongPollSettings error",
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vk.GroupsGetLongPollSettings(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetLongPollSettings() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsGetMembers(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)
	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetMembersResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vk.GroupsGetMembers(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetMembers() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetMembers() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsGetMembersFilterManagers(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetMembersFilterManagersResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vk.GroupsGetMembersFilterManagers(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetMembersFilterManagers() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetMembersFilterManagers() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsGetOnlineStatus(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetOnlineStatusResponse
		wantVkErr    Error
	}{
		{
			name: "groups.getOnlineStatus enabled",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
			wantResponse: GroupsGetOnlineStatusResponse{
				Status: "online",
			},
		},
		{
			name:      "groups.getOnlineStatus error",
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vk.GroupsGetOnlineStatus(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetOnlineStatus() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetOnlineStatus() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsGetTokenPermissions(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)

	gotResponse, gotVkErr := vk.GroupsGetTokenPermissions()
	if gotResponse.Mask == 0 {
		t.Errorf("VK.GroupsGetTokenPermissions() gotResponse = %v", gotResponse)
	}
	if gotVkErr.Code != 0 {
		t.Errorf("VK.GroupsGetTokenPermissions() gotVkErr = %v", gotVkErr)
	}
	vk = Init("")
	_, vkErr := vk.GroupsGetTokenPermissions()
	if vkErr.Code != 5 {
		t.Errorf("GroupsGetTokenPermissions error bad %d %s", vkErr.Code, vkErr.Message)
	}
}

func TestVK_GroupsSetCallbackSettings(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name:      "groups.setCallbackSettings error",
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVkErr := vk.GroupsSetCallbackSettings(tt.argParams); gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsSetCallbackSettings() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsSetLongPollSettings(t *testing.T) {
	userToken := os.Getenv("GROUP_TOKEN")
	if userToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := Init(userToken)
	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name: "groups.setLongPollSettings enabled",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
				"enabled":  "1",
			},
		},
		{
			name:      "groups.setLongPollSettings error",
			wantVkErr: Error{Code: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVkErr := vk.GroupsSetLongPollSettings(tt.argParams); gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsSetLongPollSettings() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}
