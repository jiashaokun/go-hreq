package main

import (
	"go-hreq/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/robfig/cron"
)

func main() {
	e := echo.New()
	e.Use(
		middleware.Logger(),
	)
	// service.Repre()

	c := cron.New()
	c.AddFunc("*/60 * * * * ?", service.Repre)
	c.Start()
	select {}
}
