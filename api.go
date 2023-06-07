package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
}

type APIFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}

// Creates a new API server with the given listener address
func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddr: listenAddress,
	}
}

func makeHTTPHandlerFunc(f APIFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := f(writer, request); err != nil {
			writeJSON(writer, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func writeJSON(writer http.ResponseWriter, status int, v any) error {
	writer.WriteHeader(status)
	writer.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(writer).Encode(v)
}

func (server *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandlerFunc(server.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(server.handleGetAccount))

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
	case request.Method == "DELETE":
		return server.handleDeleteAccount(writer, request)
	}

	return fmt.Errorf("Method provided not allowed: %s", request.Method)
}

// Handles the HTTP GET request for retrieving an account
func (server *APIServer) handleGetAccount(writer http.ResponseWriter, request *http.Request) error {
	account := NewAccount("Andrei", "Buiciuc")

	return writeJSON(writer, http.StatusOK, account)
}

// Handles the HTTP POST request for creating a new account
func (server *APIServer) handleCreateAccount(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

// Handles the HTTP DELETE request for deleting an existing account
func (server *APIServer) handleDeleteAccount(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
