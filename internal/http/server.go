package http

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/milankyncl/feature-toggles/internal/http/handler"
	"github.com/milankyncl/feature-toggles/internal/storage"
	"net/http"
	"path/filepath"
	"strconv"
)

type Server struct {
	uiPath  string
	storage storage.Adapter
}

func NewServer(basePath string, storage storage.Adapter) *Server {
	uiPath := filepath.Join(basePath, "ui/dist")
	return &Server{
		uiPath,
		storage,
	}
}

func (s *Server) router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Mount("/api/admin", s.adminApiRouter())
	r.Get("/api/feature-toggles", handler.GetFeatureTogglesHandler(s.storage))
	r.Mount("/", s.uiRouter())
	return r
}

func (s *Server) adminApiRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/features", handler.GetFeaturesHandler(s.storage))
	r.Post("/features", handler.CreateFeatureHandler(s.storage))
	r.Get("/features/{id}", handler.GetFeatureHandler(s.storage))
	r.Put("/features/{id}", handler.UpdateFeatureHandler(s.storage))
	r.Delete("/features/{id}", handler.DeleteFeatureHandler(s.storage))
	r.Put("/features/{id}/toggle", handler.ToggleFeatureHandler(s.storage))
	return r
}

func (s *Server) uiRouter() *chi.Mux {
	r := chi.NewRouter()

	fs := http.FileServer(http.Dir(s.uiPath))
	r.Handle("/ui/*", http.StripPrefix("/ui", fs))

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/ui/", http.StatusMovedPermanently)
	})

	return r
}

func (s *Server) ListenAndServe(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(port)), s.router())
}
