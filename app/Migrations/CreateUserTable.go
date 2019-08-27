package Migrations

import (
	"github.com/jinzhu/gorm"
	"simple-ecommerce/app/models"
)

func CreateUserTable(db gorm.DB)  {
	db.AutoMigrate(models.User{})
}
