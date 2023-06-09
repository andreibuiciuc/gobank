package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandlerFunc(server.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(server.handleAccountById))

	log.Println("Server listening on port: ", server.listenAddr)

	http.ListenAndServe(server.listenAddr, router)
}

// Handles the request in regards to the account functionalities
func (server *APIServer) handleAccount(writer http.ResponseWriter, request *http.Request) error {
	switch {
	case request.Method == "GET":
		return server.handleGetAccount(writer, request)
	case request.Method == "POST":
		return server.handleCreateAccount(writer, request)
	}

	return fmt.Errorf("method provided not allowed: %s", request.Method)
}

func (server *APIServer) handleAccountById(writer http.ResponseWriter, request *http.Request) error {
	switch {
	case request.Method == "GET":
		return server.handleGetAccountByID(writer, request)
	case request.Method == "DELETE":
		return server.handleDeleteAccount(writer, request)
	}

	return fmt.Errorf("method provided not allowed: %s", request.Method)
}

// Handles the HTTP GET request for retrieving accounts
func (server *APIServer) handleGetAccount(writer http.ResponseWriter, request *http.Request) error {
	accounts, err := server.store.GetAccounts()

	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, accounts)
}

// Handles the HTTP GET request for retrieving an account by its ID
func (server *APIServer) handleGetAccountByID(writer http.ResponseWriter, request *http.Request) error {
	id, err := getIDFromRequest(request)
	if err != nil {
		return err
	}

	account, err := server.store.GetAccountByID(id)
	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, account)
}

// Handles the HTTP POST request for creating a new account
func (server *APIServer) handleCreateAccount(writer http.ResponseWriter, request *http.Request) error {
	accountRequest := new(CreateAccountRequest)

	if err := json.NewDecoder(request.Body).Decode(accountRequest); err != nil {
		return err
	}

	account := NewAccount(accountRequest.FirstName, accountRequest.LastName)
	_, err := server.store.CreateAccount(account)

	if err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, account)
}

// Handles the HTTP DELETE request for deleting an existing account
func (server *APIServer) handleDeleteAccount(writer http.ResponseWriter, request *http.Request) error {
	id, err := getIDFromRequest(request)

	if err != nil {
		return err
	}

	if err := server.store.DeleteAccount(id); err != nil {
		return err
	}

	return writeJSON(writer, http.StatusOK, map[string]int{"deleted": id})
}

// Converts a function to a HTTP handler function
func makeHTTPHandlerFunc(f APIFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := f(writer, request); err != nil {
			writeJSON(writer, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

// Constructs a HTTP response
// Check the interface of the Response Writer for further information
func writeJSON(writer http.ResponseWriter, status int, v any) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)
	return json.NewEncoder(writer).Encode(v)
}

func getIDFromRequest(request *http.Request) (int, error) {
	idString := mux.Vars(request)["id"]
	id, err := strconv.Atoi(idString)

	if err != nil {
		return id, fmt.Errorf("invalid ID given %s", idString)
	}

	return id, nil
}
