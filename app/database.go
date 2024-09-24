package app

import (
	"database/sql"
	"fmt"
	"os"
	"simple-restaurant-web/helper"

	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {
	err := godotenv.Load()
	helper.PanicIfError(err)

	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s connect_timeout=600 sslmode=disable", user, dbname, password, host, port)
	db, err := sql.Open("postgres", connStr)
	helper.PanicIfError(err)

	return db
}
