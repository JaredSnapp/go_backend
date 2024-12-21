package HTTPHandler

import (
	"context"
	"fmt"

	"github.com/JaredSnapp/go_backend/internal/models"
	"github.com/google/uuid"
	"github.com/swaggest/usecase"
)

// Potential improvements
// Feature: Custom validator function, can pass in custom validator

type Service[T any] interface {
	Get() (*[]T, error)
	Add(goal T) (*T, error)
	Update(id uuid.UUID, goal T) (*T, error)
	Delete(id uuid.UUID) error
}

type PutInputType[T any] interface {
	GetData() T
	GetId() uuid.UUID
}

type Crud struct {
	Get    func() usecase.Interactor
	Put    func() usecase.Interactor
	Post   func() usecase.Interactor
	Delete func() usecase.Interactor
}

func NewCrud[T any, V PutInputType[T]](name string, service Service[T]) Crud {
	CRUD := Crud{}

	CRUD.Get = func() usecase.Interactor {
		f := func(ctx context.Context, input models.EmptyBody, output *[]T) error {
			list, err := service.Get()
			if err != nil {
				fmt.Println("Error", err)
			}

			*output = *list

			return nil
		}

		interactor := usecase.NewInteractor(f)
		interactor.SetDescription(name + " List")
		interactor.SetTitle(name)
		return interactor
	}

	CRUD.Post = func() usecase.Interactor {
		f := func(ctx context.Context, input T, output *T) error {
			data, err := service.Add(input)
			if err != nil {
				fmt.Println("Error", err)
			}

			*output = *data

			return err
		}

		interactor := usecase.NewInteractor(f)
		interactor.SetDescription("Create " + name)
		interactor.SetTitle("Create " + name)
		return interactor
	}

	CRUD.Put = func() usecase.Interactor {
		f := func(ctx context.Context, input V, output *T) error {
			fmt.Println("call service.Update service")
			fmt.Printf("%+v\n", input)

			inputData := input.GetData()
			inputId := input.GetId()

			data, err := service.Update(inputId, inputData)

			*output = *data

			return err
		}

		interactor := usecase.NewInteractor(f)
		interactor.SetDescription("Update " + name)
		interactor.SetTitle("Update " + name)
		return interactor
	}

	CRUD.Delete = func() usecase.Interactor {
		f := func(ctx context.Context, input *models.DeleteInput, output *models.EmptyBody) error {
			err := service.Delete(input.Id)

			return err
		}

		interactor := usecase.NewInteractor(f)
		interactor.SetDescription("Delete " + name)
		interactor.SetTitle("Delete " + name)
		return interactor
	}

	return CRUD
}
