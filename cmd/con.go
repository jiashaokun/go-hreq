package main

import (
	"time"

	"go-hreq/config"
	"go-hreq/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(
		middleware.Logger(),
	)
	// N 秒执行一次
	ticker := time.NewTicker(config.CronTimeSecond * time.Second)

	for {
		select {
		case <- ticker.C:
			service.Repre()
			continue
		}
	}
}
