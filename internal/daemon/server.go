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
}

func (s *Store) add(w http.ResponseWriter, r *http.Request) {
	log.Println("received request /add")
	var body struct{ Value string }
	json.NewDecoder(r.Body).Decode(&body)
	
	if err := parseJSON(r.Body, body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	/* s.mu.Lock()
	s.repo.Save("", app.Analysis{})
	s.Data = append(s.Data, body.Value)
	s.mu.Unlock()

	w.WriteHeader(http.StatusOK) */
}

func (s *Store) list(w http.ResponseWriter, r *http.Request) {
	log.Println("received request /list")
	s.mu.Lock()
	s.repo.Get("")
	//resp := append([]string(nil), s.Data...)
	s.mu.Unlock()

	//json.NewEncoder(w).Encode(resp)
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
