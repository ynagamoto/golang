package main

import (
	"net/http"

	"github.com/labstack/echo"

	. "./calendarApi"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.GET("json", getJson)
	e.GET("schedule", getTodaySchedule)
	e.Logger.Fatal(e.Start(":8080"))
}

func getJson(c echo.Context) error {
	type testJson struct {
		Name string `json: "name"`
		Text string `json: "text"`
	}
	context := testJson{"text", "text message"}
	return c.JSON(http.StatusOK, context)
}

func getTodaySchedule(c echo.Context) error {
	text := GetSchedule()
	return c.String(http.StatusOK, text)	
}