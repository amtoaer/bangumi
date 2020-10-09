package auth

import (
	"bangumi/helper"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// OAuth 登录流程需要的信息
type OAuth struct {
	clientID     string
	clientSecret string
	port         string
	httpClient   *http.Client
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
	firstURL := "https://bgm.tv/oauth/authorize?" + fmt.Sprintf("client_id=%s&response_type=code&redirect_url=localhost:%s/callback", o.clientID, o.port)
	log.Println("尚未授权，请在浏览器中打开" + firstURL + "进行授权...")
	var code string
	// 监听本地端口
	listener, err := net.Listen("tcp4", "127.0.0.1:"+o.port)
	if err != nil {
		log.Printf("监听%s端口失败，请检查端口占用或更换端口\n", o.port)
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
	resp, err := o.httpClient.PostForm("https://bgm.tv/oauth/access_token", url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {o.clientID},
		"client_secret": {o.clientSecret},
		"code":          {code},
		"redirect_uri":  {"localhost:" + o.port + "/callback"},
	})
	if err != nil {
		log.Println("请求token失败")
		return
	}
	data, _ := helper.ReadBody(resp)
	values, err := url.ParseQuery(data)
	if err != nil {
		log.Println("解析请求结果失败")
		return
	}
	log.Println("成功获取token")
	seconds, _ := strconv.Atoi(values.Get("expires_in"))
	return &Info{
		Token:        values.Get("access_token"),
		RefreshToken: values.Get("refresh_token"),
		ExpireTime:   time.Now().Add(time.Second * time.Duration(seconds)),
	}
}

// UpdateToken 为已获取过授权的应用更新token
func (o *OAuth) UpdateToken(info *Info) {
	resp, err := o.httpClient.PostForm("https://bgm.tv/oauth/access_token", url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {o.clientID},
		"client_secret": {o.clientSecret},
		"refresh_token": {info.RefreshToken},
		"redirect_uri":  {"localhost:" + o.port + "/callback"},
	})
	if err != nil {
		log.Println("请求更新token失败")
		return
	}
	data, _ := helper.ReadBody(resp)
	values, err := url.ParseQuery(data)
	if err != nil {
		log.Println("解析请求结果失败")
		return
	}
	log.Println("成功更新token")
	seconds, _ := strconv.Atoi(values.Get("expires_in"))
	// 将info内各项替换为新内容
	info.Token = values.Get("access_token")
	info.RefreshToken = values.Get("refresh_token")
	info.ExpireTime = time.Now().Add(time.Second * time.Duration(seconds))
}
