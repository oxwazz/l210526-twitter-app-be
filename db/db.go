package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/oxwazz/l210526-twitter-app-be/helpers"
)

var db *sql.DB
var err error

func Init() {

	port := helpers.GetEnv("DB_PORT")

	//conf := config.GetConfig()

	// connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"

	//connectionString := "postgres://" + conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@" + conf.DB_HOST + ":" + conf.DB_PORT + "/" + conf.DB_NAME + "?sslmode=disable"
	connectionString := "postgres://" + "uzdkdicohqwhzh" + ":" + "8b25c1f23ca21c0a826597ca2a0c33c3b7c9215d369b2dd1278c52161a7c3668" + "@" + "ec2-54-204-56-171.compute-1.amazonaws.com" + ":" + port + "/" + "daj6o8hv12hgm8"

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
