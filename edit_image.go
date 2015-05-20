package main

import (
	"image"
	"image/jpeg"
	"image/color"
	"image/draw"
	"github.com/disintegration/imaging"
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

	resized_main_img := imaging.Resize(img, 300, 0, imaging.Lanczos)
	image_averages := getImageAverageColors(resized_main_img)
	mosaic, err := os.Create("altered_test_image.jpeg")
	if err != nil {
		fmt.Print("Error @ img3")
		fmt.Println(err)
		return
	}

	defer mosaic.Close()

	image_color_dictionary := processMosaicTiles()
	tile_positions := make(map[image.Point]string)
	for point, color := range image_averages {
		tile_positions[point] = findClosestColorMatch(color, image_color_dictionary)
	}
	canvas := image.NewRGBA(image.Rect(0,0, img.Bounds().Max.X*5, img.Bounds().Max.Y*5))

	createTileCanvas(tile_positions, canvas)

	err = jpeg.Encode(mosaic, canvas, nil)
	if err != nil {
		fmt.Print("Error @ img4")
		fmt.Println(err)
		return
	}
}


//takes each 10x10 pixel block and returns the average color
func getImageAverageColors(main_image *image.NRGBA) map[image.Point]color.RGBA{
	average_colors := make(map[image.Point]color.RGBA)
	size := main_image.Bounds().Size()
	for x := 0; x < size.X; x += 5 {
		for y := 0; y < size.Y; y += 5 {
			start_point := image.Pt(x, y)
			end_point := image.Pt(x+5, y+5)
			r := image.Rectangle{start_point, end_point}
			m := main_image.SubImage(r)
			average_color := averageColor(m)
			average_colors[image.Pt(x*5, y*5)] = average_color
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
		fmt.Println(err)
	}
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

func createTileCanvas(tile_positions map[image.Point]string, mosaic *image.RGBA) {
	for point, path := range tile_positions {
		tile_file, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
		}
		defer tile_file.Close()
		tile_image, _, err := image.Decode(tile_file)
		if err != nil {
			fmt.Println(err)
		}

		resized_tile := imaging.Resize(tile_image, 25, 25, imaging.Lanczos)
		r := image.Rectangle{point, point.Add(resized_tile.Bounds().Size())}
		draw.Draw(mosaic, r, resized_tile, image.Pt(0,0), draw.Src)
	}
}
