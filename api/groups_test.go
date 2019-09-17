package api

import (
	"os"
	"reflect"
	"testing"
)

func TestVK_GroupsAddCallbackServer(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddCallbackServerResponse
		wantErr      bool
	}{
		{
			name:    "groups.addCallbackServer error",
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.GroupsAddCallbackServer(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsAddCallbackServer() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsAddCallbackServer() err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsDeleteAddress(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name      string
		argParams map[string]string
		wantErr   bool
	}{
		{
			name:    "groups.deleteAddress error",
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := vkGroup.GroupsDeleteAddress(tt.argParams); (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsDeleteAddress() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsDeleteCallbackServer(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name      string
		argParams map[string]string
		wantErr   bool
	}{
		{
			name:    "groups.deleteCallbackServer error",
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := vkGroup.GroupsDeleteCallbackServer(tt.argParams); (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsDeleteCallbackServer() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsEditCallbackServer(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name      string
		argParams map[string]string
		wantErr   bool
	}{
		{
			name:    "groups.editCallbackServer error",
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := vkGroup.GroupsEditCallbackServer(tt.argParams); (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsEditCallbackServer() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
func TestVK_GroupsEnableOnline(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name      string
		argParams map[string]string
		wantErr   bool
	}{
		{
			name: "groups.enableOnline",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:    "groups.enableOnline error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := vkGroup.GroupsEnableOnline(tt.argParams); (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsEnableOnline() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsDisableOnline(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name      string
		argParams map[string]string
		wantErr   bool
	}{
		{
			name: "groups.disableOnline",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:    "groups.disableOnline error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := vkGroup.GroupsDisableOnline(tt.argParams); (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsDisableOnline() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
func TestVK_GroupsGetBanned(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetBannedResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.GroupsGetBanned(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetBanned() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsGetBanned() err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsGetByID(t *testing.T) {
	needServiceToken(t)

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse GroupsGetByIDResponse
		wantErr bool
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
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vkService.GroupsGetByID(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsGetByID() err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsGetCallbackConfirmationCode(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse GroupsGetCallbackConfirmationCodeResponse
		wantErr bool
	}{
		{
			name: "groups.getCallbackConfirmationCode",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:    "groups.getCallbackConfirmationCode error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vkGroup.GroupsGetCallbackConfirmationCode(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsGetCallbackConfirmationCode() err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsGetCallbackServers(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name      string
		argParams map[string]string
		wantErr   bool
	}{
		{
			name: "groups.getCallbackSettings",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:    "groups.getCallbackSettings error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.GroupsGetCallbackServers(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsGetCallbackServers() err = %v, want %v", err, tt.wantErr)
			}
			if err != nil && (gotResponse.Count != len(gotResponse.Items)) {
				t.Errorf("VK.GroupsGetCallbackServers() gotResponse = %v", gotResponse)
			}
		})
	}
}

func TestVK_GroupsGetCallbackSettings(t *testing.T) {
	needGroupToken(t)
	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetCallbackSettingsResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.GroupsGetCallbackSettings(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetCallbackSettings() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsGetCallbackSettings() err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsGetLongPollServer(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name      string
		argParams map[string]string
		wantErr   bool
	}{
		{
			name: "groups.getLongPollServer enabled",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:    "groups.getLongPollServer error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vkGroup.GroupsGetLongPollServer(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsGetLongPollServer() err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsGetLongPollSettings(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name      string
		argParams map[string]string
		wantErr   bool
	}{
		{
			name: "groups.getLongPollSettings enabled",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:    "groups.getLongPollSettings error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vkGroup.GroupsGetLongPollSettings(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsGetLongPollSettings() err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsGetMembers(t *testing.T) {
	needGroupToken(t)
	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetMembersResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.GroupsGetMembers(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetMembers() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsGetMembers() err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsGetMembersFilterManagers(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetMembersFilterManagersResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.GroupsGetMembersFilterManagers(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetMembersFilterManagers() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsGetMembersFilterManagers() err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsGetOnlineStatus(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetOnlineStatusResponse
		wantErr      bool
	}{
		{
			name: "groups.getOnlineStatus enabled",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
			wantResponse: GroupsGetOnlineStatusResponse{
				Status: "none",
			},
		},
		{
			name:    "groups.getOnlineStatus error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.GroupsGetOnlineStatus(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetOnlineStatus() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsGetOnlineStatus() err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsGetTokenPermissions(t *testing.T) {
	needGroupToken(t)
	t.Run("GroupsGetTokenPermissions", func(t *testing.T) {
		gotResponse, err := vkGroup.GroupsGetTokenPermissions(map[string]string{})
		if gotResponse.Mask == 0 {
			t.Errorf("VK.GroupsGetTokenPermissions() gotResponse = %v", gotResponse)
		}
		if err != nil {
			t.Errorf("VK.GroupsGetTokenPermissions() err = %v", err)
		}
	})

}

func TestVK_GroupsSetCallbackSettings(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name      string
		argParams map[string]string
		wantErr   bool
	}{
		{
			name:    "groups.setCallbackSettings error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := vkGroup.GroupsSetCallbackSettings(tt.argParams); (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsSetCallbackSettings() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsSetLongPollSettings(t *testing.T) {
	needGroupToken(t)
	tests := []struct {
		name      string
		argParams map[string]string
		wantErr   bool
	}{
		{
			name: "groups.setLongPollSettings enabled",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
				"enabled":  "1",
			},
		},
		{
			name:    "groups.setLongPollSettings error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := vkGroup.GroupsSetLongPollSettings(tt.argParams); (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsSetLongPollSettings() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsGetAddresses(t *testing.T) {
	needServiceToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetAddressesResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name:    "groups.getAddresses error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkService.GroupsGetAddresses(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetAddresses() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsGetAddresses() err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsEditAddress(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsEditAddressResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name:    "groups.editAddresses error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.GroupsEditAddress(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsEditAddress() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsEditAddress() err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsAddAddress(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddAddressResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name:    "groups.addAddresses error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.GroupsAddAddress(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsAddAddress() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsAddAddress() err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestVK_GroupsIsMember(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse int
		wantErr      bool
	}{
		{
			name: "groups.isMembers",
			argParams: map[string]string{
				"group_id": "1",
				"user_id":  "117253521",
			},
			wantResponse: 1,
		},
		{
			name:      "groups.isMembers error",
			argParams: map[string]string{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.GroupsIsMember(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsIsMember() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsIsMember() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_GroupsIsMemberExtended(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsIsMemberExtendedResponse
		wantErr      bool
	}{
		{
			name: "groups.isMembers",
			argParams: map[string]string{
				"group_id": "1",
				"user_id":  "117253521",
			},
			wantResponse: GroupsIsMemberExtendedResponse{
				Member:    1,
				CanInvite: 0,
			},
		},
		{
			name:      "groups.isMembers error",
			argParams: map[string]string{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.GroupsIsMemberExtended(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsIsMemberExtended() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsIsMemberExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_GroupsIsMemberUserIDsExtended(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsIsMemberUserIDsExtendedResponse
		wantErr      bool
	}{
		{
			name: "groups.isMembers",
			argParams: map[string]string{
				"group_id": "1",
				"user_ids": "117253521",
			},
			wantResponse: GroupsIsMemberUserIDsExtendedResponse{
				{
					Member:    1,
					CanInvite: 0,
					UserID:    117253521,
				},
			},
		},
		{
			name:      "groups.isMembers error",
			argParams: map[string]string{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.GroupsIsMemberUserIDsExtended(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsIsMemberUserIDsExtended() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsIsMemberUserIDsExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_GroupsIsMemberUserIDs(t *testing.T) {
	needGroupToken(t)

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsIsMemberUserIDsResponse
		wantErr      bool
	}{
		{
			name: "groups.isMembers",
			argParams: map[string]string{
				"group_id": "1",
				"user_ids": "117253521",
			},
			wantResponse: GroupsIsMemberUserIDsResponse{
				{
					Member: 1,
					UserID: 117253521,
				},
			},
		},
		{
			name:      "groups.isMembers error",
			argParams: map[string]string{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkGroup.GroupsIsMemberUserIDs(tt.argParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.GroupsIsMemberUserIDs() err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsIsMemberUserIDs() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
