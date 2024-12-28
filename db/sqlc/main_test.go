package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/iam-benjamen/simple_bank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

// TestMain sets up the database connection and initializes `testQueries`.
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
