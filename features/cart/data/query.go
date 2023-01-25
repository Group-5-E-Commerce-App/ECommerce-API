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

<<<<<<< HEAD
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
=======
func (cq *cartQuery) Update(userID uint, cartID uint, qty int) (cart.Core, error) {
	getID := Cart{}
	err := cq.db.Where("id = ?", cartID).First(&getID).Error
	if err != nil {
		log.Println("select query error", err.Error())
		return cart.Core{}, errors.New("select query error")
	}

	prod := Product{}
	err = cq.db.Where("id = ?", getID.ProductID).First(&prod).Error
	if err != nil {
		log.Println("select query error", err.Error())
		return cart.Core{}, errors.New("select query error")
	}

	res := Cart{}
	res.Qty = qty
	res.UserID = userID
	qry := cq.db.Where("id = ?", cartID).Updates(&res)

	if qry.RowsAffected <= 0 {
		log.Println("update content query error : data not found")
		return cart.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update cart query error : ", err.Error())
	}
	return ToCore(res), nil
}

func (cq *cartQuery) Delete(userID uint, cartID uint) error {
	cart := Cart{}
	qry := cq.db.Where("id = ? and user_id = ?", cartID, userID).Delete(&cart)

	affRow := qry.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("failed to delete user content, data not found")
	}

	err := qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("delete data fail")
	}

	return nil
>>>>>>> feat : cart update and delete
}
