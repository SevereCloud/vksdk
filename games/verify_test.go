package games_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/SevereCloud/vksdk/v3/games"
)

func TestParamsVerify(t *testing.T) {
	t.Parallel()

	f := func(URL, clientSecret string, want, wantErr bool) {
		t.Helper()

		got, err := games.ParamsVerify(URL, clientSecret)
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

	// test for params with special characters in query ("%26")
	f(
		"https://example.com/?api_url=https://api.vk.ru/api.php&api_id=7869076&api_settings=0&viewer_id=117253521&viewer_type=0&sid=1a4c65d6a02dddd3e4428f1c0e3b913aee6f16074db1ae75e77b697e229e9ee38717bbc5515094a64611d&secret=148a2cffba&access_token=2f32ef1eee042c9b384f116bf6c4aca136386fcc5b8402f70091b3ffe61a824a937b8a986b5ed93f13a2d&user_id=0&is_app_user=1&language=0&parent_language=0&is_secure=1&stats_hash=2309491839c82a704b&group_id=0&ads_app_id=7869076_9ba760f1e19907aaa3&access_token_settings=&referrer=unknown&lc_name=d106d683&platform=web&is_widescreen=0&whitelist_scopes=friends,photos,video,stories,pages,status,notes,wall,docs,groups,stats,market,ads,notifications&group_whitelist_scopes=stories,photos,app_widget,messages,wall,docs,manage&auth_key=f3730660d049a2913b5e69c066b83b70&timestamp=1622546396&sign=lJeYjD3GiMBTqwqQGVyE-3zJyHWTS4XNvklIYGANps0&sign_keys=access_token,access_token_settings,ads_app_id,api_id,api_settings,api_url,auth_key,group_id,group_whitelist_scopes,is_app_user,is_secure,is_widescreen,language,lc_name,parent_language,platform,referrer,secret,sid,stats_hash,timestamp,user_id,viewer_id,viewer_type,whitelist_scopes&hash=",
		"TaWz5vVHUQnp7YgtuRlS",
		true,
		false,
	)

	f(
		"https://example.com/?api_url=https://api.vk.ru/api.php&api_id=7869076&api_settings=0&viewer_id=117253521&viewer_type=0&sid=1a4c65d6a02dddd3e4428f1c0e3b913aee6f16074db1ae75e77b697e229e9ee38717bbc5515094a64611d&secret=148a2cffba&access_token=3f32ef1eee042c9b384f116bf6c4aca136386fcc5b8402f70091b3ffe61a824a937b8a986b5ed93f13a2d&user_id=0&is_app_user=1&language=0&parent_language=0&is_secure=1&stats_hash=2309491839c82a704b&group_id=0&ads_app_id=7869076_9ba760f1e19907aaa3&access_token_settings=&referrer=unknown&lc_name=d106d683&platform=web&is_widescreen=0&whitelist_scopes=friends,photos,video,stories,pages,status,notes,wall,docs,groups,stats,market,ads,notifications&group_whitelist_scopes=stories,photos,app_widget,messages,wall,docs,manage&auth_key=f3730660d049a2913b5e69c066b83b70&timestamp=1622546396&sign=lJeYjD3GiMBTqwqQGVyE-3zJyHWTS4XNvklIYGANps0&sign_keys=access_token,access_token_settings,ads_app_id,api_id,api_settings,api_url,auth_key,group_id,group_whitelist_scopes,is_app_user,is_secure,is_widescreen,language,lc_name,parent_language,platform,referrer,secret,sid,stats_hash,timestamp,user_id,viewer_id,viewer_type,whitelist_scopes,bad&hash=",
		"TaWz5vVHUQnp7YgtuRlS",
		false,
		false,
	)
}

func TestParamsVerification_VerifyMiddleware(t *testing.T) {
	t.Parallel()

	pv := games.NewParamsVerification("TaWz5vVHUQnp7YgtuRlS")

	handlerOk := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
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
		"https://example.com/?api_url=https://api.vk.ru/api.php&api_id=7869076&api_settings=0&viewer_id=117253521&viewer_type=0&sid=1a4c65d6a02dddd3e4428f1c0e3b913aee6f16074db1ae75e77b697e229e9ee38717bbc5515094a64611d&secret=148a2cffba&access_token=3f32ef1eee042c9b384f116bf6c4aca136386fcc5b8402f70091b3ffe61a824a937b8a986b5ed93f13a2d&user_id=0&is_app_user=1&language=0&parent_language=0&is_secure=1&stats_hash=2309491839c82a704b&group_id=0&ads_app_id=7869076_9ba760f1e19907aaa3&access_token_settings=&referrer=unknown&lc_name=d106d683&platform=web&is_widescreen=0&whitelist_scopes=friends,photos,video,stories,pages,status,notes,wall,docs,groups,stats,market,ads,notifications&group_whitelist_scopes=stories,photos,app_widget,messages,wall,docs,manage&auth_key=f3730660d049a2913b5e69c066b83b70&timestamp=1622546396&sign=lJeYjD3GiMBTqwqQGVyE-3zJyHWTS4XNvklIYGANps0&sign_keys=access_token,access_token_settings,ads_app_id,api_id,api_settings,api_url,auth_key,group_id,group_whitelist_scopes,is_app_user,is_secure,is_widescreen,language,lc_name,parent_language,platform,referrer,secret,sid,stats_hash,timestamp,user_id,viewer_id,viewer_type,whitelist_scopes&hash=",
		403,
	)

	f(
		"https://example.com/?api_url=https://api.vk.ru/api.php&api_id=7869076&api_settings=0&viewer_id=117253521&viewer_type=0&sid=1a4c65d6a02dddd3e4428f1c0e3b913aee6f16074db1ae75e77b697e229e9ee38717bbc5515094a64611d&secret=148a2cffba&access_token=2f32ef1eee042c9b384f116bf6c4aca136386fcc5b8402f70091b3ffe61a824a937b8a986b5ed93f13a2d&user_id=0&is_app_user=1&language=0&parent_language=0&is_secure=1&stats_hash=2309491839c82a704b&group_id=0&ads_app_id=7869076_9ba760f1e19907aaa3&access_token_settings=&referrer=unknown&lc_name=d106d683&platform=web&is_widescreen=0&whitelist_scopes=friends,photos,video,stories,pages,status,notes,wall,docs,groups,stats,market,ads,notifications&group_whitelist_scopes=stories,photos,app_widget,messages,wall,docs,manage&auth_key=f3730660d049a2913b5e69c066b83b70&timestamp=1622546396&sign=lJeYjD3GiMBTqwqQGVyE-3zJyHWTS4XNvklIYGANps0&sign_keys=access_token,access_token_settings,ads_app_id,api_id,api_settings,api_url,auth_key,group_id,group_whitelist_scopes,is_app_user,is_secure,is_widescreen,language,lc_name,parent_language,platform,referrer,secret,sid,stats_hash,timestamp,user_id,viewer_id,viewer_type,whitelist_scopes&hash=",
		200,
	)
}
