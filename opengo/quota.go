package opengo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Quota struct stores data related to Scaffold Quota
type Quota struct {
	CurrentCount int `json:"currentCount"`
	LimitCount   int `json:"limitCount"`
}

// GetQuota retrieves scaffold quota from Open Platform
func (op *OpenGo) GetQuota(ctx context.Context) (*Quota, error) {
	op.baseURL.Path = "/api/scaffolds/quota"
	response, err := op.SendRequest(ctx, "GET", nil)
	if err != nil {
		return nil, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	quota := &Quota{}
	json.Unmarshal(data, quota)
	return quota, nil
}
