package main

import (
	"code.google.com/p/rsc/qr"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	//"io/ioutil"
	"log"
	"os"
)

func main() {
	s := "刘理想爱龙珑"
	code, err := qr.Encode(s, qr.H)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
	// png_img := code.PNG()
	// err = ioutil.WriteFile("fubao.png", png_img, 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	img := code.Image()

	f, err := os.Open("head.png")
	if err != nil {
		fmt.Println(err)
	}
	head_img, err := png.Decode(f)
	defer f.Close()

	bounds := img.Bounds()
	matrix := image.NewNRGBA(bounds)

	offsets := image.Pt(
		(img.Bounds().Dx()-head_img.Bounds().Dx())/2,
		(img.Bounds().Dy()-head_img.Bounds().Dy())/2)

	draw.Draw(matrix, bounds, img, image.ZP, draw.Src)
	draw.Draw(matrix, head_img.Bounds().Add(offsets), head_img, image.ZP, draw.Over)

	result_img, _ := os.Create("fubao.png")
	png.Encode(result_img, matrix)
	defer result_img.Close()
	fmt.Println("导出完毕")
}
