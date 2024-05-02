package delivery

import (
	"net/http"
	"repo/cmd/base"
	"repo/cmd/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	base *base.Memory
}

func (d *Delivery) NewUser(ctx echo.Context) error {
	user := &models.CreateUser{}
	err := ctx.Bind(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Были переданы не коректные параметры")
	}
	if user.Account == "" || user.Email == "" || user.Name == "" || user.Pass == "" {
		return ctx.JSON(http.StatusBadRequest, "Были переданы не коректные параметры")
	}
	return ctx.JSON(http.StatusOK, d.base.Postgres.NewUser(ctx.Request().Context(), *user))
}

func (d *Delivery) List(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, d.base.ItemsBase.List())
}

func (d *Delivery) NewItems(ctx echo.Context) error {
	item := models.CreateItem{}
	err := ctx.Bind(item)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Были переданы не коректные параметры")
	}
	if item.Name == "" || item.Price == 0 {
		return ctx.JSON(http.StatusBadRequest, "Были переданы не коректные параметры")
	}
	return ctx.JSON(http.StatusOK, d.base.ItemsBase.CreateItem(item))
}

func (d *Delivery) DeleteItems(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, "Были переданы не коректные параметры")
	}
	idNUM, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Были переданы не коректные параметры")
	}
	return ctx.JSON(http.StatusOK, d.base.ItemsBase.DeleteItem(models.DeleteItem{Id: idNUM}))
}

func New(base base.Memory) Delivery {
	return Delivery{base: &base}
}
