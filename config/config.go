package configuration

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gopkg.in/ini.v1"
)

type Config struct {
	Database struct {
		Host          string `ini:"host"`
		Port          int    `ini:"port"`
		Username      string `ini:"username"`
		Password      string `ini:"password"`
		DatabasenName string `ini:"databasename"`
	} `ini:"database"`
	DatabaseOptions struct {
		SSLMode string `ini:"sslmode"`
	} `ini:"database.options"`

	ApiService struct {
		Port int `ini:"port"`
	} `ini:"api"`

	ApiServiceOptions struct {
		TokenTTL   int    `ini:"tokenttl"`
		JWTPrivKey string `ini:"jwt_priv_key"`
	} `ini:"api.options"`
}

func ReadConfig() Config {
	absPath, err := filepath.Abs("config.ini")
	if err != nil {
		panic(err)
	}
	inidata, err := ini.Load(absPath)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	var config Config

	err = inidata.MapTo(&config)
	if err != nil {
		fmt.Printf("Fail to map file: %v", err)
		os.Exit(1)
	}
	return config
}

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
