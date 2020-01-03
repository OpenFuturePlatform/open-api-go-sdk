package opengo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type scaffoldResponse struct {
	TotalCount int                 `json:"totalCount"`
	Scaffolds  []*EthereumScaffold `json:"list"`
}

type EthereumScaffold struct {
	OpenKey          string                      `json:"openKey,omitempty"`
	Address          string                      `json:"address,omitempty"`
	ABI              string                      `json:"abi,omitempty"`
	Description      string                      `json:"description,omitempty"`
	FiatAmount       string                      `json:"fiatAmount,omitempty"`
	Currency         string                      `json:"currency,omitempty"`
	ConversionAmount string                      `json:"conversionAmount,omitempty"`
	DeveloperAddress string                      `json:"developerAddress,omitempty"`
	WebHook          string                      `json:"webHook,omitempty"`
	Properties       []*EthereumScaffoldProperty `json:"properties,omitempty"`
	Enabled          bool                        `json:"enabled,omitempty"`
}

// AddEthereumScaffoldProperty adds EthereumScaffoldProperty to EthereumScaffold object.
func (scaffold *EthereumScaffold) AddEthereumScaffoldProperty(Name, Type, DefaultValue string) error {
	scaffoldProperty := NewEthereumScaffoldProperty(Name, Type, DefaultValue)
	scaffold.Properties = append(scaffold.Properties, scaffoldProperty)
	return nil
}

type EthereumScaffoldSummary struct {
	Scaffold         *EthereumScaffold `json:"scaffold"`
	TransactionIndex int64             `json:"transactionIndex"`
	TokenBalance     int64             `json:"tokenBalance"`
	Enabled          bool              `json:"enabled"`
	ShareHolders     []*ShareHolder    `json:"shareHolders"`
}

type EthereumScaffoldProperty struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	DefaultValue string `json:"defaultValue"`
}

// NewEthereumScaffoldProperty return a new object of EthereumScaffoldProperty
func NewEthereumScaffoldProperty(Name, Type, DefaultValue string) *EthereumScaffoldProperty {
	scaffoldProperty := &EthereumScaffoldProperty{
		Name:         Name,
		Type:         Type,
		DefaultValue: DefaultValue,
	}
	return scaffoldProperty
}

// NewEthereumScaffold return a new object of NewEthereumScaffold
func NewEthereumScaffold(OpenKey, Address, ABI, Description, FiatAmount, Currency, ConversionAmount, DeveloperAddress, Webhook string, Enabled bool) *EthereumScaffold {
	scaffold := &EthereumScaffold{
		OpenKey:          OpenKey,
		Address:          Address,
		ABI:              ABI,
		Description:      Description,
		FiatAmount:       FiatAmount,
		Currency:         Currency,
		ConversionAmount: ConversionAmount,
		DeveloperAddress: DeveloperAddress,
		WebHook:          Webhook,
		Enabled:          Enabled,
	}
	return scaffold
}

// GetEthereumScaffolds function retrieves all []*EthereumScaffolds from Open Platform and returns a list of EthereumScaffold
func (op *OpenGo) GetEthereumScaffolds(ctx context.Context) ([]*EthereumScaffold, error) {
	op.baseURL.Path = "/api/ethereum-scaffolds"
	response, err := op.SendRequest(ctx, "GET", nil)

	if err != nil {
		return nil, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	sr := &scaffoldResponse{}
	json.Unmarshal(data, sr)
	return sr.Scaffolds, nil
}

// GetEthereumScaffold returns *EthereumScaffold given it's address
func (op *OpenGo) GetEthereumScaffold(ctx context.Context, address string) (*EthereumScaffold, error) {
	op.baseURL.Path = fmt.Sprintf("/api/ethereum-scaffolds/%s", address)
	response, err := op.SendRequest(ctx, "GET", nil)

	if err != nil {
		return &EthereumScaffold{}, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	scaffold := &EthereumScaffold{}
	json.Unmarshal(data, scaffold)
	return scaffold, nil
}

// GetEthereumScaffoldSummary retrieves scaffold summary from Open Platform API and returns *EthereumScaffoldSummary
func (op *OpenGo) GetEthereumScaffoldSummary(ctx context.Context, address string) (*EthereumScaffoldSummary, error) {
	op.baseURL.Path = fmt.Sprintf("/api/ethereum-scaffolds/%s/%s", address, "summary")
	response, err := op.SendRequest(ctx, "GET", nil)

	if err != nil {
		return &EthereumScaffoldSummary{}, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	scaffoldSummary := &EthereumScaffoldSummary{}
	json.Unmarshal(data, scaffoldSummary)
	return scaffoldSummary, nil
}

// DeployEthereumScaffold takes a EthereumScaffold and sends it to Open Platform. If successful, returns given EthereumScaffold.
func (op *OpenGo) DeployEthereumScaffold(ctx context.Context, scaffold EthereumScaffold) (EthereumScaffold, error) {
	op.baseURL.Path = "/api/ethereum-scaffolds/doDeploy"
	scaffoldJSON, _ := json.Marshal(scaffold)
	response, err := op.SendRequest(ctx, "POST", scaffoldJSON)

	if err != nil {
		return EthereumScaffold{}, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(data, scaffold)
	return scaffold, nil
}

// UpdateEthereumScaffold takes
func (op *OpenGo) UpdateEthereumScaffold(ctx context.Context, scaffold EthereumScaffold) (EthereumScaffold, error) {
	op.baseURL.Path = fmt.Sprintf("/api/ethereum-scaffolds/%s", scaffold.Address)
	scaffoldJSON, _ := json.Marshal(scaffold)
	response, err := op.SendRequest(ctx, "PATCH", scaffoldJSON)

	if err != nil {
		return EthereumScaffold{}, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(data, scaffold)
	return scaffold, nil
}

// DeactivateEthereumScaffold deactivates a EthereumScaffold and returns the deactivated EthereumScaffold
func (op *OpenGo) DeactivateEthereumScaffold(ctx context.Context, scaffoldAddress string) (EthereumScaffold, error) {
	op.baseURL.Path = fmt.Sprintf("/api/ethereum-scaffolds/%s", scaffoldAddress)
	response, err := op.SendRequest(ctx, "DELETE", nil)

	if err != nil {
		return EthereumScaffold{}, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	scaffold := &EthereumScaffold{}
	json.Unmarshal(data, scaffold)
	return *scaffold, nil
}
