package vkapps_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/SevereCloud/vksdk/vkapps"
)

func TestParamsVerify(t *testing.T) {
	f := func(URL, clientSecret string, want, wantErr bool) {
		t.Helper()

		got, err := vkapps.ParamsVerify(URL, clientSecret)
		if (err != nil) != wantErr {
			t.Errorf("ParamsValid() error = %v, wantErr %v", err, wantErr)
			return
		}

		if got != want {
			t.Errorf("ParamsValid() = %v, want %v", got, want)
		}
	}

	f("http://[%10::1]", "", false, true)
	f("https://example.com", "", false, false)
	f("https://example.com?sign=abc", "", false, false)
	f("https://example.com?q=abc", "", false, false)
	f("https://example.com?sign=abc&%gh&%ij", "", false, true)

	f(
		"https://example.com/?vk_user_id=494075&vk_app_id=6736218&vk_is_app_user=1&vk_are_notifications_enabled=1&vk_language=ru&vk_access_token_settings=notify&vk_platform=android&sign=exTIBPYTrAKDTHLLm2AwJkmcVcvFCzQUNyoa6wAjvW6k",
		"wvl68m4dR1UpLrVRli",
		false,
		false,
	)
	f(
		"https://example.com/?q=1&vk_user_id=494075&vk_app_id=6736218&vk_is_app_user=1&vk_are_notifications_enabled=1&vk_language=ru&vk_access_token_settings=&vk_platform=android&sign=htQFduJpLxz7ribXRZpDFUH-XEUhC9rBPTJkjUFEkRA",
		"wvl68m4dR1UpLrVRli",
		true,
		false,
	)
}

func TestParamsVerification_VerifyMiddleware(t *testing.T) {
	pv := vkapps.NewParamsVerification("wvl68m4dR1UpLrVRli")

	handlerOk := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	f := func(target string, wantStatusCode int) {
		t.Helper()

		handler := pv.VerifyMiddleware(handlerOk)

		u, _ := url.Parse(target)

		req := &http.Request{URL: u}
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != wantStatusCode {
			t.Errorf("StatusCode = %v, want %v", resp.StatusCode, wantStatusCode)
		}
	}

	f("example.com", 403)
	f("https://example.com?sign=abc&%h&%ij", 400)
	f(
		"https://example.com/?vk_user_id=494075&vk_app_id=6736218&vk_is_app_user=1&vk_are_notifications_enabled=1&vk_language=ru&vk_access_token_settings=notify&vk_platform=android&sign=exTIBPYTrAKDTHLLm2AwJkmcVcvFCzQUNyoa6wAjvW6k",
		403,
	)

	f(
		"https://example.com/?q=1&vk_user_id=494075&vk_app_id=6736218&vk_is_app_user=1&vk_are_notifications_enabled=1&vk_language=ru&vk_access_token_settings=&vk_platform=android&sign=htQFduJpLxz7ribXRZpDFUH-XEUhC9rBPTJkjUFEkRA",
		200,
	)
}
