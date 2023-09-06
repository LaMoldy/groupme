package database

import (
    "database/sql"
    "errors"
    "fmt"
    "os"

    "github.com/joho/godotenv"
)

func getDatabaseProperties() (string, string, string, error) {
    err := godotenv.Load()

    if err != nil {
	return "", "", "", errors.New("[Error]: Failed to load environment file")
    }

    name := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    databaseName := os.Getenv("DB_DATABASE")

    return name, password, databaseName, nil
}

func ConnectDatabase() *sql.DB {
    user, password, databaseName, err := getDatabaseProperties()

    if err != nil {
	fmt.Println(err.Error())
	os.Exit(1)
    }

    dbinfo := fmt.Sprintf("user=%s password=%s dbname=% sslmode=disable", user, password, databaseName)
    db, err := sql.Open("postgres", dbinfo)

    if err != nil {
	fmt.Println("[Error]: Failed to connect to the database")
	os.Exit(1)
    }

    return db
}
