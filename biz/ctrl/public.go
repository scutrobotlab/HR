package ctrl

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary 获取设置
// @Description 可获取"form", "announce", "time-frame"
// @Tags public, setting
// @Router /api/admin/public/{key} [GET]
// @Param        key	path	string	true	"获取设置的键"	Enums(form, announce, time-frame)
// @Success      200
// @Failure      404,500
func GetSetting(c *gin.Context) {
	key, ok := c.Params.Get("key")
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}
	setting, err := q.Setting.WithContext(ctx).
		Where(q.Setting.Key.Eq(key)).
		Take()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNotFound)
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Data(http.StatusOK, gin.MIMEJSON, []byte(setting.Value))
}

// @Summary 获取面试题库
// @Description 获取特定组别的面试题库
// @Tags public
// @Router /api/admin/public/exam/ [GET]
// @Param        group	path	string	true	"组别"	Enums(机械, 电控, 视觉)
// @Success      200  {object}  model.Question
// @Failure      404,500
func GetExam(c *gin.Context) {
	group, ok := c.Params.Get("group")
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}
	questions, err := q.Question.WithContext(ctx).
		Where(q.Question.Group.Eq(group)).Find()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNoContent)
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, questions)
}

// @Summary 获取面试时间
// @Description 获取特定组别的面试时间
// @Tags public
// @Router /api/admin/public/time/ [GET]
// @Param        group	path	string	true	"组别"	Enums(机械, 电控, 视觉)
// @Success      200  {object}  model.OptionalTime
// @Success      204
// @Failure      404,500
func GetTime(c *gin.Context) {
	group, ok := c.Params.Get("group")
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}
	opttime, err := q.OptionalTime.WithContext(ctx).
		Where(q.OptionalTime.Group.Eq(group)).
		Take()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNoContent)
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, opttime)
}
