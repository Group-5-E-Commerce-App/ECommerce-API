package handler

import (
	"ecommerce/dtos"
	"ecommerce/features/product"
	"ecommerce/helper"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type productHandle struct {
	srv product.ProductService
}

func New(ps product.ProductService) product.ProductHandler {
	return &productHandle{
		srv: ps,
	}
}

func (ph *productHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("product")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Select a file to upload"},
			})
		}

		input := AddProductRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid input")
		}

		cnv := ToCore(input)

		res, err := ph.srv.Add(*file, c.Get("user"), *cnv)
		if err != nil {
			log.Println("error post content : ", err.Error())
			return c.JSON(http.StatusInternalServerError, "unable to process the data")
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "success post content", res))
	}
}
