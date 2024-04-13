package main

import (
	"github.com/labstack/echo/v4"
	"log"
)

var app Config

func main() {
	e := echo.New()

	if err := e.Start("0.0.0.0:9876"); err != nil {
		log.Fatalln(err)
	}
}
