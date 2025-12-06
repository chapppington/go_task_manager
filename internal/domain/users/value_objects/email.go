package value_objects

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
)

// EmailValueObject представляет email адрес с валидацией
type EmailValueObject struct {
	value string
}

// NewEmailValueObject создает новый EmailValueObject с валидацией
func NewEmailValueObject(email string) (EmailValueObject, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return EmailValueObject{}, fmt.Errorf("email cannot be empty")
	}
	if !emailRegex.MatchString(email) {
		return EmailValueObject{}, fmt.Errorf("invalid email format: %s", email)
	}
	return EmailValueObject{value: strings.ToLower(email)}, nil
}

// Value возвращает строковое значение email
func (e EmailValueObject) Value() string {
	return e.value
}

// Equals проверяет равенство двух email
func (e EmailValueObject) Equals(other EmailValueObject) bool {
	return e.value == other.value
}
