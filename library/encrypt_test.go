package library

import (
	"fmt"
	"net/url"
	"testing"
)

// 测试加密函数：map[info:[id=1&name=2&sn=3] method:[post] num:[3] src:[1] url:[http://baidu.com]]
func TestMakeSign(t *testing.T) {

	req := url.Values{
		"info":{"id=1&name=2&sn=3"},
		"method": {"post"},
		"num": {"3"},
		"src": {"1"},
		"url": {"http://baidu.com"},
	}

	sn := makeSign(req, "test1")
	if sn != "0300cbb894bc409c1cf195716ed2434c" {
		t.Fatal(fmt.Sprintf("makeSign sn was wrong wat 0300cbb894bc409c1cf195716ed2434c1 now %s", sn))
	}
}
