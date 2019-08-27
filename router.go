package main

import (
	"github.com/gorilla/mux"
	"simple-ecommerce/app/Middleware"
	"sync"
)

type IChiRouter interface {
	InitRouter() *mux.Router
}

type router struct {

}

func (router *router) InitRouter() *mux.Router  {
	r:= mux.NewRouter()

    r.Use(Middleware.JwtAuthentication)
	apis :=r.PathPrefix("/api/v1").Subrouter()
	categoryController := ServiceContainer().InjectCategoryController()
	apis.HandleFunc("/categories", categoryController.GetCategories).Methods("GET").Name("get_categories")
	apis.HandleFunc("/categories", categoryController.CreateCategory).Methods("POST")
	apis.HandleFunc("/categories/{id}", categoryController.DeleteCategory).Methods("DELETE")
	apis.HandleFunc("/categories/{id}", categoryController.UpdateCategory).Methods("PUT")

	productController :=ServiceContainer().InjectProductController()
	apis.HandleFunc("/categories/{id}/products", productController.GetProductsOfCategory).Methods("GET").Name("get_products")
	apis.HandleFunc("/products", productController.CreateProduct).Methods("POST")
	apis.HandleFunc("/products/{id}", productController.DeleteProduct).Methods("DELETE")
	apis.HandleFunc("/products/{id}", productController.UpdateProduct).Methods("PUT")

	authController :=ServiceContainer().InjectAuthController()
	apis.HandleFunc("/users/register", authController.Register).Methods("POST").Name("register")
	apis.HandleFunc("/users/login", authController.Login).Methods("POST").Name("login")

	return  r
}

var (
	m *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return  m;
}
