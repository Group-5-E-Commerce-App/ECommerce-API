package services

import (
	"ecommerce/features/product"
	"ecommerce/helper"
	"ecommerce/mocks"
	"errors"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewProductData(t)

	inputProduct := product.Core{ProductName: "Mie Goreng", ProductImage: "indomie.jpg", Description: "Mie Goreng micinnya banyak bener", Stok: 10, Price: 3500, ImportantInfo: "Penting banget"}
	resProduct := product.Core{ID: uint(1), ProductName: "Mie Goreng", ProductImage: "https://res.cloudinary.com/dbg0177wd/image/upload/v1673881607/go-cloudinary/llhltxp55elirjkmlyld.png", Description: "Mie Goreng micinnya banyak bener", Stok: 10, Price: 3500, ImportantInfo: "Penting banget"}
	// var a multipart.FileHeader

	t.Run("success post product", func(t *testing.T) {
		postPhoto := multipart.FileHeader{
			Filename: "a",
			Size:     10,
		}
		repo.On("Add", uint(1), inputProduct).Return(resProduct, nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(postPhoto, token, inputProduct)

		assert.NotNil(t, err)
		assert.NotEqual(t, resProduct.ID, res.ID)
		assert.NotEqual(t, resProduct.ProductName, res.ProductName)
		assert.NotEqual(t, resProduct.Description, res.Description)
		repo.AssertExpectations(t)
	})

	t.Run("invalid JWT", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(0)
		err := srv.Delete(token, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
	})

	t.Run("open file error", func(t *testing.T) {
		postPhoto := multipart.FileHeader{
			Filename: "a",
			Size:     10,
		}
		repo.On("Add", uint(1), inputProduct).Return(product.Core{}, errors.New("open file error")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(postPhoto, token, inputProduct)

		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "file")
		repo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		postPhoto := multipart.FileHeader{
			Filename: "a",
			Size:     10,
		}
		srv := New(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(postPhoto, pToken, inputProduct)
		assert.NotNil(t, err)
		assert.Equal(t, res.UserID, uint(0))
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

func TestProducDetail(t *testing.T) {
	repo := mocks.NewProductData(t)
	resProduct := product.Core{ID: uint(1), ProductName: "Mie Goreng", Description: "Mie Goreng micinnya banyak bener", Price: 3500}
	t.Run("success show product detail", func(t *testing.T) {
		repo.On("ProductDetail", uint(1)).Return(resProduct, nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.ProductDetail(uint(1))
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("problem in server", func(t *testing.T) {
		repo.On("ProductDetail", uint(1)).Return(product.Core{}, errors.New("server error")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.ProductDetail(uint(1))
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)

	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("ProductDetail", uint(1)).Return(product.Core{}, errors.New("not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.ProductDetail(uint(1))
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)

	})

}

func TestProducList(t *testing.T) {
	repo := mocks.NewProductData(t)
	resProduct := []product.Core{{ID: uint(1), ProductName: "Mie Goreng", Description: "Mie Goreng micinnya banyak bener", Price: 3500}}
	t.Run("success show product detail", func(t *testing.T) {
		repo.On("ProductList").Return(resProduct, nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.ProductList()
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("problem in server", func(t *testing.T) {
		repo.On("ProductList").Return([]product.Core{}, errors.New("server error")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.ProductList()
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)

	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("ProductList").Return([]product.Core{}, errors.New("not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.ProductList()
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)

	})

}

func TestUpdate(t *testing.T) {
	repo := mocks.NewProductData(t)
	filePath := filepath.Join("..", "..", "..", "ERD_Group5.jpg")

	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	inputProduct := product.Core{ID: uint(1), ProductName: "Mie Goreng", ProductImage: imageTrueCnv.Filename, Description: "Mie Goreng micinnya banyak bener", Price: 3500}
	resProduct := product.Core{ID: uint(1), ProductName: "Mie Goreng", ProductImage: imageTrueCnv.Filename, Description: "Mie Goreng micinnya banyak bener", Price: 3500}

	t.Run("Success update data", func(t *testing.T) {
		repo.On("Update", uint(1), uint(1), inputProduct).Return(resProduct, nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(*imageTrueCnv, pToken, 1, inputProduct)
		assert.NotNil(t, err)
		assert.NotEqual(t, inputProduct.ProductName, res.ProductName)
		repo.AssertExpectations(t)
	})

	t.Run("invalid JWT", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(0)
		err := srv.Delete(token, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
	})

	t.Run("open file error", func(t *testing.T) {
		repo.On("EditProduct", uint(1), uint(1), inputProduct).Return(product.Core{}, errors.New("file error")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(*imageTrueCnv, pToken, 1, inputProduct)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "file")
		repo.AssertExpectations(t)
	})

}
