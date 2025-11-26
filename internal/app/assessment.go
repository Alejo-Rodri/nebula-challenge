package app

import (
	"github.com/Alejo-Rodri/nebula-challenge/internal/infra/api"
)

type AssessmentApp interface {
	
	Info(get api.GetAbstractRequest[api.ApiInfoResponse]) (api.ApiInfoResponse, error)
	Analyze(host string, get api.GetAbstractRequest[api.ApiAnalyzeResponse]) (api.ApiAnalyzeResponse, error)

}
