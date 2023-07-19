# FinTech API

The FinTech API is a RESTful API that provides functionality for managing accounts and performing financial transactions.

## Folder Structure

FinTech/
|- cmd/
| |- main.go
|
|- internal/
| |- handler/
| | |- account_handler.go
| |
| |- model/
| | |- account.go
| | |- transaction.go
| |
| |- repository/
| | |- account_repository.go
|
|- pkg/
| |- router/
| | |- router.go
|
|- data/
|
|- tests/
| |- home_test.go
| |- list_accounts_test.go
| |- make_transaction_test.go
|
|- go.mod
|- go.sum

## Concurrency Support

The FinTech API is designed to handle concurrent requests efficiently. It utilizes goroutines and synchronization primitives to ensure safe concurrent access to shared resources.

## API Endpoints

### Home

- **URL**: /
- **Method**: GET
- **Description**: Retrieves a welcome message.
- **Concurrency Support**: Yes

### List Accounts

- **URL**: /accounts
- **Method**: GET
- **Description**: Retrieves the list of accounts.
- **Concurrency Support**: Yes

### Make Transaction

- **URL**: /transfer
- **Method**: POST
- **Description**: Initiates a financial transaction from one account to another.
- **Concurrency Support**: Yes
- **Payload**:
  ```json
  {
    "from": "fd796d75-1bcf-4a95-bf1a-f7b296adb79f",
    "to": "3d253e29-8785-464f-8fa0-9e4b57699db9",
    "amount": 100
  }
  ```

## Running the API

1. Clone the repository:

git clone https://github.com/mr-meselmani/FinTech


2. Navigate to the project directory:

cd FinTech

3. Run the application:

go run cmd/main.go

4. The API will be accessible at `http://localhost:3000`.

## Testing

To run the tests for the API, navigate to the project directory and execute the following command:

go test ./tests/... -v

This will execute all the tests in the `tests` package and display the results.

## Dependencies

The FinTech API uses the following third-party libraries:

- `github.com/gorilla/mux`: A powerful URL router and dispatcher for building HTTP services.