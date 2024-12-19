package HTTPHandler

import (
	"github.com/JaredSnapp/go_backend/internal/models"
)

type PersonsService interface {
	GetPersons() (*[]models.Person, error)
	AddPerson(newPerson *models.Person) (*models.Person, error)
}
