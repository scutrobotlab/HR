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
		OutPath: "../../dal/query",
	})

	g.UseDB(dal.DB)

	// generate all table from database
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
