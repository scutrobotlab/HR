package ctrl

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scutrobotlab/HR/biz/mid"
	"github.com/scutrobotlab/HR/dal/model"
	"gorm.io/gorm/clause"
)

type Answer struct {
	QuestionID  uint   `json:"question_id"`
	TheQuestion string `json:"question"`
	TheAnswer   bool   `json:"answer"`
}

// @Summary		面试者获取特定组问题回答
// @Description	applicant get answers of a group
// @Tags		applicant
// @Router		/api/applicant/answer/{group} [GET]
// @Param		group		path	string	true	"group"
// @Success		200  {object}  []Answer
// @Failure		401,500
// @securityDefinitions.basic 面试者身份
func ApplicantGetAnswer(c *gin.Context) {
	a, ok := c.MustGet("applicant").(*mid.Applicant)
	if !ok {
		log.Println("can not get applicant")
		c.Status(http.StatusInternalServerError)
		return
	}

	var answers []Answer

	app := q.Applicant
	que := q.Question
	ans := q.Answer
	err := ans.WithContext(ctx).Select(que.TheQuestion, ans.TheAnswer).
		Join(que, que.ID.EqCol(ans.QuestionID)).
		Join(app, app.ID.EqCol(ans.ApplicantID)).
		Where(
			que.Group.Eq(c.Param("group")),
			app.OpenID.Eq(a.OpenID)).
		Scan(&answers)

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, answers)
}

// @Summary		面试者提交问题回答
// @Description	applicant submit answers
// @Tags		applicant
// @Router		/api/applicant/answer [POST]
// @Param		answers		body	model.Answer	true	"answers"
// @Success		204
// @Failure		401,500
// @securityDefinitions.basic 面试者身份
func ApplicantSubmitAnswer(c *gin.Context) {
	a, ok := c.MustGet("applicant").(*mid.Applicant)
	if !ok {
		log.Println("can not get applicant")
		c.Status(http.StatusInternalServerError)
		return
	}

	var answer model.Answer
	err := c.BindJSON(&answer)
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
	answer.ApplicantID = applicant.ID

	app := q.Applicant
	que := q.Question
	ans := q.Answer
	err = ans.WithContext(ctx).
		Join(que, que.ID.EqCol(ans.QuestionID)).
		Join(app, app.ID.EqCol(ans.ApplicantID)).
		Where(app.OpenID.Eq(a.OpenID)).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "ApplicantID"}, {Name: "QuestionID"}},
			DoUpdates: clause.Assignments(map[string]interface{}{"role": "TheAnswer"}),
		}).
		Create(&answer)

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}
