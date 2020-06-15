package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// DownloadedGamesGetPaidStatusResponse struct.
type DownloadedGamesGetPaidStatusResponse struct {
	IsPaid object.BaseBoolInt `json:"is_paid"`
}

// DownloadedGamesGetPaidStatus method.
//
// https://vk.com/dev/downloadedGames.getPaidStatus
func (vk *VK) DownloadedGamesGetPaidStatus(params Params) (response DownloadedGamesGetPaidStatusResponse, err error) {
	params["extended"] = false
	err = vk.RequestUnmarshal("downloadedGames.getPaidStatus", params, &response)

	return
}
