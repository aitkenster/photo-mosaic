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

//crop the main image
	crop_rect := image.NewRGBA(image.Rect(0, 0, 300, 150))
	draw.Draw(crop_rect, crop_rect.Bounds(), img, image.Point{0,0}, draw.Src)

	mini_image := getMiniImage("test_image2.jpeg")
	addMiniImages(crop_rect, mini_image)
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

func addMiniImages(main_image *image.RGBA, mini_image *image.RGBA) {
	var start_point image.Point
	size := main_image.Bounds().Size()
	for x := 0; x < size.X; x += 20 {
		for y := 0; y < size.Y; y += 10 {
			fmt.Println(x)
			fmt.Println(y)
			start_point = image.Pt(x, y)
			min_point := image.Pt(0,0)
			r := image.Rectangle{start_point, start_point.Add(mini_image.Bounds().Size())}
			draw.Draw(main_image, r, mini_image, min_point, draw.Src)
		}
	}
}

func getMiniImage(filename string) (*image.RGBA) {
	mini_file, err := os.Open(filename)
	if err != nil {
		fmt.Print("Error @ img5")
		fmt.Println(err)
	}
	defer mini_file.Close()

	mini_image, _, err  := image.Decode(mini_file)
	if err != nil {
		fmt.Print("Error @ img6")
		fmt.Println(err)
	}

//crop the mini image(this needs to be resize or something really)
	cropped_mini := image.NewRGBA(image.Rect(0, 0, 10, 10))
	draw.Draw(cropped_mini, cropped_mini.Bounds(), mini_image, image.Point{0,0}, draw.Src)
	return cropped_mini
}
