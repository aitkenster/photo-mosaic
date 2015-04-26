package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
	"image/png"
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

	img, err := png.Decode(file)
	if err != nil {
		fmt.Fprint(w, "Error @ 2")
		fmt.Fprint(w, err)
		return
	}

	out, err := os.Create("./tmp/uploadedfile.jpeg")
	if err != nil {
		fmt.Fprint(w, "Error @ 3")
		fmt.Fprint(w, err)

	}
	defer out.Close()

	err = jpeg.Encode(out, img, nil)
	if err != nil {
		fmt.Fprint(w, "Error @ 4")
		fmt.Fprint(w, err)
		return
	}



	fmt.Fprint(w, "Image uploaded successfully")
}

func main () {
	http.HandleFunc("/", uploadHandler)
	http.HandleFunc("/view", viewFileHandler)
	http.ListenAndServe(":8080", nil)

}
