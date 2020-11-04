package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math/rand"
	"strconv"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/ean"
	"github.com/disintegration/imaging"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

func createEan13Barcode(productCode string) string {

	fmt.Println("Generating Datamatrix barcode for : ", productCode)

	code, err := ean.Encode(productCode)

	if err != nil {
		fmt.Printf("String %s cannot be encoded\n", productCode)
	}

	scaledCode, err := barcode.Scale(code, 200, 200)

	if err != nil {
		fmt.Println("EAN scaling error : ", err)
	}

	draw2d.SetFontFolder(".")

	img := image.NewRGBA(image.Rect(0, 0, 250, 50))

	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	gc := draw2dimg.NewGraphicContext(img)

	gc.FillStroke()

	gc.SetFontData(draw2d.FontData{"BebasNeue", draw2d.FontFamilyMono, draw2d.FontStyleBold | draw2d.FontStyleItalic})
	gc.SetFillColor(image.Black)

	gc.SetFontSize(30)
	gc.FillStringAt(productCode, 8, 40)

	newImg := imaging.New(200, 200, color.NRGBA{255, 255, 255, 255})

	newImg = imaging.Paste(newImg, scaledCode, image.Pt(0, 0))

	newImg = imaging.Paste(newImg, img, image.Pt(0, 150))

	buf := new(bytes.Buffer)

	err = jpeg.Encode(buf, newImg, &jpeg.Options{35})

	if err != nil {
		fmt.Println(err)
	}

	imageBit := buf.Bytes()

	b64String := base64.StdEncoding.EncodeToString([]byte(imageBit))

	return b64String
}

func createRandomEan13() string {
	rand.Seed(time.Now().UnixNano())

	min := 100000000000

	max := 999999999999

	rnd := rand.Intn(max-min) + min

	return strconv.Itoa(rnd)
}
