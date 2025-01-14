package goals

import (
	"fmt"

	"github.com/JaredSnapp/go_backend/internal/models"
	"github.com/JaredSnapp/go_backend/internal/service"
	"github.com/google/uuid"
)

type Repo interface {
	service.DbGenericInterface[models.Goal]
	DeleteGoal(id uuid.UUID) error
}

type Service struct {
	repo Repo
}

func NewService(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (ps Service) Get() (*[]models.Goal, error) {
	data, err := ps.repo.GetWithRelation("Actions")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (ps Service) Add(data models.Goal) (*models.Goal, error) {
	data.ID = uuid.New()
	fmt.Println("Create new goal")
	// fmt.Println(data)
	fmt.Printf("%+v\n", data)
	err := ps.repo.Create(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (ps Service) Update(id uuid.UUID, data models.Goal) (*models.Goal, error) {
	fmt.Println("Update GoalMetaData")
	data.ID = id

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
	emptyAction := models.Action{}

	err := ps.repo.DeleteWithRelation(id, &emptyAction, "goal_id")
	if err != nil {
		return err
	}

	return nil
}
