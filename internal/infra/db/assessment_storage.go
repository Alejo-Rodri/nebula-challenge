package db

import (
	"fmt"

	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
)

type AssessmentManager struct {
	db map[string]app.Analysis
}

func NewAssessmentManager() AssessmentManager {
	return AssessmentManager{
		db: make(map[string]app.Analysis),
	}
}

// TODO assManager deberia ofrecer un lock que le permita a muchos la lectura pero solo a uno la escritura
func (a *AssessmentManager) Get(assessmentKey string) (app.Analysis, error) {
	result, ok := a.db[assessmentKey]
	if !ok {
		return result, fmt.Errorf("%w", ErrNotFound)
	}

	return result, nil
}

func (a *AssessmentManager) Save(assessmentKey string, result app.Analysis) error {
	a.db[assessmentKey] = result

	return nil
}

