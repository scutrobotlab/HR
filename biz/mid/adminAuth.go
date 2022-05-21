package mid

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 管理员登录验证中间件
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if gin.Mode() != gin.ReleaseMode {
			c.Set("admin", "test_admin")
		} else {
			session := sessions.Default(c)
			if admin := session.Get("admin"); admin == "" {
				c.AbortWithStatus(401)
				return
			} else {
				c.Set("admin", session.Get("admin"))
			}
		}

		c.Next()

		status := c.Writer.Status()
		log.Println(status)
	}
}
