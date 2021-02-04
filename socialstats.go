package coinlore

import (
	"context"
	"fmt"
)

// SocialStat for a specific coin.
type SocialStat struct {
	Reddit struct {
		AvgActiveUsers float64 `json:"avg_active_users"`
		Subscribers    int32   `json:"subscribers"`
	} `json:"reddit"`
	Twitter struct {
		FollowersCount int32 `json:"followers_count"`
		StatusCount    int32 `json:"status_count"`
	} `json:"twitter"`
}

// GetSocialStat will return the social stats for reddit and twitter for the given coin id.
func (c Client) GetSocialStat(ctx context.Context, id string) (*SocialStat, error) {
	url := fmt.Sprintf("%s/coin/social_stats/?id=%s", c.BaseURL, id)

	res, err := c.getRequest(url, ctx)
	if err != nil {
		return nil, err
	}

	var ss SocialStat
	if exErr := c.extractAndParseBody(*res, &ss); exErr != nil {
		return nil, exErr
	}

	return &ss, nil
}
