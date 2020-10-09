package helper

import (
	"io"
	"io/ioutil"
	"net/http"
)

// ReadBody 读取response内容并关闭流
func ReadBody(r *http.Response) (string, error) {
	content, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// WriteSuccessHTML 为本地web服务器显示授权成功界面
func WriteSuccessHTML(w io.Writer) {
	w.Write([]byte("<h1>成功获得验证代码！</h1>"))
}
