package opengo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type EthereumTransactionResponse struct {
	TotalCount   int                    `json:"totalCount"`
	Transactions []*EthereumTransaction `json:"list"`
}

type Event struct {
	Type                     string            `json:"type"`
	Activated                bool              `json:"activated"`
	UserAddress              string            `json:"userAddress"`
	PartnerShare             int64             `json:"partnerShare"`
	Amount                   int64             `json:"amount"`
	ToAddress                string            `json:"toAddress"`
	CustomerAddress          string            `json:"customerAddress"`
	TransactionAmount        int64             `json:"transactionAmount"`
	ScaffoldTransactionIndex int64             `json:"scaffoldTransactionIndex"`
	Properties               map[string]string `json:"properties"`
}

type EthereumTransaction struct {
	Scaffold EthereumScaffold `json:"scaffold"`
	Event    Event            `json:"event"`
	Type     string           `json:"type"`
}

func (op *OpenGo) GetEthereumTransactions(ctx context.Context, address string) ([]*EthereumTransaction, error) {
	op.baseURL.Path = fmt.Sprintf("/api/ethereum-scaffolds/%s/%s", address, "transactions")
	response, err := op.SendRequest(ctx, "GET", nil)

	if err != nil {
		return nil, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	scaffoldTransactionResponse := &EthereumTransactionResponse{}
	json.Unmarshal(data, scaffoldTransactionResponse)
	return scaffoldTransactionResponse.Transactions, nil
}
