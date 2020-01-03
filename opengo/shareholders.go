package opengo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ShareHolder struct {
	Address string `json:"address,omitempty"`
	Percent int    `json:"percent,omitempty"`
}

func NewShareHolder(address string, percent int) *ShareHolder {
	holder := &ShareHolder{
		Address: address,
		Percent: percent,
	}
	return holder
}

func (op *OpenGo) AddShareHolder(ctx context.Context, scaffoldAddress string, holder ShareHolder) (string, error) {
	op.baseURL.Path = fmt.Sprintf("/api/ethereum-scaffolds/%s/%s", scaffoldAddress, "holders")
	holderJSON, _ := json.Marshal(holder)
	response, err := op.SendRequest(ctx, "POST", holderJSON)

	if err != nil {
		return "", fmt.Errorf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return string(data), nil
}

func (op *OpenGo) DeleteShareHolder(ctx context.Context, scaffoldAddress string, holderAddress string) (string, error) {
	op.baseURL.Path = fmt.Sprintf("/api/ethereum-scaffolds/%s/%s/%s", scaffoldAddress, "holders", holderAddress)
	response, err := op.SendRequest(ctx, "DELETE", nil)

	if err != nil {
		return "", fmt.Errorf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return string(data), nil
}

func (op *OpenGo) UpdateShareHolder(ctx context.Context, scaffoldAddress string, holder ShareHolder) (string, error) {
	op.baseURL.Path = fmt.Sprintf("/api/ethereum-scaffolds/%s/%s/%s", scaffoldAddress, "holders", holder.Address)
	holder.Address = ""
	holderJSON, _ := json.Marshal(holder)
	response, err := op.SendRequest(ctx, "PUT", holderJSON)

	if err != nil {
		return "", fmt.Errorf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return string(data), nil
}
