package daemon

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
	"github.com/Alejo-Rodri/nebula-challenge/internal/daemon/dto"
)

type UnixClient struct {
	client *http.Client
}

func NewUnixClient(socket string) *UnixClient {
	return &UnixClient{
		client: &http.Client{
			Transport: &http.Transport{
				DialContext: func (_ context.Context, _, _ string) (net.Conn, error) {
					return net.Dial("unix", socket)
				},
			},
		},
	}
}

// TODO handle errors
func (u *UnixClient) AddValue(assessmentKey string, assessment app.Analysis) error {
	var req = dto.AddRequest{AssessmentKey: assessmentKey, Result: assessment}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("")
	}
	
	_, err = u.client.Post("http://unix/add", "application/json", bytes.NewReader(reqBody))

	return err
}

// gets the assessment by the key :)
func (u *UnixClient) GetAssResultByKey(assessmentKey string) (app.Analysis, error) {
	return list(u, assessmentKey, mapListByKey)
}

func (u *UnixClient) ListAllValues() (app.ListAllResults, error) {
	return list(u, "", mapListAll)
}

type mapper[T any, R any] func(T) R

func list[T any, R any](u *UnixClient, assessmentKey string, m mapper[T, R]) (R, error) {
	var result R

	resp, err := u.client.Get("http://unix/list?key=" + assessmentKey)
	if err != nil {
		return result, fmt.Errorf("%w", ErrListRequest)
	}

	defer resp.Body.Close()

	var body T
	if err := parseJSON(resp.Body, &body); err != nil {
		return result, fmt.Errorf("%w", ErrParsingToJSON)
	}

	return m(body), nil
}
