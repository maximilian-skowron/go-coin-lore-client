// package to hold client informations
package coinlore

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const BaseURL = "https://api.coinlore.net/api"

type Client struct {
	BaseURL string
}

func NewClient(url string) *Client {
	return &Client{
		BaseURL: url,
	}
}

func (c Client) getRequest(url string, ctx context.Context) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, reqErr := http.DefaultClient.Do(req)
	if reqErr != nil {
		return nil, reqErr
	}

	return res, nil
}

func (c Client) extractAndParseBody(res http.Response, parse interface{}) error {
	jsonBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if unerr := json.Unmarshal(jsonBody, parse); unerr != nil {
		return unerr
	}

	return nil
}
