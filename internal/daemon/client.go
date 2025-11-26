package daemon

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

func newUnixClient(socket string) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: func (_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", socket)
			},
		},
	}
}

func AddValue(socket, v string) error {
	// TODO handle errors
	body, _ := json.Marshal(map[string]string{"Value": v})
	
	client := newUnixClient(socket)

	_, err := client.Post("http://unix/add", "application/json", bytes.NewReader(body))

	return err
}

func ListValues(socket string) ([]string, error) {
	client := newUnixClient(socket)

	resp, err := client.Get("http://unix/list")
	if err != nil {
		return nil, fmt.Errorf("%w", ErrListRequest)
	}

	defer resp.Body.Close()

	var data []string
	json.NewDecoder(resp.Body).Decode(&data)

	return data, nil
}
