package controller

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"go_mytodo/domain"
	"go_mytodo/usecase"
	"go_mytodo/interface/config"
	"go_mytodo/interface/database"
)

type ItemController struct {
	Interactor usecase.ItemInteractor
}

// コントローラを生成
func NewItemController(dbConfig config.Database) *ItemController {
	return &ItemController{
		Interactor: usecase.ItemInteractor{
			ItemRepository: database.NewItemRepository(dbConfig),
		},
	}
}

// アイテムを追加
func (controller *ItemController) Create(c echo.Context) error {
	// 追加するアイテムの情報を取得
	item := &domain.Item{}
	if err := c.Bind(item); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// アイテムを追加
	item, err := controller.Interactor.AddItem(item)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusCreated, item)
}

// アイテムを取得
func (controller *ItemController) Get(c echo.Context) error {
	// 指定したIDのアイテムを取得
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := controller.Interactor.ItemById(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, item)
}

// 全てのアイテムを取得
func (controller *ItemController) GetAll(c echo.Context) error {
	items, err := controller.Interactor.Items()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, items)
}

// アイテムを更新
func (controller *ItemController) Update(c echo.Context) error {
	// 更新するアイテムの情報を取得
	item := new(domain.Item)
	if err := c.Bind(item); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// アイテムを更新
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := controller.Interactor.UpdateItemById(id, item)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, item)
}

// アイテムを削除
func (controller *ItemController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := controller.Interactor.DeleteItemById(id)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}
