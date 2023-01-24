package handler

import "ecommerce/features/product"

type ProductResponse struct {
	ID            uint   `json:"id"`
	ProductName   string `json:"product_name"`
	ProductImage  string `json:"product_image"`
	Description   string `json:"description"`
	Category      string `json:"category"`
	Qty           uint   `json:"qty"`
	Price         uint   `json:"price"`
	ProductDetail string `json:"product_detail"`
	ImportantInfo string `json:"important_info"`
}

func ToResponse(data product.Core) ProductResponse {
	return ProductResponse{
		ID:            data.ID,
		ProductName:   data.ProductName,
		ProductImage:  data.ProductImage,
		Description:   data.Description,
		Category:      data.Category,
		Qty:           data.Qty,
		Price:         data.Price,
		ProductDetail: data.ProductDetail,
		ImportantInfo: data.ImportantInfo,
	}
}

func CoresToResponse(dataCore product.Core) ProductResponse {
	return ProductResponse{
		ID:            dataCore.ID,
		ProductName:   dataCore.ProductName,
		ProductImage:  dataCore.ProductImage,
		Description:   dataCore.Description,
		Category:      dataCore.Category,
		Qty:           dataCore.Qty,
		Price:         dataCore.Price,
		ProductDetail: dataCore.ProductDetail,
		ImportantInfo: dataCore.ImportantInfo,
	}
}

func ListCoreToResp(data []product.Core) []ProductResponse {
	var dataResp []ProductResponse
	for _, v := range data {
		dataResp = append(dataResp, CoresToResponse(v))
	}
	return dataResp
}
