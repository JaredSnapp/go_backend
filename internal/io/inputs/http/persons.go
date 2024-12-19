package HTTPHandler

import (
	"context"
	"fmt"

	"github.com/JaredSnapp/go_backend/internal/models"
	"github.com/swaggest/usecase"
)

type getPersonsInput struct {
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
