package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func PopulateRouter(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.POST("/login", nil)    //login with an existing account
	e.POST("/register", nil) //register a new user
	e.POST("/reset", nil)    //reset password
}
