package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/LaMoldy/groupme/server/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDatabaseProperties() (string, string, string, string, error) {
    err := godotenv.Load()

    if err != nil {
	return "", "", "", "", errors.New("[Error]: Failed to load environment file")
    }

    name := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    databaseName := os.Getenv("DB_DATABASE")
    port := os.Getenv("DB_PORT")

    return name, password, databaseName, port, nil
}

func ConnectDatabase() *gorm.DB {
    user, password, databaseName, port, err := getDatabaseProperties()

    if err != nil {
	fmt.Println(err.Error())
	os.Exit(1)
    }

    dbinfo := fmt.Sprintf(
	"host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
	port,
	user,
	password,
	databaseName,
    )

    db, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})

    if err != nil {
	fmt.Println("[Error]: Failed to connect to the database")
	os.Exit(1)
    }

    db.AutoMigrate(models.User{})

    return db
}
