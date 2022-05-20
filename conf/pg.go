package conf

import (
	"fmt"
	"log"
	"os"
)

var Postgres string

func init() {
	Postgres = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	log.Println(Postgres)
}
