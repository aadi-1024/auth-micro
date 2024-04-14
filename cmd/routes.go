package main

import (
	"github.com/aadi-1024/auth-micro/pkg/handlers"
	"github.com/labstack/echo/v4"
	"net/http"
)

func PopulateRouter(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.POST("/login", handlers.LoginHandler(app.Db, app.Jwt))  //login with an existing account
	e.POST("/register", handlers.RegistrationHandler(app.Db)) //register a new user
	e.POST("/reset", handlers.ResetPasswordHandler(app.Db))   //reset password
}
