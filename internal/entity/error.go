package entity

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Cause   error  `json:"-"`
}

func NewError(status int, message string) *AppError {
	return &AppError{Status: status, Message: message}
}

func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Cause
}

func (e *AppError) GetStatus() int {
	return e.Status
}

func (e *AppError) WithCause(err error) *AppError {
	clone := *e
	clone.Cause = err
	return &clone
}

var (
	ErrUserNotFound = NewError(http.StatusNotFound, "User not found.")
)