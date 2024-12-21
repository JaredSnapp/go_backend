package models

import (
	"time"

	"github.com/google/uuid"
)

// TODO: Add User field

// No Problems:
// type GoalMetaData struct {
// 	Id    uuid.UUID `json:"id"`
// 	Name  string    `json:"name"`
// 	Type  GoalType  `json:"type"`
// 	Units string    `json:"units"`
// 	Goal  int       `json:"goal"`
// 	// CreatedAt time.Time `json:"created_at,omitempty"`
// 	// UpdatedAt time.Time `json:"updated_at,omitempty"`
// }

type GoalMetaData struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name,omitempty" gorm:"omitempty"`
	Type      GoalType  `json:"type,omitempty" gorm:"omitempty"`
	Units     string    `json:"units,omitempty" gorm:"omitempty"`
	Goal      int       `json:"goal,omitempty" gorm:"omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// TODO: Put doesn't work with generic. Not sure how to handle this. Pass in type?
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

func (g *GoalMetaData) SetID(id uuid.UUID) {
	g.Id = id
}

type BooleanGoal struct {
	Name      string    `json:"name"`
	Value     bool      `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

type CountGoal struct {
	Name      string    `json:"name"`
	Value     int       `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

type GoalType string

func (gt *GoalType) IsValid() bool {
	switch *gt {
	case "bool", "count":
		return true
	default:
		return false
	}
}
