package main

import (
	"image"
	"image/jpeg"
	"image/draw"
	"os"
	"fmt"
)

func main (){
	test_image, err := os.Open("test_image.jpeg")
	if err != nil {
		fmt.Print("Error @ img1")
		fmt.Println(err)
		return
	}
	defer test_image.Close()

	img, _, err  := image.Decode(test_image)
	if err != nil {
		fmt.Print("Error @ img2")
		fmt.Println(err)
		return
	}

	b := img.Bounds()
	fmt.Println(b.Max.X)

	fmt.Println(b.Max.Y)
	crop_rect := image.NewRGBA(image.Rect(0, 0, 300, 150))
	draw.Draw(crop_rect, crop_rect.Bounds(), img, image.Point{0,0}, draw.Src)
	out, err := os.Create("altered_test_image.jpeg")
	if err != nil {
		fmt.Print("Error @ img3")
		fmt.Println(err)
		return
	}

	defer out.Close()

	err = jpeg.Encode(out, crop_rect, nil)
	if err != nil {
		fmt.Print("Error @ img4")
		fmt.Println(err)
		return
	}


}

