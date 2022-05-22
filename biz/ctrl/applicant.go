package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApplicantInfo(c *gin.Context) {
	applicant, ok := c.Get("applicant")
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, applicant)
}
