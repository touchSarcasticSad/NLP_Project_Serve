package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)


func main() {
	fmt.Print("Hello world")
	e := echo.New()

	e.File("/", "public/index.html")


	e.Static("/public/static", "public/assets")

	e.Logger.Fatal(e.Start(":8080"))
}