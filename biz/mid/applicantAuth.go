package mid

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/scutrobotlab/HR/conf"
)

type Applicant struct {
	NickName   string `json:"nickname"`
	OpenID     string `json:"openid"`
	HeadImgUrl string `json:"headimgurl"`
}

// 面试者登录验证中间件
func ApplicantAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.DefaultMany(c, "applicant")
		if applicant := session.Get("applicant"); applicant == nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		} else {
			c.Set("applicant", applicant)
		}

		c.Next()
	}
}

// @Summary 面试者登录
// @Description 面试者登录
// @Router /api/applicant/login/{token} [POST]
// @Param        token	path	string	true	"oauth2 token"
// @Success      200	{object}	Applicant
// @Failure      401,404
func ApplicantLogin(c *gin.Context) {
	session := sessions.DefaultMany(c, "applicant")
	var applicant Applicant
	if gin.Mode() != gin.ReleaseMode {
		applicant = Applicant{
			NickName:   "本地测试",
			OpenID:     "test",
			HeadImgUrl: "https://res.wx.qq.com/a/wx_fed/webwx/res/static/img/2KriyDK.png",
		}
	} else {
		token, ok := c.Params.Get("token")
		if !ok {
			c.Status(http.StatusNotFound)
			return
		}
		url := conf.WechatUri + "/api/oauth/profile/" + token
		resp, err := http.Get(url)
		log.Printf("GET %s : %s", url, resp.Status)
		if err != nil || resp.StatusCode != http.StatusOK {
			defer resp.Body.Close()
			c.DataFromReader(resp.StatusCode, resp.ContentLength, "application/json", resp.Body, nil)
			return
		}
		defer resp.Body.Close()
		j, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(j, &applicant)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
	}
	session.Set("nickname", applicant.NickName)
	session.Set("openid", applicant.OpenID)
	session.Set("headimgurl", applicant.HeadImgUrl)
	session.Save()
	c.JSON(http.StatusOK, applicant)
}
