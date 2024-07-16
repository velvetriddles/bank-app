package usecase

import (
	"github.com/velvetriddles/bank-app/internal/domain"
	"github.com/velvetriddles/bank-app/internal/repository"
	"github.com/velvetriddles/bank-app/pkg/logger"
)

type AccountUseCase struct {
	repo repository.AccountRepository
	log  logger.Logger
}

func NewAccountUseCase(repo repository.AccountRepository, log logger.Logger) *AccountUseCase {
	return &AccountUseCase{
		repo: repo,
		log:  log,
	}
}

func (uc *AccountUseCase) CreateAccount(initialBalance float64) (*domain.Account, error) {
	createChan := make(chan *domain.Account)
	errorChan := make(chan error)

	go func() {
		account, err := uc.repo.Create(initialBalance)
		if err != nil {
			errorChan <- err
			return
		}
		createChan <- account
	}()

	select {
	case account := <-createChan:
		uc.log.Info("Account created", "id", account.ID, "initial_balance", initialBalance)
		return account, nil
	case err := <-errorChan:
		uc.log.Error("Failed to create account", "error", err)
		return nil, err
	}
}

func (uc *AccountUseCase) Deposit(id int, amount float64) (*domain.Account, error) {
	accountChan := make(chan *domain.Account)
	errorChan := make(chan error)

	go func() {
		account, err := uc.repo.GetByID(id)
		if err != nil {
			errorChan <- err
			return
		}
		if err := account.Deposit(amount); err != nil {
			errorChan <- err
			return
		}
		if err := uc.repo.Update(account); err != nil {
			errorChan <- err
			return
		}
		accountChan <- account
	}()

	select {
	case account := <-accountChan:
		uc.log.Info("Deposit successful", "id", id, "amount", amount, "new_balance", account.GetBalance())
		return account, nil
	case err := <-errorChan:
		uc.log.Error("Deposit failed", "id", id, "amount", amount, "error", err)
		return nil, err
	}
}

func (uc *AccountUseCase) Withdraw(id int, amount float64) (*domain.Account, error) {
	accountChan := make(chan *domain.Account)
	errorChan := make(chan error)

	go func() {
		account, err := uc.repo.GetByID(id)
		if err != nil {
			errorChan <- err
			return
		}
		if err := account.Withdraw(amount); err != nil {
			errorChan <- err
			return
		}
		if err := uc.repo.Update(account); err != nil {
			errorChan <- err
			return
		}
		accountChan <- account
	}()

	select {
	case account := <-accountChan:
		uc.log.Info("Withdrawal successful", "id", id, "amount", amount, "new_balance", account.GetBalance())
		return account, nil
	case err := <-errorChan:
		uc.log.Error("Withdrawal failed", "id", id, "amount", amount, "error", err)
		return nil, err
	}
}

func (uc *AccountUseCase) GetBalance(id int) (float64, error) {
	balanceChan := make(chan float64)
	errorChan := make(chan error)

	go func() {
		account, err := uc.repo.GetByID(id)
		if err != nil {
			errorChan <- err
			return
		}
		balanceChan <- account.GetBalance()
	}()

	select {
	case balance := <-balanceChan:
		uc.log.Info("Balance checked", "id", id, "balance", balance)
		return balance, nil
	case err := <-errorChan:
		uc.log.Error("Failed to get balance", "id", id, "error", err)
		return 0, err
	}
}
