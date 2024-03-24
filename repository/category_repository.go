package repository

import "github.com/Annelies11/go-unit-test.git/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
