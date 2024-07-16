package router

import (
	"github.com/gorilla/mux"
	"github.com/velvetriddles/bank-app/internal/delivery/http/handler"
	"github.com/velvetriddles/bank-app/internal/delivery/http/middleware"
	"github.com/velvetriddles/bank-app/pkg/logger"
)

func SetupRouter(accountHandler *handler.AccountHandler, log logger.Logger) *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware(log))

	r.HandleFunc("/accounts", accountHandler.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}/deposit", accountHandler.Deposit).Methods("POST")
	r.HandleFunc("/accounts/{id}/withdraw", accountHandler.Withdraw).Methods("POST")
	r.HandleFunc("/accounts/{id}/balance", accountHandler.GetBalance).Methods("GET")

	return r
}
