package routes

import (
	"net/http"

	"github.com/kokweikhong/calvary-admin/backend/internal/handlers"
)

type userRoutes struct {
	userHandler handlers.UserHandler
}

func NewUserRoutes() *userRoutes {
	return &userRoutes{
		userHandler: handlers.NewUserHandler(),
	}
}

func (u *userRoutes) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /users", u.userHandler.GetUsers)
	mux.HandleFunc("GET /users/{id}", u.userHandler.GetUser)
	mux.HandleFunc("POST /users", u.userHandler.CreateUser)
	mux.HandleFunc("PUT /users/{id}", u.userHandler.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", u.userHandler.DeleteUser)
}


