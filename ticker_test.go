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
