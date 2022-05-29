package ctrl

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary		面试者信息
// @Description 管理员获取面试者信息
// @Tags		admin
// @Router		/api/admin/applicant/{id} [GET]
// @Param		id	path	uint	true	"面试者ID"
// @Success		200  {object}  model.Applicant
// @Failure		401,404,500
// @securityDefinitions.basic 管理员身份
// TODO 分数排名状态
func AdminGetApplicantInfo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Println("get id: ", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	applicant, err := q.Applicant.WithContext(ctx).
		Where(q.Applicant.ID.Eq(uint(id))).
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
	c.JSON(http.StatusOK, applicant)
}

// @Summary		面试者名单
// @Description 管理员获取面试者名单
// @Tags		admin
// @Router		/api/admin/applicants/name-list [GET]
// @Success		200  {array}	[]string	"name list"
// @Failure		401,404,500
// @securityDefinitions.basic 管理员身份
func AdminGetNameList(c *gin.Context) {
	applicants, err := q.Applicant.WithContext(ctx).
		Select(q.Applicant.Name).
		Find()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, applicants)
}
