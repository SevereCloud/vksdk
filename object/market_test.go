package object_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/object"
)

func TestMarketMarketItem_ToAttachment(t *testing.T) {
	t.Parallel()

	f := func(market object.MarketMarketItem, want string) {
		if got := market.ToAttachment(); got != want {
			t.Errorf("MarketMarketItem.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.MarketMarketItem{ID: 10, OwnerID: 20}, "market20_10")
	f(object.MarketMarketItem{ID: 20, OwnerID: -10}, "market-10_20")
}

func TestMarketMarketAlbum_ToAttachment(t *testing.T) {
	t.Parallel()

	f := func(marketAlbum object.MarketMarketAlbum, want string) {
		if got := marketAlbum.ToAttachment(); got != want {
			t.Errorf("MarketMarketAlbum.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.MarketMarketAlbum{ID: 10, OwnerID: 20}, "market_album20_10")
	f(object.MarketMarketAlbum{ID: 20, OwnerID: -10}, "market_album-10_20")
}
