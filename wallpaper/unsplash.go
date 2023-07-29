package wallpaper

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	// "github.com/elias-gill/walldo-in-go/wallpaper"
)

const URL = "https://api.unsplash.com/"

var accessKey = os.Getenv("UNSPLASHACCESS")
var secretKey = os.Getenv("UNSPLASHSECRET")

// to parse unsplash response from json
type unsplashResponse struct {
	Url  urls  `json:"urls"`
	Link links `json:"links"`
}

type links struct {
	Download          string `json:"download"`
	Download_location string `json:"download_location"`
}

type urls struct {
	Raw       string `json:"raw"`
	Full      string `json:"full"`
	Regular   string `json:"regular"`
	Small     string `json:"small"`
	Thumbnail string `json:"thumbnail"`
}

func SetRandomImage() {
	req, err := http.NewRequest("GET", URL+"/photos/random", nil)
	req.Header.Set("Authorization", "Client-ID "+accessKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		println(accessKey)
		panic(err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(res.StatusCode)
	}

	var body unsplashResponse
	json.NewDecoder(res.Body).Decode(&body)
	// TODO: elegir una locacion con archivos temporales (probablemente al final en el config file)
	err = downloadFile(body.Link.Download, "/home/elias/random.jpg")
	if err != nil {
		panic(err)
	}
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("Authorization", "Client-ID "+accessKey)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}
