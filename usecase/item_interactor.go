package usecase

import (
	"go_mytodo/domain"
)

type ItemInteractor struct {
	ItemRepository ItemRepository
}

func (interactor *ItemInteractor) AddItem(item *domain.Item) (*domain.Item, error) {
	id, err := interactor.ItemRepository.Add(item)
	if err != nil {
		return nil, err
	}
	newItem, err := interactor.ItemRepository.Get(id)
	return newItem, nil
}

func (interactor *ItemInteractor) ItemById(id int) (*domain.Item, error) {
	item, err := interactor.ItemRepository.Get(id)
	return item, err
}

func (interactor *ItemInteractor) Items() ([]*domain.Item, error) {
	items, err := interactor.ItemRepository.GetAll()
	return items, err
}

func (interactor *ItemInteractor) UpdateItemById(id int, item *domain.Item) (*domain.Item, error) {
	err := interactor.ItemRepository.Update(id, item)
	if err != nil {
		return nil, err
	}
	newItem, err := interactor.ItemRepository.Get(id)
	return newItem, nil
}

func (interactor *ItemInteractor) DeleteItemById(id int) error {
	err := interactor.ItemRepository.Delete(id)
	return err
}
