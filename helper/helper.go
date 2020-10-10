package helper

import (
	"io"
	"io/ioutil"
	"net/http"
)

// ReadByteBody 读取response内容并关闭流（返回[]byte结果）
func ReadByteBody(r *http.Response) ([]byte, error) {
	content, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	return content, err
}

// WriteSuccessHTML 为本地web服务器显示授权成功界面
func WriteSuccessHTML(w io.Writer) {
	w.Write([]byte("<h1>成功获得验证代码！</h1>"))
}
