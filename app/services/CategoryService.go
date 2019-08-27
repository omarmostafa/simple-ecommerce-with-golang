package services

import (
	"simple-ecommerce/app/Errors"
	"simple-ecommerce/app/infrastructures"
	"simple-ecommerce/app/interfaces"
	"simple-ecommerce/app/models"
)

type CategoryService struct {
	repo interfaces.ICategoryRepository
}
func NewCategoryService(repo interfaces.ICategoryRepository)  *CategoryService{
	return &CategoryService{
		repo : repo,
	}
}

func (service *CategoryService)GetAllCategories() ([]*models.Category,error)  {
	categories ,err :=  service.repo.GetAllCategories()
	if err!=nil {
		return nil,err
	}
	return categories,nil
}

func (service *CategoryService)CreateCategory(category *models.Category) (*models.Category,error)  {
	if validationError:= infrastructures.Validate(category); validationError!=nil{
		return nil,validationError
	}
	category = service.repo.CreateCategory(category)
	return category,nil
}

func (service *CategoryService)DeleteCategory(id string) (*models.Category,error)  {
	if err := service.checkCategoryExistance(id); err!=nil {
		return  nil,err
	}
	deleted_Category :=service.repo.DeleteCategory(id)
	return deleted_Category,nil
}

func (service *CategoryService)UpdateCategory(id string,newCategory *models.Category) (*models.Category,error)  {
	if err := service.checkCategoryExistance(id); err!=nil {
		return  nil,err
	}
	if validationError:= infrastructures.Validate(newCategory); validationError!=nil{
		return nil,validationError
	}
	updatedCategory :=service.repo.UpdateCategory(id,newCategory)
	return updatedCategory,nil
}

func (service *CategoryService)checkCategoryExistance(id string) (error)  {
	category :=service.repo.FindCategory(id)
	if !category {
		err :=Errors.NotFoundError{"Category Not Found"}
		return  err
	}
	return nil
}


