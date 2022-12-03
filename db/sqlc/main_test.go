package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	// postgres driver for Go's database/sql package
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/apple_store?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
