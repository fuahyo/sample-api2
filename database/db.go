package database

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func ConnectDB() (*sqlx.DB, error) {
	// Replace with your SQL Server database credentials
	server := "CSIJKTSAPP06"
	port := 1433
	user := "webapps"
	password := "supernova"
	dbName := "Reserva2.0"

	// Create the SQL Server connection string
	connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, dbName)

	// Connect to the SQL Server database
	db, err := sqlx.Connect("mssql", connStr)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
