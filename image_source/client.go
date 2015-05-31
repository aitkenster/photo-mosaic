package image_source

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func getFlickrRecentPhotos() []string {
	resp, err := http.Get("https://api.flickr.com/services/feeds/photos_public.gne?format=json&nojsoncallback=1")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	type WrappedImage struct {
		Image string `json:"m"`
	}

	type ImageLinks struct {
		Media WrappedImage `json:"media"`
	}

	var data struct {
		Items []ImageLinks `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
		return nil
	}
	var imageList []string
	for _, entry := range data.Items {
		imageList = append(imageList, entry.Media.Image)
	}
	fmt.Println(imageList)
	return imageList
}

