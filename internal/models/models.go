package models

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}