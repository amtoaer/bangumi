package session

import (
	"net/http"
	"net/url"
)

// InfoOperation 在此处使用接口避免import cycle not allowed错误
type InfoOperation interface {
	NewSession() *API
	AccessToken() string
}

// API 辅助进行请求的结构体
type API struct {
	Client *http.Client
	Info   InfoOperation
}

// APIOperation API需要实现的操作
type APIOperation interface {
	get(string) (string, error)
	post(string, url.Values) (string, error)
	UserInfo(string) (map[string]interface{}, error)
}

var _ APIOperation = &API{}
