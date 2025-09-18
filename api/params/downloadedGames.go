package params // import "github.com/SevereCloud/vksdk/v3/api/params"

import (
	"github.com/SevereCloud/vksdk/v3/api"
)

// DownloadedGamesGetPaidStatusBuilder builder.
//
// https://dev.vk.ru/method/downloadedGames.getPaidStatus
type DownloadedGamesGetPaidStatusBuilder struct {
	api.Params
}

// NewDownloadedGamesGetPaidStatusBuilder func.
func NewDownloadedGamesGetPaidStatusBuilder() *DownloadedGamesGetPaidStatusBuilder {
	return &DownloadedGamesGetPaidStatusBuilder{api.Params{}}
}

// UserID parameter.
func (b *DownloadedGamesGetPaidStatusBuilder) UserID(v int) *DownloadedGamesGetPaidStatusBuilder {
	b.Params["user_id"] = v
	return b
}
