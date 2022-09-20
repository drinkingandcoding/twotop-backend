package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("%s not set\n", key)
	} else {
		return val, nil
	}
}
func Connect() (*gorm.DB, error) {
	dsn, err := getEnv("INTERNAL_URL")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, err
}
