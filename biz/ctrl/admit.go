package ctrl

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scutrobotlab/HR/biz/mid"
	"github.com/scutrobotlab/HR/dal/model"
	"gorm.io/gorm/clause"
)

// @Summary		录取/取消录取
// @Description 管理员录取/取消录取面试者到特定组别
// @Tags		admin
// @Router		/api/admin/admit [PUT]
// @Router		/api/admin/admit [DELETE]
// @Param		admit	body	model.Admit	true	"录取"
// @Success		204
// @Failure		400,401,500
// @securityDefinitions.basic 管理员身份
func AdminAdmit(c *gin.Context) {
	admit := model.Admit{}
	err := c.BindJSON(&admit)
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
	admit.AdminID = admin.ID

	// 检查是否有对应组别的权限
	if !checkAdminGroup(admin, admit.Group) {
		c.Status(http.StatusForbidden)
		return
	}

	switch c.Request.Method {
	case "PUT":
		err = q.Admit.WithContext(ctx).
			Clauses(clause.OnConflict{DoNothing: true}).
			Create(&admit)
	case "DELETE":
		_, err = q.Admit.WithContext(ctx).
			Where(
				q.Admit.ApplicantID.Eq(admit.ApplicantID),
				q.Admit.Group.Eq(admit.Group)).
			Delete()
	default:
		c.Status(http.StatusMethodNotAllowed)
		return
	}
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, admit)
}
