package coinlore

import (
	"context"
	"testing"
)

func TestGetCoin(t *testing.T) {
	c := NewClient(BaseURL)
	coinId := "90"

	coin, err := c.GetCoin(context.TODO(), coinId)
	if err != nil {
		t.Errorf("GetCoin errored with: %s", err.Error())
	}
	if coin.Id != coinId {
		t.Errorf("Expected coin id: %s; Got: %s", coinId, coin.Id)
	}
}

func TestGetCoins(t *testing.T) {
	c := NewClient(BaseURL)

	to := TickerOptions{
		Start: 0,
	}
	to.SetLimit(10)

	coins, err := c.GetCoins(context.TODO(), &to)
	if err != nil {
		t.Errorf("GetCoins errored with: %s", err.Error())
	}
	if len(coins.Data) != to.limit {
		t.Errorf("Expected coin count: %d; Got: %d", to.limit, len(coins.Data))
	}
}
