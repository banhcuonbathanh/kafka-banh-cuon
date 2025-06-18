package interfaces

import (
	"main-service/pkg/domain"
	"main-service/pkg/utils/models"
)

type CategoryUseCase interface {
	AddCategory(category domain.Category) (domain.Category, error)
	UpdateCategory(current string, new string) (domain.Category, error)
	DeleteCategory(categoryID string) error
	GetCategories() ([]domain.Category, error)
	GetProductDetailsInACategory(id int) ([]models.Inventories, error)
	GetBannersForUsers() ([]models.Banner, error)
}
