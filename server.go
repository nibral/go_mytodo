package main

import (
	"github.com/labstack/echo/middleware"

	"go_mytodo/interface/config"
	"go_mytodo/interface/controller"
)

func main() {
	router := controller.NewRouter(config.Config)

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.Logger.Fatal(router.Start(":3000"))
}
