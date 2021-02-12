package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Set your API Key here, your can get on account settings
	apiKey := "XXXXXX"
	// URL API
	url := "https://api.monday.com/v2"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"query":"{ boards (limit:5) {name id} }"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "query")
	//req.Header.Set("Query", "{ boards (limit:5) {name id} }")
	if err != nil {
		panic(err)
	}

	req.Header.Set("authorization", apiKey)
	if req == nil {
		fmt.Printf("Not value to work, review your URL and Headears: %v", req)

	} else {
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		//Print status and header
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", req.Header)
		body, _ := ioutil.ReadAll(resp.Body)

		//Close connection and test
		fmt.Println("response Body:", string(body))
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}
}
