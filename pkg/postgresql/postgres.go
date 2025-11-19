package postgresql

import (
	"fmt"
	"golangqatestdesu/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func NewDB(Config config.Config) (*DB, error) {
	connstr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", Config.DBHost, Config.DBUser, Config.DBPwd, Config.DBName, Config.DBPort)
	db, err := gorm.Open(postgres.Open(connstr), &gorm.Config{})
	return &DB{DB: db}, err
}
