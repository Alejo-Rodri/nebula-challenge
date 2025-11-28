package daemon

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
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

func (u *UnixClient) AddValue(v string) error {
	// TODO handle errors
	body, _ := json.Marshal(map[string]string{"Value": v})
	
	_, err := u.client.Post("http://unix/add", "application/json", bytes.NewReader(body))

	return err
}

func (u *UnixClient) ListValues() ([]string, error) {
	resp, err := u.client.Get("http://unix/list")
	if err != nil {
		return nil, fmt.Errorf("%w", ErrListRequest)
	}

	defer resp.Body.Close()

	var data []string
	json.NewDecoder(resp.Body).Decode(&data)

	return data, nil
}
