package services

import (
	"ecommerce/features/product"
	"ecommerce/helper"
	"ecommerce/mocks"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewProductData(t)

	var a multipart.FileHeader

	t.Run("success post product", func(t *testing.T) {
		inputProduct := product.Core{ProductName: "Mie Goreng", ProductImage: "indomie.jpg", Description: "Mie Goreng micinnya banyak bener", Qty: 10, Price: 3500, ImportantInfo: "Penting banget"}
		resProduct := product.Core{ID: uint(1), ProductName: "Mie Goreng", ProductImage: "https://res.cloudinary.com/dbg0177wd/image/upload/v1673881607/go-cloudinary/llhltxp55elirjkmlyld.png", Description: "Mie Goreng micinnya banyak bener", Qty: 10, Price: 3500, ImportantInfo: "Penting banget"}
		repo.On("Add", uint(1), inputProduct).Return(resProduct, nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(a, token, inputProduct)

		assert.Nil(t, err)
		assert.Equal(t, resProduct.ID, res.ID)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewProductData(t)

	t.Run("success delete product", func(t *testing.T) {
		repo.On("Delete", uint(1), uint(1)).Return(nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, 1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("invalid JWT", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(0)
		err := srv.Delete(token, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("Delete", uint(2), uint(2)).Return(errors.New("data not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(2)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, 2)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		repo.AssertExpectations(t)
	})
}
