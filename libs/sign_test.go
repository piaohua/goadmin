package libs

import (
	"testing"
)

func TestSign(t *testing.T) {
	params := make(map[string]string)
	params["i"] = "i"
	params["p"] = "p"
	params["a"] = "a"
	params["sign"] = "7782F13A9FC822CBC130D6DB3D7BD6ED"
	key := "secret"
	sign := Sign(params, key)
	if params["sign"] != sign {
		t.Errorf("sign %s\n", sign)
	}
}
