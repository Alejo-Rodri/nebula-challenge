package daemon

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
)

type Store struct {
	mu sync.Mutex
	repo app.AssessmentStorage
	Data []string
}

func (s *Store) add(w http.ResponseWriter, r *http.Request) {
	var body struct{ Value string }
	json.NewDecoder(r.Body).Decode(&body)

	s.mu.Lock()
	s.Data = append(s.Data, body.Value)
	s.mu.Unlock()

	w.WriteHeader(http.StatusOK)
}

func (s *Store) list(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	resp := append([]string(nil), s.Data...)
	s.mu.Unlock()

	json.NewEncoder(w).Encode(resp)
}

func RunServer(socket string) {
	// deletes the socket if already exists
	if _, err := os.Stat(socket); err == nil {
		os.Remove(socket)
	}

	store := &Store{}

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
