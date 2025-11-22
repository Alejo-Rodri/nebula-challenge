package api

import (
	"net/http"
	"time"
)

type ApiClient struct {
	baseURL string
	http *http.Client
}

func NewApiClient(baseURL string) *ApiClient {
	return &ApiClient{
		baseURL: baseURL,
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}