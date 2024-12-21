package models

import (
	"time"

	"github.com/google/uuid"
)

// TODO: Add User field
type GoalMetaData struct {
	Id          uuid.UUID `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name,omitempty" gorm:"omitempty"`
	Type        string    `json:"type,omitempty" gorm:"unique,omitempty"`
	BooleanGoal bool      `json:"boolean_goal,omitempty" gorm:"omitempty"`
	Units       string    `json:"units,omitempty" gorm:"omitempty"`
	Goal        int       `json:"goal,omitempty" gorm:"omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type GoalPutInput struct {
	Id uuid.UUID `path:"id" required:"true" description:"document id of shift to retrieve"`
	GoalMetaData
}

func (in GoalPutInput) GetData() GoalMetaData {
	return in.GoalMetaData
}

func (in GoalPutInput) GetId() uuid.UUID {
	return in.Id
}

// func (g *GoalMetaData) SetID(id uuid.UUID) {
// 	g.Id = id
// }
