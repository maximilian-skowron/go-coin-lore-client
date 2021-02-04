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
