package models

import (
	"time"

	"github.com/google/uuid"
)

type Recruiter struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name,omitempty" gorm:"omitempty"`
	Company     string    `json:"company,omitempty" gorm:"omitempty"`
	Status      string    `json:"status,omitempty" gorm:"omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	LastContact time.Time `json:"last_contact,omitempty"`
	RichText    RichText  `gorm:"foreignKey:RecruiterID;constraint:OnDelete:CASCADE;"`
}

type RecruiterPutInput struct {
	Id uuid.UUID `path:"id" required:"true" description:"document id of shift to retrieve"`
	Recruiter
}

func (in RecruiterPutInput) GetData() Recruiter {
	return in.Recruiter
}

func (in RecruiterPutInput) GetId() uuid.UUID {
	return in.Id
}
