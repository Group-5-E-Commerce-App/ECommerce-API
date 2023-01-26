package handler

import (
	cart "ecommerce/features/cart"
	helper "ecommerce/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type cartHandler struct {
	srv cart.CartService
}

func New(srv cart.CartService) cart.CartHandler {
	return &cartHandler{
		srv: srv,
	}
}

func (ch *cartHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		paramID := c.Param("id")
		cartID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "Invalid input")
		}

		input := UpdateFormat{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadGateway, "invalid input")
		}

		res, err := ch.srv.Update(token, uint(cartID), *ToCore(input))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success edit product quantity in cart",
		})

	}
}

func (ch *cartHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		paramID := c.Param("id")
		cartID, err := strconv.Atoi(paramID)

		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "Invalid input")
		}

		err = ch.srv.Delete(token, uint(cartID))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusAccepted, "success delete cart")
	}
}

func (ch *cartHandler) AddCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := AddCartReq{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		res, err := ch.srv.AddCart(token, input.IdProduct, *ToCore(input))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "berhasil menambahkan", res))
	}
}
