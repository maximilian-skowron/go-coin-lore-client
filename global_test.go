package coinlore

import (
	"context"
	"testing"
)

func TestGetMarketInfo(t *testing.T) {
	c := NewClient(BaseURL)

	market, err := c.GetCryptoMarket(context.TODO())
	if err != nil {
		t.Errorf("GetCryptoMarket errored with: %s", err.Error())
	}
	if market.ActiveMarkets == 0 {
		t.Errorf("Expected active market >0; Got: %d", market.ActiveMarkets)
	}
}
