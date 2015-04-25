package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadFile("uploads.html")
	fmt.Fprint(w, string(body))
}

func main () {
	http.HandleFunc("/", uploadHandler)
	http.ListenAndServe(":8080", nil)
}
