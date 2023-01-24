package product

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID            uint
	ProductName   string
	ProductImage  string
	Description   string
	Category      string
	Qty           uint
	Price         uint
	ProductDetail string
	ImportantInfo string
	UserID        uint
}

type ProductHandler interface {
	Add() echo.HandlerFunc
	// ProductDetail() echo.HandlerFunc
	// ProductList() echo.HandlerFunc
	// Update() echo.HandlerFunc
	// Delete() echo.HandlerFunc
}

type ProductService interface {
	Add(file multipart.FileHeader, token interface{}, newProduct Core) (Core, error)
	// ProductDetail(contentID uint) (interface{}, error)
	// ProductList() ([]Core, error)
	// Update(token interface{}, contentID uint, updatedContent Core) (Core, error)
	// Delete(token interface{}, contentID uint) error
}

type ProductData interface {
	Add(userID uint, newProduct Core) (Core, error)
	// ProductDetail(productID uint) (interface{}, error)
	// ProductList() ([]Core, error)
	// Update(userID uint, contentID uint, updatedContent Core) (Core, error)
	// Delete(userID uint, contentID uint) error
}
