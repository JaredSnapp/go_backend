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

type GoalsService interface {
	Get() (*[]models.GoalMetaData, error)
	Add(goal models.GoalMetaData) (*models.GoalMetaData, error)
	Update(id uuid.UUID, goal models.GoalMetaData) (*models.GoalMetaData, error)
	Delete(id uuid.UUID) error
}

type ActionService interface {
	Get() (*[]models.Action, error)
	Add(data models.Action) (*models.Action, error)
	Update(id uuid.UUID, data models.Action) (*models.Action, error)
	Delete(id uuid.UUID) error
}
