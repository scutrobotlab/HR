package faker

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/scutrobotlab/HR/dal/query"
)

type Faker struct {
	FirstName  []string
	SecondName []string
	SingleName []string
	Major      []string
	Class      []string
	Group      []string
}

var q = query.Q
var ctx = context.Background()

// Read and parse json file
func GetFaker(path string) (*Faker, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var d Faker
	if err := json.Unmarshal(file, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (f *Faker) getAdminName() string {
	first := rand.Int() % len(f.FirstName)
	second := rand.Int() % len(f.SecondName)
	return f.FirstName[first] + f.SecondName[second]
}

func (f *Faker) getApplicantName() string {
	first := rand.Int() % len(f.FirstName)
	second := rand.Int() % len(f.SingleName)
	if rand.Int()%3 == 0 {
		return f.FirstName[first] + f.SingleName[second]
	}
	third := rand.Int() % len(f.SingleName)
	return f.FirstName[first] + f.SingleName[second] + f.SingleName[third]
}

func (f *Faker) getClassAndSchool() (string, string) {
	major := f.Major[rand.Int()%len(f.Major)]
	return major + f.Class[rand.Int()%len(f.Class)], major + "学院"
}

func (f *Faker) getPhone() string {
	return fmt.Sprintf("13%d%d%d%d%d%d%d%d%d%d",
		rand.Int()%10, rand.Int()%10, rand.Int()%10, rand.Int()%10,
		rand.Int()%10, rand.Int()%10, rand.Int()%10, rand.Int()%10,
		rand.Int()%10, rand.Int()%10)
}

// get md5 from uint value
func (f *Faker) getMD5(i uint) string {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%d", i)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
