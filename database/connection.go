package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connStr := "user=postgres password=123456 dbname=api-with-go sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Setting Pool Connections
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(0)

	// Testing Connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected in PostgreSQL...")

	return db
}
