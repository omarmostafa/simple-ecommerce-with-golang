package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"simple-ecommerce/app/Errors"
	"simple-ecommerce/app/Transformers"
	"simple-ecommerce/app/interfaces"
	"simple-ecommerce/app/models"
)

type CategoryController struct {
	service interfaces.ICategoryService
}

func NewCategoryController(service interfaces.ICategoryService) *CategoryController {
	return &CategoryController{
		service: service,
	}
}

func (controller *CategoryController) GetCategories(res http.ResponseWriter, req *http.Request) {
	categories, _ := controller.service.GetAllCategories()
	categoriesTransformed := Transformers.TransformCollectionOfCategory(categories)
	RespondAccepted(res,req,"",categoriesTransformed)
}

func (controller *CategoryController) CreateCategory(res http.ResponseWriter, req *http.Request) {
	category := &models.Category{}
	err := json.NewDecoder(req.Body).Decode(&category)
	category, err = controller.service.CreateCategory(category)
	if err != nil {
		data := Transformers.ParsingValidationError(err)
		RespondBadRequest(res, req, "Input Validations Errors", data)
		return
	}
	transformedCategory:=Transformers.CategoryTransformer{}
	Transformers.TransformItem(&transformedCategory,category)
	RespondCreated(res, req, "Category created succssfully", transformedCategory)
}

func (controller *CategoryController) DeleteCategory(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	_, err := controller.service.DeleteCategory(id)
	if err != nil {
		RespondNotFound(res, req, err.Error(), nil)
		return
	}
	RespondCreated(res, req, "Category deleted successfully", nil)
}

func (controller *CategoryController) UpdateCategory(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	category := &models.Category{}
	err := json.NewDecoder(req.Body).Decode(&category)
	category, err = controller.service.UpdateCategory(id, category)
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
	RespondCreated(res, req, "Category updated successfully", nil)
}
