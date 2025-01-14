package service

import "github.com/google/uuid"

type DbGenericInterface[T any] interface {
	Get() (*[]T, error)
	GetWithRelation(name string) (*[]T, error)
	Create(Data *T) error
	// GetPersonById(personID string) (T, error)
	Update(Data *T) (*T, error)
	Delete(id uuid.UUID) error
	// DeleteWithRelation(id uuid.UUID, name string) error
	DeleteWithRelation(id uuid.UUID, emptyType any, idName string) error
}
