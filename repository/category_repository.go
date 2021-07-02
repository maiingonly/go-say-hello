package repository

import "github.com/maiingonly/go-say-hello/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
