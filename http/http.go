package http_test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Data `json:"results"`
}

type Data struct {
	Name string
	Url  string
}

func Http() {

	response, err := http.Get("https://pokeapi.co/api/v2/ability/?limit=20&offset=20")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var ResponseObject Response
	json.Unmarshal(responseData, &ResponseObject)

	for _, res := range ResponseObject.Results {
		fmt.Println(res.Name)
	}
}
