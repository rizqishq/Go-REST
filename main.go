// @title Go REST User API
// @version 1.0
// @description This is a sample REST API for managing users using Go.
// @contact.name Your Name
// @contact.email your.email@example.com
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/rizqishq/Go-REST/config"
	"github.com/rizqishq/Go-REST/controllers"
	_ "github.com/rizqishq/Go-REST/docs"
	"github.com/rizqishq/Go-REST/middleware"
	"github.com/rizqishq/Go-REST/repositories"
	"github.com/rizqishq/Go-REST/services"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	cfg := config.LoadConfig()

	router := mux.NewRouter()

	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.RecoveryMiddleware)

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	userRepo := repositories.NewInMemoryUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	registerRoutes(apiRouter, userController)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	// Start server in a goroutine
	go func() {
		fmt.Printf("Starting server on :%s\n", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :%s: %v\n", cfg.Server.Port, err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Server is shutting down...")

	// Create a deadline to wait for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server exited properly")
}

// registerRoutes sets up all API routes
// @Summary Health Check
// @Description Returns API health status
// @Tags system
// @Produce plain
// @Success 200 {string} string "API is healthy"
// @Router /health [get]
func registerRoutes(router *mux.Router, userController *controllers.UserController) {
	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API is healthy"))
	}).Methods("GET")

	userController.RegisterRoutes(router)
}
