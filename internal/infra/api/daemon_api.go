package api

import (
	"net/http"
	"net/url"

	"github.com/Alejo-Rodri/nebula-challenge/configs"
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
    endpoint := "/analyze"

    base := configs.Envs.BaseApiURL

    baseURL, err := url.Parse(base)
    if err != nil {
        return app.Analysis{}, printError("GET", endpoint, ErrParsingUrlRequest, err)
    }

    query := url.Values{}
    query.Set("host", host)
    query.Set("all", "done")

    result, err := Get[app.Analysis](d.client, baseURL.String(), endpoint, query)
    if err != nil {
        return result, err
    }

    return result, nil
}

