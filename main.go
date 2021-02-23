package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Board struct {
	Data struct {
		Boards []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"boards"`
	} `json:"data"`
	AccountID int `json:"account_id"`
}

func GetBoards(){
	// export your API Token with ENV KEYMONDAY
	apiKey := os.Getenv("KEYMONDAY")
	// URL API
	url := "https://api.monday.com/v2"

	var jsonStr = []byte(`{"query":"{ boards (limit:5) {name id} }"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}

	if req.Header.Set("authorization", apiKey); req == nil {
		fmt.Printf("Not value to work, review your URL and Headears: %v", req)
	} else {
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		//Print status and header
		fmt.Println("Status:", resp.Status)
		//fmt.Println("response Headers:", req.Header)
		body, _ := ioutil.ReadAll(resp.Body)

		// create a data container
		var board Board
		// unmarshal `data`
		json.Unmarshal( body, &board )
		fmt.Printf("\nBoards:\n")
		for _, value := range(board.Data.Boards) {
			fmt.Println(value)
		}

		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}
}

func main() {
	GetBoards()
}
