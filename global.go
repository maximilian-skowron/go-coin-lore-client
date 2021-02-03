package coinlore

import (
	"context"
	"fmt"
)

type CryptoMarketInfo struct {
	CoinsCount       int     `json:"coins_count"`
	ActiveMarkets    int     `json:"active_markets"`
	TotalMcap        float64 `json:"total_mcap"`
	TotalVolume      float64 `json:"total_volume"`
	BtcD             string  `json:"btc_d"`
	EthD             string  `json:"eth_d"`
	McapChange       string  `json:"mcap_change"`
	VolumeChange     string  `json:"volume_change"`
	AvgChangePercent string  `json:"avg_change_percent"`
	VolumeAth        int64   `json:"volume_ath"`
	McapAth          float64 `json:"mcap_ath"`
}

func (c Client) GetCryptoMarket(ctx context.Context) (*CryptoMarketInfo, error) {
	url := fmt.Sprintf("%s/global/", c.BaseURL)

	res, err := c.getRequest(url, ctx)
	if err != nil {
		return nil, err
	}

	var cm []CryptoMarketInfo
	if exErr := c.extractAndParseBody(*res, &cm); exErr != nil {
		return nil, exErr
	}

	return &cm[0], nil
}
