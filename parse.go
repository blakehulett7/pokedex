package main

import (
	"encoding/json"
	"fmt"
	"internal/api"
)

func ReadLocations(rawData []byte) {
	locations := api.Location{}
	json.Unmarshal(rawData, &locations)
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
}
