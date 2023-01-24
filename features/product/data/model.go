package data

import (
	"ecommerce/features/product"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName   string
	ProductImage  string
	Description   string
	Category      string
	Qty           uint
	Price         uint
	ProductDetail string
	ImportantInfo string
	UserID        uint
}

func ToCore(data Product) product.Core {
	return product.Core{
		ID:            data.ID,
		ProductName:   data.ProductName,
		ProductImage:  data.ProductImage,
		Description:   data.Description,
		Category:      data.Category,
		Qty:           data.Qty,
		Price:         data.Price,
		ProductDetail: data.ProductDetail,
		ImportantInfo: data.ImportantInfo,
		UserID:        data.UserID,
	}
}

func CoreToData(data product.Core) Product {
	return Product{
		Model:         gorm.Model{ID: data.ID},
		ProductName:   data.ProductName,
		ProductImage:  data.ProductImage,
		Description:   data.Description,
		Category:      data.Category,
		Qty:           data.Qty,
		Price:         data.Price,
		ProductDetail: data.ProductDetail,
		ImportantInfo: data.ImportantInfo,
		UserID:        data.UserID,
	}
}
