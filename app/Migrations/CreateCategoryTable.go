package Migrations

import (
	"github.com/jinzhu/gorm"
	"simple-ecommerce/app/models"
)

func CreateCategoriesTable(db gorm.DB)  {
	db.AutoMigrate(models.Category{})
}
