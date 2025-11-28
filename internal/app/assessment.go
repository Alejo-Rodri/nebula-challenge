package app

import (
	"net/http"
	"net/url"
)

type GetRequest[T any] interface {
	Do(
		c *http.Client,
		baseUrl,
		endpoint string,
		query url.Values,
	) (T, error)
}

type AssessmentApp interface {
	Info(get GetRequest[Info]) (Info, error)
	Analyze(
		host string,
		execBackgraund bool,
		get GetRequest[Analysis],
	) (Analysis, error)
}

type AssessmentStorage interface {
	Save(assessmentKey string, result Analysis) error
	GetAll() map[string]Analysis
	GetByKey(assessmentKey string) (Analysis, error)
}
