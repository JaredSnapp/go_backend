package persons

import (
	"fmt"

	"github.com/JaredSnapp/go_backend/internal/models"
)

type Service struct {
	Peeps *[]models.Person
}

type PersonsRepo interface {
	GetPersons() (*[]models.Person, error)
	AddPerson(newPerson *models.Person) (*models.Person, error)
}

func NewService() *Service {
	peeps := make([]models.Person, 0)
	return &Service{
		Peeps: &peeps,
	}
}

func (ps Service) GetPersons() (*[]models.Person, error) {
	return ps.Peeps, nil
}

func (ps Service) AddPerson(newPerson *models.Person) (*models.Person, error) {
	fmt.Println("persons.AddPerson")
	fmt.Printf("%+v\n", newPerson)

	*ps.Peeps = append(*ps.Peeps, *newPerson)

	return newPerson, nil
}
