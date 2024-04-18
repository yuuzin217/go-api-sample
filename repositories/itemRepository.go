package repositories

import (
	"fmt"
	"yuuzin217/go-api-sample/models"
)

type I_ItemRepository interface {
	FindAll() ([]*models.Item, error)
	FindByID(itemID uint) (*models.Item, error)
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
	return nil, fmt.Errorf("Item not found")
}
