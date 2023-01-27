package order

import "github.com/labstack/echo"

type Core struct {
	ID           uint
	TotalProduct int
	TotalPrice   int
	PaymentUrl   string
	OrderDate    string
	Status       string
	OrderName    string
	OrderCode    string
}

type OrderHandler interface {
	AddOrder() echo.HandlerFunc
}

type OrderService interface {
	AddOrder(token interface{}, paymentURL string, orderCode string) (Core, error)
}

type OrderData interface {
	AddOrder(userID uint, paymentURL string, orderCode string) (Core, error)
}
