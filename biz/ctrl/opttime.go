package ctrl

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/scutrobotlab/HR/dal/model"
)

var OptTimeCSVHeader = [...]string{"组别", "日期", "时间", "地点", "第几志愿", "人数"}

// @Summary		设置面试时间文件
// @Description	管理员上传面试时间csv文件
// @Tags	admin
// @Accept	multipart/form-data
// @Param	group	path		string	true "组"
// @Param	file	formData	file	true "csv 文件"
// @Produce	json
// @Success 204
// @Router	/api/admin/time/{group}	[PUT]
func SetOptTime(c *gin.Context) {
	file_ptr, err := c.FormFile("file")
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	log.Println(file_ptr.Filename)
	file, err := file_ptr.Open()
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	for i, head := range records[0] {
		if head != OptTimeCSVHeader[i] {
			c.String(http.StatusBadRequest,
				"csv#1: column %d should be %s, but got %s.",
				i, OptTimeCSVHeader[i], head)
			return
		}
	}
	opttimes := []*model.OptionalTime{}
	for l, line := range records {
		fmt.Println(line)
		if l == 0 {
			continue
		}
		_date, err := time.Parse("2006-01-02", line[1])
		if err != nil {
			c.String(http.StatusBadRequest,
				"csv#%d: parse date %s failed.",
				l, line[1])
			return
		}
		_time, err := time.Parse("15:04:05", line[2])
		if err != nil {
			c.String(http.StatusBadRequest,
				"csv#%d: parse time %s failed.",
				l, line[2])
			return
		}
		_rank, err := strconv.ParseUint(line[4], 10, 32)
		if err != nil {
			c.String(http.StatusBadRequest,
				"csv#%d: parse rank %s failed.",
				l, line[4])
			return
		}
		rank := uint(_rank)
		opt := &model.OptionalTime{
			Group:       line[0],
			TheDate:     _date,
			TheTime:     _time,
			TheLocation: line[3],
			IntentRank:  &rank,
		}
		opttimes = append(opttimes, opt)
	}
	err = q.OptionalTime.WithContext(ctx).Create(opttimes...)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
}
