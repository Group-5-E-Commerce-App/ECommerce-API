package cart

import "github.com/labstack/echo/v4"

type Core struct {
	ID               uint
	IdProduct        uint
	IdUser           uint
	NamaProduct      string
	QtyProduct       int
	DetailProduct    string
	DeskripsiProduct string
	InfoPenting      string
	Price            int
	ProductPicture   string
}

type CartHandler interface {
	AddCart() echo.HandlerFunc
	// Get() echo.HandlerFunc
	// Delete() echo.HandlerFunc
	// Update() echo.HandlerFunc
}

type CartService interface {
	AddCart(newProduct Core) (Core, error)
	// Get(token interface{}) (Core, error)
	// Update(token interface{}, updateData Core) (Core, error)
	// Delete(token interface{}) (Core, error)
}

type CartData interface {
	AddCart(newCart Core) (Core, error)
	// Get(id uint) (Core, error)
	// Update(id uint, updateData Core) (Core, error)
	// Delete(id uint) (Core, error)
}
