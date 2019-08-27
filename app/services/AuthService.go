package services

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"simple-ecommerce/app/Errors"
	"simple-ecommerce/app/Requests"
	"simple-ecommerce/app/infrastructures"
	"simple-ecommerce/app/interfaces"
	"simple-ecommerce/app/models"
	"time"
)

type AuthService struct {
	repo interfaces.IUserRepository
}
func NewUserService(repo interfaces.IUserRepository)  *AuthService{
	return &AuthService{
		repo : repo,
	}
}

func (service *AuthService)RegisterNewUser(user *models.User) (*models.User,error)  {
	if validationError:= infrastructures.Validate(user); validationError!=nil{
		return nil,validationError
	}
	err := service.checkEmailExistance(user.Email)
	if err !=nil {
		return  nil,err
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	user = service.repo.RegisterNewUser(user)
	return user,nil
}

func (service *AuthService)Login(request *Requests.LoginRequest) (string,*models.User,error)  {
	if validationError:= infrastructures.Validate(request); validationError!=nil{
		return "",nil,validationError
	}
	user,check := service.repo.GetUserByEmail(request.Email)
	if !check {
		return  "",nil,Errors.NotFoundError{"Email Not found"}
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return "",nil,Errors.NotFoundError{"Invalid credintials"}
	}

	expirationTime := time.Now().Add(5000 * time.Hour)

	tk := &models.Token{UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
	},}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	return tokenString,user,nil
}


func (service *AuthService)checkEmailExistance(email string) (error)  {
	user :=service.repo.CheckEmailExistence(email)
	if user {
		err :=Errors.ValidationError{"Email Already Exist"}
		return  err
	}
	return nil
}