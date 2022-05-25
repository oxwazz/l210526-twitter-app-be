package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var db2 *gorm.DB
var err2 error

func Init2() {
	dsn := "host=localhost user=postgres password=postgres dbname=twitter port=5432 sslmode=disable"
	db2, err2 = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err2 != nil {
		panic("connectionerror..")
	}
}

func CreateCon2() *gorm.DB {
	return db2
}
