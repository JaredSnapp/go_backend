package models

import (
	"time"

	"github.com/google/uuid"
)

// TODO: Add user
// TODO: Look into relations. Type should probably be relation to goal
type Action struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey"`
	Type      string    `json:"type,omitempty" gorm:"omitempty"`
	Done      bool      `json:"done,omitempty" gorm:"omitempty"`
	Value     int       `json:"value,omitempty" gorm:"omitempty"`
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
