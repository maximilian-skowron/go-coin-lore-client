package coinlore

import (
	"context"
	"fmt"
	"log"
)

const (
	defaultLimit = 100
	defaultStart = 0
)

type TickerOptions struct {
	Start int64
	limit int
}

func (to *TickerOptions) SetLimit(limit int) {
	if limit > defaultLimit {
		log.Println("limit max is 100 and will set automatically to 100")
		limit = defaultLimit
	}

	to.limit = limit
}

func (to *TickerOptions) SetStart(s int64) {
	to.Start = s
}

type Coin struct {
	Id               string `json:"id"`
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	NameId           string `json:"nameid"`
	Rank             int    `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PercentChange24h string `json:"percent_change_24h"`
	PercentChange1h  string `json:"percent_change_1h"`
	PercentChange7d  string `json:"percent_change_7d"`
	PriceBtc         string `json:"price_btc"`
	MarketCapUsd     string `json:"market_cap_usd`
	Csupply          string `json:"csupply"`
	Tsupply          string `json:"tsupply"`
	Msupply          string `json:"msupply"`
}

// the volume fields are floats if you get all and strings of you only want one coin
type CoinsRes struct {
	Coin
	Volume24  float64 `json:"volume24"`
	Volume24a float64 `json:"volume24a"`
}

type CoinRes struct {
	Coin
	Volume24  string `json:"volume24"`
	Volume24a string `json:"volume24a"`
}

type Coins struct {
	Data []Coins `json:"data"`
	Info struct {
		CoinsNum uint  `json:"coins_num"`
		Time     int64 `json:"time"`
	} `json:"info"`
}

func (c Client) GetCoins(ctx context.Context, to *TickerOptions) (*Coins, error) {
	var (
		limit int
		start int64
	)

	if to == nil {
		limit = defaultLimit
		start = defaultStart
	} else {
		limit = to.limit
		start = to.Start
	}

	url := fmt.Sprintf("%s/tickers/?start=%d&limit=%d", c.BaseURL, start, limit)

	res, err := c.getRequest(url, ctx)
	if err != nil {
		return nil, err
	}

	var coins Coins
	if exErr := c.extractAndParseBody(*res, &coins); exErr != nil {
		return nil, exErr
	}

	return &coins, nil
}

func (c Client) GetCoin(ctx context.Context, id string) (*CoinRes, error) {
	url := fmt.Sprintf("%s/ticker/?id=%s", c.BaseURL, id)

	res, err := c.getRequest(url, ctx)
	if err != nil {
		return nil, err
	}

	var coin []CoinRes
	if exErr := c.extractAndParseBody(*res, &coin); exErr != nil {
		return nil, exErr
	}

	return &coin[0], nil
}
