package interfaces

import "simple-ecommerce/app/models"

type ICategoryService interface {
	GetAllCategories() ([]*models.Category, error)
	CreateCategory(category *models.Category) (*models.Category, error)
	DeleteCategory(id string) (*models.Category, error)
	UpdateCategory(id string,category *models.Category) (*models.Category, error)
}

