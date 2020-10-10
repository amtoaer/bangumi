package session

import (
	"fmt"
)

// UserInfo 获取用户信息
func (a *API) UserInfo(username string) (m map[string]interface{}, err error) {
	url := fmt.Sprintf("https://api.bgm.tv/user/%s", username)
	return a.getJSON(url)
}
