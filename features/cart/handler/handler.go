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

		res, err := ch.srv.Update(token, uint(cartID), input.QtyProduct)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success edit product quanity in cart",
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

// func (ch *cartHandler) AddCart() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		input := AddCartReq{}
// 		if err := c.Bind(&input); err != nil {
// 			return c.JSON(http.StatusBadRequest, "format inputan salah")
// 		}

// 		idProduct, err := strconv.Atoi(c.Param("idProduct"))
// 		if err != nil {
// 			return c.JSON(helper.PrintErrorResponse(err.Error()))
// 		}
// 		_, err = ch.srv.AddCart(*ToCore(idProduct))
// 		if err != nil {
// 			return c.JSON(helper.PrintErrorResponse(err.Error()))
// 		}
// 		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "berhasil menambahkan"))
// 	}
// }

// func (ch *cartHandler) Get() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		id := helper.ExtractToken(c)

// 		res, err := ch.srv.Get(uint(id))
// 		if err != nil {
// 			return c.JSON(helper.PrintErrorResponse(err.Error()))
// 		}
// 		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil lihat profil", ToResponse(res, "sukses")))
// 	}
// }

// func (ch *cartHandler) Update() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		id, err := strconv.Atoi(c.Param("id"))

// 		input := UpdateRequest{}
// 		input.ID = uint(id)
// 		if err := c.Bind(&input); err != nil {
// 			return c.JSON(http.StatusBadRequest, "format inputan salah")
// 		}

// 		id2 := helper.ExtractToken(c)
// 		input.idUser = uint(id2)
// 		dataCore := *ToCore(input)

// 		res, err := ch.srv.Update(id, dataCore)

// 		if err != nil {
// 			return c.JSON(helper.PrintErrorResponse(err.Error()))
// 		}
// 		dataResp := ToResponse(res)
// 		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "berhasil mengubah data", dataResp))
// 	}
// }

// func (ch *cartHandler) Delete() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		id, err := strconv.Atoi(c.Param("id"))
// 		_, err = ch.srv.Delete(uint(id))
// 		if err != nil {
// 			return c.JSON(helper.PrintErrorResponse(err.Error()))
// 		}
// 		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil hapus"))
// 	}
// }
