package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Alejo-Rodri/nebula-challenge/types"
)

func (c *ApiClient) Info() error {
	req, err := c.http.Get(c.baseURL + "/info")
	if err != nil {
		return fmt.Errorf("%w: %s", ErrConnection, err)
	}

	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		var error_type types.ApiErrorsResponse
		if err := json.NewDecoder(req.Body).Decode(&error_type); err != nil {
			return err
		}

		return fmt.Errorf("%w: %s", ErrInvalidRequest, error_type)
	}

	var resp types.ApiInfoResponse
	if err := json.NewDecoder(req.Body).Decode(&resp); err != nil {
		return err
	}


	fmt.Println(resp)

	return nil
}