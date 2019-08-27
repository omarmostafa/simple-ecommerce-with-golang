package models

import "time"

type Product struct {
	ID   int `json:"id"`
	Name string `json:"name" validate:"required"`
	Price string `json:"price" validate:"required"`
	Categories []*Category `gorm:"many2many:products_categories;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time
	DeletedAt *time.Time
}

