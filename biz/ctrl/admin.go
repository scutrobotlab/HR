package ctrl

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scutrobotlab/HR/biz/mid"
	"github.com/scutrobotlab/HR/dal/model"
)

// @Summary		管理员信息
// @Description 获取管理员信息
// @Tags		admin
// @Router		/api/admin/info [get]
// @Success		200  {object}  mid.Admin
// @Failure		401,500
// @securityDefinitions.basic 管理员身份
func AdminInfo(c *gin.Context) {
	admin, ok := c.MustGet("admin").(*mid.Admin)
	if !ok || admin == nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, admin)
}

// @Summary 设置默认标准
// @Description 设置默认评价标准
// @Tags admin
// @Router /api/admin/standard [PUT]
// @Param        admin	body	model.Admin	true	"仅选取评价标准的ID"
// @Success      204
// @Failure      400,401,500
// @securityDefinitions.basic 管理员身份
func AdminSetStandard(c *gin.Context) {
	var std model.Admin
	if err := c.BindJSON(&std); err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	admin, ok := c.MustGet("admin").(*mid.Admin)
	if !ok {
		log.Println("cannot convert admin")
		c.Status(http.StatusInternalServerError)
		return
	}
	_, err := q.Admin.WithContext(ctx).
		Where(q.Admin.ID.Eq(admin.ID)).
		UpdateColumn(q.Admin.StandardID, std.StandardID)
	if err != nil {
		log.Println("cannot update standard")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

// 检查admin是否有group组别的权限
func checkAdminGroup(admin *mid.Admin, group string) bool {
	if admin == nil {
		return false
	}
	permit := false
	for _, g := range admin.Groups {
		if g.Name == group {
			permit = true
		}
	}
	return permit
}
