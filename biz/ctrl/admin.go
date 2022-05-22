package ctrl

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scutrobotlab/HR/biz/mid"
)

// @Summary 管理员信息
// @Description 获取管理员信息
// @Tag admin
// @Router /api/admin/info [get]
// @Success      200  {object}  model.Admin
// @Failure      401
// @Failure      500
// @securityDefinitions.basic 管理员身份
func AdminInfo(c *gin.Context) {
	a, ok := c.Get("admin")
	if !ok {
		log.Println("admin unfound")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, a)
}

type standard struct {
	StandardID uint `json:"standard_id" binding:"required"`
}

// @Summary 设置默认标准
// @Description 设置默认评价标准
// @Tag admin
// @Router /api/admin/standard [PUT]
// @Param        Standard	body	standard	true	"评价标准的ID"
// @Success      204
// @Failure      400,401,500
// @securityDefinitions.basic 管理员身份
func AdminSetStandard(c *gin.Context) {
	var std standard
	if err := c.BindJSON(&std); err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	a, ok := c.Get("admin")
	if !ok {
		log.Println("admin unfound")
		c.Status(http.StatusInternalServerError)
		return
	}
	admin, ok := a.(*mid.Admin)
	if !ok {
		log.Println("cannot convert admin")
		c.Status(http.StatusInternalServerError)
		return
	}
	qa := q.Admin
	do := qa.WithContext(ctx)
	_, err := do.Where(qa.ID.Eq(admin.ID)).UpdateColumn(qa.StandardID, std.StandardID)
	if err != nil {
		log.Println("cannot update standard")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

func getAdmin(c *gin.Context) (*mid.Admin, error) {
	a, ok := c.Get("admin")
	if !ok {
		log.Println("admin unfound")
		return nil, ErrInternalServer
	}
	admin, ok := a.(*mid.Admin)
	if !ok {
		log.Println("cannot convert admin")
		return nil, ErrInternalServer
	}
	return admin, nil
}

func getApplicantId(c *gin.Context) (uint, error) {
	a := c.Param("id")
	appid, err := strconv.ParseUint(a, 10, 32)
	if err != nil {
		log.Println("get applicant id: ", err.Error())
		return 0, ErrBadRequest
	}
	return uint(appid), nil
}
