package app

import (
	"github.com/Alejo-Rodri/nebula-challenge/internal/infra/api"
)

type AssessmentApp interface {
	Info() (api.ApiInfoResponse, error)
	Analyze(host string) (api.ApiAnalyzeResponse, error)
}

type App struct {
	ApiClient *api.ApiClient
}

func NewApp() *App {
	return &App{
		ApiClient: api.NewApiClient(""),
	}
}

func (a *App) Info() (api.ApiInfoResponse, error) {
	return a.ApiClient.Info()
}

func (a *App) Analyze(host string) (api.ApiAnalyzeResponse, error) {
	return a.ApiClient.Analyze(host)
}