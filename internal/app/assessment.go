package app

import (
	"github.com/Alejo-Rodri/nebula-challenge/internal/infra/api"
)

type AssessmentApp interface {
	Info() (api.ApiInfoResponse, error)
	Analyze(host string) (api.ApiAnalyzeResponse, error)
}