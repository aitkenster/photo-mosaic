package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
	"image"
//import file formats for the image package to decode
	_ "image/png"
	_ "image/gif"
	"image/jpeg"
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

	out, err := os.Create("./public/tmp/uploadedfile.jpg")
	if err != nil {
		fmt.Fprint(w, "Error @ 3")
		fmt.Fprint(w, err)

	}

	err = jpeg.Encode(out, img, nil)
	if err != nil {
		fmt.Fprint(w, "Error @ 4")
		fmt.Fprint(w, err)
		return
	}

	defer out.Close()

	http.ServeFile(w, r, "./public/tmp/uploadedfile.jpg")

}

func main () {
	http.HandleFunc("/", uploadHandler)
	http.HandleFunc("/view", viewFileHandler)
	http.ListenAndServe(":8080", nil)

}
