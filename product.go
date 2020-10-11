package main

import "strconv"

type Product struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Barcode string `json:"barcode"`
}

type Products []Product

func GetAllProducts() Products {

	products := Products{}

	for i := 0; i < 100; i++ {
		products = append(products, Product{Code: "Code" + strconv.Itoa(i+1), Name: "Name " + strconv.Itoa(i+1), Barcode: "1234567890"})
	}

	return products
}

func GetProduct(productCode string) Product {

	return Product{Code: productCode, Name: "A product", Barcode: "1234567890"}
}

func CreateProduct(product Product) {

}

func UpdateProduct(product Product, code string) {

}

func DeleteProduct(code string) {

}
