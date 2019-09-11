package api

import (
	"reflect"
	"testing"
)

func TestVK_OrdersCancelSubscription(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse int
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name:    "OrdersCancelSubscription error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkService.OrdersCancelSubscription(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.OrdersCancelSubscription() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.OrdersCancelSubscription() err = %v, want %v", err, tt.wantErr)
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
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name:    "OrdersChangeState error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkService.OrdersChangeState(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.OrdersChangeState() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.OrdersChangeState() err = %v, want %v", err, tt.wantErr)
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
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name:         "OrdersGet empty",
			wantResponse: OrdersGetResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkService.OrdersGet(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.OrdersGet() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.OrdersGet() err = %v, want %v", err, tt.wantErr)
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
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkUser.OrdersGetAmount(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.OrdersGetAmount() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.OrdersGetAmount() err = %v, want %v", err, tt.wantErr)
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
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name:    "OrdersGetByID error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkService.OrdersGetByID(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.OrdersGetByID() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.OrdersGetByID() err = %v, want %v", err, tt.wantErr)
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
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name:    "OrdersGetUserSubscriptionByID error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkService.OrdersGetUserSubscriptionByID(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.OrdersGetUserSubscriptionByID() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.OrdersGetUserSubscriptionByID() err = %v, want %v", err, tt.wantErr)
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
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name:    "OrdersGetUserSubscriptions error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkService.OrdersGetUserSubscriptions(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.OrdersGetUserSubscriptions() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.OrdersGetUserSubscriptions() err = %v, want %v", err, tt.wantErr)
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
		wantResponse int
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name:    "OrdersUpdateSubscription error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := vkService.OrdersUpdateSubscription(tt.argParams)
			if gotResponse != tt.wantResponse {
				t.Errorf("VK.OrdersUpdateSubscription() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("VK.OrdersUpdateSubscription() err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
