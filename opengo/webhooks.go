package opengo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type WebHook struct {
	WebHook url.URL `json:"webHook"`
}

func (op *OpenGo) SetWebHook(ctx context.Context, scaffoldAddress string, hook WebHook) (string, error) {
	op.baseURL.Path = fmt.Sprintf("/api/scaffolds/%s", scaffoldAddress)
	webHookJSON, _ := json.Marshal(hook)
	request, _ := http.NewRequest("POST", op.baseURL.String(), bytes.NewBuffer(webHookJSON))
	request = request.WithContext(ctx)
	request.Header.Set("Authorization", op.token)
	request.Header.Set("Content-Type", "application/json")

	response, err := op.httpClient.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return "", err
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return string(data), nil
	}
}
