package main

import (
	_ "ProjectManagementService/docs"
	"ProjectManagementService/internal/handlers"
	"ProjectManagementService/internal/models"
	"context"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Project Management Service API
// version 1.0
// @description This is project management service API
// @host projectmanagementservice.onrender.com
// @BasePath /
func main() {

	db, err := initializeDB()
	if err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Could not close the database connection: ", err)
		}
	}(db)
	userHandler := handlers.NewUserHandler(models.NewUserModel(db))
	taskHandler := handlers.NewTaskHandler(models.NewTaskModel(db))
	projectHandler := handlers.NewProjectHandler(models.NewProjectModel(db))

	router := mux.NewRouter()

	SetupRouter(router, userHandler, taskHandler, projectHandler)

	port := "8080"
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// graceful shutdown
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

		<-signals

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Graceful shutdown failed: %v\n", err)
		}
	}()

	log.Printf("Server is starting on port %s\n", port)
	// start server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server startup failed: %v\n", err)
	}

	log.Println("Server gracefully stopped")

}
