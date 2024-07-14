package marusia // import "github.com/SevereCloud/vksdk/v3/marusia"

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
	"strings"
)

// CreateDeepLink returns a link to go to the skill, as well as transfer data to your skill.
//
// See https://dev.vk.com/ru/marusia/deep-link
func CreateDeepLink(marusiaID string, params map[string]string) string {
	query := url.Values{}
	query.Add("event_name", "external."+marusiaID+".start")

	if params != nil {
		b, _ := json.Marshal(params)

		base64Params := base64.StdEncoding.EncodeToString(b)
		base64Params = strings.ReplaceAll(base64Params, "+", "-")
		base64Params = strings.ReplaceAll(base64Params, "/", "_")
		base64Params = strings.TrimRight(base64Params, "=")
		query.Add("event_data", base64Params)
	}

	u := &url.URL{
		Scheme:   "https",
		Host:     "vk.me",
		Path:     "marusia",
		RawQuery: query.Encode(),
	}

	return u.String()
}
