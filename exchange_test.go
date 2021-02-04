package coinlore

import (
	"context"
	"testing"
)

func TestGetAllExchanges(t *testing.T) {
	c := NewClient(BaseURL)

	exchanges, err := c.GetAllExchanges(context.TODO())
	if err != nil {
		t.Errorf("GetAllExchanges errored with: %s", err.Error())
	}
	if len(exchanges) == 0 {
		t.Errorf("Returned length should greater than zero; Got: %d", len(exchanges))
	}
}

func TestGetExchange(t *testing.T) {
	c := NewClient(BaseURL)

	name := "Binance"

	se, err := c.GetExchange(context.TODO(), "5")
	if err != nil {
		t.Errorf("GetExchange errored with: %s", err.Error())
	}
	if se.Name != name {
		t.Errorf("Expected name: %s; Got: %s", name, se.Name)
	}
}
