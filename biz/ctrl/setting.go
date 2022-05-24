package ctrl

import (
	"encoding/json"
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

// @Summary 设置报名表
// @Description 修改报名表设置
// @Tags admin, setting
// @Router /api/setting/form [PUT]
// @Param        form	body	form	true	"报名表"
// @Success      200
// @Failure      400,401,500
func updateForm(c *gin.Context) ([]byte, error) {
	var f form
	if err := c.BindJSON(&f); err != nil {
		return []byte{}, err
	}
	return json.Marshal(f)
}

// @Summary 设置公告
// @Description 修改公告设置
// @Tags admin, setting
// @Router /api/setting/announce [PUT]
// @Param        announce	body	announce	true	"公告"
// @Success      200
// @Failure      400,401,500
func updateAnnounce(c *gin.Context) ([]byte, error) {
	var a announce
	if err := c.BindJSON(&a); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a)
}

// @Summary 设置时间节点
// @Description 修改时间节点设置
// @Tags admin, setting
// @Router /api/setting/time-frame [PUT]
// @Param        timeframe	body	timeframe	true	"时间节点"
// @Success      200
// @Failure      400,401,500
func UpdateTimeFrame(c *gin.Context) ([]byte, error) {
	var t timeframe
	if err := c.BindJSON(&t); err != nil {
		return []byte{}, err
	}
	return json.Marshal(&t)
}

func UpdateSetting(c *gin.Context) {
	key, ok := c.Params.Get("key")
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}
	var (
		v   []byte
		err error
	)
	switch key {
	case "form":
		v, err = updateForm(c)
	case "announce":
		v, err = updateAnnounce(c)
	case "time-frame":
		v, err = UpdateTimeFrame(c)
	}
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	_, err = q.Setting.WithContext(ctx).
		Where(q.Setting.Key.Eq(key)).
		Update(q.Setting.Value, v)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}
