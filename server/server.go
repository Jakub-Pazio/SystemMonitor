package main

import (

	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/", "../front-end/dist")

	tracker := CpuTracker{}
	go tracker.StartTracking()

	e.GET("/cpu", func(c echo.Context) error {
		e.Logger.Infof("call to /cpu endpoint: %v", c.ParamValues())
		name := "cpu"
		result := tracker.GetCpuPercentUsage()
		rs := &ResponseStat{name, result}
		return c.JSON(http.StatusOK, rs)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
