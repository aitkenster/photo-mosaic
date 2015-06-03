package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"image"
	_ "image/png"
	_ "image/gif"
	_ "image/jpeg"
	"github.com/aitkenster/photo-mosaic/edit_image"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadFile("uploads.html")
	fmt.Fprint(w, string(body))
}

func viewFileHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("image")
	if err != nil {
		fmt.Fprint(w, "Error @ 1")
		fmt.Fprint(w, err)
		return
	}

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Fprint(w, "Error @ 2")
		fmt.Fprint(w, err)
		return
	}

	edit_image.CreateMosaic(img)

	http.ServeFile(w, r, "altered_test_image.jpeg")
}

func main () {
	http.HandleFunc("/", uploadHandler)
	http.HandleFunc("/view", viewFileHandler)
	http.ListenAndServe(":8080", nil)

}
