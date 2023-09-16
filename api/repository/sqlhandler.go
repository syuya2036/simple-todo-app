package repository

import (
	"fmt"
	"os"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlHandler struct {
	*gorm.DB
}

func NewSqlHandler() SqlHandler {
	dbEnv := getDbEnv()
	db, err := gorm.Open(postgres.Open(dbEnv), &gorm.Config{})
	if err != nil {
		msg := fmt.Sprintf("DB connection error: %s", err)
		panic(msg)
	}

	sqlHandler := SqlHandler{db}
	return sqlHandler
}

func getDbEnv() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	sslmode := os.Getenv("DB_SSLMODE")

	// pingを打って確かめる
	

	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)
}
