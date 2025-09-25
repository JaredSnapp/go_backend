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

type RecruiterService interface {
	Get() (*[]models.Recruiter, error)
	Add(data models.Recruiter) (*models.Recruiter, error)
	Update(id uuid.UUID, data models.Recruiter) (*models.Recruiter, error)
	Delete(id uuid.UUID) error
}

type RichTextService interface {
	Get() (*[]models.RichText, error)
	Add(data models.RichText) (*models.RichText, error)
	Update(id uuid.UUID, data models.RichText) (*models.RichText, error)
	Delete(id uuid.UUID) error
}
