package api

import (
	"net/http"
	"net/url"

	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
)

type DaemonApi struct {
	client *http.Client
}

func NewDaemonApi(client *http.Client) *DaemonApi {
	return &DaemonApi{
		client: client,
	}
}

func (d *DaemonApi) Analyze(host string) (app.Analysis, error) {

	var endpoint string = "/analyze"

	baseURL, err := url.Parse(host + endpoint)
	if err != nil {
		return app.Analysis{}, printError("GET", endpoint, ErrParsingUrlRequest, err)
	}

	// first request
	// This parameter should be used only once to initiate a new assessment; further invocations should omit it to avoid causing an assessment loop.
	query := baseURL.Query()
	query.Set("host", host)
	query.Set("all", "done")

	result, err := Get[app.Analysis](d.client, host, endpoint, query)
	if err != nil {
		return result, err
	}

	return result, nil
}
