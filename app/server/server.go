package server

import (
	"encoding/json"
	"github.com/p0poff/mock/app/circular_stack"
	"github.com/p0poff/mock/app/storage"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	port  string
	db    *storage.SQLiteDB
	stack *circular_stack.CircularStack
}

func NewServer(port string, db *storage.SQLiteDB, stack *circular_stack.CircularStack) *Server {
	s := &Server{
		port:  port,
		db:    db,
		stack: stack,
	}

	return s
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

func (s *Server) adminSaveRouteHandler(w http.ResponseWriter, r *http.Request) {
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

	if route.Id < 1 {
		err = s.db.AddRoute(route)
	} else {
		err = s.db.EditRoute(route)
	}

	if err != nil {
		log.Println("[ERROR] Failed to add route:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if route.Id < 1 {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (s *Server) adminDeleteRouteHandler(w http.ResponseWriter, r *http.Request) {
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

	err = s.db.DeleteRoute(route)
	if err != nil {
		log.Println("[ERROR] Failed to delete route:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) adminGetRouteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var ro storage.Route
	err := json.NewDecoder(r.Body).Decode(&ro)
	if err != nil {
		log.Println("[ERROR] Failed to decode JSON request:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	route, err := s.db.GetRouteById(ro.Id)
	if err != nil {
		log.Println("[ERROR] Failed to get route:", err)

		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(route)
	if err != nil {
		log.Println("[ERROR] Failed to marshal JSON response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (s *Server) adminGetRoutesHandler(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) adminLogRoutesHandler(w http.ResponseWriter, r *http.Request) {
	routers := s.stack.All()
	jsonResponse, err := json.Marshal(routers)
	if err != nil {
		log.Println("[ERROR] Failed to marshal JSON response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (s *Server) mockFaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not found", http.StatusNotFound)
}

func (s *Server) mockHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO] request uri (method): %s (%s)", r.URL.Path, r.Method)

	s.stack.Push(storage.Request{Date: time.Now().Format("2006-01-02 15:04:05"), Url: r.URL.Path, Method: r.Method})

	// Get the route from the database
	route, err := s.db.GetRoute(r.URL.Path, r.Method)
	if err != nil {
		log.Println("[ERROR] Failed to get route:", err)

		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Set the headers
	for key, value := range route.Headers {
		w.Header().Set(key, value)
	}

	// Set the status code
	w.WriteHeader(route.StatusCode)

	// Write the body
	w.Write([]byte(route.Body))
}

func (s *Server) Start() error {
	addr := ":" + s.port
	log.Printf("[INFO] Server start! port: %s", addr)

	//routing
	http.HandleFunc("/admin", s.adminindexHandler)
	http.HandleFunc("/admin/save-route", s.adminSaveRouteHandler)
	http.HandleFunc("/admin/get-routes", s.adminGetRoutesHandler)
	http.HandleFunc("/admin/get-route", s.adminGetRouteHandler)
	http.HandleFunc("/admin/delete-route", s.adminDeleteRouteHandler)
	http.HandleFunc("/admin/log-route", s.adminLogRoutesHandler)
	http.HandleFunc("/favicon.ico", s.mockFaviconHandler)
	http.HandleFunc("/", s.mockHandler)

	//static
	fs := http.FileServer(http.Dir("html"))
	http.Handle("/html/", http.StripPrefix("/html/", fs))

	return http.ListenAndServe(addr, nil)
}
