package models

import (
	"time"

	"github.com/google/uuid"
)

type Goal struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name,omitempty" gorm:"omitempty"`
	BooleanGoal bool      `json:"boolean_goal,omitempty" gorm:"omitempty"`
	Units       string    `json:"units,omitempty" gorm:"omitempty"`
	Goal        int       `json:"goal,omitempty" gorm:"omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Actions     []Action  `gorm:"foreignKey:GoalID;constraint:OnDelete:CASCADE;"`
}

type GoalPutInput struct {
	Id uuid.UUID `path:"id" required:"true" description:"document id of shift to retrieve"`
	Goal
}

func (in GoalPutInput) GetData() Goal {
	return in.Goal
}

func (in GoalPutInput) GetId() uuid.UUID {
	return in.Id
}

// func (g *Goal) BeforeDelete(tx *gorm.DB) (err error) {
// 	fmt.Println("before delete called")
// 	tx.Clauses(clause.Returning{}).Where("goal_id = ?", g.Id).Delete(&Action{})
// 	return
// }

// func (g *GoalMetaData) SetID(id uuid.UUID) {
// 	g.Id = id
// }
