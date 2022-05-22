package main

import (
	"fmt"
	"log"

	"github.com/scutrobotlab/HR/biz"
	"github.com/scutrobotlab/HR/conf"
	"github.com/scutrobotlab/HR/dal"
	"github.com/scutrobotlab/HR/dal/query"
)

func init() {
	dal.DB = dal.ConnectDB(conf.Postgres).Debug()
}

func main() {
	log.SetFlags(log.Lshortfile)
	// start your project here
	fmt.Println("hello world")
	defer fmt.Println("bye~")

	query.SetDefault(dal.DB)
	biz.Run()
}
