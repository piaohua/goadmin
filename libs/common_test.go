package libs

import (
	"testing"
)

func TestUcFirst(t *testing.T) {
	str := UcFirst("test")
	t.Logf("str %s\n", str)
}

func TestMD5(t *testing.T) {
	str := MD5("test")
	t.Logf("str %s\n", str)
}

func TestCmd(t *testing.T) {
	c := "cat /dev/urandom | od -x | tr -d ' ' | head -n 1"
	out, err := ExecCmd(c)
	t.Logf("out %s, err %v\n", string(out), err)
}

func TestLoad(t *testing.T) {
	//filePath := "shop.json"
	//list := make([]Shop, 0)
	//err := Load(filePath, &list)
	//t.Logf("err %v\n", err)
	//t.Logf("list %#v\n", list)
}
