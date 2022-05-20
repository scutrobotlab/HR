package main

import (
	conf "github.com/scutrobotlab/HR/conf"
	dal "github.com/scutrobotlab/HR/dal"
	"gorm.io/gen"
)

func init() {
	dal.DB = dal.ConnectDB(conf.Postgres).Debug()

	prepare(dal.DB) // prepare table for generate
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "/tmp/gentest/query",
	})

	g.UseDB(dal.DB)

	g.GenerateAllTable()

	g.Execute()
}
