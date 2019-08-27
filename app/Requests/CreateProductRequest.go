package Requests

type CreateProductRequest struct {
	ID   int `json:"id"`
	Name string `json:"name" validate:"required"`
	Price string `json:"price" validate:"required"`
	Categories []string `json:"categories" validate:"required"`
}
