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

func (cq *cartQuery) AddCart(newCart cart.Core) (cart.Core, error) {
	var cnv cart.Core
	var compare Product
	if err := cq.db.Where("id_user = ? AND id = ?", cnv.IdUser, cnv.IdProduct).First(&compare).Error; err == nil {
		log.Print(errors.New("cannot buy own product"))
		return domain.Core{}, errors.New("cannot buy own product")
	}

	if err := rq.db.Where("id_product=? AND id_user=?", cnv.IdProduct, cnv.IdUser).First(&cnv).Error; err == nil {
		cnv.ProductQty += 1
		if err := rq.db.Model(&Cart{}).Where("id_product = ?", cnv.IdProduct).Update("product_qty", cnv.ProductQty).Error; err != nil {
			log.Print(errors.New("error udpdate quantity"))
			return domain.Core{}, err

		}
	} else {
		if err := rq.db.Where("id = ? AND product_qty>=?", cnv.IdProduct, cnv.ProductQty).First(&compare).Error; err != nil {
			log.Print(errors.New("stock product tidak cukup"))
			return domain.Core{}, errors.New("stock product tidak cukup")
		}

		if err := rq.db.Select("id_product", "id_user", "product_qty").Create(&cnv).Error; err != nil {
			return domain.Core{}, err
		}

	}

	// selesai dari DB
	newCart = ToDomain(cnv)
	return newCart, nil
}
