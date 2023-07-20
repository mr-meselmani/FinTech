# FinTech API

The FinTech API is a RESTful API that provides functionality for managing accounts and performing financial transactions.

# Folder Structure

- `cmd`:

  - `main.go`: Initializes the application and serves as the main entry point.

- `internal`:

  - `handler`:
    - `account_handler.go`: Handles account-related requests and responses.
  - `model`:
    - `account.go`: Defines the account model.
    - `transaction.go`: Defines the transaction model.
  - `repository`:
    - `account_repository.go`: Implements the repository for account data.

- `pkg`:

  - `router`:
    - `router.go`: Implements the router for request handling.
  - `banner`:
    - Contains a package called `banner` for displaying banners in the application.
    - This package uses the following third-party dependencies:
      - `github.com/common-nighthawk/go-figure`
      - `github.com/fatih/color`

- `data`:

  - `accounts.json`: A JSON file that stores the account information.

- `tests`:

  - `home_test.go`: Tests the home endpoint.
  - `list_accounts_test.go`: Tests the listing of accounts.
  - `make_transaction_test.go`: Tests the making of transactions.

- `go.mod` and `go.sum`: Files related to Go module management.

- `runApi.bat`: Batch script to automate the server start process.

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

```
git clone https://github.com/mr-meselmani/FinTech
```

2. Navigate to the project directory:

```
cd FinTech
```

3. Execute **runApi.bat** to automate running server processes OR skip this line & continue manually.

4. Run the following:

```
go mod download
```

5. Run the application:

```
go run cmd/main.go
```

6. The API will be accessible at `http://localhost:3000`.

## Testing

To run the tests for the API, navigate to the project directory and execute the following command:

1. Test GET `http://localhost:3000/` will send 100 req to the endpoint by running:

```
go test ./tests/home_test.go -v
```

2. Test GET `http://localhost:3000/accounts` will send 100 req by running:

```
go test ./tests/list_accounts_test.go -v
```

3. Test POST `http://localhost:3000/transfer` will send 100 req by running:

```
go test ./tests/make_transaction_test.go -v
```

4. Test them all ( this will depend on your cpu ) by running:

```
go test ./tests/... -v
```

## Dependencies

The FinTech API uses the following third-party libraries:

- `github.com/gorilla/mux`: A powerful URL router and dispatcher for building HTTP services.
- `github.com/common-nighthawk/go-figure`: A package for displaying banners in the application.
- `github.com/fatih/color`: A package for colorizing terminal output.
