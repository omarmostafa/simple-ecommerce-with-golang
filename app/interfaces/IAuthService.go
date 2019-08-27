package interfaces

import (
	"simple-ecommerce/app/Requests"
	"simple-ecommerce/app/models"
)

type IAuthService interface {
	RegisterNewUser(user *models.User) (*models.User,error)
	Login(request *Requests.LoginRequest) (string,*models.User,error)
}

