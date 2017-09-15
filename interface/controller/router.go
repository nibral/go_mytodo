package controller

import (
	"github.com/labstack/echo"
	"go_mytodo/interface/config"
)

type Router struct{}

func NewRouter(config config.AppConfig) *echo.Echo {
	router := echo.New()

	itemController := NewItemController(config.Database)
	router.POST("/items", itemController.Create)
	router.GET("/items", itemController.GetAll)
	router.GET("/items/:id", itemController.Get)
	router.PUT("/items/:id", itemController.Update)
	router.DELETE("/items/:id", itemController.Delete)

	return router
}
