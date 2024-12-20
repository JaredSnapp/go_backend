package HTTPHandler

import (
	"github.com/JaredSnapp/go_backend/internal/models"
	"github.com/google/uuid"
)

type PersonsService interface {
	GetPersons() (*[]models.Person, error)
	AddPerson(newPerson *models.Person) (*models.Person, error)
	UpdatePerson(newPerson *models.Person) (*models.Person, error)
	DeletePerson(id uuid.UUID) error
}
