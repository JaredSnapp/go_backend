package postgres

import (
	"github.com/JaredSnapp/go_backend/internal/models"
	"github.com/google/uuid"
)

func (pg *Service) CreatePerson(person *models.Person) error {
	result := pg.db.Create(person)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pg *Service) GetPersons() (*[]models.Person, error) {
	var persons []models.Person
	result := pg.db.Find(&persons)
	if result.Error != nil {
		return nil, result.Error
	}
	return &persons, nil
}

func (pg *Service) GetPersonById(personID string) (*models.Person, error) {
	var person models.Person
	result := pg.db.First(&person, personID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &person, nil
}

func (pg *Service) UpdatePerson(person *models.Person) (*models.Person, error) {
	result := pg.db.Updates(person)
	if result.Error != nil {
		return nil, result.Error
	}
	return person, nil
}

func (pg *Service) DeletePerson(id uuid.UUID) error {
	result := pg.db.Delete(models.Person{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
