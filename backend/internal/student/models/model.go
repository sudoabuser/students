package models

import "github.com/google/uuid"

type Student struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type StudentEntity struct {
	ID      uuid.UUID `gorm:"primary_key;type:uuid"`
	Name    string
	Surname string
}

func (Student) TableName() string { // By default, plural of struct's name ('students') is the table name used.
	return "students" // using this syntax to specify the tablename.
}

func (StudentEntity) TableName() string {
	return "students"
}

type Page struct {
	Number   int `json:"pageNumber"`
	Size     int `json:"pageSize"`
	Elements int `json:"totalElements"`
	Pages    int `json:"totalPages"`
}

type PaginationResponse struct {
	Students []Student `json:"students"`
	Page     Page      `json:"page"`
}
