package Mocks

import (
	"github.com/stretchr/testify/mock"
	"simple-ecommerce/app/models"
)

type ICategoryRepository struct {
	mock.Mock
}

func (_m *ICategoryRepository)CreateCategory(category *models.Category) *models.Category {
	ret := _m.Called(category)

	var r0 *models.Category
	if rf, ok := ret.Get(0).(func(*models.Category) *models.Category); ok {
		r0 = rf(category)
	} else {
		r0 = ret.Get(0).(*models.Category)
	}
	return r0
}

func (_m *ICategoryRepository) GetAllCategories() ([]*models.Category ,error) {
	ret := _m.Called()

	var r0 []*models.Category
	if rf, ok := ret.Get(0).(func() []*models.Category); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).([]*models.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}
	return r0,r1
}

func (_m *ICategoryRepository)FindCategory(id string) (bool){
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}
	return r0
}

func (_m *ICategoryRepository)DeleteCategory(id string) (*models.Category){
	ret := _m.Called()

	var r0 *models.Category
	if rf, ok := ret.Get(0).(func(string) *models.Category); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(*models.Category)
	}
	return r0
}

func (_m *ICategoryRepository)UpdateCategory(id string,category *models.Category) (*models.Category){
	ret := _m.Called()

	var r0 *models.Category
	if rf, ok := ret.Get(0).(func(string,*models.Category) *models.Category); ok {
		r0 = rf(id,category)
	} else {
		r0 = ret.Get(0).(*models.Category)
	}
	return r0
}

func (_m *ICategoryRepository)FindCategoriesByIds(ids []string) []*models.Category{
	ret := _m.Called()

	var r0 []*models.Category
	if rf, ok := ret.Get(0).(func([]string) []*models.Category); ok {
		r0 = rf(ids)
	} else {
		r0 = ret.Get(0).([]*models.Category)
	}
	return r0
}