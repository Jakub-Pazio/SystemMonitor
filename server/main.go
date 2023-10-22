package main

import (
	"gonitor/pkg/cpu"
	"gonitor/pkg/temp"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/", "../front-end/dist")

	// Initialize CPU tracker
	trackerCPU := cpu.CpuTracker{}
	go trackerCPU.StartTracking()

	// Initialize Temp tracker
	trackerTemp := temp.TempTracker{}
	go trackerTemp.StartTracking()

	e.GET("/cpu", func(c echo.Context) error {
		e.Logger.Infof("call to /cpu endpoint: %v", c.ParamValues())
		name := "cpu"
		result := trackerCPU.GetCpuPercentUsage()
		rs := &cpu.ResponseStat{Name: name, Value: result}
		return c.JSON(http.StatusOK, rs)
	})
	e.GET("/temp", func(c echo.Context) error {
		e.Logger.Infof("call to /temp endpoint: %v", c.ParamValues())
		name := "temp"
		result := trackerTemp.GetAvgTemp()
		rs := &temp.ResponseStat{Name: name, Value: result}
		return c.JSON(http.StatusOK, rs)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
