package coinlore

import (
	"context"
	"encoding/json"
	"fmt"
)

// Exchange isn't the exact object of the get all exchanges endpoint.
// It is only one exchange out of this stupid large object.
type Exchange struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	NameID      string  `json:"name_id"`
	VolumeUsd   float64 `json:"volume_usd"`
	ActivePairs float64 `json:"active_pairs"`
	URL         string  `json:"url"`
	Country     string  `json:"country"`
	// there are more stupid fields that arn't dokumented on there site
}

// SpecExchange is a specific exchange with trimmed fields and a slice of pairs.
type SpecExchange struct {
	Name     string `json:"name"`
	DateLive string `json:"date_live"`
	URL      string `json:"url"`
	Pairs    []Pair `json:"pairs"`
}

// Pair represents a pair of a specific exchange.
type Pair struct {
	Base     string  `json:"base"`
	Quote    string  `json:"quote"`
	Volume   float64 `json:"volume"`
	Price    float64 `json:"price"`
	PriceUsd float64 `json:"price_usd"`
	Time     int64   `json:"time"`
}

// GetAllExchanges will return a slice of all exchanges.
// baseurl/exchanges
func (c Client) GetAllExchanges(ctx context.Context) ([]Exchange, error) {
	url := fmt.Sprintf("%s/exchanges", c.BaseURL)

	res, err := c.getRequest(url, ctx)
	if err != nil {
		return nil, err
	}

	var objMap map[string]json.RawMessage
	if exErr := c.extractAndParseBody(*res, &objMap); exErr != nil {
		return nil, exErr
	}

	exchanges := make([]Exchange, 0)
	for _, v := range objMap {
		var e Exchange
		if uErr := json.Unmarshal(v, &e); uErr != nil {
			return nil, uErr
		}
		exchanges = append(exchanges, e)
	}

	return exchanges, nil
}

// GetExchange will return a SpecExchange for the given id.
func (c Client) GetExchange(ctx context.Context, id string) (*SpecExchange, error) {
	url := fmt.Sprintf("%s/exchange/?id=%s", c.BaseURL, id)

	res, err := c.getRequest(url, ctx)
	if err != nil {
		return nil, err
	}

	var objMap map[string]json.RawMessage
	if exErr := c.extractAndParseBody(*res, &objMap); exErr != nil {
		return nil, exErr
	}

	var se SpecExchange
	if uErr := json.Unmarshal(objMap["0"], &se); uErr != nil {
		return nil, uErr
	}

	p := make([]Pair, 0)
	if uErr := json.Unmarshal(objMap["pairs"], &p); uErr != nil {
		return nil, uErr
	}

	se.Pairs = p

	return &se, nil
}
