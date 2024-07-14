package api // import "github.com/SevereCloud/vksdk/v3/api"

import (
	"github.com/SevereCloud/vksdk/v3/object"
)

// DownloadedGamesGetPaidStatusResponse struct.
type DownloadedGamesGetPaidStatusResponse struct {
	IsPaid object.BaseBoolInt `json:"is_paid"`
}

// DownloadedGamesGetPaidStatus method.
//
// https://dev.vk.com/method/downloadedGames.getPaidStatus
func (vk *VK) DownloadedGamesGetPaidStatus(params Params) (response DownloadedGamesGetPaidStatusResponse, err error) {
	err = vk.RequestUnmarshal("downloadedGames.getPaidStatus", &response, params, Params{"extended": false})

	return
}
