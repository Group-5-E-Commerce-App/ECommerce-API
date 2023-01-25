package services

import (
	cart "ecommerce/features/cart"
	"ecommerce/helper"
	"log"

	// helper "ecommerce/helper"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

type cartService struct {
	qry cart.CartData
	vld *validator.Validate
}

func New(cd cart.CartData) cart.CartService {
	return &cartService{
		qry: cd,
		vld: validator.New(),
	}
}

func (cs *cartService) Update(token interface{}, cartID uint, qty int) (cart.Core, error) {
	id := helper.ExtractToken(token)

	if id <= 0 {
		return cart.Core{}, errors.New("data not found")
	}

	res, err := cs.qry.Update(uint(id), cartID, qty)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Failed to update, no new record or data not found"
		} else if strings.Contains(err.Error(), "Unauthorized") {
			msg = "Unauthorized request"
		} else {
			msg = "unable to process the data"
		}
		return cart.Core{}, errors.New(msg)
	}
	return res, nil
}

func (cs *cartService) Delete(token interface{}, cartID uint) error {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return errors.New("data not found")
	}
	err := cs.qry.Delete(uint(id), cartID)

	if err != nil {
		log.Println("delete query error", err.Error())
		return err
	}
	return nil

}

// func (cs *cartService) AddCart(newProduct cart.Core) (cart.Core, error) {

// 	err := cs.vld.Struct(newProduct)
// 	if err != nil {
// 		if _, ok := err.(*validator.InvalidValidationError); ok {
// 			log.Println(err)
// 		}
// 		return cart.Core{}, errors.New("validation error")
// 	}
// 	if err != nil {
// 		log.Println("bcrypt error ", err.Error())
// 		return cart.Core{}, errors.New("password process error")
// 	}

// 	res, err := cs.qry.AddCart(newProduct)
// 	if err != nil {
// 		msg := ""
// 		if strings.Contains(err.Error(), "duplicated") {
// 			msg = "data sudah terdaftar"
// 		} else {
// 			msg = "terdapat masalah pada server"
// 		}
// 		return cart.Core{}, errors.New(msg)
// 	}

// 	return res, nil
// }

// func (cs *cartService) Get(token interface{}) (cart.Core, error) {
// 	id := helper.ExtractToken(token)
// 	if id <= 0 {
// 		return cart.Core{}, errors.New("data tidak ditemukan")
// 	}
// 	res, err := cs.qry.Get(uint(id))
// 	if err != nil {
// 		msg := ""
// 		if strings.Contains(err.Error(), "not found") {
// 			msg = "data tidak ditemukan"
// 		} else {
// 			msg = "terdapat masalah pada server"
// 		}
// 		return cart.Core{}, errors.New(msg)
// 	}
// 	return res, nil
// }

// func (cs *cartService) Update(token interface{}, updateData cart.Core) (cart.Core, error) {
// 	id := helper.ExtractToken(token)

// 	res, err := cs.qry.Update(uint(id), updateData)

// 	if err != nil {
// 		msg := ""
// 		if strings.Contains(err.Error(), "not found") {
// 			msg = "data tidak ditemukan"
// 		} else if strings.Contains(err.Error(), "not valid") {
// 			msg = "format tidak sesuai"
// 		} else {
// 			msg = "terdapat masalah pada server"
// 		}
// 		return cart.Core{}, errors.New(msg)
// 	}

// 	return res, nil
// }

// func (cs *cartService) Delete(token interface{}) (cart.Core, error) {

// 	id := helper.ExtractToken(token)
// 	if id <= 0 {
// 		return cart.Core{}, errors.New("id user not found")
// 	}
// 	data, err := cs.qry.Delete(uint(id))
// 	if err != nil {
// 		msg := ""
// 		if strings.Contains(err.Error(), "not found") {
// 			msg = "data tidak ditemukan"
// 		} else {
// 			msg = "internal server error"
// 		}
// 		return cart.Core{}, errors.New(msg)
// 	}
// 	return data, nil

// }
