package HTTPHandler

import (
	"context"
	"fmt"

	"github.com/JaredSnapp/go_backend/internal/models"
	"github.com/google/uuid"
	"github.com/swaggest/usecase"
)

type getPersonsInput struct {
}

type PutPersonInput struct {
	Id   uuid.UUID `path:"id" required:"true" description:"document id of shift to retrieve"`
	Data models.Person
}

func (h Handler) getPersons() usecase.Interactor {
	f := func(ctx context.Context, input getPersonsInput, output *[]models.Person) error {
		persons, err := h.PersonsService.GetPersons()
		if err != nil {
			fmt.Println("Error", err)
		}

		*output = *persons

		return nil
	}

	interactor := usecase.NewInteractor(f)
	interactor.SetDescription("Person List")
	interactor.SetTitle("Persons")
	return interactor
}

func (h Handler) postPerson() usecase.Interactor {
	f := func(ctx context.Context, input *models.Person, output *models.Person) error {

		person, err := h.PersonsService.AddPerson(input)

		*output = *person

		return err
	}

	interactor := usecase.NewInteractor(f)
	interactor.SetDescription("Create Person")
	interactor.SetTitle("Create Person")
	return interactor
}

func (h Handler) putPerson() usecase.Interactor {
	f := func(ctx context.Context, input PutPersonInput, output *models.Person) error {

		input.Data.SetID(input.Id)
		person, err := h.PersonsService.UpdatePerson(&input.Data)

		*output = *person

		return err
	}

	interactor := usecase.NewInteractor(f)
	interactor.SetDescription("Update Person")
	interactor.SetTitle("Update Person")
	return interactor
}

func (h Handler) deletePerson() usecase.Interactor {
	f := func(ctx context.Context, input *models.DeleteInput, output *models.EmptyBody) error {

		err := h.PersonsService.DeletePerson(input.Id)

		return err
	}

	interactor := usecase.NewInteractor(f)
	interactor.SetDescription("Delete Person")
	interactor.SetTitle("Delete Person")
	return interactor
}
