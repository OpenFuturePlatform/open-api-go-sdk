package opengo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TransactionResponse struct {
	TotalCount   int            `json:"totalCount"`
	Transactions []*Transaction `json:"list"`
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

type Transaction struct {
	Scaffold Scaffold `json:"scaffold"`
	Event    Event    `json:"event"`
	Type     string   `json:"type"`
}

func (op *OpenGo) GetTransactions(ctx context.Context, address string) ([]*Transaction, error) {
	op.baseURL.Path = fmt.Sprintf("/api/scaffolds/%s/%s", address, "transactions")
	request, _ := http.NewRequest("GET", op.baseURL.String(), nil)
	request.Header.Set("Authorization", op.token)
	request = request.WithContext(ctx)

	response, err := op.httpClient.Do(request)

	if err != nil {
		return nil, fmt.Errorf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		scaffoldTransactionResponse := &TransactionResponse{}
		json.Unmarshal(data, scaffoldTransactionResponse)
		return scaffoldTransactionResponse.Transactions, nil
	}
	return nil, err
}
