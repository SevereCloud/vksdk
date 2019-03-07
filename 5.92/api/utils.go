package api // import "github.com/severecloud/vksdk/5.92/api"

import "encoding/json"

// UtilsCheckLinkResponse struct
type UtilsCheckLinkResponse struct{}

// TODO: utils.checkLink checks whether a link is blocked in VK.
// https://vk.com/dev/utils.checkLink

// UtilsDeleteFromLastShortenedResponse struct
type UtilsDeleteFromLastShortenedResponse struct{}

// TODO: utils.deleteFromLastShortened deletes shortened link from user's list.
// https://vk.com/dev/utils.deleteFromLastShortened

// UtilsGetLastShortenedLinksResponse struct
type UtilsGetLastShortenedLinksResponse struct{}

// TODO: utils.getLastShortenedLinks returns a list of user's shortened links.
// https://vk.com/dev/utils.getLastShortenedLinks

// UtilsGetLinkStatsResponse struct
type UtilsGetLinkStatsResponse struct{}

// TODO: utils.getLinkStats returns stats data for shortened link.
// https://vk.com/dev/utils.getLinkStats

// UtilsGetServerTimeResponse struct
type UtilsGetServerTimeResponse struct{}

// TODO: utils.getServerTime returns the current time of the VK server.
// https://vk.com/dev/utils.getServerTime

// UtilsGetShortLinkResponse struct
type UtilsGetShortLinkResponse struct{}

// TODO: utils.getShortLink allows to receive a link shortened via vk.cc.
// https://vk.com/dev/utils.getShortLink

// UtilsResolveScreenNameResponse struct
type UtilsResolveScreenNameResponse struct {
	Type     string `json:"type"`
	ObjectID int    `json:"object_id"`
}

// UtilsResolveScreenName detects a type of object (e.g., user, community, application) and its ID by screen name.
// https://vk.com/dev/utils.resolveScreenName
func (vk VK) UtilsResolveScreenName(params map[string]string) (response UtilsResolveScreenNameResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("utils.resolveScreenName", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}
