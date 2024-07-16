# BANK APP
This project is a REST API in Golang that simulates the operation of an ATM. The API supports account creation, deposit, withdrawal and balance check operations.

## Features
- Create new bank accounts
- Top up
- Withdrawing funds
- Balance check
- Using goroutines to process transactions
- Logging all transactions to the console
- Thread-safe balance transactions

## Configuration
Create a `config.yaml` file in the root directory of the project with the following contents:
```yaml
server:
  port: “:8080”
database:
  type: “memory”
```
## Start
To start the server, run the following command from the root directory of the project:
```
go run cmd/main.go
```
## Using the API
### Create an account
```
POST /accounts
Content-Type: application/json
{
  "initial_balance": 1000.00
}
```
### Balance Replenishment
```
POST /accounts/{id}/deposit
Content-Type: application/json
{
    "amount": 200.00
}
```
### Withdrawal
```
POST /accounts/{id}/withdraw
Content-Type: application/json
{
    "amount": 200.00
}
```
### Balance check
```
GET /accounts/{id}/balance
```