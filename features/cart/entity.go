package cart

import "github.com/labstack/echo/v4"

type Core struct {
	ID          uint   `json:"id" form:"id"`
	IdProduct   uint   `json:"product_id" form:"product_id"`
	IdUser      uint   `json:"user_id" form:"user_id"`
	NamaProduct string `json:"product_name" form:"product_name"`
	QtyProduct  int    `json:"qty_product" form:"qty_product"`
	Price       int    `json:"price" form:"price"`
}

type CartHandler interface {
	AddCart() echo.HandlerFunc
	Get() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type CartService interface {
	AddCart(token interface{}, productId uint, newCart Core) (Core, error)
	Get(token interface{}) ([]Core, error)
	Update(token interface{}, cartID uint, updatedCart Core) (Core, error)
	Delete(token interface{}, cartID uint) error
}

type CartData interface {
	AddCart(productId uint, UserID uint, newCart Core) (Core, error)
	Get(userID uint) ([]Core, error)
	Update(userID uint, cartID uint, updatedCart Core) (Core, error)
	Delete(userID uint, cartID uint) error
}
