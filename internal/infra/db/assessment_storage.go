package db

import (
	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
)

type AssessmentManager struct {

}

func NewAssessmentManager() AssessmentManager {
	return AssessmentManager{}
}

func (a *AssessmentManager) Get(assessmentKey string) (app.Analysis, error) {
	var result app.Analysis
	return result, nil
}

func (a *AssessmentManager) Save(assessmentKey string, result app.Analysis) error {
	return nil
}

