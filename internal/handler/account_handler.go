package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/mr-meselmani/FinTech/internal/model"
	"github.com/mr-meselmani/FinTech/internal/repository"
)

type AccountHandler struct {
	accountRepository repository.AccountRepository
	computationMutex  sync.Mutex
	computationWG     sync.WaitGroup
	accountsMutex     sync.RWMutex
}

func NewAccountHandler() *AccountHandler {
	// Initialize and return a new AccountHandler instance
	accountRepository := repository.NewAccountRepository()
	return &AccountHandler{
		accountRepository: *accountRepository,
	}
}

// HomeHandler handles the home route
func (ah *AccountHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Perform any necessary computations or tasks
	result := performComputation()

	// Acquire the mutex to ensure exclusive access when writing the response
	ah.computationMutex.Lock()
	defer ah.computationMutex.Unlock()

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, result)
}

// ListAccounts handles the list accounts route
func (ah *AccountHandler) ListAccounts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}

	// Acquire a read lock for the accounts data to allow concurrent access
	ah.accountsMutex.RLock()
	defer ah.accountsMutex.RUnlock()

	// Increment the wait group for each goroutine
	ah.computationWG.Add(1)

	// Launch a goroutine to encode and write the accounts data to the response
	go func() {
		defer ah.computationWG.Done()

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(ah.accountRepository.GetAccounts())
		if err != nil {
			fmt.Printf("Failed to encode accounts response: %v", err)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}()

	// Wait for all goroutines to complete
	ah.computationWG.Wait()
}

// MakeTransaction handles the transaction route
func (ah *AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}

	var req model.Transaction
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Checking the from account in data
	fromAccount := ah.accountRepository.GetAccountByID(req.From)
	if fromAccount == nil {
		http.Error(w, "Invalid 'from' account", http.StatusBadRequest)
		return
	}

	// Checking the to account in data
	toAccount := ah.accountRepository.GetAccountByID(req.To)
	if toAccount == nil {
		http.Error(w, "Invalid 'to' account", http.StatusBadRequest)
		return
	}

	// Getting the req amount
	amount := req.Amount

	// Converting fromAccount balance from string to float64 to match the amount type
	fromBalance, err := strconv.ParseFloat(fromAccount.Balance, 64)
	if err != nil {
		http.Error(w, "Invalid balance for 'from' account", http.StatusBadRequest)
		return
	}

	// Converting toAccount balance from string to float64 to match amount type
	toBalance, err := strconv.ParseFloat(toAccount.Balance, 64)
	if err != nil {
		http.Error(w, "Invalid balance for 'to' account", http.StatusBadRequest)
		return
	}

	// Checking the amount if less than or equal to zero
	if amount <= 0 {
		http.Error(w, "Invalid transaction amount", http.StatusBadRequest)
		return
	}

	// Checking the account balance if it's less than the amount transferred
	if fromBalance < amount {
		http.Error(w, "Insufficient funds", http.StatusBadRequest)
		return
	}

	// Create a WaitGroup to synchronize the completion of concurrent transactions
	var wg sync.WaitGroup
	wg.Add(2) // Two concurrent transactions: deduct from 'from' account and add to 'to' account

	// Execute the transactions concurrently using Goroutines
	go func() {
		defer wg.Done()
		fromAccount.Balance = strconv.FormatFloat(fromBalance-amount, 'f', 2, 64)
		ah.accountRepository.UpdateAccount(*fromAccount)
	}()

	go func() {
		defer wg.Done()
		toAccount.Balance = strconv.FormatFloat(toBalance+amount, 'f', 2, 64)
		ah.accountRepository.UpdateAccount(*toAccount)
	}()

	// Wait for both transactions to complete
	wg.Wait()

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Transaction successful",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Printf("Failed to encode response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// performComputation simulates a computationally intensive task
func performComputation() string {
	// Simulate the computation taking some time
	time.Sleep(2 * time.Second)

	return "Hello!"
}
