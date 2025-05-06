package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_itemShopService "github.com/saksit339/isekai-shop-api-tutorial/pkg/itemShop/service"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewitemShopControllerImpl(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(ctx echo.Context) error {
	itemModelList, err := c.itemShopService.Listing()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, itemModelList)
}
