package Mocks

import (
	"github.com/stretchr/testify/mock"
	"simple-ecommerce/app/models"
)

type ICategoryService struct {
	mock.Mock
}


func (_m *ICategoryService) GetAllCategories() ([]*models.Category ,error) {
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

func (_m *ICategoryService) CreateCategory(category *models.Category) (*models.Category ,error) {
	ret := _m.Called()

	var r0 *models.Category
	if rf, ok := ret.Get(0).(func(*models.Category) *models.Category); ok {
		r0 = rf(category)
	} else {
		r0 = ret.Get(0).(*models.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}
	return r0,r1
}

func (_m *ICategoryService) DeleteCategory(id string) (*models.Category ,error) {
	ret := _m.Called()

	var r0 *models.Category
	if rf, ok := ret.Get(0).(func(string) *models.Category); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(*models.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}
	return r0,r1
}

func (_m *ICategoryService) UpdateCategory(id string,category models.Category) (*models.Category ,error) {
	ret := _m.Called()

	var r0 *models.Category
	if rf, ok := ret.Get(0).(func(string,models.Category) *models.Category); ok {
		r0 = rf(id,category)
	} else {
		r0 = ret.Get(0).(*models.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}
	return r0,r1
}