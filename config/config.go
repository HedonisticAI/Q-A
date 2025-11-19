package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBName   string
	DBPort   string
	DBHost   string
	DBPwd    string
	DBUser   string
	HttpPort string
}

func NewCondig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	DBHost, exist := os.LookupEnv("DB_HOST")
	if !exist {
		return nil
	}
	DBPort, exist := os.LookupEnv("DB_PORT")
	if !exist {
		return nil
	}
	DBUser, exist := os.LookupEnv("DB_USER")
	if !exist {

		return nil
	}
	DBPwd, exist := os.LookupEnv("DB_PWD")
	if !exist {
		return nil
	}
	DBName, exist := os.LookupEnv("DB_NAME")
	if !exist {
		return nil
	}
	HttpPort, exist := os.LookupEnv("HTTP_PORT")
	if !exist {
		return nil
	}
	return &Config{DBName: DBName, DBPort: DBPort, DBHost: DBHost, DBPwd: DBPwd, DBUser: DBUser, HttpPort: HttpPort}
}
