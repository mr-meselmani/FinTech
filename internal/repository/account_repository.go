package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/mr-meselmani/FinTech/internal/model"
)

type AccountRepository struct {
	accountsData  model.Accounts
	accountsMutex sync.RWMutex
}

func NewAccountRepository() *AccountRepository {
	// Initialize and return a new AccountRepository instance
	// Parse the accounts.json file and populate the accountsData
	filePath := "./data/accounts.json"
	accountsData, err := ParseAccounts(filePath)
	if err != nil {
		// Handle the error appropriately, e.g., log an error and exit the application
		panic(err)
	}

	return &AccountRepository{
		accountsData: accountsData,
	}
}

func (ar *AccountRepository) GetAccounts() model.Accounts {
	// Acquire a read lock for the accounts data to allow concurrent access
	ar.accountsMutex.RLock()
	defer ar.accountsMutex.RUnlock()

	return ar.accountsData
}

func (ar *AccountRepository) GetAccountByID(id string) *model.Account {
	// Acquire a read lock for the accounts data to allow concurrent access
	ar.accountsMutex.RLock()
	defer ar.accountsMutex.RUnlock()

	for _, account := range ar.accountsData {
		if account.ID == id {
			return &account
		}
	}
	return nil
}

func (ar *AccountRepository) UpdateAccount(account model.Account) {
	// Acquire a write lock for the accounts data to ensure exclusive access
	ar.accountsMutex.Lock()
	defer ar.accountsMutex.Unlock()

	for i, acc := range ar.accountsData {
		if acc.ID == account.ID {
			ar.accountsData[i].Balance = account.Balance
			return
		}
	}
}

func ParseAccounts(filePath string) (model.Accounts, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return model.Accounts{}, err
	}

	defer file.Close()

	var accountsData model.Accounts

	err = json.NewDecoder(file).Decode(&accountsData)
	if err != nil {
		return model.Accounts{}, err
	}

	fmt.Printf("Data ingested successfully. Total accounts: %d", len(accountsData))

	return accountsData, nil
}
