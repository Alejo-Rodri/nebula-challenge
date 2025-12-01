package daemon

import (
	l "container/list"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
	"github.com/Alejo-Rodri/nebula-challenge/internal/daemon/dto"
)

type Store struct {
	qmu sync.Mutex
	repo app.AssessmentStorage
	queue *l.List
	analyzer app.Analize
}

type AssessmentTime struct {
	key string
	host string
	status string
	time time.Time
}

func (s *Store) add(w http.ResponseWriter, r *http.Request) {
	log.Println("received request /add")

	var body dto.AddRequest
	if err := parseJSON(r.Body, &body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	status := body.Result.Status
	if status == "IN_PROGRESS" || status == "DNS" {
		s.queue.PushBack(&AssessmentTime{
			key:    body.AssessmentKey,
			host:   body.Result.Host,
			status: status,
			time:   time.Now(),
		})
	}

	s.repo.Save(body.AssessmentKey, body.Result)

	w.WriteHeader(http.StatusOK)
}

func (s *Store) list(w http.ResponseWriter, r *http.Request) {
	log.Println("received request /list")
	q := r.URL.Query()
	assKey := q.Get("key")

	if assKey != "" {
		s.listByKey(assKey, w)
		return
	}

	// solamente actualiza si se imprimen todos los assessments
	s.qmu.Lock()
	s.processQueue()
	s.qmu.Unlock()

	s.listAllResults(w)
}


func RunServer(socket string, db app.AssessmentStorage, analyzer app.Analize) {
	// deletes the socket if already exists
	if _, err := os.Stat(socket); err == nil {
		os.Remove(socket)
	}

	store := &Store{
		repo: db,
		queue: l.New(),
		analyzer: analyzer,
	}

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
