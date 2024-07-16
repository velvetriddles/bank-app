package domain

import "sync"

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
	mu      sync.Mutex
}

func NewAccount(id int, initialBalance float64) (*Account, error) {
	if initialBalance < 0 {
		return nil, ErrInvalidAmount
	}
	return &Account{
		ID:      id,
		Balance: initialBalance,
	}, nil
}


func (a *Account) Deposit(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount <= 0 {
		return ErrInvalidAmount
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if a.Balance < amount {
		return ErrInsufficientFunds
	}
	a.Balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.Balance
}
