package postgres

import (
	"fmt"

	"github.com/google/uuid"
)

type Repo[T any] struct {
	Pg *Service
}

func (r Repo[T]) Get() (*[]T, error) {
	var data []T
	result := r.Pg.db.Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (r Repo[T]) Create(data *T) error {
	fmt.Println("Create DB Level")
	result := r.Pg.db.Create(data)
	if result.Error != nil {
		fmt.Println("db error")
		fmt.Println(result.Error)
		return result.Error
	}
	fmt.Println("Create DB Level Successful")
	return nil
}

func (r Repo[T]) Update(data *T) (*T, error) {
	result := r.Pg.db.Updates(data)
	if result.Error != nil {
		return new(T), result.Error
	}
	return data, nil
}

func (r Repo[T]) Delete(id uuid.UUID) error {
	result := r.Pg.db.Delete(new(T), id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
