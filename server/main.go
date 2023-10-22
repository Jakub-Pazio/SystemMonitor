package main

import (
	"gonitor/pkg/cpu"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/", "../front-end/dist")

	tracker := cpu.CpuTracker{}
	go tracker.StartTracking()

	e.GET("/cpu", func(c echo.Context) error {
		e.Logger.Infof("call to /cpu endpoint: %v", c.ParamValues())
		name := "cpu"
		result := tracker.GetCpuPercentUsage()
		rs := &cpu.ResponseStat{Name: name, Value: result}
		return c.JSON(http.StatusOK, rs)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
