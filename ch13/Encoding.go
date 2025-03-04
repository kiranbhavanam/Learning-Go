package ch13

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
}

func Encoding() {

	cat := Animal{"cat", "beluga"}
	jsonData, _ := json.Marshal(cat)
	fmt.Println(string(jsonData))
}