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
	mu sync.Mutex
	repo app.AssessmentStorage
	queue *l.List
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
	if err := parseJSON(r.Body, body); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if body.Result.Status == "IN_PROGRESS" || body.Result.Status == "DNS" {
		s.queue.PushBack(&AssessmentTime{
			key: body.AssessmentKey,
			host: body.Result.Host,
			status: body.Result.Status,
			time: time.Now(),
		})
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

	// TODO volver ese dos una variable de entorno
	// Todo es mejor aumentar el contador del bucle cada vez que se haya hecho la request porque puede que las condiciones no se cumplan
	for range 2 {
		front := s.queue.Front()
		ass := front.Value.(AssessmentTime)

		if time.Since(ass.time).Seconds() > 15 {
			// TODO aqui se llama a la request
			var result app.Analysis
			s.queue.Remove(front)

			dbAss, err := s.repo.GetByKey(ass.key)
			if err != nil {
				// TODO
			}

			if result.Status != dbAss.Status {
				s.mu.Lock()
				s.repo.Save(ass.key, result)
				s.mu.Unlock()
			} else {
				s.queue.PushBack(&AssessmentTime{
					key: ass.key,
					host: ass.host,
					status: result.Status,
					time: time.Now(),
				})
			}
		}
		
	}

	s.listAllResults(w)
}

func RunServer(socket string, db app.AssessmentStorage) {
	// deletes the socket if already exists
	if _, err := os.Stat(socket); err == nil {
		os.Remove(socket)
	}

	store := &Store{
		repo: db,
		queue: l.New(),
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
