package main

import (
	"BookShop/internal/db"
	"BookShop/internal/handler"
	"BookShop/internal/repository/postgres"
	"BookShop/internal/service"
	"log"
	"net/http"

	"BookShop/internal/config"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.Load()

	conn, err := db.PostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	bookRepo := postgres.InitBookRepository(conn)
	bookService := service.InitBookService(bookRepo)
	bookHandler := handler.InitBookHandler(bookService)

	mux := chi.NewRouter()
	handler.RegisterRoutes(mux, bookHandler)

	log.Printf("Server starting on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
