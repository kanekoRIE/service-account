package config

import (
	"database/sql"
	"fmt"
	"service-account/logger"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func getDBConfig() string {
	host := DB_HOST
	port := DB_PORT
	user := DB_USER
	password := DB_PASSWORD
	dbname := DB_NAME

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return connStr
}

func InitDB() (*sql.DB, error) {
	connStr := getDBConfig()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.LogCritical("Failed to open database", logrus.Fields{"error": err})
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.LogCritical("Failed to connect to database", logrus.Fields{"error": err})
		return nil, err
	}

	logger.LogInfo("Successfully connected to the database", logrus.Fields{})
	return db, nil
}
