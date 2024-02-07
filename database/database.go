package database

import (
	configuration "api/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect(conf configuration.Config) {
	var err error

	// host := os.Getenv("DB_HOST")
	// username := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// databaseName := os.Getenv("DB_NAME")
	// port := os.Getenv("DB_PORT")

	host := conf.Database.Host
	username := conf.Database.Username
	password := conf.Database.Password
	databaseName := conf.Database.DatabasenName
	port := conf.Database.Port
	sslmode := conf.DatabaseOptions.SSLMode

	dsn := fmt.Sprintf("host=%s  user=%s  password =%s dbname=%s port =%d sslmode=%s", host, username, password, databaseName, port, sslmode)

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}
