package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Board struct {
	Data struct {
		Boards []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"boards"`
	} `json:"data"`
	AccountID int `json:"account_id"`
}

type Columns struct {
	Data struct {
		Boards []struct {
			Owner struct {
				ID int `json:"id"`
			} `json:"owner"`
			Columns []struct {
				ID    string `json:"id"`
				Title string `json:"title"`
				Type  string `json:"type"`
			} `json:"columns"`
		} `json:"boards"`
	} `json:"data"`
	AccountID int `json:"account_id"`
}

type Items struct {
	Data struct {
		Boards []struct {
			Item []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"items"`
		} `json:"boards"`
	} `json:"data"`
	AccountID int `json:"account_id"`
}


func GetAction(v []byte, option int){
	// export your API Token with ENV KEYMONDAY
	apiKey := os.Getenv("KEYMONDAY")
	if apiKey == "" {
		fmt.Printf("KEYMONDAY is empty \n You need export your API Token env var\n")
		return
	}

	url := "https://api.monday.com/v2"
	method := "GET"
	payload := strings.NewReader(string(v))

	req, err := http.NewRequest(method, url, payload)
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
		if option == 1 {
			var columns Columns
			json.Unmarshal( body, &columns )
			fmt.Printf("\nColumns:\n")
			for _, value := range(columns.Data.Boards) {
				fmt.Printf("\n %v", value )
			}
		} else if option == 2 {
			var board Board
			json.Unmarshal( body, &board )
			fmt.Printf("\nBoards:\n")
			for _, value := range(board.Data.Boards) {
				fmt.Println(value)
			}
		} else if option == 3 {
			var items Items
			json.Unmarshal( body, &items )
			fmt.Printf("\nItems:\n")
			//fmt.Println(items)
			for i, value := range(items.Data.Boards ) {
				fmt.Printf("%T", value)
				fmt.Println(i)
			}
		}

		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}
}



// Help menu
func menu(){
	fmt.Println(`-----------------
GoLang Monday CLI:

--help
--boards -> Show your boards Name, ID
--column -> Show your Board Column Name, ID, Type
--itens  -> Show 15 itens of board
`)
}

func main() {
	// Count the arguments
	if len(os.Args) < 2 {
		fmt.Println("Error: Faltou parametro")
		menu()
		os.Exit(0)
	} else {
		// Parameter valid options
		switch os.Args[1] {
		case "--boards":
			b := []byte(`{"query":"{ boards (limit:5) {name id} }"}`)
			GetAction(b, 2)
		case "--column":
			b := []byte(`{"query":" query { boards (ids: 1050856417) { owner { id } columns { id title type } } }"}`)
			GetAction(b, 1)
		case "--items":
			b := []byte(`{"query":"query { boards (ids: 1050856417) { items (limit: 50) {id name} }}"}`)
			GetAction(b, 3)
		default:
			menu()
		}
	}
}
