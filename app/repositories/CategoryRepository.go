package repositories

import (
	"simple-ecommerce/app/interfaces"
	"simple-ecommerce/app/models"
)

type CategoryRepository struct {
   DBConnection interfaces.IDatabaseConnection
}

func NewCategoryRepository(conn interfaces.IDatabaseConnection)  *CategoryRepository{
	return &CategoryRepository{
		DBConnection : conn,
	}
}

func (repo *CategoryRepository) GetAllCategories() ([]*models.Category ,error)  {
	var categories []*models.Category
	repo.DBConnection.GetDB().Table("categories").Find(&categories)
	return categories,nil
}

func (repo *CategoryRepository)CreateCategory(category *models.Category) *models.Category {
	repo.DBConnection.GetDB().Table("categories").Create(&category)
	return category
}

func (repo *CategoryRepository)FindCategory(id string) (bool) {
	var category = &models.Category{}
	return !repo.DBConnection.GetDB().Table("categories").Where("id = ?", id).First(&category).RecordNotFound()
}
func (repo *CategoryRepository)GetCategory(id string) *models.Category {
	var category = &models.Category{}
	repo.DBConnection.GetDB().Table("categories").Where("id = ?", id).First(&category)
	return category
}
func (repo *CategoryRepository)DeleteCategory(id string) (*models.Category){
	category := &models.Category{}
	repo.DBConnection.GetDB().Table("categories").Where("id= ?",id).Delete(&category)
	return category
}


func (repo *CategoryRepository)UpdateCategory(id string,category *models.Category) (*models.Category){
	repo.DBConnection.GetDB().Table("categories").Where("id= ?",id).Updates(&category)
	return category
}

func (repo *CategoryRepository)FindCategoriesByIds(ids []string) []*models.Category{
	var categories []*models.Category
	repo.DBConnection.GetDB().Table("categories").Where("id in (?)",ids).Find(&categories)
	return categories
}

