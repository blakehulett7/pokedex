package api

import (
	"fmt"
	"net/http"
)

func fetch(url string) {
	response, error := http.Get(url)
	if error != nil {
		fmt.Println("Fetch Error Occured")
		return
	}
	fmt.Println(response)
}
