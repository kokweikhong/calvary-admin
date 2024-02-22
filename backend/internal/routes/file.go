package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/kokweikhong/calvary-admin/backend/internal/handlers"
)

type FileRoutes interface {
	RegisterRoutes(r *chi.Mux)
}

type fileRoutes struct {
	fileHandler handlers.FileHandler
}

func NewFileRoutes() *fileRoutes {
	return &fileRoutes{
		fileHandler: handlers.NewFileHandler(),
	}
}

func (u *fileRoutes) RegisterRoutes(r *chi.Mux) {
	r.Route("/files", func(r chi.Router) {
		r.Post("/upload", u.fileHandler.UploadFile)
	})
}
