package main

import (
	conf "github.com/scutrobotlab/HR/conf"
	dal "github.com/scutrobotlab/HR/dal"
	"github.com/scutrobotlab/HR/dal/model"
	"gorm.io/gen"
)

func init() {
	dal.DB = dal.ConnectDB(conf.Postgres).Debug()
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithDefaultQuery,

		WithUnitTest: true,

		FieldNullable:     true,
		FieldCoverable:    true,
		FieldWithIndexTag: true,
	})

	g.UseDB(dal.DB)

	g.ApplyBasic(
		model.Admin{},
		model.Admit{},
		model.Applicant{},
		model.Group{},
		model.Intent{},
		model.OptionalTime{},
		model.Question{},
		model.Remark{},
		model.Score{},
		model.Setting{},
		model.Standard{})

	g.Execute()
}
