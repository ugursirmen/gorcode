package main

import (
	"strconv"
	"time"
)

type Product struct {
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	Barcode    string    `json:"barcode"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

type Products []Product

func GetAllProducts() Products {

	products := Products{}

	for i := 0; i < 100; i++ {
		products = append(products, Product{Code: "Code" + strconv.Itoa(i+1), Name: "Name " + strconv.Itoa(i+1), Barcode: "1234567890", ModifiedAt: time.Now()})
	}

	return products
}

func GetProduct(barcode string) Product {

	return Product{Code: barcode, Name: "A product", Barcode: "1234567890"}
}

func CreateProduct(product Product) {

}

func UpdateProduct(product Product, barcode string) {

}

func DeleteProduct(barcode string) {

}
