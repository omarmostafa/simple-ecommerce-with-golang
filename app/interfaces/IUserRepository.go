package interfaces

import "simple-ecommerce/app/models"

type IUserRepository interface {
	RegisterNewUser(user *models.User) *models.User
	CheckEmailExistence(email string) bool
	GetUserByEmail(email string) (*models.User,bool)
}

