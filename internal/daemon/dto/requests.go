package dto

import "github.com/Alejo-Rodri/nebula-challenge/internal/app"

type AddRequest struct {
	AssessmentKey string 
	Result app.Analysis
}

// si va vacio imprime todos los resultados
type ListRequest struct {
	AssessmentKey string
}

