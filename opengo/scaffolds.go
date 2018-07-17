package opengo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type scaffoldResponse struct {
	TotalCount int         `json:"totalCount"`
	Scaffolds  []*Scaffold `json:"list"`
}

type Scaffold struct {
	OpenKey          string              `json:"openKey,omitempty"`
	Address          string              `json:"address,omitempty"`
	ABI              string              `json:"abi,omitempty"`
	Description      string              `json:"description,omitempty"`
	FiatAmount       string              `json:"fiatAmount,omitempty"`
	Currency         string              `json:"currency,omitempty"`
	ConversionAmount string              `json:"conversionAmount,omitempty"`
	DeveloperAddress string              `json:"developerAddress,omitempty"`
	WebHook          string              `json:"webHook,omitempty"`
	Properties       []*ScaffoldProperty `json:"properties,omitempty"`
	Enabled          bool                `json:"enabled,omitempty"`
}

// AddProperty adds ScaffoldProperty to Scaffold object.
func (scaffold *Scaffold) AddProperty(Name, Type, DefaultValue string) error {
	scaffoldProperty := NewScaffoldProperty(Name, Type, DefaultValue)
	scaffold.Properties = append(scaffold.Properties, scaffoldProperty)
	return nil
}

type ScaffoldSummary struct {
	Scaffold         *Scaffold      `json:"scaffold"`
	TransactionIndex int64          `json:"transactionIndex"`
	TokenBalance     int64          `json:"tokenBalance"`
	Enabled          bool           `json:"enabled"`
	ShareHolders     []*ShareHolder `json:"shareHolders"`
}

type ScaffoldProperty struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	DefaultValue string `json:"defaultValue"`
}

// NewScaffoldProperty return a new object of ScaffoldProperty
func NewScaffoldProperty(Name, Type, DefaultValue string) *ScaffoldProperty {
	scaffoldProperty := &ScaffoldProperty{
		Name:         Name,
		Type:         Type,
		DefaultValue: DefaultValue,
	}
	return scaffoldProperty
}

// NewScaffold return a new object of NewScaffold
func NewScaffold(OpenKey, Address, ABI, Description, FiatAmount, Currency, ConversionAmount, DeveloperAddress, Webhook string, Enabled bool) *Scaffold {
	scaffold := &Scaffold{
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

// GetScaffolds function retrieves all []*Scaffolds from Open Platform and returns a list of Scaffold
func (op *OpenGo) GetScaffolds(ctx context.Context) ([]*Scaffold, error) {
	op.baseURL.Path = "/api/scaffolds"
	response, err := op.SendRequest(ctx, "GET", nil)

	if err != nil {
		return nil, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	sr := &scaffoldResponse{}
	json.Unmarshal(data, sr)
	return sr.Scaffolds, nil
}

// GetScaffold returns *Scaffold given it's address
func (op *OpenGo) GetScaffold(ctx context.Context, address string) (*Scaffold, error) {
	op.baseURL.Path = fmt.Sprintf("/api/scaffolds/%s", address)
	response, err := op.SendRequest(ctx, "GET", nil)

	if err != nil {
		return &Scaffold{}, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	scaffold := &Scaffold{}
	json.Unmarshal(data, scaffold)
	return scaffold, nil
}

// GetScaffoldSummary retrieves scaffold summary from Open Platform API and returns *ScaffoldSummary
func (op *OpenGo) GetScaffoldSummary(ctx context.Context, address string) (*ScaffoldSummary, error) {
	op.baseURL.Path = fmt.Sprintf("/api/scaffolds/%s/%s", address, "summary")
	response, err := op.SendRequest(ctx, "GET", nil)

	if err != nil {
		return &ScaffoldSummary{}, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	scaffoldSummary := &ScaffoldSummary{}
	json.Unmarshal(data, scaffoldSummary)
	return scaffoldSummary, nil
}

// DeployScaffold takes a Scaffold and sends it to Open Platform. If successful, returns given Scaffold.
func (op *OpenGo) DeployScaffold(ctx context.Context, scaffold Scaffold) (Scaffold, error) {
	op.baseURL.Path = "/api/scaffolds/doDeploy"
	scaffoldJSON, _ := json.Marshal(scaffold)
	response, err := op.SendRequest(ctx, "POST", scaffoldJSON)

	if err != nil {
		return Scaffold{}, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(data, scaffold)
	return scaffold, nil
}

// UpdateScaffold takes
func (op *OpenGo) UpdateScaffold(ctx context.Context, scaffold Scaffold) (Scaffold, error) {
	op.baseURL.Path = fmt.Sprintf("/api/scaffolds/%s", scaffold.Address)
	scaffoldJSON, _ := json.Marshal(scaffold)
	response, err := op.SendRequest(ctx, "PATCH", scaffoldJSON)

	if err != nil {
		return Scaffold{}, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(data, scaffold)
	return scaffold, nil
}

// DeactivateScaffold deactivates a Scaffold and returns the deactivated Scaffold
func (op *OpenGo) DeactivateScaffold(ctx context.Context, scaffoldAddress string) (Scaffold, error) {
	op.baseURL.Path = fmt.Sprintf("/api/scaffolds/%s", scaffoldAddress)
	response, err := op.SendRequest(ctx, "DELETE", nil)

	if err != nil {
		return Scaffold{}, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	scaffold := &Scaffold{}
	json.Unmarshal(data, scaffold)
	return *scaffold, nil
}
