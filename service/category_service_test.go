package service

import (
	"testing"
	"unit_testing/entity"
	"unit_testing/repository"

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

func TestCategoryService_GetData(t *testing.T) {
	category := entity.Category{Name: "Peter", Id: "1"}

	categoryRepository.Mock.On("FindById", category.Id).Return(category)

	result, err := categoryService.Get("2")

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, category.Id, result.Id, "success")
	assert.Equal(t, category.Name, result.Name, "success")

}
