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

func (cq *cartQuery) AddCart(userID uint, productID uint, newCart cart.Core) (cart.Core, error) {
	produk := Product{}
	log.Println(productID)
	err := cq.db.Where("id=?", productID).First(&produk).Error
	if err != nil {
		log.Println("query error", err.Error())
		return cart.Core{}, errors.New("server error")
	}
	cnv := CoreToData(newCart)
	cnv.ProductID = produk.ID
	cnv.UserID = userID
	err = cq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return cart.Core{}, errors.New("server error")
	}
	result := ToCore(cnv)
	return result, nil
}

func (cq *cartQuery) Update(userID uint, cartID uint, updatedCart cart.Core) (cart.Core, error) {
	cnv := CoreToData(updatedCart)
	cnv.UserID = userID
	cnv.ID = cartID

	qry := cq.db.Where("id = ? AND user_id = ?", cartID, userID).Updates(cnv)
	if qry.RowsAffected <= 0 {
		log.Println("update cart query error : data not found")
		return cart.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update cart query error : ", err.Error())
	}
	return ToCore(cnv), nil
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
}

func (cq *cartQuery) Get(userID uint) ([]cart.Core, error) {
	allCart := []cart.Core{}
	err := cq.db.Raw("SELECT carts.user_id, carts.product_id, products.price, carts.qty FROM carts JOIN users ON carts.user_id = users.id JOIN products ON carts.product_id = products.id").Scan(&allCart).Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return []cart.Core{}, err
	}
	return allCart, nil
}
