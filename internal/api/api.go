package api

import (
	//"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Fetch(url string) {
	response, error := http.Get(url)
	if error != nil {
		fmt.Println("Fetch Error Occured")
		return
	}
	body, error := io.ReadAll(response.Body)
	response.Body.Close()
	if error != nil {
		fmt.Println("Read Error Occured")
		return
	}
	fmt.Println(body)
	return
}
