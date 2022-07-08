package models

type Book struct {
	Model
	Title    string `json:"title" validate:"required,gte=6" gorm:"unique"`
	Author   string `json:"author" validate:"required"`
	Quantity int    `json:"quantity"`
}
