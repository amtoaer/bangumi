package session

import (
	"net/http"
)

// InfoOperation 在此处使用接口避免import cycle not allowed错误
type InfoOperation interface {
	NewSession() *API
	AccessToken() string
	APPID() string
}

// API 辅助进行请求的结构体
type API struct {
	Client *http.Client
	Info   InfoOperation
}
