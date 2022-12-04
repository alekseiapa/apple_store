package main

import (
	"database/sql"
	"log"

	"github.com/alekseiapa/apple_store/api"
	db "github.com/alekseiapa/apple_store/db/sqlc"

	// DONT remove! postgres driver for Go's database/sql package
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/apple_store?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server")
	}
}
