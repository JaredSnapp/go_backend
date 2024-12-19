package persons

import (
	"fmt"

	"github.com/JaredSnapp/go_backend/internal/models"
	"github.com/google/uuid"
)

type Service struct {
	Peeps *[]models.Person
}

type PersonsRepo interface {
	GetPersons() (*[]models.Person, error)
	AddPerson(newPerson *models.Person) (*models.Person, error)
	UpdatePerson(newPerson *models.Person) (*models.Person, error)
	DeletePerson(id string) error
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

	newPerson.Id = uuid.New().String()
	*ps.Peeps = append(*ps.Peeps, *newPerson)

	return newPerson, nil
}

// TODO: Update the person when we move to DB
func (ps Service) UpdatePerson(person *models.Person) (*models.Person, error) {
	fmt.Println("Update person")
	fmt.Println(person.Id)

	return person, nil
}

// TODO: delete when we move to DB
func (ps Service) DeletePerson(id string) error {
	fmt.Println("Delete person")
	fmt.Println(id)

	return nil
}
