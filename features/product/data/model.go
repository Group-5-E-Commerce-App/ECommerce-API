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

func (dataModel *Product) ModelsToCore() product.Core {
	return product.Core{
		ID:            dataModel.ID,
		ProductName:   dataModel.ProductName,
		ProductImage:  dataModel.ProductImage,
		Description:   dataModel.Description,
		Category:      dataModel.Category,
		Qty:           dataModel.Qty,
		Price:         dataModel.Price,
		ProductDetail: dataModel.ProductDetail,
		ImportantInfo: dataModel.ImportantInfo,
		UserID:        dataModel.UserID,
	}
}

func ListToCore(data []Product) []product.Core {
	var dataCore []product.Core
	for _, v := range data {
		dataCore = append(dataCore, v.ModelsToCore())
	}
	return dataCore
}
