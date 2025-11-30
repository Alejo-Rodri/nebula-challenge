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

func (a *AssessmentManager) Save(assessmentKey string, result app.Analysis) {
	a.db[assessmentKey] = result
}

func (a *AssessmentManager) GetAll() map[string]app.Analysis {
	return a.db
}

func (a *AssessmentManager) GetByKey(assessmentKey string) (app.Analysis, error) {
	result, ok := a.db[assessmentKey]
	if !ok {
		return result, fmt.Errorf("%w", ErrNotFound)
	}

	return result, nil
}
