package main

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code    string `json:"code"`
	Name    string `json:"name"`
	Barcode string `json:"barcode"`
}

func GetAllProducts() []Product {

	var products []Product

	db.Find(&products)

	return products
}

func GetProduct(id int) Product {

	var product Product

	if err := db.Where("ID = ?", id).First(&product).Error; err != nil {
		panic(err)
	}

	return product
}

func CreateProduct(product Product) {

	product.Barcode = createRandomEan13()

	db.Create(&product)
}

func UpdateProduct(product Product, id int) {

	var oldProduct = GetProduct(id)

	oldProduct.Code = product.Code
	oldProduct.Name = product.Name

	db.Save(&oldProduct)

}

func DeleteProduct(id int) {

	var oldProduct = GetProduct(id)

	db.Delete(&oldProduct)
}
