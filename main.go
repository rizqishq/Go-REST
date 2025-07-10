package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rizqishq/Go-REST/config"
	"github.com/rizqishq/Go-REST/controllers"
	"github.com/rizqishq/Go-REST/middleware"
	"github.com/rizqishq/Go-REST/repositories"
	"github.com/rizqishq/Go-REST/services"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	router := mux.NewRouter().StrictSlash(true)

	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.RecoveryMiddleware)

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	userRepo := repositories.NewInMemoryUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	userController.RegisterRoutes(apiRouter)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API is healthy"))
	}).Methods("GET")

	addr := ":" + cfg.Server.Port
	fmt.Printf("Server running on %s\n", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
