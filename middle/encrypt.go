package middle

import (
	"fmt"

	"go-hreq/library"
	"go-hreq/response"

	"github.com/labstack/echo"
)

type OutputLayout struct {
	Code int `json:"code, string"`
	Msg string `json:"msg, string"`
	Data map[string]interface{} `json:"data"`
}

func Encrypt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		sign, err := library.Encrypt(c)
		if err == false {
			response.Response(c, 500, "", fmt.Sprintln("Sign was Wrong sn was %s", sign))
			return nil
		}
		return next(c)
	}
}