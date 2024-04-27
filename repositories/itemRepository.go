package repositories

import (
	"fmt"
	"yuuzin217/go-api-sample/models"
)

type I_ItemRepository interface {
	FindAll() ([]*models.Item, error)
	FindByID(itemID uint) (*models.Item, error)
	Create(newItem *models.Item) (*models.Item, error)
	Update(updateItem *models.Item) (*models.Item, error)
}

type ItemMemoryRepository struct {
	items []*models.Item
}

func NewItemMemoryRepository(items []*models.Item) I_ItemRepository {
	return &ItemMemoryRepository{items: items}
}

func (repo *ItemMemoryRepository) FindAll() ([]*models.Item, error) {
	return repo.items, nil
}

func (repo *ItemMemoryRepository) FindByID(itemID uint) (*models.Item, error) {
	for _, item := range repo.items {
		if itemID == item.ID {
			return item, nil
		}
	}
	return nil, fmt.Errorf("item not found")
}

func (repo *ItemMemoryRepository) Create(newItem *models.Item) (*models.Item, error) {
	newItem.ID = uint(len(repo.items) + 1)
	repo.items = append(repo.items, newItem)
	return newItem, nil
}

func (repo *ItemMemoryRepository) Update(updateItem *models.Item) (*models.Item, error) {
	for i, item := range repo.items {
		if item.ID == updateItem.ID {
			repo.items[i] = updateItem
			return repo.items[i], nil
		}
	}
	return nil, fmt.Errorf("unexpected error")
}
