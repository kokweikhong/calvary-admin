package routes

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Routes interface {
	RegisterRoutes() *chi.Mux
	ListenAndServe(router *chi.Mux, port string)
}

type routes struct {
	fileRoutes FileRoutes
	userRoutes UserRoutes
}

func NewRoutes() *routes {
	return &routes{
		fileRoutes: NewFileRoutes(),
		userRoutes: NewUserRoutes(),
	}
}

func (r *routes) RegisterRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.userRoutes.RegisterRoutes(router)
	r.fileRoutes.RegisterRoutes(router)

	return router
}

func (r *routes) ListenAndServe(router *chi.Mux, port string) {
	slog.Info("Server is running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
