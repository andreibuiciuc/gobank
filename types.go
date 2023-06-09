package main

import (
	"math/rand"
	"net/http"
	"time"
)

// Entities and constructor functions

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

// Constructor function: function that returns a pointer to an instantiated component
// In this case, the function constructor will return a pointer to a new instance of the Account component
func NewAccount(firstName, lastname string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastname,
		Number:    int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}
}

type APIServer struct {
	listenAddr string
	store      Storage
}

// Constructor function: function that returns a pointer to an instantiated component
// In this case, the function constructor will return a pointer to a new instance of the APIServer component
func NewAPIServer(listenAddress string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddress,
		store:      store,
	}
}

type APIFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string `json:"error"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
