package ctrl

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scutrobotlab/HR/biz/mid"
	"github.com/scutrobotlab/HR/dal/model"
)

type Intent struct {
	Group          string `json:"group"`
	IntentRank     uint   `json:"rank"`
	OptionalTimeID uint   `json:"time_id"`
}

// @Summary		面试者选择面试时间
// @Description	applicant choose optional time for intent
// @Tags		applicant
// @Router		/api/applicant/time [POST]
// @Param		time		body	model.Intent	true	"包括组别、轮次、面试时间ID"
// @Success		204
// @Failure		401,500
// @securityDefinitions.basic 面试者身份
func ApplicantChooseTime(c *gin.Context) {
	a, ok := c.MustGet("applicant").(*mid.Applicant)
	if !ok {
		log.Println("can not get applicant")
		c.Status(http.StatusInternalServerError)
		return
	}

	var intent model.Intent
	err := c.BindJSON(&intent)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	applicant, err := q.Applicant.WithContext(ctx).
		Where(q.Applicant.OpenID.Eq(a.OpenID)).
		Select(q.Applicant.ID).
		Take()
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	_, err = q.Intent.WithContext(ctx).
		Where(
			q.Intent.ApplicantID.Eq(applicant.ID),
			q.Intent.Group.Eq(intent.Group),
			q.Intent.IntentRank.Eq(intent.IntentRank)).
		Delete()
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	err = q.Intent.WithContext(ctx).
		Create(&intent)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}
