package models

import "github.com/google/uuid"

type Person struct {
	Id   uuid.UUID `json:"id,omitempty" gorm:"primaryKey"`
	Name string    `json:"name,omitempty" gorm:"omitempty"`
	Age  int       `json:"age,omitempty"`
}

func (p *Person) SetID(id uuid.UUID) {
	p.Id = id
}

type DeleteInput struct {
	Id uuid.UUID `path:"id" required:"true" description:"document id to delete"`
}

type EmptyBody struct{}
