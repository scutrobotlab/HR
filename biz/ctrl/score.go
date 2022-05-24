package ctrl

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scutrobotlab/HR/biz/mid"
	"github.com/scutrobotlab/HR/dal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// @Summary		获取分数
// @Description 管理员获取面试者分数
// @Tags		admin
// @Router		/api/admin/score/:id [GET]
// @Param		id	path	uint	true	"面试者ID"
// @Success		200 {object} model.Score
// @Failure		400,401,500
// @securityDefinitions.basic 管理员身份
func GetScore(c *gin.Context) {
	// 设置 Admin ID
	admin, ok := c.MustGet("admin").(*mid.Admin)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Println("get id: ", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	score, err := q.Score.WithContext(ctx).Where(
		q.Score.AdminID.Eq(admin.ID),
		q.Score.ApplicantID.Eq(uint(id))).
		Take()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, score)
}

// @Summary		设置分数
// @Description 管理员设置面试者分数
// @Tags		admin
// @Router		/api/admin/score [PUT]
// @Param		admit	body	model.Score	true	"分数"
// @Success		204
// @Failure		400,401,500
// @securityDefinitions.basic 管理员身份

// @Summary		删除分数
// @Description 管理员删除面试者分数
// @Tags		admin
// @Router		/api/admin/score [DELETE]
// @Param		score	body	model.Score	true	"分数"
// @Success		204
// @Failure		400,401,500
// @securityDefinitions.basic 管理员身份
func SetScore(c *gin.Context) {
	score := model.Score{}
	err := c.BindJSON(&score)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// 设置 Admin ID
	admin, ok := c.MustGet("admin").(*mid.Admin)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	score.AdminID = admin.ID

	if !checkAdminGroup(admin, score.Group) {
		c.Status(http.StatusForbidden)
		return
	}

	switch c.Request.Method {
	case "PUT":
		err = q.Score.WithContext(ctx).
			Clauses(clause.OnConflict{UpdateAll: true}).
			Create(&score)
	case "DELETE":
		_, err = q.Score.WithContext(ctx).Where(
			q.Score.AdminID.Eq(score.AdminID),
			q.Score.ApplicantID.Eq(score.ApplicantID),
			q.Score.Group.Eq(score.Group)).Delete()
	default:
		c.Status(http.StatusMethodNotAllowed)
		return
	}
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, score)
}
