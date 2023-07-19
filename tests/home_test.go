package tests

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

// Test Home Requests
func TestHomeHandlerConcurrent(t *testing.T) {
	// Number of concurrent requests to simulate
	numRequests := 100

	// Running server URL
	serverURL := "http://localhost:3000" // Replace with the actual URL of your running server

	// Create a wait group to synchronize the completion of all requests
	var wg sync.WaitGroup
	wg.Add(numRequests)

	// Send concurrent requests
	for i := 0; i < numRequests; i++ {
		go func(i int) {
			defer wg.Done()

			// Send a GET request to the server
			resp, err := http.Get(serverURL)
			if err != nil {
				t.Errorf("Failed to send GET request: %v", err)
				return
			}
			defer resp.Body.Close()

			// Print the request URL and status
			fmt.Printf("Request %d: %s - Status: %s\n", i+1, serverURL, resp.Status)

			// Verify the response status code
			if resp.StatusCode != http.StatusAccepted {
				t.Errorf("Expected status code %d, but got %d", http.StatusAccepted, resp.StatusCode)
			}
		}(i)
	}

	// Wait for all requests to complete
	wg.Wait()
}
