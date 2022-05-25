package db

import (
	"database/sql"

	"github.com/oxwazz/l210526-twitter-app-be/config"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	// connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"

	connectionString := "postgres://" + conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@" + conf.DB_HOST + ":" + conf.DB_PORT + "/" + conf.DB_NAME + "?sslmode=disable"

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic("connectionString error..")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
		// panic("DSN Invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
