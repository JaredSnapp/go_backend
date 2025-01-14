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
	err := r.Pg.db.Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r Repo[T]) GetWithRelation(name string) (*[]T, error) {
	var temp T
	var data []T
	err := r.Pg.db.Model(&temp).Preload(name).Find(&data).Error
	if err != nil {
		return nil, err
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

func (r Repo[T]) DeleteWithRelation(id uuid.UUID, emptyType any, idName string) error {
	r.Pg.db.Model(emptyType).Where(idName+" = ?", id).Delete(emptyType)

	fmt.Println("Done deleting associated record...")

	// result := r.Pg.db.Select(name).Delete(new(T), id)
	result := r.Pg.db.Delete(new(T), id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
