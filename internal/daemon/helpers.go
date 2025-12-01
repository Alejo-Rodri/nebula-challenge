package daemon

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Alejo-Rodri/nebula-challenge/configs"
	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
	"github.com/Alejo-Rodri/nebula-challenge/internal/daemon/dto"
)

func mapListAll(l dto.ListResponseAll) app.ListAllResults {
	assessments := make([]app.AssessmentData, len(l.Assessments))

	for i, a := range l.Assessments {
		assessments[i] = app.AssessmentData{
			AssessmentKey:    a.AssessmentKey,
			AssessmentStatus: a.AssessmentStatus,
		}
	}

	return app.ListAllResults{
		Assessments: assessments,
	}
}

func mapListByKey(l dto.ListResponseKey) app.Analysis {
	endpoints := make([]app.Endpoint, len(l.Endpoints))

	for i, e := range l.Endpoints {
		endpoints[i] = app.Endpoint{
			IPAddress:         e.IPAddress,
			ServerName:        e.ServerName,
			StatusMessage:     e.StatusMessage,
			Grade:             e.Grade,
			GradeTrustIgnored: e.GradeTrustIgnored,
			HasWarnings:       e.HasWarnings,
			IsExceptional:     e.IsExceptional,
			Progress:          e.Progress,
			Duration:          e.Duration,
			ETA:               e.Eta,
			Delegation:        e.Delegation,
		}
	}

	return app.Analysis{
		Host:            l.Host,
		Port:            l.Port,
		Protocol:        l.Protocol,
		IsPublic:        l.IsPublic,
		Status:          l.Status,
		StartTime:       l.StartTime,
		TestTime:        l.TestTime,
		EngineVersion:   l.EngineVersion,
		CriteriaVersion: l.CriteriaVersion,
		Endpoints:       endpoints,
	}
}

func parseJSON(r io.Reader, v any) error {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	return json.NewDecoder(r).Decode(v)
}

func writeError(message string, w http.ResponseWriter, statusCode int) error {
	errMessage := dto.UnixError{ Message: message }

	jsonErr, err := json.Marshal(errMessage)
	if err != nil {
		return fmt.Errorf("error marshaling struct error, %w", ErrMarshaling)
	}

	_, err = w.Write(jsonErr)
	if err != nil {
		return fmt.Errorf("%w", ErrWritingBody)
	}

	w.WriteHeader(statusCode)

	log.Println(message)
	return nil
}

func (s *Store) listByKey(assessmentKey string, w http.ResponseWriter) {
	result, err := s.repo.GetByKey(assessmentKey)
	if err != nil {
		writeError("could not get an assessment with that key", w, http.StatusNotFound)
		return
	}

	resBody, err := json.Marshal(result)
	if err != nil {
		writeError("error marshaling the response", w, http.StatusInternalServerError)
		return
	}

	_, err = w.Write(resBody)
	if err != nil {
		writeError("error writing response body", w, http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Store) listAllResults(w http.ResponseWriter) {
	var response dto.ListResponseAll
	
	for key, result := range s.repo.GetAll() {
		response.Assessments = append(response.Assessments, 
			dto.AssessmentStatus{
				AssessmentKey: key,
				AssessmentStatus: result.Status,
			},
		)
	}

	resBody, err := json.Marshal(response)
	if err != nil {
		writeError("error marshaling the response", w, http.StatusInternalServerError)
		return
	}

	_, err = w.Write(resBody)
	if err != nil {
		writeError("error writing response body", w, http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Store) processQueue() {
	max := configs.Envs.ConcurrentRequestsRetry

	for range max {
		front := s.queue.Front()
		if front == nil {
			return
		}

		ass := front.Value.(*AssessmentTime)

		if time.Since(ass.time) < 15*time.Second {
			s.queue.MoveToBack(front)
			continue
		}

		result, err := s.analyzer.Analyze(ass.host)
		if err != nil {
			log.Println("request error:", err)
			continue
		}

		s.queue.Remove(front)

		dbAss, err := s.repo.GetByKey(ass.key)
		if err != nil {
			log.Println("repo error:", err)
			continue
		}

		if result.Status != dbAss.Status {
			s.repo.Save(ass.key, result)
		} else {
			s.queue.PushBack(&AssessmentTime{
				key:    ass.key,
				host:   ass.host,
				status: result.Status,
				time:   time.Now(),
			})
		}
	}
}

