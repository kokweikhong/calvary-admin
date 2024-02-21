package routes

import (
	"net/http"

	"github.com/kokweikhong/calvary-admin/backend/internal/handlers"
)

type fileRoutes struct {
	fileHandler handlers.FileHandler
}

func NewFileRoutes() *fileRoutes {
	return &fileRoutes{
		fileHandler: handlers.NewFileHandler(),
	}
}

func (u *fileRoutes) RegisterRoutes(mux *http.ServeMux) {
	// mux.HandleFunc("GET /users", u.userHandler.GetUsers)
	// mux.HandleFunc("GET /users/{id}", u.userHandler.GetUser)
	mux.HandleFunc("POST /files/upload", u.fileHandler.UploadFile)
	// mux.HandleFunc("PUT /users/{id}", u.userHandler.UpdateUser)
	// mux.HandleFunc("DELETE /users/{id}", u.userHandler.DeleteUser)
}
