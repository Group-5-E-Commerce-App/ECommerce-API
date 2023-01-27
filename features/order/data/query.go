package data

import (
	"ecommerce/features/order"

	"gorm.io/gorm"
)

type orderQry struct {
	db *gorm.DB
}

func New(db *gorm.DB) order.OrderData {
	return &orderQry{
		db: db,
	}
}

func (oq *orderQry) AddOrder(userID uint, paymentURL string, orderCode string) (order.Core, error) {

}
