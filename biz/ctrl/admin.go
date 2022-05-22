package ctrl

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scutrobotlab/HR/biz/mid"
)

func AdminInfo(c *gin.Context) {
	a, ok := c.Get("admin")
	if !ok {
		log.Println("admin unfound")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, a)
}

func AdminSetStandard(c *gin.Context) {
	var standard struct{ StandardID uint }
	if err := c.BindJSON(&standard); err != nil {
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
	_, err := do.Where(qa.ID.Eq(admin.ID)).UpdateColumn(qa.StandardID, standard.StandardID)
	if err != nil {
		log.Println("cannot update standard")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}
