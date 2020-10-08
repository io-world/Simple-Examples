package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	api_key := "acc_d5f7b792a3f7835"
	api_secret := "bea243448afd35d48ada582bd68f8a8d"
	image_upload_id := "i19db7c4f6f85d76abae3192815GZYTv"

	req, _ := http.NewRequest("GET", "https://api.imagga.com/v2/categories/personal_photos?image_upload_id="+image_upload_id, nil)
	req.SetBasicAuth(api_key, api_secret)

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
