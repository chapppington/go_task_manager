package v1

import (
	"crud/internal/presentation/api/v1/tasks"
	"crud/internal/presentation/api/v1/users"

	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"
)

// SetupRoutes настраивает маршруты API v1
func SetupRoutes(r chi.Router, container *dig.Container) error {
	// Настраиваем маршруты для пользователей
	if err := users.SetupRoutes(r, container); err != nil {
		return err
	}

	// Настраиваем маршруты для задач
	if err := tasks.SetupRoutes(r, container); err != nil {
		return err
	}

	return nil
}
