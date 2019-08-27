package Transformers

import (
	"simple-ecommerce/app/models"
	"time"
)

type UserTransformer struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (transformer *UserTransformer)SetId(category *models.User)  {
	transformer.Id=category.ID
}

func (transformer *UserTransformer)SetName(category *models.User)  {
	transformer.Name=category.Name
}

func (transformer *UserTransformer)SetCreatedAt(category *models.User)  {
	transformer.CreatedAt=category.CreatedAt.Format(time.RFC822)
}

func (transformer *UserTransformer)SetEmail(category *models.User)  {
	transformer.Email=category.Email
}

