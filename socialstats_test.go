package coinlore

import (
	"context"
	"testing"
)

func TestGetSocialStats(t *testing.T) {
	c := NewClient(BaseURL)

	ss, err := c.GetSocialStat(context.TODO(), "80")
	if err != nil {
		t.Errorf("GetSocialStats errored with: %s", err.Error())
	}
	if ss.Reddit.Subscribers <= 0 {
		t.Errorf("Returned count should greater than zero; Got: %d", ss.Reddit.Subscribers)
	}
}
