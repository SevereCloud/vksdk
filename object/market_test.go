package object

import (
	"testing"
)

func TestMarketMarketItem_ToAttachment(t *testing.T) {
	type fields struct {
		ID      int
		OwnerID int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "market20_10",
			fields: fields{10, 20},
			want:   "market20_10",
		},
		{
			name:   "market-10_20",
			fields: fields{20, -10},
			want:   "market-10_20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			market := MarketMarketItem{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := market.ToAttachment(); got != tt.want {
				t.Errorf("MarketMarketItem.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketMarketAlbum_ToAttachment(t *testing.T) {
	type fields struct {
		ID      int
		OwnerID int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "market_album20_10",
			fields: fields{10, 20},
			want:   "market_album20_10",
		},
		{
			name:   "market_album-10_20",
			fields: fields{20, -10},
			want:   "market_album-10_20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			marketAlbum := MarketMarketAlbum{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := marketAlbum.ToAttachment(); got != tt.want {
				t.Errorf("MarketMarketAlbum.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}
