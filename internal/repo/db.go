package repo

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// InitializeDB sets up the PostgreSQL connection
func InitializeDB() (*sql.DB, error) {
	// Connection string to connect to the PostgreSQL database
	connStr := "host=localhost port=5432 user=myuser password=mypassword dbname=master-financial-db sslmode=disable"

	// Open a connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to open DB connection: %v", err)
	}

	// Verify the connection is working
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to ping DB: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
	return db, nil
}
