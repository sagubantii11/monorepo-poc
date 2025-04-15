package utils

import (
	"database/sql"
	"fmt"
	"go-basic-poc/internal/config"
	"io"
	"os"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var DB *sql.DB

func InitDB(cfg *config.Config) error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	DB = db
	Logger.Info("Successfully connected to database")
	return nil
}

func ExecDBScript(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening SQL script: %w", err)
	}
	defer file.Close()

	script, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading SQL script: %w", err)
	}

	_, err = DB.Exec(string(script))
	if err != nil {
		return fmt.Errorf("error executing SQL script: %w", err)
	}
	Logger.Info("Successfully executed SQL script", zap.String("filename", filename))
	return nil
}

func QueryRowDB(query string, args ...interface{}) *sql.Row {
	return DB.QueryRow(query, args...)
}

func QueryDB(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error performing database query: %w", err)
	}
	return rows, nil
}

func ExecDB(query string, args ...interface{}) (sql.Result, error) {
	result, err := DB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing SQL statement: %w", err)
	}
	return result, nil
}
