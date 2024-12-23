package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_"github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

// TestMain sets up the database connection and initializes `testQueries`.
// Logs a fatal error if the connection fails.

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
