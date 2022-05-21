package mid

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 面试者登录验证中间件
func ApplicantAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		if applicant := session.Get("applicant"); applicant == "" {
			c.AbortWithStatus(401)
			return
		} else {
			c.Set("applicant", session.Get("applicant"))
		}

		c.Next()

		status := c.Writer.Status()
		log.Println(status)
	}
}
