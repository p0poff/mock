package server

import (
	"encoding/json"
	"github.com/p0poff/mock/app/storage"
	"log"
	"net/http"
	"os"
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
	method := r.Method

	log.Printf("[INFO] request method: %s", method)

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

func (s *Server) adminindexHandler(w http.ResponseWriter, r *http.Request) {
	// Open the HTML file
	file, err := os.Open("./html/index.html")

	if err != nil {
		// If there's an error opening the file, log it and send an internal server error
		log.Println("Error opening file:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		// If there's an error getting file info, log it and send an internal server error
		log.Println("Error getting file info:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the header to HTML (this is not always necessary but good for correctness)
	w.Header().Set("Content-Type", "text/html")

	// Serve the HTML file
	http.ServeContent(w, r, "index.html", fileInfo.ModTime(), file)
}

func (s *Server) adminAddRouteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var route storage.Route
	err := json.NewDecoder(r.Body).Decode(&route)
	if err != nil {
		log.Println("[ERROR] Failed to decode JSON request:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = s.db.AddRoute(route)
	if err != nil {
		log.Println("[ERROR] Failed to add route:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) getRoutersHandler(w http.ResponseWriter, r *http.Request) {
	routers, err := s.db.GetRoutes()
	if err != nil {
		log.Println("[ERROR] Failed to get routers:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(routers)
	if err != nil {
		log.Println("[ERROR] Failed to marshal JSON response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (s *Server) Start() error {
	addr := ":" + s.port
	log.Printf("[INFO] Server start! port: %s", addr)

	//routing
	http.HandleFunc("/hello", s.helloHandler)
	http.HandleFunc("/", s.defaultHandler)
	http.HandleFunc("/admin", s.adminindexHandler)
	http.HandleFunc("/admin/add-route", s.adminAddRouteHandler)
	http.HandleFunc("/admin/get-routers", s.getRoutersHandler)

	//static
	fs := http.FileServer(http.Dir("html"))
	http.Handle("/html/", http.StripPrefix("/html/", fs))

	return http.ListenAndServe(addr, nil)
}

func (s *Server) Stop() {
	// Add your server shutdown logic here
}
