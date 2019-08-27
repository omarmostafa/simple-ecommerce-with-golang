package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"simple-ecommerce/app/Errors"
	"simple-ecommerce/app/Requests"
	"simple-ecommerce/app/Transformers"
	"simple-ecommerce/app/interfaces"
)

type ProductController struct {
	service interfaces.IProductService
}

func NewProductController(service interfaces.IProductService) *ProductController {
	return &ProductController{
		service: service,
	}
}

func (controller *ProductController) GetProductsOfCategory(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	products, err := controller.service.GetAllProductsOfCategories(id)
	if err!=nil{
		RespondNotFound(res, req, err.Error(), nil)
		return
	}
	productsTransformed := Transformers.TransformCollectionOfProduct(products)
	RespondAccepted(res,req,"",productsTransformed)
}

func (controller *ProductController) CreateProduct(res http.ResponseWriter, req *http.Request) {
	productRequest := &Requests.CreateProductRequest{}
	err := json.NewDecoder(req.Body).Decode(&productRequest)
	product, err := controller.service.CreateProduct(productRequest)
	if err != nil {
		switch err.(type) {
		case Errors.NotFoundError:
			RespondNotFound(res, req, err.Error(), nil)
			return
		default:
			data := Transformers.ParsingValidationError(err)
			RespondBadRequest(res, req, "Input Validations Errors", data)
			return
		}
	}
	transformedProduct :=Transformers.ProductTransformer{}
	Transformers.TransformItem(&transformedProduct,product)
	RespondCreated(res, req, "Product created succssfully", transformedProduct)
}

func (controller *ProductController) DeleteProduct(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	_, err := controller.service.DeleteProduct(id)
	if err != nil {
		RespondNotFound(res, req, err.Error(), nil)
		return
	}
	RespondCreated(res, req, "Product deleted successfully", nil)
}

func (controller *ProductController) UpdateProduct(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	productRequest := &Requests.CreateProductRequest{}
	err := json.NewDecoder(req.Body).Decode(&productRequest)
	_, err  = controller.service.UpdateProduct(id, productRequest)
	if err != nil {
		switch err.(type) {
		case Errors.NotFoundError:
			RespondNotFound(res, req, err.Error(), nil)
			return
		default:
			data := Transformers.ParsingValidationError(err)
			RespondBadRequest(res, req, "Input Validations Errors", data)
			return
		}
	}
	RespondCreated(res, req, "Product updated successfully", nil)
}
