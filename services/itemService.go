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
	Update(itemID uint, updateItemInput *dto.UpdateItemInput) (*models.Item, error)
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

func (service *ItemService) Update(itemID uint, updateItemInput *dto.UpdateItemInput) (*models.Item, error) {
	targetItem, err := service.FindByID(itemID)
	if err != nil {
		return nil, err
	}
	if updateItemInput.Name != nil {
		targetItem.Name = *updateItemInput.Name
	}
	if updateItemInput.Price != nil {
		targetItem.Price = *updateItemInput.Price
	}
	if updateItemInput.Description != nil {
		targetItem.Description = *updateItemInput.Description
	}
	if updateItemInput.SoldOut != nil {
		targetItem.SoldOut = *updateItemInput.SoldOut
	}
	return service.repository.Update(targetItem)
}
