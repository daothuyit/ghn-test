package main

import (
	"ghn-test/configs"
	"ghn-test/routes"
    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()
	configs.ConnectDB()
	routes.ProductRoute(e)
    e.Logger.Fatal(e.Start(":" + configs.GetEnv("PORT")))
}