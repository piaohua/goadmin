package libs

import (
	"testing"
)

func TestCurlGet(t *testing.T) {
	url := "http://localhost:8080/debug/vars"
	params := make(map[string]string)
	out, err := CurlGet(url, params)
	t.Logf("out %s, err %v\n", out, err)
}

func TestCurlPost(t *testing.T) {
}
