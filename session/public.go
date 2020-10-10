package session

import (
	"fmt"
)

// UserInfo 获取用户信息
func (a *API) UserInfo(username string) (m map[string]interface{}, err error) {
	url := fmt.Sprintf("https://api.bgm.tv/user/%s", username)
	return a.getJSON(url)
}

// UserCollection 获取用户收藏
func (a *API) UserCollection(username string, withBook, simplify bool) (m []map[string]interface{}, err error) {
	var cat, responseGroup string
	if withBook {
		cat = "all_watching"
	} else {
		cat = "watching"
	}
	if simplify {
		responseGroup = "small"
	} else {
		responseGroup = "medium"
	}
	return a.userCollection(username, cat, responseGroup)
}
