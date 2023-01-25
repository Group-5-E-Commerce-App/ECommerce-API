package data

import (
	cart "ecommerce/features/cart"
	"errors"
	"log"

	"gorm.io/gorm"
)

type cartQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) cart.CartData {
	return &cartQuery{
		db: db,
	}
}

// func (cq *cartQuery) AddCart(newCart cart.Core) (cart.Core, error) {
// 	var compare Product
// 	if err := cq.db.Where("id_user = ? AND id = ?", newCart.IdUser, newCart.IdProduct).First(&compare).Error; err == nil {
// 		log.Print(errors.New("cannot buy own product"))
// 		return cart.Core{}, errors.New("cannot buy own product")
// 	}

// 	if err := cq.db.Where("id_product=?", newCart.IdProduct).First(&newCart).Error; err == nil {
// 		newCart.QtyProduct += 1
// 		if err := cq.db.Model(&Cart{}).Where("id_product = ?", newCart.IdProduct).Update("product_qty", newCart.QtyProduct).Error; err != nil {
// 			log.Print(errors.New("error udpdate quantity"))
// 			return cart.Core{}, err

// 		}
// 	} else {
// 		if err := cq.db.Where("id = ? AND product_qty>=?", newCart.IdProduct, newCart.QtyProduct).First(&compare).Error; err != nil {
// 			log.Print(errors.New("stock product tidak cukup"))
// 			return cart.Core{}, errors.New("stock product tidak cukup")
// 		}

// 		if err := cq.db.Select("id_product", "id_user", "product_qty").Create(&newCart).Error; err != nil {
// 			return cart.Core{}, err
// 		}

// 	}

//		// selesai dari DB
//		return ToCore(newCart), nil
//	}
func (cq *cartQuery) AddCart(newCart cart.Core) (cart.Core, error) {
	produk := Product{}
	err := cq.db.Where("id=?", newCart.IdProduct).First(&produk).Error
	if err != nil {
		log.Println("query error", err.Error())
		return cart.Core{}, errors.New("server error")
	}
	cnv := CoreToData(newCart)
	cnv.ProductID = produk.ID
	cnv.UserID = newCart.IdUser
	cnv.Qty = 1
	err = cq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return cart.Core{}, errors.New("server error")
	}
	result := ToCore(cnv)
	return result, nil
}
