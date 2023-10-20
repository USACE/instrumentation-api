package model

import (
	"fmt"
)

const (
	ErrNotFound = "data not found"
	ErrQuery    = "problem executing query"
)

type ErrorService interface {
	Error() string
	Unwrap() error
}

type errorService struct {
	message string
	err     error
}

type NotFoundError struct {
	errorService
}

type QueryError struct {
	errorService
}

func NewNotFoundError(wrappedErr error) error {
	err := errorService{message: ErrNotFound, err: wrappedErr}
	return &NotFoundError{err}
}
func NewQueryError(wrappedErr error) error {
	err := errorService{message: ErrQuery, err: wrappedErr}
	return &QueryError{err}
}

func (e errorService) Error() string {
	return fmt.Sprintf("service error: %s", e.message)
}

func (e errorService) Unwrap() error { return e.err }
