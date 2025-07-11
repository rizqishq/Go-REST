package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rizqishq/Go-REST/models"
	"github.com/rizqishq/Go-REST/services"
)

// UserController handles user-related endpoints
type UserController struct {
	userService *services.UserService
}

// Create new UserController
func NewUserController(s *services.UserService) *UserController {
	return &UserController{userService: s}
}

// RegisterRoutes hooks controller into router
func (c *UserController) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users", c.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", c.GetUserByID).Methods("GET")
	r.HandleFunc("/users", c.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", c.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", c.DeleteUser).Methods("DELETE")
}

// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} models.UserResponse
// @Router /users [get]
func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.userService.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

// @Summary Get a user by ID
// @Description Get user details by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.UserResponse
// @Failure 404 {object} middleware.ErrorResponse
// @Router /users/{id} [get]
func (c *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user, err := c.userService.GetUserByID(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User Data"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} middleware.ErrorResponse
// @Router /users [post]
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := c.userService.CreateUser(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}

// @Summary Update an existing user
// @Description Update user data by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UpdateUserRequest true "Updated data"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} middleware.ErrorResponse
// @Router /users/{id} [put]
func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	var req models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	user, err := c.userService.UpdateUser(r.Context(), uint(id), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 404 {object} middleware.ErrorResponse
// @Router /users/{id} [delete]
func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	if err := c.userService.DeleteUser(r.Context(), uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
