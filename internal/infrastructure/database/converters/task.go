package converters

import (
	"crud/internal/domain/tasks"
	"crud/internal/domain/tasks/value_objects"
	"crud/internal/infrastructure/database/models"
)

// TaskModelToEntity конвертирует GORM модель в domain entity
func TaskModelToEntity(model *models.Task) (*tasks.Task, error) {
	if model == nil {
		return nil, nil
	}

	title, err := value_objects.NewTaskTitleValueObject(model.Title)
	if err != nil {
		return nil, err
	}

	status, err := value_objects.NewTaskStatusValueObject(model.Status)
	if err != nil {
		return nil, err
	}

	return &tasks.Task{
		ID:          model.ID,
		UserID:      model.UserID,
		Title:       title,
		Description: model.Description,
		Status:      status,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}, nil
}

// TaskEntityToModel конвертирует domain entity в GORM модель
func TaskEntityToModel(task *tasks.Task) *models.Task {
	if task == nil {
		return nil
	}

	return &models.Task{
		ID:          task.ID,
		UserID:      task.UserID,
		Title:       task.Title.Value(),
		Description: task.Description,
		Status:      task.Status.Value(),
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}
