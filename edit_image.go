package main

import (
	"image"
	"image/jpeg"
	"image/color"
	"image/draw"
	"os"
	"path/filepath"
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
	image_averages := getImageAverageColors(crop_rect)
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
	sample_color := image_averages[70]
	image_color_dictionary := processMosaicTiles()
	fmt.Println(findClosestColorMatch(sample_color, image_color_dictionary))

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
func getImageAverageColors(main_image *image.RGBA) []color.RGBA{
	var average_colors []color.RGBA
	size := main_image.Bounds().Size()
	for x := 0; x < size.X; x += 10 {
		for y := 0; y < size.Y; y += 10 {
			start_point := image.Pt(x, y)
			end_point := image.Pt(x+10, y+10)
			r := image.Rectangle{start_point, end_point}
			m := main_image.SubImage(r)
			average_color := averageColor(m)
			average_colors = append(average_colors, average_color)
			draw.Draw(main_image, r, &image.Uniform{average_color}, image.ZP, draw.Src)
		}
	}
	return average_colors
}

func processMosaicTiles()(map[color.RGBA] string) {
	image_color_dictionary := make(map[color.RGBA]string)

	overallColorAvg := func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			tile_image, err := os.Open("./" + path)
			if err != nil {
				fmt.Print("Error @ mos img1")
				fmt.Println(err)
				return nil
			}
			defer tile_image.Close()

			img, _, err  := image.Decode(tile_image)
			if err != nil {
				fmt.Print("Error @ mos img2")
				fmt.Println(err)
				return nil
			}
			image_color_dictionary[averageColor(img)] = "./" + path
		}
		return nil
	}

	path := "./public/mosaic_tiles"
	err := filepath.Walk(path, overallColorAvg)


	if err != nil {
		fmt.Println("yo")
		fmt.Println(err)
	}
	fmt.Println(image_color_dictionary)
	return image_color_dictionary
}

func findClosestColorMatch(average_color color.RGBA, image_color_dictionary map[color.RGBA]string) string {
	var tile_palette color.Palette
	for color, _ := range image_color_dictionary {
		tile_palette = append(tile_palette, color)
	}
	closet_color_match := tile_palette.Convert(average_color)

	return image_color_dictionary[color.RGBAModel.Convert(closet_color_match).(color.RGBA)]
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
	return averageColor

}

func cropToSquare(img *image.RGBA) (*image.RGBA) {
	var cropped_img *image.RGBA
	var side_length int
	if img.Bounds().Max.X > img.Bounds().Max.Y {
		side_length = img.Bounds().Max.Y
	} else {
		side_length = img.Bounds().Max.X
	}

	cropped_img = image.NewRGBA(image.Rect(0, 0, side_length, side_length))
	draw.Draw(cropped_img, cropped_img.Bounds(), img, image.Point{0,0}, draw.Src)

	return cropped_img
}

//func getMatchingImage(average_color color.RGBA, image_color_averages map[string]color.RGBA) (*image.RGBA) {

	//return matching_image
//}
