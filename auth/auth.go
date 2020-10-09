package auth

import (
	"bangumi/helper"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

// OAuth 登录流程需要的信息
type OAuth struct {
	ClientID     string
	ClientSecret string
	Port         string
	HTTPClient   *http.Client
}

// OAuthOperation Oauth登录实现的操作
type OAuthOperation interface {
	GetToken() *Info
	UpdateToken(*Info)
}

// Info 需要保存的登陆信息
type Info struct {
	Token        string
	RefreshToken string
	ExpireTime   time.Time
}

var _ OAuthOperation = &OAuth{}

// GetToken 如果以前未进行过授权，则使用该方法得到token
func (o *OAuth) GetToken() (info *Info) {
	firstURL := "https://bgm.tv/oauth/authorize?" + fmt.Sprintf("client_id=%s&response_type=code&redirect_uri=localhost:%s/callback", o.ClientID, o.Port)
	log.Println("尚未授权，请在浏览器中打开 " + firstURL + " 进行授权...")
	var code string
	// 监听本地端口
	listener, err := net.Listen("tcp4", "127.0.0.1:"+o.Port)
	if err != nil {
		log.Printf("监听%s端口失败，请检查端口占用或更换端口\n", o.Port)
		return
	}
	_ = http.Serve(listener, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/callback" {
			w.WriteHeader(404)
			return
		}
		defer listener.Close()
		code = r.URL.Query().Get("code")
		log.Println("成功获得验证代码：", code)
		helper.WriteSuccessHTML(w)
	}))
	log.Println("开始获取token")
	resp, err := o.HTTPClient.PostForm("https://bgm.tv/oauth/access_token", url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {o.ClientID},
		"client_secret": {o.ClientSecret},
		"code":          {code},
		"redirect_uri":  {"localhost:" + o.Port + "/callback"},
	})
	if err != nil {
		log.Println("请求token失败")
		return
	}
	data, _ := helper.ReadBody(resp)
	container := make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &container)
	if err != nil {
		log.Println("解析请求结果失败")
		return
	}
	log.Println("成功获取token")
	seconds := container["expires_in"].(float64)
	return &Info{
		Token:        container["access_token"].(string),
		RefreshToken: container["refresh_token"].(string),
		ExpireTime:   time.Now().Add(time.Second * time.Duration(seconds)),
	}
}

// UpdateToken 为已获取过授权的应用更新token
func (o *OAuth) UpdateToken(info *Info) {
	resp, err := o.HTTPClient.PostForm("https://bgm.tv/oauth/access_token", url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {o.ClientID},
		"client_secret": {o.ClientSecret},
		"refresh_token": {info.RefreshToken},
		"redirect_uri":  {"localhost:" + o.Port + "/callback"},
	})
	if err != nil {
		log.Println("请求更新token失败")
		return
	}
	data, _ := helper.ReadBody(resp)
	container := make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &container)
	if err != nil {
		log.Println("解析请求结果失败")
		return
	}
	log.Println("成功更新token")
	seconds := container["expires_in"].(float64)
	// 将info内各项替换为新内容
	info.Token = container["access_token"].(string)
	info.RefreshToken = container["refresh_token"].(string)
	info.ExpireTime = time.Now().Add(time.Second * time.Duration(seconds))
}
