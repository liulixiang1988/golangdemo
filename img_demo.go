package main

import (
	"fmt"

	"github.com/disintegration/gift"
	"image"
	"image/jpeg"
	"os"
)

func main() {
	f, err := os.Open("origin.jpg")
	if err != nil {
		fmt.Println(err)
	}

	img, err := jpeg.Decode(f)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	g := gift.New(
		gift.Resize(800, 0, gift.LanczosResampling),
		gift.UnsharpMask(1.0, 1.0, 0.0))

	dst := image.NewNRGBA(g.Bounds(img.Bounds()))

	g.Draw(dst, img)
	result_img, _ := os.Create("img_demo.jpg")
	jpeg.Encode(result_img, dst, &jpeg.Options{100})
	defer result_img.Close()
	fmt.Println("导出完毕")

}
