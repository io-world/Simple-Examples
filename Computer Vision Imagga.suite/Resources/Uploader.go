package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	path, _ := os.Getwd()
	path = "/Users/randyhesse/Documents/Simple_Examples/Computer Vision Imagga.suite/Images/header.png"

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("image", fi.Name())
	if err != nil {
		fmt.Println(err)
		return
	}

	part.Write(fileContents)

	client := &http.Client{}
	api_key := "acc_d5f7b792a3f7835"
	api_secret := "bea243448afd35d48ada582bd68f8a8d"

	req, _ := http.NewRequest("POST", "https://api.imagga.com/v2/uploads", body)
	req.SetBasicAuth(api_key, api_secret)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
}
