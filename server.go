package main

import (
	"mytodo/interface/config"
	"mytodo/interface/controller"
)

func main() {
	router := controller.NewRouter(config.Config)
	router.Logger.Fatal(router.Start(":3000"))
}
