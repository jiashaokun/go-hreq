package main

import (
	"fmt"

	"go-hreq/service"

	"github.com/robfig/cron"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		// middleware.Recover(),
		)
	c := cron.New()
	c.AddFunc("*/60 * * * * ?", service.Repre)
	c.Start()
	select {}
	fmt.Println("ccc")
}
