package login

import (
	"bangumi/auth"
	"bangumi/config"
	"log"
	"net/http"
	"time"
)

// Login 登陆bangumi帐号并获取token信息
func Login(clientID, clientSecret string, port string) *auth.Info {
	info, err := config.ParseToken()
	// 读取、解析本地缓存未出现错误且token未过时
	if err == nil && !time.Now().After(info.ExpireTime) {
		return info
	}
	loginService := &auth.OAuth{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Port:         port,
		HTTPClient: &http.Client{
			Timeout: 8 * time.Second,
		},
	}
	// 读取文件或解析内容失败则重新授权
	if err != nil {
		info = loginService.GetToken()

	} else {
		// 否则是token失效，只需要进行更新
		loginService.UpdateToken(info)
	}
	err = config.WriteToken(info)
	if err != nil {
		log.Fatal("保存到文件失败")
	}
	return info
}
