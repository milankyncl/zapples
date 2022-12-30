package http

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/milankyncl/feature-toggles/internal/http/handler"
	"github.com/milankyncl/feature-toggles/internal/storage"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type Server struct {
	storage storage.Adapter
}

func NewServer(storage storage.Adapter) *Server {
	return &Server{
		storage,
	}
}

func (s *Server) router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Mount("/api", s.apiRouter())
	r.Mount("/", s.uiRouter())
	return r
}

func (s *Server) apiRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/features", handler.GetFeaturesHandler(s.storage))
	r.Post("/features", handler.CreateFeatureHandler(s.storage))
	r.Put("/features/{id}", handler.UpdateFeatureHandle(s.storage))
	r.Put("/features/{id}/toggle", handler.ToggleFeatureHandle(s.storage))
	return r
}

func (s *Server) uiRouter() *chi.Mux {
	r := chi.NewRouter()
	wd, _ := os.Getwd()

	fs := http.FileServer(http.Dir(filepath.Join(wd, "/ui/dist")))
	r.Handle("/ui/*", http.StripPrefix("/ui", fs))

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/ui/", http.StatusMovedPermanently)
	})

	return r
}

func (s *Server) ListenAndServe(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(port)), s.router())
}
