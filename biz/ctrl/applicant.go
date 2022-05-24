package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		面试者微信信息
// @Description 面试者获取自己的微信信息
// @Tags		applicant
// @Router		/api/applicant/info [get]
// @Success		200  {object}  mid.Applicant
// @Failure		401,500
// @securityDefinitions.basic 面试者身份
func ApplicantInfo(c *gin.Context) {
	applicant, ok := c.Get("applicant")
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, applicant)
}
