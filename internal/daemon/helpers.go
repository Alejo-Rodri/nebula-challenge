package daemon

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Alejo-Rodri/nebula-challenge/internal/daemon/dto"
)

func parseJSON(r io.Reader, v any) error {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	return json.NewDecoder(r).Decode(v)
}

// TODO improve errors
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

func (s *Store) listByKey(r *dto.ListRequest, w http.ResponseWriter) {
// TODO que pasa si hago return sin soltar el lock?
	s.mu.Lock()
	result, err := s.repo.GetByKey(r.AssessmentKey)
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

	s.mu.Unlock()
}

func (s *Store) listAllResults(w http.ResponseWriter) {
	var response dto.ListResponseAll
	
	s.mu.Lock()

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

	s.mu.Unlock()
}
