package interfaces

import (
	"simple-ecommerce/app/Requests"
	"simple-ecommerce/app/models"
)

type IProductService interface {
	GetAllProductsOfCategories(category_id string) ([]*models.Product, error)
	CreateProduct(product *Requests.CreateProductRequest) (*models.Product, error)
	DeleteProduct(id string) (*models.Product, error)
	UpdateProduct(id string,product *Requests.CreateProductRequest) (*models.Product, error)
}

