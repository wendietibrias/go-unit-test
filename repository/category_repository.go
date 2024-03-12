package repository

import (
	"unit_testing/entity"
)

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
