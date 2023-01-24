package product

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID            uint   `json:"id" form:"id"`
	ProductName   string `json:"product_name" form:"product_name"`
	ProductImage  string `json:"product_image" form:"product_image"`
	Description   string `json:"description" form:"description"`
	Category      string `json:"category" form:"category"`
	Qty           uint   `json:"qty" form:"qty"`
	Price         uint   `json:"price" form:"price"`
	ProductDetail string `json:"product_detail" form:"product_detail"`
	ImportantInfo string `json:"important_info" form:"important_info"`
	UserID        uint   `json:"user_id" form:"user_id"`
}

type ProductHandler interface {
	Add() echo.HandlerFunc
	ProductDetail() echo.HandlerFunc
	ProductList() echo.HandlerFunc
	// Update() echo.HandlerFunc
	// Delete() echo.HandlerFunc
}

type ProductService interface {
	Add(file multipart.FileHeader, token interface{}, newProduct Core) (Core, error)
	ProductDetail(productID uint) (Core, error)
	ProductList() ([]Core, error)
	// Update(token interface{}, contentID uint, updatedContent Core) (Core, error)
	// Delete(token interface{}, contentID uint) error
}

type ProductData interface {
	Add(userID uint, newProduct Core) (Core, error)
	ProductDetail(productID uint) (Core, error)
	ProductList() ([]Core, error)
	// Update(userID uint, contentID uint, updatedContent Core) (Core, error)
	// Delete(userID uint, contentID uint) error
}
