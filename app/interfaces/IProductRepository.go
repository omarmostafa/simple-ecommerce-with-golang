package interfaces

import "simple-ecommerce/app/models"

type IProductRepository interface {
	GetAllProducts(category *models.Category) ([]*models.Product, error)
    CreateProduct(product *models.Product,categories []*models.Category) *models.Product
    FindProduct(id string) bool
    DeleteProduct(id string) (*models.Product)
    UpdateProduct(id string,product *models.Product,categories []*models.Category) (*models.Product)
}

