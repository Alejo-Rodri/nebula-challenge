package api

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
	"github.com/stretchr/testify/assert"
)

// PASO 1: Definición del tipo de función (si no lo hiciste antes)
type FuncGetRequest[T any] func (
    c *http.Client,
    baseUrl,
    endpoint string,
    query url.Values,
) (T, error)

// PASO 2: Implementación del método Do (si no lo hiciste antes)
func (f FuncGetRequest[T]) Do(
    c *http.Client,
    baseUrl,
    endpoint string,
    query url.Values,
) (T, error) {
    return f(c, baseUrl, endpoint, query)
}

func mockSequence(states []string, finalErr error) (app.GetRequest[app.Analysis], *int) {
	count := 0

	mockFn := func (
		c *http.Client,
		baseUrl,
		endpoint string,
		query url.Values,
	) (app.Analysis, error) {
		if finalErr != nil && count == len(states) {
			return app.Analysis{}, finalErr
		}

		s := states[count]
		count++
		return app.Analysis{Status: s}, nil
	}

	mockImplementer := FuncGetRequest[app.Analysis](mockFn)

	return mockImplementer, &count
}

func TestAnalyze(t *testing.T) {
	// mock sleep
	/* var sleepFn = func(d time.Duration) {}
    defer func() { sleepFn = time.Sleep }() */
	//sleepFn(0 * time.Second)

	client := NewApiClient("http://fake")

	tests := []struct {
		name string
		states []string
		finalErr error
		wantStatus string
		wantErr bool
		wantCalls int
	} {
		{
			name: "READY immediate",
			states: []string{"READY"},
			wantStatus: "READY",
			wantErr: false,
			wantCalls: 1,
		},
		{
            name: "ERROR immediate",
            states: []string{"ERROR"},
            wantErr: true,
            wantCalls: 1,
        },
        {
            name: "DNS then READY",
            states: []string{"DNS", "READY"},
            wantStatus: "READY",
            wantCalls: 2,
        },
        {
            name: "IN_PROGRESS then READY",
            states: []string{"IN_PROGRESS", "READY"},
            wantStatus: "READY",
            wantCalls: 2,
        },
        {
            name: "UNKNOWN then READY",
            states: []string{"WTF", "READY"},
            wantStatus: "READY",
            wantCalls: 2,
        },
        {
            name: "multiple DNS then READY",
            states: []string{"DNS", "DNS", "DNS", "READY"},
            wantStatus: "READY",
            wantCalls: 4,
        },
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T)  {
			mock, counter := mockSequence(tt.states, tt.finalErr)
			result, err := client.Analyze("example.com", "", false, mock)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.wantStatus != "" {
				// el estatus esperado definido en el struct
				assert.Equal(t, tt.wantStatus, result.Status)
			}

			// la cantidad de llamadas esperadas
			assert.Equal(t, tt.wantCalls, *counter)
		})
	}
	
}