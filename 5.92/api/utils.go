package api // import "github.com/severecloud/vksdk/5.92/api"

import (
	"encoding/json"

	"github.com/severecloud/vksdk/5.92/object"
)

// UtilsCheckLinkResponse struct
type UtilsCheckLinkResponse object.UtilsLinkChecked

// UtilsCheckLink checks whether a link is blocked in VK.
// https://vk.com/dev/utils.checkLink
// FIXME: Если короткое имя screen_name не занято, то будет возвращён пустой объект.
func (vk VK) UtilsCheckLink(params map[string]string) (response UtilsCheckLinkResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("utils.checkLink", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// UtilsDeleteFromLastShortened deletes shortened link from user's list.
// https://vk.com/dev/utils.deleteFromLastShortened
func (vk VK) UtilsDeleteFromLastShortened(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("utils.deleteFromLastShortened", params)

	return
}

// UtilsGetLastShortenedLinksResponse struct
type UtilsGetLastShortenedLinksResponse struct {
	Count int                             `json:"count"`
	Items []object.UtilsLastShortenedLink `json:"items"`
}

// UtilsGetLastShortenedLinks returns a list of user's shortened links.
// https://vk.com/dev/utils.getLastShortenedLinks
func (vk VK) UtilsGetLastShortenedLinks(params map[string]string) (response UtilsGetLastShortenedLinksResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("utils.getLastShortenedLinks", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// UtilsGetLinkStatsResponse struct
type UtilsGetLinkStatsResponse object.UtilsLinkStats

// UtilsGetLinkStats returns stats data for shortened link.
// https://vk.com/dev/utils.getLinkStats
func (vk VK) UtilsGetLinkStats(params map[string]string) (response UtilsGetLinkStatsResponse, vkErr Error) {
	params["extended"] = "0"
	rawResponse, vkErr := vk.Request("utils.getLinkStats", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// UtilsGetLinkStatsExtendedResponse struct
type UtilsGetLinkStatsExtendedResponse object.UtilsLinkStatsExtended

// UtilsGetLinkStatsExtended returns stats data for shortened link.
// https://vk.com/dev/utils.getLinkStats
func (vk VK) UtilsGetLinkStatsExtended(params map[string]string) (response UtilsGetLinkStatsExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	rawResponse, vkErr := vk.Request("utils.getLinkStats", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// UtilsGetServerTimeResponse struct
type UtilsGetServerTimeResponse int

// UtilsGetServerTime returns the current time of the VK server.
// https://vk.com/dev/utils.getServerTime
func (vk VK) UtilsGetServerTime() (response UtilsGetServerTimeResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("utils.getServerTime", make(map[string]string))
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// UtilsGetShortLinkResponse struct
type UtilsGetShortLinkResponse object.UtilsShortLink

// UtilsGetShortLink allows to receive a link shortened via vk.cc.
// https://vk.com/dev/utils.getShortLink
func (vk VK) UtilsGetShortLink(params map[string]string) (response UtilsGetShortLinkResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("utils.getShortLink", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// UtilsResolveScreenNameResponse struct
type UtilsResolveScreenNameResponse object.UtilsDomainResolved

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
