package repositories

import (
	"simple-ecommerce/app/interfaces"
	"simple-ecommerce/app/models"
)

type UserRepository struct {
   DBConnection interfaces.IDatabaseConnection
}

func NewUserRepository(conn interfaces.IDatabaseConnection)  *UserRepository{
	return &UserRepository{
		DBConnection : conn,
	}
}

func (repo *UserRepository) RegisterNewUser(user *models.User) *models.User  {
	repo.DBConnection.GetDB().Table("users").Create(&user)
	return user
}

func (repo *UserRepository) CheckEmailExistence(email string) bool  {
	var user = &models.User{}
	return !repo.DBConnection.GetDB().Table("users").Where("email = ?", email).First(&user).RecordNotFound()
}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User,bool)  {
	var user = &models.User{}
	check :=repo.DBConnection.GetDB().Table("users").Where("email = ?", email).First(&user).RecordNotFound()
	return user,!check
}