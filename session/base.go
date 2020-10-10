package session

import (
	"fmt"
	"net/url"
)

func (a *API) userCollection(username, cat string, responseGroup string) (m []map[string]interface{}, err error) {
	args := &url.Values{}
	args.Add("cat", cat)
	if responseGroup != "" {
		args.Add("responseGroup", responseGroup)
	}
	url := fmt.Sprintf("https://api.bgm.tv/user/%s/collection?", username) + args.Encode()
	return a.getJSONSlice(url)
}
