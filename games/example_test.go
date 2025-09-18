package games_test

import (
	"fmt"
	"net/url"

	"github.com/SevereCloud/vksdk/v3/games"
)

func ExampleParamsVerification() {
	link := "https://example.com/?api_url=https://api.vk.ru/api.php&api_id=7869076&api_settings=0&viewer_id=117253521&viewer_type=0&sid=1a4c65d6a02dddd3e4428f1c0e3b913aee6f16074db1ae75e77b697e229e9ee38717bbc5515094a64611d&secret=148a2cffba&access_token=2f32ef1eee042c9b384f116bf6c4aca136386fcc5b8402f70091b3ffe61a824a937b8a986b5ed93f13a2d&user_id=0&is_app_user=1&language=0&parent_language=0&is_secure=1&stats_hash=2309491839c82a704b&group_id=0&ads_app_id=7869076_9ba760f1e19907aaa3&access_token_settings=&referrer=unknown&lc_name=d106d683&platform=web&is_widescreen=0&whitelist_scopes=friends,photos,video,stories,pages,status,notes,wall,docs,groups,stats,market,ads,notifications&group_whitelist_scopes=stories,photos,app_widget,messages,wall,docs,manage&auth_key=f3730660d049a2913b5e69c066b83b70&timestamp=1622546396&sign=lJeYjD3GiMBTqwqQGVyE-3zJyHWTS4XNvklIYGANps0&sign_keys=access_token,access_token_settings,ads_app_id,api_id,api_settings,api_url,auth_key,group_id,group_whitelist_scopes,is_app_user,is_secure,is_widescreen,language,lc_name,parent_language,platform,referrer,secret,sid,stats_hash,timestamp,user_id,viewer_id,viewer_type,whitelist_scopes&hash="

	pv := games.NewParamsVerification("TaWz5vVHUQnp7YgtuRlS") // Client secret

	u, _ := url.Parse(link)

	v, _ := pv.Verify(u)

	fmt.Println(v)

	// Output:
	// true
}

func ExampleParamsVerify() {
	link := "https://example.com/?api_url=https://api.vk.ru/api.php&api_id=7869076&api_settings=0&viewer_id=117253521&viewer_type=0&sid=1a4c65d6a02dddd3e4428f1c0e3b913aee6f16074db1ae75e77b697e229e9ee38717bbc5515094a64611d&secret=148a2cffba&access_token=2f32ef1eee042c9b384f116bf6c4aca136386fcc5b8402f70091b3ffe61a824a937b8a986b5ed93f13a2d&user_id=0&is_app_user=1&language=0&parent_language=0&is_secure=1&stats_hash=2309491839c82a704b&group_id=0&ads_app_id=7869076_9ba760f1e19907aaa3&access_token_settings=&referrer=unknown&lc_name=d106d683&platform=web&is_widescreen=0&whitelist_scopes=friends,photos,video,stories,pages,status,notes,wall,docs,groups,stats,market,ads,notifications&group_whitelist_scopes=stories,photos,app_widget,messages,wall,docs,manage&auth_key=f3730660d049a2913b5e69c066b83b70&timestamp=1622546396&sign=lJeYjD3GiMBTqwqQGVyE-3zJyHWTS4XNvklIYGANps0&sign_keys=access_token,access_token_settings,ads_app_id,api_id,api_settings,api_url,auth_key,group_id,group_whitelist_scopes,is_app_user,is_secure,is_widescreen,language,lc_name,parent_language,platform,referrer,secret,sid,stats_hash,timestamp,user_id,viewer_id,viewer_type,whitelist_scopes&hash="

	v, _ := games.ParamsVerify(link, "TaWz5vVHUQnp7YgtuRlS")

	fmt.Println(v)

	// Output:
	// true
}
