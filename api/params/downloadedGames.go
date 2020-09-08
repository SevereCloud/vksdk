package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// DownloadedGamesGetPaidStatusBuilder builder.
//
// https://vk.com/dev/downloadedGames.getPaidStatus
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
