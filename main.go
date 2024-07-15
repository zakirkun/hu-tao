package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Any("/api/*", func(c echo.Context) error {

		r := c.Request().URL

		c.JSON(http.StatusOK, map[string]string{
			"path": r.Path,
		})
		return nil
	})

	e.Logger.Fatal(e.Start(":1337"))
}
