package control

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"studentgit.kata.academy/Zhodaran/go-kata/internal/repository"
)

type UserController struct {
	userRepo repository.UserRepository
}

func NewUserController(userRepo repository.UserRepository) *UserController {
	return &UserController{userRepo: userRepo}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user repository.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := uc.userRepo.Create(context.Background(), user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := uc.userRepo.GetByID(context.Background(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user repository.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := uc.userRepo.Update(context.Background(), user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := uc.userRepo.Delete(context.Background(), id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	limit := 10 // Установите значение по умолчанию
	offset := 0 // Установите значение по умолчанию
	users, err := uc.userRepo.List(context.Background(), limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
