package main

import (
	"math/rand"
	"net/http"
)

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstName, lastname string) *Account {
	return &Account{
		ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastname,
		Number:    int64(rand.Intn(1000000)),
	}
}

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddr: listenAddress,
	}
}

type APIFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}
