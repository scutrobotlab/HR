package ctrl

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scutrobotlab/HR/biz/mid"
	"github.com/scutrobotlab/HR/dal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// @Summary		面试者微信信息
// @Description 面试者获取自己的微信信息
// @Tags		applicant
// @Router		/api/applicant/wechat [GET]
// @Success		200  {object}  mid.Applicant
// @Failure		401,500
// @securityDefinitions.basic 面试者身份
func ApplicantWechat(c *gin.Context) {
	applicant, ok := c.MustGet("applicant").(*mid.Applicant)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, applicant)
}

// @Summary		面试者状态信息
// @Description 面试者获取自己的状态信息
// @Tags		applicant
// @Router		/api/applicant/status [GET]
// @Success		200  {object}  model.Applicant
// @Failure		401,500
// @securityDefinitions.basic 面试者身份
func ApplicantStatus(c *gin.Context) {
	app, ok := c.MustGet("applicant").(*mid.Applicant)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	applicant, err := q.Applicant.WithContext(ctx).
		Where(q.Applicant.OpenID.Eq(app.OpenID)).
		Take()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNoContent)
		return
	}
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, applicant)
}

// @Summary 面试者提交报名表
// @Description 面试者提交报名表
// @Tags		applicant
// @Router		/api/applicant/apply [PUT]
// @Param		body		body 	model.Applicant	true		"报名表"
// @Success		204
// @Failure		401,500
// @securityDefinitions.basic 面试者身份
func ApplicantApply(c *gin.Context) {
	app, ok := c.MustGet("applicant").(*mid.Applicant)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	var app_form model.Applicant
	err := c.BindJSON(&app_form)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	app_form.OpenID = app.OpenID
	err = q.Applicant.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "open_id"}},
			UpdateAll: true,
		}).Create(&app_form)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}
