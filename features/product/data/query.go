package data

import (
	"ecommerce/features/product"
	"log"

	"gorm.io/gorm"
)

type productData struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.ProductData {
	return &productData{
		db: db,
	}
}

func (pd *productData) Add(userID uint, newProduct product.Core) (product.Core, error) {
	cnv := CoreToData(newProduct)
	cnv.UserID = uint(userID)
	err := pd.db.Create(&cnv).Error
	if err != nil {
		return product.Core{}, err
	}
	newProduct.ID = cnv.ID

	return newProduct, nil
}

func (pd *productData) ProductDetail(productID uint) (product.Core, error) {
	res := Product{}
	if err := pd.db.Where("id = ?", productID).Find(&res).Error; err != nil {
		log.Println("get product details by id query error : ", err.Error())
		return product.Core{}, err
	}

	return ToCore(res), nil
}
