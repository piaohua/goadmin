package libs

import (
	"time"

	"github.com/astaxie/beego/httplib"
)

// CurlGet ...
func CurlGet(url string, params map[string]string) (string, error) {
	req := httplib.Get(url).SetTimeout(2*time.Second, 3*time.Second)
	for k, v := range params {
		req.Param(k, v)
	}
	str, err := req.String()
	if err != nil {
		return "", err
	}
	return str, nil
}

// CurlPost ...
func CurlPost(url string, params map[string]string) (string, error) {
	req := httplib.Post(url).SetTimeout(2*time.Second, 3*time.Second)
	for k, v := range params {
		req.Param(k, v)
	}
	str, err := req.String()
	if err != nil {
		return "", err
	}
	return str, nil
}
