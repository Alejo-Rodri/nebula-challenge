package api

import (
	"net/url"
	"testing"

	"github.com/Alejo-Rodri/nebula-challenge/configs"
	"github.com/stretchr/testify/assert"
)

func mockAbstractGetError(c *ApiClient, endpoint string, query url.Values) (ApiAnalyzeResponse, error) {
	return ApiAnalyzeResponse{
		Status: "ERROR",
	}, nil
}

func mockAbstractGetReady(c *ApiClient, endpoint string, query url.Values) (ApiAnalyzeResponse, error) {
	return ApiAnalyzeResponse{
		Status: "READY",
	}, nil
}


func TestAnalyze(t *testing.T) {
	client := NewApiClient(configs.Envs.BaseApiURL)

	t.Run("valid ERROR state", func(t *testing.T) {
		_, err := client.Analyze("", mockAbstractGetError)
		assert.Error(t, err)
	})

	expectedReadyState := ApiAnalyzeResponse{
		Status: "READY",
	}

	t.Run("valid READY state", func(t *testing.T) {
		result, err := client.Analyze("", mockAbstractGetReady)
		assert.NoError(t, err)
		assert.Equal(t, expectedReadyState, result)
	})
}