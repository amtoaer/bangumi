package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/amtoaer/bangumi/auth"

	"github.com/mitchellh/go-homedir"
)

func dir() string {
	path, _ := homedir.Expand("~/.config/bangumi")
	return path
}

func tokenDir() string {
	return path.Join(dir(), "token.json")
}

// ParseToken 读取本地的token缓存文件
func ParseToken() (*auth.Info, error) {
	result := &auth.Info{}
	content, err := ioutil.ReadFile(tokenDir())
	if err != nil {
		os.MkdirAll(dir(), os.FileMode(0777))
		return result, err
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// WriteToken 将info保存到本地文件
func WriteToken(info *auth.Info) error {
	content, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(tokenDir(), content, os.FileMode(0777))
}
