package models

import (
	"time"

	"github.com/google/uuid"
)

type RichText struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	TextObject  string    `json:"text_object,omitempty" gorm:"omitempty"`
	RecruiterID uuid.UUID `json:"recruiter_id,omitempty" gorm:""`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type RichTextPutInput struct {
	Id uuid.UUID `path:"id" required:"true" description:"document id of shift to retrieve"`
	RichText
}

func (in RichTextPutInput) GetData() RichText {
	return in.RichText
}

func (in RichTextPutInput) GetId() uuid.UUID {
	return in.Id
}
