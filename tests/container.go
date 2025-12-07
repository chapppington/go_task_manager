package tests

import (
	"crud/config"
	application_tasks "crud/internal/application/tasks/usecases"
	application_users "crud/internal/application/users/usecases"
	"crud/internal/domain/tasks"
	"crud/internal/domain/users"
	"crud/internal/infrastructure/database/repositories/dummy"

	"go.uber.org/dig"
)

// NewTestContainer создает новый тестовый контейнер зависимостей для каждого теста
func NewTestContainer() *dig.Container {
	c := dig.New()
	initTestContainer(c)
	return c
}

// initTestContainer регистрирует все зависимости в тестовом контейнере
func initTestContainer(c *dig.Container) {
	// Регистрируем конфиг
	c.Provide(config.NewConfig)

	// Регистрируем in-memory репозитории
	c.Provide(dummy.NewTasksRepository, dig.As(new(tasks.BaseTasksRepository)))
	c.Provide(dummy.NewUsersRepository, dig.As(new(users.BaseUsersRepository)))

	// Регистрируем use cases
	c.Provide(application_tasks.NewCreateTaskUseCase)
	c.Provide(application_tasks.NewGetTaskByIDUseCase)
	c.Provide(application_tasks.NewListTasksUseCase)
	c.Provide(application_tasks.NewUpdateTaskUseCase)
	c.Provide(application_tasks.NewDeleteTaskUseCase)
	c.Provide(application_users.NewCreateUserUseCase)
	c.Provide(application_users.NewGetUserByIDUseCase)
	c.Provide(application_users.NewGetUserByEmailUseCase)
	c.Provide(application_users.NewListUsersUseCase)
	c.Provide(application_users.NewUpdateUserUseCase)
	c.Provide(application_users.NewDeleteUserUseCase)
}

// ResolveFromContainer получает зависимость из тестового контейнера по типу
func ResolveFromContainer[T any](container *dig.Container) (T, error) {
	var result T
	err := container.Invoke(func(dep T) {
		result = dep
	})
	return result, err
}
