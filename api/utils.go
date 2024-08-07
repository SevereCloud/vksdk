package api // import "github.com/SevereCloud/vksdk/v3/api"

import (
	"fmt"

	"github.com/SevereCloud/vksdk/v3/object"
	"github.com/vmihailenco/msgpack/v5"
)

// UtilsCheckLinkResponse struct.
type UtilsCheckLinkResponse object.UtilsLinkChecked

// UtilsCheckLink checks whether a link is blocked in VK.
//
// https://dev.vk.com/method/utils.checkLink
func (vk *VK) UtilsCheckLink(params Params) (response UtilsCheckLinkResponse, err error) {
	err = vk.RequestUnmarshal("utils.checkLink", &response, params)
	return
}

// UtilsDeleteFromLastShortened deletes shortened link from user's list.
//
// https://dev.vk.com/method/utils.deleteFromLastShortened
func (vk *VK) UtilsDeleteFromLastShortened(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("utils.deleteFromLastShortened", &response, params)
	return
}

// UtilsGetLastShortenedLinksResponse struct.
type UtilsGetLastShortenedLinksResponse struct {
	Count int                             `json:"count"`
	Items []object.UtilsLastShortenedLink `json:"items"`
}

// UtilsGetLastShortenedLinks returns a list of user's shortened links.
//
// https://dev.vk.com/method/utils.getLastShortenedLinks
func (vk *VK) UtilsGetLastShortenedLinks(params Params) (response UtilsGetLastShortenedLinksResponse, err error) {
	err = vk.RequestUnmarshal("utils.getLastShortenedLinks", &response, params)
	return
}

// UtilsGetLinkStatsResponse struct.
type UtilsGetLinkStatsResponse object.UtilsLinkStats

// UtilsGetLinkStats returns stats data for shortened link.
//
//	extended=0
//
// https://dev.vk.com/method/utils.getLinkStats
func (vk *VK) UtilsGetLinkStats(params Params) (response UtilsGetLinkStatsResponse, err error) {
	err = vk.RequestUnmarshal("utils.getLinkStats", &response, params, Params{"extended": false})

	return
}

// UtilsGetLinkStatsExtendedResponse struct.
type UtilsGetLinkStatsExtendedResponse object.UtilsLinkStatsExtended

// UtilsGetLinkStatsExtended returns stats data for shortened link.
//
//	extended=1
//
// https://dev.vk.com/method/utils.getLinkStats
func (vk *VK) UtilsGetLinkStatsExtended(params Params) (response UtilsGetLinkStatsExtendedResponse, err error) {
	err = vk.RequestUnmarshal("utils.getLinkStats", &response, params, Params{"extended": true})

	return
}

// UtilsGetServerTime returns the current time of the VK server.
//
// https://dev.vk.com/method/utils.getServerTime
func (vk *VK) UtilsGetServerTime(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("utils.getServerTime", &response, params)
	return
}

// UtilsGetShortLinkResponse struct.
type UtilsGetShortLinkResponse object.UtilsShortLink

// UtilsGetShortLink allows to receive a link shortened via vk.cc.
//
// https://dev.vk.com/method/utils.getShortLink
func (vk *VK) UtilsGetShortLink(params Params) (response UtilsGetShortLinkResponse, err error) {
	err = vk.RequestUnmarshal("utils.getShortLink", &response, params)
	return
}

// UtilsResolveScreenNameResponse struct.
type UtilsResolveScreenNameResponse object.UtilsDomainResolved

// UnmarshalJSON UtilsResolveScreenNameResponse.
//
// BUG(VK): UtilsResolveScreenNameResponse return [].
func (resp *UtilsResolveScreenNameResponse) UnmarshalJSON(data []byte) error {
	var p object.UtilsDomainResolved

	err := p.UnmarshalJSON(data)
	if err != nil {
		return fmt.Errorf("api.UtilsResolveScreenNameResponse: %w", err)
	}

	*resp = UtilsResolveScreenNameResponse(p)

	return nil
}

// DecodeMsgpack UtilsResolveScreenNameResponse.
//
// BUG(VK): UtilsResolveScreenNameResponse return [].
func (resp *UtilsResolveScreenNameResponse) DecodeMsgpack(dec *msgpack.Decoder) error {
	var p object.UtilsDomainResolved

	err := p.DecodeMsgpack(dec)
	if err != nil {
		return fmt.Errorf("api.UtilsResolveScreenNameResponse: %w", err)
	}

	*resp = UtilsResolveScreenNameResponse(p)

	return nil
}

// UtilsResolveScreenName detects a type of object (e.g., user, community, application) and its ID by screen name.
//
// https://dev.vk.com/method/utils.resolveScreenName
func (vk *VK) UtilsResolveScreenName(params Params) (response UtilsResolveScreenNameResponse, err error) {
	err = vk.RequestUnmarshal("utils.resolveScreenName", &response, params)
	return
}
