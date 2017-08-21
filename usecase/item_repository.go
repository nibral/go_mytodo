package usecase

import (
	"mytodo/domain"
)

// UseCase(内)からDB(外)を呼ぶためのインターフェイス
type ItemRepository interface {
	Add(*domain.Item) (int, error)
	Get(int) (*domain.Item, error)
	GetAll() ([]*domain.Item, error)
	Update(id int, item *domain.Item) error
	Delete(int) error
}
