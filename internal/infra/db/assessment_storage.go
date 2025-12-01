package db

import (
	"maps"
	"fmt"
	"sync"

	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
)

type AssessmentManager struct {
	mu sync.RWMutex
	db map[string]app.Analysis
}

func NewAssessmentManager() AssessmentManager {
	return AssessmentManager{
		db: make(map[string]app.Analysis),
	}
}

func (a *AssessmentManager) Save(assessmentKey string, result app.Analysis) {
	a.mu.Lock()
	a.db[assessmentKey] = result
	a.mu.Unlock()
}

func (a *AssessmentManager) GetAll() map[string]app.Analysis {
	a.mu.RLock()
	defer a.mu.RUnlock()

	// devolver una copia es encapsulamiento
	copy := make(map[string]app.Analysis)
	maps.Copy(copy, a.db)
	return copy
}

func (a *AssessmentManager) GetByKey(assessmentKey string) (app.Analysis, error) {
	a.mu.RLock()
	result, ok := a.db[assessmentKey]
	a.mu.RUnlock()

	if !ok {
		return result, fmt.Errorf("%w", ErrNotFound)
	}
	return result, nil
}
