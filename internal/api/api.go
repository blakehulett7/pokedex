package api

import (
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func Fetch(url string) []byte {
	response, error := http.Get(url)
	if error != nil {
		fmt.Println("Fetch Error Occured")
		return nil
	}
	body, error := io.ReadAll(response.Body)
	defer response.Body.Close()
	if error != nil {
		fmt.Println("Read Error Occured")
		return nil
	}
	return body
}
