package tests

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

// Test List Accounts
func TestListAccountsConcurrent(t *testing.T) {
	// Number of concurrent requests to simulate
	numRequests := 100

	// Running server URL
	serverURL := "http://localhost:3000/accounts" // Replace with the actual URL of your running server

	// Send concurrent requests
	for i := 0; i < numRequests; i++ {
		// Create a new WaitGroup for each request
		var wg sync.WaitGroup
		wg.Add(1)

		go func(requestNum int) {
			// Decrement the wait group counter when the goroutine completes
			defer wg.Done()

			// Send a GET request to the server
			resp, err := http.Get(serverURL)
			if err != nil {
				t.Errorf("Failed to send GET request: %v", err)
				return
			}
			defer resp.Body.Close()

			// Print the request status to the console
			fmt.Printf("Request %d: Status: %s\n", requestNum+1, resp.Status)

			// Verify the response status code
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
			}
		}(i)

		// Wait for the current request to complete before starting the next iteration
		wg.Wait()
	}
}
