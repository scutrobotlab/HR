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
)

// @Summary		修改评价
// @Description 管理员修改对特定面试者的评价
// @Tags		admin
// @Router		/api/admin/remark/:id [PUT]
// @Param		id	path	uint	true	"面试者ID"
// @Param		remark	body	model.Remark	true	"评价，仅remark本身，其他内容将被忽略"
// @Success		204
// @Failure		400,401,500
// @securityDefinitions.basic 管理员身份
func SetRemark(c *gin.Context) {
	var remark model.Remark
	err := c.BindJSON(&remark)
	if err != nil {
		log.Println("get remark: ", err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
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
	q.Remark.WithContext(ctx).Where(
		q.Remark.AdminID.Eq(admin.ID),
		q.Remark.ApplicantID.Eq(uint(id))).
		UpdateColumn(q.Remark.TheRemark, remark.TheRemark)
	c.Status(http.StatusNoContent)
}

// @Summary		获取评价
// @Description 管理员获取对特定面试者的评价
// @Tags		admin
// @Router		/api/admin/remark/:id [GET]
// @Param		id	path	uint	true	"面试者ID"
// @Success		204	{object}	model.Remark
// @Failure		400,401,500
// @securityDefinitions.basic 管理员身份
func GetRemark(c *gin.Context) {
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
	remark, err := q.Remark.WithContext(ctx).Where(
		q.Remark.AdminID.Eq(admin.ID),
		q.Remark.ApplicantID.Eq(uint(id))).
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
	c.JSON(http.StatusOK, remark.TheRemark)
}
