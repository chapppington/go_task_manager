package domain

import (
	"testing"

	"crud/internal/domain/users"
	vo "crud/internal/domain/users/value_objects"
)

func TestUserEntity_Creation(t *testing.T) {
	email, err := vo.NewEmailValueObject("test@example.com")
	if err != nil {
		t.Fatalf("Failed to create email: %v", err)
	}

	name, err := vo.NewUserNameValueObject("Test User")
	if err != nil {
		t.Fatalf("Failed to create name: %v", err)
	}

	user := users.NewUser(email, name)

	if user.Email.Value() != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got '%s'", user.Email.Value())
	}

	if user.Name.Value() != "Test User" {
		t.Errorf("Expected name 'Test User', got '%s'", user.Name.Value())
	}

	if user.OID == (users.User{}.OID) {
		t.Error("Expected OID to be set")
	}

	if user.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set")
	}

	if user.UpdatedAt.IsZero() {
		t.Error("Expected UpdatedAt to be set")
	}
}

func TestUserEntity_Equality(t *testing.T) {
	email1, _ := vo.NewEmailValueObject("test@example.com")
	email2, _ := vo.NewEmailValueObject("test@example.com")
	email3, _ := vo.NewEmailValueObject("other@example.com")

	name, _ := vo.NewUserNameValueObject("Test User")

	// Создаем пользователей с одинаковым OID
	user1 := users.NewUser(email1, name)
	user2 := users.NewUser(email2, name)
	user2.OID = user1.OID // Устанавливаем одинаковый OID

	// Создаем пользователя с другим OID
	user3 := users.NewUser(email3, name)

	// Пользователи с одинаковым OID должны быть равны
	if !user1.Equals(user2) {
		t.Error("Expected user1 and user2 to be equal (same OID)")
	}

	// Пользователи с разным OID не должны быть равны
	if user1.Equals(user3) {
		t.Error("Expected user1 and user3 to be different (different OID)")
	}
}

func TestUserEntity_Hash(t *testing.T) {
	email1, _ := vo.NewEmailValueObject("test@example.com")
	email2, _ := vo.NewEmailValueObject("test@example.com")
	name, _ := vo.NewUserNameValueObject("Test User")

	user1 := users.NewUser(email1, name)
	user2 := users.NewUser(email2, name)
	user2.OID = user1.OID // Устанавливаем одинаковый OID

	// Хеши должны быть одинаковыми для пользователей с одинаковым OID
	hash1 := user1.Hash()
	hash2 := user2.Hash()

	if hash1 != hash2 {
		t.Errorf("Expected hash1 (%s) == hash2 (%s) for users with same OID", hash1, hash2)
	}

	// Проверяем, что хеш не пустой
	if hash1 == "" {
		t.Error("Expected hash to be non-empty")
	}
}

func TestUserEntity_EmailValueObject(t *testing.T) {
	// Тест валидного email
	email, err := vo.NewEmailValueObject("test@example.com")
	if err != nil {
		t.Fatalf("Expected no error for valid email, got: %v", err)
	}
	if email.Value() != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got '%s'", email.Value())
	}

	// Тест пустого email
	_, err = vo.NewEmailValueObject("")
	if err == nil {
		t.Error("Expected error for empty email")
	}

	// Тест невалидного email
	_, err = vo.NewEmailValueObject("invalid-email")
	if err == nil {
		t.Error("Expected error for invalid email")
	}

	// Тест сравнения email
	email1, _ := vo.NewEmailValueObject("test@example.com")
	email2, _ := vo.NewEmailValueObject("test@example.com")
	email3, _ := vo.NewEmailValueObject("other@example.com")

	if !email1.Equals(email2) {
		t.Error("Expected email1 and email2 to be equal")
	}

	if email1.Equals(email3) {
		t.Error("Expected email1 and email3 to be different")
	}
}

func TestUserEntity_UserNameValueObject(t *testing.T) {
	// Тест валидного имени
	name, err := vo.NewUserNameValueObject("Test User")
	if err != nil {
		t.Fatalf("Expected no error for valid name, got: %v", err)
	}
	if name.Value() != "Test User" {
		t.Errorf("Expected name 'Test User', got '%s'", name.Value())
	}

	// Тест пустого имени
	_, err = vo.NewUserNameValueObject("")
	if err == nil {
		t.Error("Expected error for empty name")
	}

	// Тест слишком короткого имени
	_, err = vo.NewUserNameValueObject("A")
	if err == nil {
		t.Error("Expected error for name shorter than 2 characters")
	}

	// Тест слишком длинного имени
	longName := make([]byte, 101)
	for i := range longName {
		longName[i] = 'A'
	}
	_, err = vo.NewUserNameValueObject(string(longName))
	if err == nil {
		t.Error("Expected error for name longer than 100 characters")
	}

	// Тест сравнения имен
	name1, _ := vo.NewUserNameValueObject("Test User")
	name2, _ := vo.NewUserNameValueObject("Test User")
	name3, _ := vo.NewUserNameValueObject("Other User")

	if !name1.Equals(name2) {
		t.Error("Expected name1 and name2 to be equal")
	}

	if name1.Equals(name3) {
		t.Error("Expected name1 and name3 to be different")
	}
}
