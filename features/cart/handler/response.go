package handler

import (
	cart "ecommerce/features/cart"
)

type CartResponse struct {
	ID              uint   `json:"id"`
	ProductID       uint   `json:"product_id"`
	UserID          uint   `json:"user_id"`
	ProductName     string `json:"product_name"`
	QuantityProduct int    `json:"qty_product"`
}

type AddCartResponse struct {
	IdProduct  uint `json:"id_product"`
	IdUser     uint `json:"id_user"`
	QtyProduct int  `json:"product_qty"`
}

func ToResponse(data cart.Core) CartResponse {
	return CartResponse{
		ID:              data.ID,
		ProductID:       data.IdProduct,
		UserID:          data.IdUser,
		ProductName:     data.NamaProduct,
		QuantityProduct: data.QtyProduct,
	}
}

func CoresToResponse(dataCore cart.Core) CartResponse {
	return CartResponse{
		ID:              dataCore.ID,
		ProductID:       dataCore.IdProduct,
		UserID:          dataCore.IdUser,
		ProductName:     dataCore.NamaProduct,
		QuantityProduct: dataCore.QtyProduct,
	}
}
func ListCoreToResp(data []cart.Core) []CartResponse {
	var dataResp []CartResponse
	for _, v := range data {
		dataResp = append(dataResp, CoresToResponse(v))
	}
	return dataResp
}

// func fromCoreList(dataCore []cart.Core) []AddCartResponse {
// 	var dataResponse []AddCartResponse

// 	for _, v := range dataCore {
// 		dataResponse = append(dataResponse, ToResponse(v))
// 	}
// 	return dataResponse
// }
