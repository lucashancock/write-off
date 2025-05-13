package main

import (
	"encoding/json"
	"fmt"
	"journal/api"
)

func main() {
	api.ApiServer()
}

func PrintJSON(obj interface{}) {
	bytes, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(bytes))
}
