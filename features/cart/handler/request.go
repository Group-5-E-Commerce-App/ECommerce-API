package handler

import cart "ecommerce/features/cart"

type AddCartReq struct {
	IdProduct        uint   `json:"id_product" form:"id_product"`
	IdUser           uint   `json:"id_user" form:"id_user"`
	NamaProduct      string `json:"product_name" form:"product_name"`
	QtyProduct       int    `json:"product_qty" form:"product_qty"`
	DetailProduct    string `json:"product_detail" form:"product_detail"`
	DeskripsiProduct string `json:"product_deskripsi" form:"product_deskripsi"`
	InfoPenting      string `json:"info_penting" form:"info_penting"`
	Price            int    `json:"price" form:"price"`
	ProductPicture   string `json:"product_picture" form:"product_picture"`
}

type UpdateFormat struct {
	ID               uint   `json:"id" form:"id"`
	IdProduct        uint   `json:"id_product" form:"id_product"`
	IdUser           uint   `json:"id_user" form:"id_user"`
	NamaProduct      string `json:"product_name" form:"product_name"`
	QtyProduct       int    `json:"product_qty" form:"product_qty"`
	DetailProduct    string `json:"product_detail" form:"product_detail"`
	DeskripsiProduct string `json:"product_deskripsi" form:"product_deskripsi"`
	InfoPenting      string `json:"info_penting" form:"info_penting"`
	Price            int    `json:"price" form:"price"`
	ProductPicture   string `json:"product_picture" form:"product_picture"`
}

type GetId struct {
	id               uint   `param:"id"`
	IdProduct        uint   `json:"id_product" form:"id_product"`
	IdUser           uint   `json:"id_user" form:"id_user"`
	NamaProduct      string `json:"product_name" form:"product_name"`
	QtyProduct       int    `json:"product_qty" form:"product_qty"`
	DetailProduct    string `json:"product_detail" form:"product_detail"`
	DeskripsiProduct string `json:"product_deskripsi" form:"product_deskripsi"`
	InfoPenting      string `json:"info_penting" form:"info_penting"`
	Price            int    `json:"price" form:"price"`
	ProductPicture   string `json:"product_picture" form:"product_picture"`
}

func ToCore(data interface{}) *cart.Core {
	res := cart.Core{}

	switch data.(type) {
	case AddCartReq:
		cnv := data.(AddCartReq)
		res.IdProduct = cnv.IdProduct
		res.IdUser = cnv.IdUser
		res.NamaProduct = cnv.NamaProduct
		res.QtyProduct = cnv.QtyProduct
		res.DetailProduct = cnv.DetailProduct
		res.DeskripsiProduct = cnv.DeskripsiProduct
		res.InfoPenting = res.InfoPenting
		res.Price = cnv.Price
		res.ProductPicture = cnv.ProductPicture
	case GetId:
		cnv := data.(GetId)
		res.IdProduct = cnv.IdProduct
		res.IdUser = cnv.IdUser
		res.NamaProduct = cnv.NamaProduct
		res.QtyProduct = cnv.QtyProduct
		res.DetailProduct = cnv.DetailProduct
		res.DeskripsiProduct = cnv.DeskripsiProduct
		res.InfoPenting = res.InfoPenting
		res.Price = cnv.Price
		res.ProductPicture = cnv.ProductPicture
	case UpdateFormat:
		cnv := data.(UpdateFormat)
		res.IdProduct = cnv.IdProduct
		res.IdUser = cnv.IdUser
		res.NamaProduct = cnv.NamaProduct
		res.QtyProduct = cnv.QtyProduct
		res.DetailProduct = cnv.DetailProduct
		res.DeskripsiProduct = cnv.DeskripsiProduct
		res.InfoPenting = res.InfoPenting
		res.Price = cnv.Price
		res.ProductPicture = cnv.ProductPicture
	default:
		return nil
	}

	return &res
}
