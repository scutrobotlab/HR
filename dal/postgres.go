package dal

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dsn string) (db *gorm.DB) {
	var err error

	db, err = gorm.Open(postgres.Open(dsn))
	fmt.Printf("connect to db: %s", dsn)

	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}

	return db
}
