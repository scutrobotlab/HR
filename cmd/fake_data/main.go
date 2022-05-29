package main

import (
	"log"
	"os"
	"path"

	"github.com/scutrobotlab/HR/cmd/fake_data/faker"
	"github.com/scutrobotlab/HR/conf"
	"github.com/scutrobotlab/HR/dal"
	"github.com/scutrobotlab/HR/dal/model"
	"github.com/scutrobotlab/HR/dal/query"
	"gorm.io/gorm/clause"
)

func init() {
	dal.DB = dal.ConnectDB(conf.Postgres).Debug()
	query.SetDefault(dal.DB)
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Panicln(err.Error())
	}
	p := path.Join(dir, "data.json")
	f, err := faker.GetFaker(p)
	if err != nil {
		panic(err)
	}
	var (
		admin_cnt    uint = 50
		app_cnt      uint = 200
		question_cnt uint = 80
	)
	dal.DB.Migrator().DropTable(
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
	f.GenAdmin(admin_cnt)
	f.GenApplicant(app_cnt)
	f.GenAdmit(admin_cnt, app_cnt)
	f.GenQuestion(question_cnt)
	f.GenAnswer(app_cnt, question_cnt)
	f.GenOptionalTime()
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
