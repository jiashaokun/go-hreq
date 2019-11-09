package response

import (
	"github.com/labstack/echo"
)

type OutputLayout struct {
	Code int `json:"code, string"`
	Msg string `json:"msg, string"`
	Data interface{} `json:"data"`
}

type (
	codeCustom interface {
		Code() int
	}

	jsonOutput interface {
		Json() ([]byte, error)
	}
)

func Response (c echo.Context, code int, v interface{}, msg string) error {
	statusCode := code

	outData := &OutputLayout{
		Code: statusCode,
		Msg:  CodeMsg[code],
	}
	if msg != "" {
		outData.Msg = msg
	}
	outData.Data = v


	return c.JSON(statusCode, outData)
}