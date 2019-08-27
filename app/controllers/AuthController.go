package controllers

import (
	"encoding/json"
	"net/http"
	"simple-ecommerce/app/Errors"
	"simple-ecommerce/app/Requests"
	"simple-ecommerce/app/Transformers"
	"simple-ecommerce/app/interfaces"
	"simple-ecommerce/app/models"
)
type AuthController struct {
	service interfaces.IAuthService
}

func NewAuthController(service interfaces.IAuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (controller *AuthController) Register(res http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(req.Body).Decode(&user)
	user,err = controller.service.RegisterNewUser(user)
	if err != nil {
		switch err.(type) {
		case Errors.ValidationError:
			RespondBadRequest(res, req, err.Error(), nil)
			return
		default:
			data := Transformers.ParsingValidationError(err)
			RespondBadRequest(res, req, "Input Validations Errors", data)
			return
		}
	}

	RespondCreated(res, req, "User Created successfully", nil)
}

func (controller *AuthController) Login(res http.ResponseWriter, req *http.Request) {
	request := &Requests.LoginRequest{}
	err := json.NewDecoder(req.Body).Decode(&request)
	token,user,err := controller.service.Login(request)
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

	transformedUser:=Transformers.UserTransformer{}
	Transformers.TransformItem(&transformedUser,user)
	RespondCreated(res, req, "Logged in", map[string]interface{}{"token":token,"user":transformedUser})
}