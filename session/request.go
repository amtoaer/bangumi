package session

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/amtoaer/bangumi/helper"
)

// Request 基本的get/post操作
type Request interface {
	getByte(string) ([]byte, error)
	get(string) (string, error)
	getJSON(string) (map[string]interface{}, error)
	getJSONSlice(string) ([]map[string]interface{}, error)
	post(string, url.Values) (string, error)
}

var _ Request = &API{}

func (a *API) getByte(url string) (b []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	if a.Info.AccessToken() != "" {
		head := fmt.Sprintf("Bearer %s", a.Info.AccessToken())
		req.Header.Add("Authorization", head)
	}
	resp, err := a.Client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return
	}
	return helper.ReadByteBody(resp)
}

func (a *API) get(url string) (string, error) {
	result, err := a.getByte(url)
	return string(result), err
}

func (a *API) getJSON(url string) (m map[string]interface{}, err error) {
	content, err := a.getByte(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &m)
	return
}

func (a *API) getJSONSlice(url string) (m []map[string]interface{}, err error) {
	content, err := a.getByte(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &m)
	return
}

// TODO 完成post基本功能
func (a *API) post(url string, content url.Values) (string, error) {
	return "", nil
}
