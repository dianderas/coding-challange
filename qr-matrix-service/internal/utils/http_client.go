package utils

import (
	"errors"
	"strings"

	"github.com/go-resty/resty/v2"
)

type HTTPClient struct {
	Client *resty.Client
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		Client: resty.New(),
	}
}

func (h *HTTPClient) MakeAuthenticatedRequest(method, url, token string, body interface{}) (map[string]interface{}, error) {
	var response map[string]interface{}
	authToken := strings.TrimPrefix(token, "Bearer ")
	resp, err := h.Client.R().
		SetAuthToken(authToken).
		SetBody(body).
		SetResult(&response).
		Execute(method, url)

	if err != nil {
		return nil, errors.New("error making request: " + err.Error())
	}

	if resp.IsError() {
		return nil, errors.New("request return error: " + string(resp.Body()))
	}

	return response, nil
}
