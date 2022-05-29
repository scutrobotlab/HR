package faker

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/scutrobotlab/HR/dal/model"
)

func (f *Faker) GenAdmin(cnt uint) {
	admins := make([]*model.Admin, cnt)
	for i := uint(0); i < cnt; i++ {
		admins[i] = &model.Admin{
			Name: f.getAdminName(),
		}
	}
	err := q.Admin.WithContext(ctx).Create(admins...)
	if err != nil {
		panic(err)
	}
}

func (f *Faker) GenApplicant(cnt uint) {
	applicants := make([]*model.Applicant, cnt)
	for i := uint(0); i < cnt; i++ {
		gender := rand.Int()%2 == 0
		phone := f.getPhone()
		class, school := f.getClassAndSchool()

		applicants[i] = &model.Applicant{
			Name:   f.getApplicantName(),
			OpenID: f.getMD5(i),
			Gender: &gender,
			Phone:  &phone,
			Form: map[string]string{
				"class":  class,
				"school": school},
		}
	}
	err := q.Applicant.WithContext(ctx).Create(applicants...)
	if err != nil {
		panic(err)
	}
}

func (f *Faker) GenAdmit(admin_cnt uint, app_cnt uint) {
	admits := make([]*model.Admit, 0, app_cnt/2)

	applicants, err := q.Applicant.WithContext(ctx).Select(q.Applicant.ID).Find()
	if err != nil {
		panic(err)
	}

	admins, err := q.Admin.WithContext(ctx).Select(q.Admin.ID).Find()
	if err != nil {
		panic(err)
	}

	for _, app := range applicants {
		if rand.Int()%2 == 0 {
			continue
		}
		admit := &model.Admit{
			AdminID:     admins[rand.Int()%len(admins)].ID,
			ApplicantID: app.ID,
			Group:       f.Group[rand.Int()%len(f.Group)],
		}
		admits = append(admits, admit)
	}
	err = q.Admit.WithContext(ctx).Create(admits...)
	if err != nil {
		panic(err)
	}
}

func (f *Faker) GenQuestion(q_cnt uint) {
	questions := make([]*model.Question, q_cnt)
	for i := 0; i < int(q_cnt); i++ {
		questions[i] = &model.Question{
			TheQuestion: "了解" + f.getApplicantName() + "吗？",
			Group:       f.Group[rand.Int()%len(f.Group)],
		}
	}
	err := q.Question.WithContext(ctx).Create(questions...)
	if err != nil {
		panic(err)
	}
}

func (f *Faker) GenAnswer(app_cnt uint, q_cnt uint) {
	answers := make([]*model.Answer, 0, app_cnt*q_cnt/2)

	applicants, err := q.Applicant.WithContext(ctx).Select(q.Applicant.ID).Find()
	if err != nil {
		panic(err)
	}

	questions, err := q.Question.WithContext(ctx).Select(q.Question.ID).Find()
	if err != nil {
		panic(err)
	}

	for _, app := range applicants {
		for _, q := range questions {
			if rand.Int()%2 == 0 {
				continue
			}
			answer := &model.Answer{
				ApplicantID: app.ID,
				QuestionID:  q.ID,
				TheAnswer:   rand.Int()%2 == 0,
			}
			answers = append(answers, answer)
		}
	}
	err = q.Answer.WithContext(ctx).Create(answers...)
	if err != nil {
		panic(err)
	}
}

func (f *Faker) GenOptionalTime() int {
	times := make([]*model.OptionalTime, 0, 3*(11-9)*60/15*len(f.Group)*2)
	now := time.Now()
	for d := 0; d < 3; d++ {
		for h := 9; h < 11; h++ {
			for m := 0; m < 60; m += 15 {
				datetime := time.Date(now.Year(), now.Month(), now.Day()+d, h, m, 0, 0, time.Local)
				for g := 0; g < len(f.Group); g++ {
					for r := uint(1); r <= 2; r++ {
						time := &model.OptionalTime{
							TheDate:     datetime,
							TheTime:     datetime,
							TheLocation: fmt.Sprintf("33号楼%d0%d", rand.Int()%9+1, rand.Int()%9+1),
							Group:       f.Group[g],
							IntentRank:  r,
							TotalCount:  2,
						}
						times = append(times, time)
					}
				}
			}
		}
	}
	err := q.OptionalTime.WithContext(ctx).Create(times...)
	if err != nil {
		panic(err)
	}
	return len(times)
}
