package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/",  func(c echo.Context) error {
		text_ := c.FormValue("text_content")
		return c.String(http.StatusOK, "Text: " + text_)
	})
	e.Logger.Fatal(e.Start(":3001"))
}