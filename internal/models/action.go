package models

import (
	"time"

	"github.com/google/uuid"
)

type Action struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	Done      bool      `json:"done,omitempty" gorm:"omitempty"`
	Value     int       `json:"value,omitempty" gorm:"omitempty"`
	GoalID    uuid.UUID `json:"goal_id,omitempty" gorm:""`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type ActionPutInput struct {
	Id uuid.UUID `path:"id" required:"true" description:"document id of shift to retrieve"`
	Action
}

func (in ActionPutInput) GetData() Action {
	return in.Action
}

func (in ActionPutInput) GetId() uuid.UUID {
	return in.Id
}
