package utils

import (
	"database/sql"
	"fmt"

	"github.com/JWindy92/wanna-go-api/internal/logging"
)

var logger = logging.GetLogger()

func Connect() {
	dbUser := "john"
	dbPass := "12345"
	dbName := "wannago"
	dbHost := "localhost"
	dbPort := "3306"

	// Create a DSN (Data Source Name) string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Sugar().Errorf("Error opening database connection: %v\n", err)
		return
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		logger.Sugar().Errorf("Error connecting to database: %v\n", err)
		return
	}

	logger.Info("Connected to MySQL database!")
}
