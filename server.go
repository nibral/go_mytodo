package main

import (
	"github.com/labstack/echo/middleware"

	"mytodo/interface/config"
	"mytodo/interface/controller"
)

func main() {
	router := controller.NewRouter(config.Config)

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.Logger.Fatal(router.Start(":3000"))
}
