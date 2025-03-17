package db

import (
	"fmt"
	"kssyncservice_go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *config.Config) *Db {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.DbConfig.DB_HOST, conf.DbConfig.DB_USER, conf.DbConfig.DB_PASSWORD, conf.DbConfig.DB_DATABASE, conf.DbConfig.DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	defer fmt.Println("Database connect")

	return &Db{db}
}