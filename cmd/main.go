package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	// Создаем HTTP сервер
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	// Запускаем сервер в горутине
	go func() {
		log.Printf("Server starting on port %d", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Ожидаем сигнал для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
