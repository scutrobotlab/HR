package biz

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/scutrobotlab/HR/biz/ctrl"
	"github.com/scutrobotlab/HR/biz/mid"
	"github.com/scutrobotlab/HR/conf"
	"github.com/scutrobotlab/HR/dal"
)

func init() {
	dal.DB = dal.ConnectDB(conf.Postgres).Debug()
}

func Run() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	public := r.Group("/api/public") // 公开接口
	{
		public.GET("/time/:group")           // 面试时间
		public.GET("/exam/:group")           // 面试题库
		public.GET("/:key", ctrl.GetSetting) // 报名表信息
	}

	r.POST("/api/applicant/login/:token")  // 报名者登录
	applicant := r.Group("/api/applicant") // 报名者登录后接口
	applicant.Use(mid.ApplicantAuth())
	{
		applicant.GET("/wechat") // 我的微信信息
		applicant.GET("/status") // 我的状态信息（步骤+面试时间+面试结果）
		applicant.GET("/exam")   // 我的题库回答

		applicant.POST("/apply") // 提交表单
		applicant.POST("/exam")  // 提交题库回答
		applicant.POST("/time")  // 选择面试时间
	}

	r.POST("/api/admin/login/:token") // 面试官登录
	admin := r.Group("/api/admin")    // 面试官登录后接口
	{
		admin.GET("/info")    // 我的信息
		admin.POST("/logout") // 登出

		admin.PUT("/standard") // 设置默认评价标准

		form := admin.Group("/form") // 报名表管理
		{
			form.POST("/")      // 新增表项
			form.PUT("/:id")    // 修改表项
			form.DELETE("/:id") // 删除表项
		}
		group := admin.Group("/group") // 组管理
		{
			group.POST("/")      // 新增组
			group.PUT("/:id")    // 修改组
			group.DELETE("/:id") // 删除组
		}
		aa := admin.Group("/applicant") // 面试者管理
		{
			aa.GET("/:id")    // 获取面试者信息，包括分数、排名和状态
			aa.PUT("/admit")  // 录取/取消录取
			aa.DELETE("/:id") // 删除面试者
		}
		remark := admin.Group("/remark") // 我对特定面试者的评价
		{
			remark.GET("/:id")    // 获取评价
			remark.POST("/:id")   // 修改评价
			remark.DELETE("/:id") // 清除评价
		}
		score := admin.Group("/score") // 我对特定面试者的评分
		{
			score.GET("/:id")    // 获取评分
			score.POST("/:id")   // 修改评分
			score.DELETE("/:id") // 清除评分
		}
		aas := admin.Group("/applicants") // 面试者批量管理
		{
			aas.GET("/:group")    // 获取组别名单（用于展示）
			aas.PUT("/name-list") // 全部名单（用于搜索）
		}
		exam := admin.Group("/exam") // 面试题库管理
		{
			exam.POST("/")      // 新增题目
			exam.PUT("/:id")    // 更新题目
			exam.DELETE("/:id") // 删除题目
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
		admin.PUT("/time-frame/:key") // 设置时间节点
		admin.PUT("/announce/:key")   // 设置公告
	}

	r.Run()
}
