package main

import (
	"encoding/json"
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

func ProductAdd(w http.ResponseWriter, r *http.Request) {

	var product Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	product.Barcode = product.Code + "1234567890"

	CreateProduct(product)

	w.WriteHeader(http.StatusOK)
}

func ProductUpdate(w http.ResponseWriter, r *http.Request) {
	var product Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

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
