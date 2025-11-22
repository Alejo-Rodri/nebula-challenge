package cmd

import (
	"errors"
	"fmt"

	"github.com/Alejo-Rodri/nebula-challenge/internal/infra/api"
)

func HumanizeError(err error) string {
	switch {
	case errors.Is(err, api.ErrConnection):
		return "couldn't connect to the server"
	case errors.Is(err, api.ErrInvalidRequest):
		return "internal app error"
	case errors.Is(err, api.ErrInvalidResponse):
		return "unexpected server response"
	case errors.Is(err, api.ErrTimeout):
		return "the server took too long to respond"
	default:
		return fmt.Sprintf("unexpected error: %v", err)
	}
}
