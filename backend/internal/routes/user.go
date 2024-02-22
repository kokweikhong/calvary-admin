package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/kokweikhong/calvary-admin/backend/internal/handlers"
)

type UserRoutes interface {
	RegisterRoutes(r *chi.Mux)
}

type userRoutes struct {
	userHandler handlers.UserHandler
}

func NewUserRoutes() UserRoutes {
	return &userRoutes{
		userHandler: handlers.NewUserHandler(),
	}
}

func (u *userRoutes) RegisterRoutes(r *chi.Mux) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", u.userHandler.GetUsers)
		r.Get("/{id}", u.userHandler.GetUser)
		r.Post("/", u.userHandler.CreateUser)
		r.Put("/{id}", u.userHandler.UpdateUser)
		r.Delete("/{id}", u.userHandler.DeleteUser)
	})
}
