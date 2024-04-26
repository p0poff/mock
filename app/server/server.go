package server

import (
	"encoding/json"
	"github.com/p0poff/mock/app/storage"
	"log"
	"net/http"
)

type Server struct {
	port string
	db   *storage.SQLiteDB
}

func NewServer(port string, db *storage.SQLiteDB) *Server {
	s := &Server{
		port: port,
		db:   db,
	}

	return s
}

func (s *Server) helloHandler(w http.ResponseWriter, r *http.Request) {
	// Add your hello logic here
	_ = s.db.AddProduct("Apple", 11)
	w.Write([]byte("Hello, World!"))
}

func (s *Server) defaultHandler(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Path

	// Return JSON response
	response := map[string]string{"uri": uri}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("[ERROR] Failed to marshal JSON response:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (s *Server) Start() error {
	addr := ":" + s.port
	log.Printf("[INFO] Server start! port: %s", addr)

	http.HandleFunc("/hello", s.helloHandler)
	http.HandleFunc("/", s.defaultHandler)

	return http.ListenAndServe(addr, nil)
}

func (s *Server) Stop() {
	// Add your server shutdown logic here
}
