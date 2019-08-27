package services

import (
	"simple-ecommerce/app/Errors"
	"simple-ecommerce/app/Requests"
	"simple-ecommerce/app/infrastructures"
	"simple-ecommerce/app/interfaces"
	"simple-ecommerce/app/models"
)

type ProductService struct {
	repo interfaces.IProductRepository
	categoryRepo interfaces.ICategoryRepository
}
func NewProductService(repo interfaces.IProductRepository,categoryRepo interfaces.ICategoryRepository)  *ProductService{
	return &ProductService{
		repo : repo,
		categoryRepo:categoryRepo,
	}
}

func (service *ProductService)GetAllProductsOfCategories(category_id string) ([]*models.Product,error)  {
	if err := service.checkCategoryExistance(category_id); err!=nil {
		return  nil,err
	}
	category :=service.categoryRepo.GetCategory(category_id)
	products ,err :=  service.repo.GetAllProducts(category)
	if err!=nil {
		return nil,err
	}
	return products,nil
}

func (service *ProductService)CreateProduct(productRequest *Requests.CreateProductRequest) (*models.Product,error)  {
	if validationError:= infrastructures.Validate(productRequest); validationError!=nil{
		return nil,validationError
	}
	product := &models.Product{Name:productRequest.Name,Price:productRequest.Price}
	categories := service.categoryRepo.FindCategoriesByIds(productRequest.Categories)
	if len(categories) < len(productRequest.Categories){
		return  nil,Errors.NotFoundError{"Some Categories not found"}
	}
	product = service.repo.CreateProduct(product,categories)
	return product,nil
}

func (service *ProductService)DeleteProduct(id string) (*models.Product,error)  {
	if err := service.checkProductExistance(id); err!=nil {
		return  nil,err
	}
	deleted_Product :=service.repo.DeleteProduct(id)
	return deleted_Product,nil
}

func (service *ProductService)UpdateProduct(id string,productRequest *Requests.CreateProductRequest) (*models.Product,error)  {
	if err := service.checkProductExistance(id); err!=nil {
		return  nil,err
	}
	if validationError:= infrastructures.Validate(productRequest); validationError!=nil{
		return nil,validationError
	}
	product := &models.Product{Name:productRequest.Name,Price:productRequest.Price}
	categories := service.categoryRepo.FindCategoriesByIds(productRequest.Categories)
	if len(categories) < len(productRequest.Categories){
		return  nil,Errors.NotFoundError{"Some Categories not found"}
	}
	updatedProduct :=service.repo.UpdateProduct(id,product,categories)
	return updatedProduct,nil
}

func (service *ProductService)checkProductExistance(id string) (error)  {
	product :=service.repo.FindProduct(id)
	if !product {
		err :=Errors.NotFoundError{"Product Not Found"}
		return  err
	}
	return nil
}

func (service *ProductService)checkCategoryExistance(id string) (error)  {
	category :=service.categoryRepo.FindCategory(id)
	if !category {
		err :=Errors.NotFoundError{"Category Not Found"}
		return  err
	}
	return nil
}



