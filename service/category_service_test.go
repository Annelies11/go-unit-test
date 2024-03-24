package service

import (
	"testing"

	"github.com/Annelies11/go-unit-test.git/entity"
	"github.com/Annelies11/go-unit-test.git/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Arduino",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)

	res, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, category.Id, res.Id)
	assert.Equal(t, category.Name, res.Name)
}
