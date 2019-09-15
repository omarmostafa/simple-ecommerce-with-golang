package services

import (
	"github.com/stretchr/testify/assert"
	"simple-ecommerce/app/interfaces/Mocks"
	"simple-ecommerce/app/models"
	"testing"
)

func TestCreateCategoryWithOutRequiredData(t *testing.T) {

	categoryRepository := new(Mocks.ICategoryRepository)
	var category *models.Category
	category = &models.Category{}
	category.ID = 101
	category.Name = ""
	categoryRepository.On("CreateCategory", categoryRepository).Return(nil, nil)

	categoryService := CategoryService{categoryRepository}

	expectedResult := "Forty-Fifteen"

	actualResult, _ := categoryService.CreateCategory(category)

	assert.Equal(t, expectedResult, actualResult)
}
