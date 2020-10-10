package main

type Product struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Barcode string `json:"barcode"`
}

type Products []Product
