package data

import (
	"ecommerce/features/order"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	TotalProduct int
	TotalPrice   int
	PaymentUrl   string
	Status       string
	OrderCode    string
	UserID       uint
	OrderItem    []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductID uint
	OrderID   uint
	Qty       int
	Total     int
}

type Cart struct {
	gorm.Model
	Qty       int
	Total     int
	UserID    uint
	ProductID uint
}

func DataToCore(data Order) order.Core {
	return order.Core{
		ID:           data.ID,
		TotalProduct: data.TotalProduct,
		TotalPrice:   data.TotalPrice,
		OrderCode:    data.OrderCode,
		Status:       data.Status,
		PaymentUrl:   data.PaymentUrl,
	}
}

func CoreToData(core order.Core) Order {
	return Order{
		Model:        gorm.Model{ID: core.ID},
		TotalProduct: core.TotalProduct,
		TotalPrice:   core.TotalPrice,
		OrderCode:    core.OrderCode,
		Status:       core.Status,
		PaymentUrl:   core.PaymentUrl,
	}
}
