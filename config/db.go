package config

import (
	cart "ecommerce/features/cart/data"
	product "ecommerce/features/product/data"
	user "ecommerce/features/user/data"
	"fmt"

	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(ac AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ac.DBUser, ac.DBPass, ac.DBHost, ac.DBPort, ac.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection error : ", err.Error())
		return nil
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(user.User{})
	db.AutoMigrate(product.Product{})
	db.AutoMigrate(cart.Cart{})
}
