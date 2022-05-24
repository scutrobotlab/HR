package ctrl

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scutrobotlab/HR/biz/mid"
	"github.com/scutrobotlab/HR/dal/model"
)

// @Summary		添加题目
// @Description 管理员添加面试题目
// @Tags		admin
// @Router		/api/admin/question [POST]
// @Param		question	body	model.Question	true	"题目"
// @Success		204
// @Failure		400,401,500
// @securityDefinitions.basic 管理员身份
func AddQuestion(c *gin.Context) {
	var question model.Question
	err := c.BindJSON(&question)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	// 检查是否有对应组别的权限
	admin, ok := c.MustGet("admin").(*mid.Admin)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	if !checkAdminGroup(admin, question.Group) {
		c.Status(http.StatusForbidden)
		return
	}
	if err := q.Question.WithContext(ctx).Create(&question); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary		修改题目
// @Description 管理员修改面试题目
// @Tags		admin
// @Router		/api/admin/question/:id [PUT]
// @Param		id	path	uint	true	"题目ID"
// @Param		question	body	string	true	"题目，仅question本身，其他内容将被忽略"
// @Success		204
// @Failure		400,401,500
// @securityDefinitions.basic 管理员身份
func UpdateQuestion(c *gin.Context) {
	var question model.Question // m
	err := c.BindJSON(&question)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	// 检查是否有对应组别的权限
	admin, ok := c.MustGet("admin").(*mid.Admin)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	if !checkAdminGroup(admin, question.Group) {
		c.Status(http.StatusForbidden)
		return
	}
	// 获取qid
	_qid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	qid := uint(_qid)
	_, err = q.Question.WithContext(ctx).
		Where(q.Question.ID.Eq(qid)).
		UpdateColumn(q.Question.TheQuestion, question.TheQuestion)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary		删除分数
// @Description 管理员删除面试题目
// @Tags		admin
// @Router		/api/admin/question/:id [DELETE]
// @Param		id	path	uint	true	"题目ID"
// @Success		204
// @Failure		400,401,500
// @securityDefinitions.basic 管理员身份
func DeleteQuestion(c *gin.Context) {
	_qid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	qid := uint(_qid)
	_, err = q.Question.WithContext(ctx).
		Where(q.Question.ID.Eq(qid)).Delete()
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}
