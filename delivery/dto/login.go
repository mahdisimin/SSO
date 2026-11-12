package dto

import "strings"

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

func (l LoginRequest) Validate() error {
	if strings.TrimSpace(l.Username) == "" {
		return ValidationError{Field: "username", Message: "is required"}
	}
	if l.Password == "" {
		return ValidationError{Field: "password", Message: "is required"}
	}

	if len(l.Username) > 64 {
		return ValidationError{Field: "username", Message: "must be less than 64 characters"}
	}

	if len(l.Password) > 72 {
		return ValidationError{Field: "password", Message: "must be between 8 and 72 characters"}
	}

	return nil
}
