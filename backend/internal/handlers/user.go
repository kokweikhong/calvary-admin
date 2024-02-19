package handlers

import (
	"net/http"
	"strconv"

	"github.com/kokweikhong/calvary-admin/backend/internal/models"
	"github.com/kokweikhong/calvary-admin/backend/internal/services"
	"github.com/kokweikhong/calvary-admin/backend/internal/utils"
)

type UserHandler interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	srv   services.UserService
	jsonH utils.JSONHandler
}

func NewUserHandler() UserHandler {
	return &userHandler{
		srv:   services.NewUserService(),
		jsonH: utils.NewJSONHandler(),
	}
}

func (u *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.srv.GetUsers()
	if err != nil {
		u.jsonH.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	u.jsonH.WriteJSON(w, http.StatusOK, users)
}

func (u *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		u.jsonH.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := u.srv.GetUser(id)
	if err != nil {
		u.jsonH.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	u.jsonH.WriteJSON(w, http.StatusOK, user)
}

func (u *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	err := u.jsonH.ReadJSON(w, r, user)
	if err != nil {
		u.jsonH.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	id, err := u.srv.CreateUser(user)
	if err != nil {
		u.jsonH.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	u.jsonH.WriteJSON(w, http.StatusCreated, id)
}

func (u *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		u.jsonH.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	user := &models.User{}

	err = u.jsonH.ReadJSON(w, r, user)
	if err != nil {
		u.jsonH.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	user.ID = id

	user, err = u.srv.UpdateUser(user)
	if err != nil {
		u.jsonH.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	u.jsonH.WriteJSON(w, http.StatusOK, user)
}

func (u *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		u.jsonH.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = u.srv.DeleteUser(id)
	if err != nil {
		u.jsonH.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	u.jsonH.WriteJSON(w, http.StatusNoContent, nil)
}
