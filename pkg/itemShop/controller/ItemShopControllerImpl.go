package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saksit339/isekai-shop-api-tutorial/pkg/custom"
	_itemShopModel "github.com/saksit339/isekai-shop-api-tutorial/pkg/itemShop/model"
	_itemShopService "github.com/saksit339/isekai-shop-api-tutorial/pkg/itemShop/service"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewitemShopControllerImpl(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(ctx echo.Context) error {

	filter := _itemShopModel.ItemFilter{}
	if err := ctx.Bind(filter); err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	itemModelList, err := c.itemShopService.Listing()
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, itemModelList)
}
