package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	//auto remove if we dont have _ underscore because we dont use the package directly
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:1234@localhost:5432/gobanking?sslmode=disable"
)

// entry testpoint for all golang tests
func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())

}
