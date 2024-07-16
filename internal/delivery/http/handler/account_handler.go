package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/velvetriddles/bank-app/internal/domain"
	"github.com/velvetriddles/bank-app/internal/usecase"
	"github.com/velvetriddles/bank-app/pkg/logger"
)

type AccountHandler struct {
	useCase *usecase.AccountUseCase
	log     logger.Logger
}

func NewAccountHandler(useCase *usecase.AccountUseCase, log logger.Logger) *AccountHandler {
	return &AccountHandler{
		useCase: useCase,
		log:     log,
	}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req struct {
		InitialBalance *float64 `json:"initial_balance"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Invalid request body", "error", err)
		h.sendErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	initialBalance := 0.0
	if req.InitialBalance != nil {
		initialBalance = *req.InitialBalance
	}

	account, err := h.useCase.CreateAccount(initialBalance)
	if err != nil {
		h.log.Error("Failed to create account", "error", err)
		switch err {
		case domain.ErrInvalidAmount:
			h.sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		default:
			h.sendErrorResponse(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	
	h.log.Info("Account created successfully", "account_id", account.ID)
	h.sendJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "Account created successfully",
		"account": account,
	})
}

func (h *AccountHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	
	var req struct {
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Invalid request body", "error", err)
		h.sendErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	updatedAccount, err := h.useCase.Deposit(id, req.Amount)
	if err != nil {
		h.log.Error("Deposit failed", "account_id", id, "amount", req.Amount, "error", err)
		switch err {
		case domain.ErrAccountNotFound:
			h.sendErrorResponse(w, err.Error(), http.StatusNotFound)
		case domain.ErrInvalidAmount:
			h.sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		default:
			h.sendErrorResponse(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	
	h.log.Info("Deposit successful", "account_id", id, "amount", req.Amount)
	h.sendJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Deposit successful",
		"account": updatedAccount,
	})
}

func (h *AccountHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	
	var req struct {
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Invalid request body", "error", err)
		h.sendErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	updatedAccount, err := h.useCase.Withdraw(id, req.Amount)
	if err != nil {
		h.log.Error("Withdrawal failed", "account_id", id, "amount", req.Amount, "error", err)
		switch err {
		case domain.ErrAccountNotFound:
			h.sendErrorResponse(w, err.Error(), http.StatusNotFound)
		case domain.ErrInvalidAmount, domain.ErrInsufficientFunds:
			h.sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		default:
			h.sendErrorResponse(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	
	h.log.Info("Withdrawal successful", "account_id", id, "amount", req.Amount)
	h.sendJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Withdrawal successful",
		"account": updatedAccount,
	})
}

func (h *AccountHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	
	balance, err := h.useCase.GetBalance(id)
	if err != nil {
		h.log.Error("Failed to get balance", "account_id", id, "error", err)
		if err == domain.ErrAccountNotFound {
			h.sendErrorResponse(w, err.Error(), http.StatusNotFound)
		} else {
			h.sendErrorResponse(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	
	h.log.Info("Balance retrieved", "account_id", id, "balance", balance)
	h.sendJSONResponse(w, http.StatusOK, map[string]interface{}{
		"account_id": id,
		"balance":    balance,
	})
}

func (h *AccountHandler) sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.log.Error("Failed to encode response", "error", err)
	}
}

func (h *AccountHandler) sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": message}); err != nil{
		h.log.Error("Failed to encode error response", "error", err)	
	}
}