package ctrl

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/scutrobotlab/HR/biz/mid"
	"github.com/scutrobotlab/HR/dal/model"
)

type OptionalTime struct {
	TheDate     string `json:"date" example:"2022-05-20"`
	TheTime     string `json:"time" example:"13:14:20"`
	TheLocation string `json:"location" example:"五山路情缘酒店"`
	Group       string `json:"group" example:"机械"`
	IntentRank  *int16 `json:"intent" example:"1"`
}

func (t *OptionalTime) Get() (*model.OptionalTime, error) {
	date, err := time.Parse("2006-01-02", t.TheDate)
	if err != nil {
		return nil, err
	}
	time, err := time.Parse("15:04:05", t.TheTime)
	if err != nil {
		return nil, err
	}
	return &model.OptionalTime{
		TheDate:     date,
		TheTime:     time,
		TheLocation: t.TheLocation,
		Group:       t.Group,
		IntentRank:  t.IntentRank,
	}, nil
}

// @Summary		设置面试时间文件
// @Description	管理员上传面试时间csv文件
// @Tags	admin
// @Accept	json
// @Param	times	body	[]OptionalTime	true	"数据"
// @Success 200
// @Router	/api/admin/time/	[POST]
func SetOptTime(c *gin.Context) {
	admin, ok := c.MustGet("admin").(*mid.Admin)
	if !ok {
		log.Println("cannot convert admin")
		log.Println("admin")
		c.Status(http.StatusInternalServerError)
		return
	}
	_opts := []*OptionalTime{}
	err := c.BindJSON(&_opts)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	opts := make([]*model.OptionalTime, 0, len(_opts))
	for _, o := range _opts {
		_o, err := o.Get()
		if err != nil {
			log.Println(err.Error())
			c.String(http.StatusBadRequest,
				"can not convert %v to %v",
				o, _o)
			return
		}
		opts = append(opts, _o)
	}

	// 检查内容中的组权限
	for _, opt := range opts {
		if !checkAdminGroup(admin, opt.Group) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "No group permission.",
				"opttime": opt,
				"groups":  admin.Groups})
			return
		}
	}
	err = q.OptionalTime.WithContext(ctx).Create(opts...)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary		统计面试时间
// @Description	统计所有组别面试时间的数量
// @Tags	admin
// @Accept	json
// @Success 200
// @Router	/api/admin/time/cnt	[GET]
func CntOptTime(c *gin.Context) {
	// err := q.OptionalTime.WithContext(ctx).
	// if err != nil {
	// 	c.Status(http.StatusInternalServerError)
	// 	return
	// }
}
