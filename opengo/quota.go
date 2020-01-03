package opengo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// EthereumScaffoldQuota struct stores data related to EthereumScaffold EthereumScaffoldQuota
type EthereumScaffoldQuota struct {
	CurrentCount int `json:"currentCount"`
	LimitCount   int `json:"limitCount"`
}

// GetQuota retrieves scaffold quota from Open Platform
func (op *OpenGo) GetEthereumScaffoldQuota(ctx context.Context) (*EthereumScaffoldQuota, error) {
	op.baseURL.Path = "/api/ethereum-scaffolds/quota"
	response, err := op.SendRequest(ctx, "GET", nil)
	if err != nil {
		return nil, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	quota := &EthereumScaffoldQuota{}
	json.Unmarshal(data, quota)
	return quota, nil
}
