package opengo

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
)

// NewOpenGo creates OpenGo objects
func NewOpenGo(token string) (client *OpenGo) {
	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "api.openfuture.io",
	}
	op := &OpenGo{token: token, baseURL: baseUrl, httpClient: &http.Client{}}
	return op
}

// OpenGo is an object for communicating with Open Platform API
type OpenGo struct {
	token      string
	baseURL    *url.URL
	httpClient *http.Client
}

// SendRequest function makes requests to Open Platform API
func (op *OpenGo) SendRequest(ctx context.Context, method string, data []byte) (*http.Response, error) {
	var body []byte
	if data != nil {
		body = data
	}
	request, _ := http.NewRequest(method, op.baseURL.String(), bytes.NewBuffer(body))
	request = request.WithContext(ctx)
	request.Header.Set("Authorization", op.token)
	request.Header.Set("Content-Type", "application/json")
	response, err := op.httpClient.Do(request)
	return response, err
}
