package httpapi

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"library/internal/models"
	"library/internal/services"
)

type server struct {
	svc *services.LibraryService
	mux *http.ServeMux
}

func NewServer() http.Handler {
	s := &server{svc: services.NewLibraryService(), mux: http.NewServeMux()}
	s.routes()
	return cors(s.mux)
}

func (s *server) routes() {
	s.mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	s.mux.HandleFunc("/api/users", s.handleUsers)
	s.mux.HandleFunc("/api/books", s.handleBooks)
	s.mux.HandleFunc("/api/books/search", s.handleBookSearch)
	s.mux.HandleFunc("/api/loans/borrow", s.handleBorrow)
	s.mux.HandleFunc("/api/loans/return", s.handleReturn)
}

func (s *server) handleUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var u models.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if u.ID == "" || u.Name == "" {
			http.Error(w, "missing fields", 400)
			return
		}
		s.svc.AddUser(u)
		respond(w, 201, u)
		return
	}
	if r.Method == http.MethodDelete {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "missing id", 400)
			return
		}
		if err := s.svc.RemoveUser(id); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		respond(w, 200, map[string]string{"status": "deleted"})
		return
	}
	if r.Method == http.MethodGet {
		respond(w, 200, s.svc.ListUsers())
		return
	}
	http.NotFound(w, r)
}

func (s *server) handleBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var b models.Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if b.ID == "" || b.Title == "" || b.Author == "" {
			http.Error(w, "missing fields", 400)
			return
		}
		s.svc.AddBook(b)
		respond(w, 201, b)
		return
	}
	if r.Method == http.MethodDelete {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "missing id", 400)
			return
		}
		if err := s.svc.RemoveBook(id); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		respond(w, 200, map[string]string{"status": "deleted"})
		return
	}
	if r.Method == http.MethodGet {
		respond(w, 200, s.svc.ListBooks())
		return
	}
	http.NotFound(w, r)
}

func (s *server) handleBookSearch(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	respond(w, 200, s.svc.SearchBooks(q))
}

func (s *server) handleBorrow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	var req models.LoanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if req.UserID == "" || req.BookID == "" {
		http.Error(w, "missing fields", 400)
		return
	}
	if err := s.svc.Borrow(req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	respond(w, 200, map[string]string{"status": "borrowed"})
}

func (s *server) handleReturn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	var req models.LoanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if req.UserID == "" || req.BookID == "" {
		http.Error(w, "missing fields", 400)
		return
	}
	if err := s.svc.Return(req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	respond(w, 200, map[string]string{"status": "returned"})
}

func respond(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("encode error:", err)
	}
}

func cors(next http.Handler) http.Handler {
	origin := os.Getenv("CORS_ORIGIN")
	if origin == "" {
		origin = "*"
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
