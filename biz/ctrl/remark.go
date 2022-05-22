package ctrl

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scutrobotlab/HR/dal/model"
	"gorm.io/gorm"
)

func SetRemark(c *gin.Context) {
	var m model.Remark
	err := c.BindJSON(&m)
	if err != nil {
		log.Println("get remark: ", err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	admin, err := getAdmin(c)
	if err != nil {
		c.Status(errorHandle(err))
		return
	}
	appid, err := getApplicantId(c)
	if err != nil {
		c.Status(errorHandle(err))
		return
	}
	r := q.Remark
	r.WithContext(ctx).Where(r.AdminID.Eq(admin.ID), r.ApplicantID.Eq(appid)).UpdateColumn(r.TheRemark, m.TheRemark)
	c.Status(http.StatusNoContent)
}

func GetRemark(c *gin.Context) {
	admin, err := getAdmin(c)
	if err != nil {
		c.Status(errorHandle(err))
		return
	}
	appid, err := getApplicantId(c)
	if err != nil {
		c.Status(errorHandle(err))
		return
	}
	r := q.Remark
	m, err := r.WithContext(ctx).Where(r.AdminID.Eq(admin.ID), r.ApplicantID.Eq(appid)).Take()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNoContent)
	} else if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, m.TheRemark)
}
