package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-hreq/service"
)

func main() {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		// middleware.Recover(),
		)
	service.Repre()
	//c := cron.New()
	//c.AddFunc("*/60 * * * * ?", service.Repre)
	//c.Start()
	//select {}
	//fmt.Println("ccc")

}
