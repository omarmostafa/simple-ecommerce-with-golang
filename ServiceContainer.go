package main

import (
	"simple-ecommerce/app/controllers"
	"simple-ecommerce/app/infrastructures"
	"simple-ecommerce/app/repositories"
	"simple-ecommerce/app/services"
	"sync"
)

type IServiceContainer interface {
	InjectCategoryController() *controllers.CategoryController
	InjectProductController() *controllers.ProductController
	InjectAuthController() *controllers.AuthController
}

type kernel struct{}

func (k *kernel) InjectCategoryController() *controllers.CategoryController {
	dbConnection := infrastructures.NewDatabaseConnection()
	categoryRepository := repositories.NewCategoryRepository(dbConnection)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryService)

	return categoryController
}

func (k *kernel) InjectProductController() *controllers.ProductController  {
	dbConnection := infrastructures.NewDatabaseConnection()
	productRepository := repositories.NewProductRepository(dbConnection)
	categoryRepo :=repositories.NewCategoryRepository(dbConnection)
	productService := services.NewProductService(productRepository,categoryRepo)
	productController := controllers.NewProductController(productService)

	return productController
}

func (k *kernel) InjectAuthController() *controllers.AuthController  {
	dbConnection := infrastructures.NewDatabaseConnection()
	userRepository := repositories.NewUserRepository(dbConnection)
	authService := services.NewUserService(userRepository)
	authController := controllers.NewAuthController(authService)
	return authController
}

var (
	k             *kernel
	containerOnce sync.Once
)


func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
