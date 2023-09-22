package model

import (
	"fmt"
)

const (
	ErrNotFound = "data not found"
	ErrQuery    = "problem executing query"
)

type ErrorStore interface {
	Error() string
	Unwrap() error
}

type errorStore struct {
	message string
	err     error
}

type NotFoundError struct {
	errorStore
}

type QueryError struct {
	errorStore
}

func NewNotFoundError(wrappedErr error) error {
	err := errorStore{message: ErrNotFound, err: wrappedErr}
	return &NotFoundError{err}
}
func NewQueryError(wrappedErr error) error {
	err := errorStore{message: ErrQuery, err: wrappedErr}
	return &QueryError{err}
}

func (e errorStore) Error() string {
	return fmt.Sprintf("store error: %s", e.message)
}

func (e errorStore) Unwrap() error { return e.err }
