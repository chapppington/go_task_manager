package application

import (
	"sync"

	"crud/config"
	tasks_usecases "crud/internal/application/tasks/usecases"
	users_usecases "crud/internal/application/users/usecases"
	tasks_domain "crud/internal/domain/tasks"
	users_domain "crud/internal/domain/users"
	"crud/internal/infrastructure/database/gateways"
	"crud/internal/infrastructure/database/repositories"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

var (
	containerOnce sync.Once
	container     *dig.Container
)

// InitContainer инициализирует контейнер зависимостей (singleton)
func InitContainer() *dig.Container {
	containerOnce.Do(func() {
		container = dig.New()
		initContainer(container)
	})
	return container
}

// initContainer регистрирует все зависимости в контейнере
func initContainer(c *dig.Container) {
	// Регистрируем конфиг
	c.Provide(config.NewConfig)

	// Регистрируем gateway для подключения к БД
	c.Provide(gateways.NewPostgresGateway)

	// Провайдер для *gorm.DB из gateway
	c.Provide(func(gw *gateways.PostgresGateway) *gorm.DB {
		return gw.DB()
	})

	// Регистрируем репозитории
	c.Provide(repositories.NewUsersRepository, dig.As(new(users_domain.BaseUsersRepository)))
	c.Provide(repositories.NewTasksRepository, dig.As(new(tasks_domain.BaseTasksRepository)))

	// Регистрируем use cases для пользователей
	c.Provide(users_usecases.NewCreateUserUseCase)
	c.Provide(users_usecases.NewGetUserByIDUseCase)
	c.Provide(users_usecases.NewGetUserByEmailUseCase)
	c.Provide(users_usecases.NewListUsersUseCase)
	c.Provide(users_usecases.NewUpdateUserUseCase)
	c.Provide(users_usecases.NewDeleteUserUseCase)

	// Регистрируем use cases для задач
	c.Provide(tasks_usecases.NewCreateTaskUseCase)
	c.Provide(tasks_usecases.NewGetTaskByIDUseCase)
	c.Provide(tasks_usecases.NewListTasksUseCase)
	c.Provide(tasks_usecases.NewUpdateTaskUseCase)
	c.Provide(tasks_usecases.NewDeleteTaskUseCase)
}

// ResolveFromContainer получает зависимость из переданного контейнера по типу
func ResolveFromContainer[T any](container *dig.Container) (T, error) {
	var result T
	err := container.Invoke(func(dep T) {
		result = dep
	})
	return result, err
}
