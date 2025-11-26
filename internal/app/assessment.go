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
		host,
		assessmentKey string,
		execBackgraund bool,
		get GetRequest[Analysis],
	) (Analysis, error)
}

type AssessmentStorage interface {
	Save(assessmentKey string, result Analysis) error
	Get(assessmentKey string) (Analysis, error)
}
