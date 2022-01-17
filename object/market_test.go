package object_test

import (
	"encoding/json"
	"testing"

	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/vmihailenco/msgpack/v5/msgpcode"
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

func TestMarketMarketItem_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(data []byte, wantMarket object.MarketMarketItem) {
		var market object.MarketMarketItem

		err := json.Unmarshal(data, &market)
		assert.NoError(t, err)
		assert.Equal(t, wantMarket, market)
	}

	f([]byte("false"), object.MarketMarketItem{})
	f([]byte(`{"id":1}`), object.MarketMarketItem{ID: 1})

	var market object.MarketMarketItem
	err := json.Unmarshal([]byte("0"), &market)
	assert.Error(t, err)
}

func TestMarketMarketItem_DecodeMsgpack(t *testing.T) {
	t.Parallel()

	f := func(data []byte, wantMarket object.MarketMarketItem, wantErr string) {
		var market object.MarketMarketItem

		err := msgpack.Unmarshal(data, &market)
		if err != nil || wantErr != "" {
			assert.EqualError(t, err, wantErr)
		}

		assert.Equal(t, wantMarket, market)
	}

	f([]byte{msgpcode.False}, object.MarketMarketItem{}, "")
	f([]byte{0x81, 0xA2, 0x69, 0x64, 0x01}, object.MarketMarketItem{ID: 1}, "")

	f([]byte("\xc3"), object.MarketMarketItem{}, "msgpack: unexpected code=c3 decoding map length")
	f(nil, object.MarketMarketItem{}, "EOF")
}

func TestMarketPrice_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(data []byte, wantMarket object.MarketPrice) {
		var market object.MarketPrice

		err := json.Unmarshal(data, &market)
		assert.NoError(t, err)
		assert.Equal(t, wantMarket, market)
	}

	f([]byte("[]"), object.MarketPrice{})
	f([]byte(`{"text":"a"}`), object.MarketPrice{Text: "a"})

	var market object.MarketPrice
	err := json.Unmarshal([]byte("0"), &market)
	assert.Error(t, err)
}

func TestMarketPrice_DecodeMsgpack(t *testing.T) {
	t.Parallel()

	f := func(data []byte, wantMarket object.MarketPrice, wantErr string) {
		var market object.MarketPrice

		err := msgpack.Unmarshal(data, &market)
		if err != nil || wantErr != "" {
			assert.EqualError(t, err, wantErr)
		}

		assert.Equal(t, wantMarket, market)
	}

	f([]byte{msgpcode.FixedArrayLow}, object.MarketPrice{}, "")
	f([]byte{
		0x81, 0xA4, 0x74, 0x65, 0x78, 0x74, 0xA1, 0x61,
	}, object.MarketPrice{Text: "a"}, "")

	f([]byte("\xc2"), object.MarketPrice{}, "msgpack: unexpected code=c2 decoding map length")
	f(nil, object.MarketPrice{}, "EOF")
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
