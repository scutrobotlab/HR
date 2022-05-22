package ctrl

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type timeframe struct {
	FormStart string `json:"form-start" binding:"required"`
	FormEnd   string `json:"form-end" binding:"required"`
	TimeStart string `json:"time-start" binding:"required"`
	TimeEnd   string `json:"time-end" binding:"required"`
	Done      string `json:"done" binding:"required"`
}

type field struct {
	Name     string   `json:"name" binding:"required"`
	Key      string   `json:"key" binding:"required"`
	Type     string   `json:"type" binding:"required"`
	Required bool     `json:"required" binding:"required"`
	Options  []string `json:"option"`
	Regexp   string   `json:"regexp" binding:"required"`
}

type intent struct {
	Min      uint `json:"min" binding:"required"`
	Max      uint `json:"max" binding:"required"`
	Parallel bool `json:"parallel" binding:"required"`
}

type form struct {
	Intent intent  `json:"intent" binding:"required"`
	Fields []field `json:"fields" binding:"required"`
}

type announce struct {
	HaventAppliedForm  string
	HaventSelectedTime string
	HaventInterview    string
	Interviewed        string
	Admitted           string
	Failed             string
}

func UpdateSetting(c *gin.Context) {
	key, ok := c.Params.Get("key")
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}
	s := q.Setting
	do := s.WithContext(ctx).Where(s.Key.Eq(key))

	var err error
	switch key {
	case "form":
		var f form
		if err := c.BindJSON(&f); err != nil {
			log.Println(err.Error())
			c.Status(http.StatusBadRequest)
			return
		}
		_, err = do.UpdateColumn(s.Value, f)
	case "announce":
		var a announce
		if err := c.BindJSON(&a); err != nil {
			log.Println(err.Error())
			c.Status(http.StatusBadRequest)
			return
		}
		_, err = do.UpdateColumn(s.Value, a)
	case "time-frame":
		var t timeframe
		if err := c.BindJSON(&t); err != nil {
			log.Println(err.Error())
			c.Status(http.StatusBadRequest)
			return
		}
		_, err = do.UpdateColumn(s.Value, t)
	}
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}
