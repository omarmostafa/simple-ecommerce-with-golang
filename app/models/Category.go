package models

import "time"

type Category struct {
	ID   int `json:"id"`
	Name string `json:"name" validate:"required"`
	Products []Product `gorm:"many2many:products_categories;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time
	DeletedAt *time.Time
}


func NewCategory(name string) *Category {
	return &Category{
		Name: name,
	}
}

func (u *Category) GetName() string {
	return u.Name
}