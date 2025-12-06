package presentation

import (
	v1 "crud/internal/presentation/api/v1"
	"crud/tests"

	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"
)

// NewTestRouter создает новый тестовый chi роутер с настроенными маршрутами
func NewTestRouter(container *dig.Container) chi.Router {
	r := chi.NewRouter()

	// Настраиваем API v1 с тестовым контейнером
	r.Route("/api/v1", func(r chi.Router) {
		if err := v1.SetupRoutes(r, container); err != nil {
			panic(err)
		}
	})

	return r
}

// NewTestRouterWithContainer создает новый тестовый роутер с новым тестовым контейнером
func NewTestRouterWithContainer() chi.Router {
	container := tests.NewTestContainer()
	return NewTestRouter(container)
}
