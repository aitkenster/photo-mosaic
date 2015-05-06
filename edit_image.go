package main

import (
	"image"
	"image/jpeg"
	"image/color"
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
	getImageAverageColors(crop_rect)
	//mini_image := getMiniImage("test_image2.jpeg")
	//addMiniImages(crop_rect, mini_image)
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

//func addMiniImages(main_image *image.RGBA, mini_image *image.RGBA) {
	//var start_point image.Point
	//size := main_image.Bounds().Size()
	//for x := 0; x < size.X; x += 20 {
		//for y := 0; y < size.Y; y += 10 {
			//start_point = image.Pt(x, y)
			//min_point := image.Pt(0,0)
			//r := image.Rectangle{start_point, start_point.Add(mini_image.Bounds().Size())}
			//draw.Draw(main_image, r, mini_image, min_point, draw.Src)
		//}
	//}
//}

//takes each 10x10 pixel block and returns the average color
func getImageAverageColors(main_image *image.RGBA) {
	size := main_image.Bounds().Size()
	for x := 0; x < size.X; x += 10 {
		for y := 0; y < size.Y; y += 10 {
			start_point := image.Pt(x, y)
			end_point := image.Pt(x+10, y+10)
			fmt.Println(start_point)
			r := image.Rectangle{start_point, end_point}
			m := main_image.SubImage(r)
			average_color := averageColor(m)
			fmt.Println(average_color)
			draw.Draw(main_image, r, &image.Uniform{average_color}, image.ZP, draw.Src)
		}
	}
}

//func getMiniImage(filename string) (*image.RGBA) {
	//mini_file, err := os.Open(filename)
	//if err != nil {
		//fmt.Print("Error @ img5")
		//fmt.Println(err)
	//}
	//defer mini_file.Close()

	//mini_image, _, err  := image.Decode(mini_file)
	//if err != nil {
		//fmt.Print("Error @ img6")
		//fmt.Println(err)
	//}

////crop the mini image(this needs to be resize or something really)
	//cropped_mini := image.NewRGBA(image.Rect(0, 0, 10, 10))
	//draw.Draw(cropped_mini, cropped_mini.Bounds(), mini_image, image.Point{0,0}, draw.Src)
	//analyseColors(cropped_mini)
	//return cropped_mini
//}

func averageColor(img image.Image) (color.RGBA) {
	bounds := img.Bounds()

	minX := bounds.Min.X
	maxX := bounds.Max.X

	minY := bounds.Min.Y
	maxY := bounds.Max.Y

	var r, g, b, number_pixels int
	for x := minX; x < maxX; x ++ {
		for y := minY; y < maxY; y++ {
			pixel := img.At(x,y)
			rgbaColor := color.RGBAModel.Convert(pixel).(color.RGBA)
			r += int(rgbaColor.R)
			g += int(rgbaColor.G)
			b += int(rgbaColor.B)
			number_pixels ++
		}
	}
	averageColor := color.RGBA{
		R: uint8(r/number_pixels),
		G: uint8(g/number_pixels),
		B: uint8(b/number_pixels),
		A: 1,
	}
	fmt.Println(averageColor)
	return averageColor

}
