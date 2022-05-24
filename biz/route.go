package biz

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/scutrobotlab/HR/biz/ctrl"
	"github.com/scutrobotlab/HR/biz/mid"
	"github.com/scutrobotlab/HR/conf"
	"github.com/scutrobotlab/HR/dal"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/scutrobotlab/HR/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title HR API
// @version 1.0
// @description 面试系统后端程序

// @host localhost:8080
// @BasePath /api

func init() {
	dal.DB = dal.ConnectDB(conf.Postgres).Debug()
}

func Run() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	sessionNames := []string{"admin", "applicant"}
	r.Use(sessions.SessionsMany(sessionNames, store))

	public := r.Group("/api/public") // 公开接口
	{
		public.GET("/time/:group", ctrl.GetTime) // 面试时间
		public.GET("/exam/:group", ctrl.GetExam) // 面试题库
		public.GET("/:key", ctrl.GetSetting)     // 报名表信息
	}

	r.POST("/api/applicant/login/:token", mid.ApplicantLogin) // 报名者登录
	applicant := r.Group("/api/applicant")                    // 报名者登录后接口
	applicant.Use(mid.ApplicantAuth())
	{
		applicant.GET("/info", ctrl.ApplicantInfo) // 我的微信信息
		applicant.GET("/status")                   // 我的状态信息（步骤+面试时间+面试结果）
		applicant.GET("/exam/:group")              // 我的题库回答

		applicant.POST("/apply") // 提交表单
		applicant.POST("/exam")  // 提交题库回答
		applicant.POST("/time")  // 选择面试时间
	}

	r.POST("/api/admin/login/:code", mid.AdminLogin) // 面试官登录
	admin := r.Group("/api/admin")                   // 面试官登录后接口
	admin.Use(mid.AdminAuth())
	{
		admin.GET("/info", ctrl.AdminInfo) // 我的信息
		admin.POST("/logout")              // 登出

		admin.PUT("/standard", ctrl.AdminSetStandard) // 设置默认评价标准

		admin.PUT("/setting/:key", ctrl.UpdateSetting) // 修改设置

		admin.GET("/applicant:id", ctrl.AdminGetApplicantInfo) // 获取面试者信息，包括分数、排名、状态
		admin.PUT("/admit", ctrl.AdminAdmit)                   // 录取
		admin.DELETE("/admit", ctrl.AdminAdmit)                // 取消录取
		admin.GET("/remark/:id", ctrl.GetRemark)               // 获取我对特定面试者的评价
		admin.POST("/remark/:id", ctrl.SetRemark)              // 修改我对特定面试者的评价
		score := admin.Group("/score")                         // 我对特定面试者的评分
		{
			score.GET("/:id", ctrl.GetScore)    // 获取评分
			score.POST("/:id", ctrl.SetScore)   // 修改评分
			score.DELETE("/:id", ctrl.SetScore) // 清除评分
		}
		aas := admin.Group("/applicants") // 面试者批量管理
		{
			aas.GET("/:group")    // 获取组别名单（用于展示）
			aas.PUT("/name-list") // 全部名单（用于搜索）
		}
		question := admin.Group("/question") // 面试题库管理
		{
			question.POST("/")      // 新增题目
			question.PUT("/:id")    // 更新题目
			question.DELETE("/:id") // 删除题目
		}
		standard := admin.Group("/standard") // 评价标准管理
		{
			standard.GET("/")       // 获取所有评价标准
			standard.GET("/:id")    // 获取特定评价标准
			standard.POST("/")      // 新增评价标准
			standard.PUT("/:id")    // 更新评价标准
			standard.DELETE("/:id") // 删除评价标准
		}
		time := admin.Group("/time") // 面试时间管理
		{
			time.GET("/:group")    // 导出面试时间(CSV)
			time.PUT("/:group")    // 上传面试时间(CSV)
			time.DELETE("/:group") // 清空该组面试时间
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
