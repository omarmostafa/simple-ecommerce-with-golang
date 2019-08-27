package Migrations

import (
	"github.com/jinzhu/gorm"
	"simple-ecommerce/app/models"
)

func CreateProductTable(db gorm.DB)  {
	db.AutoMigrate(models.Product{})
	db.Table("products_categories").AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
	db.Table("products_categories").AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")
}
