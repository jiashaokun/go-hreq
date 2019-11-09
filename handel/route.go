package handel

import (
	"hreq/middle"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Route() {
	e := echo.New()

	e.Use(
		middle.Encrypt,
		middleware.Logger(),
	)

	e.POST("/req", addReq)

	e.Start(":1323")
}
