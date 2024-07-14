package vkapps_test

import (
	"fmt"
	"net/url"

	"github.com/SevereCloud/vksdk/v3/vkapps"
)

func ExampleParamsVerification() {
	link := "https://example.com/?vk_user_id=494075&vk_app_id=6736218&vk_is_app_user=1&vk_are_notifications_enabled=1&vk_language=ru&vk_access_token_settings=&vk_platform=android&sign=htQFduJpLxz7ribXRZpDFUH-XEUhC9rBPTJkjUFEkRA"

	pv := vkapps.NewParamsVerification("wvl68m4dR1UpLrVRli") // Client secret

	u, _ := url.Parse(link)

	v, _ := pv.Verify(u)

	fmt.Println(v)

	// Output:
	// true
}

func ExampleParamsVerify() {
	link := "https://example.com/?vk_user_id=494075&vk_app_id=6736218&vk_is_app_user=1&vk_are_notifications_enabled=1&vk_language=ru&vk_access_token_settings=&vk_platform=android&sign=htQFduJpLxz7ribXRZpDFUH-XEUhC9rBPTJkjUFEkRA"

	v, _ := vkapps.ParamsVerify(link, "wvl68m4dR1UpLrVRli")

	fmt.Println(v)

	// Output:
	// true
}
