package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

// Test Transaction
func TestMakeTransactionConcurrent(t *testing.T) {
	// Number of concurrent requests to simulate
	numRequests := 100

	// Payload for the transaction request
	payload := []byte(`{
		"from": "fd796d75-1bcf-4a95-bf1a-f7b296adb79f",
		"to": "3d253e29-8785-464f-8fa0-9e4b57699db9",
		"amount": 10
	}`)

	// Send concurrent requests
	for i := 0; i < numRequests; i++ {
		go func() {
			// Create a new request
			req, err := http.NewRequest("POST", "http://localhost:3000/transfer", bytes.NewBuffer(payload))
			if err != nil {
				t.Errorf("Failed to create request: %v", err)
				return
			}

			// Send the request
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("Failed to send request: %v", err)
				return
			}
			defer resp.Body.Close()

			// Verify the response status code
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)

				// Read and print the response body
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					t.Errorf("Failed to read response body: %v", err)
					return
				}
				fmt.Printf("Response Body: %s\n", body)
				return
			}

			// Read the response body
			var response struct {
				Message string `json:"message"`
			}
			err = json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				t.Errorf("Failed to decode response: %v", err)
				return
			}

			// Print the response message to the console
			fmt.Printf("Response Message: %s\n", response.Message)
		}()
	}

	// Wait for all requests to complete
	time.Sleep(2 * time.Second) // Add a delay to ensure all requests are processed
}
