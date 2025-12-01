package api

import (
	"net/http"
	"net/url"
	"time"

	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
)

type GetInfoRequest struct {
    client  *ApiClient
}

func NewInfoRequest(client *ApiClient) GetInfoRequest {
	return GetInfoRequest{
		client: client,
	}
}

func (r GetInfoRequest) Do(
	c *http.Client,
	baseUrl,
	endpoint string,
	query url.Values,
) (app.Info, error) {
    ext, err := Get[ApiInfoResponse](c, baseUrl, endpoint, query)
    if err != nil {
        return app.Info{}, err
    }

    // mapea infra → dominio
    return mapInfo(ext), nil
}

type GetAnalyzeRequest struct {
	client *ApiClient
}

func NewAnalyzeRequest(client *ApiClient) GetAnalyzeRequest {
	return GetAnalyzeRequest{
		client: client,
	}
}

func (r GetAnalyzeRequest) Do(
	c *http.Client,
	baseUrl,
	endpoint string,
	query url.Values,
) (app.Analysis, error) {
	ext, err := Get[ApiAnalyzeResponse](c, baseUrl, endpoint, query)

	if err != nil {
        return app.Analysis{}, err
    }

    // mapea infra → dominio
    return mapAnalysis(ext), nil
}

type ApiClient struct {
	baseURL string
	Http *http.Client
}

func NewApiClient(baseURL string) *ApiClient {
	return &ApiClient{
		baseURL: baseURL,
		Http: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}