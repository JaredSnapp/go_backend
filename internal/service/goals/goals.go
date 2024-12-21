package goals

import (
	"fmt"

	"github.com/JaredSnapp/go_backend/internal/models"
	"github.com/google/uuid"
)

// type GoalsService interface {
// 	Get() (*[]*models.GoalMetaData, error)
// 	Add(goal *models.GoalMetaData) (*models.GoalMetaData, error)
// 	Update(goal *models.GoalMetaData) (*models.GoalMetaData, error)
// 	Delete(id uuid.UUID) error
// }

type Service struct {
	repo Repo[models.GoalMetaData]
}

type Repo[T any] interface {
	Get() (*[]T, error)
	Create(Data *T) error
	// GetPersonById(personID string) (T, error)
	Update(Data *T) (*T, error)
	Delete(id uuid.UUID) error
}

func NewService(repo Repo[models.GoalMetaData]) *Service {
	return &Service{
		repo: repo,
	}
}

func (ps Service) Get() (*[]models.GoalMetaData, error) {
	data, err := ps.repo.Get()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (ps Service) Add(data models.GoalMetaData) (*models.GoalMetaData, error) {
	data.Id = uuid.New()
	fmt.Println("Create new goal")
	// fmt.Println(data)
	fmt.Printf("%+v\n", data)
	err := ps.repo.Create(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (ps Service) Update(id uuid.UUID, data models.GoalMetaData) (*models.GoalMetaData, error) {
	fmt.Println("Update GoalMetaData")
	data.Id = id

	fmt.Printf("%+v\n", data)

	updateData, err := ps.repo.Update(&data)
	if err != nil {
		return nil, err
	}

	return updateData, nil
}

func (ps Service) Delete(id uuid.UUID) error {
	fmt.Println("Delete person")
	fmt.Println(id)

	err := ps.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
