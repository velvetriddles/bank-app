package memory

import (
	"sync"

	"github.com/velvetriddles/bank-app/internal/domain"
	"github.com/velvetriddles/bank-app/internal/repository"
)

type InMemoryAccountRepository struct {
	accounts map[int]*domain.Account
	mu       sync.RWMutex
	nextID   int
}

func NewInMemoryAccountRepository() repository.AccountRepository {
	return &InMemoryAccountRepository{
		accounts: make(map[int]*domain.Account),
		nextID:   1,
	}
}

func (r *InMemoryAccountRepository) Create(initialBalance float64) (*domain.Account, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	account := &domain.Account{ID: r.nextID, Balance: initialBalance}
	r.accounts[r.nextID] = account
	r.nextID++

	return account, nil
}

func (r *InMemoryAccountRepository) GetByID(id int) (*domain.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	account, ok := r.accounts[id]
	if !ok {
		return nil, domain.ErrAccountNotFound
	}

	return account, nil
}

func (r *InMemoryAccountRepository) Update(account *domain.Account) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.accounts[account.ID] = account
	return nil
}
