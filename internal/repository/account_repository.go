package repository

import "github.com/velvetriddles/bank-app/internal/domain"

type AccountRepository interface {
	Create(initialBalance float64) (*domain.Account, error)
	GetByID(id int) (*domain.Account, error)
	Update(account *domain.Account) error
}
