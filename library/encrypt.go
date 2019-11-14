package library

import (
	"crypto/md5"
	"fmt"
	"net/url"

	"go-hreq/config"

	"github.com/labstack/echo"
)

func Encrypt (c echo.Context) (string, bool) {
	params, err := c.FormParams()

	if err != nil {
		return "", false
	}
	sn := params.Get("sn")
	params.Del("sn")
	srcKey := config.ApiCinfig[params.Get("src")]
	sign := makeSign(params, srcKey)

	if sign != sn {
		return sign, false
	}
	return sign, true
}

func makeSign(v url.Values, k string) string {
	if len(v) == 0 {
		return ""
	}
	s := v.Encode()
	s = fmt.Sprintf("%s%s", s, k)

	ms := md5.Sum([]byte(s))
	sign := fmt.Sprintf("%x", ms)

	return sign
}
