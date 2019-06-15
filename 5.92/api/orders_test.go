package api

import (
	"reflect"
	"testing"

	"github.com/SevereCloud/vksdk/5.92/object"
)

func TestVK_OrdersCancelSubscription(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse OrdersCancelSubscriptionResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "OrdersCancelSubscription error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkService.OrdersCancelSubscription(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.OrdersCancelSubscription() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.OrdersCancelSubscription() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_OrdersChangeState(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse OrdersChangeStateResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "OrdersChangeState error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkService.OrdersChangeState(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.OrdersChangeState() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.OrdersChangeState() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_OrdersGet(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse OrdersGetResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:         "OrdersGet empty",
			wantResponse: OrdersGetResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkService.OrdersGet(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.OrdersGet() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.OrdersGet() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_OrdersGetAmount(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse OrdersGetAmountResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkUser.OrdersGetAmount(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.OrdersGetAmount() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.OrdersGetAmount() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_OrdersGetByID(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse OrdersGetByIDResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "OrdersGetByID error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkService.OrdersGetByID(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.OrdersGetByID() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.OrdersGetByID() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_OrdersGetUserSubscriptionByID(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse OrdersGetUserSubscriptionByIDResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "OrdersGetUserSubscriptionByID error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkService.OrdersGetUserSubscriptionByID(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.OrdersGetUserSubscriptionByID() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.OrdersGetUserSubscriptionByID() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_OrdersGetUserSubscriptions(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse OrdersGetUserSubscriptionsResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "OrdersGetUserSubscriptions error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkService.OrdersGetUserSubscriptions(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.OrdersGetUserSubscriptions() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.OrdersGetUserSubscriptions() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_OrdersUpdateSubscription(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse OrdersUpdateSubscriptionResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "OrdersUpdateSubscription error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkService.OrdersUpdateSubscription(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.OrdersUpdateSubscription() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.OrdersUpdateSubscription() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}
