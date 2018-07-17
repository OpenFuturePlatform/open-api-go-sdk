package opengo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	op.baseURL.Path = fmt.Sprintf("/api/scaffolds/%s/%s", scaffoldAddress, "holders")
	holderJSON, _ := json.Marshal(holder)
	request, _ := http.NewRequest("POST", op.baseURL.String(), bytes.NewBuffer(holderJSON))
	request = request.WithContext(ctx)
	request.Header.Set("Authorization", op.token)
	request.Header.Set("Content-Type", "application/json")

	response, err := op.httpClient.Do(request)

	if err != nil {
		return "", fmt.Errorf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return string(data), nil
}

func (op *OpenGo) DeleteShareHolder(ctx context.Context, scaffoldAddress string, holderAddress string) (string, error) {
	op.baseURL.Path = fmt.Sprintf("/api/scaffolds/%s/%s/%s", scaffoldAddress, "holders", holderAddress)
	request, _ := http.NewRequest("DELETE", op.baseURL.String(), nil)
	request = request.WithContext(ctx)
	request.Header.Set("Authorization", op.token)
	request.Header.Set("Content-Type", "application/json")

	response, err := op.httpClient.Do(request)

	if err != nil {
		return "", fmt.Errorf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return string(data), nil
}

func (op *OpenGo) UpdateShareHolder(ctx context.Context, scaffoldAddress string, holder ShareHolder) (string, error) {
	op.baseURL.Path = fmt.Sprintf("/api/scaffolds/%s/%s/%s", scaffoldAddress, "holders", holder.Address)
	holder.Address = ""
	holderJSON, _ := json.Marshal(holder)
	request, _ := http.NewRequest("PUT", op.baseURL.String(), bytes.NewBuffer(holderJSON))
	request = request.WithContext(ctx)
	request.Header.Set("Authorization", op.token)
	request.Header.Set("Content-Type", "application/json")

	response, err := op.httpClient.Do(request)

	if err != nil {
		return "", fmt.Errorf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return string(data), nil
}
