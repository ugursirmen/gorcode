package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return router
}

var routes = Routes{
	Route{
		"ProductList",
		"GET",
		"/products",
		ProductList,
	},
	Route{
		"ProductDetail",
		"GET",
		"/products/{id}",
		ProductDetail,
	},
	Route{
		"ProductAdd",
		"POST",
		"/products",
		ProductAdd,
	},
	Route{
		"ProductUpdate",
		"PUT",
		"/products/{id}",
		ProductUpdate,
	},
	Route{
		"ProductRemove",
		"DELETE",
		"/products/{id}",
		ProductRemove,
	},
	Route{
		"BarcodeImage",
		"GET",
		"/barcode/{barcode}",
		GetBarcodeImage,
	},
}
