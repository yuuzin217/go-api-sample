package services

import (
	"yuuzin217/go-api-sample/dto"
	"yuuzin217/go-api-sample/models"
	"yuuzin217/go-api-sample/repositories"
)

type I_ItemService interface {
	FindAll() ([]*models.Item, error)
	FindByID(itemID uint) (*models.Item, error)
	Create(createItemInput *dto.CreateItemInput) (*models.Item, error)
}

type ItemService struct {
	repository repositories.I_ItemRepository // interface を持たせる（実装の変更を容易にするため）
}

func NewItemService(repository repositories.I_ItemRepository) I_ItemService {
	return &ItemService{repository: repository}
}

func (service *ItemService) FindAll() ([]*models.Item, error) {
	return service.repository.FindAll()
}

func (service *ItemService) FindByID(itemID uint) (*models.Item, error) {
	return service.repository.FindByID(itemID)
}

func (service *ItemService) Create(createItemInput *dto.CreateItemInput) (*models.Item, error) {
	newItem := &models.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
	}
	return service.repository.Create(newItem)
}
