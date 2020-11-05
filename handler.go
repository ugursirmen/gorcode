package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ProductList(w http.ResponseWriter, r *http.Request) {
	products := GetAllProducts()

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(products); err != nil {
		panic(err)
	}
}

func ProductDetail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	product := GetProduct(id)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		panic(err)
	}
}

func GetBarcodeImage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	barcode := vars["barcode"]

	barcodeImageString := createEan13Barcode(barcode)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(barcodeImageString))
}

func ProductAdd(w http.ResponseWriter, r *http.Request) {

	var product Product

	image, _, _ := r.FormFile("image")

	if image != nil {
		imageBytes, err := ioutil.ReadAll(image)
		if err != nil {
			panic(err)
		}

		defer image.Close()

		b64String := base64.StdEncoding.EncodeToString(imageBytes)
		product.Image = b64String
	}

	product.Code = r.FormValue("code")
	product.Name = r.FormValue("name")

	product.Barcode = createRandomEan13()

	CreateProduct(product)

	w.WriteHeader(http.StatusOK)
}

func ProductUpdate(w http.ResponseWriter, r *http.Request) {

	var product Product

	image, _, _ := r.FormFile("image")

	if image != nil {
		imageBytes, err := ioutil.ReadAll(image)
		if err != nil {
			panic(err)
		}

		defer image.Close()

		b64String := base64.StdEncoding.EncodeToString(imageBytes)
		product.Image = b64String
	}

	product.Code = r.FormValue("code")
	product.Name = r.FormValue("name")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	UpdateProduct(product, id)

	w.WriteHeader(http.StatusOK)
}

func ProductRemove(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	DeleteProduct(id)

	w.WriteHeader(http.StatusOK)
}
