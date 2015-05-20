package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func main() {
	getFlickrRecentPhotos()
}

type PhotoStream struct {
	title string `json:"name"`
	link string `json:"link"`
	description string `json:"description"`
	modified string `json:"modified"`
	generator string `json:"generator"`
	items []map[string]PhotoInfo
}

type PhotoInfo struct {
	title string `json:"title"`
	link string `json:"link"`
	media map[string]string `json:"media"`
	date_taken string `json:"data_taken"`
	description string `json:"description"`
	published string `json:"published"`
	author string `json:"author"`
	author_id string `json:"author_id"`
	tags string `json:"tags"`

}

func getFlickrRecentPhotos() []string {
	resp, err := http.Get("https://api.flickr.com/services/feeds/photos_public.gne?format=json&nojsoncallback=1")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var flickrData map[string][]PhotoStream
	err = json.Unmarshal(body, &flickrData)
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range flickrData["items"] {
		fmt.Println(item)
	}
	return nil
}



