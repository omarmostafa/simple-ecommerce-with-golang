package Transformers

import (
	"simple-ecommerce/app/models"
	"time"
)

type CategoryTransformer struct {
	Id int `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func (transformer *CategoryTransformer)SetId(category *models.Category)  {
	transformer.Id=category.ID
}

func (transformer *CategoryTransformer)SetName(category *models.Category)  {
	transformer.Name=category.Name
}

func (transformer *CategoryTransformer)SetCreatedAt(category *models.Category)  {
	transformer.CreatedAt=category.CreatedAt.Format(time.RFC822)
}

func TransformCollectionOfCategory(data []*models.Category) []CategoryTransformer {
	var collection []CategoryTransformer
	for _,item :=range data {
		itemTransformer := CategoryTransformer{}
		TransformItem(&itemTransformer,item)
		collection = append(collection, itemTransformer)
	}
	return collection
}
