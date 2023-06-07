package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) (*Account, error)
	DeleteAccount(int) error
	UpdateAccount(*Account) (*Account, error)
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

func (store *PostgresStore) Init(isAccDropEnable bool) error {
	var err error = nil

	if isAccDropEnable {
		err = store.dropAccountTable()
	}

	if err != nil {
		return err
	}

	return store.createAccountTable()
}

func (store *PostgresStore) createAccountTable() error {
	query :=
		`create table if not exists account (
			id serial primary key,
			first_name varchar(50),
			last_name varchar(50),
			number serial,
			balance bigint,
			created_at timestamp
		)`

	_, err := store.db.Exec(query)
	return err
}

func (store *PostgresStore) dropAccountTable() error {
	query := "drop table if exists account"

	_, err := store.db.Exec(query)
	return err
}

func (store *PostgresStore) CreateAccount(account *Account) (*Account, error) {
	query := ` insert into account (first_name, last_name, number, balance, created_at) values($1, $2, $3, $4, $5)`

	resp, err := store.db.Query(
		query,
		account.FirstName,
		account.LastName,
		account.Number,
		account.Balance,
		account.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", resp)

	return nil, nil
}

func (store *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (store *PostgresStore) UpdateAccount(account *Account) (*Account, error) {
	return nil, nil
}

func (store *PostgresStore) GetAccounts() ([]*Account, error) {
	query := "select * from account"

	rows, err := store.db.Query(query)

	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (store *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
