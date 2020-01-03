package opengo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type WebHook struct {
	WebHook url.URL `json:"webHook"`
}

func (op *OpenGo) SetEthereumScaffoldWebHook(ctx context.Context, scaffoldAddress string, hook WebHook) (string, error) {
	op.baseURL.Path = fmt.Sprintf("/api/ethereum-scaffolds/%s", scaffoldAddress)
	webHookJSON, _ := json.Marshal(hook)
	response, err := op.SendRequest(ctx, "POST", webHookJSON)

	if err != nil {
		return "", fmt.Errorf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return string(data), nil
	}
}
