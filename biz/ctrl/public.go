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
	s := q.Setting
	do := s.WithContext(ctx)
	m, err := do.Where(s.Key.Eq(key)).Take()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNotFound)
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Data(http.StatusOK, gin.MIMEJSON, []byte(m.Value))
}

// @Summary 获取面试题库
// @Description 获取特定组别的面试题库
// @Tags public
// @Router /api/admin/public/exam/{group} [GET]
// @Param        group	path	string	true	"组别"	Enums(机械, 电控, 视觉)
// @Success      200  {object}  model.Question
// @Failure      404,500
func GetExam(c *gin.Context) {
	group, ok := c.Params.Get("group")
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}
	question := q.Question
	do := question.WithContext(ctx)
	m, err := do.Where(question.Group.Eq(group)).Take()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNotFound)
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, m)
}

// @Summary 获取面试时间
// @Description 获取特定组别的面试时间
// @Tags public
// @Router /api/admin/public/time/{group} [GET]
// @Param        group	path	string	true	"组别"	Enums(机械, 电控, 视觉)
// @Success      200  {object}  model.OptionalTime
// @Failure      404,500
func GetTime(c *gin.Context) {
	group, ok := c.Params.Get("group")
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}
	t := q.OptionalTime
	do := t.WithContext(ctx)
	m, err := do.Where(t.Group.Eq(group)).Take()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNotFound)
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, m)
}
