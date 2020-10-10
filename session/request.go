package session

import (
	"bangumi/helper"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (a *API) getByte(url string) (b []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	head := fmt.Sprintf("Bearer %s", a.Info.AccessToken())
	req.Header.Add("Authorization", head)
	resp, err := a.Client.Do(req)
	if err != nil {
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
	m = make(map[string]interface{})
	err = json.Unmarshal(content, &m)
	return
}

func (a *API) post(url string, content url.Values) (string, error) {
	return "", nil
}
