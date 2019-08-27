package interfaces

import "simple-ecommerce/app/models"

type ICategoryRepository interface {
	GetAllCategories() ([]*models.Category, error)
    CreateCategory(category *models.Category) *models.Category
    FindCategory(id string) bool
    GetCategory(id string) *models.Category
    DeleteCategory(id string) (*models.Category)
    UpdateCategory(id string,category *models.Category) (*models.Category)
	FindCategoriesByIds(ids []string) []*models.Category
}

