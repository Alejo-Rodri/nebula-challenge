package daemon

import (
	//"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
	"github.com/Alejo-Rodri/nebula-challenge/internal/daemon/dto"
)

type Store struct {
	mu sync.Mutex
	repo app.AssessmentStorage
}

func (s *Store) add(w http.ResponseWriter, r *http.Request) {
	log.Println("received request /add")
	
	var body dto.AddRequest
	if err := parseJSON(r.Body, body); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	s.repo.Save(body.AssessmentKey, body.Result)
	s.mu.Unlock()

	w.WriteHeader(http.StatusOK)
}

func (s *Store) list(w http.ResponseWriter, r *http.Request) {
	log.Println("received request /list")

	var reqBody dto.ListRequest
	if err := parseJSON(r.Body, reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if reqBody.AssessmentKey != "" {
		s.listByKey(&reqBody, w)
		return
	}

	s.listAllResults(w)
}

func RunServer(socket string, db app.AssessmentStorage) {
	// deletes the socket if already exists
	if _, err := os.Stat(socket); err == nil {
		os.Remove(socket)
	}

	store := &Store{repo: db}

	mux := http.NewServeMux()
	mux.HandleFunc("/add", store.add)
	mux.HandleFunc("/list", store.list)

	l, err := net.Listen("unix", socket)
	if err != nil {
		log.Fatal("listen unix:", err)
	}

	log.Println("daemon running on", socket)
	http.Serve(l, mux)
}
