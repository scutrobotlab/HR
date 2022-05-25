package ctrl

import (
	"bytes"
	"encoding/csv"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/scutrobotlab/HR/biz/mid"
	"github.com/scutrobotlab/HR/dal/model"
	"gorm.io/gorm"
)

type OptionalTime struct {
	TheDate     string `json:"date" example:"2022-01-02"`
	TheTime     string `json:"time" example:"15:04:05"`
	TheLocation string `json:"location" example:"春茧体育馆"`
	Group       string `json:"group" example:"机械"`
	IntentRank  string `json:"rank" example:"1"`
	TotalCount  string `json:"total_cnt" example:"1"`
}

func (t *OptionalTime) Get() (*model.OptionalTime, error) {
	date, err := time.Parse("2006/01/02", t.TheDate)
	if err != nil {
		return nil, err
	}
	time, err := time.Parse("15:04:05", t.TheTime)
	if err != nil {
		return nil, err
	}
	rank, err := strconv.ParseUint(t.IntentRank, 10, 32)
	if err != nil {
		return nil, err
	}
	cnt, err := strconv.ParseUint(t.TotalCount, 10, 32)
	if err != nil {
		return nil, err
	}
	return &model.OptionalTime{
		TheDate:     date,
		TheTime:     time,
		TheLocation: t.TheLocation,
		Group:       t.Group,
		IntentRank:  uint(rank),
		TotalCount:  uint(cnt),
	}, nil
}

// @Summary		设置面试时间
// @Description	管理员设置面试时间
// @Tags	admin
// @Accept	json
// @Param	times	body	[]OptionalTime	true	"面试时间"
// @Success 204
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
	var result []struct {
		Group string `json:"group"`
		Count int    `json:"count"`
	}
	err := q.OptionalTime.WithContext(ctx).Select(
		q.OptionalTime.Group,
		q.OptionalTime.ID.Count().As("count")).
		Group(q.OptionalTime.Group).
		Scan(&result)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary		清空组别面试时间
// @Description	清空给定组别的所有面试时间
// @Tags	admin
// @Accept	json
// @Param	groups	body	[]string	true	"给定的组别"
// @Success 204
// @Router	/api/admin/time	[DELETE]
func DelOptTime(c *gin.Context) {
	var groups []string
	err := c.BindJSON(&groups)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	admin, ok := c.MustGet("admin").(*mid.Admin)
	if !ok {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	for _, g := range groups {
		if !checkAdminGroup(admin, g) {
			c.String(http.StatusBadRequest,
				"No %s group permission.", g)
			return
		}
	}
	_, err = q.OptionalTime.WithContext(ctx).
		Where(q.OptionalTime.Group.In(groups...)).
		Delete()
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary		下载组别面试时间csv文件
// @Description	下载给定组别的所有面试时间为csv文件
// @Tags	admin
// @Accept	json
// @Param	groups	body	[]string	true	"给定的组别"
// @Success 200
// @Router	/api/admin/time	[DELETE]
func GetOptTimeCSV(c *gin.Context) {
	var groups []string
	err := c.BindJSON(&groups)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	opttimes, err := q.OptionalTime.WithContext(ctx).
		Where(q.OptionalTime.Group.In(groups...)).
		Find()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.String(http.StatusNoContent, "No Record Found")
		return
	}
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	buf := &bytes.Buffer{}
	w := csv.NewWriter(buf)

	record := []string{
		"组别",
		"日期",
		"时间",
		"地点",
		"第几志愿",
		"人数"}
	if err := w.Write(record); err != nil {
		log.Println("error writing record to csv: ", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	for _, opt := range opttimes {
		record := []string{
			opt.Group,
			opt.TheDate.Format("2006/01/02"),
			opt.TheTime.Format("15:04:05"),
			opt.TheLocation,
			strconv.FormatUint(uint64(opt.IntentRank), 10),
			strconv.FormatUint(uint64(opt.TotalCount), 10)}
		if err := w.Write(record); err != nil {
			log.Println("error writing record to csv: ", err)
			c.Status(http.StatusInternalServerError)
			return
		}
	}
	c.Header("Content-Disposition", "attachment; filename=export.csv")
	c.Data(http.StatusOK, "text/csv", buf.Bytes())
}
