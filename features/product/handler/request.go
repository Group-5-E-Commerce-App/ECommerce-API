package handler

import "ecommerce/features/product"

type AddProductRequest struct {
	ProductName   string `json:"product_name" form:"product_name"`
	ProductImage  string `json:"product_image" form:"product_image"`
	Description   string `json:"description" form:"description"`
	Stok          uint   `json:"stok" form:"stok"`
	Price         uint   `json:"price" form:"price"`
	ImportantInfo string `json:"important_info" form:"important_info"`
}

type UpdateProductRequest struct {
	ProductName   string `json:"product_name" form:"product_name"`
	ProductImage  string `json:"product_image" form:"product_image"`
	Description   string `json:"description" form:"description"`
	Stok          uint   `json:"stok" form:"stok"`
	Price         uint   `json:"price" form:"price"`
	ImportantInfo string `json:"important_info" form:"important_info"`
}

func ToCore(data interface{}) *product.Core {
	res := product.Core{}

	switch data.(type) {
	case AddProductRequest:
		cnv := data.(AddProductRequest)
		res.ProductName = cnv.ProductName
		res.ProductImage = cnv.ProductImage
		res.Description = cnv.Description
		res.Stok = cnv.Stok
		res.Price = cnv.Price
		res.ImportantInfo = cnv.ImportantInfo
	case UpdateProductRequest:
		cnv := data.(UpdateProductRequest)
		res.ProductName = cnv.ProductName
		res.ProductImage = cnv.ProductImage
		res.Description = cnv.Description
		res.Stok = cnv.Stok
		res.Price = cnv.Price
		res.ImportantInfo = cnv.ImportantInfo
	default:
		return nil
	}

	return &res
}
