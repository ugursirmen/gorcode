package main

import (
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("gorcode.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	barcodeCreator()

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func barcodeCreator() {
	code, _ := code128.Encode("1234567890")

	scaledCode, _ := barcode.Scale(code, 300, 100)

	file, _ := os.Create("barcode.png")
	defer file.Close()

	png.Encode(file, scaledCode)
}
