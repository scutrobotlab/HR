package mid

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/scutrobotlab/HR/conf"
	"github.com/scutrobotlab/HR/dal/model"
	"github.com/scutrobotlab/HR/dal/query"
	"gorm.io/gorm"
)

const testAccessToken = "Test Access Token"

type Group struct {
	Name string `json:"name"`
}

type Admin struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Avatar string  `json:"avatar"`
	Groups []Group `json:"groups"`
}

// 管理员登录验证中间件
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.DefaultMany(c, "admin")
		access_token := session.Get("access_token")
		if access_token == nil {
			log.Println("access_token unfound")
			c.AbortWithStatus(401)
			return
		}
		token := fmt.Sprint(access_token)
		admin, err := getAdmin(token)
		if err != nil {
			log.Println("get admin err: ", err.Error())
			c.AbortWithStatus(401)
			return
		}
		c.Set("admin", admin)

		c.Next()
	}
}

// 管理员登录
func AdminLogin(c *gin.Context) {
	session := sessions.DefaultMany(c, "admin")
	code, ok := c.Params.Get("code")
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}
	access_token, err := getAccessToken(code)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	session.Set("access_token", access_token)
	session.Save()
	admin, err := getAdmin(access_token)
	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}
	c.JSON(http.StatusOK, admin)
}

func getAccessToken(code string) (string, error) {
	if gin.Mode() != gin.ReleaseMode {
		return testAccessToken, nil
	}
	u, err := url.Parse(conf.AdminUri)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, "/api/oauth/token")
	q := u.Query()
	q.Add("grant_type", "authorization_code")
	q.Add("client_id", conf.OAuthClientID)
	q.Add("client_secret", conf.OAuthSecret)
	q.Add("redirect_uri", conf.RedirectUri)
	q.Add("code", code)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	log.Printf("GET %s : %s", u.String(), resp.Status)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", errors.Wrap(err, resp.Status)
	}
	defer resp.Body.Close()
	j, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	r := struct {
		AccessToken string `json:"access_token"`
	}{}
	err = json.Unmarshal(j, &r)
	if err != nil {
		return "", err
	}
	return r.AccessToken, nil
}

func getAdmin(access_token string) (*Admin, error) {
	if gin.Mode() != gin.ReleaseMode && access_token == testAccessToken {
		return &Admin{
			ID:     1,
			Name:   "测试员",
			Groups: []Group{{"机械"}},
		}, nil
	}
	u, err := url.Parse(conf.AdminUri)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "/api/oauth/userinfo")
	q := u.Query()
	q.Add("access_token", access_token)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	log.Printf("GET %s : %s", u.String(), resp.Status)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, resp.Status)
	}
	defer resp.Body.Close()
	j, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	a := Admin{}
	err = json.Unmarshal(j, &a)
	if err != nil {
		return nil, err
	}
	return &a, updateAdmin(a)
}

func updateAdmin(a Admin) error {
	q := query.Q.Admin
	do := q.WithContext(context.Background())
	m, err := do.Where(q.ID.Eq(a.ID)).Take()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		m := model.Admin{ID: a.ID, Name: a.Name}
		return do.Create(&m)
	}
	if err != nil {
		return err
	}
	if m.Name == a.Name {
		return nil
	}
	m.Name = a.Name
	_, err = do.Update(q.Name, a.Name)
	return err
}
