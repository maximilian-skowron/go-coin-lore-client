package coinlore

import (
	"context"
	"log"
	"testing"
)

func TestGetAllExchanges(t *testing.T) {
	c := NewClient(BaseURL)

	exchanges, err := c.GetAllExchanges(context.TODO())
	log.Println(exchanges)
	if err != nil {
		t.Errorf("GetCryptoMarket errored with: %s", err.Error())
	}
	if len(exchanges) == 0 {
		t.Errorf("Returned length should greater than zero; Got: %d", len(exchanges))
	}
}
