package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
    _ "github.com/lib/pq"
)

// the database driver and the source url from the docker in the makefile
const (
    driverName = "postgres"
    dataSourceName = "postgresql://postgres:password@localhost:5433/user_golang?sslmode=disable"
)

// testQueries is a pointer to *Queries struct in the ./db/sqlc/db.go
var testQueries *Queries

func TestMain(m *testing.M) {
    conn, err := sql.Open(driverName, dataSourceName)
    if err != nil {
        log.Fatal("Cannot Connect to the Database because ", err)
    }

    testQueries = New(conn)
    os.Exit(m.Run())
}
