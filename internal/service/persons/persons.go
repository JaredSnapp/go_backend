package persons

import (
	"fmt"

	"github.com/JaredSnapp/go_backend/internal/models"
	"github.com/google/uuid"
)

type Service struct {
	Peeps *[]models.Person
	repo  PersonRepo
}

type PersonRepo interface {
	CreatePerson(person *models.Person) error
	GetPersons() (*[]models.Person, error)
	GetPersonById(personID string) (*models.Person, error)
	UpdatePerson(person *models.Person) (*models.Person, error)
	DeletePerson(id uuid.UUID) error
}

func NewService(repo PersonRepo) *Service {
	peeps := make([]models.Person, 0)
	return &Service{
		Peeps: &peeps,
		repo:  repo,
	}
}

func (ps Service) GetPersons() (*[]models.Person, error) {
	peeps, err := ps.repo.GetPersons()
	if err != nil {
		return nil, err
	}

	return peeps, nil
}

func (ps Service) AddPerson(newPerson *models.Person) (*models.Person, error) {
	newPerson.Id = uuid.New()
	err := ps.repo.CreatePerson(newPerson)
	if err != nil {
		return nil, err
	}

	return newPerson, nil
}

func (ps Service) UpdatePerson(person *models.Person) (*models.Person, error) {
	fmt.Println("Update person")
	fmt.Println(person.Id)

	person, err := ps.repo.UpdatePerson(person)
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (ps Service) DeletePerson(id uuid.UUID) error {
	fmt.Println("Delete person")
	fmt.Println(id)

	err := ps.repo.DeletePerson(id)
	if err != nil {
		return err
	}

	return nil
}
