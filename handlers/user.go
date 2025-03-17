package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserStore struct {
	sync.RWMutex
	users    map[string]User
	nextID   int
}

func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[string]User),
		nextID: 1,
	}
}

type UserHandler struct {
	store *UserStore
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		store: NewUserStore(),
	}
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	h.store.RLock()
	users := make([]User, 0, len(h.store.users))
	for _, user := range h.store.users {
		users = append(users, user)
	}
	h.store.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.store.Lock()
	user.ID = fmt.Sprintf("%d", h.store.nextID)
	h.store.users[user.ID] = user
	h.store.nextID++
	h.store.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request, id string) {
	h.store.RLock()
	user, exists := h.store.users[id]
	h.store.RUnlock()

	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
