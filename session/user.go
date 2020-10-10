package session

import (
	"fmt"
	"net/url"
)

// UserInfo 获取用户信息
func (a *API) UserInfo(username string) (m map[string]interface{}, err error) {
	url := fmt.Sprintf("https://api.bgm.tv/user/%s", username)
	return a.getJSON(url)
}

// UserCollection 获取用户收藏
// withBook:是否展示书籍 simplify:是否使用简略模式
func (a *API) UserCollection(username string, withBook, simplify bool) ([]map[string]interface{}, error) {
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

// UserCollectionPreview 获取用户收藏概览
// subjectType可选：book:书籍，anime:动画，music:音乐，game:游戏，real:三次元
// maxResults 显示的条数
func (a *API) UserCollectionPreview(username, subjectType, maxResults string) ([]map[string]interface{}, error) {
	args := &url.Values{}
	args.Add("app_id", a.Info.APPID())
	if maxResults != "" {
		args.Add("max_results", maxResults)
	}
	url := fmt.Sprintf("https://api.bgm.tv/user/%s/collections/%s?", username, subjectType) + args.Encode()
	return a.getJSONSlice(url)
}

// UserCollectionStatus 获取用户所有收藏信息
func (a *API) UserCollectionStatus(username string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("https://api.bgm.tv/user/%s/collections/status?app_id=%s", username, a.Info.APPID())
	return a.getJSONSlice(url)
}

// UserProgress 获取用户收视进度（需登陆）
func (a *API) UserProgress(username, subjectID string) ([]map[string]interface{}, error) {
	var url string
	if subjectID == "" {
		url = fmt.Sprintf("https://api.bgm.tv/user/%s/progress", username)
	} else {
		url = fmt.Sprintf("https://api.bgm.tv/user/%s/progress?subject_id=%s", username, subjectID)
	}
	return a.getJSONSlice(url)
}
