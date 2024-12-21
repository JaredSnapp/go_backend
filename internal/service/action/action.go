package action

import (
	"fmt"

	"github.com/JaredSnapp/go_backend/internal/models"
	"github.com/JaredSnapp/go_backend/internal/service"
	"github.com/google/uuid"
)

type Service struct {
	repo service.DbGenericInterface[models.Action]
}

func NewService(repo service.DbGenericInterface[models.Action]) *Service {
	return &Service{
		repo: repo,
	}
}

func (ps Service) Get() (*[]models.Action, error) {
	data, err := ps.repo.Get()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (ps Service) Add(data models.Action) (*models.Action, error) {
	data.Id = uuid.New()
	fmt.Println("Create new action")
	// fmt.Println(data)
	fmt.Printf("%+v\n", data)
	err := ps.repo.Create(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (ps Service) Update(id uuid.UUID, data models.Action) (*models.Action, error) {
	fmt.Println("Update action")
	data.Id = id

	fmt.Printf("%+v\n", data)

	updateData, err := ps.repo.Update(&data)
	if err != nil {
		return nil, err
	}

	return updateData, nil
}

func (ps Service) Delete(id uuid.UUID) error {
	fmt.Println("Delete action")
	fmt.Println(id)

	err := ps.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
