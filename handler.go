package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PageVariables struct {
	PageTitle string
}

func ProductList(w http.ResponseWriter, r *http.Request) {
	products := Products{}

	for i := 0; i < 100; i++ {
		products = append(products, Product{Code: "Code" + strconv.Itoa(i+1), Name: "Name " + strconv.Itoa(i+1), Barcode: "1234567890"})
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(products); err != nil {
		panic(err)
	}
}

func ProductDetail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	code := vars["code"]
	fmt.Fprintln(w, "Product code:", code)
}

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	var product Product
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &product); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	//todo create new product and return it
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		panic(err)
	}
}
