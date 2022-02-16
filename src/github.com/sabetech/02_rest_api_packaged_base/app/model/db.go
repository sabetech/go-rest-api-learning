package model

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var db *sql.DB

type dbConnection struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Init() {
	err := godotenv.Load("../config/.env")

	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err.Error())
		return
	}

	connectionInfo := dbConnection{
		Host:     os.Getenv("DB_CONNECTION"),
		Port:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_Password"),
		Database: os.Getenv("DB_DATABASE"),
	}

	//lets open an mysql connection
	db, err = sql.Open("mysql", connToString(connectionInfo))
	if err != nil {
		fmt.Printf("Error connecting to the DB: %s\n", err.Error())
		return
	} else {
		fmt.Printf("DB is open\n")
	}
}

func connToString(info dbConnection) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.Database)

}
