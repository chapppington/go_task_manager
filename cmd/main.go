package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"crud/config"
	"crud/internal/application"
	v1 "crud/internal/presentation/api/v1"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Инициализируем контейнер зависимостей
	container := application.InitContainer()

	// Создаем chi роутер
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Настраиваем API v1
	r.Route("/api/v1", func(r chi.Router) {
		if err := v1.SetupRoutes(r, container); err != nil {
			log.Fatalf("Failed to setup routes: %v", err)
		}
	})

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Получаем конфиг для порта
	var cfg *config.Config
	if err := container.Invoke(func(c *config.Config) {
		cfg = c
	}); err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	port := cfg.APIPort
	if port == 0 {
		port = 8000
	}

	// Запускаем сервер
	log.Printf("Server starting on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
