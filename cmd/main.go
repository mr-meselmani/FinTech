package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mr-meselmani/FinTech/internal/handler"
	"github.com/mr-meselmani/FinTech/pkg/banner"
	"github.com/mr-meselmani/FinTech/pkg/router"
)

func main() {
	// Write Banner
	banner.WriteBanner()

	// Setup Router
	r := mux.NewRouter().StrictSlash(true)

	// Initialize the account handler
	accountHandler := handler.NewAccountHandler()

	// Define the routes
	r.HandleFunc("/", accountHandler.HomeHandler).Methods("GET")
	r.HandleFunc("/accounts", accountHandler.ListAccounts).Methods("GET")
	r.HandleFunc("/transfer", accountHandler.MakeTransaction).Methods("POST")

	// Use the custom router with middleware
	router := router.NewRouter(r)

	// Start the server
	log.Fatal(http.ListenAndServe(":3000", router))
}
