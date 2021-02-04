package coinlore

import (
	"context"
	"fmt"
)

// Market represents a market for crypto currency.
type Market struct {
	Name      string  `json:"name"`
	Base      string  `json:"base"`
	Quote     string  `json:"quote"`
	Price     float64 `json:"price"`
	PriceUsd  float64 `json:"price_usd"`
	Volume    float64 `json:"volume"`
	VolumeUsd float64 `json:"volume_usd"`
	Time      int64   `json:"time"`
}

// GetMarketsForCoin will return a slice of Market that handle the provided coin.
func (c Client) GetMarketsForCoin(ctx context.Context, id string) ([]Market, error) {
	url := fmt.Sprintf("%s/coin/markets/?id=%s", c.BaseURL, id)

	res, err := c.getRequest(url, ctx)
	if err != nil {
		return nil, err
	}

	var m []Market
	if exErr := c.extractAndParseBody(*res, &m); exErr != nil {
		return nil, exErr
	}

	return m, nil
}
