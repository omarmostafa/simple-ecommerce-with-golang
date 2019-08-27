package repositories

import (
	"simple-ecommerce/app/interfaces"
	"simple-ecommerce/app/models"
)

type ProductRepository struct {
   DBConnection interfaces.IDatabaseConnection
}

func NewProductRepository(conn interfaces.IDatabaseConnection)  *ProductRepository{
	return &ProductRepository{
		DBConnection : conn,
	}
}

func (repo *ProductRepository) GetAllProducts(category *models.Category) ([]*models.Product ,error)  {
	var products []*models.Product
	repo.DBConnection.GetDB().Model(&category).Preload("Categories").Related(&products,"Products")
	return products,nil
}

func (repo *ProductRepository)CreateProduct(product *models.Product,categories []*models.Category) *models.Product {
	repo.DBConnection.GetDB().Table("products").Create(&product)
	repo.DBConnection.GetDB().Model(&product).Association("Categories").Append(categories)
	return product
}

func (repo *ProductRepository)FindProduct(id string) (bool) {
	var product = &models.Product{}
	return !repo.DBConnection.GetDB().Table("products").Where("id = ?", id).First(&product).RecordNotFound()
}

func (repo *ProductRepository)DeleteProduct(id string) (*models.Product){
	product := &models.Product{}
	repo.DBConnection.GetDB().Table("products").Where("id= ?",id).Delete(&product)
	return product
}


func (repo *ProductRepository)UpdateProduct(id string,product *models.Product,categories []*models.Category) (*models.Product){
	repo.DBConnection.GetDB().Table("products").Where("id= ?",id).Updates(&product)
	repo.DBConnection.GetDB().Table("products").Where("id= ?",id).First(&product).Association("Categories").Replace(&categories)
	return product
}