package repositories

import "yuuzin217/go-api-sample/models"

type I_ItemRepository interface {
	FindAll() ([]*models.Item, error)
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
