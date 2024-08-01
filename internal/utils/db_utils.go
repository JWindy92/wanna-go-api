package utils

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/JWindy92/wanna-go-api/internal/logging"
	"github.com/JWindy92/wanna-go-api/internal/models"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// TODO: move to models
type User struct {
	Username string `json:"username"`
}

var logger = logging.GetLogger()
var db *sql.DB

func Connect() (*sql.DB, error) {
	dbUser := "john"
	dbPass := "12345"
	dbName := "wannago"
	dbHost := "localhost"
	dbPort := "3306"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		logger.Sugar().Errorf("Error opening database connection: %v\n", err)
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		logger.Sugar().Errorf("Error connecting to database: %v\n", err)
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	err = runMigrations(db)
	if err != nil {
		logger.Sugar().Errorf("Error running migrations: %v\n", err)
		return nil, fmt.Errorf("error running migrations: %v", err)
	}

	logger.Info("Connected to MySQL database!")
	return db, nil
}

func CreateUser(r models.RegistrationRequest) error {
	user := models.User{
		Username: r.Username,
	}

	query := "INSERT INTO users (username) VALUES (?)"
	result, err := db.Exec(query, user.Username)
	if err != nil {
		return fmt.Errorf("error inserting data: %v", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting the last insert ID: %v", err)
	}
	logger.Sugar().Infof("Inserted user with ID: %d\n", lastInsertID)
	return nil
}

func runMigrations(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}
	logger.Sugar().Infof("Driver: %d\n", driver)
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"wannago", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
