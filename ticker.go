package coinlore

import (
	"context"
	"fmt"
	"log"
	"time"
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
	Id               string  `json:"id"`
	Symbol           string  `json:"id"`
	Name             string  `json:"name"`
	NameId           string  `json:"nameid"`
	Rank             int     `json:"rank"`
	PriceUsd         string  `json:"price_usd"`
	PercentChange24h string  `json:"percent_change_24h"`
	PercentChange1h  string  `json:"percent_change_1h"`
	PercentChange7d  string  `json:"percent_change_7d"`
	PriceBtc         string  `json:"price_btc"`
	MarketCapUsd     string  `json:"market_cap_usd`
	Volume24         float64 `json:"volume24"`
	Volume24a        float64 `json:"volume24a"`
	Csupply          string  `json:"csupply"`
	Tsupply          string  `json:"tsupply"`
	Msupply          string  `json:"msupply"`
}

type Coins struct {
	Data []Coin `json:"data"`
	Info struct {
		CoinsNum uint      `json:"coins_num"`
		Time     time.Time `json:"time"`
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

	url := fmt.Sprintf("%s/tickers/?start=%s&limit=%s", c.BaseURL, start, limit)

	res, err := c.getRequest(url, ctx)
	if err != nil {
		return nil, err
	}

	var coins *Coins
	if exErr := c.extractAndParseBody(*res, coins); exErr != nil {
		return nil, exErr
	}

	return coins, nil
}

func (c Client) GetCoin(ctx context.Context, id string) (*Coin, error) {
	url := fmt.Sprintf("%s/tickers/?id=%s", c.BaseURL, id)

	res, err := c.getRequest(url, ctx)
	if err != nil {
		return nil, err
	}

	var coin *Coin
	if exErr := c.extractAndParseBody(*res, coin); exErr != nil {
		return nil, exErr
	}

	return coin, nil
}
