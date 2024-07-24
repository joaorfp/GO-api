package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joaorfp/GO-api/internal/model"
	"github.com/joaorfp/GO-api/internal/service"
)

type UserHandler struct {
    service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
    return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
    users := h.service.GetAllUsers()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"message": "Invalid ID"})
        return
    }

    user, found := h.service.GetUserById(id)
    if found {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(user)
        return
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"message": "User not found"})
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user model.User
    _ = json.NewDecoder(r.Body).Decode(&user)
    createdUser := h.service.CreateUser(&user)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(createdUser)
}

func (h *UserHandler) UpdateUserById(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"message": "Invalid ID"})
        return
    }

    var user model.User
    _ = json.NewDecoder(r.Body).Decode(&user)
    updatedUser, found := h.service.UpdateUserById(id, &user)
    if found {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(updatedUser)
        return
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"message": "User not found"})
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"message": "Invalid ID"})
        return
    }

    deleted := h.service.DeleteUser(id)
    if deleted {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"})
        return
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"message": "User not found"})
}
