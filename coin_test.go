package coinlore

import (
	"context"
	"testing"
)

func TestClient_GetMarketsForCoin(t *testing.T) {
	c := NewClient(BaseURL)

	const coinId = "90"

	markets, err := c.GetMarketsForCoin(context.TODO(), coinId)
	if err != nil {
		t.Errorf("GetCryptoMarket errored with: %s", err.Error())
	}
	if len(markets) == 0 {
		t.Errorf("Expected active market >0; Got: %d", len(markets))
	}
}

