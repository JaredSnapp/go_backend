package service

import "github.com/google/uuid"

type DbGenericInterface[T any] interface {
	Get() (*[]T, error)
	Create(Data *T) error
	// GetPersonById(personID string) (T, error)
	Update(Data *T) (*T, error)
	Delete(id uuid.UUID) error
}
