package models

type Person struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func (p *Person) SetID(id string) {
	p.Id = id
}

type DeleteInput struct {
	Id string `path:"id" required:"true" description:"document id to delete"`
}

type EmptyBody struct{}
