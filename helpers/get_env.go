package helpers

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func GetEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	// Set the file name of the configurations file
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		fmt.Println("error ges")
	}

	if len(value) == 0 {
		return key
	}
	return value
}
