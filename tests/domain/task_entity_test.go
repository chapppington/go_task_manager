package domain

import (
	"testing"

	"crud/internal/domain/tasks"
	vo "crud/internal/domain/tasks/value_objects"

	"github.com/google/uuid"
)

func TestTaskEntity_Creation(t *testing.T) {
	title, err := vo.NewTaskTitleValueObject("Learn gRPC")
	if err != nil {
		t.Fatalf("Failed to create title: %v", err)
	}

	status, err := vo.NewTaskStatusValueObject("todo")
	if err != nil {
		t.Fatalf("Failed to create status: %v", err)
	}

	userID := uuid.New()
	task := tasks.NewTask(userID, title, "Study gRPC basics", status)

	if task.Title.Value() != "Learn gRPC" {
		t.Errorf("Expected title 'Learn gRPC', got '%s'", task.Title.Value())
	}

	if task.Description != "Study gRPC basics" {
		t.Errorf("Expected description 'Study gRPC basics', got '%s'", task.Description)
	}

	if task.Status.Value() != "todo" {
		t.Errorf("Expected status 'todo', got '%s'", task.Status.Value())
	}

	if task.UserID != userID {
		t.Errorf("Expected UserID %s, got %s", userID, task.UserID)
	}

	if task.ID == (tasks.Task{}.ID) {
		t.Error("Expected ID to be set")
	}

	if task.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set")
	}

	if task.UpdatedAt.IsZero() {
		t.Error("Expected UpdatedAt to be set")
	}
}

func TestTaskEntity_Equality(t *testing.T) {
	title1, _ := vo.NewTaskTitleValueObject("Learn gRPC")
	title2, _ := vo.NewTaskTitleValueObject("Learn gRPC")
	title3, _ := vo.NewTaskTitleValueObject("Master Go")

	status, _ := vo.NewTaskStatusValueObject("todo")
	userID := uuid.New()

	// Создаем задачи с одинаковым ID
	task1 := tasks.NewTask(userID, title1, "Description 1", status)
	task2 := tasks.NewTask(userID, title2, "Description 2", status)
	task2.ID = task1.ID // Устанавливаем одинаковый ID

	// Создаем задачу с другим ID
	task3 := tasks.NewTask(userID, title3, "Description 3", status)

	// Задачи с одинаковым ID должны быть равны
	if !task1.Equals(task2) {
		t.Error("Expected task1 and task2 to be equal (same ID)")
	}

	// Задачи с разным ID не должны быть равны
	if task1.Equals(task3) {
		t.Error("Expected task1 and task3 to be different (different ID)")
	}
}

func TestTaskEntity_TitleValueObject(t *testing.T) {
	// Тест валидного заголовка
	title, err := vo.NewTaskTitleValueObject("Learn gRPC")
	if err != nil {
		t.Fatalf("Expected no error for valid title, got: %v", err)
	}
	if title.Value() != "Learn gRPC" {
		t.Errorf("Expected title 'Learn gRPC', got '%s'", title.Value())
	}

	// Тест пустого заголовка
	_, err = vo.NewTaskTitleValueObject("")
	if err == nil {
		t.Error("Expected error for empty title")
	}

	// Тест слишком длинного заголовка
	longTitle := make([]byte, 201)
	for i := range longTitle {
		longTitle[i] = 'A'
	}
	_, err = vo.NewTaskTitleValueObject(string(longTitle))
	if err == nil {
		t.Error("Expected error for title longer than 200 characters")
	}

	// Тест сравнения заголовков
	title1, _ := vo.NewTaskTitleValueObject("Learn gRPC")
	title2, _ := vo.NewTaskTitleValueObject("Learn gRPC")
	title3, _ := vo.NewTaskTitleValueObject("Master Go")

	if !title1.Equals(title2) {
		t.Error("Expected title1 and title2 to be equal")
	}

	if title1.Equals(title3) {
		t.Error("Expected title1 and title3 to be different")
	}
}

func TestTaskEntity_StatusValueObject(t *testing.T) {
	// Тест валидных статусов
	validStatuses := []string{"todo", "in_progress", "done"}
	for _, statusStr := range validStatuses {
		status, err := vo.NewTaskStatusValueObject(statusStr)
		if err != nil {
			t.Fatalf("Expected no error for valid status '%s', got: %v", statusStr, err)
		}
		if status.Value() != statusStr {
			t.Errorf("Expected status '%s', got '%s'", statusStr, status.Value())
		}
		if !status.IsValid() {
			t.Errorf("Expected status '%s' to be valid", statusStr)
		}
	}

	// Тест пустого статуса
	_, err := vo.NewTaskStatusValueObject("")
	if err == nil {
		t.Error("Expected error for empty status")
	}

	// Тест невалидного статуса
	_, err = vo.NewTaskStatusValueObject("invalid_status")
	if err == nil {
		t.Error("Expected error for invalid status")
	}

	// Тест сравнения статусов
	status1, _ := vo.NewTaskStatusValueObject("todo")
	status2, _ := vo.NewTaskStatusValueObject("todo")
	status3, _ := vo.NewTaskStatusValueObject("done")

	if !status1.Equals(status2) {
		t.Error("Expected status1 and status2 to be equal")
	}

	if status1.Equals(status3) {
		t.Error("Expected status1 and status3 to be different")
	}
}
