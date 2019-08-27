package Transformers

import (
	"simple-ecommerce/app/models"
	"time"
)

type ProductTransformer struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Price string `json:"price"`
	Categories []CategoryTransformer `json:"categories"`
	CreatedAt string `json:"created_at"`
}

func (transformer *ProductTransformer)SetId(product *models.Product)  {
	transformer.Id=product.ID
}

func (transformer *ProductTransformer)SetName(product *models.Product)  {
	transformer.Name=product.Name
}

func (transformer *ProductTransformer)SetCreatedAt(product *models.Product)  {
	transformer.CreatedAt=product.CreatedAt.Format(time.RFC822)
}


func (transformer *ProductTransformer)SetPrice(product *models.Product)  {
	transformer.Price=product.Price
}

func (transformer *ProductTransformer)SetCategories(product *models.Product)  {
	transformer.Categories=TransformCollectionOfCategory(product.Categories)
}

func TransformCollectionOfProduct(data []*models.Product) []ProductTransformer  {
	var collection []ProductTransformer
	for _,item :=range data {
		itemTransformer := ProductTransformer{}
		TransformItem(&itemTransformer,item)
		collection = append(collection, itemTransformer)
	}
	return collection
}
