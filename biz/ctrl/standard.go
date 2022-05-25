package ctrl

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scutrobotlab/HR/biz/mid"
	"github.com/scutrobotlab/HR/dal/model"
)

// @Summary		所有标准
// @Description 管理员获取所有评价标准
// @Tags		admin
// @Router		/api/admin/standard [GET]
// @Success		200	{object}	[]model.Standard
// @Failure		401,500
// @securityDefinitions.basic 管理员身份
func GetStandards(c *gin.Context) {
	stds, err := q.Standard.WithContext(ctx).Find()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, stds)
}

// @Summary		单个标准
// @Description 管理员获取特定评价标准
// @Tags		admin
// @Router		/api/admin/standard/{id} [GET]
// @Param		id	path	uint	true	"评价标准ID"
// @Success		200	{object}	model.Standard
// @Failure		400,401,500
// @securityDefinitions.basic 管理员身份
func GetStandard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	std, err := q.Standard.WithContext(ctx).
		Where(q.Standard.ID.Eq(uint(id))).
		Take()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, std)
}

// @Summary		添加标准
// @Description 管理员添加评价标准
// @Tags		admin
// @Router		/api/admin/standard [POST]
// @Param		standard	body	model.Standard	true	"标准"
// @Success		204
// @Failure		400,401,500
// @securityDefinitions.basic 管理员身份
func AddStandard(c *gin.Context) {
	var std model.Standard
	err := c.BindJSON(&std)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	// 设置管理员ID
	admin, ok := c.MustGet("admin").(*mid.Admin)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	std.AdminID = admin.ID

	if err := q.Standard.WithContext(ctx).Create(&std); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary		更新标准
// @Description 管理员更新特定评价标准
// @Tags		admin
// @Router		/api/admin/standard/{id} [PUT]
// @Param		id	path	uint	true	"评价标准ID"
// @Param		standard	body	model.Standard	true	"标准"
// @Success		204
// @Failure		400,401,500
// @securityDefinitions.basic 管理员身份
func UpdateStandard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	var std model.Standard
	err = c.BindJSON(&std)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	// 设置管理员ID
	admin, ok := c.MustGet("admin").(*mid.Admin)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	std.AdminID = admin.ID

	_, err = q.Standard.WithContext(ctx).
		Where(q.Standard.ID.Eq(uint(id))).Select(
		q.Standard.Name,
		q.Standard.Content,
		q.Standard.AdminID).
		Updates(std)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, std)
}

// @Summary		删除标准
// @Description 管理员删除特定评价标准
// @Tags		admin
// @Router		/api/admin/standard/{id} [DELETE]
// @Param		id	path	uint	true	"评价标准ID"
// @Success		204
// @Failure		400,401,500
// @securityDefinitions.basic 管理员身份
func DeleteStandard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	_, err = q.Standard.WithContext(ctx).
		Where(q.Standard.ID.Eq(uint(id))).Delete()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}
