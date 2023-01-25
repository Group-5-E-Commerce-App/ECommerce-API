package handler

import cart "ecommerce/features/cart"

type AddCartReq struct {
	IdProduct  uint `json:"id_product" form:"id_product"`
	IdUser     uint `json:"id_user" form:"id_user"`
	QtyProduct int  `json:"product_qty" form:"product_qty"`
}

type UpdateFormat struct {
	ID         uint `json:"id" form:"id"`
	IdProduct  uint `json:"id_product" form:"id_product"`
	IdUser     uint `json:"id_user" form:"id_user"`
	QtyProduct int  `json:"product_qty" form:"product_qty"`
}

type GetId struct {
	id         uint `param:"id"`
	IdProduct  uint `json:"id_product" form:"id_product"`
	IdUser     uint `json:"id_user" form:"id_user"`
	QtyProduct int  `json:"product_qty" form:"product_qty"`
}

func ToCore(data interface{}) *cart.Core {
	res := cart.Core{}

	switch data.(type) {
	case AddCartReq:
		cnv := data.(AddCartReq)
		res.IdProduct = cnv.IdProduct
		res.IdUser = cnv.IdUser
		res.QtyProduct = cnv.QtyProduct
	case GetId:
		cnv := data.(GetId)
		res.IdProduct = cnv.IdProduct
		res.IdUser = cnv.IdUser
		res.QtyProduct = cnv.QtyProduct
	case UpdateFormat:
		cnv := data.(UpdateFormat)
		res.IdProduct = cnv.IdProduct
		res.IdUser = cnv.IdUser
		res.QtyProduct = cnv.QtyProduct
	default:
		return nil
	}

	return &res
}
