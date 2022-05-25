package databases

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db2 *gorm.DB
var err2 error

func Init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	dsn := "host=ec2-54-204-56-171.compute-1.amazonaws.com user=uzdkdicohqwhzh password=8b25c1f23ca21c0a826597ca2a0c33c3b7c9215d369b2dd1278c52161a7c3668 dbname=daj6o8hv12hgm8 port=5432"
	//dsn := "host=localhost user=postgres password=postgres dbname=twitter port=5432 sslmode=disable"
	db2, err2 = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err2 != nil {
		panic("connectionerror..")
	}
}

func CreateConnection() *gorm.DB {
	return db2
}
