package main

import (
	"log"
	"os"
	"path"

	conf "github.com/scutrobotlab/HR/conf"
	dal "github.com/scutrobotlab/HR/dal"
	"github.com/scutrobotlab/HR/dal/model"
	"gorm.io/gorm/clause"
)

func init() {
	dal.DB = dal.ConnectDB(conf.Postgres).Debug()
}

func main() {
	dal.DB.AutoMigrate(
		&model.Admin{},
		&model.Admit{},
		&model.Applicant{},
		model.Answer{},
		&model.Intent{},
		&model.OptionalTime{},
		&model.Question{},
		&model.Remark{},
		&model.Score{},
		&model.Setting{},
		&model.Standard{})

	var users = []model.Setting{
		getSetting("time-frame"),
		getSetting("announce"),
		getSetting("form"),
	}
	dal.DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(&users)
}

func getSetting(key string) model.Setting {
	dir, err := os.Getwd()
	if err != nil {
		log.Panicln(err.Error())
	}
	path := path.Join(dir, key+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		log.Panicln(err.Error())
	}
	return model.Setting{
		Key:   key,
		Value: string(b),
	}
}
